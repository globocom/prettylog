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

*Leia em outros idiomas: [Inglês](README.md), [Português](README.pt-br.md).*

Ferramenta para exibição de logs estruturados em JSON em formato compatível com seres humanos.

![Prettylog](https://github.com/globocom/prettylog/raw/master/prettylog.png)

## Instalação

Go 1.17+:

    go install github.com/globocom/prettylog

Go 1.16 ou inferior:

    curl https://github.com/globocom/prettylog/raw/master/install.sh | sh

O Prettylog será instalado em `$GOPATH/bin`. Certifique-se de adicioná-lo ao seu `PATH` para que você possa executar o `prettylog` em qualquer lugar.

## Funcionamento

Prettylog processa logs contendo um número arbitrário de campos, e produz uma saída amigável no seguinte formato:

    <TIMESTAMP> <LOGGER> <CALLER> <LEVEL> <MESSAGE> <FIELD1>=<VALUE> <FIELD2>=<VALUE> ...

Os campos inexistentes serão ignorados e os logs não formatadas como JSON serão impressos sem nenhuma modificação.

## Utilização

A ferramenta foi projetada para ler diretamente o `stdout` de uma aplicação que produza logs em formato estruturado:

    app | prettylog

Se a aplicação escrever logs no `stderr` ao invés do `stdout`, um redirecionamento é necessário para a ferramenta funcionar corretamente:

    app 2>&1 | prettylog

## Configuração

Você pode configurar o funcionamento do Prettylog criando um arquivo `.prettylog.yml` localmente (diretório onde a ferramenta é executada) ou globalmente (em `$HOME`):


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

Cada campo tem sua própria chave e as seguintes propriedades estão disponíveis:

| Nome | Descrição |
| - | - |
|**key**| Nome do campo. |
|**visible**| Atributo que indica se o campo será impresso. |
|**padding**| Número de espaços em branco que serão adicionados à direita do campo. |
|**color/colors**| Atributos de cor. Podem ser usados ​​até 3 valores (fg, bg e effects). Mais Informações [aqui](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors). |
|**format**| (exclusivo para Timestamp) define o formato da data a ser exibido na tela. O valor desse atributo deve seguir as [especificações da linguagem Go](https://golang.org/pkg/time/#pkg-constants) para formato de datas. |

## Utilização com outras ferramentas de linha de comando

Prettylog pode ser utilizado em conjunto com outras ferramentas de processamento de output, como o `grep`. Entretanto, para que a formatação da saída seja feita corretamente, é necessário desligar qualquer buffer que não seja por linha. Por exemplo, com o `grep` basta utilizar a opção `--line-buffered`:

    app | grep --line-buffered -v debug | prettylog

Se a ferramenta fizer uso de um buffer e não fornecer uma forma nativa de desligá-lo, tente usar o
[stdbuff](https://www.gnu.org/software/coreutils/manual/html_node/stdbuf-invocation.html).
