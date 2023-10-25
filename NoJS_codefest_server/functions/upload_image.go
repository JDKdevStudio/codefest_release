package functions

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func UploadImageFunction(file *multipart.FileHeader, allowedExtensions []string) (string, error) {
	// Obtén la extensión del archivo
	ext := filepath.Ext(file.Filename)

	// Comprueba si la extensión está en la lista de extensiones permitidas
	var isAllowedExtension bool
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			isAllowedExtension = true
			break
		}
	}

	if !isAllowedExtension {
		return "", fmt.Errorf("la extensión del archivo no es válida")
	}

	// Obtén la ruta del ejecutable
	executablePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	// Carpeta de destino donde se almacenará la imagen
	destinationFolder := filepath.Join(filepath.Dir(executablePath), "public")

	// Nombre de archivo deseado
	destinationFilename := generateUniqueName() + ext

	// Ruta completa del archivo de destino
	destinationPath := filepath.Join(destinationFolder, destinationFilename)

	// Crea la carpeta de destino si no existe
	if err := os.MkdirAll(destinationFolder, 0755); err != nil {
		return "", err
	}

	// Abre el archivo de origen
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Crea el archivo de destino
	dst, err := os.Create(destinationPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copia el contenido desde el archivo de origen al archivo de destino
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	// Retorna la ruta completa del archivo de destino para su posterior uso si es necesario
	return destinationFilename, nil
}

func generateUniqueName() string {
	// Crear un generador de números aleatorios local
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Obtener un número aleatorio para evitar colisiones
	randomNumber := random.Intn(1000)

	// Obtener la hora actual en nanosegundos
	timestamp := time.Now().UnixNano()

	// Construir la cadena única
	uniqueName := fmt.Sprintf("%d_%d", timestamp, randomNumber)

	return uniqueName
}
