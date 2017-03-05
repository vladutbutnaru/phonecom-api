package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "phonecom-go-sdk"
  "errors"
)

var param CliParams

func main() {

  app := cli.NewApp()

  app.Flags = getCliFlags()
	var cliConfig CliConfig

  app.Action = func(c *cli.Context) error {
    var err error
    err, _ = execute(c, cliConfig)
    return err
  }

  app.Run(os.Args)
}

func execute(
    c *cli.Context,
    cliConfig CliConfig) (error, map[string] interface{}) {

  var err error
	param, err = createCliParams(c);

	if (err != nil) {
		return err, nil
	}

  showDryRunVerbose(param)
  if (param.dryRun) {
    return nil, nil
  }

	cliConfig, err = cliConfig.createOrReadCliConfig()

	if (err != nil) {
    return err, nil
	}

  if (param.verbose) {
    fmt.Printf("CLI configuration: %+v\n\n", cliConfig)
  }

  var sampleJsonCreator SampleJsonCreator
  sampleJsonCreator.createSampleInOutIfNeeded(param.inputFormat)

  if param.samplein != "" || param.sampleout != "" {
    return nil, nil
  }

	swaggerConfig := cliConfig.createSwaggerConfig(cliConfig)
  apiResolver := ApiResolver{swaggerConfig, param.command}
  api := apiResolver.resolve()

	param.accountId = cliConfig.AccountId

  if (api == nil) {
    return errors.New(msgCouldNotGetResponse), nil
  }

	return invokeCommand(api)
}

