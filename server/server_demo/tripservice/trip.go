package trip

import (
	"context"

	trippb "happy-car/server_demo/proto/gen/go"
)

// type TripServiceServer interface {
// 	GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
// }

// Service implements trip service.
type Service struct{}

// GetTrip 一般来说， 接收者都是一个指针类型的参数
// 既然是一个服务，一般都用指针接收者
func (s *Service) GetTrip(ctx context.Context, request *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: request.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "def",
			DurationSec: 3600,
			FeeCent:     10000,
			StartPos: &trippb.Location{
				Longitude: 30,
				Latitude:  120,
			},
			EndPos: &trippb.Location{
				Longitude: 35,
				Latitude:  120,
			},
			PathLocations: []*trippb.Location{
				{
					Longitude: 30,
					Latitude:  120,
				},
				{
					Longitude: 32,
					Latitude:  118,
				},
			},
			Status: trippb.TripStatus_IN_PROGRESS,
		},
	}, nil
}
