package item

func NewItemApple() (Item, error) {
	return NewItem("apple", 12.0)
}

func NewItemWithoutName() (Item, error) {
	return NewItem("", 10.0)
}

func NewItemWithZeroOrNegativePrice() (Item, error) {
	return NewItem("apple", 0)
}

func NewItemOrange() (Item, error) {
	return NewItem("orange", 10.0)
}

func NewItemPineapple() (Item, error) {
	return NewItem("pineapple", 18.0)
}
