# Stickerl

A QRCode generator web service for personal url shortener, based on the [Traffic web framework](https://github.com/pilu/traffic) and
[qpliu/qrencode-go](https://github.com/qpliu/qrencode-go) library.

## Installation

    go get github.com/pilu/stickerl

## Usage

    stickerl http://example.com

The above command starts the service on port 7000 with `http://example.com` as base redirect url.
This means that visiting http://localhost:7000/foo, the generated qrcode will have the url `http://example.com/foo`.

Use `stickerl -h` for more options.
