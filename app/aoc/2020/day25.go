package main

import (
	"errors"
	"fmt"
)

func day25(publicKey1, publicKey2 int) (int, error) {
	loopSize1, err := getLoopSize(publicKey1)
	if err != nil {
		return 0, err
	}

	loopSize2, err := getLoopSize(publicKey2)
	if err != nil {
		return 0, err
	}

	encryptionKey1 := transform(publicKey1, loopSize2)
	encryptionKey2 := transform(publicKey2, loopSize1)

	if encryptionKey1 == encryptionKey2 {
		return encryptionKey1, nil
	}
	return 0, errors.New("generated encryption keys don't match")
}

func transform(subjectNumber, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subjectNumber
		value %= 20201227
	}

	fmt.Printf("subject number: %d, loop size: %d -> encryption key: %d\n", subjectNumber, loopSize, value)
	return value
}

func getLoopSize(publicKey int) (int, error) {
	value := 1
	loopSize := 0
	for value != publicKey {
		value *= 7
		value %= 20201227
		loopSize++
		//fmt.Printf("val: %d -> loop size: %d\n", value, loopSize)
	}

	fmt.Printf("public key: %d -> loop size: %d\n", publicKey, loopSize)
	if value == publicKey {
		return loopSize, nil
	}

	return 0, errors.New("failed to get loop size")
}
