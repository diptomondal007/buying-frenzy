# buying-frenzy

### Deployment
Find the deployed app [here](https://buying-frenzy-dipto.fly.dev) .

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
This app will be running on port 8080.
```shell
make development-serve
```

### Api
#### Open Restaurants
> /api/v1/restaurant/open

Query Params:
> optional
> * date_time (ex - date_time=04/07/2019%2010:48%20PM)

##### Response - 200
```json
{
  "success": true,
  "message": "request successful!",
  "status_code": 200,
  "data": [
    {
      "id": "29881887-10a4-4d50-a573-112973dfe4fd",
      "name": "Danton's Gulf Coast Seafood Kitchen"
    },
    {
      "id": "5566994d-e97c-4f3c-98fd-5e78e497ee23",
      "name": "The Local House"
    }
  ]
}
```

##### Response - 400
```json
{
   "success": false,
   "message": "bad format of date time",
   "status_code": 400
}
```

#### Open Restaurants
> api/v1/restaurant/list?less_than=2&price_low=20&price_high=400

Query Params:
> Optional
> * less_than (ex - 2)
> * more_than (ex - 10)

> Required
> * price_low (ex-20) // low range of price
> * price_high (ex-400) // high range of price

##### Response - 200
```json
{
   "success": true,
   "message": "request successful!",
   "status_code": 200,
   "data": [
      {
         "id": "a069a467-fefd-41e8-a3f8-555e32e0ddb4",
         "name": "2G Japanese Brasserie"
      },
      {
         "id": "3d525611-3988-4ae8-b60d-c25c806aeec9",
         "name": "60 Degrees Mastercrafted"
      }
   ]
}
```
##### Response - 400
```json
{
   "success": false,
   "message": "both more_than and less_than param can't be empty",
   "status_code": 400
}
```