package locdb

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// func (dbs *DBS) SaveResource() error {

// }

func (dbs *DBS) LoadResourceCsv(name string) ([]OutPut, error) {
	basedir, _ := os.Executable()

	path := filepath.Dir(basedir)

	filename := name + ".csv"

	new_resource := filepath.Join(path, "data", "csv", filename)
	f, err := os.Open(new_resource)
	if err != nil {
		log.Println("failed to open file ", err)
		return nil, err
	}

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Println("failed to marshall csv", err)
		return nil, err
	}

	out := make([]OutPut, 0)

	for i, line := range data {
		if i > 0 { // omit header line
			var rec OutPut

			if len(line) != 14 {
				fmt.Println(len(line))
				fmt.Println("not 13")
			} else {
				rec.Name = line[0]
				rec.Phone = line[1]
				rec.Email = strings.ToUpper(line[2])
				rec.Mc = line[3]
				rec.Dot = line[4]
				rec.Street = line[5]
				rec.State = line[6]
				rec.City = line[7]
				rec.DotApplyDate = line[8]
				rec.McAppyDate = line[9]
				rec.McGrantDate = line[10]
				rec.PowerUnit = line[11]
				rec.DriverTotal = line[12]
				rec.McsFileDate = line[13]
			}

			out = append(out, rec)
			fmt.Println(rec.Email)
		}
	}
	log.Println("len of new output, ", len(out))
	return out, nil
}

func (dbs *DBS) GetClientResourceList(emails []string) ([]ClientResource, error) {

	out := make([]ClientResource, 0)
	err := dbs.DB.Where("email IN ?", emails).Find(&out)

	if err.Error != nil {
		fmt.Println("failed to find emails")
		return nil, err.Error
	}

	return out, err.Error

}

func (dbs *DBS) GetAllResource() ([]ClientResource, error) {

	out := make([]ClientResource, 0)

	re := dbs.DB.Find(&out)

	if re.Error != nil {
		return nil, re.Error
	}

	return out, nil

}

func (dbs *DBS) LoadResourceMain(name string) error {

	all_in_db, err := dbs.GetAllResource()

	if err != nil {
		fmt.Println("failed to get all resource")
		return err
	}

	new_from_csv, err := dbs.LoadResourceCsv(name)

	if err != nil {
		fmt.Println("failed to load csv files")

		return err
	}

	search_db := dbs.CreateSearchDB(all_in_db)

	//search_csv := dbs.CreateSearchCsv(new_from_csv)

	searchmapdb := make(map[int][]int)
	last := 0

	for _, v := range search_db {

		val, ok := searchmapdb[v.Value]

		if !ok {
			mas := make([]int, 0)
			searchmapdb[v.Value] = mas
			searchmapdb[v.Value] = append(searchmapdb[v.Value], v.Key)
			last = v.Value

		} else {
			searchmapdb[v.Value] = append(val, v.Key)
		}
	}

	fmt.Println("searh name size , ", len(searchmapdb))
	fmt.Println("last map", searchmapdb[last])

	//	found_list := make([]OutPut, 0)
	not_found_list := make([]OutPut, 0)
	totalfound := 0
	totalnotfound := 0
	notfoundprev := false

	for k, v := range new_from_csv {
		foundinmap := false
		ascn := ToAsciiTotal(v.Email)
		//fmt.Println(k)

		if notfoundprev {
			not_found_list = append(not_found_list, new_from_csv[k-1])
			notfoundprev = false
			totalnotfound++
		} else {
			//notfoundprev = true

			val, ok := searchmapdb[ascn]

			if !ok {
				//emil does not exist

				//add to not found list

				//not_found_list = append(not_found_list, v)
				notfoundprev = true
				foundinmap = false

			} else {
				//email already exists to not add to not found list

				cleanemail := strings.ReplaceAll(v.Email, " ", "")
				for _, vv := range val {

					cleanemaildb := strings.ReplaceAll(all_in_db[vv].Email, " ", "")

					if cleanemail == cleanemaildb {
						foundinmap = true
						//var found OutPut
						//fmt.Println(found)
						log.Println(cleanemail, "  dirt-> ", cleanemaildb)
						totalfound++
						break
					}
				}

			}

			if !foundinmap {
				notfoundprev = true
			}

			if k < len(new_from_csv) {
				//notfoundprev = false
				fmt.Println("k ", k)
				fmt.Println("Email ", v.Email)
				fmt.Println("notfound ", totalnotfound)
				fmt.Println("found ", totalfound)
				fmt.Println("foundinmap  ", foundinmap)
				fmt.Println("notfoundprev ", notfoundprev)
			}

			continue

		}

	}

	log.Println("total found ", totalfound)
	log.Println("total not found, ", totalnotfound)

	err = dbs.SaveNotFound(not_found_list)

	return err

}

