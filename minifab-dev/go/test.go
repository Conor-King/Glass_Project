package main

import (
	"fmt"
	"strings"
)

func main() {
	//a := []string{"QmUF8tfcEQb26n6YmqctmogeUiKaWiiayPkRKMmGQByct1", "QmdVuEn92CAzLduo35FyFvw1vwUwDTdLizHg5xfqvJFwzV", "QmSA5vwC9NNSmmXK6cmxHhjrSdxgqv2LgSKcGjG5aXdbXt", "QmSnRa72fs3WdUuJaP3UQ28HUkje4djcapksV4SqPXFSaW"}
	//fmt.Println(a)
	
	//string got from the ledger
	var string1 = "uniqueid1:cid1 uniqueid2:cid2 uniquid3:cid3"
	
	//slice the string to multiple elements
	sliceS1 := strings.Split(string1, " ")

	//loop each elemnt and search for the uniqueid
	for i := range sliceS1 {
		res := strings.HasPrefix(sliceS1[i], "uniquid3")
		//res is true when the element begins with unidquid3
		if  res  {
		//do whatever when the element is found
			fmt.Println(sliceS1[i])
			break
		}
	}


//	fmt.Println(sliceS1)

	// for index, value := range a {
	// 	id := uuid.New()
	// 	idStr := id.String()
	// 	value = idStr + ":" + value 
	// 	fmt.Println("Index:", index, "Value:", value)

	// }
}