package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
)

var filenames = []string{"one.txt", "two.txt", "sample.json"}

type filebytes []byte

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/:filename", func(c *gin.Context) {
		files := make(map[string]filebytes)
		filename := c.Param("filename")

		for _, f := range filenames {
			fileContent, err := ioutil.ReadFile(f)
			if err != nil {
				panic(err)
			}
			files[f] = fileContent
		}

		CreateArchiveStream(c, filename, files)

	})
	r.Run()
}

func CreateArchiveStream(c *gin.Context, name string, files map[string]filebytes) {
	c.Writer.Header().Set("Content-Type", "appliation/octet-stream")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", name))
	c.Stream(func(w io.Writer) bool {
		archive := zip.NewWriter(w)

		for k := range files {
			arfile, _ := archive.Create(k)
			content := bytes.NewReader(files[k])
			io.Copy(arfile, content)
		}

		archive.Close()

		return false
	})
}
