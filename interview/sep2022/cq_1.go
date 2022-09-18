package sep2022

func getMinMoves(plates []int32) int32 {
	lightestPlate, heaviestPlate := plates[0], plates[0]
	lightestIdx, heaviestIdx := 0, 0
	for i := range plates {
		if lightestPlate > plates[i] {
			lightestPlate = plates[i]
			lightestIdx = i
		}
		if heaviestPlate < plates[i] {
			heaviestPlate = plates[i]
			heaviestIdx = i
		}
	}
	heaviestMoves := len(plates) - heaviestIdx - 1
	lightestMoves := lightestIdx
	if lightestIdx > heaviestIdx {
		lightestMoves--
	}
	return int32(heaviestMoves + lightestMoves)
}
