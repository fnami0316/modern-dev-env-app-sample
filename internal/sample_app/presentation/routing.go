package presentation

import (
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
	"modern-dev-env-app-sample/internal/sample_app/presentation/sample"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// RegisterRPCServices protoc都合のRPCサービス構造体を登録
// これによりいい感じにリクエストがルーティングされる
func RegisterRPCServices(grpcServer *grpc.Server, db *gorm.DB) error {
	// SampleService
	sampleService, err := sample.NewSampleServiceServer(db)
	if err != nil {
		return fmt.Errorf("failed to NewSampleServiceServer(): %w", err)
	}
	pb.RegisterSampleServiceServer(grpcServer, sampleService)

	return nil
}
