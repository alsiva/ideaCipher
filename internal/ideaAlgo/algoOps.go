package ideaAlgo

// Исключающее или
func xor(a, b uint16) uint16 {
	return a ^ b
}

// Сложение по модулю 2^16
func sum(a, b uint16) uint16 {
	return uint16(a + b)
}

// Умножение по модулю 2^16 + 1
func mul(a, b uint16) uint16 {
	return uint16((uint32(a) * uint32(b)) % ((1 << 16) + 1))
}
