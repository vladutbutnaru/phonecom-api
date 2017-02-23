package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"phonecom-go-sdk"
)

func getAccountIdFromFile(inputFile string) int32 {

	dat := readAndUnmarshal(inputFile)
	return int32(dat["account_id"].(float64))
}

func readAndUnmarshal(inputFile string) map[string]interface {} {

	file, e := ioutil.ReadFile(inputFile)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(file, &dat); err != nil {
		panic(err)
	}

	return dat;
}

func getLimitFromFile(inputFile string) int32 {

	dat := readAndUnmarshal(inputFile)
	return int32(dat["limit"].(float64))
}

func getOffsetFromFile(inputFile string) int32 {

	dat := readAndUnmarshal(inputFile)
	return int32(dat["offset"].(float64))
}

func getFieldsFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)
	return dat["fields"].(string)
}

func getExtensionIdFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)
	return dat["extension_id"].(string)
}

func getFiltersIdFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[id]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersExtensionFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[extension]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersGroupIdFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[group_id]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersUpdatedAtFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[updated_at]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}

}

func getFiltersPhoneNumberFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[phone_number]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}

}

func getFiltersNameFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[name]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}

}

func getFiltersNumberFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[number]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersDirectionFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[direction]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersCalledNumberFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[called_number]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersTypeFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[type]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersCountryCodeFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[country_code]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersCountryFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[countries]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersNpaFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[npa]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersNxxFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[nxx]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersXxxxFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[xxxx]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersCityFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[city]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersProvinceFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[province]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersPriceFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[price]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersCategoryFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[category]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersFromFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[from]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersToFromFile(inputFile string) []string {

	dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[to]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getSortIdFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)

	return dat["sort[id]"].(string)
}

func getSortPhoneNumberFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)

	return dat["sort[phone_number]"].(string)
}

func getSortNumberFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)

	return dat["sort[number]"].(string)
}

func getSortNameFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)

	return dat["sort[name]"].(string)

}

func getSortInternalFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)

	return dat["sort[internal]"].(string)

}
func getSortPriceFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)

	return dat["sort[price]"].(string)
}

func getSortExtensionFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)

	return dat["sort[extension]"].(string)
}

func getSortUpdatedAtFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)

	return dat["sort[updated_at]"].(string)
}

func getSortCreatedAtFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)
	return dat["sort[created_at]"].(string)
}

func getSortStartTimeFromFile(inputFile string) string {

	dat := readAndUnmarshal(inputFile)

	return dat["sort[start_time]"].(string)
}

func createDeviceParams(inputFile string) swagger.CreateDeviceParams {

	var params swagger.CreateDeviceParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func readFile(inputFile string) []byte {

	file, e := ioutil.ReadFile(inputFile)

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	fmt.Printf("%s\n", string(file))

	return file
}

func createMenuParams(inputFile string) swagger.CreateMenuParams {

	var params swagger.CreateMenuParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func replaceMenuParams(inputFile string) swagger.ReplaceMenuParams {

	var params swagger.ReplaceMenuParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func createExtensionsParams(inputFile string) swagger.CreateExtensionParams {

	var params swagger.CreateExtensionParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func createQueueParams(inputFile string) swagger.CreateQueueParams {

	var params swagger.CreateQueueParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func createRouteParams(inputFile string) swagger.CreateRouteParams {

	var params swagger.CreateRouteParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func createSmsParams(inputFile string) swagger.CreateSmsParams {

	var params swagger.CreateSmsParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func createSubaccountParams(inputFile string) swagger.CreateSubaccountParams {

	var params swagger.CreateSubaccountParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func createTrunkParams(inputFile string) swagger.CreateTrunkParams {

	var params swagger.CreateTrunkParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func replaceExtensionParams(inputFile string) swagger.ReplaceExtensionParams {

	var params swagger.ReplaceExtensionParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func createContactParams(inputFile string) swagger.CreateContactParams {

	var params swagger.CreateContactParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func createGroupParams(inputFile string) swagger.CreateGroupParams {

	var params swagger.CreateGroupParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func createPhoneNumberParams(inputFile string) swagger.CreatePhoneNumberParams {

	var params swagger.CreatePhoneNumberParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}

func replacePhoneNumberParams(inputFile string) swagger.ReplacePhoneNumberParams {

	var params swagger.ReplacePhoneNumberParams
	json.Unmarshal(readFile(inputFile), &params)

	return params
}
