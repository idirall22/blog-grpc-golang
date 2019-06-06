CREATE TABLE IF NOT EXISTS users
(
  id serial PRIMARY KEY,
  username VARCHAR(128) NOT NULL,
  email VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS posts
(
  id serial PRIMARY KEY,
  author INTEGER REFERENCES users(id),
  title VARCHAR(128) NOT NULL,
  content TEXT NOT NULL,
  timestemp date NOT NULL DEFAULT CURRENT_TIMESTAMP,
  published BOOLEAN NOT NULL,
  comments_list JSONB NOT NULL DEFAULT '{}'::jsonb
)
