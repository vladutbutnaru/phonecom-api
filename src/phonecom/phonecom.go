package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "phonecom-go-sdk"
  "errors"
  "strconv"
  "encoding/json"
  "io/ioutil"
)

var configPath = "config.xml" // Used as a variable. To be changed in tests.
var param CliParams

func main() {

  app := cli.NewApp()

  app.Flags = getCliFlags()

  app.Action = func(c *cli.Context) error {
    var err error
    err, _ = execute(c)
    return err
  }

  app.Run(os.Args)
}

type CliParams struct {
  slice       []string
  filtersId   []string
  limit       int32
  offset      int32
  id          int32
  idString    string
  dryRun      bool
  verbose     bool
  input       string
  command     string
  idSecondary int32
  accountId   int32
  fields      string
  contact     string
  billingContact string
  filterType  string
  filterValue string
}

func execute(
c *cli.Context) (error, map[string] interface{}) {

  slice := make([]string, 0)
  limit := int32(c.Int("limit"))
  offset := int32(c.Int("offset"))
  idString := c.String("id")
  var id int32 = 0

  if _, err := strconv.Atoi(idString); err == nil {
    idInt := 0
    idInt, err = strconv.Atoi(idString);
    id = int32(idInt);
  }

  verbose := c.Bool("verbose")
  dryRun := c.Bool("dryrun")
  input := c.String("input")
  command := c.String("command")
  idSecondary := int32(c.Int("id-secondary"))
  accountId := int32(c.Int("account"))
  contact := c.String("contact")
  billingContact := c.String("billing-contact")
  fields := ""
  filtersType := c.String("filtersType")
  filtersValue := c.String("filtersValue")
  sortType := c.String("sortType")
  sortValue := c.String("sortValue")
  sample := c.String("sample")


  var filtersId []string
  var groupBy []string
  var extensionId int32
  var deviceId int32
  var codeId int32
  var contactId int32
  var groupId int32
  var recordingId int32
  var menuId int32
  var numberId int32
  var queueId int32
  var routeId int32
  var scheduleId int32
  var smsId string
  var trunkId int32

  var from string
  var to string
  var text string
  from = c.String("from")
  to = c.String("to")
  text = c.String("text")

  from = defaultFrom
  to = defaultTo
  text = defaultText

  var trunkName = c.String("name")
  var trunkUri = c.String("uri")
  var trunkConcurrentCalls = int32(c.Int("max-concurrent-calls"))
  var trunkMaxMinutes = int32(c.Int("max-minutes-per-month"))

  trunkName = defaultTrunkName
  trunkUri = defaultTrunkUri
  trunkConcurrentCalls = int32(defaultTrunkConcurrentCalls)
  trunkMaxMinutes = int32(defaultTrunkMaxMinutes)

  var filterParams FilterParams
  var sortParams SortParams
  var otherParams OtherParams


  if (input != "") {
    var err error
    var listParams ListParams
    err, listParams = getListParams(input)
    if (err != nil) {
      return err, nil
    }

    accountId = listParams.accountId
    limit = listParams.limit
    offset = listParams.offset
    fields = listParams.fields

    err, filterParams = getFiltersParams(input)

    if (err != nil) {
      return err, nil
    }

    filtersId = filterParams.filtersId

    err, sortParams = getSortParams(input)

    if (err != nil) {
      return err, nil
    }

    err, otherParams = getOtherParams(input)
    groupBy = otherParams.groupBy
    extensionId = otherParams.extensionId
    deviceId = otherParams.deviceId
    codeId = otherParams.codeId
    contactId = otherParams.contactId
    groupId = otherParams.groupId
    recordingId = otherParams.recordingId
    menuId = otherParams.menuId
    numberId = otherParams.numberId
    queueId = otherParams.queueId
    routeId = otherParams.routeId
    scheduleId = otherParams.scheduleId
    smsId = otherParams.smsId
    trunkId = otherParams.trunkId

    if (err != nil) {
      return err, nil
    }
  }

  param.slice = slice
  param.limit = limit
  param.offset = offset
  param.id = id
  param.idString = idString
  param.dryRun = dryRun
  param.verbose = verbose
  param.input = input
  param.command = command
  param.idSecondary = idSecondary
  param.accountId = accountId
  param.fields = fields
  param.contact = contact
  param.billingContact = billingContact

  showDryRunVerbose(param)
  if (param.dryRun) {
    return nil, nil
  }

  api, config := getApi(command)
  accountId = config.AccountId

  if (api == nil) {
    return errors.New(msgCouldNotGetResponse), nil
  }

  if sortType != "" && sortValue != "" {
    switch sortType {
    case "id":
      sortParams.sortId = sortValue
    case "name":
      sortParams.sortName = sortValue
    case "start_time":
      sortParams.sortStartTime = sortValue
    case "created_at":
      sortParams.sortCreatedAt = sortValue
    case "extension":
      sortParams.sortExtension = sortValue
    case "number":
      sortParams.sortNumber = sortValue
    case "updated_at":
      sortParams.sortUpdatedAt = sortValue
    case "phone_number":
      sortParams.sortPhoneNumber = sortValue
    case "internal":
      sortParams.sortInternal = sortValue
    case "price":
      sortParams.sortPrice = sortValue
    case "npa":
      sortParams.sortNpa = sortValue
    case "nxx":
      sortParams.sortNxx = sortValue
    case "is_toll_free":
      sortParams.sortIsTollFree = sortValue
    case "city":
      sortParams.sortCity = sortValue
    case "province_postal_code":
      sortParams.sortProvincePostalCode = sortValue
    case "country_postal_code":
      sortParams.sortCountryPostalCode = sortValue
    }
  }

  if filtersType != "" && filtersValue != "" {
    slice := append(slice, filtersValue)
    switch filtersType {
    case "id":
      filtersId = slice
    case "name":
      filterParams.filtersName = slice
    case "start_time":
      filterParams.filtersStartTime = slice
    case "created_at":
      filterParams.filtersCreatedAt = filtersValue
    case "direction":
      filterParams.filtersDirection = filtersValue
    case "called_number":
      filterParams.filtersCalledNumber = filtersValue
    case "type":
      filterParams.filtersType = filtersValue
    case "extension":
      filterParams.filtersExtension = slice
    case "number":
      filterParams.filtersNumber = slice
    case "group_id":
      filterParams.filtersGroupId = slice
    case "updated_at":
      filterParams.filtersUpdatedAt = slice
    case "phone_number":
      filterParams.filtersPhoneNumber = slice
    case "from":
      filterParams.filtersFrom = filtersValue
    case "to":
      filterParams.filtersTo = slice
    case "country_code":
      filterParams.filtersCountryCode = slice
    case "npa":
      filterParams.filtersNpa = slice
    case "nxx":
      filterParams.filtersNxx = slice
    case "xxxx":
      filterParams.filtersXxxx = slice
    case "city":
      filterParams.filtersCity = slice
    case "province":
      filterParams.filtersProvince = slice
    case "country":
      filterParams.filtersCountry = slice
    case "price":
      filterParams.filtersPrice = slice
    case "category":
      filterParams.filtersCategory = slice
    case "is_toll_free":
      filterParams.filtersIsTollFree = slice
    case "province_postal_code":
      filterParams.filtersProvincePostalCode = slice
    case "country_postal_code":
      filterParams.filtersCountryPostalCode = slice
    }
  }

  if sample != "" {
    randomName := randomString(12)
    randomPassword := randomString(12)
    randomNum := randomNumber(10, 9999999)
    switch sample {
      case createDevice:
        marshalJson(CreateDeviceJson{defaultAccountId, randomName}, "createDevice.json")

      case createExtension:
        marshalJson(
          CreateExtensionJson{defaultAccountId, "+12019570328", "unlimited", true, int32(randomNum), true, "The name", "The full name", "America/Los_Angeles", 619, true, false, true, false, 12345, "standard", "automated", "+18587741111", "+18587748888"},
          "createExtension.json")

      case createContact:
        marshalJson(CreateContactJson{defaultAccountId, 1764590, "Geordi", "middle name", "last name", "prefix", "phoneticFirstName", "phoneticMiddleName", "phoneticLastName", "suffix", "nickname", "company", "department", "jobTitle"},
          "createContact.json")

      case createGroup:
        marshalJson(CreateGroupJson{defaultAccountId, 1764590, "Ferengi Traders"}, "createGroup.json")

      case createMenu:
        marshalJson(CreateMenuJson{defaultAccountId, randomName, true, 3}, "createMenu.json")

      case createPhoneNumber:
        marshalJson(CreatePhoneNumberJson{defaultAccountId, "+12546551377", "Phone Name Now", true, true, "Phone N", "business", "extension", "+18587740222"},
          "createPhoneNumber.json")
      case createQueue:
        marshalJson(CreateQueueJson{defaultAccountId, randomName, 60, "called_number", 10}, "createQueue.json")

      case createRoute:
        marshalJson(CreateRouteJson{defaultAccountId, randomName, []RulesJson{RulesJson{[]ActionsJson{ActionsJson{"queue", QueueJson{int32(22035), "ntud62prqbl7"}}}}}},
          "createRoute.json")

      case createSms:
        marshalJson(CreateSmsJson{defaultAccountId, "+16309624775", "+12019570328", "Another message for create", 1767963}, "createSms.json")

      case createSubaccount:
        marshalJson(CreateSubaccountJson{defaultAccountId, randomName, randomPassword, ContactJson{"Bobby", AddressJson{"100 Main St", "San Diego", "CA", "92129", "US"}, "+18585553333", "asd@sd.co"},
          ContactJson{"Bobby", AddressJson{"100 Main St", "San Diego", "CA", "92129", "US"}, "+18585553333", "asd@sd.co"}},
          "createSubaccount.json")

      case createTrunk:
        marshalJson(CreateTrunkJson{defaultAccountId, randomName, "SIP/1234@phone.com:5060", int32(60), int32(800)}, "createTrunk.json")

      case replaceDevice:
        marshalJson(ReplaceDeviceJson{defaultAccountId, int32(145392), randomName}, "replaceDevice.json")

      case replaceExtension:
        marshalJson(ReplaceExtensionJson{defaultAccountId, int32(1767963), randomName, "America/Los_Angeles", true, 111, true, "unlimited", 12344, "bobby McFerrin", true, "standard", "private", 619, true, true, "automated", "+18587741111", "+18587748888"},
          "replaceExtension.json")

      case replaceMenu:
        marshalJson(ReplaceMenuJson{defaultAccountId, int32(88519), randomName, false, 5}, "replaceMenu.json")

      case replacePhoneNumber:
        marshalJson(ReplacePhoneNumberJson{defaultAccountId, int32(2131572), "Robert", true, true, "Phone N", "business", "extension", "+18587740222"}, "replacePhoneNumber.json")

      case replaceQueue:
        marshalJson(ReplaceQueueJson{defaultAccountId, int32(22035), randomName, 60, "called_number", 10}, "replaceQueue.json")

      case replaceRoute:
        marshalJson(ReplaceRouteJson{defaultAccountId, int32(4705073), randomName, []RulesJson{RulesJson{[]ActionsJson{ActionsJson{"queue", QueueJson{22026, "61kkjklmin74"}}}}}},
          "replaceRoute.json")

      case replaceTrunk:
        marshalJson(ReplaceTrunkJson{defaultAccountId, int32(1515), randomName, "SIP/1234@phone.com:5060", int32(80), int32(800)}, "createTrunk.json")

    }
  }

  switch api := api.(type) {

  case swagger.MediaApi:

    switch (command) {

    case listMedia:

      return handle(api.ListAccountMedia(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getRecording:

      if (recordingId != 0) {
        id = recordingId
      }

      return handle(api.GetAccountMedia(accountId, id))
    }

  case swagger.MenusApi:

    switch (command) {

    case listMenus:

      return handle(api.ListAccountMenus(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getMenu:

      if (menuId != 0) {
        id = menuId
      }

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

    switch (command) {

    case listQueues:

      return handle(api.ListAccountQueues(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getQueue:

      if (queueId != 0) {
        id = queueId
      }

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

    switch (command) {

    case listRoutes:

      return handle(api.ListAccountRoutes(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getRoute:

      if (routeId != 0) {
        id = routeId
      }

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

    switch (command) {

    case listSchedules:

      return handle(api.ListAccountSchedules(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getSchedule:

      if (scheduleId != 0) {
        id = scheduleId
      }

      return handle(api.GetAccountSchedule(accountId, id))
    }

  case swagger.SmsApi:

    switch (command) {

    case listSms:

      return handle(api.ListAccountSms(accountId, filtersId, filterParams.filtersDirection, filterParams.filtersFrom, sortParams.sortId, sortParams.sortCreatedAt, limit, offset, fields))

    case getSms:

      if (smsId != "") {
        idString = smsId
      }

      return handle(api.GetAccountSms(accountId, idString))

    case createSms:

      params := createSmsParams(from,to,text, extensionId)
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

      params := createSubaccountParams(input, contact, billingContact)
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

      return handle(api.ListAvailablePhoneNumberRegions(filterParams.filtersCountryCode, filterParams.filtersNpa, filterParams.filtersNxx, filterParams.filtersIsTollFree, filterParams.filtersCity, filterParams.filtersProvincePostalCode, filterParams.filtersCountryPostalCode, sortParams.sortCountryCode, sortParams.sortNpa, sortParams.sortNxx, sortParams.sortIsTollFree, sortParams.sortCity, sortParams.sortProvincePostalCode, sortParams.sortCountryPostalCode, limit, offset, fields, groupBy))
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

    switch (command) {

    case listDevices:

      return handle(api.ListAccountDevices(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getDevice:

      if (deviceId != 0) {
        id = deviceId
      }

      return handle(api.GetAccountDevice(accountId, id))

    case createDevice:

      params := createDeviceParams(input)
      return handle(api.CreateAccountDevice(accountId, params))

    case replaceDevice:

      params := createDeviceParams(input)
      return handle(api.ReplaceAccountDevice(accountId, id, params))
    }

  case swagger.ExpressservicecodesApi:

    switch (command) {

    case listExpressServiceCodes:

      return handle(api.ListAccountExpressSrvCodes(accountId, slice))

    case getExpressServiceCode:

      if (codeId != 0) {
        id = codeId
      }

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

    switch (command) {

    case listContacts:

      return handle(api.ListAccountExtensionContacts(accountId, id, filtersId, filterParams.filtersGroupId, filterParams.filtersUpdatedAt, sortParams.sortId, sortParams.sortUpdatedAt, limit, offset, fields))

    case getContact:

      if (extensionId != 0) {
        id = extensionId
      }

      if (contactId != 0) {
        idSecondary = contactId
      }

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

    switch (command) {

    case listGroups:

      return handle(api.ListAccountExtensionContactGroups(accountId, id, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getGroup:

      if (extensionId != 0) {
        id = extensionId
      }

      if (groupId != 0) {
        idSecondary = groupId
      }

      return handle(api.GetAccountExtensionContactGroup(accountId, id, idSecondary))

    case createGroup:

      params := createGroupParams(input)
      return handle(api.CreateAccountExtensionContactGroup(accountId, id, params))

    case replaceGroup:

      //params := createGroupParams()
      return handle(api.ReplaceAccountExtensionContactGroup(accountId, id, idSecondary))
    case deleteGroup:

      return handle(api.DeleteAccountExtensionContactGroup(accountId, id, idSecondary))
    }

  case swagger.PhonenumbersApi:

    switch (command) {

    case listPhoneNumbers:

      return handle(api.ListAccountPhoneNumbers(accountId, filtersId, filterParams.filtersName, filterParams.filtersPhoneNumber, sortParams.sortId, sortParams.sortName, sortParams.sortPhoneNumber, limit, offset, fields))

    case getPhoneNumber:

      if (numberId != 0) {
        id = numberId
      }

      return handle(api.GetAccountPhoneNumber(accountId, id))

    case createPhoneNumber:

      params := createPhoneNumberParams(input)
      return handle(api.CreateAccountPhoneNumber(accountId, params))

    case replacePhoneNumber:

      params := replacePhoneNumberParams(input)
      return handle(api.ReplaceAccountPhoneNumber(accountId, id, params))
    }

  case swagger.TrunksApi:

    switch (command) {

    case listTrunks:

      return handle(api.ListAccountTrunks(accountId, filtersId, filterParams.filtersName, sortParams.sortId, sortParams.sortName, limit, offset, fields))

    case getTrunk:

      if (trunkId != 0) {
        id = trunkId
      }

      return handle(api.GetAccountTrunk(accountId, id))

    case createTrunk:
      params := createTrunkParams(trunkName, trunkUri, trunkConcurrentCalls, trunkMaxMinutes )
      return handle(api.CreateAccountTrunk(accountId, params))

    case replaceTrunk:

      params := createTrunkParams(trunkName, trunkUri, trunkConcurrentCalls, trunkMaxMinutes )
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
command string) (interface{}, Config) {

  var api interface{}
  var config = swagger.NewConfiguration()

  var xmlConfig Config = getConfig()
  var baseApiPath string = xmlConfig.BaseApiPath
  var apiKeyPrefix string = xmlConfig.ApiKeyPrefix
  var apiKey string = xmlConfig.ApiKey

  if (len(apiKeyPrefix) == 0 || len(apiKey) == 0) {
    return nil, xmlConfig
  }

  if (len(baseApiPath) > 0) {
    config.BasePath = baseApiPath
  }

  config.APIKeyPrefix["Authorization"] = apiKeyPrefix
  config.APIKey["Authorization"] = apiKey

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
    return nil, xmlConfig
  }

  return api, xmlConfig
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

  json := validateJson(string(payload))

  if (json == nil) {

    if (param.verbose) {
      fmt.Println("Could not unmarshal API json response")
    }

    return errors.New(msgInvalidJson), nil
  }

  message := validateResponse(json)

  if (message != "") {

    if (param.verbose) {
      fmt.Printf("%+v\n%s\n", x, response)
    }

    return errors.New(message), nil
  } else {
    fmt.Printf("%+v\n%s\n", x, response)
  }

  return nil, json
}

func marshalJson(param interface{}, fileName string) {

  b, err := json.Marshal(param)
  err = ioutil.WriteFile(fileName, b, 0644)
  _ = err
}