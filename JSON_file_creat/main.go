package main

import (
	// "context"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	shell "github.com/ipfs/go-ipfs-api"
)

// TimeSeriesDatum is the structure used to store a single time series data
type TimeSeriesDatum struct {
	Id_number    uint64 `json:"id"`
	Name string  `json:"name"`
	
}	

func main() {
	//
	// Connect to your local IPFS deamon running in the background.
	//

	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")

	//
	// Add the file to IPFS
	//

	tsd := &TimeSeriesDatum{
		Id_number: 2123198,
		Name: "Ben Brown",
		
	}
	tsdBin, _ := json.Marshal(tsd)
	reader := bytes.NewReader(tsdBin)

	cid, err := sh.Add(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s\n", cid)

	//
	// Get the data from IPFS and output the contents into `struct` format.
	//

	data, err := sh.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	// ...so we convert it to a string by passing it through
	// a buffer first. A 'costly' but useful process.
	// https://golangcode.com/convert-io-readcloser-to-a-string/
	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	newStr := buf.String()

	res := &TimeSeriesDatum{}
	json.Unmarshal([]byte(newStr), &res)
	fmt.Println(res)
}

