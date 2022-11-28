CREATE TABLE "config"
(
    "id"        varchar PRIMARY KEY NOT NULL,
    "is_live"   boolean             NOT NULL,
    "sync_type" varchar             NOT NULL
);

CREATE TABLE "urlFirst"
(
    "id"       varchar PRIMARY KEY NOT NULL,
    "url_hash" varchar
);

CREATE TABLE "urlSecond"
(
    "id"           varchar PRIMARY KEY NOT NULL,
    "url_hash"     varchar,
    "regex"        varchar,
    "start_index"  int,
    "finish_index" int
);

CREATE TABLE "requestUrl"
(
    "id"          varchar PRIMARY KEY NOT NULL,
    "request_url" varchar
);

ALTER TABLE "urlFirst"
    ADD FOREIGN KEY ("id") REFERENCES "config" ("id");

ALTER TABLE "urlSecond"
    ADD FOREIGN KEY ("id") REFERENCES "config" ("id");

ALTER TABLE "requestUrl"
    ADD FOREIGN KEY ("id") REFERENCES "config" ("id");