-- +goose Up
-- +goose StatementBegin
CREATE TABLE shops (
    id INTEGER PRIMARY KEY,
    name TEXT,
    go_search_template TEXT
);
CREATE TABLE products (
    id INTEGER PRIMARY KEY,
    barcode VARCHAR(255),
    go_search_template TEXT
);
CREATE TABLE prices (
    shop_id INTEGER,
    barcode VARCHAR(255),
    price_min INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE prices;
DROP TABLE shops;
DROP TABLE products;
-- +goose StatementEnd
