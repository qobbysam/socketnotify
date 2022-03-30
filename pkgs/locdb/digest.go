package locdb

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"os"
// )

// // type OutPut struct {
// // 	Name         string
// // 	Phone        string
// // 	Email        string
// // 	Mc           string
// // 	Dot          string
// // 	Street       string
// // 	State        string
// // 	City         string //Phone string
// // 	DotApplyDate string
// // 	McAppyDate   string
// // 	McGrantDate  string
// // 	PowerUnit    string
// // 	DriverTotal  string
// // 	McsFileDate  string
// // }

// type Check struct {
// 	Name  string
// 	Email string
// }

// func CsvToList(filepath string) []Check {
// 	out := make([]Check, 0)
// 	f, err := os.Open(filepath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	csvReader := csv.NewReader(f)
// 	data, err := csvReader.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for i, line := range data {
// 		if i > 0 { // omit header line
// 			//var rec Input_dot

// 			if len(line) != 5 {
// 				fmt.Println("not 5")
// 			} else {
// 				var nn Check
// 				nn.Name = line[0]
// 				nn.Email = line[1]
// 				out = append(out, nn)

// 			}
// 		}
// 	}

// 	return out
// }

// func CompareList(A, B []Check) []Check {
// 	out := make([]Check, 0)

// 	mb := make(map[string]struct{}, len(B))

// 	for _, x := range B {
// 		key := x.Email
// 		mb[key] = struct{}{}
// 	}

// 	for _, v := range A {
// 		key := v.Email
// 		if _, found := mb[key]; !found {
// 			out = append(out, v)
// 		}
// 	}

// 	return out
// }
