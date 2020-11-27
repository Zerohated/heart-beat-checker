package controller

import (
	conf "github.com/Zerohated/heart-beat-checker/configs"
	"github.com/Zerohated/tools/pkg/logger"
)

var (
	config = conf.Config
	log    = logger.Logger
)

// Return Codes
const (
	CodeOK  = 0
	CodeErr = -1
)

// RespOK returns when no error happen
type respOK struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func newRespOK(code int, data interface{}) respOK {
	return respOK{Code: code, Data: data}
}

// RespErr returns when error happen
type respErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newRespErr(message string) respErr {
	return respErr{Code: CodeErr, Message: message}
}

// Controller example
type Controller struct {
	URL string
}

// NewController example
func NewController() *Controller {
	return &Controller{
		URL: config.URL,
	}
}
