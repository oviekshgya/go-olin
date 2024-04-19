CREATE TABLE orders (id SERIAL PRIMARY KEY, user_id INT, amount NUMERIC(10,2),
                     created_at TIMESTAMP);

INSERT INTO orders (user_id, amount, created_at) VALUES (1, 100.00, now()), (2,
                                                                             200.00, now()), (1, 50.00, now());