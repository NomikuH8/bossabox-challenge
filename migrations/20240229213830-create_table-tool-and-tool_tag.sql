CREATE TABLE tool (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  link TEXT NOT NULL,
  description TEXT NOT NULL
)

CREATE TABLE tool_tag (
  id SERIAL PRIMARY KEY,
  tool_id INTEGER NOT NULL,
  tag_name TEXT
)