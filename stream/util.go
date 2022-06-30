package stream

// CoerceAtLeast Ensures that orig value is not less than the specified minimum.
// returns orig if it's greater than or equal to the minimum, or minimum otherwise.
func coerceAtLeast(orig, minimum int) int {
	if orig < minimum {
		return minimum
	} else {
		return orig
	}
}
