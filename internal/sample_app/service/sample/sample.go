package sample

import pb "modern-dev-env-app-sample/internal/sample_app/pb/api/proto"

// SampleServiceServer SampleService のサーバーAPI実装
type SampleServiceServer struct {
	pb.UnimplementedSampleServiceServer
}
