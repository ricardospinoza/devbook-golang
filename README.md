# devbook-golang
Rede social Dev Book feito em GoLang

Curso Aprenda Golang do Zero! Desenvolva uma APLICAÇÃO COMPLETA!, Udemy
https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa

# banco mysql

## scrips de criação do banco
    /sql
        criacao_banco.sql

##  dados conexao 
    na raiz arquivo '.env'
        DB_USUARIO=seu usuario de banco
        DB_SENHA=senha do usuario de banco
        DB_NOME=nome da base


## API
    na raiz arquivo '.env'

    API_PORTA=numero da porta que irá rodar a api

# Libs de apoio, instalar via "go get" + url seguintes:

    # tratamento de rotas
    github.com/gorilla/mux

    # leitura do arquivo .env
    github.com/joho/godotenv

    #drive sql
    github.com/go-sql-driver/mysql
    
    #validacao de email
    github.com/badoux/checkmail
    

