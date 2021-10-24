# dt (DeepL Translation)

A simple command line application to translate using deepl API.

## Installing

```bash
go get -u github.com/kinpoko/dt
```

## Usage

1. Get authentication key for Deepl API. See [here](https://www.deepl.com/ja/docs-api/).
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
