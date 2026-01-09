package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	 

	"github.com/gin-gonic/gin"
)

// func readCsvFile(filePath string) [][]string {
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatal("Impossible de lire le fichier : ", err)
// 	}
// 	defer f.Close()

// 	csvReader := csv.NewReader(f)
// 	csvReader.Comma = ',' 

// 	records, err := csvReader.ReadAll()
// 	if err != nil {
// 		log.Fatal("Erreur de parsing CSV : ", err)
// 	}

// 	return records
// }

// func main() {
	
// 	filePath := "data/2018_01_Sites_mobiles_2G_3G_4G_France_metropolitaine_L93_ver2(3).csv"
// 	records := readCsvFile(filePath) //

// 	fmt.Println("Nombre total de lignes :", len(records))

// 	for i := 1; i <= 10; i++ { 
		
// 		line := records[i] 
		
// 		fmt.Println("Ligne :", i)
// 		fmt.Println("Opérateur :", line[0])
// 		fmt.Println("X  :", line[1])
// 		fmt.Println("Y  :", line[2])
// 		fmt.Println("2G :", line[3])
// 		fmt.Println("3G :", line[4])
// 		fmt.Println("4G :", line[5])
// 	}
// }

type Response struct {
    Type     string    `json:"type"`
    Features []Feature `json:"features"`
}

type Feature struct {
    Type     string   `json:"type"`
    Geometry Geometry `json:"geometry"`
	Coordinates Coordinates `json:"coordinates"`
}

type Geometry struct {
    Type        string    `json:"type"`
    Coordinates []float64 `json:"coordinates"`
}

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Impossible de lire le fichier : ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ',' 

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Erreur de parsing CSV : ", err)
	}

	return records
}

func main() {

	records := readCsvFile("data/2018_01_Sites_mobiles_2G_3G_4G_France_metropolitaine_L93_ver2(3).csv")
	r := gin.Default()

	r.GET("/XY", func(c *gin.Context) {
		response, err := http.Get("https://api-adresse.data.gouv.fr/search/?q=Lille&limit=1")
		if err != nil {
			fmt.Printf("The http request failed")
			return
		}
		defer response.Body.Close()

		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))

		var result Response
		err = json.Unmarshal(data, &result)
		fmt.Println(result.Features[0].Geometry.Coordinates[0], result.Features[0].Geometry.Coordinates[1])
		if err != nil {
			log.Fatal(err)
		}

    
		
	})

	r.Run(":8081")
}
