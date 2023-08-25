package campaigns

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestEntitySuite struct {
	suite.Suite

	campaignEntity Campaign
}

// SetupTest sets up often used objects
func (test *TestEntitySuite) SetupTest() {

	test.campaignEntity = Campaign{
		ID:                     1235123,
		Name:                   "test campaign",
		ProductCode:            "TEST-SKU",
		Duration:               10,
		PriceManipulationLimit: 5,
		TargetSalesCount:       20,
	}
}

// TestEntityTestSuite Runs the testsuite
func TestEntityTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TestEntitySuite))
}

func (test *TestEntitySuite) TestGetCampaignProperties() {
	result := test.campaignEntity

	test.Equal(result.ID, int64(1235123))
	test.Equal(result.Name, "test campaign")
}

func (test *TestEntitySuite) TestGetCampaignPropertiesFailure() {
	result := test.campaignEntity

	test.NotEqual(result.ID, int64(2))
}
