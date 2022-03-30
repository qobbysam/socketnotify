package locdb

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"time"
// )

// func CreateOut(state string) {
// 	basedir, _ := os.Executable()

// 	path := filepath.Dir(basedir)

// 	filename := state + ".csv"

// 	input_dot := filepath.Join(path, "data", "input", "dot", filename)

// 	input_mc := filepath.Join(path, "data", "input", "mc", filename)

// 	t := time.Now()
// 	out_name := state + string(t.Format("2006")) + ".csv"

// 	found_location := filepath.Join(path, "data", "output", "ffound", out_name)
// 	not_found_location := filepath.Join(path, "data", "output", "nnotfound", out_name)

// 	dot_in := CreateInputDot(input_dot)

// 	// fmt.Println(s)

// 	mc_in := CreateInputMc(input_mc)

// 	//searchmc := CreateSearchMC(mc_in)

// 	searchdotname := CreateSearchDotName(dot_in)

// 	searchdotdba := CreateSearchDotDBA(dot_in)

// 	searchmapname := make(map[int][]int)

// 	searchmapdba := make(map[int][]int)
// 	lastname := 0
// 	for _, v := range searchdotname {

// 		val, ok := searchmapname[v.Value]

// 		if !ok {
// 			mas := make([]int, 0)
// 			searchmapname[v.Value] = mas
// 			searchmapname[v.Value] = append(searchmapname[v.Value], v.Key)

// 			lastname = v.Value
// 		} else {
// 			searchmapname[v.Value] = append(val, v.Key)
// 		}
// 	}

// 	last := 0
// 	for _, v := range searchdotdba {
// 		val, ok := searchmapdba[v.Value]

// 		if !ok {
// 			mas := make([]int, 0)
// 			searchmapdba[v.Value] = mas
// 			searchmapdba[v.Value] = append(searchmapdba[v.Value], v.Key)
// 			last = v.Value
// 		} else {
// 			searchmapdba[v.Value] = append(val, v.Key)
// 		}
// 	}
// 	// for k, v := range searchmapdba {
// 	// 	fmt.Println("KEy ", k, " value ", v)
// 	// }
// 	// for k, v := range searchmapname {
// 	// 	fmt.Println("KEy ", k, " value ", v)
// 	// }
// 	fmt.Println("searh name size , ", len(searchmapname))
// 	fmt.Println("last map", searchmapdba[last])
// 	fmt.Println("last map name, ", searchmapname[lastname])
// 	//fmt.Println("last name, " , search)

// 	foundlist := make([]OutPut, 0)
// 	notfoundlist := make([]Input_mc, 0)
// 	total := 0
// 	notfoundprev := false

// 	for k, v := range mc_in {
// 		found_in_name := false
// 		found_in_dba := false
// 		ascn := ToAsciiTotal(v.Name)
// 		if notfoundprev {
// 			notfoundlist = append(notfoundlist, mc_in[k-1])
// 			notfoundprev = false
// 		} else {
// 			notfoundprev = true
// 			val, ok := searchmapname[ascn]

// 			if !ok {
// 				// found_in_name = true
// 				vall, okk := searchmapdba[ascn]

// 				if !okk {
// 					//found_in_map = true
// 				} else {
// 					cleanmcname := strings.ReplaceAll(v.Name, " ", "")
// 					//found_one := false
// 					for _, vvv := range vall {
// 						cleandotdba := strings.ReplaceAll(dot_in[vvv].DBA, " ", "")

// 						if cleanmcname == cleandotdba {

// 							var found OutPut

// 							found.Name = dot_in[vvv].Name
// 							found.Phone = dot_in[vvv].Phone
// 							found.Email = dot_in[vvv].Email
// 							found.Mc = mc_in[k].MC
// 							found.Dot = dot_in[vvv].Dot
// 							found.Street = dot_in[vvv].Street
// 							found.State = dot_in[vvv].State
// 							found.City = dot_in[vvv].City
// 							found.DotApplyDate = dot_in[vvv].ApplyDate
// 							found.McAppyDate = mc_in[k].ApplyDate
// 							found.McGrantDate = mc_in[k].GrantDate
// 							found.PowerUnit = dot_in[vvv].PowerUnit
// 							found.DriverTotal = dot_in[vvv].DriverTotal
// 							found.McsFileDate = dot_in[vvv].McsFileDate

// 							total++

// 							foundlist = append(foundlist, found)
// 							found_in_dba = true
// 							//notfoundprev = false
// 							//found_one = true
// 							fmt.Println("found one dba ", v.Name, "  ", dot_in[vvv].DBA)
// 							break
// 						}
// 					}

