package model

type Articke struct {
	Title string    `db:"type:varchar(100);not null" json:"title"`
	Tag Tag 
	Tid int `db:"type:int" json:"tid"`
	Content string `db:"type:longtext" json:"content"`
	Img string `db:"type:verchar(100)" json:"img"`
}