package sample

import (
	"context"
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	sample2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
)

// ListSamples (protoc依存のRPCメソッド実装) サンプルデータのリストを取得
// protoc都合のリクエストパラメータ構造体をユースケース層都合の構造体に変換した上で、本質的な処理はユースケース層にあるメソッドへとルーティング
// ユースケース層のメソッドから返ってきた結果は、protoc都合のレスポンスパラメータ構造体に変換して返す
func (s *SampleServiceServer) ListSamples(_ context.Context, req *pb.ListSamplesRequest) (*pb.ListSamplesResponse, error) {
	// protoc都合のリクエストパラメータ構造体をユースケース層都合のものに変換
	useCaseReq, err := s.convertToListSamplesRequestForUseCase(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convertToListSamplesRequestForUseCase(): %w", err)
	}

	// ユースケースを実行
	useCaseRes, err := s.iListSamplesUseCase.Run(useCaseReq)
	if err != nil {
		return nil, fmt.Errorf("failed to Run(): %w", err)
	}

	// ユースケース都合のレスポンスパラメータ構造体をprotoc都合のものに変換
	pbRes, err := s.convertToListSamplesResponseForProtoc(useCaseRes)
	if err != nil {
		return nil, fmt.Errorf("failed to convertToListSamplesResponseForProtoc(): %w", err)
	}
	return pbRes, nil
}

// convertToListSamplesRequestForUseCase protoc都合のListSamplesのリクエストパラメータ構造体をユースケース都合のものに変換
func (s *SampleServiceServer) convertToListSamplesRequestForUseCase(pbReq *pb.ListSamplesRequest) (*sample.ListSamplesRequest, error) {
	// 各パラメータを値オブジェクトへ地道に変換
	sampleIDs := make([]value.SampleID, len(pbReq.Ids))
	for i, pbReqID := range pbReq.Ids {
		sampleID, err := value.NewSampleID(pbReqID)
		if err != nil {
			return nil, fmt.Errorf("failed to NewSampleID(): %w", err)
		}
		sampleIDs[i] = sampleID
	}

	// ユースケース層都合のリクエストパラメータ構造体を生成
	useCaseReq, err := sample.NewListSamplesRequest(sampleIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to NewListSamplesRequest(): %w", err)
	}
	return useCaseReq, nil
}

// convertToListSamplesResponseForProtoc ユースケース都合のListSamplesのレスポンスパラメータ構造体をprotoc都合のものに変換
func (s *SampleServiceServer) convertToListSamplesResponseForProtoc(useCaseRes *sample2.ListSamplesResponse) (*pb.ListSamplesResponse, error) {
	// 各パラメータをprotoc都合の型に地道に変換
	pbSamples := make([]*pb.Sample, len(useCaseRes.Samples()))
	for i, sampleEntity := range useCaseRes.Samples() {
		pbSample := &pb.Sample{
			Id:   sampleEntity.ID().ToString(),
			Name: sampleEntity.Name().ToString(),
		}
		pbSamples[i] = pbSample
	}

	// protoc都合のレスポンスパラメータ構造体を生成
	pbRes := &pb.ListSamplesResponse{Samples: pbSamples}
	return pbRes, nil
}
