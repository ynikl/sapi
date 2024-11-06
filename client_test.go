package sip

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func Test_innerClient_Get(t *testing.T) {
	type fields struct {
		client  *http.Client
		timeout time.Duration
		reqLogs []*RequestLog
	}
	type args struct {
		url    string
		params map[string]string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantResBody []byte
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name: "normal request without error",
			fields: fields{
				client:  client,
				timeout: 0,
				reqLogs: nil,
			},
			args: args{
				url:    "https://www.baidu.com/search",
				params: map[string]string{"q": "hello"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &innerClient{
				client:  tt.fields.client,
				timeout: tt.fields.timeout,
				reqLogs: tt.fields.reqLogs,
			}
			gotResBody, err := i.Get(tt.args.url, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResBody, tt.wantResBody) {
				t.Errorf("Get() gotResBody = %v, want %v", gotResBody, tt.wantResBody)
			}
		})
	}
}
