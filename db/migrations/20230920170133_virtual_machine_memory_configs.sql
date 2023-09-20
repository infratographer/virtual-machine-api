-- +goose Up
-- create "virtual_machine_memory_configs" table
CREATE TABLE "virtual_machine_memory_configs" ("id" character varying NOT NULL, "size" bigint NOT NULL DEFAULT 8:::INT8, PRIMARY KEY ("id"));
-- modify "virtual_machines" table
ALTER TABLE "virtual_machines" ADD COLUMN "vm_memory_config_id" character varying NOT NULL, ADD CONSTRAINT "virtual_machines_virtual_machine_memory_configs_virtual_machine" FOREIGN KEY ("vm_memory_config_id") REFERENCES "virtual_machine_memory_configs" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- create index "virtual_machines_vm_memory_config_id_key" to table: "virtual_machines"
CREATE UNIQUE INDEX "virtual_machines_vm_memory_config_id_key" ON "virtual_machines" ("vm_memory_config_id");

-- +goose Down
-- reverse: create index "virtual_machines_vm_memory_config_id_key" to table: "virtual_machines"
DROP INDEX "virtual_machines_vm_memory_config_id_key";
-- reverse: modify "virtual_machines" table
ALTER TABLE "virtual_machines" DROP CONSTRAINT "virtual_machines_virtual_machine_memory_configs_virtual_machine", DROP COLUMN "vm_memory_config_id";
-- reverse: create "virtual_machine_memory_configs" table
DROP TABLE "virtual_machine_memory_configs";
