# Wallet Sandbox 

This is a test bed/sandbox project that encompases technologies needed/used to create a modern digital wallet system. 

This should ideally only be used in development by its developers. Ideas are meant to be tried out, this project allows (encourages) rewriting or reinventing as the case may be.  

## Setup

To setup your dev environment the following dependencies are required: 

- Docker
- PostgresDB
- make 
- Go 1.18

That's it. In order to startup the server the database needs to be populated. 

- `make setup-db` creates the wallet schema using the sql script `db/setup.sql`
- `make teardow-db` removes the wallet schema along with all related objects i.e tables, functions etc.



To spin up the web and db containers use:

```bash
$ docker compose up -d
```

To stop web and db containers use: 

```bash
$ docker compose stop
```

## Test 

Run tests using  `make test`

## Contribution 

To contribute, please create an issue or PR and it will be merged or looked over by @beesaferoot. 