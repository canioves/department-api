-- +goose Up
CREATE TABLE "departments" (
    "id" bigserial,
    "name" text,
    "parent_id" bigint,
    "created_at" timestamptz,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_departments_childern" FOREIGN KEY ("parent_id") REFERENCES "departments"("id")
);

-- +goose Down
DROP TABLE "departments" CASCADE;
