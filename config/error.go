package config

import "errors"

var (
	ErrInvalidJson            = errors.New("invalid json response")
	ErrUnknownSignatureScheme = errors.New("unknown scheme sign scheme flag")
	ErrInvalidEncryptFlag     = errors.New("invalid encrypt flag")
	ErrInvalidAddress         = errors.New("invalid address")
)

//keypair

var (
	ErrPrivatekeyPrefixInvalid = errors.New("invalid private key prefix")
)
