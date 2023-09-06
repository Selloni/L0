CREATE TABLE orders (
                        order_uid TEXT,
                        track_number TEXT,
                        entry TEXT,
                        delivery JSONB,
                        payment JSONB,
                        items JSONB,
                        locale TEXT,
                        internal_signature TEXT,
                        customer_id TEXT,
                        delivery_service TEXT,
                        shardkey TEXT,
                        sm_id INTEGER,
                        date_created TIMESTAMP,
                        oof_shard TEXT
);

CREATE TABLE delivery (
                          name TEXT,
                          phone TEXT,
                          zip TEXT,
                          city TEXT,
                          address TEXT,
                          region TEXT,
                          email TEXT
);


CREATE TABLE payment (
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