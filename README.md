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

---
Method : `GET`
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

#### List Restaurants

---
Method : `GET`
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

#### Search Restaurants

---
Method : `GET`
> /api/v1/restaurant/search?q=piz

Query Params:
> Optional
> * q (ex - piz)

##### Response - 200
```json
{
   "success": true,
   "message": "request successful!",
   "status_code": 200,
   "data": [
      {
         "id": "b6e35136-ab76-49cf-8ae2-a39b962c03af",
         "name": "Pizza Burg"
      },
      {
         "id": "40ef288d-bbf3-4947-a8e8-3b8d033e094d",
         "name": "Pier W"
      }
   ]
}
```
##### Response - 400
```json
{
   "success": false,
   "message": "search term 'q' missing",
   "status_code": 400
}
```

#### Search Dishes

---

Method : `GET`
> /api/v1/dish/search?q=piz

Query Params:
> Optional
> * q (ex - piz)

##### Response - 200
```json
{
   "success": true,
   "message": "request successful!",
   "status_code": 200,
   "data": [
      {
         "id": "96f84b6b-a7a3-4fd1-a819-2e9468675773",
         "name": "Pie",
         "price": 10.1
      },
      {
         "id": "708a3685-2be4-4bcd-8149-cf8a4988d81e",
         "name": "Pie",
         "price": 10.3
      }
   ]
}
```
##### Response - 400
```json
{
   "success": false,
   "message": "search term 'q' missing",
   "status_code": 400
}
```

#### Purchase Dish

---

Method : `POST`
> /api/v1/user/purchase/:user_id

Params:
> Required
> * user_id (ex - 552)

Request Body:
```json
{
   "restaurant_id": "00017a27-5fcc-4e01-acab-b791aa0a6292",
   "menu_id": "0d9625a0-1298-40f2-b168-167b4ad70d74"
}
```

##### Response - 200
```json
{
   "success": true,
   "message": "purchased successfully!",
   "status_code": 202,
   "data": {
      "current_balance": 3.5099998
   }
}
```
##### Response - 400
```json
{
   "success": false,
   "message": "not a valid request body",
   "status_code": 400
}
```

##### Response - 404
```json
{
    "success": false,
    "message": "restaurant or dish does not exist",
    "status_code": 404
}
```

##### Response - 406
```json
{
    "success": false,
    "message": "you don't have enough cash to buy this dish! you have $3.51",
    "status_code": 406
}
```