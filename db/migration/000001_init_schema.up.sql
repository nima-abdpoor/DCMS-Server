CREATE TABLE "config"
(
    "id"        varchar PRIMARY KEY NOT NULL,
    "is_live"   boolean             NOT NULL,
    "sync_type" varchar             NOT NULL
);

CREATE TABLE "urlfirst"
(
    "id"       varchar PRIMARY KEY NOT NULL,
    "url_hash" varchar
);

CREATE TABLE "urlsecond"
(
    "id"           varchar PRIMARY KEY NOT NULL,
    "url_hash"     varchar,
    "regex"        varchar,
    "start_index"  int,
    "finish_index" int
);

CREATE TABLE "requesturl"
(
    "id"          varchar PRIMARY KEY NOT NULL,
    "request_url" varchar
);

ALTER TABLE "urlfirst"
    ADD FOREIGN KEY ("id") REFERENCES "config" ("id");

ALTER TABLE "urlsecond"
    ADD FOREIGN KEY ("id") REFERENCES "config" ("id");

ALTER TABLE "requesturl"
    ADD FOREIGN KEY ("id") REFERENCES "config" ("id");