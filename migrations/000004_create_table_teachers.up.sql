CREATE TABLE teachers (
  id uuid PRIMARY KEY,
  first_name VARCHAR,
  last_name VARCHAR,
  subject_id uuid,
  start_working timestamp,
  phone VARCHAR,
  mail VARCHAR,
  created_at timestamp NOT NULL DEFAULT now(),
  updated timestamp
  );