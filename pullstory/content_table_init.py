import sqlite3
import traceback


def create(story_id):
  # create story_content table
  try:
    open('sql/story_content.sql', 'r') as sql_script:
      sql = sql_script.read()
      with sqlite3.connect('db/content/'+str(story_id)) as conn:
        with conn.cursor() as cur:
          cur.executescript(sql)
          conn.commit()
  except Exception as e:
    traceback.print_exc()
