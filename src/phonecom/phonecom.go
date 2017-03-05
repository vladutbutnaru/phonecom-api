package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "phonecom-go-sdk"
  "errors"
  "encoding/json"
  "io/ioutil"
  "strings"
  "time"
  "encoding/xml"
)

var outputType = "json"
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

	param, err := createCliParams(c);

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

	swaggerConfig := cliConfig.createSwaggerConfig(cliConfig)
  api := getApi(swaggerConfig, param.command)

	param.accountId = cliConfig.AccountId

  if (api == nil) {
    return errors.New(msgCouldNotGetResponse), nil
  }

  if param.sortType != "" && param.sortValue != "" {
    switch param.sortType {
    case "id":
      param.sortParams.sortId = param.sortValue
    case "name":
			param.sortParams.sortName = param.sortValue
    case "start_time":
			param.sortParams.sortStartTime = param.sortValue
    case "created_at":
			param.sortParams.sortCreatedAt = param.sortValue
    case "extension":
			param.sortParams.sortExtension = param.sortValue
    case "number":
			param.sortParams.sortNumber = param.sortValue
    case "updated_at":
			param.sortParams.sortUpdatedAt = param.sortValue
    case "phone_number":
			param.sortParams.sortPhoneNumber = param.sortValue
    case "internal":
			param.sortParams.sortInternal = param.sortValue
    case "price":
			param.sortParams.sortPrice = param.sortValue
    case "npa":
			param.sortParams.sortNpa = param.sortValue
    case "nxx":
			param.sortParams.sortNxx = param.sortValue
    case "is_toll_free":
			param.sortParams.sortIsTollFree = param.sortValue
    case "city":
			param.sortParams.sortCity = param.sortValue
    case "province_postal_code":
			param.sortParams.sortProvincePostalCode = param.sortValue
    case "country_postal_code":
			param.sortParams.sortCountryPostalCode = param.sortValue
    }
  }

  if param.filterType != "" && param.filterValue != "" {
    slice1 := make([]string, 0)
    slice1 = append(slice1, param.filterValue)
    switch param.filterType {
    case "id":
			param.filtersId = slice1
    case "name":
			param.filterParams.filtersName = slice1
    case "start_time":
			param.filterParams.filtersStartTime = slice1
    case "created_at":
			param.filterParams.filtersCreatedAt = param.filterValue
    case "direction":
			param.filterParams.filtersDirection = param.filterValue
    case "called_number":
			param.filterParams.filtersCalledNumber = param.filterValue
    case "type":
			param.filterParams.filtersType = param.filterValue
    case "extension":
			param.filterParams.filtersExtension = slice1
    case "number":
			param.filterParams.filtersNumber = slice1
    case "group_id":
			param.filterParams.filtersGroupId = slice1
    case "updated_at":
			param.filterParams.filtersUpdatedAt = slice1
    case "phone_number":
			param.filterParams.filtersPhoneNumber = slice1
    case "from":
			param.filterParams.filtersFrom = param.filterValue
    case "to":
			param.filterParams.filtersTo = slice1
    case "country_code":
			param.filterParams.filtersCountryCode = slice1
    case "npa":
			param.filterParams.filtersNpa = slice1
    case "nxx":
			param.filterParams.filtersNxx = slice1
    case "xxxx":
			param.filterParams.filtersXxxx = slice1
    case "city":
			param.filterParams.filtersCity = slice1
    case "province":
			param.filterParams.filtersProvince = slice1
    case "country":
			param.filterParams.filtersCountry = slice1
    case "price":
			param.filterParams.filtersPrice = slice1
    case "category":
			param.filterParams.filtersCategory = slice1
    case "is_toll_free":
			param.filterParams.filtersIsTollFree = slice1
    case "province_postal_code":
			param.filterParams.filtersProvincePostalCode = slice1
    case "country_postal_code":
			param.filterParams.filtersCountryPostalCode = slice1
    }
  }

  inputType := "json"

  if (strings.EqualFold(param.inputFormat, "xml")) {
    inputType = "xml"
  }

  if (strings.EqualFold(param.inputFormat, "csv")) {
    inputType = "csv"
  }

  if (strings.EqualFold(param.outputFormat, "csv")) {
    outputType = "csv"
  }

  if param.samplein != "" {

    stringEmailSlice := make([]string, 0)
    stringEmailSlice = append(stringEmailSlice, "asd@asd.com")

    switch param.samplein {
    case createDevice:
      createDeviceParamsSample := swagger.CreateDeviceParams{randomString(12), nil}
      marshalInput(createDeviceParamsSample, "createDevice", inputType)

    case createExtension:
      createExtensionParamsSample := swagger.CreateExtensionParams{"+12019570328", "unlimited", true, int32(randomNumber(10, 9999999)), true, "The name", "The full name", "America/Los_Angeles", swagger.MediaSummary{int32(randomNumber(10, 99999)), randomString(12)}, swagger.MediaSummary{int32(randomNumber(10, 99999)), randomString(12)}, 619, true, false, true, false, 12345, "standard", swagger.MediaSummary{int32(randomNumber(10, 99999)), randomString(12)}, "automated", stringEmailSlice, "+18587741111", stringEmailSlice, "+18587748888"}
      marshalInput(createExtensionParamsSample, "createExtension", inputType)

    case createContact:
      createContactParamsSample := swagger.CreateContactParams{"Geordi", "middle name", "last name", "prefix", "phoneticFirstName", "phoneticMiddleName", "phoneticLastName", "suffix", "nickname", "company", "department", "jobTitle", nil, nil, nil, nil}
      marshalInput(createContactParamsSample, "createContact", inputType)

    case createGroup:
      createGroupParamsSample := swagger.CreateGroupParams{"Ferengi Traders"}
      marshalInput(createGroupParamsSample, "createGroup", inputType)

    case createMenu:
      createMenuParamsSample := swagger.CreateMenuParams{randomString(12), nil, nil, true, 3, nil, nil}
      marshalInput(createMenuParamsSample, "createMenu", inputType)

    case createPhoneNumber:
      createPhoneNumberParamsSample := swagger.CreatePhoneNumberParams{"+12546551377", swagger.RouteSummary{123, randomString(12)}, "Phone Name Now", true, true, "Phone N", "business", "extension", swagger.ApplicationSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.ExtensionSummary{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999))}, stringEmailSlice, "+18587740222"}
      marshalInput(createPhoneNumberParamsSample, "createPhoneNumber", inputType)

    case createQueue:
      createQueueParamsSample := swagger.CreateQueueParams{randomString(12), swagger.MediaSummary{123, randomString(12)}, swagger.MediaSummary{123, randomString(12)}, 60, "called_number", 10, nil}
      marshalInput(createQueueParamsSample, "createQueue", inputType)

    case createRoute:
      createRouteJsonSample := CreateRouteJson{randomString(12), []RulesJson{RulesJson{[]ActionsJson{ActionsJson{"queue", QueueJson{int32(22035), "ntud62prqbl7"}}}}}}
      marshalInput(createRouteJsonSample, "createRoute", inputType)

    case createSms:
      createSmsParamsSample := swagger.CreateSmsParams{"+16309624775", "+12019570328", "Another message for create", 1767963}
      marshalInput(createSmsParamsSample, "createSms", inputType)

    case createSubaccount:
      createSubaccountJsonSample := CreateSubaccountJson{randomString(12), randomString(12), ContactJson{"Bobby", AddressJson{"100 Main St", "San Diego", "CA", "92129", "US"}, "+18585553333", "asd@sd.co"},
        ContactJson{"Bobby", AddressJson{"100 Main St", "San Diego", "CA", "92129", "US"}, "+18585553333", "asd@sd.co"}}
      marshalInput(createSubaccountJsonSample, "createSubaccount", inputType)

    case createTrunk:
      createTrunkParamsSample := swagger.CreateTrunkParams{randomString(12), "SIP/1234@phone.com:5060", int32(60), int32(800), swagger.MediaSummary{123, randomString(12)}, swagger.MediaSummary{123, randomString(12)}, nil}
      marshalInput(createTrunkParamsSample, "createTrunk", inputType)

    case replaceDevice:
      marshalInput(createDeviceParams, "replaceDevice", inputType)

    case replaceExtension:
      replaceExtensionParamsSample := swagger.ReplaceExtensionParams{nil, nil, randomString(12), "America/Los_Angeles", true, 111, true, "unlimited", 12344, "bobby McFerrin", true, nil, "standard", "private", 619, true, true, "automated", nil, "+18587741111", nil, "+18587748888", nil}
      marshalInput(replaceExtensionParamsSample, "replaceExtension", inputType)

    case replaceMenu:
      replaceMenuParamsSample := swagger.ReplaceMenuParams{randomString(12), nil, nil, false, 5, nil, nil}
      marshalInput(replaceMenuParamsSample, "replaceMenu", inputType)

    case replacePhoneNumber:
      replacePhoneNumberParamsSample := swagger.ReplacePhoneNumberParams{swagger.RouteSummary{123, randomString(12)}, "Robert", true, true, "Phone N", "business", "extension", swagger.ApplicationSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.ExtensionSummary{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999))}, nil, stringEmailSlice, "+18587740222"}
      marshalInput(replacePhoneNumberParamsSample, "replacePhoneNumber", inputType)

    case replaceQueue:
      createQueueParamsSample := swagger.CreateQueueParams{randomString(12), swagger.MediaSummary{123, randomString(12)}, swagger.MediaSummary{123, randomString(12)}, 60, "called_number", 10, nil}
      marshalInput(createQueueParamsSample, "replaceQueue", inputType)

    case replaceRoute:
      replaceRouteJsonSample := ReplaceRouteJson{int32(4705073), randomString(12), []RulesJson{RulesJson{[]ActionsJson{ActionsJson{"queue", QueueJson{22026, "61kkjklmin74"}}}}}}
      marshalInput(replaceRouteJsonSample, "replaceRoute", inputType)

    case replaceTrunk:
      createTrunkParamsSample := swagger.CreateTrunkParams{randomString(12), "SIP/1234@phone.com:5060", int32(60), int32(800), swagger.MediaSummary{123, randomString(12)}, swagger.MediaSummary{123, randomString(12)}, nil}
      marshalInput(createTrunkParamsSample, "replaceTrunk", inputType)

    }

    return nil, nil
  }

  if (param.sampleout != "") {
    stringEmailSlice := make([]string, 0)
    stringEmailSlice = append(stringEmailSlice, "asd@asd.com")

    switch param.sampleout {
    case getAccount:
      accountFullSample := swagger.AccountFull{int32(randomNumber(10, 9999)), randomString(12), randomString(12), randomString(12), swagger.AccountSummary{int32(randomNumber(10, 9999)), randomString(12)}, swagger.ContactAccount{randomString(12), randomString(12), swagger.Address{randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), strings.ToUpper(randomAlphaString(2))}, randomString(12), randomString(12), "primary@email.com", "alternate@email.com"}, swagger.ContactAccount{randomString(12), randomString(12), swagger.Address{randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), strings.ToUpper(randomAlphaString(2))}, randomString(12), randomString(12), "primary@email.com", "alternate@email.com"}}
      marshalInput(accountFullSample, "getAccount", inputType)

    case getApplication:
      applicationFullSample := swagger.ApplicationFull{int32(randomNumber(10, 9999)), randomString(12)}
      marshalInput(applicationFullSample, "getApplication", inputType)

    case getDevice:
      deviceFullSample := swagger.DeviceFull{int32(randomNumber(10, 9999)), randomString(12), swagger.SipAuthentication{randomString(12), int32(randomNumber(10, 9999)), randomString(12), randomString(12)}, nil/*[]swagger.Line{int32(randomNumber(1,9999)), swagger.ExtensionSummary{int32(randomNumber(1,9999)), randomString(12), int32(randomNumber(1,9999))}}*/ }
      marshalInput(deviceFullSample, "getDevice", inputType)

    case getExpressServiceCode:
      expressServiceCodeFullSample := swagger.ExpressServiceCodeFull{int32(randomNumber(1, 9999)), randomNumericString(8), int32(randomNumber(9999, 9999999999))}
      marshalInput(expressServiceCodeFullSample, "getExpressServiceCode", inputType)

    case getExtension:
      extensionFullSample := swagger.ExtensionFull{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999)), randomString(12), randomString(12), swagger.DeviceMembership{int32(randomNumber(1, 9999)), swagger.DeviceSummary{int32(randomNumber(1, 9999)), randomString(12)}}, randomString(12), swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, true, "+454654564534", randomString(12), true, true, swagger.Voicemail{true, randomString(12), swagger.Greeting{randomString(12), swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, true}, randomString(12), swagger.Notification{stringEmailSlice, randomString(12)}, randomString(12)}, swagger.Notification{stringEmailSlice, randomString(12)}, swagger.RouteSummary{int32(randomNumber(1, 9999)), randomString(12)}}
      marshalInput(extensionFullSample, "getExtension", inputType)

    case getContact:
      emailSlice := make([]swagger.Email, 0)
      emailSlice = append(emailSlice, swagger.Email{"primary", "asd@asd.com"})
      phoneNumberContactslice := make([]swagger.PhoneNumberContact, 0)
      phoneNumberContactslice = append(phoneNumberContactslice, swagger.PhoneNumberContact{"home", "+45454684545", "+45454684545"})
      addressListContacts := make([]swagger.AddressListContacts, 0)
      addressListContacts = append(addressListContacts, swagger.AddressListContacts{"home", randomString(12), randomString(12), randomString(12), randomNumericString(5), strings.ToUpper(randomAlphaString(2))})
      contactFullSample := swagger.ContactFull{int32(randomNumber(1, 9999)), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), emailSlice, phoneNumberContactslice, addressListContacts, swagger.GroupListContacts{int32(randomNumber(1, 9999)), randomString(12)}, int32(randomNumber(9999, 9999999999)), int32(randomNumber(9999, 9999999999))}
      marshalInput(contactFullSample, "getContact", inputType)

    case getGroup:
      groupFullSample := swagger.GroupFull{int32(randomNumber(1, 9999)), randomString(12)}
      marshalInput(groupFullSample, "getGroup", inputType)

    case getRecording:
      mediaFullSample := swagger.MediaFull{int32(randomNumber(1, 9999)), randomString(12), "hold_music"}
      marshalInput(mediaFullSample, "getRecording", inputType)

    case getMenu:
      optionSlice := make([]swagger.Option, 0)
      optionSlice = append(optionSlice, swagger.Option{randomNumericString(1), swagger.RouteSummary{int32(randomNumber(1,9999)), randomString(12)}})
      menuFullSample := swagger.MenuFull{int32(randomNumber(1, 9999)), randomString(12), true, int32(randomNumber(1, 9999)), swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.RouteSummary{int32(randomNumber(1, 9999)), randomString(12)}, optionSlice}
      marshalInput(menuFullSample, "getMenu", inputType)

    case getPhoneNumber:
      phoneNumberFullSample := swagger.PhoneNumberFull{int32(randomNumber(1, 9999)), randomString(12), "+54654612511", true, true, swagger.RouteSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.CallerIdPhoneNumber{randomString(12), "bussiness"}, swagger.SmsForwarding{"extension", swagger.ExtensionSummary{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999))}, swagger.ApplicationSummary{int32(randomNumber(1, 9999)), randomString(12)}}, swagger.CallNotifications{stringEmailSlice, "+45456486464"}}
      marshalInput(phoneNumberFullSample, "getPhoneNumber", inputType)

    case getQueue:
      memberSlice := make([]swagger.Member, 0)
      memberSlice = append(memberSlice, swagger.Member{swagger.ExtensionSummary{int32(randomNumber(1,9999)), randomString(12), int32(randomNumber(1,9999))}, randomString(12)})
      queueFullSample := swagger.QueueFull{int32(randomNumber(1, 9999)), randomString(12), swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.HoldMusic{int32(randomNumber(1, 9999)), randomString(12)}, 300, randomString(12), 20, memberSlice}
      marshalInput(queueFullSample, "getQueue", inputType)

    case getRoute:
      ruleSetForwardItemSlice := make([]swagger.RuleSetForwardItem, 0)
      ruleSetForwardItemSlice = append(ruleSetForwardItemSlice, swagger.RuleSetForwardItem{"+154514214546", swagger.ExtensionSummary{int32(randomNumber(1,9999)), randomString(12), int32(randomNumber(1,9999))}, "+453484845122", true, "calling_number", randomString(12), "DEFAULT"})
      ruleSetActionSlice := make([]swagger.RuleSetAction, 0)
      ruleSetActionSlice = append(ruleSetActionSlice, swagger.RuleSetAction{randomString(12), swagger.ExtensionSummary{int32(randomNumber(1,9999)), randomString(12), int32(randomNumber(1,9999))}, ruleSetForwardItemSlice, int32(randomNumber(5,90)), swagger.MediaSummary{int32(randomNumber(1,9999)), randomString(12)}, swagger.MediaSummary{int32(randomNumber(1,9999)), randomString(12)}, int32(randomNumber(1, 60)), swagger.MenuSummary{int32(randomNumber(1,9999)), randomString(12)}, swagger.QueueSummary{int32(randomNumber(1,9999)), randomString(12)}, swagger.TrunkSummary{int32(randomNumber(1,9999)), randomString(12)}})
      rulesetSlice := make([]swagger.RuleSet, 0)
      rulesetSlice = append(rulesetSlice, swagger.RuleSet{swagger.RuleSetFilter{"schedule", swagger.ScheduleSummary{int32(randomNumber(1,9999)), randomString(12)}, swagger.ContactSummary{int32(randomNumber(1,9999)), "Mr", randomString(12), randomString(12), randomString(12), "Jr", randomString(12), randomString(12)}, swagger.GroupSummary{int32(randomNumber(1,9999)), randomString(12)}}, ruleSetActionSlice})
      routeFullSample := swagger.RouteFull{int32(randomNumber(1, 9999)), randomString(12), swagger.ExtensionSummary{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999))}, rulesetSlice}
      marshalInput(routeFullSample, "getRoute", inputType)

    case getSchedule:
      scheduleFullSample := swagger.ScheduleFull{int32(randomNumber(1, 9999)), randomString(12)}
      marshalInput(scheduleFullSample, "getSchedule", inputType)

    case getSms:
      recipientSlice := make([]swagger.Recipient, 0)
      recipientSlice = append(recipientSlice, swagger.Recipient{"+12454354513", "sent"})
      smsFullSample := swagger.SmsFull{randomSmsId(), "+5646517686", recipientSlice, "in", int32(randomNumber(9999, 9999999999)), time.Time{}, randomString(12)}
      marshalInput(smsFullSample, "getSms", inputType)

    case getTrunk:
      stringCodeSlice := make([]string, 0)
      stringCodeSlice = append(stringCodeSlice, "g711u 64k")
      trunkFullSample := swagger.TrunkFull{int32(randomNumber(1, 9999)), randomString(12), "SIP/01%e@243.1.45.52:5060", int32(randomNumber(1, 100)), int32(randomNumber(500, 2000)), swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, stringCodeSlice}
      marshalInput(trunkFullSample, "getTrunk", inputType)

    }
    return nil, nil
  }

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

