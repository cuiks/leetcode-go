package main

func countOperations(num1 int, num2 int) int {
	count := 0
	for num1 != 0 && num2 != 0 {
		count++
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
	}
	return count
}

func main() {

}
