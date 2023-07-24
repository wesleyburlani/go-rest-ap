CREATE TABLE IF NOT EXISTS users (
  id        BIGSERIAL NOT NULL PRIMARY KEY,
  username  TEXT      NOT NULL UNIQUE,
  email     TEXT      NOT NULL UNIQUE,
  password  TEXT      NOT NULL,
  role      TEXT      NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

