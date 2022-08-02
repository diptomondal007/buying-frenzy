CREATE TABLE "user_info" (
    id serial primary key unique,
    name varchar(100),
    cash_balance float4
);

CREATE TABLE "restaurant" (
    id varchar(100) primary key unique,
    name varchar(500),
    cash_balance float4
);

CREATE TABLE "open_hour" (
    id serial primary key unique,
    week_name varchar(20),
    start_time time,
    closing_time time,
    restaurant_id varchar(100),

    CONSTRAINT fk_restaurant_id FOREIGN KEY (restaurant_id) REFERENCES restaurant(id)
);

CREATE TABLE "dish" (
    id varchar(100) primary key unique,
    name varchar(500),
    price float4,
    restaurant_id varchar(100),

    CONSTRAINT fk_restaurant_id FOREIGN KEY (restaurant_id) REFERENCES restaurant(id)
);

CREATE TABLE "purchase_history" (
    id serial primary key unique,
    transaction_amount float4,
    transaction_date timestamptz,
    restaurant_id varchar(100),
    dish_id varchar(100),
    user_id int,

    CONSTRAINT fk_restaurant_id FOREIGN KEY (restaurant_id) REFERENCES restaurant(id),
    CONSTRAINT fk_dish_id FOREIGN KEY (dish_id) REFERENCES dish(id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES user_info(id)
);