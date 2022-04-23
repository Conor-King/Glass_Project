package main

import (
	// "context"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"github.com/google/uuid"
	shell "github.com/ipfs/go-ipfs-api"
	//crypto
	 "crypto/aes"
	 "crypto/cipher"
 	"encoding/base64"
 	
)



var myslice []string

// TimeSeriesDatum is the structure used to store a single time series data
type TimeSeriesDatum struct {
	Id_number string `json:"NINO"`
	Name      string `json:"name"`
	Address   string `json:"address"`
}




//cypto stuff




var cbytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"
func Encode(b []byte) string {
 return base64.StdEncoding.EncodeToString(b)
}


//decrypt
func Decode(s string) []byte {
 data, err := base64.StdEncoding.DecodeString(s)
 if err != nil {
  panic(err)
 }
 return data
} 


// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
 block, err := aes.NewCipher([]byte(MySecret))
 if err != nil {
  return "", err
 }
 cipherText := Decode(text)
 cfb := cipher.NewCFBDecrypter(block, cbytes)
 plainText := make([]byte, len(cipherText))
 cfb.XORKeyStream(plainText, cipherText)
 return string(plainText), nil
}



// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
 block, err := aes.NewCipher([]byte(MySecret))
 if err != nil {
  return "", err
 }
 plainText := []byte(text)
 cfb := cipher.NewCFBEncrypter(block, cbytes)
 cipherText := make([]byte, len(plainText))
 cfb.XORKeyStream(cipherText, plainText)
 return Encode(cipherText), nil
}



func IPFS_add(name_cal string, address_cal string, id_num_cal string) {
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
		Address:   address_cal,
		Name:      name_cal,
	}

	tsdBin, _ := json.Marshal(tsd)
	reader := bytes.NewReader(tsdBin)

	cid, err := sh.Add(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	id := uuid.New()
	idStr := id.String()
	pair := idStr + ":" + cid
	myslice = append(myslice, pair)
	fmt.Printf("id:cid %s:%s\n", idStr, cid)
	fmt.Printf("Data: %s \n \n", tsdBin)

	//
	// Get the data from IPFS and output the contents into `struct` format.
	//

	//data, err := sh.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	// ...so we convert it to a string by passing it through
	// a buffer first. A 'costly' but useful process.
	// https://golangcode.com/convert-io-readcloser-to-a-string/
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(data)
	//newStr := buf.String()

	//res := &TimeSeriesDatum{}
	//json.Unmarshal([]byte(newStr), &res)
	//fmt.Println(res)
}

func main() {

    
	arg := os.Args[1]

	if arg == "1" {
		fmt.Println("1. Please Enter Name: ")
		var name string
		fmt.Scanln(&name)

		fmt.Println("2. Please Enter Address: ")
		var address string
		fmt.Scanln(&address)
		fmt.Println("3. Please Enter Nino: ")
		var nino string
		fmt.Scanln(&nino)
		//name encrypt
		encname, err := (Encrypt(name, MySecret))
			if err != nil {
     				fmt.Println("error encrypting your classified text: ", err)
    				} 				   				
    				
    		//address encrypt
		encaddress, err := (Encrypt(address, MySecret))
			if err != nil {
     				fmt.Println("error encrypting your classified text: ", err)
    				}
    		//Nino encrypt
		encNino, err := (Encrypt(nino, MySecret))
			if err != nil {
     				fmt.Println("error encrypting your classified text: ", err)
    				}				
    		IPFS_add(encname, encaddress, encNino)
    		//Decrypt name
		decname, err := Decrypt(encname, MySecret)
 		fmt.Println("Encrypted name:",encname ,"		Decrypt name:",decname)
 		
 		//Decrypt address
 		decaddress, err := Decrypt(encaddress, MySecret)
 		fmt.Println("Encrypted address:",encaddress ,"		Decrypt address:",decaddress)
 		
 		//Decrypt Nino
 		decaNino, err := Decrypt(encNino, MySecret)
 		fmt.Println("Encrypted Nino:",encaddress ,"		Decrypt Nino:",decaNino)


		
	} else {
		
		//data from https://www.fakenamegenerator.com/gen-random-en-uk.php
		//moved from address A to adddress B
		IPFS_add("Lola Kirby", "9 Hart Road NORTHBOURNE CT14 5LS ", "EN 94 73 70 D")
		IPFS_add("Mathias Aachen", "Gotzkowskystraße 14 46348 Raesfeld", "NA")
		IPFS_add("Mujahid Juwain Najjar", "Coquimbo 9385 65002 Nuevo Berlín", "NA")
		IPFS_add("Baba Qga", "2/7 Grove Street", "NA")
		fmt.Println(myslice)
	}

}


