package ideaAlgo

// Сдвигает ключ циклически на 25 бит влево
func shiftKey(hk uint64, lk uint64) (uint64, uint64) {
	newHk := (hk << 25) | (lk >> 39)
	newLk := (lk << 25) | (hk >> 39)

	return newHk, newLk
}

// Формирует матрицу ключей из изходного 128 битного ключа
func encryptionKeyMatrix(hk uint64, lk uint64) [9][6]uint16 {
	var matrix [9][6]uint16
	var keyVector [52]uint16
	genNum := 0
	for i := 0; i < 52; i++ {
		genNum++
		switch genNum {
		case 1:
			keyVector[i] = uint16(hk >> 48)
		case 2:
			keyVector[i] = uint16(hk >> 32)
		case 3:
			keyVector[i] = uint16(hk >> 16)
		case 4:
			keyVector[i] = uint16(hk)
		case 5:
			keyVector[i] = uint16(lk >> 48)
		case 6:
			keyVector[i] = uint16(lk >> 32)
		case 7:
			keyVector[i] = uint16(lk >> 16)
		case 8:
			keyVector[i] = uint16(lk)
			hk, lk = shiftKey(hk, lk)
			genNum = 0
		}

	}

	idx := 0
OuterLoop:
	for i := 0; i < 9; i++ {
		for j := 0; j < 6; j++ {
			matrix[i][j] = keyVector[idx]
			idx += 1
			if idx == 52 {
				break OuterLoop
			}
		}
	}

	return matrix
}

// Вычисление мультипликативной инверсии по модулю 65537
//func modInverse(a uint16) uint16 {
//	if a == 0 {
//		return 0 // IDEA определяет 0 как 65536
//	}
//
//	// Расширенный алгоритм Евклида
//	t, newT := 0, 1
//	r, newR := 65537, int(a)
//
//	for newR != 0 {
//		quotient := r / newR
//		t, newT = newT, t-quotient*newT
//		r, newR = newR, r-quotient*newR
//	}
//
//	if t < 0 {
//		t += 65537
//	}
//	return uint16(t)
//}

//// Формирование матрицы подключей для дешифрации
//func decryptionKeyMatrix(keyMatrix [9][6]uint16) [9][6]uint16 {
//	var result [9][6]uint16
//
//	// Обратный порядок и вычисление инверсий
//	for i := 0; i < 9; i++ {
//		result[i][0] = modInverse(keyMatrix[8-i][0]) // K1⁻¹
//		result[i][3] = modInverse(keyMatrix[8-i][3]) // K4⁻¹
//
//		if i == 8 {
//			result[i][1] = -keyMatrix[0][1] // K2
//			result[i][2] = -keyMatrix[0][2] // K3
//		} else if i == 0 {
//			result[i][1] = -keyMatrix[8][1]
//			result[i][2] = -keyMatrix[8][2]
//		} else {
//			result[i][1] = -keyMatrix[8-i][2] // K3
//			result[i][2] = -keyMatrix[8-i][1] // K2
//		}
//
//		if i < 8 {
//			result[i][4] = keyMatrix[7-i][4] // K5
//			result[i][5] = keyMatrix[7-i][5] // K6
//		}
//	}
//
//	return result
//}
