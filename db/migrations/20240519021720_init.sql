-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS book (
    id INTEGER PRIMARY KEY,
    title TEXT UNIQUE,
    subtitle TEXT,
    description TEXT,
    purchase_link TEXT,
    published_date TEXT,
    language TEXT,
    page_count INTEGER
);

CREATE TABLE IF NOT EXISTS author (
    id INTEGER PRIMARY KEY,
    name TEXT
);

CREATE TABLE IF NOT EXISTS book_author (
    book_id INTEGER,
    author_id INTEGER,
    FOREIGN KEY (book_id) REFERENCES book(id),
    FOREIGN KEY (author_id) REFERENCES author(id)
);

CREATE TABLE IF NOT EXISTS category (
    id INTEGER,
    name TEXT
);

CREATE TABLE IF NOT EXISTS book_category (
    book_id INTEGER,
    category_id INTEGER,
    FOREIGN KEY (book_id) REFERENCES book(id),
    FOREIGN KEY (category_id) references category(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE book_category;
DROP TABLE category;
DROP TABLE book_author;
DROP TABLE author;
DROP TABLE book;
-- +goose StatementEnd
