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


CREATE OR REPLACE FUNCTION trigger_delivery_id()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO orders(delivery_id) VALUES (NEW.id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_delivery
    AFTER INSERT ON delivery
    FOR EACH ROW
EXECUTE FUNCTION trigger_delivery_id();



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


CREATE OR REPLACE FUNCTION trigger_payment_id()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO orders(payment_id) VALUES (NEW.id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_payment
    AFTER INSERT ON payment
    FOR EACH ROW
EXECUTE FUNCTION trigger_payment_id();

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


CREATE OR REPLACE FUNCTION trigger_items_id()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO orders(item_id) VALUES (NEW.id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_items
    AFTER INSERT ON items
    FOR EACH ROW
EXECUTE FUNCTION trigger_delivery_id();


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
);



