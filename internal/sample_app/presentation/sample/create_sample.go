package sample

import (
	"context"
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	sample2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
)

// CreateSample (protoc依存のRPCメソッド実装) サンプルデータを追加
func (s *SampleServiceServer) CreateSample(_ context.Context, req *pb.CreateSampleRequest) (*pb.CreateSampleResponse, error) {
	// protoc都合のリクエストパラメータ構造体をユースケース層都合のものに変換
	useCaseReq, err := s.convertToCreateSampleRequestForUseCase(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convertToCreateSampleRequestForUseCase(): %w", err)
	}

	// ユースケースを実行
	useCaseRes, err := s.iCreateSampleUseCase.Run(useCaseReq)
	if err != nil {
		return nil, fmt.Errorf("failed to Run(): %w", err)
	}

	// ユースケース都合のレスポンスパラメータ構造体をprotoc都合のものに変換
	pbRes, err := s.convertToCreateSampleResponseForProtoc(useCaseRes)
	if err != nil {
		return nil, fmt.Errorf("failed to convertToCreateSampleResponseForProtoc(): %w", err)
	}
	return pbRes, nil
}

// convertToCreateSampleRequestForUseCase protoc都合のCreateSampleのリクエストパラメータ構造体をユースケース都合のものに変換
func (s *SampleServiceServer) convertToCreateSampleRequestForUseCase(pbReq *pb.CreateSampleRequest) (*sample.CreateSampleRequest, error) {
	// 各パラメータを値オブジェクトへ地道に変換
	sampleName, err := value.NewSampleName(pbReq.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to NewSampleName(): %w", err)
	}

	// ユースケース層都合のリクエストパラメータ構造体を生成
	useCaseReq, err := sample.NewCreateSampleRequest(sampleName)
	if err != nil {
		return nil, fmt.Errorf("failed to NewCreateSampleRequest(): %w", err)
	}
	return useCaseReq, nil
}

// convertToCreateSampleResponseForProtoc ユースケース都合のCreateSampleのレスポンスパラメータ構造体をprotoc都合のものに変換
func (s *SampleServiceServer) convertToCreateSampleResponseForProtoc(useCaseRes *sample2.CreateSampleResponse) (*pb.CreateSampleResponse, error) {
	// 各パラメータをprotoc都合の型に地道に変換
	pbSample := &pb.Sample{
		Id:   useCaseRes.Sample().ID().ToString(),
		Name: useCaseRes.Sample().Name().ToString(),
	}

	// protoc都合のレスポンスパラメータ構造体を生成
	pbRes := &pb.CreateSampleResponse{Sample: pbSample}
	return pbRes, nil
}
