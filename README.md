# Pretty Log

A tool to display structured logs in JSON, in a human-compatible format.

![Prettylog](https://github.com/globocom/prettylog/raw/master/prettylog.png)

## Installation

    curl https://github.com/globocom/prettylog/raw/master/install.sh | sh 

Assuming the `$ GOPATH / bin` folder is included in the current user's `PATH`, the app will be available for
use immediately after installation.

## Functionality

Prettylog processes logs that contain an arbitrary number of fields and produces friendly results in the following format:

    <TIMESTAMP> <LOGGER> <CALLER> <LEVEL> <MESSAGE> <FIELD1>=<VALUE> <FIELD2>=<VALUE> ...

If a given field does not exist in the log, it will be ignored in the generated output.

**NOTA**: Currently, only JSON format logs are supported. Logs in other formats, or without any format, will be printed without any modification.

## utilization

The tool is designed to read directly from an application's `stdout` which produces logs in structured format:

    app | prettylog

If the application writes logs to `stderr` instead of` stdout`, a redirect is required for the tool to work properly:

    app 2>&1 | prettylog

## Settings

The tool can be configured through the `.prettylog.yml` file, which can be located either locally (in the folder where the tool runs) or globally (in the `$ HOME` folder). The file structure is as follows:

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

Each key configures the formatting of a log field, and the meaning of each property is described below:

- **key**:Name of the field to extract from the application log.
- **visible**: Flag indicating if the field will be displayed by the tool.
- **padding**: Amount of whitespace to add to the right of the field text.
- **color/colors**: Color attributes used to color the field text. Up to 3 values can be entered (foreground, background e effects), according to [tabela para cores ASCII](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors).

## Using with other command-line tools

Prettylog can be used in conjunction with other output processing tools such as `grep`. however,
For the formatting of the output to be done correctly, it is necessary to turn off any non-line buffer.
For example, with `grep` just use the option `--line-buffered`:

    app | grep --line-buffered -v debug | prettylog

If the tool makes use of a buffer and does not provide a native way to turn it off, try using the  
[stdbuff](https://www.gnu.org/software/coreutils/manual/html_node/stdbuf-invocation.html).
