package main

import (
  "log"
  "fmt"
  "strings"
  "net/http"
  "image/png"
  "encoding/json"
  "github.com/pilu/traffic"
  "github.com/qpliu/qrencode-go/qrencode"
)

const VERSION = "0.1.0"

var baseUrl string

func init() {
  var ok bool
  baseUrl, ok = traffic.GetVar("base_redirect_url").(string)
  if !ok || baseUrl == "" {
    log.Fatal("Base Redirect URL is blank. Set the environment variable TRAFFIC_BASE_REDIRECT_URL")
  }

  if !strings.HasSuffix(baseUrl, "/") {
    baseUrl = fmt.Sprintf("%s/", baseUrl)
  }
}

func rootHandler(w traffic.ResponseWriter, r *http.Request) {
  response := map[string]string{
    "version": VERSION,
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func qrcodesHandler(w traffic.ResponseWriter, r *http.Request) {
  code  := r.URL.Query().Get("code")
  url   := fmt.Sprintf("%s%s", baseUrl, code)

  grid, err := qrencode.Encode(url, qrencode.ECLevelQ)
  if err != nil {
    panic(err)
  }

  w.Header().Set("Content-Type", "image/png")
  png.Encode(w, grid.Image(8))
}

func main() {
  app := traffic.New()
  app.Get("/", rootHandler)
  app.Get("/qrcodes/:code", qrcodesHandler)
  app.Run()
}
