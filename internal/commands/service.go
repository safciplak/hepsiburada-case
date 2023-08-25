package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hepsiburada/hepsiburada-command/app"
	"github.com/hepsiburada/hepsiburada-command/internal/campaigns"
	"github.com/hepsiburada/hepsiburada-command/internal/orders"
	"github.com/hepsiburada/hepsiburada-command/internal/products"
	"github.com/hepsiburada/hepsiburada-command/internal/uniqid"
)

const (
	CREATE_PRODUCT    = "create_product"
	CREATE_CAMPAIGN   = "create_campaign"
	GET_CAMPAGIN_INFO = "get_campaign_info"
	GET_PRODUCT_INFO  = "get_product_info"
	INCREASE_TIME     = "increase_time"
	CREATE_ORDER      = "create_order"
)

var (
	ErrorNotfound = app.BusinessError("command not found")
)

type commandsService interface {
	GetCommands(ctx context.Context) ([]Command, error)
}

type productService interface {
	CreateProduct(ctx context.Context, c *products.Product) error
	GetProduct(ctx context.Context, productName string) (*products.Product, error)
}

type campaignService interface {
	CreateCampaign(ctx context.Context, c *campaigns.Campaign) error
	GetCampaign(ctx context.Context, campaignName string) (*campaigns.Campaign, error)
}

type ordersService interface {
	CreateOrder(ctx context.Context, o *orders.Order) error
	GetOrder(ctx context.Context, orderName string) (*orders.Order, error)
}

type Service struct {
	Commands  commandsService
	Products  productService
	Campaigns campaignService
	Orders    ordersService
}

type Command struct {
	Command string `json:"command"`
}

type Commands struct {
	Commands []Command `json:"commands"`
}

func (s *Service) GetCommands(ctx context.Context) ([]Command, error) {
	// Open our jsonFile

	jsonFile, err := os.Open("commands.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened commands.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var commands Commands

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &commands)

	return commands.Commands, nil
}

func (s *Service) ProcessEachCommand(ctx context.Context) {
	commandList, err := s.Commands.GetCommands(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, _ := range commandList {
		var command = commandList[k].Command
		if strings.HasPrefix(command, CREATE_PRODUCT) {
			// create product
			s.createProduct(ctx, command)
		}

		if strings.HasPrefix(command, CREATE_CAMPAIGN) {
			// create campaign
			s.createCampaign(ctx, command)
		}

		if strings.HasPrefix(command, GET_CAMPAGIN_INFO) {
			// get campaign info
			s.getCampaignInfo(ctx, command)
		}

		if strings.HasPrefix(command, GET_PRODUCT_INFO) {
			// get product info
			s.getProductInfo(ctx, command)
		}

		if strings.HasPrefix(command, INCREASE_TIME) {
			// increase time
		}

		if strings.HasPrefix(command, CREATE_ORDER) {
			// create order
			now := time.Now()
			res := strings.Split(command, " ")

			productCode := res[1]
			quantity, _ := strconv.ParseInt(res[2], 10, 64)

			var o orders.Order
			o.ID = uniqid.Generate()
			o.ProductCode = productCode
			o.Quantity = quantity
			o.CreatedAt = &now
			o.UpdatedAt = &now

			err := s.Orders.CreateOrder(ctx, &o)
			if err != nil {
				log.Println("create order error")
			}
		}
	}
}

func (s *Service) getProductInfo(ctx context.Context, command string) {
	res := strings.Split(command, " ")
	name := res[1]

	p, err := s.Products.GetProduct(ctx, name)
	if err != nil {
		log.Println("get product error")
	}

	fmt.Println("--- product information ----")
	fmt.Println("product code", p.ProductCode)
	fmt.Println("product price", p.Price)
	fmt.Println("product stock", p.Stock)
}

func (s *Service) getCampaignInfo(ctx context.Context, command string) {
	res := strings.Split(command, " ")
	name := res[1]

	c, err := s.Campaigns.GetCampaign(ctx, name)
	if err != nil {
		fmt.Println("get campaign error")
	}

	fmt.Println("--- campaign information ----")
	fmt.Println("campaign name", c.Name)
	fmt.Println("campaign duration", c.Duration)
}

func (s *Service) createCampaign(ctx context.Context, command string) {
	now := time.Now()
	res := strings.Split(command, " ")

	name := res[1]
	productCode := res[2]
	duration, _ := strconv.ParseInt(res[3], 10, 64)
	priceManipulationLimit, _ := strconv.ParseInt(res[4], 10, 64)
	targetSalesCount, _ := strconv.ParseInt(res[5], 10, 64)

	var c campaigns.Campaign
	c.ID = uniqid.Generate()
	c.Name = name
	c.ProductCode = productCode
	c.Duration = duration
	c.PriceManipulationLimit = priceManipulationLimit
	c.TargetSalesCount = targetSalesCount
	c.CreatedAt = &now
	c.UpdatedAt = &now

	err := s.Campaigns.CreateCampaign(ctx, &c)
	if err != nil {
		log.Println("create campaign error")
	}
}

func (s *Service) createProduct(ctx context.Context, command string) {
	now := time.Now()
	res := strings.Split(command, " ")

	productCode := res[1]
	price, _ := strconv.ParseInt(res[2], 10, 64)
	stock, _ := strconv.ParseInt(res[3], 10, 64)

	var p products.Product
	p.ID = uniqid.Generate()
	p.ProductCode = productCode
	p.Price = price
	p.Stock = stock
	p.CreatedAt = &now
	p.UpdatedAt = &now

	err := s.Products.CreateProduct(ctx, &p)
	if err != nil {
		log.Println("create product error")
	}
}
