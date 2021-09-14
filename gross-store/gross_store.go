package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
	}
}

// NewBill create a new bill
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem add item to customer bill
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, unitExists := units[unit]

	if !unitExists {return false}

	bill[item] = units[unit]
	return true
}

// RemoveItem remove item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	quantity, itemExists := bill[item]
	quantityToRemove, unitExists := units[unit]

	if !unitExists || !itemExists {
		return false
	} else if quantity < quantityToRemove {
		return false
	} else if quantity == quantityToRemove {
		delete(bill, item)
		return true
	}

	bill[item] -= quantityToRemove
	return true
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, itemExists := bill[item]
	return quantity, itemExists
}
