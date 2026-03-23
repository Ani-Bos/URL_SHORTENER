package utilities

import (
	"math"
)

const bas62str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func EncodeBase62(number int) string {
	ans := ""
	for number > 0 {
		rem := number % 62
		ans = string(bas62str[rem]) + ans
		number /= 62
	}
	return ans
}

func DecodeBase62(str string) int {
	res := 0
	n:=len(str)
	//make a map of char , int
	mp:=make(map[rune]int)
	for i, ch := range bas62str{
		mp[ch] = i
	}
	for i, ch := range str{
		idx := mp[ch]
		res = res + idx*int(math.Pow(62, float64(n-i-1)))
	}
	return res
}