// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// Problem: What is the 1 millionth prime number?
package main 

import (
	"fmt"
	"math"
)

func main(){
	// The number of the prime we are trying to get
	// This is a constant, as denoted by the "const" keyword
	const target_prime uint32 = 1000000;

	// Use the info the problem gave; start with prime #6, which is 13
	// These variables are mutable, as defined by the "var" keyword
	var current_prime uint32 = 6;
	var current_num uint32 = 13;

	// Keep going until we are at the target prime number
	// Go doesn't have a "while" loop, but rather has a variant "for" loop to fulfill the same role
	for current_prime < target_prime {
		// Increase the value of the current number that we are on
		current_num += 2;

		// If the current number is prime, increment the number of primes that we have hit
		if is_prime(current_num) {
			current_prime += 1;
		}
	}
	fmt.Printf("%v%v%v%v", "Prime number #", current_prime, " found: ", current_num);
}

// Returns true if number is prime, false if number is not
func is_prime(num uint32) bool {
	// Iterates from 3 up to the sqrt of the number + 1, because:
	// If n is not divisible by any number from 0 to sqrt(n), then n is prime
	var start uint32 = 3
	var end uint32 = uint32(math.Sqrt(float64(num))) + 1

	for start <= end  {
		if num % start == 0 {
			return false
		}
		start++
	}
	return true
}