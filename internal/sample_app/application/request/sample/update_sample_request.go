package sample

import "modern-dev-env-app-sample/internal/sample_app/domain/value"

// UpdateSampleRequest ユースケース: サンプルデータを更新 のリクエストパラメータ
type UpdateSampleRequest struct {
	// id ID
	id value.SampleID
	// name 名前
	name value.SampleName
}

// NewUpdateSampleRequest UpdateSampleRequestのコンストラクタ
func NewUpdateSampleRequest(
	id value.SampleID,
	name value.SampleName,
) (*UpdateSampleRequest, error) {
	req := &UpdateSampleRequest{
		id:   id,
		name: name,
	}
	if err := req.validate(); err != nil {
		return nil, err
	}
	return req, nil
}

// validate バリデーション
func (c *UpdateSampleRequest) validate() error {
	return nil
}

// ID IDを取得
func (c *UpdateSampleRequest) ID() value.SampleID {
	return c.id
}

// Name 名前を取得
func (c *UpdateSampleRequest) Name() value.SampleName {
	return c.name
}
