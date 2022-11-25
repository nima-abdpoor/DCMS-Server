CREATE TABLE "config"
(
    "id"                bigserial PRIMARY KEY,
    "is_live"           boolean,
    "synctype"          varchar,
    "valid_request_url" varchar,
    "url_id_first"      varchar,
    "url_id_second"     varchar,
    "regex"             varchar,
    "start_index"       int,
    "finish_index"      int
);