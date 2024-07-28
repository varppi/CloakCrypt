package main

import (
	"cloakcrypt/internal/encryption"
	"cloakcrypt/internal/filesystem"
	"context"
	"log"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Wails default
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

////

var CloakFile string
var EncryptFilePath string
var EncryptedFile string

func (a *App) SelectFile(mode int) string {
	var err error
	var vari *string

	if mode == 0 {
		vari = &CloakFile
	} else if mode == 1 {
		vari = &EncryptFilePath
	} else {
		vari = &EncryptedFile
	}

	*vari, err = runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Base(*vari)
}

func (a *App) EncryptFile(password string) bool {
	encryption.Key = password

	err := filesystem.Write(CloakFile, EncryptFilePath, a.ctx)
	if err != nil {
		log.Println("Error while encrypting file:", err)
		return false
	}

	return true
}

func (a *App) DecryptFile(password string) bool {
	encryption.Key = password

	err := filesystem.Read(EncryptedFile, a.ctx)
	if err != nil {
		log.Println("Error while decrypting file:", err)
		return false
	}

	return true
}
