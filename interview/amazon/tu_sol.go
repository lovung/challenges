package amazon

// func getMaximumPoints5(days []int32, k int32) int64 {
// 	var (
// 		totalDays int32 = 0
// 	)
// 	for _, day := range days {
// 		totalDays += day
// 	}

// 	var (
// 		sprint                      int32 = 1
// 		point                       int32 = 1
// 		leftDay, rightDay           int32 = 1, 1
// 		leftDaySprint, leftDayPoint int32 = 1, 1
// 		maxPoints                   int64 = -1
// 		windowSum                   int64 = 0
// 	)
// 	for i := int32(1); i <= totalDays+(k-1); i++ {
// 		windowSum += int64(point)

// 		if rightDay-leftDay+1 > k {
// 			windowSum -= int64(leftDayPoint)

// 			leftDay++
// 			leftDayPoint++

// 			if leftDayPoint > days[leftDaySprint-1] {
// 				leftDayPoint = 1
// 				leftDaySprint++
// 			}
// 		}
// 		if rightDay-leftDay+1 == k && windowSum > maxPoints {
// 			maxPoints = windowSum
// 		}

// 		rightDay++
// 		if rightDay > totalDays {
// 			rightDay = rightDay % totalDays
// 		}

// 		point++
// 		if point > days[sprint-1] {
// 			sprint++
// 			if sprint > int32(len(days)) {
// 				sprint = sprint % int32(len(days))
// 			}
// 			point = 1
// 		}
// 	}

// 	return maxPoints
// }
