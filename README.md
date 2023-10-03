<p align="center">
  <img src="https://img.shields.io/github/workflow/status/globocom/prettylog/Go?style=flat-square">
  <img src="https://goreportcard.com/badge/github.com/globocom/prettylog?style=flat-square">
  <a href="https://github.com/globocom/prettylog/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/globocom/prettylog?color=blue&style=flat-square">
  </a>
  <img src="https://img.shields.io/github/go-mod/go-version/globocom/prettylog?style=flat-square">
  <a href="https://pkg.go.dev/github.com/globocom/prettylog">
    <img src="https://img.shields.io/badge/Go-reference-blue?style=flat-square">
  </a>
</p>

# Prettylog

Command line tool that displays JSON logs in a human-friendly format.

![Prettylog](https://github.com/globocom/prettylog/raw/master/prettylog.png)

## Installation

Go 1.18+:

    go install github.com/globocom/prettylog@latest

Go 1.17:

    go install github.com/globocom/prettylog

Go 1.16 and older:

    curl https://github.com/globocom/prettylog/raw/master/install.sh | sh

Prettylog will be installed to `$GOPATH/bin`. Make sure to add it to your `PATH` to run `prettylog` anywhere.
If the path is not set, it defaults to `$HOME/go/bin/prettylog`.

## How it works

Prettylog parses log messages that contain an arbitrary number of fields and generates a nice output in the
following format:

    <TIMESTAMP> <LOGGER> <CALLER> <LEVEL> <MESSAGE> <FIELD1>=<VALUE> <FIELD2>=<VALUE> ...

Non-existent fields will be ignored, and messages not encoded as JSON will be printed as is.

## Usage

Simply pipe the `stdout` of an application that outputs structured log messages into `prettylog`:

    app | prettylog

You might need to redirect `stderr` to `stdout` if the application doesn't log to the standard output:

    app 2>&1 | prettylog

## Configuration

You can configure how Prettylog works by creating a `.prettylog.yml` file either locally (per directory)
or globally (in `$HOME`):


```yaml
timestamp:
  key:     <string>
  visible: <bool>
  color:   <list of int>
  format:  <string>

logger:
  key:     <string>
  visible: <bool>
  padding: <int>
  color:   <list of int>

caller:
  key:     <string>
  visible: <bool>
  padding: <int>
  color:   <list of int>

level:
  key:     <string>
  visible: <bool>
  padding: <int>
  colors:
    debug: <list of int>
    info:  <list of int>
    warn:  <list of int>
    error: <list of int>
    fatal: <list of int>

message:
  key:     <string>
  padding: <int>
  color:   <list of int>
```

Each field has its own key and the following properties are available:

| Name | Description |
| - | - |
|**key**| Field name. |
|**visible**| Flag indicating whether the field will be printed. |
|**padding**| Number of whitespaces that will be added to the right of the field. |
|**color/colors**| Color attributes. Up to 3 values can be used (fg, bg and effects). More information [here](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors). |
|**format**| (timestamp field only) Layout that will be used to print timestamp values. It must follow the rules of the [time package](https://golang.org/pkg/time/#pkg-constants). |

## Using with other tools

Prettylog can be used along with other command line tools. Just make sure no buffer is enabled. For instance, `grep`
has a `--line-buffered` flag:

    app | grep --line-buffered -v debug | prettylog

If the tool you want to use buffers its output and does not offer such a flag, you can try
[stdbuff](https://www.gnu.org/software/coreutils/manual/html_node/stdbuf-invocation.html).
