package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRpnHandler(t *testing.T) {

	tests := []struct {
		jsonStr  []byte
		expected string
		status   int
		err      error
	}{
		{
			[]byte(`{"rpns":[{"input":"10 -1 +"},{"input":"10 -2 +"}]}`),
			`{"rpn":[{"name":"10 -1 +","result":9},{"name":"10 -2 +","result":8}]}`,
			200,
			nil,
		},

		{
			[]byte(`{"rpns":[{"input":"10 -1"},{"input":"10 -2 +"}]}`),
			`{"rpn":[{"name":"10 -1 +","result":9},{"name":"10 -2 +","result":8}]}`,
			500,
			errors.New("invailid rpn input"),
		},
		{
			[]byte{},
			`{"rpn":[{"name":"10 -1 +","result":9},{"name":"10 -2 +","result":8}]}`,
			500,
			errors.New("unexpected end of JSON input"),
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest("POST", "/parse", bytes.NewBuffer(test.jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(RpnHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.status {
			t.Errorf("handler returned wrong status code: got %v want %v", status, test.status)
		}
		if test.status == 200 {
			if strings.TrimRight(rr.Body.String(), "\n") != test.expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), test.expected)
			}
		} else {
			if strings.TrimRight(rr.Body.String(), "\n") != test.err.Error() {
				t.Errorf("hander return expected error: want %v get %v", test.err.Error(), rr.Body.String())
			}
		}
	}
}

func TestHealthHander(t *testing.T) {
	tests := []struct {
		expected string
		status   int
		err      error
	}{
		{
			`{"status":"OK","code":200}`,
			200,
			nil,
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/health", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HealthHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.status {
			t.Errorf("handler returned wrong status code: got %v want %v", status, test.status)
		}
		if test.status == 200 {
			if strings.TrimRight(rr.Body.String(), "\n") != test.expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), test.expected)
			}
		} else {
			if strings.TrimRight(rr.Body.String(), "\n") != test.err.Error() {
				t.Errorf("hander return expected error: got %v want %v", test.err.Error(), rr.Body.String())
			}
		}
	}

}
