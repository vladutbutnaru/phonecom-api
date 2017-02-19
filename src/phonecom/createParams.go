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