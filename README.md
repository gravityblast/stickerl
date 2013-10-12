# Stickerl

A QRCode generator web service for personal url shortener, base on [Traffic](https://github.com/pilu/traffic).

## Installation

  go get github.com/pilu/stickerl

## Usage

  stickerl http://example.com

The above command starts the service on port 7000 with `http://example.com` as base redirect url.
This means that visiting http://localhost:7000/qrcodes/foo, the generated qrcode will have the url `http://example.com/foo`.

Use `stickerl -h` for more options.
