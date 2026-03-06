-- +goose Up
CREATE TABLE "employees" (
    "id" bigserial,
    "full_name" text,
    "position" text,
    "hired_at" timestamptz,
    "created_at" timestamptz,
    "department_id" bigint,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_departments_employees" FOREIGN KEY ("department_id") REFERENCES "departments"("id")
);

-- +goose Down
DROP TABLE "employees" CASCADE;
