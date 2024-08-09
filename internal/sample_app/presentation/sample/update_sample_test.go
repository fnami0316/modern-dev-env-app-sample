package sample

import (
	"context"
	"reflect"
	"testing"

	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	"modern-dev-env-app-sample/internal/sample_app/application/usecase/sample"
	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
)

func TestSampleServiceServer_UpdateSample(t *testing.T) {
	type fields struct {
		iListSamplesUseCase              sample.IListSamplesUseCase
		iCreateSampleUseCase             sample.ICreateSampleUseCase
		iUpdateSampleUseCase             sample.IUpdateSampleUseCase
		iDeleteSampleUseCase             sample.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		in0 context.Context
		req *pb.UpdateSampleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UpdateSampleResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SampleServiceServer{
				iListSamplesUseCase:              tt.fields.iListSamplesUseCase,
				iCreateSampleUseCase:             tt.fields.iCreateSampleUseCase,
				iUpdateSampleUseCase:             tt.fields.iUpdateSampleUseCase,
				iDeleteSampleUseCase:             tt.fields.iDeleteSampleUseCase,
				UnimplementedSampleServiceServer: tt.fields.UnimplementedSampleServiceServer,
			}
			got, err := s.UpdateSample(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateSample() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSample() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleServiceServer_convertToUpdateSampleRequestForUseCase(t *testing.T) {
	type fields struct {
		iListSamplesUseCase              sample.IListSamplesUseCase
		iCreateSampleUseCase             sample.ICreateSampleUseCase
		iUpdateSampleUseCase             sample.IUpdateSampleUseCase
		iDeleteSampleUseCase             sample.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		pbReq *pb.UpdateSampleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *application.UpdateSampleRequest
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				pbReq: &pb.UpdateSampleRequest{
					Id:   "id1",
					Name: "name1",
				},
			},
			want:    newUpdateSampleRequestForTest(t, "id1", "name1"),
			wantErr: false,
		},
		{
			name: "[NG]値オブジェクトSampleID生成時にエラー",
			args: args{
				pbReq: &pb.UpdateSampleRequest{
					Id:   "", // エラー
					Name: "name1",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[NG]値オブジェクトSampleName生成時にエラー",
			args: args{
				pbReq: &pb.UpdateSampleRequest{
					Id:   "id1",
					Name: "", // エラー
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SampleServiceServer{
				iListSamplesUseCase:              tt.fields.iListSamplesUseCase,
				iCreateSampleUseCase:             tt.fields.iCreateSampleUseCase,
				iUpdateSampleUseCase:             tt.fields.iUpdateSampleUseCase,
				iDeleteSampleUseCase:             tt.fields.iDeleteSampleUseCase,
				UnimplementedSampleServiceServer: tt.fields.UnimplementedSampleServiceServer,
			}
			got, err := s.convertToUpdateSampleRequestForUseCase(tt.args.pbReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToUpdateSampleRequestForUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToUpdateSampleRequestForUseCase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleServiceServer_convertToUpdateSampleResponseForProtoc(t *testing.T) {
	type fields struct {
		iListSamplesUseCase              sample.IListSamplesUseCase
		iCreateSampleUseCase             sample.ICreateSampleUseCase
		iUpdateSampleUseCase             sample.IUpdateSampleUseCase
		iDeleteSampleUseCase             sample.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		useCaseRes *application2.UpdateSampleResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UpdateSampleResponse
		wantErr bool
	}{
		{
			name: "[OK]インスタンスを生成できる",
			args: args{
				useCaseRes: newUpdateSampleResponseForTest(t, newSampleForTest(t, "id1", "name1")),
			},
			want: &pb.UpdateSampleResponse{
				Sample: &pb.Sample{
					Id:   "id1",
					Name: "name1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SampleServiceServer{
				iListSamplesUseCase:              tt.fields.iListSamplesUseCase,
				iCreateSampleUseCase:             tt.fields.iCreateSampleUseCase,
				iUpdateSampleUseCase:             tt.fields.iUpdateSampleUseCase,
				iDeleteSampleUseCase:             tt.fields.iDeleteSampleUseCase,
				UnimplementedSampleServiceServer: tt.fields.UnimplementedSampleServiceServer,
			}
			got, err := s.convertToUpdateSampleResponseForProtoc(tt.args.useCaseRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToUpdateSampleResponseForProtoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToUpdateSampleResponseForProtoc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// newUpdateSampleRequestForTest UpdateSampleRequestを生成(エラー発生時はテスト失敗扱い)
func newUpdateSampleRequestForTest(t *testing.T, id value.SampleID, name value.SampleName) *application.UpdateSampleRequest {
	req, err := application.NewUpdateSampleRequest(id, name)
	if err != nil {
		t.Fatalf("failed to NewUpdateSampleRequest(): %v", err)
	}
	return req
}

// newUpdateSampleResponseForTest UpdateSampleResponseを生成(エラー発生時はテスト失敗扱い)
func newUpdateSampleResponseForTest(t *testing.T, sample *entity.Sample) *application2.UpdateSampleResponse {
	res, err := application2.NewUpdateSampleResponse(sample)
	if err != nil {
		t.Fatalf("failed to NewUpdateSampleResponse(): %v", err)
	}
	return res
}
