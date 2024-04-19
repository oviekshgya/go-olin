CREATE TABLE users (id SERIAL PRIMARY KEY, name VARCHAR(255), email
                       VARCHAR(255));

INSERT INTO users (name, email) VALUES ('John Doe', 'john.doe@example.com'),
                                       ('Jane Smith', 'jane.smith@example.com');