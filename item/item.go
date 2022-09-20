package item

import "errors"

type Item struct {
	name  string
	price float64
}

func NewItem(name string, price float64) (Item, error) {
	if ok := validateName(name); !ok {
		return Item{}, errors.New("item name cannot be empty")
	}

	if ok := validatePrice(price); !ok {
		return Item{}, errors.New("item price cannot be less than or equal to zero")
	}
	return Item{name: name, price: price}, nil
}

func validateName(name string) bool {
	return len(name) != 0
}

func validatePrice(price float64) bool {
	return price > 0
}

func (i Item) GetPrice() float64 {
	return i.price
}

func (i Item) GetName() string {
	return i.name
}
