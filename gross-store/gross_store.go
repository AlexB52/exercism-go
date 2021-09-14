package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] 	 	= 6
	units["dozen"] 						 	= 12
	units["small_gross"] 			 	= 120
	units["gross"] 						 	= 144
	units["great_gross"] 			 	= 1728

	return units
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
