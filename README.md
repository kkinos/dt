# dt (DeepL Translation)

[![Go](https://github.com/kinpoko/dt/actions/workflows/go.yml/badge.svg)](https://github.com/kinpoko/dt/actions/workflows/go.yml)
![License](https://img.shields.io/github/license/kinpoko/dt?color=blue)

A simple command line application to translate with DeepL API.

This application only support for DeepL API Free.

## Installing

```bash
go install github.com/kinpoko/dt@latest
```

## Usage

1. Get authentication key for DeepL API. See [here](https://www.deepl.com/ja/docs-api/).
2. Set the key as `DEEPL_TOKEN`.

`.bashrc`

```bash
export DEEPL_TOKEN=<authentication key>
```

3. Translate.

```bash
dt Hello World
ハロー・ワールド
```

```bash
dt -h
A simple command line application to translate with deepl API.

Usage:
  dt [text] [flags]

Flags:
  -h, --help            help for dt
  -s, --source string   source language   (default "en")
  -t, --target string   target language   (default "ja")
```

## Using Pipe

e.g.

```bash
cat text.txt | xargs dt
```
