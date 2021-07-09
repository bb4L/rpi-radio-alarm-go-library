# rpi-radio-alarm-go-library

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/bb4L/rpi-radio-alarm-go-library)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/bb4l/rpi-radio-alarm-go-library)
[![Go Reference](https://pkg.go.dev/badge/github.com/bb4L/rpi-radio-alarm-go-library.svg)](https://pkg.go.dev/github.com/bb4L/rpi-radio-alarm-go-library)
![GitHub](https://img.shields.io/github/license/bb4l/rpi-radio-alarm-go-library)
![GitHub Release Date](https://img.shields.io/github/release-date/bb4l/rpi-radio-alarm-go-library)
![GitHub last commit](https://img.shields.io/github/last-commit/bb4l/rpi-radio-alarm-go-library)
[![Go Report Card](https://goreportcard.com/badge/github.com/bb4L/rpi-radio-alarm-go-library)](https://goreportcard.com/report/github.com/bb4L/rpi-radio-alarm-go-library)
![GitHub issues](https://img.shields.io/github/issues-raw/bb4l/rpi-radio-alarm-go-library)
![Lines of code](https://img.shields.io/tokei/lines/github/bb4l/rpi-radio-alarm-go-library)

Wrapper library to communicate with the [rpi-radio-alarm-go](https://github.com/bb4L/rpi-radio-alarm-go)

## Installation
- get the package with `go get`
- ensure `./rpi_data.yaml` is available (if using the helper interacting with storage)
  ```yaml
  settings:
    port: 8000
    run_api: true
    run_telegram_bot: true
    run_discord_bot: false
  alarms:
  - name: Test
    hour: 7
    minute: 0
    days:
    - 0
    - 1
    active: false
  ...
  radio:
    running: false
    pid: -1

  ```

# License
[GPLv3](LICENSE)
