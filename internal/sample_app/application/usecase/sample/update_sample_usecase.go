package sample

import (
	"errors"
	"fmt"

	usecase "modern-dev-env-app-sample/internal/sample_app/application/repository"
	usecase2 "modern-dev-env-app-sample/internal/sample_app/application/repository/transaction"
	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
)

var _ IUpdateSampleUseCase = &UpdateSampleUseCase{}

type UpdateSampleUseCase struct {
	iCon        usecase2.IConnection
	iSampleRepo usecase.ISampleRepository
}

// NewUpdateSampleUseCase UpdateSampleUsecaseのコンストラクタ
func NewUpdateSampleUseCase(iCon usecase2.IConnection, iSampleRepo usecase.ISampleRepository) (*UpdateSampleUseCase, error) {
	useCase := &UpdateSampleUseCase{
		iCon:        iCon,
		iSampleRepo: iSampleRepo,
	}
	if err := useCase.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate(): %w", err)
	}
	return useCase, nil
}

// validate UpdateSampleUsecaseのバリデーション
func (l *UpdateSampleUseCase) validate() error {
	if l.iCon == nil {
		return errors.New("iCon is nil")
	}
	if l.iSampleRepo == nil {
		return errors.New("iSampleRepo is nil")
	}
	return nil
}

// Run ユースケース: サンプルデータを更新 を実行
func (l *UpdateSampleUseCase) Run(req *application.UpdateSampleRequest) (*application2.UpdateSampleResponse, error) {
	id := req.ID()
	name := req.Name()
	var updatedSample *entity.Sample
	if err := l.iCon.Transaction(func(iTx usecase2.ITransaction) error {
		samples, err := l.iSampleRepo.FindByIDs(value.SampleIDs{id}, iTx)
		if err != nil {
			return fmt.Errorf("failed to FindByIDs(): %w", err)
		}
		if len(samples) == 0 {
			return fmt.Errorf("failed to FindByIDs(): request id not exist in data store")
		}

		sampleEntity := samples[0]
		updatedSample, err = sampleEntity.Update(name)
		if err != nil {
			return fmt.Errorf("failed to Update(): %w", err)
		}

		if err := l.iSampleRepo.Save(updatedSample, iTx); err != nil {
			return fmt.Errorf("failed to Save(): %w", err)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to Transaction(): %w", err)
	}

	res, err := application2.NewUpdateSampleResponse(updatedSample)
	if err != nil {
		return nil, fmt.Errorf("failed to NewUpdateSampleResponse(): %w", err)
	}
	return res, nil
}
