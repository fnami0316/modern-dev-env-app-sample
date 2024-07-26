package sample

import (
	"errors"
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
	usecase2 "modern-dev-env-app-sample/internal/sample_app/usecase/repository"
	usecase "modern-dev-env-app-sample/internal/sample_app/usecase/repository/transaction"
)

// SampleServiceServer protocで自動生成されたSampleServiceServerのインターフェースをみたす構造体
type SampleServiceServer struct {
	iCon        usecase.IConnection
	iSampleRepo usecase2.ISampleRepository
	pb.UnimplementedSampleServiceServer
}

// NewSampleServiceServer SampleServiceServerのコンストラクタ
func NewSampleServiceServer(iCon usecase.IConnection, iRepo usecase2.ISampleRepository) (*SampleServiceServer, error) {
	sampleServiceServer := &SampleServiceServer{
		iCon:        iCon,
		iSampleRepo: iRepo,
	}
	if err := sampleServiceServer.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate() :%w", err)
	}
	return sampleServiceServer, nil
}

// validate SampleServiceServerのバリデーション
func (s *SampleServiceServer) validate() error {
	if s.iCon == nil {
		return errors.New("db connection is nil")
	}
	if s.iSampleRepo == nil {
		return errors.New("sample repository is nil")
	}
	return nil
}
