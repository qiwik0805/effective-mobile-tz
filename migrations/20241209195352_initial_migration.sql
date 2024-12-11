-- +goose Up
-- +goose StatementBegin
CREATE TABLE song (
    id SERIAL PRIMARY KEY,
    song TEXT NOT NULL,
    group_name TEXT NOT NULL,
    release_date TEXT NOT NULL,
    text TEXT NOT NULL,
    link TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE song;
-- +goose StatementEnd
