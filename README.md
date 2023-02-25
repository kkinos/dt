# dt (DeepL Translation)

[![Go](https://github.com/kinpoko/dt/actions/workflows/go.yml/badge.svg)](https://github.com/kinpoko/dt/actions/workflows/go.yml)
![License](https://img.shields.io/github/license/kinpoko/dt?color=blue)

A simple command line application to translate with DeepL API.

This application only supports DeepL API Free, and does not currently support other APIs.

## Installing

```bash
go install github.com/kinpoko/dt@latest
```

## Usage

1. Obtain an authentication key for the DeepL API. For information on how to obtain an authentication key, please see the [official documentation](https://www.deepl.com/ja/docs-api/).
2. Set the authentication key as an environment variable named `DEEPL_TOKEN`.

`.bashrc`

```bash
export DEEPL_TOKEN=<authentication key>
```

3. Translate text using the following command:

```bash
dt Hello World
ハロー・ワールド
```

Use the `-h` flag to view the help menu and command line options:

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

You can use `dt` with other command line utilities by piping text to the application. For example, to translate the contents of a file named text.txt, you can use the following command:

```bash
cat text.txt | xargs dt
```

This will translate the text in the `text.txt` file and output the resulting translation to the command line.
