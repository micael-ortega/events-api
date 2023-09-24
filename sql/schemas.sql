CREATE TABLE IF NOT EXISTS course(
  id INTEGER PRIMARY KEY,
  course TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS instructor(
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS attendee(
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  cpf TEXT NOT NULl,
  role TEXT NOT NULL,
  board TEXT NOT NULL,
  company  TEXT NOT NULL,
  branch TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS event(
 id INTEGER PRIMARY KEY,
 begin_date DATE NOT NULL,
 end_date DATE NOT NULL,
 modality TEXT NOT NULL,
 duration REAL NOT NULL,
 instructor_id  INTEGER NOT NULL,
 course_id INTEGER NOT NULL,
 FOREIGN KEY (instructor_id) REFERENCES instructor(id),
 FOREIGN KEY (course_id) REFERENCES course(id)
);

CREATE TABLE IF NOT EXISTS event_attendee(
  id INTEGER PRIMARY KEY,
  event_id INTEGER NOT NULL,
  attendee_id INTEGER NOT NULL,
  status INTEGER NOT NULL,
  FOREIGN KEY (event_id) REFERENCES event(id),
  FOREIGN KEY (attendee_id) REFERENCES attendee(id)
);

CREATE TABLE IF NOT EXISTS user(
  id INTEGER PRIMARY KEY,
  username TEXT UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
);