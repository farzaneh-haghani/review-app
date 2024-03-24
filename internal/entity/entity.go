package entity

type DBModel interface{
	table() string
}