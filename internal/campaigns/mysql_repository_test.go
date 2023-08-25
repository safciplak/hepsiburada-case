package campaigns

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

type TestCampaignRepositorySuite struct {
	suite.Suite
}

// SetupTest sets up often used objects
func (test *TestCampaignRepositorySuite) SetupTest() {

}

// TestClientTestSuite Runs the testsuite
func TestCampaignRepositoryTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TestCampaignRepositorySuite))
}

func TestGetCampaign(m *testing.T) {
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

	mock.ExpectQuery("SELECT * FROM campaigns WHERE name = ? ORDER BY id LIMIT 1").WithArgs("test-campaign")

	db.QueryContext(context.Background(), "SELECT * FROM campaigns WHERE name = ? ORDER BY id LIMIT 1", "test-campaign")

	err = mock.ExpectationsWereMet()
	assert.NoError(m, err)
}

func TestCreateCampaign(m *testing.T) {
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

	mock.ExpectQuery("INSERT INTO campaigns (id, name, product_code, duration, price_manipulation_limit, target_sales_count,created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)").
		WithArgs(1, "test campaign", "PRODUCT-SKU", 1, 2, 3, "2023-08-24", "2023-08-24")

	db.QueryContext(context.Background(), "INSERT INTO campaigns (id, name, product_code, duration, price_manipulation_limit, target_sales_count,created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", 1, "test campaign", "PRODUCT-SKU", 1, 2, 3, "2023-08-24", "2023-08-24")

	err = mock.ExpectationsWereMet()
	assert.NoError(m, err)
}
