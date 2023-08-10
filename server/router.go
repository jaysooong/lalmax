package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/q191201771/naza/pkg/nazalog"
)

func (s *LalMaxServer) InitRouter(router *gin.Engine) {
	if router != nil {
		router.Use(s.Cors())
		// whip
		router.POST("/whip", s.HandleWHIP)
		router.OPTIONS("/whip", s.HandleWHIP)

		// whep
		router.POST("/whep", s.HandleWHEP)
		router.OPTIONS("/whep", s.HandleWHEP)

		// http-fmp4/hls/dash
		router.GET("/live/m4s/:streamid", s.HandleHttpM4s)
	}
}
func (s *LalMaxServer) Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		//服务器支持的所有跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		//允许跨域设置可以返回其他子段，可以自定义字段
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
		// 允许浏览器（客户端）可以解析的头部 （重要）
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		//设置缓存时间
		c.Header("Access-Control-Max-Age", "172800")
		//允许客户端传递校验信息比如 cookie (重要)
		c.Header("Access-Control-Allow-Credentials", "true")

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}
		c.Next()
	}
}
func (s *LalMaxServer) HandleWHIP(c *gin.Context) {
	switch c.Request.Method {
	case "OPTIONS":
		c.Status(http.StatusOK)
	case "POST":
		if s.rtcsvr != nil {
			s.rtcsvr.HandleWHIP(c)
		}
	}
}

func (s *LalMaxServer) HandleWHEP(c *gin.Context) {
	switch c.Request.Method {
	case "OPTIONS":
		c.Status(http.StatusOK)
	case "POST":
		if s.rtcsvr != nil {
			s.rtcsvr.HandleWHEP(c)
		}
	}
}

func (s *LalMaxServer) HandleHttpM4s(c *gin.Context) {
	if strings.HasSuffix(c.Request.URL.Path, ".m3u8") {
		s.handleM3u8(c)
	} else if strings.HasSuffix(c.Request.URL.Path, ".mpd") {
		s.handleDash(c)
	} else if strings.HasSuffix(c.Request.URL.Path, ".mp4") {
		s.handleHttpFmp4(c)
	} else {
		c.Status(http.StatusBadRequest)
		return
	}
}

func (s *LalMaxServer) handleHttpFmp4(c *gin.Context) {
	nazalog.Info("handleHttpFmp4")
	if s.httpfmp4svr != nil {
		s.httpfmp4svr.HandleRequest(c)
	}
}

func (s *LalMaxServer) handleM3u8(c *gin.Context) {
	// TODO 支持hls-fmp4/llhls
	nazalog.Info("handle m3u8")
	c.Status(http.StatusOK)
}

func (s *LalMaxServer) handleDash(c *gin.Context) {
	// TODO 支持dash
	nazalog.Info("handle dash")
	c.Status(http.StatusOK)
}