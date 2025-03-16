import sqlite3
import traceback


def persist(story_id, story_name, page_no, content, page_desc):
  try:
    with sqlite3.connect('db/content/'+str(story_id)) as conn:
      with conn.cursor() as cur:
        sql = 'INSERT INTO story_content (page_no, page_desc, content) VALUES(?, ?, ?)'
        cur.execute(sql, (page_no, page_desc, content))
        conn.commit()
    print('insert success: %s_%s' % (story_name, page_desc))
  except Exception as e:
    print('insert failed: %s_%s' % (story_name, page_desc))
    traceback.print_exc()
