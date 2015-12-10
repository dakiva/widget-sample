-- +goose Up

CREATE TABLE widget (
       widget_id      bigserial,
       widget_uuid    uuid NOT NULL,
       widget_name    text NOT NULL,
       created_on_ts  timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
       modified_on_ts timestamp NOT NULL DEFAULT (now() at time zone 'utc'),
       CONSTRAINT pk_widget PRIMARY KEY(widget_id)
);

CREATE UNIQUE INDEX ix_widget_widget_uuid ON widget (
       widget_uuid
);
