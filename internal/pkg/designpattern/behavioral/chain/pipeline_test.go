package chain

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	request := new(Request)
	request.SetTagId("test")

	tagId := request.GetTagId()

	fmt.Println(tagId)
	fmt.Println(request.GetTagId())
}

func TestPipelineImpl(t *testing.T) {

	var pip Pipeline = new(PipelineImpl)
	metricsHandler := new(MetricsHandler)
	globalHandler := new(GlobalHandler)

	pip.AddLast(metricsHandler)
	pip.AddLast(globalHandler)

	request := new(Request)
	request.SetTagId("test-xxx")

	response := new(Response)

	pip.DoHandler(*request, *response)
}
