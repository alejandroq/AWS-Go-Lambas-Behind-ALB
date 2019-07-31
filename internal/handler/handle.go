package handler

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

//HandleRequest that is incoming
func HandleRequest(ctx context.Context, event events.ALBTargetGroupRequest) (string, error) {
	log.Printf("HTTP %v %v\n%v", event.HTTPMethod, event.Path, event.Body)
	return "hello world", nil
}