func getApi(
    config *swagger.Configuration,
    command string) interface{} {

  var api interface{}

  switch (command) {

  case listMedia, getRecording:

    mediaApi := *swagger.NewMediaApi()
    mediaApi.Configuration = config
    api = mediaApi

  case listMenus, getMenu, createMenu, replaceMenu, deleteMenu:

    menusApi := *swagger.NewMenusApi()
    menusApi.Configuration = config
    api = menusApi

  case listQueues, getQueue, createQueue, replaceQueue, deleteQueue:

    queuesApi := *swagger.NewQueuesApi()
    queuesApi.Configuration = config
    api = queuesApi

  case listRoutes, getRoute, createRoute, replaceRoute, deleteRoute:

    routesApi := *swagger.NewRoutesApi()
    routesApi.Configuration = config
    api = routesApi

  case listSchedules, getSchedule:

    schedulesApi := *swagger.NewSchedulesApi()
    schedulesApi.Configuration = config
    api = schedulesApi

  case listSms, getSms, createSms:

    smsApi := *swagger.NewSmsApi()
    smsApi.Configuration = config
    api = smsApi

  case listAvailablePhoneNumbers:

    availableNumbersApi := *swagger.NewAvailablenumbersApi()
    availableNumbersApi.Configuration = config
    api = availableNumbersApi

  case listAvailablePhoneNumberRegions:

    availableNumbersApi := *swagger.NewNumberregionsApi()
    availableNumbersApi.Configuration = config
    api = availableNumbersApi

  case listSubaccounts, createSubaccount:

    subAccountsApi := *swagger.NewSubaccountsApi()
    subAccountsApi.Configuration = config
    api = subAccountsApi

  case listAccounts, getAccount:

    accountsApi := *swagger.NewAccountsApi()
    accountsApi.Configuration = config
    api = accountsApi

  case listApplications, getApplication:

    applicationsApi := *swagger.NewApplicationsApi()
    applicationsApi.Configuration = config
    api = applicationsApi

  case listCallLogs:

    callLogsApi := *swagger.NewCalllogsApi()
    callLogsApi.Configuration = config
    api = callLogsApi

  case listDevices, getDevice, createDevice, replaceDevice:

    devicesApi := *swagger.NewDevicesApi()
    devicesApi.Configuration = config
    api = devicesApi

  case listExpressServiceCodes, getExpressServiceCode:

    serviceCodesApi := *swagger.NewExpressservicecodesApi()
    serviceCodesApi.Configuration = config
    api = serviceCodesApi

  case listExtensions, getExtension, createExtension, replaceExtension:

    extensionsApi := *swagger.NewExtensionsApi()
    extensionsApi.Configuration = config
    api = extensionsApi

  case getCallerId:

    callerIdsApi := *swagger.NewCalleridsApi()
    callerIdsApi.Configuration = config
    api = callerIdsApi

  case listContacts, getContact, createContact, replaceContact, deleteContact:

    contactsApi := *swagger.NewContactsApi()
    contactsApi.Configuration = config
    api = contactsApi

  case listGroups, getGroup, createGroup, replaceGroup, deleteGroup:

    groupsApi := *swagger.NewGroupsApi()
    groupsApi.Configuration = config
    api = groupsApi

  case listPhoneNumbers, getPhoneNumber, createPhoneNumber, replacePhoneNumber:

    phoneNumbersApi := *swagger.NewPhonenumbersApi()
    phoneNumbersApi.Configuration = config
    api = phoneNumbersApi

  case listTrunks, getTrunk, createTrunk, deleteTrunk, replaceTrunk:

    trunksApi := *swagger.NewTrunksApi()
    trunksApi.Configuration = config
    api = trunksApi

  default:
    fmt.Printf("Invalid command: %v\n", command)
    return nil
  }

  return api
}

