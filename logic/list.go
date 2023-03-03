package logic

import (
	"stocket/api"
)

func Get_Stock_List() api.List {
	list := api.Get_All_Stock_Code_Api()
	// fmt.Println(list.OTC_List)
	// fmt.Println(list.OTC_Total_Num)
	return list
}

func Get_Price_Test() {

	// fmt.Println(list.OTC_List[1])
	//api.Get_stocket_price(api.All_Listing_Stock[0])
}
