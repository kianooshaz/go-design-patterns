package main

type maverick struct {
	gun
}

func newMaverick() iGun {
	return &maverick{
		gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}
