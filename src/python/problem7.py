# By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
# Problem: What is the 1 millionth prime number?

# Returns true if number is prime, false if number is not
def is_prime(num):
	# Iterates from 3 up to the sqrt of the number + 1, because:
	# If n is not divisible by any number from 0 to sqrt(n), then n is prime
	for i in range(3, (int)(num**0.5) + 1):
		if num % i == 0:
			return False
	return True

# The number of the prime we are trying to get
target_prime = 1000000

# Use the info the problem gave start with prime #6, which is 13
# These variables are mutable, as defined by the "mut" keyword
# All variables in Rust are "final" by default, unless we declare them "mut"
current_prime = 6
current_num = 13

# Keep going until we are at the target prime number
while current_prime < target_prime:
	# Increase the value of the current number that we are on by 2
	# Increase by 2 because there is no sense in ever checking even numbers
	current_num+=2

	# If the current number is prime, increment the number of primes that we have hit
	if is_prime(current_num):
		current_prime+=1
		print 'Prime number #%(prime)s found: %(num)s' % {'prime': current_prime, 'num': current_num}