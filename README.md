# Pretty Log

This tool was created to exhibit structured JSON logs to humans view.

!Ferramenta para exibição de logs estruturados em JSON em formato compatível com seres humanos.

![Prettylog](https://github.com/globocom/prettylog/raw/master/prettylog.png)

## Install

    curl https://github.com/globocom/prettylog/raw/master/install.sh | sh 

Assumindo que a pasta `$GOPATH/bin` esteja adicionada ao `PATH` do usuário atual, a aplicação ficará disponível para 
utilização imediatamente após a instalação.

## Runn

Prettylog processa logs contendo um número arbitrário de campos, e produz uma saída amigável no seguinte formato:

    <TIMESTAMP> <LOGGER> <CALLER> <LEVEL> <MESSAGE> <FIELD1>=<VALUE> <FIELD2>=<VALUE> ...

Se um determinado campo não existir no log, ele será ignorado na saída gerada.

**NOTA**: Atualmente apenas logs no formato JSON são suportados. Logs em outros formatos, ou sem formato algum, serão
impressos sem nenhuma modificação.

## Execution

A ferramenta foi projetada para ler diretamente o `stdout` de uma aplicação que produza logs em formato estruturado:

    app | prettylog

Se a aplicação escrever logs no `stderr` ao invés do `stdout`, um redirecionamento é necessário para a ferramenta 
funcionar corretamente:

    app 2>&1 | prettylog

## Configuration

A ferramenta pode ser configurada através do arquivo `.prettylog.yml`, que pode estar localizado tanto localmente (na
pasta onde a ferramenta é executada), quando globalmente (na pasta `$HOME`). A estrutura do arquivo é a seguinte:

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

Cada chave configura a formatação de um campo do log, e o significado de cada propriedade é descrito abaixo:

- **key**: Nome do campo a ser extraído do log da aplicação.
- **visible**: Flag indicando se o campo será exibido pela ferramenta.
- **padding**: Quantidade de espaços em branco a serem adicionados à direita do texto do campo.
- **color/colors**: Atributos de cor usados para colorir o texto do campo. Até 3 valores podem ser informados 
(foreground, background e effects), de acordo com a [tabela para cores ASCII](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors).
- **format**: Atributo exclusivo para Timestamp que define o formato da data a ser exibido na tela. O valor desse atributo
deve seguir as [especificações da linguagem Go](https://golang.org/pkg/time/#pkg-constants) para formato de datas.

## Use with other command line tools

Prettylog pode ser utilizado em conjunto com outras ferramentas de procesamento de output, como o `grep`. Entretanto, 
para que a formatação da saída seja feita corretamente, é necessário desligar qualquer buffer que não seja por linha. 
Por exemplo, com o `grep` basta utilizar a opção `--line-buffered`:

    app | grep --line-buffered -v debug | prettylog

Se a ferramenta fizer uso de um buffer e não fornecer uma forma nativa de desligá-lo, tente usar o 
[stdbuff](https://www.gnu.org/software/coreutils/manual/html_node/stdbuf-invocation.html).
