CREATE TABLE "user_info" (
    id serial primary key unique,
    name varchar(100),
    cash_balance float4
);

CREATE TABLE "restaurant" (
    id serial primary key unique,
    name varchar(100),
    cash_balance float4
);

CREATE TABLE "open_hour" (
    id serial primary key unique,
    week_name varchar(20),
    start_time time,
    closing_time time,
    restaurant_id int,

    CONSTRAINT fk_restaurant_id FOREIGN KEY (restaurant_id) REFERENCES restaurant(id)
);

CREATE TABLE "dish" (
    id serial primary key unique,
    name varchar(100),
    price float4,
    restaurant_id int,

    CONSTRAINT fk_restaurant_id FOREIGN KEY (restaurant_id) REFERENCES restaurant(id)
);

CREATE TABLE "purchase_history" (
    id serial primary key unique,
    restaurant_id int,
    dish_id int,

    CONSTRAINT fk_restaurant_id FOREIGN KEY (restaurant_id) REFERENCES restaurant(id),
    CONSTRAINT fk_dish_id FOREIGN KEY (dish_id) REFERENCES dish(id)
);