ALTER TABLE students
ADD CONSTRAINT phone UNIQUE (phone);

ALTER TABLE teachers
ADD CONSTRAINT phone UNIQUE (phone);

ALTER TABLE students
ADD CONSTRAINT mail UNIQUE (mail);


ALTER TABLE teachers
ADD CONSTRAINT mail UNIQUE (mail);
