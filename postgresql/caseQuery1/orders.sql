CREATE TABLE orders (id INT PRIMARY KEY, user_id INT, amount DECIMAL(10,2),
                     created_at TIMESTAMP);

INSERT INTO orders (id, user_id, amount, created_at) VALUES (1, 1, 100.00, '2022
01-02 10:30:00'), (2, 2, 50.00, '2022-01-03 09:00:00'), (3, 1, 150.00, '2022-01-04
14:15:00'), (4, 3, 200.00, '2022-01-05 17:45:00'), (5, 2, 75.00, '2022-01-06 11:20:00');