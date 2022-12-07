CREATE TABLE "config"
(
    "id"        bigserial PRIMARY KEY,
    "is_live"   boolean NOT NULL,
    "sync_type" varchar NOT NULL
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