// http://qiita.com/taizo/items/32d895e35397336bf285
package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Response struct {
	path, query, contenttype, body string
}

func TestOutputResponse(t *testing.T) {
	response := &Response{
		path:        "/",
		contenttype: "text/plain",
		body:        `foo`,
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Check request.
		if g, w := r.URL.Path, response.path; g != w {
			t.Errorf("request got path %s, want %s", g, w)
		}

		// Send response.
		w.Header().Set("Content-Type", response.contenttype)
		io.WriteString(w, response.body)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	ret, err := outputResponse(server.URL + response.path)
	assert.Nil(t, err)
	assert.Equal(t, `foo`, ret, "return body should be foo")
}
