# GophKeeper password manager

<img src="/assets/31.svg" width="256" alt="gopher">

Author of the illustration - [link](https://github.com/MariaLetta/free-gophers-pack).

## Server

### Migrations

1. [Install](https://github.com/pressly/goose#install) `goose` database migration tool.

2. [Apply](https://github.com/pressly/goose#usage) all migrations:

    ```sh
    GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres:postgres@localhost:5432/gophkeeper?sslmode=disable" \
    make migrate-up
    ```

Check migration status:

```sh
GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres:postgres@localhost:5432/gophkeeper?sslmode=disable" \
make migrate-status
```

To [roll back](https://github.com/pressly/goose#usage) migrations, run the following command:

```sh
GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres:postgres@localhost:5432/gophkeeper?sslmode=disable" \
make migrate-down
```
