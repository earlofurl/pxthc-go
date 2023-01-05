-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-12-29T08:31:19.114Z

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "hashed_password" varchar(64) NOT NULL,
  "username" varchar(64) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "first_name" varchar(255) NOT NULL DEFAULT '',
  "last_name" varchar(255) NOT NULL DEFAULT '',
  "phone" varchar(26) UNIQUE,
  "role" varchar(255) NOT NULL DEFAULT 'user',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar(64) NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar(45) NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "strains" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar(255) NOT NULL DEFAULT '',
  "type" varchar(255) NOT NULL DEFAULT '',
  "yield_average" numeric(9,6),
  "terp_average_total" numeric(9,6),
  "terp_1" varchar(255),
  "terp_1_value" numeric(9,6),
  "terp_2" varchar(255),
  "terp_2_value" numeric(9,6),
  "terp_3" varchar(255),
  "terp_3_value" numeric(9,6),
  "terp_4" varchar(255),
  "terp_4_value" numeric(9,6),
  "terp_5" varchar(255),
  "terp_5_value" numeric(9,6),
  "thc_average" numeric(9,6),
  "total_cannabinoid_average" numeric(9,6),
  "light_dep_2022" varchar NOT NULL DEFAULT 'false',
  "fall_harvest_2022" varchar NOT NULL DEFAULT 'false',
  "quantity_available" numeric(9,6) NOT NULL DEFAULT 0
);

CREATE TABLE "retailer_locations" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar(255) UNIQUE NOT NULL,
  "address" varchar(255) NOT NULL DEFAULT '',
  "city" varchar(255) NOT NULL DEFAULT '',
  "state" varchar(255) NOT NULL DEFAULT '',
  "zip" varchar(10) NOT NULL DEFAULT '',
  "latitude" numeric(9,6) NOT NULL DEFAULT 90,
  "longitude" numeric(9,6) NOT NULL DEFAULT 135,
  "note" varchar(1024) NOT NULL DEFAULT '',
  "website" varchar(2083),
  "sells_flower" boolean NOT NULL DEFAULT false,
  "sells_prerolls" boolean NOT NULL DEFAULT false,
  "sells_pressed_hash" boolean NOT NULL DEFAULT false
);

CREATE TABLE "items" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "description" varchar(255) NOT NULL DEFAULT '',
  "is_used" boolean NOT NULL DEFAULT false,
  "item_type_id" bigint NOT NULL,
  "strain_id" bigint NOT NULL
);

CREATE TABLE "item_types" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "product_form" varchar(255) NOT NULL DEFAULT '',
  "product_modifier" varchar(255) NOT NULL DEFAULT '',
  "uom_default" bigint NOT NULL,
  "product_category_id" bigint NOT NULL
);

CREATE TABLE "package_tags" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "tag_number" varchar(24) NOT NULL,
  "is_assigned" boolean NOT NULL DEFAULT false,
  "is_provisional" boolean NOT NULL DEFAULT true,
  "is_active" boolean NOT NULL DEFAULT false,
  "assigned_package_id" bigint
);

CREATE TABLE "packages" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "tag_id" bigint,
  "package_type" varchar(255) NOT NULL DEFAULT '',
  "is_active" boolean NOT NULL DEFAULT true,
  "quantity" numeric(19,6) NOT NULL DEFAULT 0,
  "notes" varchar(1024) NOT NULL DEFAULT '',
  "packaged_date_time" timestamptz NOT NULL DEFAULT (now()),
  "harvest_date_time" timestamptz,
  "lab_testing_state" varchar(255) NOT NULL DEFAULT 'Untested',
  "lab_testing_state_date_time" timestamptz,
  "is_trade_sample" boolean NOT NULL DEFAULT false,
  "is_testing_sample" boolean NOT NULL DEFAULT false,
  "product_requires_remediation" boolean NOT NULL DEFAULT false,
  "contains_remediated_product" boolean NOT NULL DEFAULT false,
  "remediation_date_time" timestamptz,
  "received_date_time" timestamptz,
  "received_from_manifest_number" varchar(255),
  "received_from_facility_license_number" varchar(255),
  "received_from_facility_name" varchar(255),
  "is_on_hold" boolean NOT NULL DEFAULT false,
  "archived_date" timestamptz,
  "finished_date" timestamptz,
  "item_id" bigint,
  "provisional_label" varchar(255),
  "is_provisional" boolean NOT NULL DEFAULT false,
  "is_sold" boolean NOT NULL DEFAULT false,
  "ppu_default" numeric(19,4) NOT NULL DEFAULT 0,
  "ppu_on_order" numeric(19,4) NOT NULL DEFAULT 0,
  "total_package_price_on_order" numeric(19,4) NOT NULL DEFAULT 0,
  "ppu_sold_price" numeric(19,4) NOT NULL DEFAULT 0,
  "total_sold_price" numeric(19,4) NOT NULL DEFAULT 0,
  "packaging_supplies_consumed" boolean NOT NULL DEFAULT false,
  "is_line_item" boolean NOT NULL DEFAULT false,
  "order_id" bigint,
  "uom_id" bigint NOT NULL,
  "facility_location_id" bigint NOT NULL DEFAULT 1
);

