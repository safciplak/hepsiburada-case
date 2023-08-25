package campaigns

import (
	"context"

	"github.com/hepsiburada/hepsiburada-command/app"
)

var (
	ErrorNotfound = app.BusinessError("campaign not found")
)

type campaigns interface {
	CreateCampaign(ctx context.Context, c *Campaign) error
	GetCampaign(ctx context.Context, campaignName string) (*Campaign, error)
}

type Service struct {
	Campaigns campaigns
}

func (s *Service) CreateCampaign(ctx context.Context, c *Campaign) error {
	return s.Campaigns.CreateCampaign(ctx, c)
}

func (s *Service) GetCampaign(ctx context.Context, campaignName string) (*Campaign, error) {
	return s.Campaigns.GetCampaign(ctx, campaignName)
}
