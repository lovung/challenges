package preparekit

// Link: https://www.hackerrank.com/challenges/one-week-preparation-kit-plus-minus/problem
import (
	"fmt"
	"io"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

var w io.Writer

func plusMinus(arr []int32) {
	// Write your code here
	cntP, cntN, cntZ := 0, 0, 0
	for i := range arr {
		switch {
		case arr[i] > 0:
			cntP++
		case arr[i] == 0:
			cntZ++
		case arr[i] < 0:
			cntN++
		}
	}
	sum := cntP + cntN + cntZ
	Print(fmt.Sprintf("%.6f\n", float64(cntP)/float64(sum)))
	Print(fmt.Sprintf("%.6f\n", float64(cntN)/float64(sum)))
	Print(fmt.Sprintf("%.6f\n", float64(cntZ)/float64(sum)))
}

// Print to writer
func Print(s string) {
	fmt.Fprintf(w, s)
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	n := int32(nTemp)

// 	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

// 	var arr []int32

// 	for i := 0; i < int(n); i++ {
// 		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
// 		checkError(err)
// 		arrItem := int32(arrItemTemp)
// 		arr = append(arr, arrItem)
// 	}

// 	plusMinus(arr)
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
