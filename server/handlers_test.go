package server

import (
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	s := setup()

	r := ReqConf{
		Method: "GET",
		Path:   "/",
		Finaliser: func(r *httptest.ResponseRecorder) {
			t.Errorf("r=%+v", r)
			//
		},
	}

	RunReq(s, r)
}

func TestVersion(t *testing.T) {
	s := setup()

	r := ReqConf{
		Method: "GET",
		Path:   "/v1/version",
		Finaliser: func(r *httptest.ResponseRecorder) {
			t.Logf("r=%+v", r)
		},
	}

	RunReq(s, r)
}
