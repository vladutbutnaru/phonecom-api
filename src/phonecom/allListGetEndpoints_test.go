package main

import (
  "testing"
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

func TestListAvailablePhoneNumbers2(t *testing.T) {

  var err error

  phoneNumberSlice := make([]string, 0)
  phoneNumberSlice = append(phoneNumberSlice, "+12234687")
  countryCodeSlice := make([]string, 0)
  countryCodeSlice = append(countryCodeSlice, "1")
  npaSlice := make([]string, 0)
  npaSlice = append(npaSlice, "gt:10")
  nxxSlice := make([]string, 0)
  nxxSlice = append(nxxSlice, "831")
  xxxxSlice := make([]string, 0)
  xxxxSlice = append(xxxxSlice, "7863")
  citySlice := make([]string, 0)
  citySlice = append(citySlice, "Martinsville")
  provinceSlice := make([]string, 0)
  provinceSlice = append(provinceSlice, "IN")
  countrySlice := make([]string, 0)
  countrySlice = append(countrySlice, "US")
  priceSlice := make([]string, 0)
  priceSlice = append(priceSlice, "0")
  categorySlice := make([]string, 0)
  categorySlice = append(categorySlice, "1")


  err, _ = createCliListAvailablePhoneNumbers(listAvailablePhoneNumbers, phoneNumberSlice, countryCodeSlice, npaSlice, nxxSlice, xxxxSlice, citySlice, provinceSlice, countrySlice, priceSlice, categorySlice)
  assertErrorNotNull(t, err)
}

func TestListAvailablePhoneNumberRegions2(t *testing.T) {

  var err error

  countryCodeSlice := make([]string, 0)
  countryCodeSlice = append(countryCodeSlice, "1")
  npaSlice := make([]string, 0)
  npaSlice = append(npaSlice, "not-empty")
  nxxSlice := make([]string, 0)
  nxxSlice = append(nxxSlice, "249")
  isTollFreeSlice := make([]string, 0)
  isTollFreeSlice = append(isTollFreeSlice, "true")
  citySlice := make([]string, 0)
  citySlice = append(citySlice, "Prilep")
  provincePostalCodeSlice := make([]string, 0)
  provincePostalCodeSlice = append(provincePostalCodeSlice, "2345")
  countryPostalCodeSlice := make([]string, 0)
  countryPostalCodeSlice = append(countryPostalCodeSlice, "2345")

  err, _ = createCliListAvailablePhoneNumberRegions(listAvailablePhoneNumberRegions, countryCodeSlice, npaSlice, nxxSlice, isTollFreeSlice, citySlice, provincePostalCodeSlice, countryPostalCodeSlice)
  assertErrorNotNull(t, err)
}
