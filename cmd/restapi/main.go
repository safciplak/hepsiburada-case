package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/hepsiburada/hepsiburada-command/internal/campaigns"
	"github.com/hepsiburada/hepsiburada-command/internal/commands"
	"github.com/hepsiburada/hepsiburada-command/internal/orders"
	"github.com/hepsiburada/hepsiburada-command/internal/products"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/gorilla/mux"
	"github.com/hepsiburada/hepsiburada-command/app"
	"github.com/rs/cors"
	"github.com/subosito/gotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultPORT = "8080"

func init() {
	gotenv.Load()
	app.Init(os.Getenv("APP_ENV"), os.Getenv("APP_URL"))
}

func main() {

	db := initGORM()

	autoMigrate(db)

	// services
	var productService *products.Service
	{
		repo := products.NewMysqlRepo(db)
		productService = &products.Service{Products: repo}
	}

	var campaignService *campaigns.Service
	{
		repo := campaigns.NewMysqlRepo(db)
		campaignService = &campaigns.Service{Campaigns: repo}
	}

	var orderService *orders.Service
	{
		repo := orders.NewMysqlRepo(db)
		orderService = &orders.Service{Orders: repo}
	}

	var commandsService *commands.Service

	commandsService = &commands.Service{
		Commands:  commandsService,
		Products:  productService,
		Campaigns: campaignService,
		Orders:    orderService,
	}

	commandsService.ProcessEachCommand(context.TODO())

	r := mux.NewRouter()

	r.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	// start the http server
	log.Printf("Listening on port: %s", getADDR())

	sentryHandler := sentryhttp.New(sentryhttp.Options{Repanic: true})

	handler := cors.AllowAll().Handler(sentryHandler.Handle(r))

	if err := http.ListenAndServe(getADDR(), handler); err != nil {
		log.Fatal(err)
	}
}

func getADDR() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return ":" + defaultPORT
}

func mustGetEnv(key string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	panic("env required: " + key)
}

func mysqlDSN(dsn string) string {
	u, err := url.Parse(dsn)
	if err != nil {
		log.Fatalf("mysql dsn parse error: %s", err)
	}
	userPass := u.User.Username()
	if pass, ok := u.User.Password(); ok {
		userPass += ":" + pass
	}
	return fmt.Sprintf("%s@tcp(%s)%s?charset=utf8&parseTime=True&loc=Local", userPass, u.Host, u.Path)
}

func initGORM() *gorm.DB {
	logmode := logger.Info
	if app.ENV == app.PROD {
		logmode = logger.Silent
	}

	db, err := gorm.Open(mysql.Open(mysqlDSN(mustGetEnv("MYSQL_DSN"))), &gorm.Config{
		Logger: logger.Default.LogMode(logmode),
	})
	if err != nil {
		log.Fatalf("gorm: %s", err)
	}

	return db
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&products.Product{}, campaigns.Campaign{}, orders.Order{})
}
