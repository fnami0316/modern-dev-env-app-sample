package repository

import (
	usecase "modern-dev-env-app-sample/internal/sample_app/application/repository/transaction"
	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
)

// ISampleRepository Sample集約リポジトリのインターフェース
type ISampleRepository interface {
	// Save 1件のSampleエンティティを保存
	Save(sample *entity.Sample, iTx usecase.ITransaction) error
	// FindByIDs 指定したID群でSampleエンティティ群を取得
	FindByIDs(ids value.SampleIDs, iTx usecase.ITransaction) ([]*entity.Sample, error)
	// FindAll 全てのSampleエンティティ群を取得
	FindAll(iTx usecase.ITransaction) ([]*entity.Sample, error)
	// Delete 1件のSampleエンティティを論理削除
	Delete(sample *entity.Sample, iTx usecase.ITransaction) error
}
