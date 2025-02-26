package textOps

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readkey(filename string) (uint64, uint64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	keyAsString := string(data)
	if len(keyAsString) != 34 || keyAsString[:2] != "0x" {
		return 0, 0, fmt.Errorf("неверный формат\n" +
			"Правильный формат: 0xN, где N -- 32 16=ричные цифры")
	}

	hkAsString := keyAsString[3:19]
	lkAsString := keyAsString[19:]
	hk, err := strconv.ParseUint(hkAsString, 16, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования 16-ричного числа %w", err)
	}

	lk, err := strconv.ParseUint(lkAsString, 16, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования 16-ричного числа %w", err)
	}

	return hk, lk, nil
}

func TxtToUint32Blocks(filename string) ([]uint32, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var result []uint32
	for _, block := range []rune(string(data)) {
		result = append(result, uint32(block))
	}

	return result, nil
}

func TxtToString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return string(data), nil
}

func Uint32BlocksAsBase16NumsToString(blocks []uint32) string {
	var sb strings.Builder

	for _, block := range blocks {
		sb.WriteString(fmt.Sprintf("0x%X ", block))
	}

	return sb.String()
}

func Uint32BlocksToString(blocks []uint32) string {
	var runes []rune

	for _, block := range blocks {
		runes = append(runes, rune(block))
	}

	return string(runes)
}
