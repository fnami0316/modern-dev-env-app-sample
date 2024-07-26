package sample

import (
	"errors"
	"fmt"

	domain "modern-dev-env-app-sample/internal/sample_app/usecase/repository"
)

// ListSamplesUseCase ユースケース: サンプルデータのリストを取得
type ListSamplesUseCase struct {
	iSampleRepo domain.ISampleRepository
}

// NewListSamplesUseCase ListSamplesUseCaseのコンストラクタ
func NewListSamplesUseCase(iSampleRepo domain.ISampleRepository) (*ListSamplesUseCase, error) {
	useCase := &ListSamplesUseCase{iSampleRepo: iSampleRepo}
	if err := useCase.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate(): %w", err)
	}
	return useCase, nil
}

// Run ユースケース: サンプルデータのリストを取得 を実行
func (l *ListSamplesUseCase) Run(req *ListSamplesRequest) (*ListSamplesResponse, error) {
	samples, err := l.iSampleRepo.FindByIDs(req.ids, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to FindSamples(): %w", err)
	}

	res, err := NewListSamplesResponse(samples)
	if err != nil {
		return nil, fmt.Errorf("failed to NewListSamplesResponse(): %w", err)
	}
	return res, nil
}

// validate ListSamplesUseCaseのバリデーション
func (l *ListSamplesUseCase) validate() error {
	if l.iSampleRepo == nil {
		return errors.New("iSampleRepo is nil")
	}
	return nil
}
