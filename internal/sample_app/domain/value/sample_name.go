package value

import "fmt"

type SampleName string

// NewSampleName SampleNameのコンストラクタ
func NewSampleName(name string) (SampleName, error) {
	sampleName := SampleName(name)
	if err := sampleName.validate(); err != nil {
		return "", fmt.Errorf("failed to validate(): %w", err)
	}
	return sampleName, nil
}

// validate SampleNameのバリデーション
func (s SampleName) validate() error {
	if len(s) == 0 {
		return fmt.Errorf("SampleName must not be empty")
	}
	return nil
}

// ToString SampleNameを文字列に変換
func (s SampleName) ToString() string {
	return string(s)
}
