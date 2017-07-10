DROP TABLE comments;
DROP TABLE posts;

CREATE TABLE posts (
  id      SERIAL PRIMARY KEY,
  content TEXT,
  author  VARCHAR(255)
);

CREATE TABLE comments (
  id      SERIAL PRIMARY KEY,
  content TEXT,
  author  VARCHAR(255),
  post_id INTEGER REFERENCES posts (id)
);

