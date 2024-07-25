package sample

import (
	"fmt"

	"modern-dev-env-app-sample/internal/sample_app/domain/value"
)

// ListSamplesRequest ユースケース: サンプルデータのリストを取得 のリクエストパラメータ
type ListSamplesRequest struct {
	// ids IDのリスト
	ids []value.SampleID
}

// NewListSamplesRequest ListSamplesRequestのコンストラクタ
func NewListSamplesRequest(ids []value.SampleID) (*ListSamplesRequest, error) {
	req := &ListSamplesRequest{ids: ids}
	if err := req.validate(); err != nil {
		return nil, fmt.Errorf("failed to validate(): %w", err)
	}
	return req, nil
}

// validate バリデーション
func (l *ListSamplesRequest) validate() error {
	return nil
}
