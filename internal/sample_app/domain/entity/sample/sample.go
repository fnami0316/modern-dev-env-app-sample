package sample

import (
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/domain/value"
)

// Sample Sampleエンティティ
type Sample struct {
	// id ID
	id value.SampleID
	// name 名前
	name value.SampleName
}

// CreateDefaultSample デフォルトのSampleを生成するファクトリ
// idは自動採番
func CreateDefaultSample(name value.SampleName) (*Sample, error) {
	id, err := value.CreateRandomSampleID()
	if err != nil {
		return nil, fmt.Errorf("failed to CreateRandomSampleID(): %w", err)
	}
	sample, err := NewSample(id, name)
	if err != nil {
		return nil, fmt.Errorf("failed to NewSample(): %w", err)
	}
	return sample, nil
}

// NewSample コンストラクタ
func NewSample(id value.SampleID, name value.SampleName) (*Sample, error) {
	s := &Sample{
		id:   id,
		name: name,
	}

	if err := s.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate(): %w", err)
	}
	return s, nil
}

// validate バリデーション
func (s *Sample) validate() error {
	return nil
}

// ID IDを取得
func (s *Sample) ID() value.SampleID {
	return s.id
}

// Name 名前を取得
func (s *Sample) Name() value.SampleName {
	return s.name
}

// Update フィールド値を更新
// 外部から更新可能なのはnameのみ
func (s *Sample) Update(name value.SampleName) (*Sample, error) {
	return s.update(name)
}

// update 更新
func (s *Sample) update(name value.SampleName) (*Sample, error) {
	updatedSample, err := NewSample(s.id, name)
	if err != nil {
		return nil, fmt.Errorf("failed to NewSample(): %w", err)
	}
	return updatedSample, nil
}

// Samples サンプルデータのリスト
type Samples []*Sample
