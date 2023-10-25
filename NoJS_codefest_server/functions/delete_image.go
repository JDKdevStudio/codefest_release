package functions

import (
	"os"
	"path/filepath"
)

func DeleteImageHandler(filename string) error {
	// Obtén la ruta del ejecutable
	executablePath, err := os.Executable()
	if err != nil {
		return err
	}

	// Carpeta donde se encuentra la imagen
	imageFolder := filepath.Join(filepath.Dir(executablePath), "public")

	// Ruta completa de la imagen a borrar
	imagePath := filepath.Join(imageFolder, filename)

	// Verifica si el archivo existe antes de intentar borrarlo
	_, err = os.Stat(imagePath)
	if os.IsNotExist(err) {
		return err
	}
	// Intenta borrar la imagen
	if err := os.Remove(imagePath); err != nil {
		return err
	}
	return nil
}
