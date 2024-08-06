package sample

import (
	"errors"
	"fmt"

	usecase "modern-dev-env-app-sample/internal/sample_app/application/repository"
	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
)

var _ ICreateSampleUseCase = &CreateSampleUsecase{}

type CreateSampleUsecase struct {
	iSampleRepo usecase.ISampleRepository
}

// NewCreateSampleUsecase CreateSampleUsecaseのコンストラクタ
func NewCreateSampleUsecase(iSampleRepo usecase.ISampleRepository) (*CreateSampleUsecase, error) {
	useCase := &CreateSampleUsecase{iSampleRepo: iSampleRepo}
	if err := useCase.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate(): %w", err)
	}
	return useCase, nil
}

// validate CreateSampleUsecaseのバリデーション
func (l *CreateSampleUsecase) validate() error {
	if l.iSampleRepo == nil {
		return errors.New("iSampleRepo is nil")
	}
	return nil
}

// Run ユースケース: サンプルデータを作成 を実行
func (l *CreateSampleUsecase) Run(req *application.CreateSampleRequest) (*application2.CreateSampleResponse, error) {
	return nil, nil
}
