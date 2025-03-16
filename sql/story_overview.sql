-- story_overview definition

CREATE TABLE story_overview (
  namespace varchar(32),
  story_id INTEGER,
  story_name varchar(50),
  db_path varchar(50));

CREATE INDEX idx_space_story ON story_overview (namespace,story_id)
;