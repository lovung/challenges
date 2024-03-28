package coupang

/*
Let's imagine a media file (for example, audio or video) that has a variable playback rate, that is each of its samples can be played back for a different duration. The playback timeline of the file is represented as a list of 2-tuples of positive integers. A 2-tuple (N, D) represents N consecutive samples each of which will be played back for D units of time.

As an example, the list (2,3), (5,4) represents a timeline with two 2-tuples. In this timeline, the first two samples are each played for 3 units of time and the next five samples are each played back for 4 units of time, resulting in a total of playback time of 26 (3+3+4+4+4+4+4) units.

A media editor could cut out an initial portion of time (X, a positive integer) from the original media file resulting in a new media timeline that is truncated from the beginning. For example, given a timeline (2,3), (5, 4) and X =1, the final media timeline will be (1,2), (1, 3), (5, 4) resulting in a total playback time of 25 (2+3+4+4+4+4+4) units. Some example test cases are included below:

input: timeline = (2,3), (5, 4) X=1: output: (1, 2), (1,3), (5, 4)
input: timeline = (2,3), (5, 4) X=2: output: (1, 1), (1, 3), (5, 4)
input: timeline = (2,3), (5, 4) X=3: output: (1, 3), (5, 4)
input: timeline = (2,3), (5, 4) X=4: output: (1, 2), (5, 4)

input: timeline = (2,3), (5, 4) X=5: output: (1, 1), (5, 4)
input: timeline = (2,3), (5, 4) X=6: output: (5, 4)
input: timeline = (2,3), (5, 4) X=7: output: (1, 3), (4, 4)
input: timeline = (2,3), (5, 4) X > =26: output: empty list

(2,3), (5,4) => (3, 3, 4, 4, 4, 4, 4) => (X=1) => (2, 3, 4, 4, 4, 4, 4) => (1, 2), (1,3), (5,4)
(2,3), (5,4) => (3, 3, 4, 4, 4, 4, 4) => (X=2) => (1, 3, 4, 4, 4, 4, 4) => (1, 1), (1,3), (5,4)
(2,3), (5,4) => (3, 3, 4, 4, 4, 4, 4) => (X=3) => (3, 4, 4, 4, 4, 4) => (1,3), (5,4)
(2,3), (5,4) => (3, 3, 4, 4, 4, 4, 4) => (X=6) => (4, 4, 4, 4, 4) => (5,4)
*/

type Tuple struct {
	N uint64 // count
	D uint64 // duration
}

func CutInitialPortion(input []Tuple, x uint64) []Tuple {
	// Solution 1: Flat all input
	// Solution 2: Greedy
	// remain of X > 0:
	//      try to get the first tuple (t) from input
	//      if t.N * t.D < remainX
	// 5*4 = 20
	// x=1 -> (1,3),(4,4)
	// x=1 -> (1,(20-x)%t.D),((20-x)/t.D,t.D)
	// x=4 -> (1,0),(4,4)
	// x=20 -> (1,0),(4,4)
	for i, t := range input {
		if x >= t.N*t.D {
			x -= t.N * t.D
			continue
		}
		tuple1 := Tuple{N: 1, D: (t.N*t.D - x) % t.D}
		tuple2 := Tuple{N: (t.N*t.D - x) / t.D, D: t.D}
		output := make([]Tuple, 0, len(input[i+1:])+2)
		if tuple1.D != 0 {
			output = append(output, tuple1)
		}
		if tuple2.N != 0 {
			output = append(output, tuple2)
		}
		output = append(output, input[i+1:]...)
		return output
	}
	return input
}
