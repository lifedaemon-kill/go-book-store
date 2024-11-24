CREATE TABLE books
(
    id        SERIAL PRIMARY KEY,
    title     VARCHAR(100)   NOT NULL,
    author    VARCHAR(100)   NOT NULL,
    price     DECIMAL(10, 2) NOT NULL,
    purchased INTEGER        NOT NULL
);
