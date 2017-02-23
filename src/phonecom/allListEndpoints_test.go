package main

import (
  "testing"
)

func TestListAccounts(t *testing.T) {

  var err = createCli(listAccounts)

  assertErrorNotNull(t, err)
}

func TestListApplications(t *testing.T) {

  var err = createCli(listApplications)

  assertErrorNotNull(t, err)
}

func TestListCallLogs(t *testing.T) {

  var err = createCli(listCallLogs)

  assertErrorNotNull(t, err)
}

func TestListDevices(t *testing.T) {

  var err = createCli(listDevices)

  assertErrorNotNull(t, err)
}

func TestExpressServiceCodes(t *testing.T) {

  var err = createCli(listExpressServiceCodes)

  assertErrorNotNull(t, err)
}

func TestListExtensions(t *testing.T) {

  var err = createCli(listExtensions)

  assertErrorNotNull(t, err)
}

func TestListCallerIds(t *testing.T) {

  var err = createCli(getCallerId)

  assertErrorNotNull(t, err)
}

func TestListContacts(t *testing.T) {

  var err = createCli(listContacts)

  assertErrorNotNull(t, err)
}

func TestListGroups(t *testing.T) {

  var err = createCli(listGroups)

  assertErrorNotNull(t, err)
}

func TestListMedia(t *testing.T) {

  var err = createCli(listMedia)

  assertErrorNotNull(t, err)
}

func TestListMenus(t *testing.T) {

  var err = createCli(listMenus)

  assertErrorNotNull(t, err)
}

func TestListPhoneNumbers(t *testing.T) {

  var err = createCli(listPhoneNumbers)

  assertErrorNotNull(t, err)
}

func TestListQueues(t *testing.T) {

  var err = createCli(listQueues)

  assertErrorNotNull(t, err)
}

func TestListRoutes(t *testing.T) {

  var err = createCli(listRoutes)

  assertErrorNotNull(t, err)
}

func TestListSchedules(t *testing.T) {

  var err = createCli(listSchedules)

  assertErrorNotNull(t, err)
}

func TestListSms(t *testing.T) {

  var err = createCli(listSms)

  assertErrorNotNull(t, err)
}

func TestListSubaccounts(t *testing.T) {

  var err = createCli(listSubaccounts)

  assertErrorNotNull(t, err)
}

func TestListTrunks(t *testing.T) {

  var err = createCli(listTrunks)

  assertErrorNotNull(t, err)
}

func TestListAvailablePhoneNumbers(t *testing.T) {

  var err = createCli(listAvailablePhoneNumbers)

  assertErrorNotNull(t, err)
}

func TestListAvailablePhoneNumberRegions(t *testing.T) {

  var err = createCli(listAvailablePhoneNumberRegions)

  assertErrorNotNull(t, err)
}
