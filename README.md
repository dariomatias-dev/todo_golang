# API de lista de tarefas feita em Golang

Bem-vindo à TaskList API, uma aplicação de lista de tarefas construída em Go utilizando os frameworks Gin para a camada web, SQLC para acesso ao banco de dados e Docker para facilitar a execução e implantação.

## Tecnologias usadas

- [Go](https://go.dev/)
- [Gin](https://gin-gonic.com/)
- [SQLC](https://sqlc.dev/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)

## Requisitos

Certifique-se de ter os seguintes itens instalados antes de iniciar:

- Go (versão utilizada no projeto: 1.21.4)
- Docker
- Imagem do PostgreSQL baixada
- SQLC (foi instalado na maquina)

## Instalação e Execução

Clone o repositório:

```bash
git clone https://github.com/dariomatias-dev/todo_golang
cd todo_golang
```

Abra o arquivo `.env_example` e remova do seu nome `_example`, depois preencha os campos que estão faltando.

Instale as depedências:

```bash
go mod download
```

Crie o container com o banco de dados:

```bash
docker-compose up -d
```

Obs.: Se estiver no Linux e tiver dado erro, provavelmente irá precisar colocar `sudo` antes do comando.

Rode a aplicação:

```bash
go run main.go
```

## Endpoints

A API fornece os seguintes endpoints:

- POST `/todo`: Criação de uma nova tarefa. Por padrão a tarefa é criada com o valor `false` no campo `status`.
- GET `/todos/:id`: Retorna a tarefa do ID enviado, caso não exista será retornado o valor `null`.
- GET `/todos`: Retorna todas as tarefas, caso não exista nenhuma no banco de dados será retornado `null`.
- PATCH `/todos/:id`: Atualiza uma tarefa existente.
- PATCH `/todo-status/:id`: Atualiza o status da tarefa.
- DELETE `/todos/:id`: Exclui uma tarefa.

## Licença

Este projeto é licenciado sob a Licença MIT.

Divirta-se usando a To-Do API construída em Go!
