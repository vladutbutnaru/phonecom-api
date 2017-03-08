package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestListAccounts(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listAccounts)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getAccount, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListApplications(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listApplications)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getApplication, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListCallLogs(t *testing.T) {

  var err error

  err, _ = createCli(listCallLogs)
  assertErrorNotNull(t, err)
}

func TestListDevices(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listDevices)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getDevice, int(firstId))

  assertErrorNotNull(t, err)
}

func TestExpressServiceCodes(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listExpressServiceCodes)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getExpressServiceCode, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListExtensions(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listExtensions)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getExtension, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListCallerIds(t *testing.T) {

  var err error
	var json map[string] interface{}

	err, json = createCli(listExtensions)
	assertErrorNotNull(t, err)
	extensionId := getFirstId(json)

  err, _ = createCliWithId(getCallerId, extensionId)
  assertErrorNotNull(t, err)
}

func TestListContacts(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listExtensions)
  assertErrorNotNull(t, err)
  extensionId := getFirstId(json)

  err, json = createCliWithId(listContacts, extensionId)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithIdIdSecondary(getContact, extensionId, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListGroups(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listExtensions)
  assertErrorNotNull(t, err)
  extensionId := getFirstId(json)

  err, json = createCliWithId(listGroups, extensionId)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithIdIdSecondary(getGroup, extensionId, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListMedia(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listMedia)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getRecording, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListMenus(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listMenus)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getMenu, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListPhoneNumbers(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listPhoneNumbers)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getPhoneNumber, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListQueues(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listQueues)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getQueue, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListRoutes(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listRoutes)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getRoute, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListSchedules(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listSchedules)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getSchedule, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListSms(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listSms)
  assertErrorNotNull(t, err)

  firstId := getFirstIdString(json)
  err, json = createGetCliStringId(getSms, firstId)

  assertErrorNotNull(t, err)
}

func TestListSubaccounts(t *testing.T) {

  var err error

  err, _ = createCli(listSubaccounts)
  assertErrorNotNull(t, err)
}

func TestListTrunks(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listTrunks)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createCliWithId(getTrunk, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListAvailablePhoneNumbers(t *testing.T) {

  var err error

  err, _ = createCli(listAvailablePhoneNumbers)
  assertErrorNotNull(t, err)
}

func TestListAvailablePhoneNumberRegions(t *testing.T) {

  var err error

  err, _ = createCli(listAvailablePhoneNumberRegions)
  assertErrorNotNull(t, err)
}

func TestFilterSortListAvailablePhoneNumbers(t *testing.T) {

  phoneNumberSlice := make([]string, 0)
  expectedPhoneNumber := "+12234687"
  phoneNumberSlice = append(phoneNumberSlice, expectedPhoneNumber)
  countryCodeSlice := make([]string, 0)
  expectedCountryCode := "1"
  countryCodeSlice = append(countryCodeSlice, expectedCountryCode)
  npaSlice := make([]string, 0)
  expectedNpa := "gt:10"
  npaSlice = append(npaSlice, expectedNpa)
  nxxSlice := make([]string, 0)
  expectedNxx := "831"
  nxxSlice = append(nxxSlice, expectedNxx)
  xxxxSlice := make([]string, 0)
  expectedXxxx := "7863"
  xxxxSlice = append(xxxxSlice, expectedXxxx)
  citySlice := make([]string, 0)
  expectedCity := "Martinsville"
  citySlice = append(citySlice, expectedCity)
  provinceSlice := make([]string, 0)
  expectedProvince := "IN"
  provinceSlice = append(provinceSlice, expectedProvince)
  countrySlice := make([]string, 0)
  expectedCountry := "US"
  countrySlice = append(countrySlice, expectedCountry)
  priceSlice := make([]string, 0)
  expectedPrice := "0"
  priceSlice = append(priceSlice, expectedPrice)
  categorySlice := make([]string, 0) // check filter
  expectedCategory := "10"
  categorySlice = append(categorySlice, expectedCategory)

  err, response := createCliListAvailablePhoneNumbers(listAvailablePhoneNumbers, phoneNumberSlice, countryCodeSlice, npaSlice, nxxSlice, xxxxSlice, citySlice, provinceSlice, countrySlice, priceSlice, categorySlice)
  assertErrorNotNull(t, err)

  filters := getFilters(response)
  phoneNumber := filters["phone_number"].(string)
  assert.Equal(t, expectedPhoneNumber, phoneNumber)
  countryCode := filters["country_code"].(string)
  assert.Equal(t, expectedCountryCode, countryCode)
  npa := filters["npa"].(string)
  assert.Equal(t, expectedNpa, npa)
  nxx := filters["nxx"].(string)
  assert.Equal(t, expectedNxx, nxx)
  xxxx := filters["xxxx"].(string)
  assert.Equal(t, expectedXxxx, xxxx)
  city := filters["city"].(string)
  assert.Equal(t, expectedCity, city)
  province := filters["province"].(string)
  assert.Equal(t, expectedProvince, province)
  country := filters["country"].(string)
  assert.Equal(t, expectedCountry, country)
  price := filters["price"].(string)
  assert.Equal(t, expectedPrice, price)
  //category := filters["category"].(string)
  //assert.Equal(t, expectedCategory, category)
}

func TestFilterSortListAvailablePhoneNumberRegions(t *testing.T) {

  countryCodeSlice := make([]string, 0)
  expectedCountryCode := "1"
  countryCodeSlice = append(countryCodeSlice, expectedCountryCode)
  npaSlice := make([]string, 0)
  expectedNpa := "gt:10"
  npaSlice = append(npaSlice, expectedNpa)
  nxxSlice := make([]string, 0)
  expectedNxx := "249"
  nxxSlice = append(nxxSlice, expectedNxx)
  isTollFreeSlice := make([]string, 0)
  expectedIsTollFree := "true"
  isTollFreeSlice = append(isTollFreeSlice, expectedIsTollFree)
  citySlice := make([]string, 0)
  expectedCity := "Prilep"
  citySlice = append(citySlice, expectedCity)
  provincePostalCodeSlice := make([]string, 0)
  expectedProvincePostalCode := "2345"
  provincePostalCodeSlice = append(provincePostalCodeSlice, expectedProvincePostalCode)
  countryPostalCodeSlice := make([]string, 0)
  expectedCountryPostalCode := "23453"
  countryPostalCodeSlice = append(countryPostalCodeSlice, expectedCountryPostalCode)

  err, response := createCliListAvailablePhoneNumberRegions(listAvailablePhoneNumberRegions, countryCodeSlice, npaSlice, nxxSlice, isTollFreeSlice, citySlice, provincePostalCodeSlice, countryPostalCodeSlice)
  assertErrorNotNull(t, err)

  filters := getFilters(response)
  countryCode := filters["country_code"].(string)
  assert.Equal(t, expectedCountryCode, countryCode)
  npa := filters["npa"].(string)
  assert.Equal(t, expectedNpa, npa)
  nxx := filters["nxx"].(string)
  assert.Equal(t, expectedNxx, nxx)
  isTollFree := filters["is_toll_free"].(string)
  assert.Equal(t, expectedIsTollFree, isTollFree)
  city := filters["city"].(string)
  assert.Equal(t, expectedCity, city)
  provincePostalCode := filters["province_postal_code"].(string)
  assert.Equal(t, expectedProvincePostalCode, provincePostalCode)
  countryPostalCode := filters["country_postal_code"].(string)
  assert.Equal(t, expectedCountryPostalCode, countryPostalCode)
}
