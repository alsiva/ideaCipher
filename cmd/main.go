package main

import (
	"fmt"
	"ideaCipher/internal/cfb"
	"ideaCipher/internal/textOps"
	"math/rand"
)

func main() {

	fileName := "data/input.txt"
	keyFileName := "data/key.txt"

	iv := rand.Uint64()

	hk, lk, err := textOps.Readkey(keyFileName)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}

	sourceText, err := textOps.TxtToString(fileName)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	fmt.Printf("Содержимое файла: \n%s\n",
		sourceText)

	plainTextBlocks, err := textOps.TxtToUint32Blocks(fileName)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}

	cipherBlocks := cfb.Crypt(hk, lk, iv, plainTextBlocks, true)
	fmt.Printf("Зашифрованный набор данных\n%s\n",
		textOps.Uint32BlocksAsBase16NumsToString(cipherBlocks))

	decodedPlainTextBlocks := cfb.Crypt(hk, lk, iv, cipherBlocks, false)

	decodedText := textOps.Uint32BlocksToString(decodedPlainTextBlocks)
	fmt.Printf("Декодированный текст\n%s\n",
		decodedText)

	if sourceText == decodedText {
		fmt.Println("Исходный и декодированный текст совпадают")
	} else {
		fmt.Println("Исходный и декодированный текст не совпадают")
	}

}
