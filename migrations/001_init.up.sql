CREATE TABLE public.person (
	"id" serial NOT NULL,
	"first_name" VARCHAR(255) NOT NULL,
	"last_name" VARCHAR(255) NOT NULL,
	"patronymic" VARCHAR(255) NOT NULL,
	"login" VARCHAR(255) NOT NULL unique ,
	"role_id" integer NOT NULL,
	"passwd" VARCHAR(255) NOT NULL,
	"updated_at" TIMESTAMP NOT NULL default current_timestamp,
	"deleted" BOOLEAN NOT NULL default false,
	CONSTRAINT "person_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.contact (
	"id" serial NOT NULL,
	"person_id" integer NOT NULL,
	"contact_type_id" integer NOT NULL,
	"contact_value" VARCHAR(255) NOT NULL,
	CONSTRAINT "contact_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.mentor (
	"id" integer NOT NULL,
    CONSTRAINT "mentor_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.mentor_notes (
	"id" serial NOT NULL,
	"student_id" integer NOT NULL,
	"mentor_id" integer NOT NULL,
	"note" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	CONSTRAINT "mentor_notes_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.course (
	"id" serial NOT NULL,
	"title" VARCHAR(255) NOT NULL,
	"start_date" TIMESTAMP NOT NULL,
	"end_date" TIMESTAMP NOT NULL,
	"course_status_id" integer NOT NULL,
	CONSTRAINT "course_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.group (
	"id" serial NOT NULL,
	"title" VARCHAR(255) NOT NULL,
	"course_id" integer NOT NULL,
	"mentor_id" integer NOT NULL,
	CONSTRAINT "group_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.homework (
	"id" serial NOT NULL,
	"lecture_id" integer NOT NULL,
	"title" VARCHAR(255) NOT NULL,
	"task" VARCHAR(255) NOT NULL,
	"max_score" integer NOT NULL,
	CONSTRAINT "homework_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.project (
	"id" serial NOT NULL,
	"title" VARCHAR(255) NOT NULL,
	"description" TEXT NOT NULL,
	"group_id" integer NOT NULL,
	CONSTRAINT "project_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.courses_status (
	"id" serial NOT NULL,
	"title" VARCHAR(255) NOT NULL,
	CONSTRAINT "courses_status_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.student_group (
	"student_id" integer NOT NULL,
	"group_id" integer NOT NULL,
	CONSTRAINT "student_group_pk" PRIMARY KEY ("student_id","group_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.operation_log (
	"id" serial NOT NULL,
	"description" text NOT NULL,
	"operation_id" integer NOT NULL,
	"person_id" integer NOT NULL,
	"role_id" integer NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	CONSTRAINT "operation_log_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.role (
	"id" serial NOT NULL,
	"role_name" VARCHAR(255) NOT NULL,
	"is_read_only" BOOLEAN NOT NULL,
	CONSTRAINT "role_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.contact_type (
	"id" integer NOT NULL,
	"type" VARCHAR(255) NOT NULL,
	CONSTRAINT "contact_type_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.lecture (
	"id" serial NOT NULL,
	"title" VARCHAR(255) NOT NULL,
	CONSTRAINT "lecture_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.course_lecture (
	"course_id" integer NOT NULL,
	"lecture_id" integer NOT NULL,
	CONSTRAINT "course_lecture_pk" PRIMARY KEY ("course_id","lecture_id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.operation (
	"id" serial NOT NULL,
	"operation_type_id" integer NOT NULL,
	"title" VARCHAR(255) NOT NULL,
	"is_logging" BOOLEAN NOT NULL,
	CONSTRAINT "operation_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.operation_type (
	"id" serial NOT NULL,
	"title" VARCHAR(255) NOT NULL,
	CONSTRAINT "operation_type_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.student_homework (
	"id" serial NOT NULL,
	"student_id" integer NOT NULL,
	"homework_id" integer NOT NULL,
	"started_at" TIMESTAMP NOT NULL,
	"finished_at" TIMESTAMP NOT NULL,
	"score" integer NOT NULL,
	CONSTRAINT "student_homework_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.student (
	"id" integer NOT NULL,
    CONSTRAINT "student_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.interview (
	"id" serial NOT NULL,
	"student_id" integer NOT NULL,
	"mentor_id" integer NOT NULL,
	"note" VARCHAR(255) NOT NULL,
	"score" integer NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	CONSTRAINT "interview_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.certificate (
	"id" serial NOT NULL,
	"template" VARCHAR(255) NOT NULL,
	CONSTRAINT "certificate_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE public.student_certificate (
	"id" serial NOT NULL,
	"student_id" integer NOT NULL,
	"certificate_id" integer NOT NULL,
	"course_id" integer NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	CONSTRAINT "student_certificate_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



ALTER TABLE person ADD CONSTRAINT "person_fk0" FOREIGN KEY ("role_id") REFERENCES "role"("id");

ALTER TABLE contact ADD CONSTRAINT "contact_fk0" FOREIGN KEY ("person_id") REFERENCES "person"("id");
ALTER TABLE contact ADD CONSTRAINT "contact_fk1" FOREIGN KEY ("contact_type_id") REFERENCES "contact_type"("id");

ALTER TABLE mentor ADD CONSTRAINT "mentor_fk0" FOREIGN KEY ("id") REFERENCES "person"("id");

ALTER TABLE mentor_notes ADD CONSTRAINT "mentor_notes_fk0" FOREIGN KEY ("student_id") REFERENCES "student"("id");
ALTER TABLE mentor_notes ADD CONSTRAINT "mentor_notes_fk1" FOREIGN KEY ("mentor_id") REFERENCES "mentor"("id");

ALTER TABLE course ADD CONSTRAINT "course_fk0" FOREIGN KEY ("course_status_id") REFERENCES "courses_status"("id");

ALTER TABLE "group" ADD CONSTRAINT "group_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");
ALTER TABLE "group" ADD CONSTRAINT "group_fk1" FOREIGN KEY ("mentor_id") REFERENCES "mentor"("id");

ALTER TABLE homework ADD CONSTRAINT "homework_fk0" FOREIGN KEY ("lecture_id") REFERENCES "lecture"("id");

ALTER TABLE project ADD CONSTRAINT "project_fk0" FOREIGN KEY ("group_id") REFERENCES "group"("id");


ALTER TABLE student_group ADD CONSTRAINT "student_group_fk0" FOREIGN KEY ("student_id") REFERENCES "student"("id");
ALTER TABLE student_group ADD CONSTRAINT "student_group_fk1" FOREIGN KEY ("group_id") REFERENCES "group"("id");

ALTER TABLE operation_log ADD CONSTRAINT "operation_log_fk0" FOREIGN KEY ("id") REFERENCES "person"("id");
ALTER TABLE operation_log ADD CONSTRAINT "operation_log_fk1" FOREIGN KEY ("operation_id") REFERENCES "operation"("id");
ALTER TABLE operation_log ADD CONSTRAINT "operation_log_fk2" FOREIGN KEY ("person_id") REFERENCES "person"("id");
ALTER TABLE operation_log ADD CONSTRAINT "operation_log_fk3" FOREIGN KEY ("role_id") REFERENCES "role"("id");




ALTER TABLE course_lecture ADD CONSTRAINT "course_lecture_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");
ALTER TABLE course_lecture ADD CONSTRAINT "course_lecture_fk1" FOREIGN KEY ("lecture_id") REFERENCES "lecture"("id");

ALTER TABLE operation ADD CONSTRAINT "operation_fk0" FOREIGN KEY ("operation_type_id") REFERENCES "operation_type"("id");


ALTER TABLE student_homework ADD CONSTRAINT "student_homework_fk0" FOREIGN KEY ("student_id") REFERENCES "student"("id");
ALTER TABLE student_homework ADD CONSTRAINT "student_homework_fk1" FOREIGN KEY ("homework_id") REFERENCES "homework"("id");

ALTER TABLE student ADD CONSTRAINT "student_fk0" FOREIGN KEY ("id") REFERENCES "person"("id");

ALTER TABLE interview ADD CONSTRAINT "interview_fk0" FOREIGN KEY ("student_id") REFERENCES "student"("id");
ALTER TABLE interview ADD CONSTRAINT "interview_fk1" FOREIGN KEY ("mentor_id") REFERENCES "mentor"("id");


ALTER TABLE student_certificate ADD CONSTRAINT "student_certificate_fk0" FOREIGN KEY ("student_id") REFERENCES "student"("id");
ALTER TABLE student_certificate ADD CONSTRAINT "student_certificate_fk1" FOREIGN KEY ("certificate_id") REFERENCES "certificate"("id");
ALTER TABLE student_certificate ADD CONSTRAINT "student_certificate_fk2" FOREIGN KEY ("course_id") REFERENCES "course"("id");


INSERT INTO public.role (role_name, is_read_only) VALUES
                                                      ('admin', false),
                                                      ('mentor', false),
                                                      ('student', false),
                                                      ('admin', true),
                                                      ('mentor', true),
                                                      ('student', true);

INSERT INTO public.person (first_name, last_name, patronymic, login, role_id, passwd) VALUES
    ('Ivan', 'Ivanov', 'Ivanovich', 'Ivan777', 3, '12345678');