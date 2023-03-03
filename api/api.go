package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type List struct {
	OTC_List      []string `json:"OTCList"`
	OTC_Total_Num int      `json:"OTClNum"`
}

type Result struct {
	Date_List      []string
	Listing_Result map[string][]Detail
}

type Detail struct {
	Close float32
	MA5   float32
	MA10  float32
	MA20  float32
	MA60  float32
}

func Get_All_Stock_Code_Api() List {

	url := "https://www.tpex.org.tw/openapi/v1/mopsfin_t187ap03_O"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("date", " Wed,15 Feb 2023 04:06:39 GMT")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res_json []map[string]interface{}
	json.Unmarshal(result, &res_json)
	//fmt.Printf("%s", result)

	var list List

	for _, r := range res_json {

		list.OTC_List = append(list.OTC_List, r["SecuritiesCompanyCode"].(string))

		// fmt.Println(i, r["SecuritiesCompanyCode"])
		// fmt.Println()
		// fmt.Println()

	}
	list.OTC_Total_Num = len(res_json)

	return list
}

func (r *Result) Get_stocket_price(code string) {
	url := fmt.Sprintf("https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=%s&stockNo=%s", "20230101", code)

	req, err := http.NewRequest("GET", url, nil)
	//req.Header.Add("If-Modified-Since", "Mon, 26 Jul 1997 05:00:00 GMT")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res_json map[string]interface{}
	json.Unmarshal(result, &res_json)
	//fmt.Printf("%s", result)
	tmp := res_json["data"].([]interface{})
	for _, r := range tmp {

		k := r.([]interface{})

		fmt.Printf("%+v:%+v\n", k[0], k[6])

	}
	fmt.Println()

}
