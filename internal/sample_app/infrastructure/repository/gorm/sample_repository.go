package gorm

import (
	"errors"
	"fmt"
	"time"

	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	"modern-dev-env-app-sample/internal/sample_app/infrastructure/repository/gorm/transaction"
	usecase2 "modern-dev-env-app-sample/internal/sample_app/usecase/repository"
	usecase "modern-dev-env-app-sample/internal/sample_app/usecase/repository/transaction"

	"gorm.io/gorm"
)

var _ usecase2.ISampleRepository = &SampleRepository{}

// SampleRepository Sample集約リポジトリ
type SampleRepository struct {
	con *gorm.DB
}

// CreateSampleRepository SampleRepositoryのファクトリ
func CreateSampleRepository(iCon usecase.IConnection) (usecase2.ISampleRepository, error) {
	con, err := transaction.Con(iCon)
	if err != nil {
		return nil, fmt.Errorf("failed to Con(): %w", err)
	}
	return newSampleRepository(con)
}

// newSampleRepository SampleRepositoryのコンストラクタ
func newSampleRepository(con *gorm.DB) (*SampleRepository, error) {
	sampleRepo := &SampleRepository{
		con: con,
	}
	if err := sampleRepo.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate(): %w", err)
	}
	return sampleRepo, nil
}

// validate SampleRepositoryのバリデーション
func (s *SampleRepository) validate() error {
	if s.con == nil {
		return errors.New("con is nil")
	}
	return nil
}

// Save 1件のSampleエンティティを保存
// sampleがnilの場合は即エラー扱い
func (s *SampleRepository) Save(sampleEntity *entity.Sample, iTx usecase.ITransaction) error {
	if sampleEntity == nil {
		return errors.New("sampleEntity is nil")
	}
	sampleGORM, err := s.convEntityToGORM(sampleEntity)
	if err != nil {
		return fmt.Errorf("failed to convEntityToGORM(): %w", err)
	}

	conWithTx, err := transaction.ConWithTx(s.con, iTx)
	if err != nil {
		return fmt.Errorf("failed to conWithTx(): %w", err)
	}
	result := conWithTx.Save(sampleGORM)
	if result.Error != nil {
		return fmt.Errorf("failed to Save(): %w", result.Error)
	}
	return nil
}

// FindByIDs 指定したID群でSampleエンティティ群を取得
// idsのサイズが0の場合は即エラー扱い
func (s *SampleRepository) FindByIDs(ids value.SampleIDs, iTx usecase.ITransaction) ([]*entity.Sample, error) {
	if len(ids) == 0 {
		return nil, errors.New("ids is empty")
	}
	sampleGORMs := make([]*SampleGORM, 0, len(ids))

	conWithTx, err := transaction.ConWithTx(s.con, iTx)
	if err != nil {
		return nil, fmt.Errorf("failed to conWithTx(): %w", err)
	}
	result := conWithTx.Where("id IN ?", ids.ToInt64()).Find(&sampleGORMs)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to Find(): %w", result.Error)
	}

	sampleEntities, err := s.convGORMListToEntityList(sampleGORMs)
	if err != nil {
		return nil, fmt.Errorf("failed to convGORMListToEntityList(): %w", err)
	}
	return sampleEntities, nil
}

// FindAll 全てのSampleエンティティ群を取得
func (s *SampleRepository) FindAll(iTx usecase.ITransaction) ([]*entity.Sample, error) {
	sampleGORMs := make([]*SampleGORM, 0, 0)
	conWithTx, err := transaction.ConWithTx(s.con, iTx)
	if err != nil {
		return nil, fmt.Errorf("failed to conWithTx(): %w", err)
	}
	result := conWithTx.Find(&sampleGORMs)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to Find(): %w", result.Error)
	}

	sampleEntities, err := s.convGORMListToEntityList(sampleGORMs)
	if err != nil {
		return nil, fmt.Errorf("failed to convGORMListToEntityList(): %w", err)
	}
	return sampleEntities, nil
}

// Delete 1件のSampleエンティティを論理削除
// sampleがnilの場合は即エラー扱い
func (s *SampleRepository) Delete(sample *entity.Sample, iTx usecase.ITransaction) error {
	if sample == nil {
		return errors.New("sample is nil")
	}
	conWithTx, err := transaction.ConWithTx(s.con, iTx)
	if err != nil {
		return fmt.Errorf("failed to conWithTx(): %w", err)
	}
	conWithTx.Where("id = ?", sample.ID().ToInt64()).Delete(&SampleGORM{})
	return nil
}

// convEntityToGORM エンティティをGORM用の構造体に変換
// エンティティがnilの場合は即エラー扱い
func (s *SampleRepository) convEntityToGORM(sampleEntity *entity.Sample) (*SampleGORM, error) {
	if sampleEntity == nil {
		return nil, errors.New("sample is nil")
	}

	sampleGORM, err := NewSampleGORM(
		int64(sampleEntity.ID()),
		string(sampleEntity.Name()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to NewSampleGORM(): %w", err)
	}
	return sampleGORM, nil
}

// convGORMToEntity GORM用の構造体をエンティティに変換
// GORM用の構造体がnilの場合は即エラー扱い
func (s *SampleRepository) convGORMToEntity(sampleGORM *SampleGORM) (*entity.Sample, error) {
	if sampleGORM == nil {
		return nil, errors.New("sampleGORM is nil")
	}

	sampleEntityID, err := value.NewSampleID(sampleGORM.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to NewSampleID(): %w", err)
	}
	sampleEntityName, err := value.NewSampleName(sampleGORM.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to NewSampleName(): %w", err)
	}
	sampleEntity, err := entity.NewSample(sampleEntityID, sampleEntityName)
	if err != nil {
		return nil, fmt.Errorf("failed to NewSample(): %w", err)
	}
	return sampleEntity, nil
}

// convGORMListToEntityList GORM用の構造体リストをエンティティリストに変換
func (s *SampleRepository) convGORMListToEntityList(sampleGORMs []*SampleGORM) ([]*entity.Sample, error) {
	sampleEntities := make([]*entity.Sample, len(sampleGORMs))
	for i, sampleGORM := range sampleGORMs {
		sampleEntity, err := s.convGORMToEntity(sampleGORM)
		if err != nil {
			return nil, fmt.Errorf("failed to convGORMToEntity(): %w", err)
		}
		sampleEntities[i] = sampleEntity
	}
	return sampleEntities, nil
}

// GORMに必要な実装
// TODO:
// - ここに記述すべきか判断に迷っている。infrastructure/persistence/sample/を切った上で、repository.go, gorm.go とした方がいいかもしれない。
// - IDの扱いをどうするか迷っている。ドメインとしてのユニークキーであるID と GORMのID(自動採番) を総合的に考える必要がある。
// - gorm.Modelは使わず、同等のカラムを手動でおく。パッと一目でどういう挙動になるかを判断できるようにするため。

// SampleGORM GORM経由での永続化に必要になる構造体
type SampleGORM struct {
	ID        int64 `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// NewSampleGORM SampleGORMのコンストラクタ
// CreatedAt, UpdatedAt, DeletedAtはGORMが自動で設定するため、指定しない
// TODO: IDは大丈夫なんだろうか
func NewSampleGORM(id int64, name string) (*SampleGORM, error) {
	sampleGORM := &SampleGORM{
		ID:   id,
		Name: name,
	}
	if err := sampleGORM.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate(): %w", err)
	}
	return sampleGORM, nil
}

// validate SampleGORMのバリデーション
func (s *SampleGORM) validate() error {
	return nil
}
