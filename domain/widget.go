package domain

import "time"

type Widget struct {
	Id         int64     `db:"widget_id" json:"-"`
	Uuid       string    `db:"widget_uuid" json:"widget_id" description:"The unique identifier."`
	Name       string    `db:"widget_name" json:"widget_name" description:"The name of the widget."`
	CreatedOn  time.Time `db:"created_on_ts" json:"-"`
	ModifiedOn time.Time `db:"modified_on_ts" json:"-"`
}
