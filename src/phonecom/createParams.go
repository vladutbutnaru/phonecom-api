package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
    "phonecom-go-sdk"
)

//func createCallParams(inputFile string) swagger.CreateCallParams {
//
///*
//	var params swagger.CreateCallParams
//
//	params.CallerPhoneNumber = "+17324568888"
//	params.CallerExtension = 1749841
//	params.CallerCallerId = "+17329980138"
//	params.CallerPrivate = true
//	params.CalleePhoneNumber = "+17329980138"
//	params.CalleeExtension = 1749842
//	params.CalleeCallerId = "+17324568888"
//	params.CalleePrivate = true
//*/
//
//
//     file, e := ioutil.ReadFile(inputFile)
//    if e != nil {
//        fmt.Printf("File error: %v\n", e)
//        os.Exit(1)
//    }
//    fmt.Printf("%s\n", string(file))
//
//    var params  swagger.CreateCallParams
//    json.Unmarshal(file, &params)
//
//	return params
//}


func getAccountIdFromFile(inputFile string) int32 {
    var params int32
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = int32(dat["account_id"].(float64))
    
    return params
    
    
}

func getLimitFromFile(inputFile string) int32 {
      var params int32
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = int32(dat["limit"].(float64))
    
    return params
    
    
}

func getOffsetFromFile(inputFile string) int32 {
      var params int32
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = int32(dat["offset"].(float64))
    
    return params
    
    
}

func getFieldsFromFile(inputFile string) string {
      var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["fields"].(string)
    
    return params
    
    
}

func getExtensionIdFromFile(inputFile string) string {
      var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["extension_id"].(string)
    
    return params
    
    
}


func getFiltersIdFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[id]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}


func getFiltersExtensionFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[extension]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}


func getFiltersGroupIdFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[group_id]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersUpdatedAtFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[updated_at]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersPhoneNumberFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[phone_number]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}


func getFiltersNameFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[name]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersNumberFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[number]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}


func getFiltersDirectionFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[direction]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersCalledNumberFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[called_number]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersTypeFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[type]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersCountryCodeFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[country_code]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersCountryFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[countries]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}
func getFiltersNpaFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[npa]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersNxxFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[nxx]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersXxxxFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[xxxx]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersCityFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[city]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersProvinceFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[province]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersPriceFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[price]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersCategoryFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[category]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersFromFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[from]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getFiltersToFromFile(inputFile string) []string {
      //var params []string
    var dat map[string]interface{}

     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }



    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }

	  filterIdRaw := dat["filters[to]"].([]interface{})

	  str1 := filterIdRaw[0].(string)

    
    return []string{str1}
    
    
}

func getSortIdFromFile(inputFile string) string {
      var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["sort[id]"].(string)
    
    return params
    
    
}

func getSortPhoneNumberFromFile(inputFile string) string {
      var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["sort[phone_number]"].(string)
    
    return params
    
    
}

func getSortNumberFromFile(inputFile string) string {
      var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["sort[number]"].(string)
    
    return params
    
    
}

func getSortNameFromFile(inputFile string) string {
        var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["sort[name]"].(string)
    
    return params
    
    
}

func getSortInternalFromFile(inputFile string) string {
        var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["sort[internal]"].(string)
    
    return params
    
    
}
func getSortPriceFromFile(inputFile string) string {
        var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["sort[price]"].(string)
    
    return params
    
    
}

func getSortExtensionFromFile(inputFile string) string {
        var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["sort[extension]"].(string)
    
    return params
    
    
}

func getSortUpdatedAtFromFile(inputFile string) string {
        var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["sort[updated_at]"].(string)
    
    return params
    
    
    
    
}

func getSortCreatedAtFromFile(inputFile string) string {
        var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["sort[created_at]"].(string)
    
    return params
    
    
    
    
}


func getSortStartTimeFromFile(inputFile string) string {
        var params string
    var dat map[string]interface{}
    
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        return ""
    }
    

    
    if err := json.Unmarshal(file, &dat); err != nil {
        panic(err)
    }
    
    params  = dat["sort[start_time]"].(string)
    
    return params
    
    
}

func createDeviceParams(inputFile string) swagger.CreateDeviceParams {

	

	
     file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateDeviceParams
    json.Unmarshal(file, &params)

	return params
}
func createMenuParams(inputFile string) swagger.CreateMenuParams {
    
    file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateMenuParams
    json.Unmarshal(file, &params)

	return params


}

func replaceMenuParams(inputFile string) swagger.ReplaceMenuParams{
    
	 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.ReplaceMenuParams
    json.Unmarshal(file, &params)

	return params





}
func createExtensionsParams(inputFile string) swagger.CreateExtensionParams {

	 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateExtensionParams
    json.Unmarshal(file, &params)

	return params
}
func createQueueParams(inputFile string) swagger.CreateQueueParams {
	 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateQueueParams
    json.Unmarshal(file, &params)

	return params

}

func createRouteParams(inputFile string) swagger.CreateRouteParams {
	 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateRouteParams
    json.Unmarshal(file, &params)

	return params


}

func createSmsParams(inputFile string) swagger.CreateSmsParams {
		 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateSmsParams
    json.Unmarshal(file, &params)

	return params
}

func createSubaccountParams(inputFile string) swagger.CreateSubaccountParams{
		 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateSubaccountParams
    json.Unmarshal(file, &params)

	return params


}

func createTrunkParams(inputFile string) swagger.CreateTrunkParams{
		 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateTrunkParams
    json.Unmarshal(file, &params)

	return params


}


func replaceExtensionParams(inputFile string) swagger.ReplaceExtensionParams {
	 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.ReplaceExtensionParams
    json.Unmarshal(file, &params)

	return params
}

func createContactParams(inputFile string) swagger.CreateContactParams {

 file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateContactParams
    json.Unmarshal(file, &params)

	return params
}

func createGroupParams(inputFile string) swagger.CreateGroupParams {

	file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreateGroupParams
    json.Unmarshal(file, &params)

	return params
}

func createPhoneNumberParams(inputFile string) swagger.CreatePhoneNumberParams {

	file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.CreatePhoneNumberParams
    json.Unmarshal(file, &params)

	return params
}

func replacePhoneNumberParams(inputFile string) swagger.ReplacePhoneNumberParams {

	file, e := ioutil.ReadFile(inputFile)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    
    var params  swagger.ReplacePhoneNumberParams
    json.Unmarshal(file, &params)

	return params
}
