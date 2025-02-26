package ideaAlgo

// Кодирует, либо декодирует данные, в зависимости от keyMatrix
func encryptBlock(data uint64, keyMatrix [9][6]uint16) uint64 {
	d1 := uint16(data >> 48)
	d2 := uint16(data >> 32)
	d3 := uint16(data >> 16)
	d4 := uint16(data)

	for i := 0; i < 8; i++ {
		keyVector := keyMatrix[i]

		k1 := keyVector[0]
		k2 := keyVector[1]
		k3 := keyVector[2]
		k4 := keyVector[3]
		k5 := keyVector[4]
		k6 := keyVector[5]

		a := mul(d1, k1)
		b := sum(d2, k2)
		c := sum(d3, k3)
		d := mul(d4, k4)
		e := xor(a, c)
		f := xor(b, d)
		t := mul(sum(f, mul(e, k5)), k6)

		d1 = xor(a, t)
		d2 = xor(c, t)
		d3 = xor(b, sum(mul(e, k5), t))
		d4 = xor(d, sum(mul(e, k5), t))
	}

	// Финальный раунд
	keyVector := keyMatrix[8]
	r1 := mul(d1, keyVector[0])
	r2 := sum(d3, keyVector[1])
	r3 := sum(d2, keyVector[2])
	r4 := mul(d4, keyVector[3])

	return (uint64(r1) << 48) | (uint64(r2) << 32) | (uint64(r3) << 16) | uint64(r4)
}
