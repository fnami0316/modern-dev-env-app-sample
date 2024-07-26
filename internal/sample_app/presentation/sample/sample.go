package sample

import (
	"errors"
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
	usecase "modern-dev-env-app-sample/internal/sample_app/usecase/sample"
)

// SampleServiceServer protocで自動生成されたSampleServiceServerのインターフェースをみたす構造体
type SampleServiceServer struct {
	listSamplesUseCase *usecase.ListSamplesUseCase
	pb.UnimplementedSampleServiceServer
}

// NewSampleServiceServer SampleServiceServerのコンストラクタ
func NewSampleServiceServer(listSamplesUseCase *usecase.ListSamplesUseCase) (*SampleServiceServer, error) {
	sampleServiceServer := &SampleServiceServer{
		listSamplesUseCase: listSamplesUseCase,
	}
	if err := sampleServiceServer.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate() :%w", err)
	}
	return sampleServiceServer, nil
}

// validate SampleServiceServerのバリデーション
func (s *SampleServiceServer) validate() error {
	if s.listSamplesUseCase == nil {
		return errors.New("listSamplesUseCase is nil")
	}
	return nil
}
