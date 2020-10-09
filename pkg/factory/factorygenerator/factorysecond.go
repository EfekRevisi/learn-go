package factorygenerator

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) SetName(name string) *Employee {
	return &Employee{
		name,
		f.Position,
		f.AnnualIncome,
	}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		position,
		annualIncome,
	}
}
