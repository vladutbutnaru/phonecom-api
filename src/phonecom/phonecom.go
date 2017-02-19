package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "phonecom-go-sdk"
	"encoding/xml"
	"io/ioutil"
)

const accountId = 1315091


type Query struct {
	ConfigList []Config `xml:"Config"`
}

type Config struct {
	ApiKeyPrefix string
	ApiKey string
	Type string
}

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
  }

  app.Action = func(c *cli.Context) error {

		var config Config = getConfig()
		var apiKeyPrefix string = config.ApiKeyPrefix
		var apiKey string = config.ApiKey

		if (len(apiKeyPrefix) == 0 || len(apiKey) == 0) {
			return nil
		}

		return execute(c, apiKeyPrefix, apiKey)
  }

  app.Run(os.Args)
}

func getConfig() Config {

	xmlFile, err := os.Open("config.xml")

	var noConfig Config;

	if err != nil {
		fmt.Println("Could not read config.xml", err)
		return noConfig
	}

	defer xmlFile.Close()

	content, _ := ioutil.ReadAll(xmlFile)

	var q Query
	xml.Unmarshal(content, &q)

	for _, config := range q.ConfigList {
		if (config.Type == "main") {
			return config
		}
	}

	return noConfig
}

func execute(
    c *cli.Context,
    apiKeyPrefix string,
    apiKey string) error {

	slice := make([]string, 0)
	limit := int32(c.Int("limit"))
	offset := int32(c.Int("offset"))
	id := int32(c.Int("id"))

	command := c.String("command")

	var api interface{} = getApi(command, apiKeyPrefix, apiKey)

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
			}

		case swagger.QueuesApi:

			switch (command) {
				case listQueues:
					handle(api.ListAccountQueues(accountId, slice, slice, "", "", limit, offset, ""))
				case getQueue:
					handle(api.GetAccountQueue(accountId, id))
			}

		case swagger.RoutesApi:

			switch (command) {
				case listRoutes:
					handle(api.ListAccountRoutes(accountId, slice, slice, "", "", limit, offset, ""))
				case getRoute:
					handle(api.GetAccountRoute(accountId, id))
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
			}

		case swagger.DevicesApi:

			switch (command) {

				case listDevices:
					handle(api.ListAccountDevices(accountId, slice, slice, "", "", limit, offset, ""))
				case getDevice:
					handle(api.GetAccountDevice(accountId, 144510))
				case createDevice:
					var params = createDeviceParams()
					handle(api.CreateAccountDevice(accountId, params))
				case replaceDevice:
					var params = createDeviceParams()
					handle(api.ReplaceAccountDevice(accountId, 142315, params))
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

		default:
			return nil
	}

	return nil
}

func getApi(
    command string,
    apiKeyPrefix string,
    apiKey string) interface{} {

  var api interface{}
  var config = swagger.NewConfiguration()

  config.APIKeyPrefix["Authorization"] = apiKeyPrefix
  config.APIKey["Authorization"] = apiKey

  switch (command) {

    case listMedia, getRecording:

      mediaApi := *swagger.NewMediaApi()
      mediaApi.Configuration = config
      api = mediaApi

    case listMenus, getMenu:

      menusApi := *swagger.NewMenusApi()
      menusApi.Configuration = config
      api = menusApi

    case listQueues, getQueue:

      queuesApi := *swagger.NewQueuesApi()
      queuesApi.Configuration = config
      api = queuesApi

    case listRoutes, getRoute:

      routesApi := *swagger.NewRoutesApi()
      routesApi.Configuration = config
      api = routesApi

    case listSchedules, getSchedule:

      schedulesApi := *swagger.NewSchedulesApi()
      schedulesApi.Configuration = config
      api = schedulesApi

    case listSms, getSms:

      smsApi := *swagger.NewSmsApi()
      smsApi.Configuration = config
      api = smsApi

    case listAvailablePhoneNumbers:

      availablenumbersApi := *swagger.NewAvailablenumbersApi()
      availablenumbersApi.Configuration = config
      api = availablenumbersApi

    case listSubaccounts:

      subaccountsApi := *swagger.NewSubaccountsApi()
      subaccountsApi.Configuration = config
      api = subaccountsApi

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

      calllogsApi := *swagger.NewCalllogsApi()
      calllogsApi.Configuration = config
      api = calllogsApi

    case listDevices, getDevice, createDevice, replaceDevice:

      devicesApi := *swagger.NewDevicesApi()
      devicesApi.Configuration = config
      api = devicesApi

    case listExpressServiceCodes, getExpressServiceCode:

      expressservicecodesApi := *swagger.NewExpressservicecodesApi()
      expressservicecodesApi.Configuration = config
      api = expressservicecodesApi

    case listExtensions, getExtension, createExtension, replaceExtension:

      extensionsApi := *swagger.NewExtensionsApi()
      extensionsApi.Configuration = config
      api = extensionsApi

    case getCallerId:

      calleridsApi := *swagger.NewCalleridsApi()
      calleridsApi.Configuration = config
      api = calleridsApi

    case listContacts, getContact, createContact, replaceContact, deleteContact:

      contactsApi := *swagger.NewContactsApi()
      contactsApi.Configuration = config
      api = contactsApi

    case listGroups, getGroup, createGroup, replaceGroup, deleteGroup:

      groupsApi := *swagger.NewGroupsApi()
      groupsApi.Configuration = config
      api = groupsApi

    case listPhoneNumbers, getPhoneNumber, createPhoneNumber, replacePhoneNumber:

      phonenumbersApi := *swagger.NewPhonenumbersApi()
      phonenumbersApi.Configuration = config
      api = phonenumbersApi

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
