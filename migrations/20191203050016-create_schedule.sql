-- +migrate Up
CREATE TABLE IF NOT EXISTS schedule (
  type INT,
  place_id VARCHAR(6) NOT NULL,
  date TIMESTAMP,
  created_at TIMESTAMP default current_timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS schedule;