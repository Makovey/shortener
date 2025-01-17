-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shortener(
     id SERIAL PRIMARY KEY,
     short_url VARCHAR(100) NOT NULL,
     original_url VARCHAR(255) NOT NULL,
     created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC') NOT NULL,
     owner_user_id VARCHAR(100) NOT NULL,
     is_deleted BOOLEAN DEFAULT FALSE,
     UNIQUE (original_url)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shortener;
-- +goose StatementEnd
