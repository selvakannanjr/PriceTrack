CREATE TABLE links (
  id   BIGSERIAL PRIMARY KEY,
  url text      NOT NULL,
  UNIQUE(url)
);