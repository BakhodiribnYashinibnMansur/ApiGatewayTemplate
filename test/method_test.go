package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)



func HttpClientRequest(Router *gin.Engine, url string, method string, query map[string]string, token string, body interface{}) (resp *httptest.ResponseRecorder, err error) {
	w := httptest.NewRecorder()
	byteBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	var req *http.Request
	if method == "GET" {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewReader(byteBody))
	}
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	Router.ServeHTTP(w, req)

	return w, err

}
