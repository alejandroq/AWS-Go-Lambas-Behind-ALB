package main

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/aws/aws-lambda-go/events"
	"github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

//FeatureContext for godog to evaluate features
func FeatureContext(s *godog.Suite) {
	s.Step(`^I send a GET request$`, iSendAGETRequest)
	s.Step(`^I receive a response with: (.*)$`, iReceiveAResponseWithHelloWorld)

	s.BeforeScenario(func(interface{}) { ts = scenario{} })
}

type scenario struct{ resp *[]byte }

//ts is a Test Scenario
var ts scenario

func iSendAGETRequest() error {
	payload := events.ALBTargetGroupRequest{
		HTTPMethod: "GET",
	}
	resp, err := golambdainvoke.Run(golambdainvoke.Input{
		Port:    8001,
		Payload: payload,
	})
	if err != nil {
		return fmt.Errorf("failed to RPC call http://localhost:8001 with: %v", err)
	}
	ts.resp = &resp
	return nil
}

func iReceiveAResponseWithHelloWorld(ex string) error {
	if *ts.resp == nil {
		return fmt.Errorf("expected %v; a request was never sent, therefore no responses to assert", ex)
	}

	if string(*ts.resp) != ex {
		return fmt.Errorf("expected %v; got %v", ex, string(*ts.resp))
	}

	return nil
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "should pass",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
