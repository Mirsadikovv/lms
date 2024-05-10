CREATE TABLE IF NOT EXISTS subjects (
  id uuid PRIMARY KEY,
  name VARCHAR,
  type VARCHAR,
  created_at timestamp NOT NULL DEFAULT 'now()',
  updated timestamp
);

CREATE TABLE IF NOT EXISTS time_table (
  id uuid PRIMARY KEY,
  teacher_id uuid NOT NULL,
  student_id uuid NOT NULL,
  subject_id uuid NOT NULL,
  from_date timestamp NOT NULL,
  to_date timestamp NOT NULL
);


