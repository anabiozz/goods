CREATE DATABASE art OWNER postgres;

DROP SCHEMA IF EXISTS main;
CREATE SCHEMA main;

DROP TABLE IF EXISTS main.images;
CREATE TABLE main.images (
  uuid TEXT NOT NULL,
  name TEXT,
  materials TEXT,
  year TEXT,
  size TEXT,
  type INT, 
  is_for_sale BOOLEAN,
  PRIMARY KEY(uuid)
);

CREATE TABLE IF EXISTS main.images_type (
  id INT NOT NULL,
  type TEXT NOT NULL,
  PRIMARY KEY(id),
);