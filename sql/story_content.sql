-- story_content definition

CREATE TABLE story_content (
  page_no INTEGER,
  page_desc varchar(50),
  content TEXT);

CREATE INDEX idx_page ON story_content (page_no)
;