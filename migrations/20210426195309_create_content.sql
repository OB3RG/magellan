-- +goose Up
CREATE TABLE content (
  id int NOT NULL PRIMARY KEY,
  markdown text
);

INSERT INTO content VALUES
(0, '#Hello World');

-- +goose Down
DROP TABLE content;
