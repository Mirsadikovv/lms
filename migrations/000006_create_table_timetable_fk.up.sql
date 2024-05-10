ALTER TABLE time_table ADD FOREIGN KEY (teacher_id) REFERENCES teachers (id);

ALTER TABLE time_table ADD FOREIGN KEY (student_id) REFERENCES students (id);

ALTER TABLE time_table ADD FOREIGN KEY (subject_id) REFERENCES subjects (id);

ALTER TABLE teachers ADD FOREIGN KEY (subject_id) REFERENCES subjects (id);
