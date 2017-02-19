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
      Usage: "API command that you want to execute",
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

		slice := make([]string, 0)
		limit := int32(c.Int("limit"))
		offset := int32(c.Int("offset"))
    id := int32(c.Int("id"))

    var mediaApi swagger.MediaApi
    var menusApi swagger.MenusApi
    var queuesApi swagger.QueuesApi
    var routesApi swagger.RoutesApi
    var schedulesApi swagger.SchedulesApi
    var smsApi swagger.SmsApi
    var availableNumbersApi swagger.AvailablenumbersApi
    var subAccountsApi swagger.SubaccountsApi
    var accountsApi swagger.AccountsApi
    var applicationsApi swagger.ApplicationsApi
    var callsApi swagger.CallsApi
    var callLogsApi swagger.CalllogsApi
    var devicesApi swagger.DevicesApi
    var expressServiceCodesApi swagger.ExpressservicecodesApi
    var extensionsApi swagger.ExtensionsApi
    var callerIdsApi swagger.CalleridsApi
    var contactsApi swagger.ContactsApi
    var groupsApi swagger.GroupsApi
    var phoneNumbersApi swagger.PhonenumbersApi

    command := c.String("command")

    switch (command) {

      case listMedia, getRecording:
        mediaApi = *swagger.NewMediaApi()
      case listMenus, getMenu:
        menusApi = *swagger.NewMenusApi()
      case listQueues, getQueue:
        queuesApi = *swagger.NewQueuesApi()
      case listRoutes, getRoute:
        routesApi = *swagger.NewRoutesApi()
      case listSchedules, getSchedule:
        schedulesApi = *swagger.NewSchedulesApi()
      case listSms, getSms:
        smsApi = *swagger.NewSmsApi()
      case listAvailablePhoneNumbers:
        availableNumbersApi = *swagger.NewAvailablenumbersApi()
      case listSubaccounts:
        subAccountsApi = *swagger.NewSubaccountsApi()
      case listAccounts, getAccount:
        accountsApi = *swagger.NewAccountsApi()
      case listApplications, getApplication:
        applicationsApi = *swagger.NewApplicationsApi()
      case createCall:
        callsApi = *swagger.NewCallsApi()
      case listCallLogs:
        callLogsApi = *swagger.NewCalllogsApi()
      case listDevices, getDevice, createDevice, replaceDevice:
        devicesApi = *swagger.NewDevicesApi()
      case listExpressServiceCodes, getExpressServiceCode:
        expressServiceCodesApi = *swagger.NewExpressservicecodesApi()
      case listExtensions, getExtension, createExtension, replaceExtension:
        extensionsApi = *swagger.NewExtensionsApi()
      case getCallerId:
        callerIdsApi = *swagger.NewCalleridsApi()
      case listContacts, getContact, createContact, replaceContact, deleteContact:
        contactsApi = *swagger.NewContactsApi()
      case listGroups, getGroup, createGroup, replaceGroup, deleteGroup:
        groupsApi = *swagger.NewGroupsApi()
      case listPhoneNumbers, getPhoneNumber, createPhoneNumber, replacePhoneNumber:
        phoneNumbersApi = *swagger.NewPhonenumbersApi()
      default:
        fmt.Println("Invalid command:", command)
        return nil
    }

		switch (command) {

      case listMedia:
        handle(mediaApi.ListAccountMedia(accountId, slice, slice, "", "", limit, offset, ""))

      case getRecording:
				handle(mediaApi.GetAccountMedia(accountId, id))

      case listMenus:
				handle(menusApi.ListAccountMenus(accountId, slice, slice, "", "", limit, offset, ""))

      case getMenu:
				handle(menusApi.GetAccountMenu(accountId, id))

      case listQueues:
				handle(queuesApi.ListAccountQueues(accountId, slice, slice, "", "", limit, offset, ""))

      case getQueue:
				handle(queuesApi.GetAccountQueue(accountId, id))

      case listRoutes:
				handle(routesApi.ListAccountRoutes(accountId, slice, slice, "", "", limit, offset, ""))

      case getRoute:
				handle(routesApi.GetAccountRoute(accountId, id))

      case listSchedules:
				handle(schedulesApi.ListAccountSchedules(accountId, slice, slice, "", "", limit, offset, ""))

      case getSchedule:
				handle(schedulesApi.GetAccountSchedule(accountId, id))

      case listSms:
				handle(smsApi.ListAccountSms(accountId, slice, "", "", "", "", limit, offset, ""))

      case getSms:
				handle(smsApi.GetAccountSms(accountId, id))

      case listAvailablePhoneNumbers:
				handle(availableNumbersApi.ListAvailablePhoneNumbers(slice, slice, slice, slice, slice, slice, slice, slice, slice, slice, slice, "", "", "", limit, offset, ""))

      case listSubaccounts:
				handle(subAccountsApi.ListAccountSubaccounts(accountId, slice, "", limit, offset, ""))

      case listAccounts:
				handle(accountsApi.ListAccounts(slice, "", limit, offset, ""))

      case getAccount:
				handle(accountsApi.GetAccount(accountId))

      case listApplications:
				handle(applicationsApi.ListAccountApplications(accountId, slice, slice, "", "", limit, offset, ""))

      case getApplication:
				handle(applicationsApi.GetAccountApplication(accountId, 1356077))

      case createCall:

        var params = createCallParams()
				handle(callsApi.CreateAccountCalls(accountId, params))

      case listCallLogs:
				handle(callLogsApi.ListAccountCallLogs(accountId, slice, slice, "", "", "", "", slice, "", "", "", limit, offset, ""))

      case listDevices:
				handle(devicesApi.ListAccountDevices(accountId, slice, slice, "", "", limit, offset, ""))

      case getDevice:
				handle(devicesApi.GetAccountDevice(accountId, 144510))

      case createDevice:

        var params = createDeviceParams();
				handle(devicesApi.CreateAccountDevice(accountId, params))

      case replaceDevice:

        var params = createDeviceParams();
				handle(devicesApi.ReplaceAccountDevice(accountId, 142315, params))

      case listExpressServiceCodes:
				handle(expressServiceCodesApi.ListAccountExpressSrvCodes(accountId, slice))

      case getExpressServiceCode:
				handle(expressServiceCodesApi.GetAccountExpressSrvCode(accountId, 324202))

      case listExtensions:
				handle(extensionsApi.ListAccountExtensions(accountId, slice, slice, slice, "", "", "", limit, offset, ""))

      case getExtension:
        handle(extensionsApi.GetAccountExtension(accountId, 1767129))

      case createExtension:

        var params = createExtensionsParams()
				handle(extensionsApi.CreateAccountExtension(accountId, params))

      case replaceExtension:

        var params = replaceExtensionParams()
        handle(extensionsApi.ReplaceAccountExtension(accountId, 1767129, params))

      case getCallerId:
        handle(callerIdsApi.GetCallerIds(accountId, 1764590, slice, slice, "", "", limit, offset, ""))

      case listContacts:
        handle(contactsApi.ListAccountExtensionContacts(accountId, 1764590, slice, slice, slice, "", "", limit, offset, ""))
        
      case getContact:
        handle(contactsApi.GetAccountExtensionContact(accountId, 1764590, 2074702))

      case createContact:

        var params = createContactParams()
        handle(contactsApi.CreateAccountExtensionContact(accountId, 1764590, params))

      case replaceContact:
        var params = createContactParams();
        handle(contactsApi.ReplaceAccountExtensionContact(accountId, 1764590, params))
        
      case deleteContact:
        handle(contactsApi.DeleteAccountExtensionContact(accountId, 1764590, 2072969))
        
      case listGroups:
        handle(groupsApi.ListAccountExtensionContactGroups(accountId, 1764590, slice, slice, "", "", limit, offset, ""))
        
      case getGroup:
        handle(groupsApi.GetAccountExtensionContactGroup(accountId, 1764590, 331026))
        
      case createGroup:

        var params = createGroupParams()
        handle(groupsApi.CreateAccountExtensionContactGroup(accountId, 1764590, params))
        
      case replaceGroup:

        var groupParams swagger.CreateGroupParams
        groupParams.Name = "groupName2"
        handle(groupsApi.ReplaceAccountExtensionContactGroup(accountId, 1764590, 331978))
        
      case deleteGroup:
        handle(groupsApi.DeleteAccountExtensionContactGroup(accountId, 1764590, 331978))
        
      case listPhoneNumbers:
        handle(phoneNumbersApi.ListAccountPhoneNumbers(accountId, slice, slice, slice, "", "", "", limit, offset, ""))

      case getPhoneNumber:
        handle(phoneNumbersApi.GetAccountPhoneNumber(accountId, 2116986))
        
      case createPhoneNumber:

        var params = createPhoneNumberParams()
        handle(phoneNumbersApi.CreateAccountPhoneNumber(accountId, params))
        
      case replacePhoneNumber:

        var params = replacePhoneNumberParams()
        handle(phoneNumbersApi.ReplaceAccountPhoneNumber(accountId, 2116986, params))

      default:

			  return nil
    }

    return nil
  }

  app.Run(os.Args)
}

func handle(
    x interface{},
    response *swagger.APIResponse,
    error error) {

  if (error != nil) {
    panic(error)
  }

  fmt.Printf("%+v\n%+v\n", x, response)
}
