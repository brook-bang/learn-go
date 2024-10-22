package main

func sqrtRecursive(x, guess, prevGuess, epsilon float64) float64 {
	if diff := guess*guess - x; diff < epsilon && -diff < epsilon {
		return guess
	}

}

func main() {

}
