package model

type Tag struct {
	Name string `db:"type:varchar(10);not null" json:"name"`
}