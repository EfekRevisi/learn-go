package prototypefactory

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{
			"",
			"Developer",
			6000,
		}
	case Manager:
		return &Employee{
			"",
			"Manager",
			8000,
		}
	default:
		panic("Unsupported role")
	}
}
