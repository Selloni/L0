DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS delivery;
DROP TABLE IF EXISTS payment;
DROP TABLE IF EXISTS items;



CREATE TABLE delivery (
    id SERIAL PRIMARY KEY,
    name TEXT,
    phone TEXT,
    zip TEXT,
    city TEXT,
    address TEXT,
    region TEXT,
    email TEXT
);


CREATE TABLE payment (
    id SERIAL PRIMARY KEY,
    transaction TEXT,
    request_id TEXT,
    currency TEXT,
    provider TEXT,
    amount INTEGER,
    payment_dt INTEGER,
    bank TEXT,
    delivery_cost INTEGER,
    goods_total INTEGER,
    custom_fee INTEGER
);


CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    chrt_id INTEGER,
    track_number TEXT,
    price INTEGER,
    rid TEXT,
    name TEXT,
    sale INTEGER,
    size TEXT,
    total_price INTEGER,
    nm_id INTEGER,
    brand TEXT,
    status INTEGER
);
-- REFERENCES delivery (id) NOT NULL
CREATE TABLE orders (
    order_uid TEXT,
    track_number TEXT,
    entry TEXT,
    delivery_id INTEGER REFERENCES delivery (id),
    payment_id INTEGER REFERENCES payment (id),
    item_id INTEGER REFERENCES items(id),
    locale TEXT,
    internal_signature TEXT,
    customer_id TEXT,
    delivery_service TEXT,
    shardkey TEXT,
    sm_id INTEGER,
    date_created TIMESTAMP,
    oof_shard TEXT
--     FOREIGN KEY (delivery_id) REFERENCES delivery (id),
--     FOREIGN KEY (payment_id) REFERENCES payment (id),
--     FOREIGN KEY (item_id) REFERENCES items (id)
);

