-- create database
CREATE DATABASE url_db;
\c url_db

-- creates user
CREATE USER backend WITH PASSWORD 'qwer1234';
GRANT CONNECT ON DATABASE url_db TO backend;

-- create schema for urls tables
CREATE SCHEMA urls;

-- create urls tables
CREATE TABLE IF NOT EXISTS urls.urls (
  "url_hash" VARCHAR(10) NOT NULL,
  "url" TEXT NOT NULL,
  PRIMARY KEY ("url_hash")
);

CREATE TABLE IF NOT EXISTS urls.urls_linking (
	"id" UUID NOT NULL,
	"url_hash" VARCHAR(10) NOT NULL,
	"visited_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	CONSTRAINT "FK_urls_linking_urls" FOREIGN KEY ("url_hash") REFERENCES urls.urls ("url_hash") ON UPDATE NO ACTION ON DELETE CASCADE
);

-- grant privilegies to the user
GRANT usage ON SCHEMA urls TO backend;
GRANT select, insert, update, delete, trigger ON all tables IN SCHEMA urls TO backend;
GRANT usage, select ON all sequences IN SCHEMA urls TO backend;
GRANT execute ON all functions IN SCHEMA urls TO backend;