CREATE TABLE "source_packages_child_packages" (
  "source_package_id" bigint NOT NULL,
  "child_package_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "lab_tests" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "test_name" varchar(255) NOT NULL DEFAULT '',
  "batch_code" varchar(255) NOT NULL DEFAULT '',
  "test_id_code" varchar(255) NOT NULL DEFAULT '',
  "lab_facility_name" varchar(255) NOT NULL DEFAULT '',
  "test_performed_date_time" timestamptz NOT NULL DEFAULT (now()),
  "test_completed" boolean NOT NULL DEFAULT false,
  "overall_passed" boolean NOT NULL DEFAULT false,
  "test_type_name" varchar(255) NOT NULL DEFAULT '',
  "test_passed" boolean NOT NULL DEFAULT false,
  "test_comment" varchar(255) NOT NULL DEFAULT '',
  "thc_total_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "thc_total_value" numeric(9,6) NOT NULL DEFAULT 0,
  "cbd_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "cbd_value" numeric(9,6) NOT NULL DEFAULT 0,
  "terpene_total_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "terpene_total_value" numeric(9,6) NOT NULL DEFAULT 0,
  "thc_a_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "thc_a_value" numeric(9,6) NOT NULL DEFAULT 0,
  "delta9_thc_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "delta9_thc_value" numeric(9,6) NOT NULL DEFAULT 0,
  "delta8_thc_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "delta8_thc_value" numeric(9,6) NOT NULL DEFAULT 0,
  "thc_v_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "thc_v_value" numeric(9,6) NOT NULL DEFAULT 0,
  "cbd_a_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "cbd_a_value" numeric(9,6) NOT NULL DEFAULT 0,
  "cbn_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "cbn_value" numeric(9,6) NOT NULL DEFAULT 0,
  "cbg_a_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "cbg_a_value" numeric(9,6) NOT NULL DEFAULT 0,
  "cbg_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "cbg_value" numeric(9,6) NOT NULL DEFAULT 0,
  "cbc_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "cbc_value" numeric(9,6) NOT NULL DEFAULT 0,
  "total_cannabinoid_percent" numeric(9,6) NOT NULL DEFAULT 0,
  "total_cannabinoid_value" numeric(9,6) NOT NULL DEFAULT 0
);

CREATE TABLE "lab_tests_packages" (
  "lab_test_id" bigint NOT NULL,
  "package_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "uoms" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar(32) NOT NULL DEFAULT '',
  "abbreviation" varchar(16) NOT NULL DEFAULT ''
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "scheduled_pack_date_time" timestamptz NOT NULL DEFAULT (now()),
  "scheduled_ship_date_time" timestamptz NOT NULL DEFAULT (now()),
  "scheduled_delivery_date_time" timestamptz NOT NULL DEFAULT (now()),
  "actual_pack_date_time" timestamptz,
  "actual_ship_date_time" timestamptz,
  "actual_delivery_date_time" timestamptz,
  "order_total" numeric(19,4) NOT NULL DEFAULT 0,
  "notes" varchar(1024) NOT NULL DEFAULT '',
  "status" varchar(255) NOT NULL DEFAULT '',
  "customer_name" varchar(255) NOT NULL DEFAULT ''
);

