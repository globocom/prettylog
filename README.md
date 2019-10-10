# Pretty Log

Tool for displaying structured logs in JSON in a human-compatible format.

![Prettylog](https://github.com/globocom/prettylog/raw/master/prettylog.png)

## Installation

    curl https://github.com/globocom/prettylog/raw/master/install.sh | sh 

Assuming the `$GOPATH/bin` folder is added to the current user's `PATH` folder, the application will be available for 
use immediately after installation.

## Operation

Prettylog processes logs containing an arbitrary number of fields, and produces a friendly output in the following format:

    <TIMESTAMP> <LOGGER> <CALLER> <LEVEL> <MESSAGE> <FIELD1>=<VALUE> <FIELD2>=<VALUE> ...

If a particular field does not exist in the log, it is ignored in the output generated.

**NOTE***: Currently only logs in JSON format are supported. Logs in other formats, or without any format, will be
printed without any changes.

## Usage

The tool is designed to directly read the `stdout` of an application that produces logs in structured format:

    app | prettylog

If the application writes logs in `stderr` instead of `stdout`, a redirection is required for the tool. 
function correctly:

    app 2>&1 | prettylog

## Configuration

The tool can be configured through the `.prettylog.yml` file, which can be located both locally (in the
folder where the tool is executed), when globally (in the `$HOME` folder). The file structure is as follows:

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

- **key**: Name of the field to be extracted from the application log.
- **visible**: Flag indicating whether the field will be displayed by the tool.
- **padding**: Number of blanks to be added to the right of the field text.
- **color/colors**: Color attributes used to colorize the field text. Up to 3 values can be entered 
(foreground, background e effects), according to the [ASCII color chart](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors).

## Use with other command line tools

Prettylog can be used in conjunction with other output processing tools, such as `grep`. However, 
In order for the output to be formatted correctly, you must turn off any buffers that are not per line. 
For example, with `grep` just use the `--line-buffered` option:

    app | grep --line-buffered -v debug | prettylog

If the tool makes use of a buffer and does not provide a native way to turn it off, try using the [stdbuff](https://www.gnu.org/software/coreutils/manual/html_node/stdbuf-invocation.html).
