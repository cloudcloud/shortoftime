package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/cloudcloud/shortoftime/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// RespFunc for response generation
type RespFunc func(*httptest.ResponseRecorder)

// ReqConf is essentially the data for an individual test case
type ReqConf struct {
	Body      string
	Finaliser RespFunc
	Headers   map[string]string
	Method    string
	Path      string
}

var (
	conf config.Conf
)

func setup() Serv {
	dir, err := os.Getwd()
	if err != nil {
		dir = "../public"
	}

	conf = &config.Config{Debug: true, Listener: "0.0.0.0:8000", RootDir: dir + "../public"}

	return Init(conf, gin.TestMode)
}

func TestServing(t *testing.T) {
	assert := assert.New(t)
	s := setup()

	assert.NotNil(s, "Server instance succeeded")
}

// RunReq is a helper to simplify execution of HTTP calls
func RunReq(s Serv, r ReqConf) {
	qry := ""
	if strings.Contains(r.Path, "?") {
		str := strings.Split(r.Path, "?")
		r.Path = str[0]
		qry = str[1]
	}

	body := bytes.NewBufferString(r.Body)
	req, _ := http.NewRequest(r.Method, r.Path, body)

	if len(qry) > 0 {
		req.URL.RawQuery = qry
	}

	if len(r.Headers) > 0 {
		for k, v := range r.Headers {
			req.Header.Set(k, v)
		}
	} else if r.Method == "POST" || r.Method == "PUT" {
		if strings.HasPrefix(r.Body, "{") {
			req.Header.Set("Content-Type", "application/json")
		} else {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	}

	w := httptest.NewRecorder()
	s.(*Server).GinServ.ServeHTTP(w, req)

	if r.Finaliser != nil {
		r.Finaliser(w)
	}
}
