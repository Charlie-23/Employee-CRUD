CREATE TABLE "employees" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "position" varchar NOT NULL,
  "salary" float NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "employees"."salary" IS 'must be positive';