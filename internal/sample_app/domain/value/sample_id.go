package value

import (
	"errors"
	"fmt"
)

type SampleID int64

// NewSampleID SampleIDのコンストラクタ
func NewSampleID(id int64) (SampleID, error) {
	sampleID := SampleID(id)
	if err := sampleID.validate(); err != nil {
		return 0, fmt.Errorf("failed to validate(): %w", err)
	}
	return SampleID(id), nil
}

// validate SampleIDのバリデーション
func (s SampleID) validate() error {
	if s <= 0 {
		return errors.New(fmt.Sprintf("SampleID must be greater than 0 (s:%v)", s))
	}
	return nil
}

type SampleIDs []SampleID

func (s SampleIDs) ToInt64() []int64 {
	ids := make([]int64, len(s))
	for i, id := range s {
		ids[i] = int64(id)
	}
	return ids
}
