package food

type Food struct {
	name string
	qty  int
}

func NewFood(name string, qty int) *Food {
	newFood := new(Food)
	newFood.name = name
	newFood.qty = qty
	return newFood
}
