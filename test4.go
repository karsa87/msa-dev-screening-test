package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// func createCsv(students []Student) {
// 	// data, err := json.Marshal(students)
// 	b, err := json.Marshal(students)
// 	data := string(b)

// 	// create a file
// 	file, err := os.Create("result.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	// initialize csv writer
// 	writer := csv.NewWriter(file)

// 	defer writer.Flush()

// 	// write all rows at once
// 	writer.WriteAll(data)

// 	// write single row
// 	extraData := []string{"lettuce", "raspberry"}
// 	writer.Write(extraData)
// }

type Student struct {
	VaultId                string `json:"vault_id"`
	Year                   string `json:"year"`
	School                 string `json:"school"`
	University             string `json:"university"`
	EmploymentRateOverall  string `json:"employment_rate_overall"`
	EmploymentRateFtPerm   string `json:"employment_rate_ft_perm"`
	BasicMonthlyMean       string `json:"basic_monthly_mean"`
	BasicMonthlyMedian     string `json:"basic_monthly_median"`
	GrossMonthlyMean       string `json:"gross_monthly_mean"`
	GrossMonthlyMedian     string `json:"gross_monthly_median"`
	GrossMthly25Percentile string `json:"gross_mthly_25_percentile"`
	GrossMthly75Percentile string `json:"gross_mthly_75_percentile"`
}

type Response struct {
	Code int `json:"code"`
	Data struct {
		Rows []Student `json:"rows"`
	} `json:"data"`
}

func main() {
	res, err := http.Get("https://api-production.data.gov.sg/v2/internal/api/datasets/d_3c55210de27fcccda2ed0c63fdd2b352/rows?limit=2")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var result Response
	if err := json.Unmarshal(resBody, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println(err)
		fmt.Println("Can not unmarshal JSON")
	}

	students := make(map[int][]Student)
	for i := 0; i < len(result.Data.Rows); i++ {
		student := result.Data.Rows[i]
		year, _ := strconv.Atoi(string(student.Year))
		students[year] = append(students[year], student)
	}

	data, err := json.Marshal(students[2013])
	fmt.Println(string(data))
}
