package sample

import (
	"context"
	"reflect"
	"testing"

	application "modern-dev-env-app-sample/internal/sample_app/application/request/sample"
	application2 "modern-dev-env-app-sample/internal/sample_app/application/response/sample"
	application3 "modern-dev-env-app-sample/internal/sample_app/application/usecase/sample"
	entity "modern-dev-env-app-sample/internal/sample_app/domain/entity/sample"
	"modern-dev-env-app-sample/internal/sample_app/domain/value"
	"modern-dev-env-app-sample/internal/sample_app/presentation/pb"
)

func TestSampleServiceServer_ListSamples(t *testing.T) {
	type fields struct {
		iListSamplesUseCase              application3.IListSamplesUseCase
		iCreateSampleUseCase             application3.ICreateSampleUseCase
		iUpdateSampleUseCase             application3.IUpdateSampleUseCase
		iDeleteSampleUseCase             application3.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		in0 context.Context
		req *pb.ListSamplesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ListSamplesResponse
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
			got, err := s.ListSamples(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSamples() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListSamples() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleServiceServer_convertToListSamplesRequestForUseCase(t *testing.T) {
	type fields struct {
		iListSamplesUseCase              application3.IListSamplesUseCase
		iCreateSampleUseCase             application3.ICreateSampleUseCase
		iUpdateSampleUseCase             application3.IUpdateSampleUseCase
		iDeleteSampleUseCase             application3.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		pbReq *pb.ListSamplesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *application.ListSamplesRequest
		wantErr bool
	}{
		{
			name:   "[OK]インスタンスを生成できる",
			fields: fields{},
			args: args{
				pbReq: &pb.ListSamplesRequest{
					Ids: []string{
						"id1",
						"id2",
						"id3",
					},
				},
			},
			want:    newListSamplesRequestForTest(t, value.SampleIDs{"id1", "id2", "id3"}),
			wantErr: false,
		},
		{
			name:   "[NG]idが不正な場合エラー",
			fields: fields{},
			args: args{
				pbReq: &pb.ListSamplesRequest{
					Ids: []string{
						"id1",
						"", // エラー
						"id3",
					},
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
			got, err := s.convertToListSamplesRequestForUseCase(tt.args.pbReq)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToListSamplesRequestForUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToListSamplesRequestForUseCase() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleServiceServer_convertToListSamplesResponseForProtoc(t *testing.T) {
	sample1 := newSampleForTest(t, "id1", "name1")

	type fields struct {
		iListSamplesUseCase              application3.IListSamplesUseCase
		iCreateSampleUseCase             application3.ICreateSampleUseCase
		iUpdateSampleUseCase             application3.IUpdateSampleUseCase
		iDeleteSampleUseCase             application3.IDeleteSampleUseCase
		UnimplementedSampleServiceServer pb.UnimplementedSampleServiceServer
	}
	type args struct {
		useCaseRes *application2.ListSamplesResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ListSamplesResponse
		wantErr bool
	}{
		{
			name:   "[OK]インスタンスを生成できる",
			fields: fields{},
			args: args{
				useCaseRes: newListSamplesResponseForTest(t, entity.Samples{sample1}),
			},
			want: &pb.ListSamplesResponse{
				Samples: []*pb.Sample{
					{
						Id:   "id1",
						Name: "name1",
					},
				},
			},
			wantErr: false,
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
			got, err := s.convertToListSamplesResponseForProtoc(tt.args.useCaseRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToListSamplesResponseForProtoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToListSamplesResponseForProtoc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// newListSamplesRequestForTest ListSamplesRequestを生成(エラーが発生した場合はテスト失敗扱い)
func newListSamplesRequestForTest(t *testing.T, ids value.SampleIDs) *application.ListSamplesRequest {
	req, err := application.NewListSamplesRequest(ids)
	if err != nil {
		t.Fatalf("failed to NewListSamplesRequest(): %v", err)
	}
	return req
}

// newListSamplesResponseForTest ListSamplesResponseを生成(エラーが発生した場合はテスト失敗扱い)
func newListSamplesResponseForTest(t *testing.T, samples entity.Samples) *application2.ListSamplesResponse {
	res, err := application2.NewListSamplesResponse(samples)
	if err != nil {
		t.Fatalf("failed to NewListSamplesResponse(): %v", err)
	}
	return res
}

// newSampleForTest Sampleを生成(エラーが発生した場合はテスト失敗扱い)
func newSampleForTest(t *testing.T, id value.SampleID, name value.SampleName) *entity.Sample {
	sample, err := entity.NewSample(id, name)
	if err != nil {
		t.Fatalf("failed to NewSample(): %v", err)
	}
	return sample
}
