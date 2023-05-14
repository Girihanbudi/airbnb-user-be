package log

import (
	elastic "airbnb-user-be/internal/pkg/elasticsearch"
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func LogRequestToElastic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Create a log writter
		logWriter := &GinLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = logWriter

		// Create request log data
		startTime := time.Now()
		requestBody, _ := ioutil.ReadAll(ctx.Request.Body)
		reader := ioutil.NopCloser(bytes.NewBuffer(requestBody))
		ctx.Request.Body = reader
		request := Request{
			Time:      startTime,
			Method:    ctx.Request.Method,
			Uri:       ctx.Request.RequestURI,
			Proto:     ctx.Request.Proto,
			UserAgent: ctx.Request.UserAgent(),
			Referer:   ctx.Request.Referer(),
			Body:      string(requestBody),
			Ip:        ctx.ClientIP(),
		}

		// Processing the request
		ctx.Next()

		// Create response log data
		endTime := time.Now()
		response := Response{
			Time:       endTime,
			StatusCode: ctx.Writer.Status(),
			Body:       logWriter.body.String(),
		}

		body := Log{
			ProcessTime: endTime.Sub(startTime),
			Request:     request,
			Response:    response,
		}

		go elastic.Send(body, "api")
	}
}
