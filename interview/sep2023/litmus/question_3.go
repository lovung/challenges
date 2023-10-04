package litmus

type Reptile interface {
	Lay() ReptileEgg
}

type ReptileCreator func() Reptile

type ReptileEgg struct {
	CreateReptile ReptileCreator

	created bool
}

func (egg *ReptileEgg) Hatch() Reptile {
	if egg.created {
		return nil
	}
	egg.created = true
	return egg.CreateReptile()
}

type FireDragon struct{}

func (fd FireDragon) Lay() ReptileEgg {
	return ReptileEgg{
		CreateReptile: func() Reptile {
			return FireDragon{}
		},
	}
}
