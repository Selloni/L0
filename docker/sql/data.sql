DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS delivery;
DROP TABLE IF EXISTS payment;
DROP TABLE IF EXISTS items;

CREATE TABLE delivery (
    order_uid TEXT,
    name TEXT,
    phone TEXT,
    zip TEXT,
    city TEXT,
    address TEXT,
    region TEXT,
    email TEXT
);

CREATE TABLE payment (
    order_uid TEXT,
    transaction TEXT,
    request_id TEXT,
    currency TEXT,
    provider TEXT,
    amount BIGINT,
    payment_dt BIGINT,
    bank TEXT,
    delivery_cost BIGINT,
    goods_total BIGINT,
    custom_fee BIGINT
);

CREATE TABLE items (
    order_uid TEXT,
    chrt_id BIGINT,
    track_number TEXT,
    price BIGINT,
    rid TEXT,
    name TEXT,
    sale BIGINT,
    size TEXT,
    total_price BIGINT,
    nm_id BIGINT,
    brand TEXT,
    status BIGINT
);

CREATE TABLE orders (
    order_uid TEXT,
    track_number TEXT,
    entry TEXT,
    locale TEXT,
    internal_signature TEXT,
    customer_id TEXT,
    delivery_service TEXT,
    shardkey TEXT,
    sm_id BIGINT,
    date_created TIMESTAMP,
    oof_shard TEXT
);



