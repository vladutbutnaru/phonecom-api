package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "phonecom-go-sdk"
)

const listAccounts = "list-accounts"
const listMedia = "list-media"
const listMenus = "list-menus"
const listQueues = "list-queues"
const listRoutes = "list-routes"
const listSchedules = "list-schedules"
const listSms = "list-sms"
const listAvailablePhoneNumbers = "list-available-phone-numbers"
const listSubaccounts = "list-subaccounts"
const listApplications = "list-applications"
const listCallLogs = "list-call-logs"
const listDevices = "list-devices"
const listExpressServiceCodes = "list-express-service-codes"
const listExtensions = "list-extensions"
const listContacts = "list-contacts"
const listGroups = "list-groups"
const listPhoneNumbers = "list-phone-numbers"

const getMenu = "get-menu"
const getRecording = "get-recording"
const getQueue = "get-queue"
const getRoute = "get-route"
const getSchedule = "get-schedule"
const getSms = "get-sms"
const getAccount = "get-account"
const getApplication = "get-application"
const getDevice = "get-device"
const getExtension = "get-extension"
const getExpressServiceCode = "get-express-service-code"
const getCallerId = "get-caller-id"
const getContact = "get-contact"
const getGroup = "get-group"
const getPhoneNumber = "get-phone-number"

const createPhoneNumber = "create-phone-number"
const createCall = "create-call"
const createDevice = "create-device"
const createExtension = "create-extension"
const createContact = "create-contact"
const createGroup = "create-group"

const replaceDevice = "replace-device"
const replaceExtension = "replace-extension"
const replaceContact = "replace-contact"
const replaceGroup = "replace-group"
const replacePhoneNumber = "replace-phone-number"

