package main

import (
	"fmt"
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

	if (json == nil || json["filters[id]"] == nil) {
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

  if (dat == nil) {
    return err, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil
  }

	return err,
    createStringArray(dat["filters[extension]"]),
    createStringArray(dat["filters[group_id]"]),
    createStringArray(dat["filters[updated_at]"]),
    createStringArray(dat["filters[phone_number]"]),
    createStringArray(dat["filters[name]"]),
    createStringArray(dat["filters[number]"]),
    createStringArray(dat["filters[direction]"]),
    createStringArray(dat["filters[called_number]"]),
    createStringArray(dat["filters[type]"]),
    createStringArray(dat["filters[country_code]"]),
    createStringArray(dat["filters[countries]"]),
    createStringArray(dat["filters[npa]"]),
    createStringArray(dat["filters[nxx]"]),
    createStringArray(dat["filters[xxxx]"]),
    createStringArray(dat["filters[city]"]),
    createStringArray(dat["filters[province]"]),
    createStringArray(dat["filters[price]"]),
    createStringArray(dat["filters[category]"]),
    createStringArray(dat["filters[from]"]),
    createStringArray(dat["filters[to]"])
}

func createStringArray(filter interface{}) []string {

  if (filter == nil) {
    return nil
  }

  filterArray := filter.([] interface{})
  str1 := filterArray[0].(string)

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
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func readAndUnmarshalFile(inputFile string, params interface{}) error {

	file, e := ioutil.ReadFile(inputFile)

	if e != nil {
		return e
	}

	fmt.Printf("%s\n", string(file))

  json.Unmarshal(file, params)

	return nil
}

func createMenuParams(inputFile string) swagger.CreateMenuParams {

	var params swagger.CreateMenuParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func replaceMenuParams(inputFile string) swagger.ReplaceMenuParams {

	var params swagger.ReplaceMenuParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createExtensionsParams(inputFile string) swagger.CreateExtensionParams {

	var params swagger.CreateExtensionParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createQueueParams(inputFile string) swagger.CreateQueueParams {

	var params swagger.CreateQueueParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createRouteParams(inputFile string) swagger.CreateRouteParams {

	var params swagger.CreateRouteParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createSmsParams(inputFile string) swagger.CreateSmsParams {

	var params swagger.CreateSmsParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createSubaccountParams(inputFile string) swagger.CreateSubaccountParams {

	var params swagger.CreateSubaccountParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createTrunkParams(inputFile string) swagger.CreateTrunkParams {

	var params swagger.CreateTrunkParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func replaceExtensionParams(inputFile string) swagger.ReplaceExtensionParams {

	var params swagger.ReplaceExtensionParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createContactParams(inputFile string) swagger.CreateContactParams {

	var params swagger.CreateContactParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createGroupParams(inputFile string) swagger.CreateGroupParams {

	var params swagger.CreateGroupParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createPhoneNumberParams(inputFile string) swagger.CreatePhoneNumberParams {

	var params swagger.CreatePhoneNumberParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func replacePhoneNumberParams(inputFile string) swagger.ReplacePhoneNumberParams {

	var params swagger.ReplacePhoneNumberParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}
