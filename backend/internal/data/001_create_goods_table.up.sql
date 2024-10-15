CREATE TABLE IF NOT EXISTS
    goods
(
    id SERIAL PRIMARY KEY,
    name TEXT,
    price NUMERIC(10, 2),
    city_id INT
)