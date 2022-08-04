# buying-frenzy

### Run

#### Setting Up DB schema  (Required)
1. Run this command to run a postgresql instance. Connect the db with table plus or something.
    ```shell
    docker-compose up postgres -d
    ```
2. Copy all the queries from `/infrastructure/db/migrations/000001_create_basic_tables.up.sql`
3. Run the queries at once to create all the tables.

#### ETL
Run this command to run a postgresql instance in docker container and the etl command of the app.
```shell
make etl-local
```

#### Server
```shell
make development-serve
```
