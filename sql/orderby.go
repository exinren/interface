package sql

type Order string

const (
	ASC Order = "ASC"
	DESC Order = "DESC"
)

type OrderByItem struct {
	Field string
	Order
}

