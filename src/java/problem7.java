// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// Problem: What is the 1 millionth prime number?

public class problem7 {
	public static void main(String[] args) {
		// The number of the prime we are trying to get
		int targetPrime = 1000000;

		// Use the info the problem gave; start with prime #6, which is 13
		// These variables are mutable, as defined by the "mut" keyword
		// All variables in Rust are "final" by default, unless we declare them "mut"
		int currentPrime = 6;
		int currentNum = 13;

		// Keep going until we are at the target prime number
		while (currentPrime < targetPrime) {
			// Increase the value of the current number that we are on by 2
			// Increase by 2 because there is no sense in ever checking even numbers
			currentNum+=2;

			// If the current number is prime, increment the number of primes that we have hit
			if (isPrime(currentNum)) {
				currentPrime+=1;
				System.out.println("Prime number #" + currentPrime + " found: " + currentNum);
			}
		}
	}

	// Returns true if number is prime, false if number is not
	public static boolean isPrime(int num) {
		// Iterates from 3 up to the sqrt of the number + 1, because:
		// If n is not divisible by any number from 0 to sqrt(n), then n is prime
		for (int i = 3; i < Math.sqrt(num)+1; i++) {
			if (num % i == 0) {
				return false;
			}
		}
		return true;
	}
}