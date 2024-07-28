package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/sha3"
)

/*
Encryption: AES-256 GCM

How the 256 bit key is derived: Shake256 --> p*gC2;U( <password> <first 10 bytes of password hash> ^E!4>#hw
*/

var Key string
var salt = "p*gC2;U(^E!4>#hw"

func Encrypt(data []byte) ([]byte, error) {
	ogHash := make([]byte, 32)
	sha3.ShakeSum256(ogHash, []byte(Key))

	newPassword := []byte(salt[:8])
	newPassword = append(newPassword, []byte(Key)...)
	newPassword = append(newPassword, ogHash[:10]...)
	newPassword = append(newPassword, salt[8:]...)

	secKey := make([]byte, 32)
	sha3.ShakeSum256(secKey, newPassword)

	block, err := aes.NewCipher(secKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	rand.Read(nonce)

	cipheredRaw := gcm.Seal(nonce, nonce, data, nil)

	return cipheredRaw, nil
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	ogHash := make([]byte, 32)
	sha3.ShakeSum256(ogHash, []byte(Key))

	newPassword := []byte(salt[:8])
	newPassword = append(newPassword, []byte(Key)...)
	newPassword = append(newPassword, ogHash[:10]...)
	newPassword = append(newPassword, salt[8:]...)

	secKey := make([]byte, 32)
	sha3.ShakeSum256(secKey, newPassword)

	block, err := aes.NewCipher(secKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("too short message")
	}

	nonce := ciphertext[:gcm.NonceSize()]
	cipherdata := ciphertext[gcm.NonceSize():]

	return gcm.Open(nil, nonce, cipherdata, nil)
}
