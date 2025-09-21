package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	stringArr := [3]string{primary, secondary, tertiary}
	costArr := [3]int{len((primary)), len(secondary) + len(primary), len(tertiary) + len(primary) + len(secondary)}

	return stringArr, costArr
}
