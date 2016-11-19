package src

type Length struct {
	Amount int
	Unit   int
}

const (
	Unit_M = 100
	Unit_CM = 1
)

func (l1 *Length) Equal(l2 Length) bool {
	if (l1.Amount * l1.Unit) == (l2.Amount * l2.Unit) {
		return true
	}
	return false
}
