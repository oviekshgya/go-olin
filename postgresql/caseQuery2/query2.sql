SELECT u.name AS user_name, o.amount, o.created_at
FROM users u
         JOIN orders o ON u.id = o.user_id;
