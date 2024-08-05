package sample

import (
	"fmt"

	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
)

// ListSamplesResponse ユースケース: サンプルデータのリストを取得 のレスポンスパラメータ
type ListSamplesResponse struct {
	// samples サンプルデータのリスト
	samples entity.Samples
}

// NewListSamplesResponse ListSamplesResponseのコンストラクタ
func NewListSamplesResponse(samples entity.Samples) (*ListSamplesResponse, error) {
	res := &ListSamplesResponse{samples: samples}
	if err := res.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate(): %w", err)
	}
	return res, nil
}

// validate バリデーション
func (l *ListSamplesResponse) validate() error {
	return nil
}

// Samples サンプルデータのリストを取得
func (l *ListSamplesResponse) Samples() entity.Samples {
	return l.samples
}
