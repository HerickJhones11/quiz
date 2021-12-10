package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {

		log.Fatal("Unable to parse file as CSV for"+filePath, err)
	}
	return records
}
func main() {
	records := readCsvFile("./problems.csv")
	sum := 0
	var input string
	count := len(records)
	
	fmt.Println("Pressione alguma tecla")
	fmt.Scanln(&input)
	
	if input != "" {
		timer := time.NewTimer(30 * time.Second)
		go func ()  {
			<- timer.C
			fmt.Println("Você acertou", sum, "de", count)
			os.Exit(0)
		}()
	}
	


	for i := 0; i < count; i++ {
		fmt.Println("Digite o valor de:")
		fmt.Println(records[i][0])
		fmt.Scanln(&input)
		if input == records[i][1] {
			sum++
		}
	}
	fmt.Println("Você acertou", sum, "de", count)


}
