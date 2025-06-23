package main

import (
    "math"
)

// 判断一个数是否是质数
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    sqrtN := int(math.Sqrt(float64(n)))
    for i := 2; i <= sqrtN; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func Prime(n int) []int {
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
func main() {
    
}