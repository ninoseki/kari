package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_root(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func Test_extract(t *testing.T) {
	r := setupRouter()

	values := url.Values{}
	values.Add("data", "1.1.1.1 google.com www.us-cert.org f6f8179ac71eaabff12b8c024342109b")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/extract", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	data := new(Cacadordata)
	bytes := ([]byte)(w.Body.String())
	err := json.Unmarshal(bytes, data)
	assert.Equal(t, nil, err)
	assert.Equal(t, "1.1.1.1", data.Networks.Ipv4s[0])
	assert.Equal(t, "google.com", data.Networks.Domains[0])
	assert.Equal(t, "www.us-cert.org", data.Networks.Domains[1])
	assert.Equal(t, "f6f8179ac71eaabff12b8c024342109b", data.Hashes.Md5s[0])
}
