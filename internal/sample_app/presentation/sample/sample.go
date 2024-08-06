package sample

import (
	"errors"
	"fmt"

	usecase "modern-dev-env-app-sample/internal/sample_app/application/usecase/sample"
	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
)

// SampleServiceServer protocで自動生成されたSampleServiceServerのインターフェースをみたす構造体
type SampleServiceServer struct {
	iListSamplesUseCase  usecase.IListSamplesUseCase
	iCreateSampleUseCase usecase.ICreateSampleUseCase
	pb.UnimplementedSampleServiceServer
}

// NewSampleServiceServer SampleServiceServerのコンストラクタ
func NewSampleServiceServer(
	iListSamplesUseCase usecase.IListSamplesUseCase,
	iCreateSampleUseCase usecase.ICreateSampleUseCase,
) (*SampleServiceServer, error) {
	sampleServiceServer := &SampleServiceServer{
		iListSamplesUseCase:  iListSamplesUseCase,
		iCreateSampleUseCase: iCreateSampleUseCase,
	}
	if err := sampleServiceServer.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate() :%w", err)
	}
	return sampleServiceServer, nil
}

// validate SampleServiceServerのバリデーション
func (s *SampleServiceServer) validate() error {
	if s.iListSamplesUseCase == nil {
		return errors.New("iListSamplesUseCase is nil")
	}
	if s.iCreateSampleUseCase == nil {
		return errors.New("iCreateSampleUseCase is nil")
	}
	return nil
}