const deleteContact = "delete-contact"
const deleteGroup = "delete-group"

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


    switch c.String("command") {

      case listMedia:

        var mediaApi swagger.MediaApi = *swagger.NewMediaApi()
        mySlice1 := make([]string, 0)
        x, response, callError := mediaApi.ListAccountMedia(1315091, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)

      case getRecording:

        var mediaApi swagger.MediaApi = *swagger.NewMediaApi()
        x, response, callError := mediaApi.GetAccountMedia(1315091, int32(c.Int("id")))
        handle(x, response, callError)

      case listMenus:
        var menusApi swagger.MenusApi = *swagger.NewMenusApi()
        mySlice1 := make([]string, 0)
        x, response, callError := menusApi.ListAccountMenus(1315091, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)

      case getMenu:
        var menusApi swagger.MenusApi = *swagger.NewMenusApi()

        x, response, callError := menusApi.GetAccountMenu(1315091, int32(c.Int("id")))
        handle(x, response, callError)

      case listQueues:
        var queuesApi swagger.QueuesApi = *swagger.NewQueuesApi()
        mySlice1 := make([]string, 0)
        x, response, callError := queuesApi.ListAccountQueues(1315091, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)

      case getQueue:
        var queuesApi swagger.QueuesApi = *swagger.NewQueuesApi()

        x, response, callError := queuesApi.GetAccountQueue(1315091, int32(c.Int("id")))
        handle(x, response, callError)

      case listRoutes:
        var routesApi swagger.RoutesApi = *swagger.NewRoutesApi()
        mySlice1 := make([]string, 0)
        x, response, callError := routesApi.ListAccountRoutes(1315091, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)

      case getRoute:
        var routesApi swagger.RoutesApi = *swagger.NewRoutesApi()

        x, response, callError := routesApi.GetAccountRoute(1315091, int32(c.Int("id")))
        handle(x, response, callError)

      case listSchedules:
        var schedulesApi swagger.SchedulesApi = *swagger.NewSchedulesApi()
        mySlice1 := make([]string, 0)
        x, response, callError := schedulesApi.ListAccountSchedules(1315091, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)

      case getSchedule:
        var schedulesApi swagger.SchedulesApi = *swagger.NewSchedulesApi()

        x, response, callError := schedulesApi.GetAccountSchedule(1315091, int32(c.Int("id")))
        handle(x, response, callError)

      case listSms:
        var smsApi swagger.SmsApi = *swagger.NewSmsApi()
        mySlice1 := make([]string, 0)
        x, response, callError := smsApi.ListAccountSms(1315091, mySlice1, "", "", "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)

      case getSms:
        var smsApi swagger.SmsApi = *swagger.NewSmsApi()

        x, response, callError := smsApi.GetAccountSms(1315091, int32(c.Int("id")))
        handle(x, response, callError)

      case listAvailablePhoneNumbers:
        var availableNumbersApi swagger.AvailablenumbersApi = *swagger.NewAvailablenumbersApi()
        mySlice1 := make([]string, 0)
        x, response, callError := availableNumbersApi.ListAvailablePhoneNumbers(mySlice1, mySlice1, mySlice1, mySlice1, mySlice1, mySlice1, mySlice1, mySlice1, mySlice1, mySlice1, mySlice1, "", "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)

      case listSubaccounts:
        var subaccountsApi swagger.SubaccountsApi = *swagger.NewSubaccountsApi()
        mySlice1 := make([]string, 0)
        x, response, callError := subaccountsApi.ListAccountSubaccounts(1315091, mySlice1, "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)

      case listAccounts:
        var accountsApi swagger.AccountsApi = *swagger.NewAccountsApi()
        mySlice1 := make([]string, 0)
        x, response, callError := accountsApi.ListAccounts(mySlice1, "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)

      case getAccount:
        var accountsApi swagger.AccountsApi = *swagger.NewAccountsApi()
        x, response, callError := accountsApi.GetAccount(1315091)
        handle(x, response, callError)
      case listApplications:
        var applicationsApi swagger.ApplicationsApi = *swagger.NewApplicationsApi()
        mySlice1 := make([]string, 0)
        x, response, callError := applicationsApi.ListAccountApplications(1315091, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)
      case getApplication:
        var applicationsApi swagger.ApplicationsApi = *swagger.NewApplicationsApi()
        x, response, callError := applicationsApi.GetAccountApplication(1315091, 1356077)
        handle(x, response, callError)
      case createCall:
        var callsApi swagger.CallsApi = *swagger.NewCallsApi()
        var callParams swagger.CreateCallParams
        callParams.CallerPhoneNumber = "callCallerPhoneNumber"
        callParams.CallerExtension = 222
        callParams.CallerCallerId = "callCallerCallerId"
        callParams.CallerPrivate = true
        callParams.CalleePhoneNumber = "callCalleePhoneNumber"
        callParams.CalleeExtension = 222
        callParams.CalleeCallerId = "callCalleeCallerId"
        callParams.CalleePrivate = true
        x, response, callError := callsApi.CreateAccountCalls(1315091, callParams)
        handle(x, response, callError)
      case listCallLogs:
        var calllogsApi swagger.CalllogsApi = *swagger.NewCalllogsApi()
        mySlice1 := make([]string, 0)
        x, response, callError := calllogsApi.ListAccountCallLogs(1315091, mySlice1, mySlice1, "", "", "", "", mySlice1, "", "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)
      case listDevices:
        var devicesApi swagger.DevicesApi = *swagger.NewDevicesApi()
        mySlice1 := make([]string, 0)
        x, response, callError := devicesApi.ListAccountDevices(1315091, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)
      case getDevice:
        var devicesApi swagger.DevicesApi = *swagger.NewDevicesApi()
        x, response, callError := devicesApi.GetAccountDevice(1315091, 144510)
        handle(x, response, callError)
      case createDevice:
        var devicesApi swagger.DevicesApi = *swagger.NewDevicesApi()
        var deviceParams swagger.CreateDeviceParams
        deviceParams.Name = "name1"
        //~ var line1 swagger.Line
        //~ line1.Line = 10
        //~ var line2 swagger.Line
        //~ line2.Line = 10
        //~ deviceParams.Lines = []swagger.Line{line1, line2}
        x, response, callError := devicesApi.CreateAccountDevice(1315091, deviceParams)
        handle(x, response, callError)
      case replaceDevice:
        var devicesApi swagger.DevicesApi = *swagger.NewDevicesApi()
        var deviceParams swagger.CreateDeviceParams
        deviceParams.Name = "name2"
        //~ deviceParams.Lines =
        x, response, callError := devicesApi.ReplaceAccountDevice(1315091, 142315, deviceParams)
        handle(x, response, callError)
      case listExpressServiceCodes:
        var expressservicecodesApi swagger.ExpressservicecodesApi = *swagger.NewExpressservicecodesApi()
        mySlice1 := make([]string, 0)
        x, response, callError := expressservicecodesApi.ListAccountExpressSrvCodes(1315091, mySlice1)
        handle(x, response, callError)
      case getExpressServiceCode:
        var expressservicecodesApi swagger.ExpressservicecodesApi = *swagger.NewExpressservicecodesApi()
        x, response, callError := expressservicecodesApi.GetAccountExpressSrvCode(1315091, 324202)
        handle(x, response, callError)
      case listExtensions:
        var extensionsApi swagger.ExtensionsApi = *swagger.NewExtensionsApi()
        mySlice1 := make([]string, 0)
        x, response, callError := extensionsApi.ListAccountExtensions(1315091, mySlice1, mySlice1, mySlice1, "", "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)
      case getExtension:
        var extensionsApi swagger.ExtensionsApi = *swagger.NewExtensionsApi()
        x, response, callError := extensionsApi.GetAccountExtension(1315091, 1767129)
        handle(x, response, callError)
      case createExtension:
        var extensionsApi swagger.ExtensionsApi = *swagger.NewExtensionsApi()
        var extensionParams swagger.CreateExtensionParams
        extensionParams.CallerId = "private"
        extensionParams.UsageType = "limited"
        extensionParams.AllowsCallWaiting = true
        extensionParams.Extension = 15524
        extensionParams.IncludeInDirectory = true
        extensionParams.Name = "extensionName"
        extensionParams.FullName = "extensionFullName"
        extensionParams.Timezone = "America/Los_Angeles"
        extensionParams.LocalAreaCode = 111
        extensionParams.VoicemailGreetingEnableLeaveMessagePrompt = true
        extensionParams.VoicemailEnabled = true
        extensionParams.EnableOutboundCalls = true
        extensionParams.EnableCallWaiting = true
        extensionParams.VoicemailPassword = 111
        extensionParams.VoicemailGreetingType = "extensionVoicemailGreetingType"
        extensionParams.VoicemailTranscription = "extensionVoicemailTranscription"
        extensionParams.VoicemailNotificationsEmails = []string{"extensionVoicemailNotificationsEmails1", "extensionVoicemailNotificationsEmails2"}
        extensionParams.VoicemailNotificationsSms = "extensionVoicemailNotificationsSms"
        extensionParams.CallNotificationsEmails = []string{"extensionCallNotificationsEmails1", "extensionCallNotificationsEmails2"}
        extensionParams.CallNotificationsSms = "extensionCallNotificationsSms"
        x, response, callError := extensionsApi.CreateAccountExtension(1315091, extensionParams)
        handle(x, response, callError)
      case replaceExtension:
        var extensionsApi swagger.ExtensionsApi = *swagger.NewExtensionsApi()
        var extensionParams swagger.ReplaceExtensionParams
        //~ extensionParams.Name = "extensionName2"
        //~ extensionParams.Timezone = "Europe/Skopje"
        //~ extensionParams.IncludeInDirectory = false
        //~ extensionParams.Extension = 15524
        //~ extensionParams.EnableOutboundCalls = false
        //~ extensionParams.UsageType = "extensionUsageType1" ###
        //~ extensionParams.VoicemailPassword = 222
        //~ extensionParams.FullName = "extensionFullName2"
        //~ extensionParams.EnableCallWaiting = false
        //~ extensionParams.VoicemailGreetingType = "extensionVoicemailGreetingType2"
        //~ extensionParams.CallerId = "extensionCallerId1" ###
        //~ extensionParams.LocalAreaCode = 620
        //~ extensionParams.VoicemailEnabled = false
        //~ extensionParams.VoicemailGreetingEnableLeaveMessagePrompt = false
        //~ extensionParams.VoicemailTranscription = "extensionVoicemailTranscription2"
        //~ extensionParams.VoicemailNotificationsEmails = []string{"extensionVoicemailNotificationsEmails3", "extensionVoicemailNotificationsEmails4"}
        //~ extensionParams.VoicemailNotificationsSms = "extensionVoicemailNotificationsSms2"
        //~ extensionParams.CallNotificationsEmails = []string{"extensionCallNotificationsEmails3", "extensionCallNotificationsEmails4"}
        //~ extensionParams.CallNotificationsSms = "extensionCallNotificationsSms2"
        //~ extensionParams.Route = []string{"extensionRoute3", "extensionRoute4"}
        x, response, callError := extensionsApi.ReplaceAccountExtension(1315091, 1767129, extensionParams)
        handle(x, response, callError)
      case getCallerId:
        var calleridsApi swagger.CalleridsApi = *swagger.NewCalleridsApi()
        mySlice1 := make([]string, 0)
        x, response, callError := calleridsApi.GetCallerIds(1315091, 1764590, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)
      case listContacts:
        var contactsApi swagger.ContactsApi = *swagger.NewContactsApi()
        mySlice1 := make([]string, 0)
        x, response, callError := contactsApi.ListAccountExtensionContacts(1315091, 1764590, mySlice1, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)
      case getContact:
        var contactsApi swagger.ContactsApi = *swagger.NewContactsApi()
        x, response, callError := contactsApi.GetAccountExtensionContact(1315091, 1764590, 2074702)
        handle(x, response, callError)
      case createContact:
        var contactsApi swagger.ContactsApi = *swagger.NewContactsApi()
        var contactParams swagger.CreateContactParams
        contactParams.FirstName = "contactFirstName"
        contactParams.MiddleName = "contactMiddleName"
        contactParams.LastName = "contactLastName"
        contactParams.Prefix = "contactPrefix"
        contactParams.PhoneticFirstName = "contactPhoneticFirstName"
        contactParams.PhoneticMiddleName = "contactPhoneticMiddleName"
        contactParams.PhoneticLastName = "contactPhoneticLastName"
        contactParams.Suffix = "contactSuffix"
        contactParams.Nickname = "contactNickname"
        contactParams.Company = "contactCompany"
        contactParams.Department = "contactDepartment"
        contactParams.JobTitle = "contactJobTitle"
        x, response, callError := contactsApi.CreateAccountExtensionContact(1315091, 1764590, contactParams)
        handle(x, response, callError)
      case replaceContact:
        var contactsApi swagger.ContactsApi = *swagger.NewContactsApi()
        var contactParams swagger.CreateContactParams
        contactParams.FirstName = "contactFirstName3"
        //~ contactParams.MiddleName = "contactMiddleName3"
        //~ contactParams.LastName = "contactLastName3"
        //~ contactParams.Prefix = "contactPrefix3"
        //~ contactParams.PhoneticFirstName = "contactPhoneticFirstName3"
        //~ contactParams.PhoneticMiddleName = "contactPhoneticMiddleName3"
        //~ contactParams.PhoneticLastName = "contactPhoneticLastName3"
        //~ contactParams.Suffix = "contactSuffix3"
        //~ contactParams.Nickname = "contactNickname3"
        //~ contactParams.Company = "contactCompany3"
        //~ contactParams.Department = "contactDepartment3"
        //~ contactParams.JobTitle = "contactJobTitle3"
        x, response, callError := contactsApi.ReplaceAccountExtensionContact(1315091, 1764590, contactParams)
        handle(x, response, callError)
      case deleteContact:
        var contactsApi swagger.ContactsApi = *swagger.NewContactsApi()
        x, response, callError := contactsApi.DeleteAccountExtensionContact(1315091, 1764590, 2072969)
        handle(x, response, callError)
      case listGroups:
        var groupsApi swagger.GroupsApi = *swagger.NewGroupsApi()
        mySlice1 := make([]string, 0)
        x, response, callError := groupsApi.ListAccountExtensionContactGroups(1315091, 1764590, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)
      case getGroup:
        var groupsApi swagger.GroupsApi = *swagger.NewGroupsApi()
        x, response, callError := groupsApi.GetAccountExtensionContactGroup(1315091, 1764590, 331026)
        handle(x, response, callError)
      case createGroup:
        var groupsApi swagger.GroupsApi = *swagger.NewGroupsApi()
        var groupParams swagger.CreateGroupParams
        groupParams.Name = "groupName"
        x, response, callError := groupsApi.CreateAccountExtensionContactGroup(1315091, 1764590, groupParams)
        handle(x, response, callError)
      case replaceGroup:
        var groupsApi swagger.GroupsApi = *swagger.NewGroupsApi()
        var groupParams swagger.CreateGroupParams
        groupParams.Name = "groupName2"
        x, response, callError := groupsApi.ReplaceAccountExtensionContactGroup(1315091, 1764590, 331978)
        handle(x, response, callError)
      case deleteGroup:
        var groupsApi swagger.GroupsApi = *swagger.NewGroupsApi()
        x, response, callError := groupsApi.DeleteAccountExtensionContactGroup(1315091, 1764590, 331978)
        handle(x, response, callError)
      case listPhoneNumbers:
        var phonenumbersApi swagger.PhonenumbersApi = *swagger.NewPhonenumbersApi()
        mySlice1 := make([]string, 0)
        x, response, callError := phonenumbersApi.ListAccountPhoneNumbers(1315091, mySlice1, mySlice1, mySlice1, "", "", "", int32(c.Int("limit")), int32(c.Int("offset")), "")
        handle(x, response, callError)
      case getPhoneNumber:
        var phonenumbersApi swagger.PhonenumbersApi = *swagger.NewPhonenumbersApi()
        x, response, callError := phonenumbersApi.GetAccountPhoneNumber(1315091, 2116986)
        handle(x, response, callError)
      case createPhoneNumber:
        var phonenumbersApi swagger.PhonenumbersApi = *swagger.NewPhonenumbersApi()
        var phonenumberParams swagger.CreatePhoneNumberParams
        //~ phonenumberParams.Name = "phonenumberName"
        //~ phonenumberParams.BlockIncoming = true
        //~ phonenumberParams.BlockAnonymous = true
        //~ phonenumberParams.CallerIdName = "Caller Id Name"
        //~ phonenumberParams.CallerIdType = "phonenumberCallerIdType"
        //~ phonenumberParams.SmsForwardingType = "phonenumberSmsForwardingType"
        //~ phonenumberParams.CallNotificationsEmails = []string{"phonenumberCallNotificationsEmails1", "phonenumberCallNotificationsEmails2"}
        //~ phonenumberParams.CallNotificationsSms = "phonenumberCallNotificationsSms"
        x, response, callError := phonenumbersApi.CreateAccountPhoneNumber(1315091, phonenumberParams)
        handle(x, response, callError)
      case replacePhoneNumber:
        var phonenumbersApi swagger.PhonenumbersApi = *swagger.NewPhonenumbersApi()
        var phonenumberParams swagger.ReplacePhoneNumberParams
        phonenumberParams.Name = "The Name"
        phonenumberParams.BlockIncoming = false
        phonenumberParams.BlockAnonymous = false
        var callerId swagger.CallerIdFull
        callerId.Name = "The Caller Id Name"
        callerId.Type_ = "The Caller Id Type"
        phonenumberParams.CallerIdName = callerId.Name
        phonenumberParams.CallerIdType = callerId.Type_
        phonenumberParams.SmsForwardingType = "The Sms Forwarding Type"
        phonenumberParams.CallNotificationsEmails = []string{"The Call Notifications Emails 1", "The Call Notifications Emails 2"}
        phonenumberParams.CallNotificationsSms = "The Call Notifications Sms"
        x, response, callError := phonenumbersApi.ReplaceAccountPhoneNumber(1315091, 2116986, phonenumberParams)
        handle(x, response, callError)

      default:
        fmt.Println("Command not valid")
    }
    return nil
  }

  app.Run(os.Args)
}

func handle(
    x interface{},
    response *swagger.APIResponse,
    error error) {

  if error != nil {
    panic(error)
  }

  fmt.Printf("%+v\n%+v\n", x, response)
}
