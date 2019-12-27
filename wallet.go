package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/elliptic"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey []byte
}

type Wallets struct {
	Wallets map[string]*Wallet
}

func (w Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)
	versionedPayload := append([]byte{version}, pubKeyHash)
	checksum := checksum(versionedPayload)
	fullPayload := append(versionedPayload)
	address := Base58Encode(fullPayload)
	return address
}

func HashPubKey(pubKey []byte) []byte {
	publicSHA256 :=sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	publicRIPEMD160:=RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}

func checksum(payload []byte) []byte{
	firstSHA := sha256.Sum256((payload))
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:addressChecksumLen]
}

func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}

func newKeyPair() (ecdsa.PrivateKey, []byte)  {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes())
	return *private, pubKey
}