CREATE TABLE order_items (id INT PRIMARY KEY, order_id INT, product_name
                             VARCHAR(255), price DECIMAL(10,2), quantity INT);

INSERT INTO order_items (id, order_id, product_name, price, quantity) VALUES (1, 1,
                                                                              'T-Shirt', 25.00, 2), (2, 1, 'Jeans', 50.00, 1), (3, 2, 'Socks', 10.00, 5), (4, 3, 'Shoes',
                                                                                                                                                           75.00, 2), (5, 4, 'Jacket', 100.00, 1), (6, 5, 'Sweater', 25.00, 3);