package sample

import "modern-dev-env-app-sample/internal/sample_app/domain/value"

// DeleteSampleRequest ユースケース: サンプルデータを削除 のリクエストパラメータ
type DeleteSampleRequest struct {
	// id ID
	id value.SampleID
}

// NewDeleteSampleRequest DeleteSampleRequestのコンストラクタ
func NewDeleteSampleRequest(
	id value.SampleID,
) (*DeleteSampleRequest, error) {
	req := &DeleteSampleRequest{
		id: id,
	}
	if err := req.validate(); err != nil {
		return nil, err
	}
	return req, nil
}

// validate バリデーション
func (c *DeleteSampleRequest) validate() error {
	return nil
}

// ID IDを取得
func (c *DeleteSampleRequest) ID() value.SampleID {
	return c.id
}
