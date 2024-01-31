package main


import (
    "xib/albums-api/domain"
    
    "fmt"
    "os"

    "encoding/json"
    "io/ioutil"

)

type Albums struct {
    Albums []domain.Album `json:"albums"`
}

func main(){

	 jsonFile, err := os.Open("data.json")
	 if err != nil {
		 fmt.Println(err)
	 }
 
	 fmt.Println("Successfully Opened data.json")
	 defer jsonFile.Close()
 
	 byteValue, _ := ioutil.ReadAll(jsonFile)
 
	 var albums Albums
 
	 json.Unmarshal(byteValue, &albums)
 

	for _, al := range albums.Albums{
		fmt.Println(al)
	}

	 for i := 0; i < len(albums.Albums); i++ {
		 fmt.Println("Artist: " + albums.Albums[i].Artist)
	 }
 

}