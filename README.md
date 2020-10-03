# Pretty Log

[![GoDoc](https://godoc.org/github.com/globocom/prettylog?status.svg)](https://godoc.org/github.com/globocom/prettylog)
[![Go Report Card](https://goreportcard.com/badge/github.com/globocom/prettylog)](https://goreportcard.com/report/github.com/globocom/prettylog)

Application to display JSON logs structured in a format understandable by humans.

![Prettylog](https://github.com/globocom/prettylog/raw/master/prettylog.png)

## Installation

    curl https://github.com/globocom/prettylog/raw/master/install.sh | sh 

Assuming that the `$GOPATH/bin` folder is added to the `PATH` of the current user, the application will be available for use immediately after installation.


## Running

Prettylog processes logs containing an arbitrary number of fields, and produces friendly output in the following format:

    <TIMESTAMP> <LOGGER> <CALLER> <LEVEL> <MESSAGE> <FIELD1>=<VALUE> <FIELD2>=<VALUE> ...

If a given field does not exist in the log, it will be ignored in the output generated.

**NOTE**: Currently only logs in JSON format are supported. Logs in other formats, or without any format, will be printed without any modification.

## Usage

The tool was designed to directly read the `stdout` of an application that produces logs in a structured format:

    app | prettylog

If the application writes logs to `stderr` instead of `stdout`, a redirect is required for the tool function properly:

    app 2>&1 | prettylog

## Configuration

The tool can be configured through the `.prettylog.yml` file, which can be located either locally (in the folder where the tool runs), when globally (in the `$HOME` folder). The file structure is as follows:

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

Each key configures the formatting of a field in the log, and the meaning of each property is described below:

- **key**: Name of the field to be extracted from the application log.
- **visible**: Flag indicating whether the field will be displayed by the tool.
- **padding**: Number of blanks to be added to the right of the field text.
- **color/colors**: Color attributes used to color the field text. Up to 3 values ​​can be informed (foreground, background and effects), according to the [table for ASCII colors](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors).
- **format**: Unique attribute for Timestamp that defines the date format to be displayed on the screen. The value of this attribute you must follow the [Go language specifications](https://golang.org/pkg/time/#pkg-constants) for date format.

## Use with other command line tools

Prettylog can be used in conjunction with other output processing tools, such as `grep`. However, for the output to be formatted correctly, it is necessary to turn off any buffer that is not per line. For example, with `grep` just use the `--line-buffered` option:

    app | grep --line-buffered -v debug | prettylog

If the tool uses a buffer and does not provide a native way to turn it off, try using the [stdbuff](https://www.gnu.org/software/coreutils/manual/html_node/stdbuf-invocation.html).