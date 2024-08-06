package sample

import (
	"errors"
	"fmt"

	usecase "modern-dev-env-app-sample/internal/sample_app/application/repository"
	"modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	sample2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
)

var _ IListSamplesUseCase = &ListSamplesUseCase{}

// ListSamplesUseCase ユースケース: サンプルデータのリストを取得
type ListSamplesUseCase struct {
	iSampleRepo usecase.ISampleRepository
}

// NewListSamplesUseCase ListSamplesUseCaseのコンストラクタ
func NewListSamplesUseCase(iSampleRepo usecase.ISampleRepository) (*ListSamplesUseCase, error) {
	useCase := &ListSamplesUseCase{iSampleRepo: iSampleRepo}
	if err := useCase.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate(): %w", err)
	}
	return useCase, nil
}

// Run ユースケース: サンプルデータのリストを取得 を実行
func (l *ListSamplesUseCase) Run(req *sample.ListSamplesRequest) (*sample2.ListSamplesResponse, error) {
	samples, err := l.iSampleRepo.FindByIDs(req.IDs(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to FindSamples(): %w", err)
	}

	res, err := sample2.NewListSamplesResponse(samples)
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
