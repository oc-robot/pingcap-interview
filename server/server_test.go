package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_server_ServeHTTP(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name     string
		url      string
		wantCode int
		wantResp string
	}{
		{
			name:     "1_ok",
			url:      "http://localhost/latency/200ms",
			wantCode: http.StatusOK,
			wantResp: "",
		}, {
			name:     "2_err_arg",
			url:      "http://localhost/latency/error",
			wantCode: http.StatusBadRequest,
			wantResp: "invalid request body",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()
			s := &server{}
			s.ServeHTTP(w, req)
			resp := w.Result()
			if resp.StatusCode != tt.wantCode {
				t.Errorf("want status code should be %d, but got %d", tt.wantCode, resp.StatusCode)
			}
			body, _ := ioutil.ReadAll(resp.Body)
			if string(body) != tt.wantResp {
				t.Errorf("want response should be %s, but got %s", tt.wantResp, string(body))
			}
		})
	}
}
