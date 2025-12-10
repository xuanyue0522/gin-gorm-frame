package middleware

import (
	"bytes"
	"gin-gorm-frame/consts"
	"gin-gorm-frame/utils/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"time"
)

// GetRequestBody 获取请求体
func getRequestBody(ctx *gin.Context) string {
	data, _ := io.ReadAll(ctx.Request.Body)
	return string(data)
}

func GetResponseBody(ctx *gin.Context) string {
	resp := ctx.Request.Response
	if resp == nil || resp.Body == nil {
		return ""
	}
	data, _ := io.ReadAll(resp.Body)
	return string(data)
}

type responseWriterWrapper struct {
	gin.ResponseWriter
	Writer io.Writer
}

func (w *responseWriterWrapper) Write(data []byte) (int, error) {
	return w.Writer.Write(data)
}

func AccessLogMiddleware(filter func(*gin.Context) bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if filter != nil && !filter(ctx) { // 过滤器不为空，并且过滤器返回false，则不记录日志
			ctx.Next()
			return
		}

		body := getRequestBody(ctx)

		// Gin 框架中，上下文中的body读取一遍之后后续的中间件就无法获取了，所以需要重新赋值
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(body)))
		begin := time.Now()

		fields := []zap.Field{
			zap.String("ip", ctx.RemoteIP()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("params", ctx.Request.URL.RawQuery),
			zap.Any("body", body),
			zap.String("token", ctx.GetHeader(consts.AdminTokenKey)),
		}

		var responseBody bytes.Buffer
		multiWriter := io.MultiWriter(ctx.Writer, &responseBody)
		ctx.Writer = &responseWriterWrapper{
			ResponseWriter: ctx.Writer,
			Writer:         multiWriter,
		}

		ctx.Next()
		respBody := responseBody.String()
		if len(respBody) > 1024 {
			respBody = respBody[:1024]
		}
		fields = append(fields, zap.Int64("dur_ms", time.Since(begin).Milliseconds()))
		fields = append(fields, zap.Int("status", ctx.Writer.Status()))
		fields = append(fields, zap.String("resp", respBody))
		logger.Info("access_log", fields...)
	}
}
