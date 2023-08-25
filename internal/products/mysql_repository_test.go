package products

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestProductRepositorySuite struct {
	suite.Suite
}

// SetupTest sets up often used objects
func (test *TestProductRepositorySuite) SetupTest() {

}

// TestProductRepositoryTestSuite Runs the testsuite
func TestProductRepositoryTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TestProductRepositorySuite))
}

func TestGetOrder(m *testing.T) {
	var mock sqlmock.Sqlmock
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(m, err)

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		m.Errorf("Failed to open connection to DB: %v", err)
	}

	if conn == nil {
		m.Error("Failed to open connection to DB: conn is nil")
	}

	defer db.Close()

	mock.ExpectQuery("SELECT * FROM products WHERE name = ? ORDER BY id LIMIT 1").WithArgs("product-code")

	db.QueryContext(context.Background(), "SELECT * FROM products WHERE name = ? ORDER BY id LIMIT 1", "product-code")

	err = mock.ExpectationsWereMet()
	assert.NoError(m, err)
}

func TestCreateProduct(m *testing.T) {
	var mock sqlmock.Sqlmock
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(m, err)

	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		m.Errorf("Failed to open connection to DB: %v", err)
	}

	if conn == nil {
		m.Error("Failed to open connection to DB: conn is nil")
	}

	defer db.Close()

	mock.ExpectQuery("INSERT INTO products (id,	product_code,price,	stock, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)").
		WithArgs(1, "PRODUCT-SKU", 1, 2, "2023-08-24", "2023-08-24")

	db.QueryContext(context.Background(), "INSERT INTO products (id,	product_code,price,	stock, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", 1, "PRODUCT-SKU", 1, 2, "2023-08-24", "2023-08-24")

	err = mock.ExpectationsWereMet()
	assert.NoError(m, err)
}
