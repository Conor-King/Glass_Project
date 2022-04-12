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
	Id_number    string `json:"NINO"`
	Name string  `json:"name"`
	Address string  `json:"address"`
	
}	

func IPFS_add(name_cal string , address_cal string ,id_num_cal string) {
	//
	// Connect to your local IPFS deamon running in the background.
	//

	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")

	//
	// Add the file to IPFS
	//

	tsd := &TimeSeriesDatum{
		Id_number: id_num_cal,
		Address: address_cal,
		Name: name_cal,
		
		
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

func main(){

	//data from https://www.fakenamegenerator.com/gen-random-en-uk.php
	//moved from address A to adddress B
	IPFS_add("Paige Buckley","22 Grenoble Road BREAKISH IV42 4DH","XA 50 30 34 B")
	IPFS_add("Paige Buckley","14 Grenoble Road BREAKISH IV42 4DH","XA 50 30 34 B")
	
	
	IPFS_add("Connor Elliott","6 Great North Road ALWOODLEY PARK LS17 9ND","YR 79 78 34 B")
	IPFS_add("Noah Morrison","7 Dover Road WESTHAM BN24 7WW","AR 59 07 88 ")
	IPFS_add("Lola Kirby","9 Hart Road NORTHBOURNE CT14 5LS ","EN 94 73 70 D")
	IPFS_add("Mathias Aachen","Gotzkowskystraße 14 46348 Raesfeld","NA")
	IPFS_add("Mujahid Juwain Najjar","Coquimbo 9385 65002 Nuevo Berlín","NA")
	
	
	
}
