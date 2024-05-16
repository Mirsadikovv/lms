ALTER TABLE students
ADD CONSTRAINT phone_s UNIQUE (phone);

ALTER TABLE teachers
ADD CONSTRAINT phone_t UNIQUE (phone);

ALTER TABLE students
ADD CONSTRAINT mail_s UNIQUE (mail);

ALTER TABLE teachers
ADD CONSTRAINT mail_t UNIQUE (mail);
