package filesystem

import (
	"bufio"
	"bytes"
	"cloakcrypt/internal/encoder"
	"cloakcrypt/internal/encryption"
	"context"
	"errors"
	"io"
	"math/rand"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func Write(original, encrypt string, ctx context.Context) error {
	//Initialize stuff
	ogHandle, err := os.Open(original)
	if err != nil {
		return err
	}
	defer ogHandle.Close()

	stats, err := os.Stat(original)
	if err != nil {
		return err
	}

	size := int(stats.Size())
	if size <= 1024 {
		return errors.New("too small file")
	}

	encryptHandle, err := os.Open(encrypt)
	if err != nil {
		return err
	}
	defer encryptHandle.Close()
	////

	//Read last 1024 bytes as a cover and return to start of file
	lastChunk := make([]byte, 1024)
	_, err = ogHandle.Seek(-1024, io.SeekEnd)
	if err != nil {
		return err
	}
	_, err = ogHandle.Read(lastChunk)
	if err != nil {
		return err
	}
	_, err = ogHandle.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	////

	saveFile, err := runtime.SaveFileDialog(ctx, runtime.SaveDialogOptions{})
	if err != nil {
		return err
	}

	saveHandle, err := os.OpenFile(saveFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer saveHandle.Close()

	//Write the original file to the encrypted file
	for {
		chunk := make([]byte, 1024)
		n, err := ogHandle.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		chunk = chunk[0:n]

		_, err = saveHandle.Write(chunk)
		if err != nil {
			return err
		}
	}
	////

	//Write the encrypted chunks
	for {
		chunkSize := rand.Intn(5000-500) + 500 //Random encrypted chunk size
		chunk := make([]byte, chunkSize)
		n, err := encryptHandle.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		chunk = chunk[0:n]

		encryptedChunk, err := encryption.Encrypt(chunk)
		if err != nil {
			return err
		}

		_, err = saveHandle.Write(append(encoder.Encode(encryptedChunk), 0))
		if err != nil {
			return err
		}
	}
	////

	saveHandle.Write(lastChunk) //Write the last 1024 bytes of the original file as cover

	return nil
}

func Read(encryptFile string, ctx context.Context) error {
	//Get handle and info
	encryptHandle, err := os.Open(encryptFile)
	if err != nil {
		return err
	}
	defer encryptHandle.Close()

	stats, err := os.Stat(encryptFile)
	if err != nil {
		return err
	}

	size := int(stats.Size())
	if size <= 1024*2 {
		return errors.New("not encrypted")
	}
	////

	//Get ending
	_, err = encryptHandle.Seek(-1024, io.SeekEnd)
	if err != nil {
		return err
	}

	falseEnding := make([]byte, 1024)
	_, err = encryptHandle.Read(falseEnding)
	if err != nil {
		return err
	}
	////

	//Get start
	_, err = encryptHandle.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	var start int
	chunkSize := 1024
	for {
		chunk := make([]byte, chunkSize)
		n, err := encryptHandle.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		chunk = chunk[:n]
		if len(chunk) < chunkSize && start == 0 {
			return errors.New("not encrypted")
		}
		if loc := bytes.Index(chunk, falseEnding[1000:1024]); loc != -1 {
			curLoc, err := encryptHandle.Seek(0, io.SeekCurrent)
			if err != nil {
				return err
			}
			start = int(curLoc) + loc + 24 - 1024
			break

		}
	}
	////

	//Go back to beginning of file
	_, err = encryptHandle.Seek(int64(start), io.SeekStart)
	if err != nil {
		return err
	}
	////

	//Create decrypted file
	saveFile, err := runtime.SaveFileDialog(ctx, runtime.SaveDialogOptions{})
	if err != nil {
		return err
	}

	saveHandle, err := os.Create(saveFile)
	if err != nil {
		return err
	}
	defer saveHandle.Close()
	////

	//Read encrypted chunks separated by null character, decrypt them and write them to the decrypted file.
	scanner := bufio.NewReader(encryptHandle)
	var canDecrypt bool
	for {
		chunk, err := scanner.ReadBytes(0)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		chunk = chunk[:len(chunk)-1]

		decompressed, err := encoder.Decode(chunk)
		if err != nil {
			return err
		}

		decrypted, err := encryption.Decrypt(decompressed)
		if err != nil {
			break
		}
		canDecrypt = true

		saveHandle.Write(decrypted)
	}

	if !canDecrypt {
		return errors.New("invalid key or not encrypted")
	}
	////

	return nil
}
