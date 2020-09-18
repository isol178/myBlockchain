package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%064x%064x", s.R, s.S) //デコード処理で64バイトを仮定しているので足りない場合は0を補うようにする
}

func String2BigIntTuple(s string) (big.Int, big.Int) {
	bx, err := hex.DecodeString(s[:64])
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	by, err := hex.DecodeString(s[64:])
	if err != nil {
		log.Printf("ERROR: %v", err)
	}

	var bix big.Int
	var biy big.Int

	_ = bix.SetBytes(bx)
	_ = biy.SetBytes(by)

	return bix, biy
}

func SignatureFromString(s string) *Signature {
	x, y := String2BigIntTuple(s)
	return &Signature{R: &x, S: &y}
}

func PublicKeyFromString(s string) *ecdsa.PublicKey {
	x, y := String2BigIntTuple(s)
	return &ecdsa.PublicKey{elliptic.P256(), &x, &y}
}

func PrivateKeyFromString(s string, publicKey *ecdsa.PublicKey) *ecdsa.PrivateKey {
	b, err := hex.DecodeString(s[:])
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	var bi big.Int
	_ = bi.SetBytes(b)
	return &ecdsa.PrivateKey{*publicKey, &bi}
}
