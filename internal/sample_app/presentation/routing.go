package presentation

import (
	"fmt"

	infra "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm"
	infra2 "modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm/transaction"
	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
	"modern-dev-env-app-sample/internal/sample_app/presentation/sample"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// RegisterRPCServices protoc都合のRPCサービス構造体を登録
// TODO: どのインフラ詳細を使うのか？をプレゼンテーション層でやるのはおかしい気がする。神たるmain関数の責務では。
// RegisterSampleServiceServer()のような、通信の仕組みを整えるのはプレゼンテーション層でもいい気はするけど
func RegisterRPCServices(grpcServer *grpc.Server, db *gorm.DB) error {
	// SampleService
	iCon := infra2.NewGORMConnection(db)
	iSampleRepo, err := infra.CreateSampleRepository(iCon)
	if err != nil {
		return fmt.Errorf("failed to CreateSampleRepository(): %w", err)
	}
	sampleService, err := sample.NewSampleServiceServer(iCon, iSampleRepo)
	if err != nil {
		return fmt.Errorf("failed to NewSampleServiceServer(): %w", err)
	}

	pb.RegisterSampleServiceServer(grpcServer, sampleService)
	return nil
}
