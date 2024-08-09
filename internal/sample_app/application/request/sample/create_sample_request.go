package sample

import "modern-dev-env-app-sample/internal/sample_app/domain/value"

// CreateSampleRequest ユースケース: サンプルデータを追加 のリクエストパラメータ
type CreateSampleRequest struct {
	// name 名前
	name value.SampleName
}

// NewCreateSampleRequest CreateSampleRequestのコンストラクタ
func NewCreateSampleRequest(name value.SampleName) (*CreateSampleRequest, error) {
	req := &CreateSampleRequest{name: name}
	if err := req.validate(); err != nil {
		return nil, err
	}
	return req, nil
}

// validate バリデーション
func (c *CreateSampleRequest) validate() error {
	return nil
}

// Name 名前を取得
func (c *CreateSampleRequest) Name() value.SampleName {
	return c.name
}