// func (dbs *DBS) GetDifferences(newoutputlist []OutPut) ([]OutPut, []OutPut) {

// 	new_output := make([]OutPut, 0)
// 	old_output := make([]OutPut, 0)

// }

func (dbs *DBS) SaveNotFound(inlist []OutPut) error {

	if len(inlist) == 0 {
		log.Println("saving none")
		return nil
	}
	clientresourcelist := make([]ClientResource, 0)

	for _, v := range inlist {
		resource := dbs.OutputToClientResource(v)
		clientresourcelist = append(clientresourcelist, resource)
	}

	err := dbs.DB.CreateInBatches(clientresourcelist, 100)

	if err.Error != nil {
		log.Println("failed to create in batches")
	}

	log.Println("saving ,  ", len(inlist))
	return err.Error
}

func (dbs *DBS) OutputToClientResource(in OutPut) ClientResource {

	out := NewClientResource()

	out.Name = in.Name
	out.Phone = in.Phone
	out.Email = in.Email
	out.Mc = in.Mc
	out.Dot = in.Dot
	out.Street = in.Street
	out.City = in.City
	out.State = in.State
	out.DotApplyDate = in.DotApplyDate
	out.McAppyDate = in.McAppyDate
	out.McGrantDate = in.McGrantDate
	out.PowerUnit = in.PowerUnit
	out.DriverTotal = in.DriverTotal
	out.McsFileDate = in.McsFileDate

	out.CreatedTime = time.Now()

	return *out
}

func (dbs *DBS) CreateSearchCsv(inlist []OutPut) []Search {
	out := make([]Search, 0)
	//outmap := make(map[int]int)
	for k, v := range inlist {
		//	out[k] = ToAsciiTotal(v.DBA)
		val := Search{Key: k, Value: ToAsciiTotal(v.Email)}
		out = append(out, val)
	}
	//sort.Ints(out)

	// sort.Slice(out, func(i, j int) bool {
	// 	return out[i].Value < out[j].Value
	// })
	return out
}

func (dbs *DBS) CreateSearchDB(inlist []ClientResource) []Search {
	out := make([]Search, 0)
	//outmap := make(map[int]int)
	for k, v := range inlist {
		//	out[k] = ToAsciiTotal(v.DBA)
		val := Search{Key: k, Value: ToAsciiTotal(v.Email)}
		out = append(out, val)
	}
	//sort.Ints(out)

	// sort.Slice(out, func(i, j int) bool {
	// 	return out[i].Value < out[j].Value
	// })
	return out
}

func ToAsciiTotal(str string) int {
	cleanstr := strings.ReplaceAll(str, " ", "")
	//runes := []rune(cleanstr)

	total := 0

	for _, v := range cleanstr {
		total = total + int(v)
	}
	return total
}

type Search struct {
	Key   int
	Value int
}
type OutPut struct {
	Name         string
	Phone        string
	Email        string
	Mc           string
	Dot          string
	Street       string
	State        string
	City         string //Phone string
	DotApplyDate string
	McAppyDate   string
	McGrantDate  string
	PowerUnit    string
	DriverTotal  string
	McsFileDate  string
}
