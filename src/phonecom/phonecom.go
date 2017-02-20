package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"phonecom-go-sdk"
)

const accountId = 1315091

func main() {

	app := cli.NewApp()

	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name: "command",
			Value: listAccounts,
			Usage: "Phone.com API command that you want to execute",
		},
		cli.IntFlag{
			Name: "id",
			Value: 0,
			Usage: "ID of entity you want to operate",
		},
		cli.IntFlag{
			Name: "limit",
			Value: 5,
			Usage: "Upper limit of results you want to fetch",
		},
		cli.IntFlag{
			Name: "offset",
			Value: 0,
			Usage: "Offset of results you want to fetch",
		},
		cli.StringFlag{
			Name: "dryrun",
			Value: "",
			Usage: "Print the expected action without executing the API command",
		},
	}

	app.Action = func(c *cli.Context) error {
		return execute(c)
	}

	app.Run(os.Args)
}

func execute(
    c *cli.Context) error {

	slice := make([]string, 0)
	limit := int32(c.Int("limit"))
	offset := int32(c.Int("offset"))
	id := int32(c.Int("id"))
	dryrun := c.String("dryrun")

	command := c.String("command")

	var api interface{} = getApi(command)
	if (api == nil) {
		return nil
	}

	switch api := api.(type) {

	case swagger.MediaApi:

		switch (command) {

		case listMedia:
			if(dryrun != ""){
				fmt.Println("Calling ", command, " with parameters: limit ", limit, " offset ", offset, " account ID ", accountId, " expecting a list of media in JSON format")

				return nil
			}
			handle(api.ListAccountMedia(accountId, slice, slice, "", "", limit, offset, ""))
		case getRecording:
			if(dryrun != ""){
				fmt.Println("Calling ", command, " with parameters: account ID ", accountId,  " id of recording ", id, " expecting a recording in JSON format")

				return nil
			}
			handle(api.GetAccountMedia(accountId, id))
		}

	case swagger.MenusApi:

		switch (command) {
		case listMenus:
			if(dryrun != ""){
				fmt.Println("Calling ", command, " with parameters: limit ", limit, " offset ", offset, " account ID ", accountId, " expecting a list of menus in JSON format")

				return nil
			}
			handle(api.ListAccountMenus(accountId, slice, slice, "", "", limit, offset, ""))
		case getMenu:
			if(dryrun != ""){
				fmt.Println("Calling ", command, " with parameters: account ID ", accountId,  " id of media ", id, " expecting a recording in JSON format")

				return nil
			}
			handle(api.GetAccountMenu(accountId, id))
		case createMenu:
			var params = createMenuParams()
			handle(api.CreateAccountMenu(accountId,params))
		case replaceMenu:
			var params = replaceMenuParams()
			handle(api.ReplaceAccountMenu(accountId,88295,params))
		case deleteMenu:
			handle(api.DeleteAccountMenu(accountId,88295))

		}

	case swagger.QueuesApi:

		switch (command) {
		case listQueues:
			if(dryrun != ""){
				fmt.Println("Calling ", command, " with parameters: limit ", limit, " offset ", offset, " account ID ", accountId, " expecting a list of queues in JSON format")

				return nil
			}
			handle(api.ListAccountQueues(accountId, slice, slice, "", "", limit, offset, ""))
		case getQueue:
			if(dryrun != ""){
				fmt.Println("Calling ", command, " with parameters: account ID ", accountId,  " id of media ", id, " expecting a recording in JSON format")

				return nil
			}
			handle(api.GetAccountQueue(accountId, id))
		case createQueue:
			var params = createQueueParams()
			handle(api.CreateAccountQueue(accountId, params))
		case replaceQueue:
			var params = createQueueParams()
			handle(api.ReplaceAccountQueue(accountId, 141494, params))
		case deleteQueue:
			handle(api.DeleteAccountQueue(accountId, 141494))
		}

	case swagger.RoutesApi:

		switch (command) {
			case listRoutes:
				handle(api.ListAccountRoutes(accountId, slice, slice, "", "", limit, offset, ""))
			case getRoute:
				handle(api.GetAccountRoute(accountId, id))
			case createRoute:
				var params = createRouteParams()
				handle(api.CreateRoute(accountId, params))
			case replaceRoute:
				var params = createRouteParams()
				handle(api.ReplaceAccountRoute(accountId, id, params))
			case deleteRoute:
				handle(api.DeleteAccountRoute(accountId, 14987))
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
                    var params = createSmsParams()
                    handle(api.CreateAccountSms(accountId, params))
			}

		case swagger.AvailablenumbersApi:

			switch (command) {
				case listAvailablePhoneNumbers:
					handle(api.ListAvailablePhoneNumbers(slice, slice, slice, slice, slice, slice, slice, slice, slice, slice, slice, "", "", "", limit, offset, ""))
			}

		case swagger.SubaccountsApi:

			switch (command) {
				case listSubaccounts:
					handle(api.ListAccountSubaccounts(accountId, slice, "", limit, offset, ""))
                case createSubaccount:
                    var params = createSubaccountParams()
                    handle(api.CreateAccountSubaccount(accountId, params))
			}

		case swagger.AccountsApi:

			switch (command) {

				case listAccounts:
					handle(api.ListAccounts(slice, "", limit, offset, ""))
				case getAccount:
					handle(api.GetAccount(accountId))
			}

		case swagger.ApplicationsApi:

			switch (command) {

				case listApplications:
					handle(api.ListAccountApplications(accountId, slice, slice, "", "", limit, offset, ""))
				case getApplication:
					handle(api.GetAccountApplication(accountId, 1356077))
			}

		case swagger.CallsApi:

			switch (command) {
				case createCall:
					var params = createCallParams()
					handle(api.CreateAccountCalls(accountId, params))
			}

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
					var params = createDeviceParams()
					handle(api.CreateAccountDevice(accountId, params))
				case replaceDevice:
					var params = createDeviceParams()
					handle(api.ReplaceAccountDevice(accountId, id, params))
			}

		case swagger.ExpressservicecodesApi:

			switch (command) {

				case listExpressServiceCodes:
					handle(api.ListAccountExpressSrvCodes(accountId, slice))
				case getExpressServiceCode:
					handle(api.GetAccountExpressSrvCode(accountId, 324202))
			}

		case swagger.ExtensionsApi:

			switch (command) {

				case listExtensions:
					handle(api.ListAccountExtensions(accountId, slice, slice, slice, "", "", "", limit, offset, ""))
				case getExtension:
					handle(api.GetAccountExtension(accountId, 1767129))
				case createExtension:
					var params = createExtensionsParams()
					handle(api.CreateAccountExtension(accountId, params))
				case replaceExtension:
					var params = replaceExtensionParams()
					handle(api.ReplaceAccountExtension(accountId, 1767129, params))
			}

		case swagger.CalleridsApi:

			switch (command) {

				case getCallerId:
					handle(api.GetCallerIds(accountId, 1764590, slice, slice, "", "", limit, offset, ""))
			}

		case swagger.ContactsApi:

			switch (command) {

				case listContacts:
					handle(api.ListAccountExtensionContacts(accountId, 1764590, slice, slice, slice, "", "", limit, offset, ""))
				case getContact:
					handle(api.GetAccountExtensionContact(accountId, 1764590, 2074702))
				case createContact:
					var params = createContactParams()
					handle(api.CreateAccountExtensionContact(accountId, 1764590, params))
				case replaceContact:
					var params = createContactParams()
					handle(api.ReplaceAccountExtensionContact(accountId, 1764590, params))
				case deleteContact:
					handle(api.DeleteAccountExtensionContact(accountId, 1764590, 2072969))
			}

		case swagger.GroupsApi:

			switch (command) {

				case listGroups:
					handle(api.ListAccountExtensionContactGroups(accountId, 1764590, slice, slice, "", "", limit, offset, ""))
				case getGroup:
					handle(api.GetAccountExtensionContactGroup(accountId, 1764590, 331026))
				case createGroup:
					var params = createGroupParams()
					handle(api.CreateAccountExtensionContactGroup(accountId, 1764590, params))
				case replaceGroup:
					//var params = createGroupParams()
					handle(api.ReplaceAccountExtensionContactGroup(accountId, 1764590, 331978))
				case deleteGroup:
					handle(api.DeleteAccountExtensionContactGroup(accountId, 1764590, 331978))
			}

		case swagger.PhonenumbersApi:

			switch (command) {

				case listPhoneNumbers:
					handle(api.ListAccountPhoneNumbers(accountId, slice, slice, slice, "", "", "", limit, offset, ""))
				case getPhoneNumber:
					handle(api.GetAccountPhoneNumber(accountId, 2116986))
				case createPhoneNumber:
					var params = createPhoneNumberParams()
					handle(api.CreateAccountPhoneNumber(accountId, params))
				case replacePhoneNumber:
					var params = replacePhoneNumberParams()
					handle(api.ReplaceAccountPhoneNumber(accountId, 2116986, params))
			}

        case swagger.TrunksApi:

			switch (command) {

				case listTrunks:
					handle(api.ListAccountTrunks(accountId, slice, slice, "", "", limit, offset, ""))
				case getTrunk:
					handle(api.GetAccountTrunk(accountId, 2116986))
				case createTrunk:
					var params = createTrunkParams()
					handle(api.CreateAccountTrunk(accountId, params))
				case replaceTrunk:
					var params = createTrunkParams()
					handle(api.ReplaceAccountTrunk(accountId, 2116986, params))
		    case deleteTrunk:
          handle(api.DeleteAccountTrunk(accountId, 2116986))
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

    case createCall:

      callsApi := *swagger.NewCallsApi()
      callsApi.Configuration = config
      api = callsApi

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
}
