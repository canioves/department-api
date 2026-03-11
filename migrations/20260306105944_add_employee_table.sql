-- +goose Up
CREATE TABLE "employees" (
    "id" bigserial,
    "full_name" text NOT NULL,
    "position" text NOT NULL,
    "hired_at" timestamptz,
    "created_at" timestamptz DEFAULT NOW(),
    "department_id" bigint NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_departments_employees" FOREIGN KEY ("department_id") REFERENCES "departments"("id")
);

-- +goose Down
DROP TABLE "employees" CASCADE;
