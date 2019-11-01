# Pretty Log

This tool was created to exhibit structured JSON logs to humans view.

[comment]: # (Ferramenta para exibição de logs estruturados em JSON em formato compatível com seres humanos.)

![Prettylog](https://github.com/globocom/prettylog/raw/master/prettylog.png)

## Install

    curl https://github.com/globocom/prettylog/raw/master/install.sh | sh 

With the `$GOPATH/bin` folder configured in current user `PATH` the application will be available for immediately use after installation.

[comment]: # (Assumindo que a pasta `$GOPATH/bin` esteja adicionada ao `PATH` do usuário atual, a aplicação ficará disponível para 
utilização imediatamente após a instalação.)

## Runn

Prettylog processes logs containing an arbitrary number of fields, and produces friendly output in the following format: 

[comment]: # (Prettylog processa logs contendo um número arbitrário de campos, e produz uma saída amigável no seguinte formato:)

    <TIMESTAMP> <LOGGER> <CALLER> <LEVEL> <MESSAGE> <FIELD1>=<VALUE> <FIELD2>=<VALUE> ...

If a given field does not exist in the log, it will be ignored in the generated output.

[comment]: # (Se um determinado campo não existir no log, ele será ignorado na saída gerada.)

**NOTE**: Nowadays only JSON format logs are supported. Logs in other formats, or without any format, will be printed without modification.

<!-- **NOTA**: Atualmente apenas logs no formato JSON são suportados. Logs em outros formatos, ou sem formato algum, serão
impressos sem nenhuma modificação.) -->

## Execution

The Prettylog tool is designed to directly read `stdout` from an application that produces logs in structured format:

[comment]: # (A ferramenta foi projetada para ler diretamente o `stdout` de uma aplicação que produza logs em formato estruturado:)

    app | prettylog

If the application writes logs to `stderr` instead of` stdout`, a redirect is required for the tool to work correctly:

<!-- Se a aplicação escrever logs no `stderr` ao invés do `stdout`, um redirecionamento é necessário para a ferramenta 
funcionar corretamente:
-->

    app 2>&1 | prettylog

## Configuration

The Prettylog tool could be configured from the `.prettylog.yml` file, which can be located either locally (in the folder where the tool runs) or globally (in the `$ HOME` folder). The file structure is as follows:

<!-- A ferramenta pode ser configurada através do arquivo `.prettylog.yml`, que pode estar localizado tanto localmente (na
pasta onde a ferramenta é executada), quando globalmente (na pasta `$HOME`). A estrutura do arquivo é a seguinte:
-->

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

Each key configures the formatting of a log field, and the meaning of each property is described below:

[comment]: # (Cada chave configura a formatação de um campo do log, e o significado de cada propriedade é descrito abaixo:)

**key**: the field name to be extracted from the application log.

**visible**: Flag indicating if the field will be displayed by the tool.

**padding**: Amount of whitespace to add to the right of the field text.

**color / colors**: Color attributes used to color the field text. Up to 3 values ​​can be entered.
(foreground, background, and effects) according to the [ASCII color chart] (https://en.wikipedia.org/wiki/ANSI_escape_code#Colors).

**format**: Unique attribute for Timestamp that defines the date format to be displayed on the screen. The value of this attribute
must follow the [Go language specifications] (https://golang.org/pkg/time/#pkg-constants) for date format.




<!--
**key**: Nome do campo a ser extraído do log da aplicação.
- **visible**: Flag indicando se o campo será exibido pela ferramenta.
- **padding**: Quantidade de espaços em branco a serem adicionados à direita do texto do campo.
- **color/colors**: Atributos de cor usados para colorir o texto do campo. Até 3 valores podem ser informados 
(foreground, background e effects), de acordo com a [tabela para cores ASCII](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors).
- **format**: Atributo exclusivo para Timestamp que define o formato da data a ser exibido na tela. O valor desse atributo
deve seguir as [especificações da linguagem Go](https://golang.org/pkg/time/#pkg-constants) para formato de datas.
-->

## Use with other command line tools

Prettylog can be used in conjunction with other output processing tools such as `grep`. However, in order to format the output correctly, it is necessary to turn off any non-line buffer.
For example, with `grep` just use the` --line-buffered` option:

<!-- Prettylog pode ser utilizado em conjunto com outras ferramentas de procesamento de output, como o `grep`. Entretanto, 
para que a formatação da saída seja feita corretamente, é necessário desligar qualquer buffer que não seja por linha. 
Por exemplo, com o `grep` basta utilizar a opção `--line-buffered`:) -->

    app | grep --line-buffered -v debug | prettylog

If the tool makes use of a buffer and does not provide a native way to turn it off, try using the [stdbuff](https://www.gnu.org/software/coreutils/manual/html_node/stdbuf-invocation.html).
