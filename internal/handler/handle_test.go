package handler

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {
	type args struct {
		ctx   context.Context
		event events.ALBTargetGroupRequest
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return an expected output",
			args: args{
				ctx: context.TODO(),
				event: events.ALBTargetGroupRequest{},
			},
			want: "hello world",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleRequest(tt.args.ctx, tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HandleRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
