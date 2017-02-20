package main

import "phonecom-go-sdk"

func createCallParams() swagger.CreateCallParams {

	var params swagger.CreateCallParams

	params.CallerPhoneNumber = "callCallerPhoneNumber"
	params.CallerExtension = 222
	params.CallerCallerId = "callCallerCallerId"
	params.CallerPrivate = true
	params.CalleePhoneNumber = "callCalleePhoneNumber"
	params.CalleeExtension = 222
	params.CalleeCallerId = "callCalleeCallerId"
	params.CalleePrivate = true

	return params
}

func createDeviceParams() swagger.CreateDeviceParams {

	var params swagger.CreateDeviceParams

	params.Name = "name1"

	//~ var line1 swagger.Line
	//~ line1.Line = 10
	//~ var line2 swagger.Line
	//~ line2.Line = 10
	//~ params.Lines = []swagger.Line{line1, line2}

	return params
}
func createMenuParams() swagger.CreateMenuParams {
    var params swagger.CreateMenuParams
    
    params.Name = "Menu Test"
    params.MainMessage = "Message test"
    params.InvalidKeypressMessage = "invalid press"
    params.AllowExtensionDial = false
    params.KeypressWaitTime = 3
    
    return params
    
    
}

func replaceMenuParams() swagger.ReplaceMenuParams{
    var params swagger.ReplaceMenuParams
    
    params.Name = "Teset Menu"
    params.Greeting = "Hello There"
    params.AllowExtensionDial = false
    params.KeypressWaitTime = 4
    
    return params
    
    
    
    
    
}
func createExtensionsParams() swagger.CreateExtensionParams {

	var params swagger.CreateExtensionParams

	params.CallerId = "private"
	params.UsageType = "limited"
	params.AllowsCallWaiting = true
	params.Extension = 15524
	params.IncludeInDirectory = true
	params.Name = "extensionName"
	params.FullName = "extensionFullName"
	params.Timezone = "America/Los_Angeles"
	params.LocalAreaCode = 111
	params.VoicemailGreetingEnableLeaveMessagePrompt = true
	params.VoicemailEnabled = true
	params.EnableOutboundCalls = true
	params.EnableCallWaiting = true
	params.VoicemailPassword = 111
	params.VoicemailGreetingType = "extensionVoicemailGreetingType"
	params.VoicemailTranscription = "extensionVoicemailTranscription"
	params.VoicemailNotificationsEmails = []string{"extensionVoicemailNotificationsEmails1", "extensionVoicemailNotificationsEmails2"}
	params.VoicemailNotificationsSms = "extensionVoicemailNotificationsSms"
	params.CallNotificationsEmails = []string{"extensionCallNotificationsEmails1", "extensionCallNotificationsEmails2"}
	params.CallNotificationsSms = "extensionCallNotificationsSms"

	return params
}
func createQueueParams() swagger.CreateQueueParams {
    var params swagger.CreateQueueParams
    
    params.Name = "Queue New"
    params.MaxHoldTime = 60
    params.RingTime = 10
    params.Greeting = "Barge landing"
    params.HoldMusic = "electronic"
    params.CallerIdType = "called_number"
    
    return params
    
}

func createRouteParams() swagger.CreateRouteParams {
    var params swagger.CreateRouteParams
    
    params.Name = "New route"
    
    return params
    
    
}

func createSmsParams() swagger.CreateSmsParams {
    var params swagger.CreateSmsParams
    
    params.From = "+12015880100"
    params.To = "+17324840911"
    params.Text = "Test SMS CLI"
    
    return params
}

func createSubaccountParams() swagger.CreateSubaccountParams{
    var params swagger.CreateSubaccountParams
    
    params.Username = "testUsername"
    params.Password = "testPassword"
    
    var contact swagger.ContactSubaccount
    contact.Name = "testName"
    contact.PrimaryEmail = "gigel@gigel.com"
    contact.Company = "FeelITServices"
    contact.Phone = "+12015880100"
    
    var address swagger.Address
    address.Line1 = "Line1 address"
    address.Line2 = "Line 2 address"
    address.City = "Paris"
    address.Province = "Paris"
    address.PostalCode = "700134"
    address.Country = "France"
    contact.Address = address
    
    params.Contact = contact
    params.BillingContact = contact
    
   
    
    return params
    
    
}

