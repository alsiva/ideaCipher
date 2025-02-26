package ideaAlgo

func Encrypt(hk, lk, data uint64) uint64 {
	keyMatrix := encryptionKeyMatrix(hk, lk)
	encData := encryptBlock(data, keyMatrix)

	return encData
}
