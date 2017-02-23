package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "phonecom-go-sdk"
  "errors"
)

var configPath = "config.xml"

func main() {

  app := cli.NewApp()

  app.Flags = getCliFlags()

  app.Action = func(c *cli.Context) error {
    return execute(c)
  }

  app.Run(os.Args)
}

type CliParams struct {
	slice       []string
    filtersId   []string
	limit       int32
	offset      int32
	id          int32
	dryRun      string
	verbose     string
	input       string
	command     string
	idSecondary int32
	accountId   int32
    fields      string
}

func execute(
    c *cli.Context) error {

  var param CliParams

  var slice = make([]string, 0)
  var limit = int32(c.Int("limit"))
  var offset = int32(c.Int("offset"))
  var id = int32(c.Int("id"))
	var verbose = c.String("verbose")
	var dryRun = c.String("dryrun")
  var input = c.String("input")
  var command = c.String("command")
  var idSecondary = int32(c.Int("id-secondary"))
  var accountId = int32(c.Int("account"))
  var fields string = ""
    var filtersId []string
    
  if getAccountIdFromFile(input) > 0 {
        accountId = getAccountIdFromFile(input)
        param.accountId = accountId
        }
    if getLimitFromFile(input) > 0 {
        limit = getLimitFromFile(input)
        
    }
    if getOffsetFromFile(input) > 0 {
          offset = getOffsetFromFile(input)
        
    }
  
    if  getFieldsFromFile(input) != "" {
          fields = getFieldsFromFile(input)
        
    }
    
    filtersId = getFiltersIdFromFile(input)
    
	param.slice = slice
	param.limit = limit
	param.offset = offset
	param.id = id
	param.dryRun = dryRun
	param.verbose = verbose
	param.input = input
	param.command = command
	param.idSecondary = idSecondary
	param.accountId = accountId
    param.fields = fields
    

	var api interface{} = getApi(command)
  if (api == nil) {
    return nil
  }

  if (showDryRunVerbose(param) != "continue") {
    return nil;
  }

   
  switch api := api.(type) {
 
  case swagger.MediaApi:

    switch (command) {

			case listMedia:
       
            return handle(api.ListAccountMedia(accountId,filtersId , slice, "", "", limit, offset, fields))

			case getRecording:

        return handle(api.GetAccountMedia(accountId, id))
    }

  case swagger.MenusApi:

    switch (command) {

			case listMenus:

        return handle(api.ListAccountMenus(accountId, slice, slice, "", "", limit, offset, fields))

			case getMenu:

				return handle(api.GetAccountMenu(accountId, id))

			case createMenu:

				var params = createMenuParams(input)
				return handle(api.CreateAccountMenu(accountId, params))

			case replaceMenu:

				var params = replaceMenuParams(input)
				return handle(api.ReplaceAccountMenu(accountId, id, params))

			case deleteMenu:

				return handle(api.DeleteAccountMenu(accountId, id))
		}

  case swagger.QueuesApi:

    switch (command) {

			case listQueues:

				return handle(api.ListAccountQueues(accountId, slice, slice, "", "", limit, offset, fields))

			case getQueue:

				return handle(api.GetAccountQueue(accountId, id))

			case createQueue:

				var params = createQueueParams(input)
				return handle(api.CreateAccountQueue(accountId, params))

			case replaceQueue:

				var params = createQueueParams(input)
				return handle(api.ReplaceAccountQueue(accountId, id, params))

			case deleteQueue:

				return handle(api.DeleteAccountQueue(accountId, id))
			}

  case swagger.RoutesApi:

    switch (command) {

			case listRoutes:

				return handle(api.ListAccountRoutes(accountId, slice, slice, "", "", limit, offset, fields))

			case getRoute:

				return handle(api.GetAccountRoute(accountId, id))

			case createRoute:

				var params = createRouteParams(input)
				return handle(api.CreateRoute(accountId, params))

			case replaceRoute:

				var params = createRouteParams(input)
				return handle(api.ReplaceAccountRoute(accountId, id, params))

			case deleteRoute:

				return handle(api.DeleteAccountRoute(accountId, id))
    }

  case swagger.SchedulesApi:

    switch (command) {

			case listSchedules:

				return handle(api.ListAccountSchedules(accountId, slice, slice, "", "", limit, offset, fields))

			case getSchedule:

				return handle(api.GetAccountSchedule(accountId, id))
    }

  case swagger.SmsApi:

    switch (command) {

			case listSms:

				return handle(api.ListAccountSms(accountId, slice, "", "", "", "", limit, offset, fields))

			case getSms:

				return handle(api.GetAccountSms(accountId, id))

			case createSms:

				var params = createSmsParams(input)
				return handle(api.CreateAccountSms(accountId, params))
    }

  case swagger.AvailablenumbersApi:

    switch (command) {

			case listAvailablePhoneNumbers:

				//return handle(api.ListAvailablePhoneNumbers(slice, slice, slice, slice, slice, slice, slice, slice, slice, slice, "", "", "", limit, offset, ""))
			}

  case swagger.SubaccountsApi:

    switch (command) {

			case listSubaccounts:

				return handle(api.ListAccountSubaccounts(accountId, slice, "", limit, offset, fields))

			case createSubaccount:

				var params = createSubaccountParams(input)
				return handle(api.CreateAccountSubaccount(accountId, params))
    }

  case swagger.AccountsApi:

    switch (command) {

			case listAccounts:

				return handle(api.ListAccounts(slice, "", limit, offset, fields))

			case getAccount:

				return handle(api.GetAccount(id))
    }

  case swagger.NumberregionsApi:

    switch (command) {

			case listAvailablePhoneNumberRegions:

				return handle(api.ListAvailablePhoneNumberRegions(slice, slice, slice, slice, slice, slice, slice, "", "", "", "", "", "", "", limit, offset, fields, slice))
    }

  case swagger.ApplicationsApi:

    switch (command) {

			case listApplications:

				return handle(api.ListAccountApplications(accountId, slice, slice, "", "", limit, offset, fields))

			case getApplication:

				return handle(api.GetAccountApplication(accountId, id))
    }

  //case swagger.CallsApi:
  //
  //	switch (command) {
  //		case createCall:
  //           if(dryrun != "" || verbose !=""){
  //		fmt.Println("Calling ", command, " expecting to create an call object with parameters from file ", input)
  //
  //		 if(verbose=="") { return nil }
  //	}
  //			var params = createCallParams(input)
  //			return handle(api.CreateAccountCalls(accountId, params))
  //	}

  case swagger.CalllogsApi:

    switch (command) {

			case listCallLogs:

				return handle(api.ListAccountCallLogs(accountId, slice, slice, "", "", "", "", slice, "", "", "", limit, offset, fields))
			//~ case getCallLog:
			//~ return handle(api.GetAccountCallLog(accountId, id))
    }

  case swagger.DevicesApi:

    switch (command) {

			case listDevices:

				return handle(api.ListAccountDevices(accountId, slice, slice, "", "", limit, offset, fields))

			case getDevice:

				return handle(api.GetAccountDevice(accountId, id))

			case createDevice:

				var params = createDeviceParams(input)
				return handle(api.CreateAccountDevice(accountId, params))

			case replaceDevice:

				var params = createDeviceParams(input)
				return handle(api.ReplaceAccountDevice(accountId, id, params))
    }

  case swagger.ExpressservicecodesApi:

    switch (command) {

			case listExpressServiceCodes:

				return handle(api.ListAccountExpressSrvCodes(accountId, slice))

			case getExpressServiceCode:

				return handle(api.GetAccountExpressSrvCode(accountId, id))
    }

  case swagger.ExtensionsApi:

    switch (command) {

			case listExtensions:

				return handle(api.ListAccountExtensions(accountId, slice, slice, slice, "", "", "", limit, offset, fields))

			case getExtension:

				return handle(api.GetAccountExtension(accountId, id))

			case createExtension:

				var params = createExtensionsParams(input)
				return handle(api.CreateAccountExtension(accountId, params))

			case replaceExtension:

				var params = replaceExtensionParams(input)
				return handle(api.ReplaceAccountExtension(accountId, id, params))

    }

  case swagger.CalleridsApi:

    switch (command) {

			case getCallerId:
				return handle(api.GetCallerIds(accountId, id, slice, slice, "", "", limit, offset, fields))
    }

  case swagger.ContactsApi:

    switch (command) {

			case listContacts:

				return handle(api.ListAccountExtensionContacts(accountId, id, slice, slice, slice, "", "", limit, offset, fields))

			case getContact:

				return handle(api.GetAccountExtensionContact(accountId, id, idSecondary))

			case createContact:

				var params = createContactParams(input)
				return handle(api.CreateAccountExtensionContact(accountId, id, params))

			case replaceContact:

				var params = createContactParams(input)
				return handle(api.ReplaceAccountExtensionContact(accountId, id, params))

			case deleteContact:

				return handle(api.DeleteAccountExtensionContact(accountId, id, idSecondary))
    }

  case swagger.GroupsApi:

    switch (command) {

			case listGroups:

				return handle(api.ListAccountExtensionContactGroups(accountId, id, slice, slice, "", "", limit, offset, fields))

			case getGroup:

				return handle(api.GetAccountExtensionContactGroup(accountId, id, idSecondary))

			case createGroup:

				var params = createGroupParams(input)
				return handle(api.CreateAccountExtensionContactGroup(accountId, id, params))

			case replaceGroup:

				//var params = createGroupParams()
				return handle(api.ReplaceAccountExtensionContactGroup(accountId, id, idSecondary))
			case deleteGroup:

				return handle(api.DeleteAccountExtensionContactGroup(accountId, id, idSecondary))
    }

  case swagger.PhonenumbersApi:

    switch (command) {

			case listPhoneNumbers:

				return handle(api.ListAccountPhoneNumbers(accountId, slice, slice, slice, "", "", "", limit, offset, fields))

			case getPhoneNumber:

				return handle(api.GetAccountPhoneNumber(accountId, id))

			case createPhoneNumber:

				var params = createPhoneNumberParams(input)
				return handle(api.CreateAccountPhoneNumber(accountId, params))

			case replacePhoneNumber:

				var params = replacePhoneNumberParams(input)
				return handle(api.ReplaceAccountPhoneNumber(accountId, id, params))
    }

  case swagger.TrunksApi:

    switch (command) {

    case listTrunks:

      return handle(api.ListAccountTrunks(accountId, slice, slice, "", "", limit, offset, fields))

    case getTrunk:

      return handle(api.GetAccountTrunk(accountId, id))

    case createTrunk:

      var params = createTrunkParams(input)
      return handle(api.CreateAccountTrunk(accountId, params))

    case replaceTrunk:

      var params = createTrunkParams(input)
      return handle(api.ReplaceAccountTrunk(accountId, id, params))

    case deleteTrunk:

      return handle(api.DeleteAccountTrunk(accountId, id))
    }

  default:
    return nil
  }

  return nil
}

func getApi(
    command string) interface{} {

  var api interface{}
  var config = swagger.NewConfiguration()

  var xmlConfig Config = getConfig()
  var baseApiPath string = xmlConfig.BaseApiPath
  var apiKeyPrefix string = xmlConfig.ApiKeyPrefix
  var apiKey string = xmlConfig.ApiKey

  if (len(apiKeyPrefix) == 0 || len(apiKey) == 0) {
    return nil
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

  //case createCall:
  //
  //  callsApi := *swagger.NewCallsApi()
  //  callsApi.Configuration = config
  //  api = callsApi

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
    fmt.Println("Invalid command:", command)
    return nil
  }

  return api
}

func handle(
    x interface{},
    response *swagger.APIResponse,
    error error) error {

  if (error != nil) {
    panic(error)
  }

  fmt.Printf("%+v\n%s\n", x, response)

  if (validateJson(string(response.Payload)) == nil) {
    return errors.New("Invalid json")
  }

  return nil
}
