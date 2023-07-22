package entity

type Group struct {
	ID          uint
	Title       string
	Description string

	UserIDs    []uint
	CategoryId uint
}

type Category struct {
	ID   uint
	Name string
}
