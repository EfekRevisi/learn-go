package factoryexercise

type ak47 struct {
	gun
}

func newAk47() IGun {
	return &ak47{
		gun: gun{
			name:  "AK 47",
			power: 4,
		},
	}
}
