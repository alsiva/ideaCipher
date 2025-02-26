package cfb

import (
	"ideaCipher/internal/ideaAlgo"
)

func Crypt(hk, lk, iv uint64, sourceBlocks []uint32, encrypt bool) []uint32 {
	var targetBlocks []uint32

	for _, sourceBlock := range sourceBlocks {
		enc := ideaAlgo.Encrypt(hk, lk, iv) // 64 битное значение
		encBlock := uint32(enc >> 32)
		targetBlock := encBlock ^ sourceBlock
		targetBlocks = append(targetBlocks, targetBlock)

		if encrypt {
			iv = (iv << 32) | uint64(targetBlock)
		} else {
			iv = (iv << 32) | uint64(sourceBlock)
		}

	}
	return targetBlocks
}
