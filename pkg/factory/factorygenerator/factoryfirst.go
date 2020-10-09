package factorygenerator

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			name,
			position,
			annualIncome,
		}
	}
}
