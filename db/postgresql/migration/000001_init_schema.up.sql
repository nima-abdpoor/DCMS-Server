CREATE TABLE "customer"
(
    "id"           bigserial PRIMARY KEY,
    "user"         varchar NOT NULL,
    "password"     varchar NOT NULL,
    "name"         varchar NOT NULL,
    "email"        varchar NOT NULL,
    "package_name" varchar NOT NULL,
    "sdk_uuid"     varchar NOT NULL,
    "secret_key"   varchar NOT NULL
);

CREATE TABLE "config"
(
    "id"                        bigserial PRIMARY KEY,
    "is_live"                   boolean   NOT NULL,
    "save_request"              boolean   NOT NULL,
    "save_response"             boolean   NOT NULL,
    "save_error"                boolean   NOT NULL,
    "save_success"              boolean   NOT NULL,
    "network_type"              varchar   NOT NULL,
    "repeat_interval"           bigserial NOT NULL,
    "repeat_interval_time_unit" varchar   NOT NULL,
    "requires_battery_not_low"  boolean   NOT NULL,
    "requires_storage_not_low"  boolean   NOT NULL,
    "requires_charging"         boolean   NOT NULL,
    "requires_device_idl"       boolean   NOT NULL
);

CREATE TABLE "urlfirst"
(
    "id"        bigserial PRIMARY KEY,
    "unique_id" bigserial NOT NULL,
    "url_hash"  varchar   NOT NULL
);

CREATE TABLE "urlsecond"
(
    "id"        bigserial PRIMARY KEY,
    "unique_id" bigserial NOT NULL,
    "url_hash"  varchar   NOT NULL
);

CREATE TABLE "regex"
(
    "id"           bigserial PRIMARY KEY,
    "url_id"       bigserial NOT NULL,
    "regex"        varchar   NOT NULL,
    "start_index"  int       NOT NULL,
    "finish_index" int       NOT NULL
);

CREATE TABLE "requesturl"
(
    "id"          bigserial PRIMARY KEY,
    "unique_id"   bigserial NOT NULL,
    "request_url" varchar   NOT NULL
);

ALTER TABLE "urlfirst"
    ADD FOREIGN KEY ("unique_id") REFERENCES "config" ("id");

ALTER TABLE "urlsecond"
    ADD FOREIGN KEY ("unique_id") REFERENCES "config" ("id");

ALTER TABLE "requesturl"
    ADD FOREIGN KEY ("unique_id") REFERENCES "config" ("id");

ALTER TABLE "regex"
    ADD FOREIGN KEY ("url_id") REFERENCES "urlsecond" ("id");