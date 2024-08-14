package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/internal/logging"
	"io"
	"time"
)

// RequestLoggingMiddleware 为请求添加前置请求日志中间件
func RequestLoggingMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取请求的基本信息
		method := context.Request.Method
		clientIP := context.ClientIP()

		// 获取请求体
		var requestBody []byte
		if context.Request.Body != nil {
			// 读取请求体
			requestBody, _ = io.ReadAll(context.Request.Body)
			// 由于 Request.Body 是一个 io.ReadCloser，读取后会关闭
			// 需要重新设置请求体
			context.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 开始时间
		start := time.Now()
		startTimeFormatted := start.Format("2006-01-02 15:04:05.000")

		// 继续处理请求
		context.Next()

		// 结束时间
		end := time.Now()
		endTimeFormatted := end.Format("2006-01-02 15:04:05.000")
		latency := end.Sub(start)

		logMessage := fmt.Sprintf(`
        Request Start Time: %s
        Request End Time:   %s
        Latency:            %v
        Client IP:          %s
        Request Info:       %s %s
        Request Body:       %s
        Response Status:    %d
        `,
			startTimeFormatted,
			endTimeFormatted,
			latency,
			clientIP,
			method,
			context.Request.RequestURI,
			string(requestBody),
			context.Writer.Status(),
		)

		// 打印请求日志
		logging.Logger().Debugln(logMessage)
	}
}
