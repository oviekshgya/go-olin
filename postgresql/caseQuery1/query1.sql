SELECT u.name AS user_name, SUM(oi.price * oi.quantity) AS total_spent
FROM users u
         JOIN orders o ON u.id = o.user_id
         JOIN order_items oi ON o.id = oi.order_id
WHERE o.created_at >= '2022-01-01'
GROUP BY u.name
HAVING SUM(oi.price * oi.quantity) >= 1000;
