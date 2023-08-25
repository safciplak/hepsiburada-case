package orders

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestEntitySuite struct {
	suite.Suite

	orderEntity Order
}

// SetupTest sets up often used objects
func (test *TestEntitySuite) SetupTest() {

	test.orderEntity = Order{
		ID:          1,
		ProductCode: "SKU-1",
		Quantity:    1,
	}
}

// TestEntityTestSuite Runs the testsuite
func TestEntityTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TestEntitySuite))
}

func (test *TestEntitySuite) TestGetOrderProperties() {
	result := test.orderEntity

	test.Equal(result.ID, int64(1))
	test.Equal(result.ProductCode, "SKU-1")
	test.Equal(result.Quantity, 1)
}

func (test *TestEntitySuite) TestGetOrderPropertiesFailure() {
	result := test.orderEntity

	test.NotEqual(result.ID, int64(2))
}
