package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
  "phonecom-go-sdk"
)

func main() {
  app := cli.NewApp()

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "command",
      Value: "list-accounts",
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
          case "list-media":
     
     
       
        var mediaApi swagger.MediaApi = *swagger.NewMediaApi()
        mySlice1 := make([]string, 0)
        x,response,error := mediaApi.ListAccountMedia(1315091, mySlice1, mySlice1, "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
       
        
          case "get-recording":
           var mediaApi swagger.MediaApi = *swagger.NewMediaApi()
          x,response,error := mediaApi.GetAccountMedia(1315091, int32(c.Int("id")))
          if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
          case "list-menus":
           var menusApi swagger.MenusApi = *swagger.NewMenusApi()
           mySlice1 := make([]string, 0)
        x,response,error := menusApi.ListAccountMenus(1315091, mySlice1, mySlice1, "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
          case "get-menu":
          var menusApi swagger.MenusApi = *swagger.NewMenusApi()
          
          x,response,error := menusApi.GetAccountMenu(1315091, int32(c.Int("id")))
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
          case "list-queues":
           var queuesApi swagger.QueuesApi = *swagger.NewQueuesApi()
           mySlice1 := make([]string, 0)
        x,response,error := queuesApi.ListAccountQueues(1315091, mySlice1, mySlice1, "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
           case "get-queue":
         var queuesApi swagger.QueuesApi = *swagger.NewQueuesApi()
          
          x,response,error := queuesApi.GetAccountQueue(1315091, int32(c.Int("id")))
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
                    case "list-routes":
           var routesApi swagger.RoutesApi = *swagger.NewRoutesApi()
           mySlice1 := make([]string, 0)
        x,response,error := routesApi.ListAccountRoutes(1315091, mySlice1, mySlice1, "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
           case "get-route":
         var routesApi swagger.RoutesApi = *swagger.NewRoutesApi()
          
          x,response,error := routesApi.GetAccountRoute(1315091, int32(c.Int("id")))
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
           case "list-schedules":
           var schedulesApi swagger.SchedulesApi = *swagger.NewSchedulesApi()
           mySlice1 := make([]string, 0)
        x,response,error := schedulesApi.ListAccountSchedules(1315091, mySlice1, mySlice1, "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
              case "get-schedule":
        var schedulesApi swagger.SchedulesApi = *swagger.NewSchedulesApi()
          
          x,response,error := schedulesApi.GetAccountSchedule(1315091, int32(c.Int("id")))
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
           case "list-sms":
           var smsApi swagger.SmsApi = *swagger.NewSmsApi()
           mySlice1 := make([]string, 0)
        x,response,error := smsApi.ListAccountSms(1315091, mySlice1, "","", "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
                case "get-sms":
      var smsApi swagger.SmsApi = *swagger.NewSmsApi()
          
          x,response,error := smsApi.GetAccountSms(1315091, int32(c.Int("id")))
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
           case "list-available-phone-numbers":
           var availableNumbersApi swagger.AvailablenumbersApi = *swagger.NewAvailablenumbersApi()
           mySlice1 := make([]string, 0)
        x,response,error := availableNumbersApi.ListAvailablePhoneNumbers(mySlice1, mySlice1,mySlice1,mySlice1,mySlice1,mySlice1,mySlice1,mySlice1,mySlice1,mySlice1,mySlice1,"", "", "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
              case "list-subaccounts":
           var subaccountsApi swagger.SubaccountsApi = *swagger.NewSubaccountsApi()
           mySlice1 := make([]string, 0)
        x,response,error := subaccountsApi.ListAccountSubaccounts(1315091, mySlice1, "", int32( c.Int("limit")), int32( c.Int("offset")), "")
        if error!=nil {
            panic(error)
            return nil
            
        }
           _ = response
      fmt.Printf("%+v\n",x)
          
          
          
		case "list-accounts":
			var accountsApi swagger.AccountsApi = *swagger.NewAccountsApi()
			mySlice1 := make([]string, 0) 
			x,response,error := accountsApi.ListAccounts(mySlice1, "", int32(c.Int("limit")), int32( c.Int("offset")), "") 
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "get-account":
			var accountsApi swagger.AccountsApi = *swagger.NewAccountsApi()
			x,response,error := accountsApi.GetAccount(1315091)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "list-applications":
			var applicationsApi swagger.ApplicationsApi = *swagger.NewApplicationsApi()
			mySlice1 := make([]string, 0)
			x,response,error := applicationsApi.ListAccountApplications(1315091, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32( c.Int("offset")), "")
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "get-application":
			var applicationsApi swagger.ApplicationsApi = *swagger.NewApplicationsApi()
			x,response,error := applicationsApi.GetAccountApplication(1315091, 1356077)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "create-call":
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
			x,response,error := callsApi.CreateAccountCalls(1315091, callParams)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "list-call-logs":
			var calllogsApi swagger.CalllogsApi = *swagger.NewCalllogsApi()
			mySlice1 := make([]string, 0)
			x,response,error := calllogsApi.ListAccountCallLogs(1315091, mySlice1, mySlice1, "", "", "", "", mySlice1, "", "", "", int32(c.Int("limit")), int32( c.Int("offset")), "")
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "list-devices":
			var devicesApi swagger.DevicesApi = *swagger.NewDevicesApi()
			mySlice1 := make([]string, 0)
			x,response,error := devicesApi.ListAccountDevices(1315091, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32( c.Int("offset")), "")
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "get-device":
			var devicesApi swagger.DevicesApi = *swagger.NewDevicesApi()
			x,response,error := devicesApi.GetAccountDevice(1315091, 144510)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "create-device":
			var devicesApi swagger.DevicesApi = *swagger.NewDevicesApi()
			var deviceParams swagger.CreateDeviceParams
			deviceParams.Name = "name1"
			//~ var line1 swagger.Line
			//~ line1.Line = 10
			//~ var line2 swagger.Line
			//~ line2.Line = 10
			//~ deviceParams.Lines = []swagger.Line{line1, line2}
			x,response,error := devicesApi.CreateAccountDevice(1315091, deviceParams)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "replace-device":
			var devicesApi swagger.DevicesApi = *swagger.NewDevicesApi()
			var deviceParams swagger.CreateDeviceParams
			deviceParams.Name = "name2"
			//~ deviceParams.Lines = 
			x,response,error := devicesApi.ReplaceAccountDevice(1315091, 142315, deviceParams)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "list-express-service-codes":
			var expressservicecodesApi swagger.ExpressservicecodesApi = *swagger.NewExpressservicecodesApi()
			mySlice1 := make([]string, 0)
			x,response,error := expressservicecodesApi.ListAccountExpressSrvCodes(1315091, mySlice1)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "get-express-service-code":
			var expressservicecodesApi swagger.ExpressservicecodesApi = *swagger.NewExpressservicecodesApi()
			x,response,error := expressservicecodesApi.GetAccountExpressSrvCode(1315091, 324202)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "list-extensions":
			var extensionsApi swagger.ExtensionsApi = *swagger.NewExtensionsApi()
			mySlice1 := make([]string, 0)
			x,response,error := extensionsApi.ListAccountExtensions(1315091, mySlice1, mySlice1, mySlice1, "", "", "", int32(c.Int("limit")), int32( c.Int("offset")), "")
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "get-extension":
			var extensionsApi swagger.ExtensionsApi = *swagger.NewExtensionsApi()
			x,response,error := extensionsApi.GetAccountExtension(1315091, 1767129)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "create-extension":
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
			x,response,error := extensionsApi.CreateAccountExtension(1315091, extensionParams)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "replace-extension":
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
			x,response,error := extensionsApi.ReplaceAccountExtension(1315091, 1767129, extensionParams)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "get-caller-id":
			var calleridsApi swagger.CalleridsApi = *swagger.NewCalleridsApi()
			mySlice1 := make([]string, 0)
			x,response,error := calleridsApi.GetCallerIds(1315091, 1764590, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32( c.Int("offset")), "")
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "list-contacts":
			var contactsApi swagger.ContactsApi = *swagger.NewContactsApi()
			mySlice1 := make([]string, 0)
			x,response,error := contactsApi.ListAccountExtensionContacts(1315091, 1764590, mySlice1, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32( c.Int("offset")), "")
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "get-contact":
			var contactsApi swagger.ContactsApi = *swagger.NewContactsApi()
			x,response,error := contactsApi.GetAccountExtensionContact(1315091, 1764590, 2074702)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "create-contact":
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
			x,response,error := contactsApi.CreateAccountExtensionContact(1315091, 1764590, contactParams)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "replace-contact":
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
			x,response,error := contactsApi.ReplaceAccountExtensionContact(1315091, 1764590, contactParams)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "delete-contact":
			var contactsApi swagger.ContactsApi = *swagger.NewContactsApi()
			x,response,error := contactsApi.DeleteAccountExtensionContact(1315091, 1764590, 2072969)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "list-groups":
			var groupsApi swagger.GroupsApi = *swagger.NewGroupsApi()
			mySlice1 := make([]string, 0)
			x,response,error := groupsApi.ListAccountExtensionContactGroups(1315091, 1764590, mySlice1, mySlice1, "", "", int32(c.Int("limit")), int32( c.Int("offset")), "")
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "get-group":
			var groupsApi swagger.GroupsApi = *swagger.NewGroupsApi()
			x,response,error := groupsApi.GetAccountExtensionContactGroup(1315091, 1764590, 331026)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "create-group":
			var groupsApi swagger.GroupsApi = *swagger.NewGroupsApi()
			var groupParams swagger.CreateGroupParams
			groupParams.Name = "groupName"
			x,response,error := groupsApi.CreateAccountExtensionContactGroup(1315091, 1764590, groupParams)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "replace-group":
			var groupsApi swagger.GroupsApi = *swagger.NewGroupsApi()
			var groupParams swagger.CreateGroupParams
			groupParams.Name = "groupName2"
			x,response,error := groupsApi.ReplaceAccountExtensionContactGroup(1315091, 1764590, 331978)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "delete-group":
			var groupsApi swagger.GroupsApi = *swagger.NewGroupsApi()
			x,response,error := groupsApi.DeleteAccountExtensionContactGroup(1315091, 1764590, 331978)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "list-phone-numbers":
			var phonenumbersApi swagger.PhonenumbersApi = *swagger.NewPhonenumbersApi()
			mySlice1 := make([]string, 0)
			x,response,error := phonenumbersApi.ListAccountPhoneNumbers(1315091, mySlice1, mySlice1, mySlice1, "", "", "", int32(c.Int("limit")), int32( c.Int("offset")), "")
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "get-phone-number":
			var phonenumbersApi swagger.PhonenumbersApi = *swagger.NewPhonenumbersApi()
			x,response,error := phonenumbersApi.GetAccountPhoneNumber(1315091, 2116986)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "create-phone-number":
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
			x,response,error := phonenumbersApi.CreateAccountPhoneNumber(1315091, phonenumberParams)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
		case "replace-phone-number":
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
			x,response,error := phonenumbersApi.ReplaceAccountPhoneNumber(1315091, 2116986, phonenumberParams)
			fmt.Printf("%+v\n",x)
			fmt.Println("")
			fmt.Println("Response:")
			fmt.Printf("%+v\n", response)
			fmt.Println(error)
			
		default:
			fmt.Println("Command not valid")
      }
    return nil
  }

  app.Run(os.Args)
}
