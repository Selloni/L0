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
    INSERT INTO orders VALUES (NEW.id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_delivery
    BEFORE INSERT ON delivery
    FOR EACH ROW
EXECUTE FUNCTION trigger_delivery_id();



CREATE TABLE payment (
    id SERIAL PRIMARY KEY,
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


CREATE OR REPLACE FUNCTION trigger_payment_id()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO orders(payment_id) VALUES (NEW.id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_payment
    BEFORE INSERT ON payment
    FOR EACH ROW
EXECUTE FUNCTION trigger_payment_id();

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
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


CREATE OR REPLACE FUNCTION trigger_items_id()
    RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO orders(item_id) VALUES (NEW.id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trigger_items
    BEFORE INSERT ON items
    FOR EACH ROW
EXECUTE FUNCTION trigger_delivery_id();


-- REFERENCES delivery (id) NOT NULL
CREATE TABLE orders (
    order_uid TEXT,
    track_number TEXT,
    entry TEXT,
    delivery_id BIGINT ,
    payment_id BIGINT ,
    item_id BIGINT ,
    locale TEXT,
    internal_signature TEXT,
    customer_id TEXT,
    delivery_service TEXT,
    shardkey TEXT,
    sm_id BIGINT,
    date_created TIMESTAMP,
    oof_shard TEXT
--     FOREIGN KEY (delivery_id) REFERENCES delivery (id),
--     FOREIGN KEY (payment_id) REFERENCES payment (id),
--     FOREIGN KEY (item_id) REFERENCES items (id)
);



