# Configuração e Execução

Este arquivo contém instruções para configurar e executar o aplicativo, incluindo a inicialização do banco de dados no Docker, a execução dos testes e a execução do arquivo principal (`main.go`).

## Pré-requisitos

Antes de começar, verifique se você tem os seguintes pré-requisitos instalados em sua máquina:

- Docker: [Instruções de instalação do Docker](https://docs.docker.com/get-docker/)
## Configuração e Instalação

1. Clone este repositório em sua máquina local:

```bash
git clone <git@github.com:VictorOliveiraPy/data-manipulation-v2.git>
cd <NOME_DO_DIRETORIO>
```

Para buildar o projeto e executar o aplicativo, siga as etapas abaixo:

Execute o seguinte comando para buildar o projeto:
```bash
make build
````
Em seguida, execute o comando a seguir para iniciar o aplicativo

```bash
make run
````


## Execução dos Testes

Para executar os testes, utilize o seguinte comando:

```bash
make test
```

Para remove o container e volume
```bash
make remove
```