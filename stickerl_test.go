package main

import (
  "fmt"
  "testing"
  "net/http"
  "net/http/httptest"
  assert "github.com/pilu/miniassert"
)

func newTestRequest(method, path string) *httptest.ResponseRecorder {
  request, _ := http.NewRequest(method, path, nil)
  recorder := httptest.NewRecorder()

  app.ServeHTTP(recorder, request)

  return recorder
}

func TestRootHandler(t *testing.T) {
  recorder := newTestRequest("GET", "/")

  assert.Equal(t, http.StatusOK, recorder.Code)
  assert.Equal(t, "application/json", recorder.HeaderMap.Get("Content-Type"))

  expectedBody := fmt.Sprintf(`{"version":"%s"}%s`, VERSION, "\n")
  assert.Equal(t, expectedBody, string(recorder.Body.Bytes()))
}

func TestCodesHandler(t *testing.T) {
  recorder := newTestRequest("GET", "/foo")

  assert.Equal(t, http.StatusOK, recorder.Code)
  assert.Equal(t, "image/png", recorder.HeaderMap.Get("Content-Type"))
}

func TestNotFound(t *testing.T) {
  recorder := newTestRequest("GET", "/foo/bar")

  assert.Equal(t, http.StatusNotFound, recorder.Code)
  assert.Equal(t, "text/html", recorder.HeaderMap.Get("Content-Type"))
  expectedBody := "404 page not found"
  assert.Equal(t, expectedBody, string(recorder.Body.Bytes()))
}
