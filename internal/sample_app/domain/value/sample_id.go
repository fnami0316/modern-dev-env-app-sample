package value

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type SampleID string

// CreateRandomSampleID ランダムなSampleIDを生成するファクトリ
// UUID v4で生成
func CreateRandomSampleID() (SampleID, error) {
	id := uuid.NewString()
	sampleID, err := NewSampleID(id)
	if err != nil {
		return "", fmt.Errorf("failed to NewSampleID(): %w", err)
	}
	return sampleID, nil
}

// NewSampleID SampleIDのコンストラクタ
func NewSampleID(id string) (SampleID, error) {
	sampleID := SampleID(id)
	if err := sampleID.validate(); err != nil {
		return "", fmt.Errorf("failed to validate(): %w", err)
	}
	return SampleID(id), nil
}

// validate SampleIDのバリデーション
func (s SampleID) validate() error {
	if len(s) == 0 {
		return errors.New(fmt.Sprintf("SampleID size must be greater than 0 (s:%v)", s))
	}
	return nil
}

// ToString stringに変換
func (s SampleID) ToString() string {
	return string(s)
}

type SampleIDs []SampleID

// ToString stringに変換
func (s SampleIDs) ToString() []string {
	ids := make([]string, len(s))
	for i, id := range s {
		ids[i] = id.ToString()
	}
	return ids
}
