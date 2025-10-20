-- +goose Up
-- +goose StatementBegin
CREATE TABLE spheres (
    id SERIAL PRIMARY KEY,
    name TEXT,
    user_id INTEGER REFERENCES users(id),
    parent_id INTEGER REFERENCES spheres(id),
    children_id INTEGER REFERENCES spheres(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE spheres;
-- +goose StatementEnd
