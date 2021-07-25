package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
	main  function is the entry point of the
	golang projects

	1. Creating a webserver which reads name from query Params
	and sends back the response as {"name":"<NAME_SEND_BY_CLIENT>"}
*/
func main() {
	println("Register root handler")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // register the handler
		names := r.URL.Query()["name"] // query parameters
		var name string                // i have a variable name of type string
		if len(names) == 1 {
			name = names[0]
		}
		m := map[string]string{"name": name} // dynamic variable declaration
		enc := json.NewEncoder(w)
		enc.Encode(m)
		// w.Write([]byte("Hello " + name)) // convert string to bytes , cascading
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

/*
func main() {
	// command line arguments
	path := flag.String("path", "myapp.log", "The path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for , Options are DEBUG,ERROR,LOG,CRITICAL")

	flag.Parse() // populate path and level variables , we have to  explicitly say that

	f, err := os.Open(*path) // pointer to path string memory location

	if err != nil {
		println("Error in opening file")
		log.Fatal(err)
	}

	defer f.Close()         // cleanup , when we are done with the file
	r := bufio.NewReader(f) // read the contents of the file
	for {                   // infinite loop
		s, err := r.ReadString('\n')
		if err != nil {
			println("Error in reading file")
			log.Fatal(err)
			break
		}
		if strings.Contains(s, *level) { // check if s contains ERROR
			fmt.Println(s) // print formatted string
		}
	}
}
*/
