package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "phonecom-go-sdk"
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
    //    cli.StringFlag{
    //	Name: "verbose, v",
    //	Value: "",
    //	Usage: "Activate verbose mode",
    //},

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
  verbose := c.String("verbose")
  input := c.String("input")
  command := c.String("command")
  idSecondary := int32(c.Int("id-secondary"))
  accountId := int32(c.Int("account"))

  var api interface{} = getApi(command)
  if (api == nil) {
    return nil
  }

  switch api := api.(type) {

  case swagger.MediaApi:

    switch (command) {

    case listMedia:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " with parameters: limit ", limit, " offset ", offset, " account ID ", accountId, " expecting a list of media in JSON format")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountMedia(accountId, slice, slice, "", "", limit, offset, ""))
    case getRecording:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " with parameters: account ID ", accountId, " id of recording ", id, " expecting a recording in JSON format")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.GetAccountMedia(accountId, id))
    }

  case swagger.MenusApi:

    switch (command) {
    case listMenus:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " with parameters: limit ", limit, " offset ", offset, " account ID ", accountId, " expecting a list of menus in JSON format")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountMenus(accountId, slice, slice, "", "", limit, offset, ""))
    case getMenu:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " with parameters: account ID ", accountId, " id of media ", id, " expecting a recording in JSON format")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.GetAccountMenu(accountId, id))
    case createMenu:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " with the parameters found in ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createMenuParams(input)
      handle(api.CreateAccountMenu(accountId, params))
    case replaceMenu:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " replacing the menu with the parameters found in ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = replaceMenuParams(input)
      handle(api.ReplaceAccountMenu(accountId, id, params))
    case deleteMenu:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to delete menu with id ", id)

        if (verbose == "") {
          return nil
        }
      }
      handle(api.DeleteAccountMenu(accountId, id))
    }

  case swagger.QueuesApi:

    switch (command) {
    case listQueues:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " with parameters: limit ", limit, " offset ", offset, " account ID ", accountId, " expecting a list of queues in JSON format")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountQueues(accountId, slice, slice, "", "", limit, offset, ""))
    case getQueue:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " with parameters: account ID ", accountId, " id of media ", id, " expecting a recording in JSON format")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.GetAccountQueue(accountId, id))
    case createQueue:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to create a queue with parameters found in ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createQueueParams(input)
      handle(api.CreateAccountQueue(accountId, params))
    case replaceQueue:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to replace a queue with parameters found in ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createQueueParams(input)
      handle(api.ReplaceAccountQueue(accountId, id, params))
    case deleteQueue:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to delete queue with id ", id)

        if (verbose == "") {
          return nil
        }
      }
      handle(api.DeleteAccountQueue(accountId, id))
    }

  case swagger.RoutesApi:

    switch (command) {
    case listRoutes:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of routes ")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountRoutes(accountId, slice, slice, "", "", limit, offset, ""))
    case getRoute:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a route with id ", id)

        if (verbose == "") {
          return nil
        }
      }
      handle(api.GetAccountRoute(accountId, id))
    case createRoute:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to create a route with parameters from ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createRouteParams(input)
      handle(api.CreateRoute(accountId, params))
    case replaceRoute:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to replace a route with parameters from ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createRouteParams(input)
      handle(api.ReplaceAccountRoute(accountId, id, params))
    case deleteRoute:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to delete a route with id ", id)

        if (verbose == "") {
          return nil
        }
      }
      handle(api.DeleteAccountRoute(accountId, id))
    }

  case swagger.SchedulesApi:

    switch (command) {
    case listSchedules:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a list of schedules in JSON format ")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountSchedules(accountId, slice, slice, "", "", limit, offset, ""))
    case getSchedule:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON schedule with id ", id)

        if (verbose == "") {
          return nil
        }
      }
      handle(api.GetAccountSchedule(accountId, id))
    }

  case swagger.SmsApi:

    switch (command) {
    case listSms:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of sms objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountSms(accountId, slice, "", "", "", "", limit, offset, ""))
    case getSms:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON sms object with id ", id)

        if (verbose == "") {
          return nil
        }
      }
      handle(api.GetAccountSms(accountId, id))
    case createSms:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to create an sms object with parameters from file ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createSmsParams(input)
      handle(api.CreateAccountSms(accountId, params))
    }

  case swagger.AvailablenumbersApi:

    switch (command) {
    case listAvailablePhoneNumbers:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of phone number objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAvailablePhoneNumbers(slice, slice, slice, slice, slice, slice, slice, slice, slice, slice, "", "", "", limit, offset, ""))
    }

  case swagger.SubaccountsApi:

    switch (command) {
    case listSubaccounts:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of subaccount objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountSubaccounts(accountId, slice, "", limit, offset, ""))
    case createSubaccount:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to create a subaccount object with parameters from file ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createSubaccountParams(input)
      handle(api.CreateAccountSubaccount(accountId, params))
    }

  case swagger.AccountsApi:

    switch (command) {

    case listAccounts:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of account objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccounts(slice, "", limit, offset, ""))
    case getAccount:
      handle(api.GetAccount(id))
    }

  case swagger.NumberregionsApi:

    switch (command) {
    case listAvailablePhoneNumberRegions:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of phone number region objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAvailablePhoneNumberRegions(slice, slice, slice, slice, slice, slice, slice, "", "", "", "", "", "", "", limit, offset, "", slice))
    }

  case swagger.ApplicationsApi:

    switch (command) {

    case listApplications:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of application objects")

        if (verbose == "") {
          return nil
        }
      }
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
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of call logs objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountCallLogs(accountId, slice, slice, "", "", "", "", slice, "", "", "", limit, offset, ""))
    //~ case getCallLog:
    //~ handle(api.GetAccountCallLog(accountId, id))
    }

  case swagger.DevicesApi:

    switch (command) {

    case listDevices:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of device objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountDevices(accountId, slice, slice, "", "", limit, offset, ""))
    case getDevice:
      handle(api.GetAccountDevice(accountId, id))
    case createDevice:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to create a device object with parameters from file ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createDeviceParams(input)
      handle(api.CreateAccountDevice(accountId, params))
    case replaceDevice:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to replace a device with parameters found in ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createDeviceParams(input)
      handle(api.ReplaceAccountDevice(accountId, id, params))
    }

  case swagger.ExpressservicecodesApi:

    switch (command) {

    case listExpressServiceCodes:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of express service code objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountExpressSrvCodes(accountId, slice))
    case getExpressServiceCode:
      handle(api.GetAccountExpressSrvCode(accountId, id))
    }

  case swagger.ExtensionsApi:

    switch (command) {

    case listExtensions:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of extension objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountExtensions(accountId, slice, slice, slice, "", "", "", limit, offset, ""))
    case getExtension:
      handle(api.GetAccountExtension(accountId, id))
    case createExtension:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to create an extension object with parameters from file ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createExtensionsParams(input)
      handle(api.CreateAccountExtension(accountId, params))
    case replaceExtension:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to replace an extension with parameters found in ", input)

        if (verbose == "") {
          return nil
        }
      }
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
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of contact objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountExtensionContacts(accountId, id, slice, slice, slice, "", "", limit, offset, ""))
    case getContact:
      handle(api.GetAccountExtensionContact(accountId, id, idSecondary))
    case createContact:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to create an account extension contact object with parameters from file ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createContactParams(input)
      handle(api.CreateAccountExtensionContact(accountId, id, params))
    case replaceContact:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to replace a contact with parameters found in ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createContactParams(input)
      handle(api.ReplaceAccountExtensionContact(accountId, id, params))
    case deleteContact:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a to delete an Account Extension object with id ", id)

        if (verbose == "") {
          return nil
        }
      }
      handle(api.DeleteAccountExtensionContact(accountId, id, idSecondary))
    }

  case swagger.GroupsApi:

    switch (command) {

    case listGroups:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of group objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountExtensionContactGroups(accountId, id, slice, slice, "", "", limit, offset, ""))
    case getGroup:
      handle(api.GetAccountExtensionContactGroup(accountId, id, idSecondary))
    case createGroup:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to create a group object with parameters from file ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createGroupParams(input)
      handle(api.CreateAccountExtensionContactGroup(accountId, id, params))
    case replaceGroup:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to replace a group with parameters found in ", input)

        if (verbose == "") {
          return nil
        }
      }
      //var params = createGroupParams()
      handle(api.ReplaceAccountExtensionContactGroup(accountId, id, idSecondary))
    case deleteGroup:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a to delete a Group object with id ", id)

        if (verbose == "") {
          return nil
        }
      }
      handle(api.DeleteAccountExtensionContactGroup(accountId, id, idSecondary))
    }

  case swagger.PhonenumbersApi:

    switch (command) {

    case listPhoneNumbers:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of phone number objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountPhoneNumbers(accountId, slice, slice, slice, "", "", "", limit, offset, ""))
    case getPhoneNumber:
      handle(api.GetAccountPhoneNumber(accountId, id))
    case createPhoneNumber:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to create a phone number object with parameters from file ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createPhoneNumberParams(input)
      handle(api.CreateAccountPhoneNumber(accountId, params))
    case replacePhoneNumber:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to replace a phone number with parameters found in ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = replacePhoneNumberParams(input)
      handle(api.ReplaceAccountPhoneNumber(accountId, id, params))
    }

  case swagger.TrunksApi:

    switch (command) {

    case listTrunks:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a JSON list of trunk objects")

        if (verbose == "") {
          return nil
        }
      }
      handle(api.ListAccountTrunks(accountId, slice, slice, "", "", limit, offset, ""))
    case getTrunk:
      handle(api.GetAccountTrunk(accountId, id))
    case createTrunk:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to create a trunk object with parameters from file ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createTrunkParams(input)
      handle(api.CreateAccountTrunk(accountId, params))
    case replaceTrunk:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting to replace a trunk with parameters found in ", input)

        if (verbose == "") {
          return nil
        }
      }
      var params = createTrunkParams(input)
      handle(api.ReplaceAccountTrunk(accountId, id, params))
    case deleteTrunk:
      if (dryrun != "" || verbose != "") {
        fmt.Println("Calling ", command, " expecting a to delete a Trunk object with id ", id)

        if (verbose == "") {
          return nil
        }
      }
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
}
