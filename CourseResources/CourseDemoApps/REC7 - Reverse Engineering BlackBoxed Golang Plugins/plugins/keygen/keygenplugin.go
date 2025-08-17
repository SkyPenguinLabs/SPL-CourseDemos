package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func A25410863kb318989a95307(username, password string) string {
	///// Prefix | Username [4]
	userPrefix := username
	if len(username) > 4 {
		userPrefix = username[:4]
	}

	///// Postfix | Password [4]
	passPostfix := password
	if len(password) > 4 {
		passPostfix = password[len(password)-4:]
	}

	////// Hash in sha256
	combined := userPrefix + passPostfix
	hash := sha256.Sum256([]byte(combined))
	hashStr := hex.EncodeToString(hash[:])

	///// License fmt
	part1 := userPrefix + hashStr[:12-len(userPrefix)]
	part2 := passPostfix + hashStr[12-len(userPrefix):16-len(passPostfix)]

	return part1 + "-" + part2
}

func A432410863kb318989a95307(licenseKey, username, password string) bool {
	expectedKey := A25410863kb318989a95307(username, password)
	return licenseKey == expectedKey
}
