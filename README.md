# BooksList

A BooksList API é uma API RESTful desenvolvida em Go usando o framework Echo. Ela permite gerenciar uma lista de livros, possibilitando operações como adicionar, atualizar, deletar e listar livros. A aplicação é ideal para servir como um backend para aplicativos de gerenciamento de leitura ou bibliotecas. Criado na intenção e listar/gerenciar os livros que li. 

# Utilização do GO

A principal intenção em utilizar GO é aplicar tudo que venho aprendendo e observando sobre a linguagem.

# Pré-requisitos

Go 1.20+
Cliente HTTP para testes: Insomnia ou Postman

# Inicialização da API

-> Instalação

1. Clone o repositório:

-- git clone https://github.com/seu-usuario/BooksList.git --
cd BooksList

2. Instale as dependências:

-- go mod tidy --

3. Inicie o servidor:

-- go run main.go --
O servidor estará disponível em http://localhost:8080.

# Endpoints

1. Health Check

-- GET /health
Retorna o status de saúde a API

2. Books

-- GET /books/:id
Retorna um livro específico pelo ID
// GET http://localhost:8080/books/1

-- POST /books/
Retorna a inclusão de um livro.

-- PUT /books/
Atualiza um livro existente.

-- DELETE /books/:id
Deleta um livro existente.

# Estrutura do Projeto 

BooksList/
│
├── src/
│   ├── handler/
│   │   └── bookHandler.go
│   ├── model/
│   │   └── BookList.go
│   ├── service/
│      └── book_service.go
│   
│
├── main.go
├── go.mod
├── go.sum
└── README.md
