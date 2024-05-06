CREATE TABLE subjects (
  id uuid PRIMARY KEY,
  name string,
  type string,
  created_at timestamp NOT NULL DEFAULT 'now()',
  updated timestamp
);

CREATE TABLE time_table (
  id uuid PRIMARY KEY,
  teacher_id uuid NOT NULL,
  student_id uuid NOT NULL,
  subject_id uuid NOT NULL,
  from_date timestamp NOT NULL,
  to_date timestamp NOT NULL
);

ALTER TABLE teachers ADD FOREIGN KEY (subject_id) REFERENCES subjects (id);

ALTER TABLE time_table ADD FOREIGN KEY (teacher_id) REFERENCES teachers (id);

ALTER TABLE time_table ADD FOREIGN KEY (student_id) REFERENCES students (id);

ALTER TABLE time_table ADD FOREIGN KEY (subject_id) REFERENCES subjects (id);