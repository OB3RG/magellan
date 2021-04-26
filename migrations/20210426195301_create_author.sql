-- +goose Up
CREATE TABLE author (
  id int NOT NULL,
  given_name text,
  surname text
);

INSERT INTO author VALUES
(0, 'oskar', 'oberg');

-- +goose Down
DROP TABLE author;
