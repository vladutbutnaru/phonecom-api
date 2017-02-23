package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "phonecom-go-sdk"
  "encoding/json"
)

func main() {

  app := cli.NewApp()

  app.Flags = []cli.Flag{

    cli.StringFlag{
      Name: "command, c",
      Value: listAccounts,
      Usage: "Phone.com API command that you want to execute",
    },
    cli.IntFlag{
      Name: "id",
      Value: 0,
      Usage: "ID of entity you want to operate",
    },
    cli.IntFlag{
      Name: "id-secondary, is",
      Value: 0,
      Usage: "Secondary ID of entity you want to operate",
    },
    cli.IntFlag{
      Name: "limit, l",
      Value: 5,
      Usage: "Upper limit of results you want to fetch",
    },
    cli.IntFlag{
      Name: "offset, o",
      Value: 0,
      Usage: "Offset of results you want to fetch",
    },
    cli.IntFlag{
      Name: "account, a",
      Value: 1315091,
      Usage: "Phone.com API account to use in API calls",
    },
    cli.StringFlag{
      Name: "dryrun, d",
      Value: "",
      Usage: "Print the expected action without executing the API command",
    },
		cli.StringFlag{
			Name: "input, i",
			Value: "",
			Usage: "Specify the path to the JSON file for making the API call",
		},
		cli.StringFlag{
			Name: "verbose, vr",
			Value: "",
			Usage: "Activate verbose mode",
		},

  }

  app.Action = func(c *cli.Context) error {
    return execute(c)
  }

  app.Run(os.Args)
}

