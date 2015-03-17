package hap

import (
	"bytes"
	"errors"
)

// NewPassword creates a HomeKit compatible password string used for the bridge pairing setup.
// The argument string must be excatly 8 characters e.g. '01020304'
func NewPassword(str string) (string, error) {
	var password string
	if len(str) != 8 {
		return password, errors.New("String must be 8 characters long")
	}

	runes := bytes.Runes([]byte(str))
	first := string(runes[:3])
	second := string(runes[3:5])
	third := string(runes[5:])
	password = first + "-" + second + "-" + third

	return password, nil
}