-- +goose Up
CREATE TABLE "departments" (
    "id" bigserial,
    "name" text NOT NULL,
    "parent_id" bigint,
    "created_at" timestamptz DEFAULT NOW(),
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_departments_childern" FOREIGN KEY ("parent_id") REFERENCES "departments"("id") ON DELETE CASCADE
);

-- +goose Down
DROP TABLE "departments" CASCADE;