type CliParams struct {
	slice       []string
	limit       int32
	offset      int32
	id          int32
	dryRun      string
	verbose     string
	input       string
	command     string
	idSecondary int32
	accountId   int32
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

				handle(api.ListAccountMedia(accountId, slice, slice, "", "", limit, offset, ""))

			case getRecording:

				handle(api.GetAccountMedia(accountId, id))
    }

  case swagger.MenusApi:

    switch (command) {

			case listMenus:

				handle(api.ListAccountMenus(accountId, slice, slice, "", "", limit, offset, ""))

			case getMenu:

				handle(api.GetAccountMenu(accountId, id))

			case createMenu:

				var params = createMenuParams(input)
				handle(api.CreateAccountMenu(accountId, params))

			case replaceMenu:

				var params = replaceMenuParams(input)
				handle(api.ReplaceAccountMenu(accountId, id, params))

			case deleteMenu:

				handle(api.DeleteAccountMenu(accountId, id))
		}

  case swagger.QueuesApi:

    switch (command) {

			case listQueues:

				handle(api.ListAccountQueues(accountId, slice, slice, "", "", limit, offset, ""))

			case getQueue:

				handle(api.GetAccountQueue(accountId, id))

			case createQueue:

				var params = createQueueParams(input)
				handle(api.CreateAccountQueue(accountId, params))

			case replaceQueue:

				var params = createQueueParams(input)
				handle(api.ReplaceAccountQueue(accountId, id, params))

			case deleteQueue:

				handle(api.DeleteAccountQueue(accountId, id))
			}

  case swagger.RoutesApi:

    switch (command) {

			case listRoutes:

				handle(api.ListAccountRoutes(accountId, slice, slice, "", "", limit, offset, ""))

			case getRoute:

				handle(api.GetAccountRoute(accountId, id))

			case createRoute:

				var params = createRouteParams(input)
				handle(api.CreateRoute(accountId, params))

			case replaceRoute:

				var params = createRouteParams(input)
				handle(api.ReplaceAccountRoute(accountId, id, params))

			case deleteRoute:

				handle(api.DeleteAccountRoute(accountId, id))
    }

  case swagger.SchedulesApi:

    switch (command) {

			case listSchedules:

				handle(api.ListAccountSchedules(accountId, slice, slice, "", "", limit, offset, ""))

			case getSchedule:

				handle(api.GetAccountSchedule(accountId, id))
    }

  case swagger.SmsApi:

    switch (command) {

			case listSms:

				handle(api.ListAccountSms(accountId, slice, "", "", "", "", limit, offset, ""))

			case getSms:

				handle(api.GetAccountSms(accountId, id))

			case createSms:

				var params = createSmsParams(input)
				handle(api.CreateAccountSms(accountId, params))
    }

  case swagger.AvailablenumbersApi:

    switch (command) {

			case listAvailablePhoneNumbers:

				handle(api.ListAvailablePhoneNumbers(slice, slice, slice, slice, slice, slice, slice, slice, slice, slice, "", "", "", limit, offset, ""))
			}

  case swagger.SubaccountsApi:

    switch (command) {

			case listSubaccounts:

				handle(api.ListAccountSubaccounts(accountId, slice, "", limit, offset, ""))

			case createSubaccount:

				var params = createSubaccountParams(input)
				handle(api.CreateAccountSubaccount(accountId, params))
    }

  case swagger.AccountsApi:

    switch (command) {

			case listAccounts:

				handle(api.ListAccounts(slice, "", limit, offset, ""))

			case getAccount:

				handle(api.GetAccount(id))
    }

  case swagger.NumberregionsApi:

    switch (command) {

			case listAvailablePhoneNumberRegions:

				handle(api.ListAvailablePhoneNumberRegions(slice, slice, slice, slice, slice, slice, slice, "", "", "", "", "", "", "", limit, offset, "", slice))
    }

  case swagger.ApplicationsApi:

    switch (command) {

			case listApplications:

				handle(api.ListAccountApplications(accountId, slice, slice, "", "", limit, offset, ""))

			case getApplication:

				handle(api.GetAccountApplication(accountId, id))
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
  //			handle(api.CreateAccountCalls(accountId, params))
  //	}

  case swagger.CalllogsApi:

    switch (command) {

			case listCallLogs:

				handle(api.ListAccountCallLogs(accountId, slice, slice, "", "", "", "", slice, "", "", "", limit, offset, ""))
			//~ case getCallLog:
			//~ handle(api.GetAccountCallLog(accountId, id))
    }

  case swagger.DevicesApi:

    switch (command) {

			case listDevices:

				handle(api.ListAccountDevices(accountId, slice, slice, "", "", limit, offset, ""))

			case getDevice:

				handle(api.GetAccountDevice(accountId, id))

			case createDevice:

				var params = createDeviceParams(input)
				handle(api.CreateAccountDevice(accountId, params))

			case replaceDevice:

				var params = createDeviceParams(input)
				handle(api.ReplaceAccountDevice(accountId, id, params))
    }

  case swagger.ExpressservicecodesApi:

    switch (command) {

			case listExpressServiceCodes:

				handle(api.ListAccountExpressSrvCodes(accountId, slice))

			case getExpressServiceCode:

				handle(api.GetAccountExpressSrvCode(accountId, id))
    }

  case swagger.ExtensionsApi:

    switch (command) {

			case listExtensions:

				handle(api.ListAccountExtensions(accountId, slice, slice, slice, "", "", "", limit, offset, ""))

			case getExtension:

				handle(api.GetAccountExtension(accountId, id))

			case createExtension:

				var params = createExtensionsParams(input)
				handle(api.CreateAccountExtension(accountId, params))

			case replaceExtension:

				var params = replaceExtensionParams(input)
				handle(api.ReplaceAccountExtension(accountId, id, params))

    }

  case swagger.CalleridsApi:

    switch (command) {

			case getCallerId:
				handle(api.GetCallerIds(accountId, id, slice, slice, "", "", limit, offset, ""))
    }

  case swagger.ContactsApi:

    switch (command) {

			case listContacts:

				handle(api.ListAccountExtensionContacts(accountId, id, slice, slice, slice, "", "", limit, offset, ""))

			case getContact:

				handle(api.GetAccountExtensionContact(accountId, id, idSecondary))

			case createContact:

				var params = createContactParams(input)
				handle(api.CreateAccountExtensionContact(accountId, id, params))

			case replaceContact:

				var params = createContactParams(input)
				handle(api.ReplaceAccountExtensionContact(accountId, id, params))

			case deleteContact:

				handle(api.DeleteAccountExtensionContact(accountId, id, idSecondary))
    }

  case swagger.GroupsApi:

    switch (command) {

			case listGroups:

				handle(api.ListAccountExtensionContactGroups(accountId, id, slice, slice, "", "", limit, offset, ""))

			case getGroup:

				handle(api.GetAccountExtensionContactGroup(accountId, id, idSecondary))

			case createGroup:

				var params = createGroupParams(input)
				handle(api.CreateAccountExtensionContactGroup(accountId, id, params))

			case replaceGroup:

				//var params = createGroupParams()
				handle(api.ReplaceAccountExtensionContactGroup(accountId, id, idSecondary))
			case deleteGroup:

				handle(api.DeleteAccountExtensionContactGroup(accountId, id, idSecondary))
    }

  case swagger.PhonenumbersApi:

    switch (command) {

			case listPhoneNumbers:

				handle(api.ListAccountPhoneNumbers(accountId, slice, slice, slice, "", "", "", limit, offset, ""))

			case getPhoneNumber:

				handle(api.GetAccountPhoneNumber(accountId, id))

			case createPhoneNumber:

				var params = createPhoneNumberParams(input)
				handle(api.CreateAccountPhoneNumber(accountId, params))

			case replacePhoneNumber:

				var params = replacePhoneNumberParams(input)
				handle(api.ReplaceAccountPhoneNumber(accountId, id, params))
    }

  case swagger.TrunksApi:

    switch (command) {

    case listTrunks:

      handle(api.ListAccountTrunks(accountId, slice, slice, "", "", limit, offset, ""))

    case getTrunk:

      handle(api.GetAccountTrunk(accountId, id))

    case createTrunk:

      var params = createTrunkParams(input)
      handle(api.CreateAccountTrunk(accountId, params))

    case replaceTrunk:

      var params = createTrunkParams(input)
      handle(api.ReplaceAccountTrunk(accountId, id, params))

    case deleteTrunk:

      handle(api.DeleteAccountTrunk(accountId, id))
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
    error error) {

  if (error != nil) {
    panic(error)
  }

  fmt.Printf("%+v\n%s\n", x, response)

  validateJson(string(response.Payload))
}

func validateJson(jsonString string) {

  fmt.Println();

  var message string = "JSON validator: Invalid JSON";

  if isJSON(string(jsonString)) {
    message = "JSON validator: Valid JSON";
  }

  fmt.Println(message);
}

func isJSON(jsonString string) bool {

  var js map[string] interface{}

  return json.Unmarshal([]byte(jsonString), &js) == nil
}
