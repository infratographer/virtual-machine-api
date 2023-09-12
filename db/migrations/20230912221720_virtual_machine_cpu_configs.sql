-- +goose Up
-- create "virtual_machine_cpu_configs" table
CREATE TABLE "virtual_machine_cpu_configs" ("id" character varying NOT NULL, "cores" bigint NOT NULL, "sockets" bigint NOT NULL, PRIMARY KEY ("id"));
-- modify "virtual_machines" table
ALTER TABLE "virtual_machines" ADD COLUMN "vm_cpu_config_id" character varying NOT NULL, ADD CONSTRAINT "virtual_machines_virtual_machine_cpu_configs_virtual_machine" FOREIGN KEY ("vm_cpu_config_id") REFERENCES "virtual_machine_cpu_configs" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- create index "virtual_machines_vm_cpu_config_id_key" to table: "virtual_machines"
CREATE UNIQUE INDEX "virtual_machines_vm_cpu_config_id_key" ON "virtual_machines" ("vm_cpu_config_id");

-- +goose Down
-- reverse: create index "virtual_machines_vm_cpu_config_id_key" to table: "virtual_machines"
DROP INDEX "virtual_machines_vm_cpu_config_id_key";
-- reverse: modify "virtual_machines" table
ALTER TABLE "virtual_machines" DROP CONSTRAINT "virtual_machines_virtual_machine_cpu_configs_virtual_machine", DROP COLUMN "vm_cpu_config_id";
-- reverse: create "virtual_machine_cpu_configs" table
DROP TABLE "virtual_machine_cpu_configs";
