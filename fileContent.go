package main

import (
	"net/http"
	"os"
)

// GetFileContentType recieves a file path as string, evaluates its type
// and returns the type and an error (that is nil if everything went ok).
func GetFileContentType(file string) (string, error) {

	out, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err = out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always
	// returns a valid content-type by returning "application/octet-stream"
	// if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
