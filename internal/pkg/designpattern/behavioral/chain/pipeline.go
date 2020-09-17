package chain

import (
	"fmt"
	"reflect"
)

/**
entity
*/
type Request struct {
	tagId string
}

func (request *Request) SetTagId(tagId string) {
	request.tagId = tagId
}

func (request *Request) GetTagId() string {
	return request.tagId
}

type Response struct {
}

/**
interface
*/
type Handler interface {
	DoHandler(request Request, response Response)
}

type Pipeline interface {
	AddLast(handler Handler)

	DoHandler(request Request, response Response)
}

/**
implement
*/
type PipelineImpl struct {
	handlers []Handler
}

func (pipeline *PipelineImpl) AddLast(handler Handler) {
	if pipeline.handlers == nil || len(pipeline.handlers) < 1 {
		pipeline.handlers = []Handler{}
	}
	pipeline.handlers = append(pipeline.handlers, handler)
}

func (pipeline *PipelineImpl) DoHandler(request Request, response Response) {
	for index, handler := range pipeline.handlers {
		fmt.Println("handler index: ", index, " -> ", &handler, " -> ", reflect.TypeOf(handler))
		handler.DoHandler(request, response)
	}
}

/**
metrics handler
*/
type MetricsHandler struct{}

func (handler *MetricsHandler) DoHandler(request Request, response Response) {
	fmt.Println("metrics", "->", request.tagId)
}

/**
global handler
*/
type GlobalHandler struct{}

func (handler *GlobalHandler) DoHandler(request Request, response Response) {
	fmt.Println("global", "->", request.tagId)
}
