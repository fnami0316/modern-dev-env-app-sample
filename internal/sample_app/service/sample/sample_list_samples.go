package sample

import (
	"context"

	pb "modern-dev-env-app-sample/internal/sample_app/pb/api/proto"
)

// ListSamples サンプルデータのリストを取得
func (s *SampleServiceServer) ListSamples(_ context.Context, _ *pb.ListSamplesRequest) (*pb.ListSamplesResponse, error) {
	return &pb.ListSamplesResponse{
		Samples: []*pb.Sample{
			{
				Id:   1,
				Name: "name1",
			},
			{
				Id:   2,
				Name: "name2",
			},
		},
	}, nil
}
