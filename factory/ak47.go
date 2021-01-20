package main

type ak74 struct {
	gun
}

func newAk47() iGun {
	return &ak74{
		gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
