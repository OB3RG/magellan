-- +goose Up
CREATE TABLE tag (
  id int NOT NULL PRIMARY KEY,
  name text
);

INSERT INTO tag VALUES
(0, 'default'),
(1, 'blog');

-- +goose Down
DROP TABLE tag;
