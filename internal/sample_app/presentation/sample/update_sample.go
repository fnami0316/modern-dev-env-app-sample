package sample

import (
	"context"
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	sample2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
)

// UpdateSample (protoc依存のRPCメソッド実装) サンプルデータを更新
func (s *SampleServiceServer) UpdateSample(_ context.Context, req *pb.UpdateSampleRequest) (*pb.UpdateSampleResponse, error) {
	// protoc都合のリクエストパラメータ構造体をユースケース層都合のものに変換
	useCaseReq, err := s.convertToUpdateSampleRequestForUseCase(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convertToUpdateSampleRequestForUseCase(): %w", err)
	}

	// ユースケースを実行
	useCaseRes, err := s.iUpdateSampleUseCase.Run(useCaseReq)
	if err != nil {
		return nil, fmt.Errorf("failed to Run(): %w", err)
	}

	// ユースケース都合のレスポンスパラメータ構造体をprotoc都合のものに変換
	pbRes, err := s.convertToUpdateSampleResponseForProtoc(useCaseRes)
	if err != nil {
		return nil, fmt.Errorf("failed to convertToUpdateSampleResponseForProtoc(): %w", err)
	}
	return pbRes, nil
}

// convertToUpdateSampleRequestForUseCase protoc都合のUpdateSampleのリクエストパラメータ構造体をユースケース都合のものに変換
func (s *SampleServiceServer) convertToUpdateSampleRequestForUseCase(pbReq *pb.UpdateSampleRequest) (*sample.UpdateSampleRequest, error) {
	// 各パラメータを値オブジェクトへ地道に変換
	sampleID, err := value.NewSampleID(pbReq.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to NewSampleID(): %w", err)
	}
	sampleName, err := value.NewSampleName(pbReq.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to NewSampleName(): %w", err)
	}

	// ユースケース層都合のリクエストパラメータ構造体を生成
	useCaseReq, err := sample.NewUpdateSampleRequest(sampleID, sampleName)
	if err != nil {
		return nil, fmt.Errorf("failed to NewUpdateSampleRequest(): %w", err)
	}
	return useCaseReq, nil
}

// convertToUpdateSampleResponseForProtoc ユースケース都合のUpdateSampleのレスポンスパラメータ構造体をprotoc都合のものに変換
func (s *SampleServiceServer) convertToUpdateSampleResponseForProtoc(useCaseRes *sample2.UpdateSampleResponse) (*pb.UpdateSampleResponse, error) {
	// 各パラメータをprotoc都合の型に地道に変換
	pbSample := &pb.Sample{
		Id:   useCaseRes.Sample().ID().ToString(),
		Name: useCaseRes.Sample().Name().ToString(),
	}

	// protoc都合のレスポンスパラメータ構造体を生成
	pbRes := &pb.UpdateSampleResponse{Sample: pbSample}
	return pbRes, nil
}