func handle(
    x interface{},
    response *swagger.APIResponse,
    error error) (error, map[string] interface{}) {

  if (error != nil) {

    if (param.verbose) {
      fmt.Println("Error while getting response")
    }

    return error, nil
  }

  payload := response.Payload

  if (payload == nil) {

    if (param.verbose) {
      fmt.Println("Null response payload")
    }

    return errors.New(msgCouldNotGetResponse), nil
  }

  validatedJson := validateJson(string(payload))

  if (validatedJson == nil) {

    if (param.verbose) {
      fmt.Println("Could not unmarshal API json response")
    }

    return errors.New(msgInvalidJson), nil
  }

  message := validateResponse(validatedJson)

  if (message != "") {

    if (param.verbose) {
      fmt.Printf("%+v\n%s\n", x, response)
    }

    return errors.New(message), nil
  } else {

    var jsonObject interface{}
    if (!param.fullList && validatedJson["items"] != nil) {
      jsonObject = validatedJson["items"]
    } else {
      jsonObject = validatedJson
    }

    if (param.verbose) {
      fmt.Printf("\nAPI Response [Verbose]:\n%+v\n", x)
      fmt.Println()
    }

    fmt.Println("API Response:")
    fmt.Println()

    if (outputType == "csv") {
      exportToCsv(jsonObject)
    } else if (outputType == "json") {
			jsonOutput, err := json.MarshalIndent(jsonObject, "", "  ")
			if (err != nil) {
				return err, nil
			}
      fmt.Printf("%s\n", jsonOutput)
		} else {
			fmt.Println("Invalid output type")
		}

  }

  return nil, validatedJson
}

func marshalInput(param interface{}, fileName string, outputType string) {

  var marshalled []byte
  var err error

  if (outputType == "json") {
    marshalled, err = json.MarshalIndent(param, "", "  ")
  } else if (outputType == "xml") {
    marshalled, err = xml.MarshalIndent(param, "", "  ")
  }

  if (err == nil) {
    err = ioutil.WriteFile(fileName + "." + outputType, marshalled, 0644)
  }

  if (err != nil) {

    fmt.Printf("Could not create sample %s\n", outputType)
  } else {
    fmt.Println("Sample " + outputType + " created successfully")
  }

}
