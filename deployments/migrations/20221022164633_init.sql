-- +goose Up
-- +goose StatementBegin
CREATE TABLE shops (
    id SERIAL PRIMARY KEY,
    name TEXT,
    go_search_template TEXT
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    articul VARCHAR(200),
    -- short_name VARCHAR(200),
--    barcode VARCHAR(255),
    -- go_search_template TEXT,
    url TEXT,
    shop_id INTEGER,
    price INTEGER,
    -- TODO: updated_at 
     CONSTRAINT fk_products_shops
      FOREIGN KEY (shop_id) 
	  REFERENCES shops(id)
);

CREATE TABLE barcode_products (
    barcode VARCHAR(255),
    product_id INTEGER,
    CONSTRAINT fk_barcode_products_products
      FOREIGN KEY (product_id) 
	  REFERENCES products(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE shops;
DROP TABLE products;
DROP TABLE barcode_products;
-- +goose StatementEnd
