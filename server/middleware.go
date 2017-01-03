package server

import (
	"io"
	"log"
	"strings"
	"time"

	"github.com/cloudcloud/shortoftime/config"
	"github.com/gin-gonic/gin"
)

var (
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	reset   = string([]byte{27, 91, 48, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})

	methodColour = map[string]string{
		"GET":     blue,
		"POST":    cyan,
		"PUT":     yellow,
		"DELETE":  red,
		"PATCH":   green,
		"HEAD":    magenta,
		"OPTIONS": white,
	}
	statusColour = map[int]string{
		2: green,
		3: white,
		4: yellow,
	}
)

const (
	// Name is the internal display name of the server
	Name = "shortofti.me"
)

// PrettyLogger will take the built-in Gin logger and enhance it a little
func PrettyLogger(o io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		start, req, raw := time.Now(), c.Request.URL.Path, c.Request.URL.RawQuery

		// make the logged URI useful
		if len(strings.TrimSpace(raw)) > 0 {
			req = req + "?" + raw
		}

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		ip, method, status := c.ClientIP(), c.Request.Method, c.Writer.Status()
		methodCol, statusCol := colourMethod(method), colourStatus(status)
		comm := c.Errors.ByType(gin.ErrorTypePrivate).String()

		log.Printf("[%s] %v |%s %3d %s| %13v | %s |%s %s %-7s %s\n%s",
			Name,
			end.Format("2006/01/02 - 15:04:05"),
			statusCol, status, reset,
			latency,
			ip,
			methodCol, method, reset,
			req,
			comm,
		)
	}
}

// PushConfig will make the Config object available via the context for all requests
func PushConfig(c config.Conf) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Config", c)
		c.Next()
	}
}

// colourMethod will determine a display colour for a given method
func colourMethod(m string) string {
	c, ok := methodColour[m]
	if !ok {
		c = reset
	}

	return c
}

// colourStatus will determine a display colour for a given status
func colourStatus(s int) string {
	c, ok := statusColour[s/100]
	if !ok {
		c = red
	}

	return c
}
