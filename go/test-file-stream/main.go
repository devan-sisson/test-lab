package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// ReadZipFromStream reads a zip archive from an io.Reader and returns a map where the keys are the file names
// and the values are the contents of the files as []byte.
func ReadZipFromStream(reader io.ReaderAt, size int64) (map[string][]byte, error) {
	// Create a zip reader
	zr, err := zip.NewReader(reader, size)
	if err != nil {
		return nil, fmt.Errorf("failed to create zip reader: %v", err)
	}

	// Create a map to store file contents
	filesContent := make(map[string][]byte)

	// Iterate through each file in the zip archive
	for _, f := range zr.File {
		// Open the file inside the zip archive
		rc, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %v", f.Name, err)
		}

		// Read the file contents into a []byte
		content, err := ioutil.ReadAll(rc)
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

func CompareArchive(files map[string][]byte, compareBasePath string) {
	for arName, arContent := range files {
		fmt.Printf("Comparing file: %s\n", arName)
		tstContent, err := os.ReadFile(fmt.Sprintf("%s/%s", compareBasePath, arName))
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}

		for i, b := range tstContent {
			if b != arContent[i] {
				log.Fatalf("bytes are not equal: %v control: %v", arContent[i], b)
			}
		}
	}
}

func main() {
	// The API endpoint that streams the zip file
	url := "http://localhost:8080/example"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("failed to get a successful response: %v", resp.Status)
	}

	// Read the response body into a buffer
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	// Read the zip archive from the buffer
	filesContent, err := ReadZipFromStream(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		log.Fatalf("error reading zip from stream: %v", err)
	}

	CompareArchive(filesContent, "tst")
	// Print the contents of each file
	//for fileName, content := range filesContent {
	//	fmt.Printf("File: %s\n", fileName)
	//	fmt.Printf("Content:\n%s\n", string(content))
	//}
}
