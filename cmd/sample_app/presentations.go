package main

import (
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
	"modern-dev-env-app-sample/internal/sample_app/presentation/sample"

	"google.golang.org/grpc"
)

// presentations 全プレゼンテーション層のインスタンスをまとめた構造体
type presentations struct {
	sampleServiceServer *sample.SampleServiceServer
}

// newPresentations コンストラクタ
func newPresentations(
	sampleServiceServer *sample.SampleServiceServer,
) *presentations {
	return &presentations{
		sampleServiceServer: sampleServiceServer,
	}
}

// createPresentations 全プレゼンテーション層のインスタンスのファクトリ
func createPresentations(
	useCases *useCases,
) (*presentations, error) {
	sampleServiceServer, err := sample.NewSampleServiceServer(useCases.listSamplesUseCase)
	if err != nil {
		return nil, fmt.Errorf("failed to NewSampleServiceServer(): %w", err)
	}
	return newPresentations(sampleServiceServer), nil
}

// registerProtocServices protoc都合のRPCサービス構造体を登録
// Register<サービス名>Server(grpc.Serverのポインタ, <サービス名>Serverインターフェース)関数は、
// protocコマンド実行で生成された「<元になった.proroファイル名>_grpc.pb.go」内に自動で定義されている
// ここで登録されたRPCサービスについてのみ、gRPC通信が可能になる
func (p *presentations) registerProtocServices(grpcServer *grpc.Server) {
	pb.RegisterSampleServiceServer(grpcServer, p.sampleServiceServer)
}
