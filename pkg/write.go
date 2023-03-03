package pkg

import (
	"fmt"
	"os"
)

func WriteJson(data []byte) {

	filename := "stock_list.json"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("is not exist")
	} else {
		err := os.Remove(filename)
		if err != nil {
			fmt.Println("delete failed")
		} else {
			fmt.Println("delete", filename)
		}
	}

	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Creat file", filename)
	}

	defer fp.Close()
	_, err1 := fp.Write(data)
	if err1 != nil {
		panic(err1)
	}
}
