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

	if (json == nil) {
		return nil
	}

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

func getFiltersParams(inputFile string) (
    error,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string,
    []string) {

	err, dat := readAndUnmarshal(inputFile)
	return err,
    getFiltersExtension(dat),
    getFiltersGroupId(dat),
    getFiltersUpdatedAt(dat),
    getFiltersPhoneNumber(dat),
    getFiltersName(dat),
    getFiltersNumber(dat),
    getFiltersDirection(dat),
    getFiltersCalledNumber(dat),
    getFiltersType(dat),
    getFiltersCountryCode(dat),
    getFiltersCountry(dat),
    getFiltersNpa(dat),
    getFiltersNxx(dat),
    getFiltersXxxx(dat),
    getFiltersCity(dat),
    getFiltersProvince(dat),
    getFiltersPrice(dat),
    getFiltersCategory(dat),
    getFiltersFrom(dat),
    getFiltersTo(dat)
}

func getFiltersExtension(dat map[string]interface {}) []string {

	filterIdRaw := dat["filters[extension]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersGroupId(dat map[string]interface {}) []string {

	filterIdRaw := dat["filters[group_id]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersUpdatedAt(dat map[string]interface {}) []string {

	filterIdRaw := dat["filters[updated_at]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersPhoneNumber(dat map[string]interface {}) []string {

	filterIdRaw := dat["filters[phone_number]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersName(dat map[string]interface {}) []string {

	filterIdRaw := dat["filters[name]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersNumber(dat map[string]interface {}) []string {

	filterIdRaw := dat["filters[number]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersDirection(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[direction]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersCalledNumber(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[called_number]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersType(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[type]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersCountryCode(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[country_code]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersCountry(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[countries]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersNpa(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[npa]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersNxx(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[nxx]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersXxxx(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[xxxx]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersCity(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[city]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersProvince(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[province]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersPrice(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[price]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersCategory(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[category]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersFrom(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[from]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}

func getFiltersTo(dat map[string]interface {}) []string{

	filterIdRaw := dat["filters[to]"].([]interface{})

	str1 := filterIdRaw[0].(string)

	return []string{str1}
}


func getOtherParams(inputFile string) (
    error,
    string) {

  err, dat := readAndUnmarshal(inputFile)
  return err,
    getFieldString(dat["extension_id"])
}

func getSortParams(inputFile string) (
    error,
    string,
    string,
    string,
    string,
    string,
    string,
    string,
    string,
    string,
    string) {

  err, dat := readAndUnmarshal(inputFile)
  return err,
    getFieldString(dat["sort[id]"]),
    getFieldString(dat["sort[phone_number]"]),
    getFieldString(dat["sort[number]"]),
    getFieldString(dat["sort[name]"]),
    getFieldString(dat["sort[internal]"]),
    getFieldString(dat["sort[price]"]),
    getFieldString(dat["sort[extension]"]),
    getFieldString(dat["sort[updated_at]"]),
    getFieldString(dat["sort[created_at]"]),
    getFieldString(dat["sort[start_time]"])
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
