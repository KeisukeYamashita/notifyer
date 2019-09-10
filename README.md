# Notifyer

[![License][license-badge]][license]
[![GoDoc][godoc-badge]][godoc]
[![GolangCI][golangci-badge]][golangci]
[![Go Report Card][go-report-card-badge]][go-report-card]


<br />
<p align="center"><a href="#" target="_blank" rel="noopener noreferrer"><img width="320"src="_docs/img/notifyer.png"></a></p>
<br />

> Extendable notification tool for various services

`notifyer` is a notification tool written in golang. It will read stdin and sends notification to services such as Slack, Discord and so on. You will have to configure by writting TOML file. Don't worry, it is very very simple to configure and to use.

## Install

### homebrew

```
$ brew tap KeisukeYamashita/notifyer
$ brew install notifyer
```

### git

```
$ git clone https://KeisukeYamashita/notifyer
$ make install
```

## Sample Usage

Post message to slack.

```
cat message.txt | notifyer slack
```

See below for detailed usage and supported services.

## Supports

<br />
<p align="center"><a href="#" target="_blank" rel="noopener noreferrer"><img width="100%"src="_docs/img/notifyer_app.png"></a></p>
<br />

### Input

Currently supports sdtin only.

### Output

<table style="width:100%;text-align:left">
  <tr>
    <th>Service name</th>
    <th>Status</th>
  </tr>
  <tr>
    <td>Discord</td>
    <td>○</td>
  </tr>
  <tr>
    <td>Slack</td>
    <td>○</td>
  </tr>
  <tr>
    <td>Line</td>
    <td>○</td>
  </tr>
</table>

## Usage

### Config file

The script will check the following files for configuration.

1. `NOTIFYER_PATH` env var
1. `$HOME/.notify_discord`

The config toml file should be in this format.

```toml
[slack]
url         = "hogehoge"
channel     = "poip"

[linebot]
accessToken = ""
to          = ""

[discord]
url         = "https://"
```

### Slack

```
echo "Hello fron Notifyer" | notifyer slack
```

###  Line

```
echo "Hello fron Notifyer" | notifyer linebot
```

### Discord

```
echo "Hello fron Notifyer" | notifyer discord
```

### Docker mode

Also docker images are supported.

```
echo "Hello from Notifyer" | docker run  -v $HOME/.notifyer.toml:/root/.notifyer.toml 1915keke/notifyer:0.1.0 slack 
```

## Contribution

* I welcome your contributions.
    * Please send issues and pull requests.

## License

Notifyer is released under the MIT license.

© 2019 GitHub, Inc.

## Author

* [KeisukeYamashita](https://github.com/KeisukeYamashita)

<!-- badge links -->

[license]: LICENSE
[godoc]: https://godoc.org/github.com/KeisukeYamashita/notifyer
[go-report-card]: https://goreportcard.com/report/github.com/KeisukeYamashita/notifyer
[golangci]: https://golangci.com/r/github.com/KeisukeYamashita/notifyer

[license-badge]: https://img.shields.io/badge/license-Apache%202.0-%23E93424
[circleci-badge]: https://img.shields.io/circleci/project/github/micnncim/protocol-buffers-language-server.svg?label=circleci&logo=circleci
[godoc-badge]: https://img.shields.io/badge/godoc.org-reference-blue.svg
[go-report-card-badge]: https://goreportcard.com/badge/github.com/KeisukeYamashita/notifyer
[golangci-badge]: https://golangci.com/badges/github.com/KeisukeYamashita/notifyer