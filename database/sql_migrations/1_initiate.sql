-- +migrate Up
-- +migrate StatmentBegin

CREATE TABLE category(
   id SERIAL PRIMARY KEY,
   name VARCHAR(255),
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    image_url VARCHAR(1024),
    release_year INT,
    price VARCHAR(255),
    total_page INT,
    thickness VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    category_id INT,
    FOREIGN KEY (category_id) REFERENCES category(id)
);

-- +migrate StatementEnd