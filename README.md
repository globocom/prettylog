# Pretty log
A tool for displaying JSON structured logs in a way that is compatible with human beings.

![Prettylog](https://github.com/globocom/prettylog/raw/master/prettylog.png)

## Installing
    curl https://github.com/globocom/prettylog/raw/master/install.sh | sh 

Assuming that the folder `$GOPATH/bin` is added to the current user's `PATH`, the application will be available immediately after installation.

## Functionality

Prettylog processes logs containing an arbitrary number of fields and produces a friendly output in the following format:

    <TIMESTAMP> <LOGGER> <CALLER> <LEVEL> <MESSAGE> <FIELD1>=<VALUE> <FIELD2>=<VALUE> ...

If a certain field doesn't exist in the log, it will be ignored in the generated output.

**NOTE**: Currently, only logs in JSON format are supported. Logs in other formats, or without any formatting, will be print without any modification.

## Usage

The tool was made to read directly the `stdout` of an application that writes logs in a structured format:

    app | prettylog

If an application writes logs on `stderr` instead of `stdout`, a redirect is necessary for the tool to work correctly:

    app 2>&1 | prettylog

## Configuration

The tool can be configurated through the `.prettylog.yml` file, which can be found both locally (on the directory which the tool is executed), and globally (on the `$HOME` folder). The folder structure is as follows:

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

Each key sets the formatting of a field on the log, and the meaning of each property is described below:

- **key**: Name of the field to be extracted from the application's log.
- **visible**: Flag indicating if the field will be shown by the tool.
- **padding**: Quantity of whitespaces to be added to the right of the field's text.
- **color/colors**: Color attributes used to color the field's text. Up to 3 values can be set (foreground, background and effects), according to the [ASCII color table.](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors).

## Utilization with other command-line tools

Prettylog can be utilized together with other output processing tools, such as `grep`. However, for the output's formatting to be done correctly, it is required to turn off any buffer that doesn't work by line. For example, with `grep`, you just need to use the `--line-buffered` flag:

    app | grep --line-buffered -v debug | prettylog

If the tool doesn't make use of a buffer and doesn't provide a way to turn it off, try using [stdbuff](https://www.gnu.org/software/coreutils/manual/html_node/stdbuf-invocation.html).