func invokeCommand(api interface{}) (error, map[string] interface{}) {

  var command = param.command
  var accountId = param.accountId
  var id = param.id
  var filtersId = param.filtersId
  var filterParams = param.filterParams
  var sortParams = param.sortParams
  var limit = param.limit
  var offset = param.offset
  var fields = param.fields
  var input = param.input
  var idSecondary = param.idSecondary
  var idString = param.idString

  switch api := api.(type) {

  case swagger.MediaApi:

    if (param.otherParams.recordingId > 0) {
      id = param.otherParams.recordingId
    }

    switch (command) {

    case listMedia:

      return handle(api.ListAccountMedia(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getRecording:

      return handle(api.GetAccountMedia(accountId, id))
    }

  case swagger.MenusApi:

    if (param.otherParams.menuId > 0) {
      id = param.otherParams.menuId
    }

    switch (command) {

    case listMenus:

      return handle(api.ListAccountMenus(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getMenu:

      return handle(api.GetAccountMenu(accountId, id))

    case createMenu:

      params := createMenuParams(input)
      return handle(api.CreateAccountMenu(accountId, params))

    case replaceMenu:

      params := replaceMenuParams(input)
      return handle(api.ReplaceAccountMenu(accountId, id, params))

    case deleteMenu:

      return handle(api.DeleteAccountMenu(accountId, id))
    }

  case swagger.QueuesApi:

    if (param.otherParams.queueId > 0) {
      id = param.otherParams.queueId
    }

    switch (command) {

    case listQueues:

      return handle(api.ListAccountQueues(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getQueue:

      return handle(api.GetAccountQueue(accountId, id))

    case createQueue:

      params := createQueueParams(input)
      return handle(api.CreateAccountQueue(accountId, params))

    case replaceQueue:

      params := createQueueParams(input)
      return handle(api.ReplaceAccountQueue(accountId, id, params))

    case deleteQueue:

      return handle(api.DeleteAccountQueue(accountId, id))
    }

  case swagger.RoutesApi:

    if (param.otherParams.routeId > 0) {
      id = param.otherParams.routeId
    }

    switch (command) {

    case listRoutes:

      return handle(api.ListAccountRoutes(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getRoute:

      return handle(api.GetAccountRoute(accountId, id))

    case createRoute:

      params := createRouteParams(input)
      return handle(api.CreateRoute(accountId, params))

    case replaceRoute:

      params := createRouteParams(input)
      return handle(api.ReplaceAccountRoute(accountId, id, params))

    case deleteRoute:

      return handle(api.DeleteAccountRoute(accountId, id))
    }

  case swagger.SchedulesApi:

    if (param.otherParams.scheduleId > 0) {
      id = param.otherParams.scheduleId
    }

    switch (command) {

    case listSchedules:

      return handle(api.ListAccountSchedules(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getSchedule:

      return handle(api.GetAccountSchedule(accountId, id))
    }

  case swagger.SmsApi:

    if (param.otherParams.smsId != "") {
      idString = param.otherParams.smsId
    }

    switch (command) {

    case listSms:

      return handle(api.ListAccountSms(accountId, filtersId, filterParams.filtersDirection, filterParams.filtersFrom, sortParams.sortId, sortParams.sortCreatedAt, limit, offset, fields))

    case getSms:

      return handle(api.GetAccountSms(accountId, idString))

    case createSms:

      params := createSmsParams(param.from, param.to, param.text, param.otherParams.extensionId)
      return handle(api.CreateAccountSms(accountId, params))
    }

  case swagger.AvailablenumbersApi:

    switch (command) {

    case listAvailablePhoneNumbers:

      return handle(api.ListAvailablePhoneNumbers(filterParams.filtersPhoneNumber, filterParams.filtersCountryCode, filterParams.filtersNpa, filterParams.filtersNxx, filterParams.filtersXxxx, filterParams.filtersCity, filterParams.filtersProvince, filterParams.filtersCountry, filterParams.filtersPrice, filterParams.filtersCategory, sortParams.sortInternal, sortParams.sortPrice, sortParams.sortPhoneNumber, limit, offset, fields))
    }

  case swagger.SubaccountsApi:

    switch (command) {

    case listSubaccounts:

      return handle(api.ListAccountSubaccounts(accountId, filtersId, sortParams.sortId, limit, offset, fields))

    case createSubaccount:

      params := createSubaccountParams(input, param.contact, param.billingContact)
      return handle(api.CreateAccountSubaccount(accountId, params))
    }

  case swagger.AccountsApi:

    switch (command) {

    case listAccounts:

      return handle(api.ListAccounts(filtersId, sortParams.sortId, limit, offset, fields))

    case getAccount:

      return handle(api.GetAccount(id))
    }

  case swagger.NumberregionsApi:

    switch (command) {

    case listAvailablePhoneNumberRegions:

      return handle(api.ListAvailablePhoneNumberRegions(filterParams.filtersCountryCode, filterParams.filtersNpa, filterParams.filtersNxx, filterParams.filtersIsTollFree, filterParams.filtersCity, filterParams.filtersProvincePostalCode, filterParams.filtersCountryPostalCode, sortParams.sortCountryCode, sortParams.sortNpa, sortParams.sortNxx, sortParams.sortIsTollFree, sortParams.sortCity, sortParams.sortProvincePostalCode, sortParams.sortCountryPostalCode, limit, offset, fields, param.otherParams.groupBy))
    }

  case swagger.ApplicationsApi:

    switch (command) {

    case listApplications:

      return handle(api.ListAccountApplications(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getApplication:

      return handle(api.GetAccountApplication(accountId, id))
    }

  case swagger.CalllogsApi:

    switch (command) {

    case listCallLogs:

      return handle(api.ListAccountCallLogs(accountId, filtersId, filterParams.filtersStartTime, filterParams.filtersCreatedAt, filterParams.filtersDirection, filterParams.filtersCalledNumber, filterParams.filtersType, filterParams.filtersExtension, sortParams.sortId, sortParams.sortStartTime, sortParams.sortCreatedAt, limit, offset, fields))
    }

  case swagger.DevicesApi:

    if (param.otherParams.deviceId > 0) {
      id = param.otherParams.deviceId
    }

    switch (command) {

    case listDevices:

      return handle(api.ListAccountDevices(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getDevice:

      return handle(api.GetAccountDevice(accountId, id))

    case createDevice:

      params := createDeviceParams(input)
      return handle(api.CreateAccountDevice(accountId, params))

    case replaceDevice:

      params := createDeviceParams(input)
      return handle(api.ReplaceAccountDevice(accountId, id, params))
    }

  case swagger.ExpressservicecodesApi:

    if (param.otherParams.codeId > 0) {
      id = param.otherParams.codeId
    }

    switch (command) {

    case listExpressServiceCodes:

      return handle(api.ListAccountExpressSrvCodes(accountId, param.slice))

    case getExpressServiceCode:

      return handle(api.GetAccountExpressSrvCode(accountId, id))
    }

  case swagger.ExtensionsApi:

    switch (command) {

    case listExtensions:

      return handle(api.ListAccountExtensions(accountId, filtersId, filterParams.filtersExtension, filterParams.filtersName, sortParams.sortId, sortParams.sortExtension, sortParams.sortName, limit, offset, fields))

    case getExtension:

      return handle(api.GetAccountExtension(accountId, id))

    case createExtension:

      params := createExtensionsParams(input)
      return handle(api.CreateAccountExtension(accountId, params))

    case replaceExtension:

      params := replaceExtensionParams(input)
      return handle(api.ReplaceAccountExtension(accountId, id, params))

    }

  case swagger.CalleridsApi:

    switch (command) {

    case getCallerId:
      return handle(api.GetCallerIds(accountId, id, filterParams.filtersNumber, filterParams.filtersName, sortParams.sortNumber, sortParams.sortName, limit, offset, fields))
    }

  case swagger.ContactsApi:

    if (param.otherParams.extensionId > 0) {
      id = param.otherParams.extensionId
    }

    if (param.otherParams.contactId > 0) {
      idSecondary = param.otherParams.contactId
    }

    switch (command) {

    case listContacts:

      return handle(api.ListAccountExtensionContacts(accountId, id, filtersId, filterParams.filtersGroupId, filterParams.filtersUpdatedAt, sortParams.sortId, sortParams.sortUpdatedAt, limit, offset, fields))

    case getContact:

      return handle(api.GetAccountExtensionContact(accountId, id, idSecondary))

    case createContact:

      params := createContactParams(input)
      return handle(api.CreateAccountExtensionContact(accountId, id, params))

    case replaceContact:

      params := createContactParams(input)
      return handle(api.ReplaceAccountExtensionContact(accountId, id, idSecondary, params))

    case deleteContact:

      return handle(api.DeleteAccountExtensionContact(accountId, id, idSecondary))
    }

  case swagger.GroupsApi:

    if (param.otherParams.extensionId > 0) {
      id = param.otherParams.extensionId
    }

    if (param.otherParams.groupId > 0) {
      idSecondary = param.otherParams.groupId
    }

    switch (command) {

    case listGroups:

      return handle(api.ListAccountExtensionContactGroups(accountId, id, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getGroup:

      return handle(api.GetAccountExtensionContactGroup(accountId, id, idSecondary))

    case createGroup:

      params := createGroupParams(input)
      return handle(api.CreateAccountExtensionContactGroup(accountId, id, params))

    case replaceGroup:

      params := createGroupParams(input)
      return handle(api.ReplaceAccountExtensionContactGroup(accountId, id, idSecondary, params))
    case deleteGroup:

      return handle(api.DeleteAccountExtensionContactGroup(accountId, id, idSecondary))
    }

  case swagger.PhonenumbersApi:

    if (param.otherParams.numberId > 0) {
      id = param.otherParams.numberId
    }

    switch (command) {

    case listPhoneNumbers:

      return handle(api.ListAccountPhoneNumbers(accountId, filtersId, filterParams.filtersName, filterParams.filtersPhoneNumber, sortParams.sortId, sortParams.sortName, sortParams.sortPhoneNumber, limit, offset, fields))

    case getPhoneNumber:

      return handle(api.GetAccountPhoneNumber(accountId, id))

    case createPhoneNumber:

      params := createPhoneNumberParams(input)
      return handle(api.CreateAccountPhoneNumber(accountId, params))

    case replacePhoneNumber:

      params := replacePhoneNumberParams(input)
      return handle(api.ReplaceAccountPhoneNumber(accountId, id, params))
    }

  case swagger.TrunksApi:

    if (param.otherParams.trunkId > 0) {
      id = param.otherParams.trunkId
    }

    switch (command) {

    case listTrunks:

      return handle(api.ListAccountTrunks(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getTrunk:

      return handle(api.GetAccountTrunk(accountId, id))

    case createTrunk:
      params := createTrunkParams(param.trunkName, param.trunkUri, param.trunkConcurrentCalls, param.trunkMaxMinutes)
      return handle(api.CreateAccountTrunk(accountId, params))

    case replaceTrunk:

      params := createTrunkParams(param.trunkName, param.trunkUri, param.trunkConcurrentCalls, param.trunkMaxMinutes)
      return handle(api.ReplaceAccountTrunk(accountId, id, params))

    case deleteTrunk:

      return handle(api.DeleteAccountTrunk(accountId, id))
    }

  default:
    return nil, nil
  }

  return nil, nil
}
