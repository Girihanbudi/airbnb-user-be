package elastic

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

type Log struct {
	ProcessTime time.Duration
	Request     Request
	Response    Response
}

type Request struct {
	Time      time.Time
	Method    string
	Uri       string
	Proto     string
	UserAgent string
	Referer   string
	Body      string
	Ip        string
}

type Response struct {
	Time       time.Time
	StatusCode int
	Body       string
}

type GinLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}
