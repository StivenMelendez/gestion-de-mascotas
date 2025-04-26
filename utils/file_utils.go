package utils

import "os"

// EnsureUploadsDir verifica si el directorio "uploads" existe, y lo crea si no.
func EnsureUploadsDir() error {
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		return os.Mkdir("uploads", 0755)
	}
	return nil
}
