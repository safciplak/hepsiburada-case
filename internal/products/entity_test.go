package products

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestEntitySuite struct {
	suite.Suite

	productEntity Product
}

// SetupTest sets up often used objects
func (test *TestEntitySuite) SetupTest() {

	test.productEntity = Product{
		ID:          1,
		ProductCode: "SKU-1",
		Price:       1,
		Stock:       1,
	}
}

// TestEntityTestSuite Runs the testsuite
func TestEntityTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TestEntitySuite))
}

func (test *TestEntitySuite) TestGetProductProperties() {
	result := test.productEntity

	test.Equal(result.ID, int64(1))
	test.Equal(result.ProductCode, "SKU-1")
	test.Equal(result.Price, int64(1))
}

func (test *TestEntitySuite) TestGetProductPropertiesFailure() {
	result := test.productEntity

	test.NotEqual(result.ID, int64(2))
}
