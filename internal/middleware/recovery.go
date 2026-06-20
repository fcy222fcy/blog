package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"blog/pkg/response"

	"github.com/gin-gonic/gin"
)

// Recovery Panic 恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 检查是否为连接断开
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Errorf("[Recovery] panic: %v\n%s", err, string(httpRequest))
					c.Abort()
					return
				}

				logger.Errorf("[Recovery] panic: %v\n%s\n%s", err, string(httpRequest), debug.Stack())
				response.ServerError(c, "服务器内部错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}
