# Pretty Log

Tool for exhibiting structured logs.

![Prettylog](https://github.com/globocom/prettylog/raw/master/prettylog.png)

_Ler em [Português](README.md)_

## Installation

    curl https://github.com/globocom/prettylog/raw/master/install.sh | sh 

Considering folder `$GOPATH/bin` is in current user `PATH`, the application will be available for use right after
installation.

## How does it work

Prettylog process logs with an arbitrary number of fields, resulting in a friendly output in the following format:

    <TIMESTAMP> <LOGGER> <CALLER> <LEVEL> <MESSAGE> <FIELD1>=<VALUE> <FIELD2>=<VALUE> ...

If a certain field does not exist in the log, it is ignored in the generated output.

**NOTE**: Currently, only logs in JSON format are supported. Logs in other formats, or without any format at all, will
be printed with no changes.

## How to use

The tool is projected to read directly the `stdout` of an application which produces logs in structured format:

    app | prettylog

May the application write logs to `stderr` instead of `stdout`, redirecting is necessary for the tool to work properly:

    app 2>&1 | prettylog

## Configuration

The tool read its configurations through the `.prettylog.yml` file, localized both locally (in the same folder where the tool
is executed) and globally (in the `$HOME` folder). The file structure is the following:

    timestamp:
      key:     <string>
      visible: <bool> 
      color:   <list of int>

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

Each key configures the formatting for a log field, and the meaning of each property is described bellow:

- **key**: Name of the field to be extracted from application log.
- **visible**: Flag to indicate if the field will be shown by the tool.
- **padding**: Amount of blank spaces to be added on the right of the field's text.
- **color/colors**: Color attributes to be used to color the field's text. Up to 3 values can be informed (foreground,
background and effects), following the [ASCII color table](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors).

## Using with other command line tools

Prettylog can be used together with other output processing tools, like `grep`. However, in order for the output
formatting to work correctly, it is necessary to shut down any non line buffered buffer. For example, with `grep`, it is
necessary to use the `--line-buffered` option:

    app | grep --line-buffered -v debug | prettylog

If the tool uses a buffer with no native way of shutting it down, try using the
[stdbuff](https://www.gnu.org/software/coreutils/manual/html_node/stdbuf-invocation.html).