func createTrunkParams() swagger.CreateTrunkParams{
    var params swagger.CreateTrunkParams
    
    params.Name = "Test Trrunk"
    params.Url = "SIP/1234@phone.com:5060"

    
    return params
    
    
}


func replaceExtensionParams() swagger.ReplaceExtensionParams {
	var params swagger.ReplaceExtensionParams
	//~ params.Name = "extensionName2"
	//~ params.Timezone = "Europe/Skopje"
	//~ params.IncludeInDirectory = false
	//~ params.Extension = 15524
	//~ params.EnableOutboundCalls = false
	//~ params.UsageType = "extensionUsageType1" ###
	//~ params.VoicemailPassword = 222
	//~ params.FullName = "extensionFullName2"
	//~ params.EnableCallWaiting = false
	//~ params.VoicemailGreetingType = "extensionVoicemailGreetingType2"
	//~ params.CallerId = "extensionCallerId1" ###
	//~ params.LocalAreaCode = 620
	//~ params.VoicemailEnabled = false
	//~ params.VoicemailGreetingEnableLeaveMessagePrompt = false
	//~ params.VoicemailTranscription = "extensionVoicemailTranscription2"
	//~ params.VoicemailNotificationsEmails = []string{"extensionVoicemailNotificationsEmails3", "extensionVoicemailNotificationsEmails4"}
	//~ params.VoicemailNotificationsSms = "extensionVoicemailNotificationsSms2"
	//~ params.CallNotificationsEmails = []string{"extensionCallNotificationsEmails3", "extensionCallNotificationsEmails4"}
	//~ params.CallNotificationsSms = "extensionCallNotificationsSms2"
	//~ params.Route = []string{"extensionRoute3", "extensionRoute4"}

	return params
}

func createContactParams() swagger.CreateContactParams {

	var params swagger.CreateContactParams
	params.FirstName = "contactFirstName"
	params.MiddleName = "contactMiddleName"
	params.LastName = "contactLastName"
	params.Prefix = "contactPrefix"
	params.PhoneticFirstName = "contactPhoneticFirstName"
	params.PhoneticMiddleName = "contactPhoneticMiddleName"
	params.PhoneticLastName = "contactPhoneticLastName"
	params.Suffix = "contactSuffix"
	params.Nickname = "contactNickname"
	params.Company = "contactCompany"
	params.Department = "contactDepartment"
	params.JobTitle = "contactJobTitle"

	return params
}

func createGroupParams() swagger.CreateGroupParams {

	var params swagger.CreateGroupParams
	params.Name = "groupName"

	return params
}

func createPhoneNumberParams() swagger.CreatePhoneNumberParams {

	var params swagger.CreatePhoneNumberParams
	//~ params.Name = "phonenumberName"
	//~ params.BlockIncoming = true
	//~ params.BlockAnonymous = true
	//~ params.CallerIdName = "Caller Id Name"
	//~ params.CallerIdType = "phonenumberCallerIdType"
	//~ params.SmsForwardingType = "phonenumberSmsForwardingType"
	//~ params.CallNotificationsEmails = []string{"phonenumberCallNotificationsEmails1", "phonenumberCallNotificationsEmails2"}
	//~ params.CallNotificationsSms = "phonenumberCallNotificationsSms"

	return params
}

func replacePhoneNumberParams() swagger.ReplacePhoneNumberParams {

	var params swagger.ReplacePhoneNumberParams
	params.Name = "The Name"
	params.BlockIncoming = false
	params.BlockAnonymous = false
	var callerId swagger.CallerIdFull
	callerId.Name = "The Caller Id Name"
	callerId.Type_ = "The Caller Id Type"
	params.CallerIdName = callerId.Name
	params.CallerIdType = callerId.Type_
	params.SmsForwardingType = "The Sms Forwarding Type"
	params.CallNotificationsEmails = []string{"The Call Notifications Emails 1", "The Call Notifications Emails 2"}
	params.CallNotificationsSms = "The Call Notifications Sms"

	return params
}