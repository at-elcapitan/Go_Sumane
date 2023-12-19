CREATE TABLE IF NOT EXISTS subjects (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  subject_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  admin_override BOOL NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS task (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_created_id INTEGER NOT NULL,
  text TEXT NOT NULL,
  daeadline DATE NOT NULL,
  subject_id INTEGER NOT NULL,
  FOREIGN KEY (subject_id) REFERENCES subjects (id) ON DELETE CASCADE,
  FOREIGN KEY (user_created_id) REFERENCES users (id) ON DELETE NO ACTION
);

