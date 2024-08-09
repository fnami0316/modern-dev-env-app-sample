package sample

import (
	"context"
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	sample2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteSample (protoc依存のRPCメソッド実装) サンプルデータを削除
func (s *SampleServiceServer) DeleteSample(_ context.Context, req *pb.DeleteSampleRequest) (*pb.DeleteSampleResponse, error) {
	// protoc都合のリクエストパラメータ構造体をユースケース層都合のものに変換
	useCaseReq, err := s.convertToDeleteSampleRequestForUseCase(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convertToDeleteSampleRequestForUseCase(): %w", err)
	}

	// ユースケースを実行
	useCaseRes, err := s.iDeleteSampleUseCase.Run(useCaseReq)
	if err != nil {
		return nil, fmt.Errorf("failed to Run(): %w", err)
	}

	// ユースケース都合のレスポンスパラメータ構造体をprotoc都合のものに変換
	pbRes, err := s.convertToDeleteSampleResponseForProtoc(useCaseRes)
	if err != nil {
		return nil, fmt.Errorf("failed to convertToDeleteSampleResponseForProtoc(): %w", err)
	}
	return pbRes, nil
}

// convertToDeleteSampleRequestForUseCase protoc都合のDeleteSampleのリクエストパラメータ構造体をユースケース都合のものに変換
func (s *SampleServiceServer) convertToDeleteSampleRequestForUseCase(pbReq *pb.DeleteSampleRequest) (*sample.DeleteSampleRequest, error) {
	// 各パラメータを値オブジェクトへ地道に変換
	sampleID, err := value.NewSampleID(pbReq.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to NewSampleID(): %w", err)
	}

	// ユースケース層都合のリクエストパラメータ構造体を生成
	useCaseReq, err := sample.NewDeleteSampleRequest(sampleID)
	if err != nil {
		return nil, fmt.Errorf("failed to NewDeleteSampleRequest(): %w", err)
	}
	return useCaseReq, nil
}

// convertToDeleteSampleResponseForProtoc ユースケース都合のDeleteSampleのレスポンスパラメータ構造体をprotoc都合のものに変換
func (s *SampleServiceServer) convertToDeleteSampleResponseForProtoc(_ *sample2.DeleteSampleResponse) (*pb.DeleteSampleResponse, error) {
	// protoc都合のレスポンスパラメータ構造体を生成
	pbRes := &pb.DeleteSampleResponse{
		Empty: &emptypb.Empty{},
	}
	return pbRes, nil
}
