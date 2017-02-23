package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"phonecom-go-sdk"
	"errors"
)

func getListParams(inputFile string) (error, int32, int32, int32, string, []string) {

	err, data := readAndUnmarshal(inputFile)
	return err,
		getField(data["account_id"]),
		getField(data["limit"]),
		getField(data["offset"]),
		getFieldString(data["fields"]),
		getFilterId(data);
}

func getField(field interface {}) int32 {

	if (field == nil) {
		return -1
	}

	return int32(field.(float64))
}

func getFilterId(json map[string]interface {}) []string {

	filterIdRaw := json["filters[id]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func readAndUnmarshal(inputFile string) (error, map[string]interface {}) {

	file, e := ioutil.ReadFile(inputFile)

	if (e != nil) {
		return errors.New("Could not read input file: " + inputFile), nil
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(file, &dat); err != nil {
		return errors.New("Could not unmarshal input file: " + inputFile), nil
	}

	return nil, dat;
}

func getFieldString(field interface {}) string {

	if (field == nil) {
		return ""
	}

	return field.(string)
}

func getExtensionIdFromFile(inputFile string) (error, string) {

	err, dat := readAndUnmarshal(inputFile)
	return err, getFieldString(dat["extension_id"])
}



func getFiltersExtensionFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[extension]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersGroupIdFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[group_id]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersUpdatedAtFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[updated_at]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}

}

func getFiltersPhoneNumberFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[phone_number]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}

}

func getFiltersNameFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[name]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}

}

func getFiltersNumberFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[number]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersDirectionFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[direction]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersCalledNumberFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[called_number]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersTypeFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[type]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersCountryCodeFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[country_code]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersCountryFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[countries]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersNpaFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[npa]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersNxxFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[nxx]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersXxxxFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[xxxx]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersCityFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[city]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersProvinceFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[province]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersPriceFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[price]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersCategoryFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[category]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersFromFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[from]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getFiltersToFromFile(inputFile string) (error, []string) {

	err, dat := readAndUnmarshal(inputFile)

	filterIdRaw := dat["filters[to]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return err, []string{str1}
}

func getSortIdFromFile(inputFile string) (error, string) {

	err, dat := readAndUnmarshal(inputFile)

	return err, getFieldString(dat["sort[id]"])
}

func getSortPhoneNumberFromFile(inputFile string) (error, string) {

	err, dat := readAndUnmarshal(inputFile)

	return err, getFieldString(dat["sort[phone_number]"])
}

func getSortNumberFromFile(inputFile string) (error, string) {

	err, dat := readAndUnmarshal(inputFile)

	return err, getFieldString(dat["sort[number]"])
}

func getSortNameFromFile(inputFile string) (error, string) {

	err, dat := readAndUnmarshal(inputFile)

	return err, getFieldString(dat["sort[name]"])

}

func getSortInternalFromFile(inputFile string) (error, string) {

	err, dat := readAndUnmarshal(inputFile)

	return err, getFieldString(dat["sort[internal]"])

}
func getSortPriceFromFile(inputFile string) (error, string) {

	err, dat := readAndUnmarshal(inputFile)

	return err, getFieldString(dat["sort[price]"])
}

func getSortExtensionFromFile(inputFile string) (error, string) {

	err, dat := readAndUnmarshal(inputFile)

	return err, getFieldString(dat["sort[extension]"])
}

func getSortUpdatedAtFromFile(inputFile string) (error, string) {

	err, dat := readAndUnmarshal(inputFile)

	return err, getFieldString(dat["sort[updated_at]"])
}

func getSortCreatedAtFromFile(inputFile string) (error, string) {

	err, dat := readAndUnmarshal(inputFile)
	return err, getFieldString(dat["sort[created_at]"])
}

func getSortStartTimeFromFile(inputFile string) (error, string) {

	err, dat := readAndUnmarshal(inputFile)

	return err, getFieldString(dat["sort[start_time]"])
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
