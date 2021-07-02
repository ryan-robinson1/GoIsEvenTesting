package main

func IsEven(num int) bool {
	return num%2 == 0
}

//Tried to make it faster but benchmarking showed it was around the same speed
func IsEvenBitCheck(num int) bool {
	return num&1 == 0
}
