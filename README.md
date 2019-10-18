# Pretty Log

Displaying tool for JSON structured logs in compatible format to humans.

![Prettylog](https://github.com/globocom/prettylog/raw/master/prettylog.png)

## Installation

    curl https://github.com/globocom/prettylog/raw/master/install.sh | sh

Assuming that the folder `$GOPATH/bin` is inside the user's actual `PATH`, the application will be available after the installation.

## Running

Prettylog processes logs containing an arbitrary number of fields, and creates a friendly output in the following format:

    <TIMESTAMP> <LOGGER> <CALLER> <LEVEL> <MESSAGE> <FIELD1>=<VALUE> <FIELD2>=<VALUE> ...

If a specific field does not exist in the log, it will be ignored on the output.

**NOTE**: Nowadays, only logs on JSON format are supported. Logs in another format, or without one, will be printed without any modification

## Utilization

The tool was designed to read directly the `stdout` from an application that produces logs in a structured format:

    app | prettylog

If the application writes logs on `sterr` instead of `stdout`, a redirect is required so the tool can work properly:

    app 2>&1 | prettylog

## Settings

The tool can be configured through the `.prettylog.yml` file, which can be found either locally(on the folder that the tool runs) or globally
(in the folder `$HOME`). The file structure is as follows:

timestamp:
key: <string>
visible: <bool>
color: <list of int>

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

Each key configures the format of the field in the log, and the meaning of each property is described bellow:

- **key**: Name of the field that will be extracted from the log of the application.
- **visible**: Flag indicating whether the field will be displayed by the tool or not.
- **padding**: Amount of blank spaces that will be added to the right of the field text.
- **color/colors**: Color attributes used for coloring the field text. Up to 3 values ​​can be entered (foreground, background e effects)
  according to [ASCII colors chart](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors).

## Utilization with other command line tools

Prettylog can be used with other processing output tools like `grep`. However, for the formatting of the output to be done correctly,
it is necessary to turn off any non-line buffer.
For example, with `grep` just use the `--line-buffered` option:

    app | grep --line-buffered -v debug | prettylog

If the tool makes use of a buffer and does not provide a native way to turn it off,
try using the [stdbuff](https://www.gnu.org/software/coreutils/manual/html_node/stdbuf-invocation.html).
