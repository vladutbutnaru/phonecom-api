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

		switch c.String("command") {

      case listMedia:

        var api swagger.MediaApi = *swagger.NewMediaApi()
        handle(api.ListAccountMedia(acountId, slice, slice, "", "", limit, offset, ""))

      case getRecording:

        var api swagger.MediaApi = *swagger.NewMediaApi()
				handle(api.GetAccountMedia(acountId, int32(c.Int("id"))))

      case listMenus:

        var api swagger.MenusApi = *swagger.NewMenusApi()
				handle(api.ListAccountMenus(acountId, slice, slice, "", "", limit, offset, ""))

      case getMenu:

        var api swagger.MenusApi = *swagger.NewMenusApi()
				handle(api.GetAccountMenu(acountId, int32(c.Int("id"))))

      case listQueues:

        var api swagger.QueuesApi = *swagger.NewQueuesApi()
				handle(api.ListAccountQueues(acountId, slice, slice, "", "", limit, offset, ""))

      case getQueue:

        var api swagger.QueuesApi = *swagger.NewQueuesApi()
				handle(api.GetAccountQueue(acountId, int32(c.Int("id"))))

      case listRoutes:

        var api swagger.RoutesApi = *swagger.NewRoutesApi()
				handle(api.ListAccountRoutes(acountId, slice, slice, "", "", limit, offset, ""))

      case getRoute:

        var api swagger.RoutesApi = *swagger.NewRoutesApi()
				handle(api.GetAccountRoute(acountId, int32(c.Int("id"))))

      case listSchedules:

        var api swagger.SchedulesApi = *swagger.NewSchedulesApi()
				handle(api.ListAccountSchedules(acountId, slice, slice, "", "", limit, offset, ""))

      case getSchedule:

        var api swagger.SchedulesApi = *swagger.NewSchedulesApi()
				handle(api.GetAccountSchedule(acountId, int32(c.Int("id"))))

      case listSms:

        var api swagger.SmsApi = *swagger.NewSmsApi()
				handle(api.ListAccountSms(acountId, slice, "", "", "", "", limit, offset, ""))

      case getSms:

        var api swagger.SmsApi = *swagger.NewSmsApi()
				handle(api.GetAccountSms(acountId, int32(c.Int("id"))))

      case listAvailablePhoneNumbers:

        var api swagger.AvailablenumbersApi = *swagger.NewAvailablenumbersApi()
				handle(api.ListAvailablePhoneNumbers(slice, slice, slice, slice, slice, slice, slice, slice, slice, slice, slice, "", "", "", limit, offset, ""))

      case listSubaccounts:

        var api swagger.SubaccountsApi = *swagger.NewSubaccountsApi()
				handle(api.ListAccountSubaccounts(acountId, slice, "", limit, offset, ""))

      case listAccounts:

        var api swagger.AccountsApi = *swagger.NewAccountsApi()
				handle(api.ListAccounts(slice, "", limit, offset, ""))

      case getAccount:

        var api swagger.AccountsApi = *swagger.NewAccountsApi()
				handle(api.GetAccount(acountId))

      case listApplications:

        var api swagger.ApplicationsApi = *swagger.NewApplicationsApi()
				handle(api.ListAccountApplications(acountId, slice, slice, "", "", limit, offset, ""))

      case getApplication:

        var api swagger.ApplicationsApi = *swagger.NewApplicationsApi()
				handle(api.GetAccountApplication(acountId, 1356077))

      case createCall:

        var api swagger.CallsApi = *swagger.NewCallsApi()
        var callParams swagger.CreateCallParams
        callParams.CallerPhoneNumber = "callCallerPhoneNumber"
        callParams.CallerExtension = 222
        callParams.CallerCallerId = "callCallerCallerId"
        callParams.CallerPrivate = true
        callParams.CalleePhoneNumber = "callCalleePhoneNumber"
        callParams.CalleeExtension = 222
        callParams.CalleeCallerId = "callCalleeCallerId"
        callParams.CalleePrivate = true
				handle(api.CreateAccountCalls(acountId, callParams))

      case listCallLogs:

        var api swagger.CalllogsApi = *swagger.NewCalllogsApi()
				handle(api.ListAccountCallLogs(acountId, slice, slice, "", "", "", "", slice, "", "", "", limit, offset, ""))

      case listDevices:

        var api swagger.DevicesApi = *swagger.NewDevicesApi()
				handle(api.ListAccountDevices(acountId, slice, slice, "", "", limit, offset, ""))

      case getDevice:

        var api swagger.DevicesApi = *swagger.NewDevicesApi()
				handle(api.GetAccountDevice(acountId, 144510))

      case createDevice:

        var api swagger.DevicesApi = *swagger.NewDevicesApi()
        var deviceParams swagger.CreateDeviceParams
        deviceParams.Name = "name1"
        //~ var line1 swagger.Line
        //~ line1.Line = 10
        //~ var line2 swagger.Line
        //~ line2.Line = 10
        //~ deviceParams.Lines = []swagger.Line{line1, line2}
				handle(api.CreateAccountDevice(acountId, deviceParams))

      case replaceDevice:

        var api swagger.DevicesApi = *swagger.NewDevicesApi()
        var deviceParams swagger.CreateDeviceParams
        deviceParams.Name = "name2"
        //~ deviceParams.Lines =
				handle(api.ReplaceAccountDevice(acountId, 142315, deviceParams))

      case listExpressServiceCodes:

        var api swagger.ExpressservicecodesApi = *swagger.NewExpressservicecodesApi()
				handle(api.ListAccountExpressSrvCodes(acountId, slice))

      case getExpressServiceCode:

        var api swagger.ExpressservicecodesApi = *swagger.NewExpressservicecodesApi()
				handle(api.GetAccountExpressSrvCode(acountId, 324202))

      case listExtensions:

        var api swagger.ExtensionsApi = *swagger.NewExtensionsApi()
				handle(api.ListAccountExtensions(acountId, slice, slice, slice, "", "", "", limit, offset, ""))

      case getExtension:

        var api swagger.ExtensionsApi = *swagger.NewExtensionsApi()
        handle(api.GetAccountExtension(acountId, 1767129))
        
      case createExtension:

        var api swagger.ExtensionsApi = *swagger.NewExtensionsApi()
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
				handle(api.CreateAccountExtension(acountId, extensionParams))

      case replaceExtension:

        var api swagger.ExtensionsApi = *swagger.NewExtensionsApi()
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
        handle(api.ReplaceAccountExtension(acountId, 1767129, extensionParams))

      case getCallerId:

        var api swagger.CalleridsApi = *swagger.NewCalleridsApi()
        handle(api.GetCallerIds(acountId, 1764590, slice, slice, "", "", limit, offset, ""))

      case listContacts:

        var api swagger.ContactsApi = *swagger.NewContactsApi()
        handle(api.ListAccountExtensionContacts(acountId, 1764590, slice, slice, slice, "", "", limit, offset, ""))
        
      case getContact:

        var api swagger.ContactsApi = *swagger.NewContactsApi()
        handle(api.GetAccountExtensionContact(acountId, 1764590, 2074702))

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
        handle(contactsApi.CreateAccountExtensionContact(acountId, 1764590, contactParams))
        
      case replaceContact:

        var api swagger.ContactsApi = *swagger.NewContactsApi()
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
       handle(api.ReplaceAccountExtensionContact(acountId, 1764590, contactParams))
        
      case deleteContact:

        var api swagger.ContactsApi = *swagger.NewContactsApi()
        handle(api.DeleteAccountExtensionContact(acountId, 1764590, 2072969))
        
      case listGroups:

        var api swagger.GroupsApi = *swagger.NewGroupsApi()
        handle(api.ListAccountExtensionContactGroups(acountId, 1764590, slice, slice, "", "", limit, offset, ""))
        
      case getGroup:

        var api swagger.GroupsApi = *swagger.NewGroupsApi()
        handle(api.GetAccountExtensionContactGroup(acountId, 1764590, 331026))
        
      case createGroup:

        var api swagger.GroupsApi = *swagger.NewGroupsApi()
        var groupParams swagger.CreateGroupParams
        groupParams.Name = "groupName"
        handle(api.CreateAccountExtensionContactGroup(acountId, 1764590, groupParams))
        
      case replaceGroup:

        var api swagger.GroupsApi = *swagger.NewGroupsApi()
        var groupParams swagger.CreateGroupParams
        groupParams.Name = "groupName2"
        handle(api.ReplaceAccountExtensionContactGroup(acountId, 1764590, 331978))
        
      case deleteGroup:

        var api swagger.GroupsApi = *swagger.NewGroupsApi()
        handle(api.DeleteAccountExtensionContactGroup(acountId, 1764590, 331978))
        
      case listPhoneNumbers:

        var api swagger.PhonenumbersApi = *swagger.NewPhonenumbersApi()
        handle(api.ListAccountPhoneNumbers(acountId, slice, slice, slice, "", "", "", limit, offset, ""))
        
      case getPhoneNumber:

        var api swagger.PhonenumbersApi = *swagger.NewPhonenumbersApi()
        handle(api.GetAccountPhoneNumber(acountId, 2116986))
        
      case createPhoneNumber:

        var api swagger.PhonenumbersApi = *swagger.NewPhonenumbersApi()
        var phonenumberParams swagger.CreatePhoneNumberParams
        //~ phonenumberParams.Name = "phonenumberName"
        //~ phonenumberParams.BlockIncoming = true
        //~ phonenumberParams.BlockAnonymous = true
        //~ phonenumberParams.CallerIdName = "Caller Id Name"
        //~ phonenumberParams.CallerIdType = "phonenumberCallerIdType"
        //~ phonenumberParams.SmsForwardingType = "phonenumberSmsForwardingType"
        //~ phonenumberParams.CallNotificationsEmails = []string{"phonenumberCallNotificationsEmails1", "phonenumberCallNotificationsEmails2"}
        //~ phonenumberParams.CallNotificationsSms = "phonenumberCallNotificationsSms"
       handle(api.CreateAccountPhoneNumber(acountId, phonenumberParams))
        
      case replacePhoneNumber:

        var api swagger.PhonenumbersApi = *swagger.NewPhonenumbersApi()
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
        handle(api.ReplaceAccountPhoneNumber(acountId, 2116986, phonenumberParams))

      default:
        fmt.Println("Command not valid")
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

func getApi (apiCall string) interface{} {

	switch (apiCall) {

	  case listAccounts:
		  return *swagger.NewAccountsApi()

  	case listMedia:
	  	return *swagger.NewMediaApi()
	}

	return nil;
}