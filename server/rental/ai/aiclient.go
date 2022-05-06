package ai

import (
	"context"
	"fmt"
	rentalpb "happy-car/rental/api/gen/v1"
	happyenvpb "happy-car/shared/happyenv"
)

// Client defines an AI client.
type Client struct {
	AIClient  happyenvpb.AIServiceClient
	UseRealAI bool
}

// DistanceKm calculates distance in km.
func (c *Client) DistanceKm(ctx context.Context, from *rentalpb.Location, to *rentalpb.Location) (float64, error) {
	resp, err := c.AIClient.MeasureDistance(ctx, &happyenvpb.MeasureDistanceRequest{
		From: &happyenvpb.Location{
			Latitude:  from.Latitude,
			Longitude: from.Longitude,
		},
		To: &happyenvpb.Location{
			Latitude:  to.Latitude,
			Longitude: to.Longitude,
		},
	})
	if err != nil {
		return 0, err
	}
	return resp.DistanceKm, nil
}

// Resolve resolves identity from given photo.
func (c *Client) Resolve(ctx context.Context, photo []byte) (*rentalpb.Identity, error) {
	i, err := c.AIClient.LicIdentity(ctx, &happyenvpb.IdentityRequest{
		Photo:  photo,
		RealAi: c.UseRealAI,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot resolve identity: %v", err)
	}
	return &rentalpb.Identity{
		Name:            i.Name,
		Gender:          rentalpb.Gender(i.Gender),
		BirthDateMillis: i.BirthDateMillis,
		LicNumber:       i.LicNumber,
	}, nil
}
