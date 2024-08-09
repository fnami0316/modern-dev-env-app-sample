package sample

import (
	"fmt"

	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
)

// UpdateSampleResponse ユースケース: サンプルデータを更新 のレスポンスパラメータ
type UpdateSampleResponse struct {
	// サンプルデータ
	sample *entity.Sample
}

// NewUpdateSampleResponse UpdateSampleResponseのコンストラクタ
func NewUpdateSampleResponse(sample *entity.Sample) (*UpdateSampleResponse, error) {
	res := &UpdateSampleResponse{sample: sample}
	if err := res.validate(); err != nil {
		return nil, err
	}
	return res, nil
}

// validate バリデーション
func (c *UpdateSampleResponse) validate() error {
	if c.sample == nil {
		return fmt.Errorf("sample is nil")
	}
	return nil
}

// Sample サンプルデータを取得
func (c *UpdateSampleResponse) Sample() *entity.Sample {
	return c.sample
}