// 				}
// 			} else {
// 				cleanmcname := strings.ReplaceAll(v.Name, " ", "")
// 				//found_one := false
// 				for _, vvv := range val {
// 					cleandotname := strings.ReplaceAll(dot_in[vvv].Name, " ", "")

// 					if cleanmcname == cleandotname {

// 						var found OutPut

// 						found.Name = dot_in[vvv].Name
// 						found.Phone = dot_in[vvv].Phone
// 						found.Email = dot_in[vvv].Email
// 						found.Mc = mc_in[k].MC
// 						found.Dot = dot_in[vvv].Dot
// 						found.Street = dot_in[vvv].Street
// 						found.State = dot_in[vvv].State
// 						found.City = dot_in[vvv].City
// 						found.DotApplyDate = dot_in[vvv].ApplyDate
// 						found.McAppyDate = mc_in[k].ApplyDate
// 						found.McGrantDate = mc_in[k].GrantDate
// 						found.PowerUnit = dot_in[vvv].PowerUnit
// 						found.DriverTotal = dot_in[vvv].DriverTotal
// 						found.McsFileDate = dot_in[vvv].McsFileDate

// 						total++

// 						foundlist = append(foundlist, found)
// 						found_in_name = true
// 						//notfoundprev = false
// 						//found_one = true
// 						//fmt.Println("found one name ", v.Name, "  ", dot_in[vvv].Name)
// 						break
// 					}
// 				}

// 			}

// 		}

// 		if found_in_name || found_in_dba {
// 			notfoundprev = false
// 		}

// 	}

// 	SaveFound(found_location, foundlist)
// 	SaveNotFound(not_found_location, notfoundlist)
// 	fmt.Println(total)
// 	fmt.Println(len(mc_in))

// }

// func SearchDotName(key int, searchlist []Search) (int, bool) {

// 	ans := false
// 	position := -1

// 	low := 0
// 	high := len(searchlist) - 1

// 	for low <= high {
// 		median := (low + high) / 2

// 		if searchlist[median].Value < key {
// 			low = median + 1
// 		} else {
// 			high = median - 1
// 		}
// 	}

// 	if low == len(searchlist) || searchlist[low].Value != key {
// 		return position, ans
// 	}

// 	position = searchlist[low].Key

// 	//ans = true

// 	return position, true

// }

// type Search struct {
// 	Key   int
// 	Value int
// }

// func CreateInputDot(dot_path string) []Input_dot {

// 	out := make([]Input_dot, 0)

// 	f, err := os.Open(dot_path)
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
// 			var rec Input_dot

// 			if len(line) != 26 {
// 				fmt.Println("not 26")
// 			} else {
// 				//fmt.Println(line[0])
// 				//fmt.Println(line[2])
// 				rec.Dot = line[0]
// 				rec.Name = line[1]
// 				rec.DBA = line[2]
// 				rec.Street = line[6]
// 				rec.State = line[8]
// 				rec.City = line[7]
// 				rec.Phone = line[16]
// 				rec.Email = line[18]
// 				rec.ApplyDate = line[22]
// 				rec.McsFileDate = line[19]
// 				rec.PowerUnit = line[24]
// 				rec.DriverTotal = line[25]

// 			}

// 			out = append(out, rec)
// 		}
// 	}

// 	return out
// }

// func CreateInputMc(mc_path string) []Input_mc {

// 	out := make([]Input_mc, 0)
// 	f, err := os.Open(mc_path)
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
// 			var rec Input_mc

// 			//fmt.Println(len(line))

// 			if len(line) != 8 {
// 				fmt.Println("not 8")
// 			} else {
// 				//fmt.Println(line[0])
// 				//fmt.Println(line[2])

// 				rec.Name = line[3]
// 				rec.MC = line[1]
// 				rec.GrantDate = line[6]
// 				rec.City = line[5]
// 				rec.ApplyDate = line[7]

// 			}

// 			out = append(out, rec)
// 		}
// 	}

// 	return out
// }

// type Input_dot struct {
// 	Name        string
// 	DBA         string
// 	Street      string
// 	State       string
// 	City        string
// 	Phone       string
// 	Email       string
// 	Dot         string
// 	ApplyDate   string
// 	McsFileDate string
// 	PowerUnit   string
// 	DriverTotal string
// }

// type Input_mc struct {
// 	Name      string
// 	MC        string
// 	City      string
// 	GrantDate string
// 	ApplyDate string
// }
