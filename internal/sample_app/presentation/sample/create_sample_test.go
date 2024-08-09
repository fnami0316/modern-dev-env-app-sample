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

func TestSampleServiceServer_CreateSample(t *testing.T) {
	type fields struct {
		iListSamplesUseCase              sample.IListSamplesUseCase
		iCreateSampleUseCase             sample.ICreateSampleUseCase
		iUpdateSampleUseCase             sample.IUpdateSampleUseCase
		iDeleteSampleUseCase             sample.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		in0 context.Context
		req *pb.CreateSampleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CreateSampleResponse
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
			got, err := s.CreateSample(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSample() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSample() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleServiceServer_convertToCreateSampleRequestForUseCase(t *testing.T) {
	type fields struct {
		iListSamplesUseCase              sample.IListSamplesUseCase
		iCreateSampleUseCase             sample.ICreateSampleUseCase
		iUpdateSampleUseCase             sample.IUpdateSampleUseCase
		iDeleteSampleUseCase             sample.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		pbReq *pb.CreateSampleRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *application.CreateSampleRequest
		wantErr bool
	}{
		{
			name:   "[OK]インスタンスを生成できる",
			fields: fields{},
			args: args{
				pbReq: &pb.CreateSampleRequest{
					Name: "test",
				},
			},
			want:    newCreateSampleRequestForTest(t, "test"),
			wantErr: false,
		},
		{
			name:   "[NG]値オブジェクトの生成でエラー",
			fields: fields{},
			args: args{
				pbReq: &pb.CreateSampleRequest{
					Name: "", //　エラー
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
			got, err := s.convertToCreateSampleRequestForUseCase(tt.args.pbReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToCreateSampleRequestForUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToCreateSampleRequestForUseCase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleServiceServer_convertToCreateSampleResponseForProtoc(t *testing.T) {
	type fields struct {
		iListSamplesUseCase              sample.IListSamplesUseCase
		iCreateSampleUseCase             sample.ICreateSampleUseCase
		iUpdateSampleUseCase             sample.IUpdateSampleUseCase
		iDeleteSampleUseCase             sample.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		useCaseRes *application2.CreateSampleResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CreateSampleResponse
		wantErr bool
	}{
		{
			name:   "[OK]インスタンスを生成できる",
			fields: fields{},
			args: args{
				useCaseRes: newCreateSampleResponseForTest(t, newSampleForTest(t, "id1", "name1")),
			},
			want: &pb.CreateSampleResponse{
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
			got, err := s.convertToCreateSampleResponseForProtoc(tt.args.useCaseRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToCreateSampleResponseForProtoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToCreateSampleResponseForProtoc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// newCreateSampleRequestForTest CreateSamplesRequestを生成(エラーが発生した場合はテスト失敗扱い)
func newCreateSampleRequestForTest(t *testing.T, name value.SampleName) *application.CreateSampleRequest {
	req, err := application.NewCreateSampleRequest(name)
	if err != nil {
		t.Fatalf("failed to NewCreateSampleRequest(): %v", err)
	}
	return req
}

// newCreateSampleResponseForTest CreateSamplesResponseを生成(エラーが発生した場合はテスト失敗扱い)
func newCreateSampleResponseForTest(t *testing.T, sample *entity.Sample) *application2.CreateSampleResponse {
	res, err := application2.NewCreateSampleResponse(sample)
	if err != nil {
		t.Fatalf("failed to NewCreateSampleResponse(): %v", err)
	}
	return res
}
