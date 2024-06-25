package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Example function to simulate getting an io.ReadCloser
func getReadCloser() (io.ReadCloser, error) {
    // For demonstration purposes, we use a simple reader.
    // In a real-world scenario, this could be a file or any other source.
    return os.Open("testFile.txt")
}

func main() {
    // Initialize the Gin router
    router := gin.Default()

    // Define a route to handle the streaming response
    router.GET("/stream", func(c *gin.Context) {
        // Get the io.ReadCloser
        reader, err := getReadCloser()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading file"})
				}
        defer reader.Close()

        // Set appropriate headers for the response
        c.Header("Content-Type", "text/plain")

        // Stream the data to the response
        c.Stream(func(w io.Writer) bool {
            // Buffer to hold chunks of data
            buf := make([]byte, 1024)

            // Read from the reader
            n, err := reader.Read(buf)
            if err != nil && err != io.EOF {
                // Handle read error
                c.AbortWithStatus(http.StatusInternalServerError)
                return false
            }

            if n == 0 {
                // End of file or no data read
                return false
            }

            // Write the read data to the response writer
            if _, err := w.Write(buf[:n]); err != nil {
                // Handle write error
                c.AbortWithStatus(http.StatusInternalServerError)
                return false
            }

            return true
        })
    })

    // Run the server on port 8080
    router.Run(":8080")
}


// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	byteStr := []byte(`{"result":"success","message":"Job submitted","jobKey":"f7cc94d6-cdaa-464c-9f7c-ef3dcc93e9b9","jobInfo":{"workspace":"Studio","directoryPath":"Model Library/Simulation/latency","filename":"latency.py","command":"run","resourceConfig":{"name":"3XS","cpu":"1vCore","ram":"2Gb","run_rate":0.5},"tags":"editor","timeout":-1}}`)
// 	str := string(byteStr)
// 	jsony, _ := json.Marshal(str)
// 	jStr := string(jsony)

// 	fmt.Println(jStr)

// 	// str := "Hello, World!"
//     // substr := str[1:] // Slice from index 1 to the end of the string

//     // fmt.Println(substr) // Output: "ello, World!"
// }

// import (
//     "fmt"
//     "reflect"
// )

// type Person struct {
//     Name    string
//     Age     int
//     Address string
// }

// func main() {
//     person := Person{}

//     // Get the reflect.Value of the struct
//     val := reflect.ValueOf(person)

//     // Iterate over the fields of the struct
//     for i := 0; i < val.NumField(); i++ {
//         // Get the field value
//         fieldValue := val.Field(i)

//         // Get the field type
//         fieldType := val.Type().Field(i)

//         // Print the field name and value
//         fmt.Printf("%s: %v\n", fieldType.Name, fieldValue.Interface())
//     }
// }


// import (
// 	"fmt"
// 	"github.com/akamensky/argparse"
// 	"os"
// )

// func main() {
// 	// Create new parser object
// 	parser := argparse.NewParser("print", "Prints provided string to stdout")
// 	// Create string flag
// 	s := parser.String("s", "string", &argparse.Options{Required: true, Help: "String to print"})
// 	// Parse input
// 	err := parser.Parse(os.Args)
// 	if err != nil {
// 		// In case of error print error and print usage
// 		// This can also be done by passing -h or --help flags
// 		fmt.Print(parser.Usage(err))
// 	}
// 	// Finally print the collected string
// 	fmt.Println(*s)
// }

// import (
//     "fmt"
//     "regexp"
// )

// func main() {
//     // Define the regex pattern
//     pattern := `('(\\'|[^'])*'|"(\\"|[^"])*"|\/(\\\/|[^\/])*\/|(\\ |[^ ])+|[\w-]+)`

//     // Compile the regex pattern
//     re := regexp.MustCompile(pattern)

//     // Sample text
//     text := `This is a sample 'text with' some "quoted strings" and /a regex/ pattern. Also, some \escaped \characters and some-unquoted-text`

//     // Find all matches
//     matches := re.FindAllString(text, -1)

//     // Print all matches
//     for _, match := range matches {
//         fmt.Println(match)
//     }
// }






// import (
//     "encoding/json"
//     "fmt"
// )

// // Define a struct to represent the JSON data
// type Item struct {
//     ID    int    `json:"id"`
//     Name  string `json:"name"`
//     Price int    `json:"price"`
// }

// func main() {
//     // Simulated JSON response
//     jsonResponse := `[{"id":1,"name":"Item 1","price":10},{"id":2,"name":"Item 2","price":20},{"id":3,"name":"Item 3","price":30}]`

//     // Define a slice of structs to hold the unmarshaled data
//     var items []Item

//     // Unmarshal the JSON response into the slice of structs
//     err := json.Unmarshal([]byte(jsonResponse), &items)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     // Print the unmarshaled data
//     fmt.Println("Unmarshaled Items:")
//     for _, item := range items {
//         fmt.Printf("ID: %d, Name: %s, Price: %d\n", item.ID, item.Name, item.Price)
//     }
// }
