package helper

import (
	"encoding/base64"
	"os"
)

func ImageToBase64(imagePath string) (string, error) {
	// Leer el archivo de imagen
	bytes, err := os.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	// Convertir a base64
	return base64.StdEncoding.EncodeToString(bytes), nil
}
