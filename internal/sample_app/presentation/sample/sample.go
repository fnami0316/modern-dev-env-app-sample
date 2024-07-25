package sample

import (
	"errors"
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"

	"gorm.io/gorm"
)

// SampleServiceServer protocで自動生成されたSampleServiceServerのインターフェースをみたす構造体
type SampleServiceServer struct {
	db *gorm.DB
	pb.UnimplementedSampleServiceServer
}

// NewSampleServiceServer SampleServiceServerのコンストラクタ
func NewSampleServiceServer(db *gorm.DB) (*SampleServiceServer, error) {
	sampleServiceServer := &SampleServiceServer{db: db}
	if err := sampleServiceServer.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate() :%w", err)
	}
	return sampleServiceServer, nil
}

// validate SampleServiceServerのバリデーション
func (s *SampleServiceServer) validate() error {
	if s.db == nil {
		return errors.New("db is nil")
	}
	return nil
}
