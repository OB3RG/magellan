-- +goose Up
CREATE TABLE site (
  id int NOT NULL PRIMARY KEY,
  url text
);

INSERT INTO site VALUES
(0, 'oskaroberg.com'),
(1, 'test.oskaroberg.com');

-- +goose Down
DROP TABLE site;
