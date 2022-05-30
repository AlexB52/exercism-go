package twobucket

type Bucket struct {
	name           string
	quantity, size int
}

func (b *Bucket) Fill() {
	b.quantity = b.size
}

func (b *Bucket) Empty() {
	b.quantity = 0
}

func (b *Bucket) Capacity() int {
	return b.size - b.quantity
}

func (b *Bucket) IsSize(size int) bool {
	return b.size == size
}

func (b *Bucket) IsFull() bool {
	return b.size == b.quantity
}

func (b *Bucket) IsEmpty() bool {
	return b.quantity == 0
}

func (b1 *Bucket) PourInto(b2 *Bucket) {
	var q int
	if b1.quantity <= b2.Capacity() {
		q = b1.quantity
	} else {
		q = b2.Capacity()
	}

	b1.quantity -= q
	b2.quantity += q
}
