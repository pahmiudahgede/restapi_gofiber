package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Meta struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewErrorResponse(statusCode int, message string) *Response {
	resp := &Response{}
	resp.Meta.StatusCode = statusCode
	resp.Meta.Message = message
	resp.Data = nil
	return resp
}

func NewSuccessResponse(message string, data interface{}) *Response {
	resp := &Response{}
	resp.Meta.StatusCode = fiber.StatusOK
	resp.Meta.Message = message
	resp.Data = data
	return resp
}
