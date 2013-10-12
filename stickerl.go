package main

import (
  "os"
  "fmt"
  "flag"
  "strings"
  "net/http"
  "image/png"
  "encoding/json"
  "github.com/pilu/traffic"
  "github.com/qpliu/qrencode-go/qrencode"
)

const VERSION = "0.1.0"

var (
  baseUrl string
  env     string
  app     *traffic.Router
)

func usage() {
  fmt.Println("USAGE:")
  fmt.Printf("  %s [OPTIONS] BASE_URL\n", os.Args[0])
  flag.PrintDefaults()
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
  flag.StringVar(&env, "e", "production", "Traffic ENV")
  flag.Parse()
  args := flag.Args()

  if len(args) != 1 {
    usage()
    os.Exit(1)
  }

  baseUrl = args[0]
  if !strings.HasSuffix(baseUrl, "/") {
    baseUrl = fmt.Sprintf("%s/", baseUrl)
  }

  traffic.SetVar("env", env)

  app = traffic.New()
  app.Get("/", rootHandler)
  app.Get("/qrcodes/:code", qrcodesHandler)
  app.Run()
}
