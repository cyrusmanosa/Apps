CREATE TABLE "fixinformation" (
  "user_id" INT GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY,
  "first_name" VARCHAR NOT NULL,
  "last_name" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "hashed_password" VARCHAR NOT NULL,
  "birth" VARCHAR NOT NULL,
  "country" VARCHAR NOT NULL,
  "gender" VARCHAR NOT NULL,
  "blood" VARCHAR NOT NULL,
  "age" INT NOT NULL,
  "constellation" VARCHAR NOT NULL,
  "certification" BOOLEAN DEFAULT false,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "password_changed_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "canchangeinformation" (
  "user_id" INT PRIMARY KEY,
  "nickname" VARCHAR NOT NULL,
  "city" VARCHAR NOT NULL,
  "sexual" VARCHAR NOT NULL,
  "height" FLOAT NOT NULL,
  "weight" FLOAT NOT NULL,
  "language" VARCHAR NOT NULL,
  "education" VARCHAR NOT NULL,
  "job" VARCHAR NOT NULL,
  "annual_salary" INT NOT NULL,
  "sociability" VARCHAR NOT NULL,
  "religious" VARCHAR NOT NULL,
  "introduce" VARCHAR NOT NULL,
  "info_changed_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "loverrequest" (
  "user_id" INT PRIMARY KEY,
  "min_age" INT NOT NULL,
  "max_age" INT NOT NULL,
  "city" VARCHAR NOT NULL,
  "gender" VARCHAR NOT NULL,
  "constellation" VARCHAR NOT NULL,
  "sexual" VARCHAR NOT NULL,
  "height" FLOAT NOT NULL,
  "weight" FLOAT NOT NULL,
  "language" VARCHAR NOT NULL,
  "job" VARCHAR NOT NULL,
  "annual_salary" INT NOT NULL,
  "sociability" VARCHAR NOT NULL,
  "religious" VARCHAR NOT NULL,
  "certification" BOOLEAN DEFAULT false,
  "info_changed_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "hobbyrequest" (
  "user_id" INT PRIMARY KEY,
  "era" INT NOT NULL,
  "city" VARCHAR NOT NULL,
  "gender" VARCHAR NOT NULL,
  "height" INT NOT NULL,
  "weight" INT NOT NULL,
  "language" VARCHAR NOT NULL,
  "find_type" VARCHAR NOT NULL,
  "find_target" VARCHAR NOT NULL,
  "experience" INT NOT NULL,
  "sociability" VARCHAR NOT NULL,
  "certification" BOOLEAN DEFAULT false,
  "info_changed_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "accompanyrequest" (
  "user_id" INT PRIMARY KEY,
  "era" INT NOT NULL,
  "city" VARCHAR NOT NULL,
  "language" VARCHAR NOT NULL,
  "gender" VARCHAR NOT NULL,
  "find_type" VARCHAR NOT NULL,
  "find_target" VARCHAR NOT NULL,
  "sociability" VARCHAR NOT NULL,
  "certification" BOOLEAN DEFAULT false,
  "info_changed_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "image" (
  "user_id" INT PRIMARY KEY,
  "img1" VARCHAR NOT NULL,
  "img2" VARCHAR NOT NULL,
  "img3" VARCHAR NOT NULL,
  "img4" VARCHAR NOT NULL,
  "img5" VARCHAR NOT NULL
);

CREATE TABLE "targetlist" (
  "user_id" INT PRIMARY KEY,
  "target_1_id" INT,
  "target_2_id" INT,
  "target_3_id" INT,
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "changetargetuser" (
  "user_id" INT,
  "change_user_id" INT NOT NULL,
  "reason" VARCHAR NOT NULL,
  "frequency" INT GENERATED BY DEFAULT AS IDENTITY NOT NULL DEFAULT 1,
  "change_time" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "complaint" (
  "cp_id" INT GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY,
  "user_id" INT,
  "cp_target_id" INT NOT NULL,
  "cp_tpye" VARCHAR NOT NULL,
  "cp_message" VARCHAR NOT NULL,
  "status" VARCHAR NOT NULL,
  "complaint_time" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "sideInformation" (
  "user_id" INT PRIMARY KEY,
  "type" VARCHAR,
  "role" VARCHAR,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "complaintProcess" (
  "cp_id" INT PRIMARY KEY,
  "staff_id" INT NOT NULL,
  "user_reply" VARCHAR,
  "resulf_message" VARCHAR NOT NULL,
  "status" VARCHAR NOT NULL,
  "result_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "staff" (
  "staff_id" INT GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY,
  "fullname" VARCHAR NOT NULL,
  "birth" VARCHAR NOT NULL,
  "phone_no" INT NOT NULL,
  "address" VARCHAR NOT NULL,
  "hiredate" VARCHAR NOT NULL,
  "depart_no" INT NOT NULL,
  "salary" INT NOT NULL
);

CREATE TABLE "payment" (
  "pay_id" UUID UNIQUE PRIMARY KEY,
  "username" VARCHAR NOT NULL,
  "plan" VARCHAR NOT NULL,
  "pay_type" VARCHAR NOT NULL,
  "memoey" INT NOT NULL,
  "pay_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

ALTER TABLE "canchangeinformation" ADD FOREIGN KEY ("user_id") REFERENCES "fixinformation" ("user_id");

ALTER TABLE "loverrequest" ADD FOREIGN KEY ("user_id") REFERENCES "fixinformation" ("user_id");

ALTER TABLE "hobbyrequest" ADD FOREIGN KEY ("user_id") REFERENCES "fixinformation" ("user_id");

ALTER TABLE "accompanyrequest" ADD FOREIGN KEY ("user_id") REFERENCES "fixinformation" ("user_id");

ALTER TABLE "image" ADD FOREIGN KEY ("user_id") REFERENCES "fixinformation" ("user_id");

ALTER TABLE "targetlist" ADD FOREIGN KEY ("user_id") REFERENCES "fixinformation" ("user_id");

ALTER TABLE "changetargetuser" ADD FOREIGN KEY ("user_id") REFERENCES "fixinformation" ("user_id");

ALTER TABLE "complaint" ADD FOREIGN KEY ("user_id") REFERENCES "fixinformation" ("user_id");

ALTER TABLE "sideInformation" ADD FOREIGN KEY ("user_id") REFERENCES "fixinformation" ("user_id");

ALTER TABLE "complaintProcess" ADD FOREIGN KEY ("cp_id") REFERENCES "complaint" ("cp_id");
