-- +goose Up
-- create "virtual_machines" table
CREATE TABLE "virtual_machines" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "owner_id" character varying NOT NULL, "location_id" character varying NOT NULL, "userdata" character varying NULL, PRIMARY KEY ("id"));
-- create index "virtualmachine_created_at" to table: "virtual_machines"
CREATE INDEX "virtualmachine_created_at" ON "virtual_machines" ("created_at");
-- create index "virtualmachine_owner_id" to table: "virtual_machines"
CREATE INDEX "virtualmachine_owner_id" ON "virtual_machines" ("owner_id");
-- create index "virtualmachine_updated_at" to table: "virtual_machines"
CREATE INDEX "virtualmachine_updated_at" ON "virtual_machines" ("updated_at");

-- +goose Down
-- reverse: create index "virtualmachine_updated_at" to table: "virtual_machines"
DROP INDEX "virtualmachine_updated_at";
-- reverse: create index "virtualmachine_owner_id" to table: "virtual_machines"
DROP INDEX "virtualmachine_owner_id";
-- reverse: create index "virtualmachine_created_at" to table: "virtual_machines"
DROP INDEX "virtualmachine_created_at";
-- reverse: create "virtual_machines" table
DROP TABLE "virtual_machines";
