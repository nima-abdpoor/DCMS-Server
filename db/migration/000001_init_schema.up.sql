CREATE TABLE "config"
(
    "id"        bigserial NOT NULL,
    "uid"       varchar   NOT NULL,
    "is_live"   boolean   NOT NULL,
    "sync_type" varchar   NOT NULL,
    PRIMARY KEY ("id", "uid")
);

CREATE TABLE "urlFirst"
(
    "id"       bigserial PRIMARY KEY NOT NULL,
    "uid"      varchar               NOT NULL,
    "url_hash" varchar
);

CREATE TABLE "urlSecond"
(
    "id"           bigserial PRIMARY KEY NOT NULL,
    "uid"          varchar,
    "url_hash"     varchar,
    "regex"        nvarchar,
    "start_index"  int,
    "finish_index" int
);

CREATE TABLE "requestUrl"
(
    "id"          bigserial PRIMARY KEY NOT NULL,
    "uid"         varchar               NOT NULL,
    "request_url" varchar
);

CREATE INDEX ON "config" ("uid");

CREATE INDEX ON "urlFirst" ("uid");

CREATE INDEX ON "urlSecond" ("uid");

CREATE INDEX ON "requestUrl" ("uid");

ALTER TABLE "urlFirst"
    ADD FOREIGN KEY ("uid") REFERENCES "config" ("id");

ALTER TABLE "urlSecond"
    ADD FOREIGN KEY ("uid") REFERENCES "config" ("id");

ALTER TABLE "requestUrl"
    ADD FOREIGN KEY ("uid") REFERENCES "config" ("id");