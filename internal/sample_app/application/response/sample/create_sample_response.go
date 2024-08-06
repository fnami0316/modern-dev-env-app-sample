package sample

import entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"

// CreateSampleResponse ユースケース: サンプルデータを追加 のレスポンスパラメータ
type CreateSampleResponse struct {
	// サンプルデータ
	sample *entity.Sample
}

// NewCreateSampleResponse CreateSampleResponseのコンストラクタ
func NewCreateSampleResponse(sample *entity.Sample) (*CreateSampleResponse, error) {
	res := &CreateSampleResponse{sample: sample}
	if err := res.validate(); err != nil {
		return nil, err
	}
	return res, nil
}

// validate バリデーション
func (c *CreateSampleResponse) validate() error {
	return nil
}

// Sample サンプルデータを取得
func (c *CreateSampleResponse) Sample() *entity.Sample {
	return c.sample
}
