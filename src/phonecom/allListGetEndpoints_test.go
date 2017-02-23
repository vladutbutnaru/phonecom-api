package main

import (
  "testing"
)

func TestListAccounts(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listAccounts)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func getFirstId(json map[string] interface{}) int {

  items := json["items"]
  _ = items

  return 0
}

func TestListApplications(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listApplications)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListCallLogs(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listCallLogs)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListDevices(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listDevices)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestExpressServiceCodes(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listExpressServiceCodes)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListExtensions(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listExtensions)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListCallerIds(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(getCallerId)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListContacts(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listContacts)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListGroups(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listGroups)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListMedia(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listMedia)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListMenus(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listMenus)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListPhoneNumbers(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listPhoneNumbers)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListQueues(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listQueues)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListRoutes(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listRoutes)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListSchedules(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listSchedules)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListSms(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listSms)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListSubaccounts(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listSubaccounts)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListTrunks(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listTrunks)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListAvailablePhoneNumbers(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listAvailablePhoneNumbers)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}

func TestListAvailablePhoneNumberRegions(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listAvailablePhoneNumberRegions)

  firstId := getFirstId(json)
  _ = firstId

  assertErrorNotNull(t, err)
}
