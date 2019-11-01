# Pretty Log

Tool made to show logs structured in JSON in a format compatible with human beings.

![Prettylog](https://github.com/globocom/prettylog/raw/master/prettylog.png)

## Installation

    curl https://github.com/globocom/prettylog/raw/master/install.sh | sh

Assuming that the folder `$GOPATH/bin` is already added to `PATH` for the current user, the application will be available right after the installation.

## How it works

Pretty log processes logs containing an arbitrary number of fields and delivers a friendly output in this format:

    <TIMESTAMP> <LOGGER> <CALLER> <LEVEL> <MESSAGE> <FIELD1>=<VALUE> <FIELD2>=<VALUE> ...

If a certain field does not exist in the log, it will be ignored in the output.

**NOTE**: Currently only JSON logs are supported. Logs in other formats, or logs without any format will be printed without changes.

## Utilization

The tool was made to read directly from the `stdout` of a structured log application:

    app | prettylog

If the application writes logs in `stderr` instead of `stdout`, a redirection is necessary in order for the application to run properly:

    app 2>&1 | prettylog

## Setup

The tool can be configured using the file `.prettylog.yml`, which can be found locally (in the folder where the tool is run), or globally (in the folder `$HOME`. The file is structured as

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

Each key defines a field's log, and you can see each property below:

- **key**: Field name to be extracted from application log
- **visible**: Indicate which field will be shown from Prettylog
- **padding**: Number of white spaces that will be added after the text
- **color/colors**: Allow to color textfield. You can use up to 3 values (foreground, background and effects), according to [ASCII table](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors).
- **format**: Exclusive used for Timestamp which define date format to be shown on screen. This attribute must follow [Golang specification](https://golang.org/pkg/time/#pkg-constants) for date format.

## Utilization with other command line tools

Prettylog can be used alongside other output processing tools, such as `grep`. But, in order for the information to be output correctly, it is necessary to shut off any not lined buffer. For example, in `grep`you will only need to use the option `--line-buffered`:

    app | grep --line-buffered -v debug | prettylog

If the tool makes use of a buffer and there is no native option to turn it off, try using [stdbuff](https://www.gnu.org/software/coreutils/manual/html_node/stdbuf-invocation.html).