CREATE TABLE "package_adj_entries" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "package_id" bigint NOT NULL,
  "amount" numeric(19,6) NOT NULL,
  "uom_id" bigint NOT NULL
);

CREATE TABLE "package_adjustments" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "from_package_id" bigint NOT NULL,
  "to_package_id" bigint NOT NULL,
  "amount" numeric(19,6) NOT NULL,
  "uom_id" bigint NOT NULL
);

CREATE TABLE "facilities" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar(255) UNIQUE NOT NULL DEFAULT '',
  "license_number" varchar(255) UNIQUE NOT NULL DEFAULT ''
);

CREATE TABLE "facility_locations" (
  "id" bigserial PRIMARY KEY,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "name" varchar(255) NOT NULL DEFAULT '',
  "facility_id" bigint NOT NULL
);

CREATE UNIQUE INDEX ON "users" ("email");

CREATE UNIQUE INDEX ON "package_tags" ("tag_number");

CREATE UNIQUE INDEX ON "packages" ("tag_id");

CREATE INDEX ON "source_packages_child_packages" ("source_package_id");

CREATE INDEX ON "source_packages_child_packages" ("child_package_id");

CREATE INDEX ON "source_packages_child_packages" ("source_package_id", "child_package_id");

CREATE INDEX ON "lab_tests_packages" ("lab_test_id");

CREATE INDEX ON "lab_tests_packages" ("package_id");

CREATE INDEX ON "lab_tests_packages" ("lab_test_id", "package_id");

CREATE INDEX ON "package_adj_entries" ("package_id");

CREATE INDEX ON "package_adjustments" ("from_package_id");

CREATE INDEX ON "package_adjustments" ("to_package_id");

CREATE INDEX ON "package_adjustments" ("from_package_id", "to_package_id");

COMMENT ON COLUMN "package_adj_entries"."amount" IS 'can be negative or positive';

COMMENT ON COLUMN "package_adjustments"."amount" IS 'must be positive';

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "items" ADD FOREIGN KEY ("item_type_id") REFERENCES "item_types" ("id");

ALTER TABLE "items" ADD FOREIGN KEY ("strain_id") REFERENCES "strains" ("id");

ALTER TABLE "item_types" ADD FOREIGN KEY ("uom_default") REFERENCES "uoms" ("id");

ALTER TABLE "item_types" ADD FOREIGN KEY ("product_category_id") REFERENCES "product_categories" ("id");

ALTER TABLE "package_tags" ADD FOREIGN KEY ("assigned_package_id") REFERENCES "packages" ("id");

ALTER TABLE "packages" ADD FOREIGN KEY ("tag_id") REFERENCES "package_tags" ("id");

ALTER TABLE "packages" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

ALTER TABLE "packages" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "packages" ADD FOREIGN KEY ("uom_id") REFERENCES "uoms" ("id");

ALTER TABLE "packages" ADD FOREIGN KEY ("facility_location_id") REFERENCES "facility_locations" ("id");

ALTER TABLE "source_packages_child_packages" ADD FOREIGN KEY ("source_package_id") REFERENCES "packages" ("id");

ALTER TABLE "source_packages_child_packages" ADD FOREIGN KEY ("child_package_id") REFERENCES "packages" ("id");

ALTER TABLE "lab_tests_packages" ADD FOREIGN KEY ("lab_test_id") REFERENCES "lab_tests" ("id");

ALTER TABLE "lab_tests_packages" ADD FOREIGN KEY ("package_id") REFERENCES "packages" ("id");

ALTER TABLE "package_adj_entries" ADD FOREIGN KEY ("package_id") REFERENCES "packages" ("id");

ALTER TABLE "package_adj_entries" ADD FOREIGN KEY ("uom_id") REFERENCES "uoms" ("id");

ALTER TABLE "package_adjustments" ADD FOREIGN KEY ("from_package_id") REFERENCES "packages" ("id");

ALTER TABLE "package_adjustments" ADD FOREIGN KEY ("to_package_id") REFERENCES "packages" ("id");

ALTER TABLE "package_adjustments" ADD FOREIGN KEY ("uom_id") REFERENCES "uoms" ("id");

ALTER TABLE "facility_locations" ADD FOREIGN KEY ("facility_id") REFERENCES "facilities" ("id");
