package unzipper

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

// ReadZipFile reads a zip archive and returns a map where the keys are the file names
// and the values are the contents of the files as []byte.
func ReadZipFile(zipFilePath string) (map[string][]byte, error) {
    // Open the zip file
    r, err := zip.OpenReader(zipFilePath)
    if err != nil {
        return nil, fmt.Errorf("failed to open zip file: %v", err)
    }
    defer r.Close()

    // Create a map to store file contents
    filesContent := make(map[string][]byte)

    // Iterate through each file in the zip archive
    for _, f := range r.File {
        // Open the file inside the zip archive
        rc, err := f.Open()
        if err != nil {
            return nil, fmt.Errorf("failed to open file %s: %v", f.Name, err)
        }

        // Read the file contents into a []byte
        content, err := io.ReadAll(rc)
        if err != nil {
            rc.Close()
            return nil, fmt.Errorf("failed to read file %s: %v", f.Name, err)
        }
        rc.Close()

        // Store the content in the map
        filesContent[f.Name] = content
    }

    return filesContent, nil
}

func Unzip() {
    // Example usage
    zipFilePath := "example.zip"

    // Ensure the zip file exists
    if _, err := os.Stat(zipFilePath); os.IsNotExist(err) {
        log.Fatalf("zip file %s does not exist", zipFilePath)
    }

    // Read the zip file
    filesContent, err := ReadZipFile(zipFilePath)
    if err != nil {
        log.Fatalf("error reading zip file: %v", err)
    }

    // Print the contents of each file
    for fileName, content := range filesContent {
        fmt.Printf("File: %s\n", fileName)
        fmt.Printf("Content:\n%s\n", string(content))
    }
}
