package main

import (
  "testing"
  "encoding/json"
  "time"
  "math/rand"
  "io/ioutil"
  "os"
)

type CreateDeviceJson struct {
  AccountId int32 `json:"account_id"`
  Name string `json:"name"`
}

func TestCreateDevice(t *testing.T) {

  var result map[string] interface{}
  var err error

  randomName := randomString(12)
  DeviceParamsJson := CreateDeviceJson{defaultAccountId, randomName}
  fileName := "../test/jsonin/createDevice" + randomName + ".json"
  b, err := json.Marshal(DeviceParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createDevice, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)

  if (id == 0) {
    t.FailNow()
  }

  defer os.Remove(fileName)
}

func randomString(strlen int) string {

  rand.Seed(time.Now().UTC().UnixNano())
  const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
  result := make([]byte, strlen)
  for i := 0; i < strlen; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  return string(result)
}

func randomNumber(min int, max int) int {

  rand.Seed(time.Now().UTC().UnixNano())
  return int(rand.Float64()*(float64(max) - float64(min)) + float64(min))
}

type ReplaceDeviceJson struct {
  AccountId int32 `json:"account_id"`
  DeviceId int32 `json:"device_id"`
  Name string `json:"name"`
}

func TestListReplaceDevice(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listDevices)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)


  assertErrorNotNull(t, err)

  randomName := randomString(12)
  DeviceParamsJson := ReplaceDeviceJson{defaultAccountId, int32(firstId), randomName}
  fileName := "../test/jsonin/replaceDevice" + randomName + ".json"
  b, err := json.Marshal(DeviceParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  t.Log(firstId)
  err, result = createReplaceCliWithJsonIn(replaceDevice, fileName, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

type CreateExtensionJson struct {
  AccountId int32 `json:"account_id"`
  CallerId string `json:"caller_id"`
  UsageType string `json:"usage_type"`
  AllowsCallWaiting bool `json:"allows_call_waiting"`
  Extension int32 `json:extension`
  IncludeInDirectory bool `json:"include_in_directory"`
  Name string `json:"name"`
  FullName string `json:"full_name"`
  Timezone string `json:"timezone"`
  LocalAreaCode int32 `json:"local_area_code"`
  VoicemailGreetingEnableLeaveMessagePrompt bool `json:"voicemail[greeting][enable_leave_message_prompt]"`
  VoicemailEnabled bool `json:"voicemail[enabled]"`
  EnableOutboundCalls bool `json:"enable_outbound_calls"`
  EnableCallWaiting bool `json:"enable_call_waiting"`
  VoicemailPassword int32 `json:"voicemail[password]"`
  VoicemailGreetingType string `json:"voicemail[greeting][type]"`
  VoicemailTranscription string `json:"voicemail[transcription]"`
  VoicemailNotificationsSms string `json:"voicemail[notifications][sms]"`
  CallNotificationsSms string `json:"call_notifications[sms]"`
}

func TestCreateExtension(t *testing.T) {

  var result map[string] interface{}
  var err error

  randomName := randomString(12)
  randomNum := randomNumber(10, 9999999)
  ExtensionParamsJson := CreateExtensionJson{defaultAccountId, "+12019570328", "unlimited", true, int32(randomNum), true, "The name", "The full name", "America/Los_Angeles", 619, true, false, true, false, 12345, "standard", "automated", "+18587741111", "+18587748888"}
  fileName := "../test/jsonin/createExtension" + randomName + ".json"
  b, err := json.Marshal(ExtensionParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createExtension, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }
  os.Remove(fileName)
}

type ReplaceExtensionJson struct {
  AccountId int32 `json:"account_id"`
  ExtensionId int32 `json:"extension_id"`
  Name string `json:"name"`
  Timezone string `json:"timezone"`
  IncludeInDirectory bool `json:"include_in_directory"`
  Extension int32 `json:"extension"`
  EnableOutboundCalls bool `json:"enable_outbound_calls"`
  UsageType string `json:"usage_type"`
  VoicemailPassword int32 `json:"voicemail[password]"`
  FullName string `json:"full_name"`
  EnableCallWaiting bool `json:"enable_call_waiting"`
  VoicemailGreetingType string `json:"voicemail[greeting][type]"`
  CallerId string `json:"caller_id"`
  LocalAreaCode int32 `json:"local_area_code"`
  VoicemailEnabled bool `json:"voicemail[enabled]"`
  VoicemailGreetingEnableLeaveMessagePrompt bool `json:"voicemail[greeting][enable_leave_message_prompt]"`
  VoicemailTranscription string `json:"voicemail[transcription]"`
  VoicemailNotificationsSms string `json:"voicemail[notifications][sms]"`
  CallNotificationsSms string `json:"call_notifications[sms]"`
}

func TestListReplaceExtension(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listExtensions)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)


  randomName := randomString(12)
  ExtensionParamsJson := ReplaceExtensionJson{defaultAccountId, int32(firstId), randomName, "America/Los_Angeles", true, 111, true, "unlimited", 12344, "bobby McFerrin", true, "standard", "private", 619, true, true, "automated", "+18587741111", "+18587748888"}
  fileName := "../test/jsonin/replaceExtension" + randomName + ".json"
  b, err := json.Marshal(ExtensionParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)



  err, result = createReplaceCliWithJsonIn(replaceExtension, fileName, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

type CreateContactJson struct {
  AccountId int32 `json:"account_id"`
  ExtensionId int32 `json:"extension_id"`
  FirstName string `json:"first_name"`
  MiddleName string `json:"middle_name"`
  LastName string `json:"last_name"`
  Prefix string `json:"prefix"`
  PhoneticFirstName string `json:"phonetic_first_name"`
  PhoneticMiddleName string `json:"phonetic_middle_name"`
  PhoneticLastName string `json:"phonetic_last_name"`
  Suffix string `json:"suffix"`
  Nickname string `json:"nickname"`
  Company string `json:"company"`
  Department string `json:"department"`
  JobTitle string `json:"job_title"`
}

func TestCreateDeleteContact(t *testing.T) {

  var result map[string] interface{}
  var err error



  randomName := randomString(12)
  ContactParamsJson := CreateContactJson{defaultAccountId, 1764590, "Geordi", "middle name", "last name", "prefix", "phoneticFirstName", "phoneticMiddleName", "phoneticLastName", "suffix", "nickname", "company", "department", "jobTitle"}
  fileName := "../test/jsonin/createContact" + randomName + ".json"
  b, err := json.Marshal(ContactParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)



  err, result = createCliWithJsonIn(createContact, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteContact, id);
  os.Remove(fileName)
}

type CreateGroupJson struct {
  AccountId int32 `json:"account_id"`
  ExtensionId int32 `json:"extension_id"`
  Name string `json:"name"`
}

func TestCreateDeleteGroup(t *testing.T) {

  var result map[string] interface{}
  var err error


  randomName := randomString(12)
  fileName := "../test/jsonin/createGroup" + randomName + ".json"
  GroupParamsJson := CreateGroupJson{defaultAccountId, 1764590, "Ferengi Traders"}
  b, err := json.Marshal(GroupParamsJson)
  //d1 := []byte("{\n\t\"account_id\": defaultAccountId,\n\t\"extension_id\": 1764590,\n\t\"name\": \"Ferengi Traders\"\n}")
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createCliWithJsonIn(createGroup, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteGroup, id);
  os.Remove(fileName)
}

type CreateMenuJson struct {
  AccountId int32 `json:"account_id"`
  Name string `json:"name"`
  AllowExtensionDial bool `json:"allow_extension_dial"`
  KeypressWaitTime int32 `json:"keypress_wait_time"`
}

func TestCreateDeleteMenu(t *testing.T) {

  var result map[string] interface{}
  var err error


  randomName := randomString(12)
  MenuParamsJson := CreateMenuJson{defaultAccountId, randomName, true, 3}
  fileName := "../test/jsonin/createMenu" + randomName + ".json"
  b, err := json.Marshal(MenuParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createCliWithJsonIn(createMenu, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteMenu, id);
  os.Remove(fileName)
}

type ReplaceMenuJson struct {
  AccountId int32 `json:"account_id"`
  MenuId int32 `json:"menu_id"`
  Name string `json:"name"`
  AllowExtensionDial bool `json:"allow_extension_dial"`
  KeypressWaitTime int32 `json:"keypress_wait_time"`
}

func TestListReplaceMenu(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listMenus)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)


  randomName := randomString(12)
  MenuParamsJson := ReplaceMenuJson{defaultAccountId, int32(firstId), randomName, false, 5}
  fileName := "../test/jsonin/replaceMenu" + randomName + ".json"
  b, err := json.Marshal(MenuParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createReplaceCliWithJsonIn(replaceMenu, fileName, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

type CreatePhoneNumberJson struct {
  AccountId int32 `json:"account_id"`
  PhoneNumber string `json:"phone_number"`
  Name string `json:"name"`
  BlockIncoming bool `json:"block_incoming"`
  BlockAnonymous bool `json:"block_anonymous"`
  CallerIdName string `json:"caller_id[name]"`
  CallerIdType string `json:"caller_id[type]"`
  SmsForwardingType string `json:"sms_forwarding[type]"`
  CallNotificationsSms string `json:"call_notifications[sms]"`
}

func TestCreatePhoneNumber(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listAvailablePhoneNumbers)
  assertErrorNotNull(t, err)

  firstId := getFirstAvailablePhoneNumber(result)


  randomName := randomString(12)
  PhoneNumberParamsJson := CreatePhoneNumberJson{defaultAccountId, firstId, "Phone Name Now", true, true, "Phone N", "business", "extension", "+18587740222"}
  fileName := "../test/jsonin/createPhoneNumber" + randomName + ".json"
  b, err := json.Marshal(PhoneNumberParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createCliWithJsonIn(createPhoneNumber, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }
  os.Remove(fileName)
}

type ReplacePhoneNumberJson struct {
  AccountId int32 `json:"account_id"`
  NumberId int32 `json:"number_id"`
  Name string `json:"name"`
  BlockIncoming bool `json:"block_incoming"`
  BlockAnonymous bool `json:"block_anonymous"`
  CallerIdName string `json:"caller_id[name]"`
  CallerIdType string `json:"caller_id[type]"`
  SmsForwardingType string `json:"sms_forwarding[type]"`
  CallNotificationsSms string `json:"call_notifications[sms]"`
}

func TestListReplacePhoneNumber(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listPhoneNumbers)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)


  randomName := randomString(12)
  PhoneNumberParamsJson := ReplacePhoneNumberJson{defaultAccountId, int32(firstId), "Robert", true, true, "Phone N", "business", "extension", "+18587740222"}
  fileName := "../test/jsonin/replacePhoneNumber" + randomName + ".json"
  b, err := json.Marshal(PhoneNumberParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)



  err, result = createReplaceCliWithJsonIn(replacePhoneNumber, fileName, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

type CreateQueueJson struct {
  AccountId int32 `json:"account_id"`
  Name string `json:"name"`
  MaxHoldTime int32 `json:"max_hold_time"`
  CallerIdType string `json:"caller_id_type"`
  RingTime int32 `json:"ring_time"`
}

func TestCreateDeleteQueue(t *testing.T) {

  var result map[string] interface{}
  var err error

  randomName := randomString(12)
  QueueParamsJson := CreateQueueJson{defaultAccountId, randomName, 60, "called_number", 10}
  fileName := "../test/jsonin/createQueue" + randomName + ".json"
  b, err := json.Marshal(QueueParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createQueue, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteQueue, id);
  os.Remove(fileName)
}

type ReplaceQueueJson struct {
  AccountId int32 `json:"account_id"`
  QueueId int32 `json:"queue_id"`
  Name string `json:"name"`
  MaxHoldTime int32 `json:"max_hold_time"`
  CallerIdType string `json:"caller_id_type"`
  RingTime int32 `json:"ring_time"`
}

func TestListReplaceQueue(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listQueues)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)

  randomName := randomString(12)
  QueueParamsJson := ReplaceQueueJson{defaultAccountId, int32(firstId), randomName, 60, "called_number", 10}
  fileName := "../test/jsonin/replaceQueue" + randomName + ".json"
  b, err := json.Marshal(QueueParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)



  err, result = createReplaceCliWithJsonIn(replaceQueue, fileName, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

type CreateRouteJson struct {
  AccountId int32 `json:"account_id"`
  Name string `json:"name"`
  Rules []RulesJson `json:"rules"`
}

type RulesJson struct {
  Actions []ActionsJson `json:"actions"`
}

type ActionsJson struct {
  Action string `json:"action"`
  Queue QueueJson `json:"queue"`
}

type QueueJson struct {
  Id int32 `json:"id"`
  Name string `json:"name"`
} 

func TestCreateDeleteRoute(t *testing.T) {

  var result map[string] interface{}
  var err error

  randomNameQ := randomString(12)
  QueueParamsJson := CreateQueueJson{defaultAccountId, randomNameQ, 60, "called_number", 10}
  fileNameQ := "../test/jsonin/createQueue" + randomNameQ + ".json"
  bQ, errQ := json.Marshal(QueueParamsJson)
  err = ioutil.WriteFile(fileNameQ, bQ, 0644)

  errQ, result = createCliWithJsonIn(createQueue, fileNameQ)
  assertErrorNotNull(t, errQ)

  idQ := getId(result)
  nameQ := getName(result)
  if (idQ == 0) {
    t.FailNow()
  }



  queueParam := QueueJson{int32(idQ), nameQ}
  actionParam := ActionsJson{"queue", queueParam}
  ruleJson := RulesJson{[]ActionsJson{actionParam}}
  rulesParam := []RulesJson{ruleJson}
  randomName := randomString(12)
  RouteParamsJson := CreateRouteJson{defaultAccountId, randomName, rulesParam}
  fileName := "../test/jsonin/createRoute" + randomName + ".json"
  b, err := json.Marshal(RouteParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createRoute, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }



  createGetOrRemoveCli(deleteRoute, id);
  os.Remove(fileName)

  createGetOrRemoveCli(deleteQueue, idQ);
  os.Remove(fileNameQ)
}

type ReplaceRouteJson struct {
  AccountId int32 `json:"account_id"`
  RouteId int32 `json:"route_id"`
  Name string `json:"name"`
  Rules []RulesJson `json:"rules"`
}

func TestListReplaceRoute(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listRoutes)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)

  queueParam := QueueJson{22026, "61kkjklmin74"}
  actionParam := ActionsJson{"queue", queueParam}
  ruleJson := RulesJson{[]ActionsJson{actionParam}}
  rulesParam := []RulesJson{ruleJson}
  randomName := randomString(12)
  RouteParamsJson := ReplaceRouteJson{defaultAccountId, int32(firstId), randomName, rulesParam}
  fileName := "../test/jsonin/replaceRoute" + randomName + ".json"
  b, err := json.Marshal(RouteParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createReplaceCliWithJsonIn(replaceRoute, fileName, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

type CreateSmsJson struct {
  AccountId int32 `json:"account_id"`
  From string `json:"from"`
  To string `json:"to"`
  Text string `json:"text"`
  ExtensionId int32 `json:"extension_id"`
}

func TestCreateSms(t *testing.T) {

  var err error

  from := "+16309624775"
  to := "+12019570328"
  text := "Another message for create"
  randomName := randomString(12)
  SmsParamsJson := CreateSmsJson{defaultAccountId, from, to, text, 1767963}
  fileName := "../test/jsonin/createSms" + randomName + ".json"
  b, err := json.Marshal(SmsParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result := createCreateSmsCliWithJsonIn(createSms, fileName, from, to, text)
  if (result == nil) {
    t.Log("Create message response is nil")
  }

  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

type CreateSubaccountJson struct {
  AccountId int32 `json:"account_id"`
  Username string `json:"username"`
  Password string `json:"password"`
  Contact ContactJson `json:"contact"`
  BillingContact ContactJson `json:"billing_contact"`
}

type ContactJson struct {
  Name string `json:"name"`
  Address AddressJson `json:"address"`
  Phone string `json:"phone"`
  PrimaryEmail string `json:"primary_email"`
}

type AddressJson struct {
  Line1 string `json:"line_1"`
  City string `json:"city"`
  Province string `json:"province"`
  PostalCode string `json:"postal_code"`
  Country string `json:"country"`
}

func TestCreateSubaccount(t *testing.T) {

  var result map[string] interface{}
  var err error


  randomName := randomString(12)
  randomPassword := randomString(12)
  AddressObject1 := AddressJson{"100 Main St", "San Diego", "CA", "92129", "US"}
  contactObject := ContactJson{"Bobby", AddressObject1, "+18585553333", "asd@sd.co"}
  AddressObject2 := AddressJson{"100 Main St", "San Diego", "CA", "92129", "US"}
  billingContactObject := ContactJson{"Bobby", AddressObject2, "+18585553333", "asd@sd.co"}
  SubaccountParamsJson := CreateSubaccountJson{defaultAccountId, randomName, randomPassword, contactObject, billingContactObject}
  fileName := "../test/jsonin/createSubaccount" + randomName + ".json"
  b, err := json.Marshal(SubaccountParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createCliWithJsonIn(createSubaccount, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }
  os.Remove(fileName)
}

type CreateTrunkJson struct {
  AccountId int32 `json:"account_id"`
  Name string `json:"name"`
  Uri string `json:"uri"`
  MaxConcurrentCalls int32 `json:"max_concurrent_calls"`
  MaxMinutesPerMonth int32 `json:"max_minutes_per_month"`
}

func TestCreateDeleteTrunk(t *testing.T) {

  var result map[string] interface{}
  var err error


  randomName := randomString(12)
  trunkName := randomName
  trunkUri := "SIP/1234@phone.com:5060"
  trunkConcurrentCalls := 60
  trunkMaxMinutes := 800
  TrunkParamsJson := CreateTrunkJson{defaultAccountId, trunkName, trunkUri, int32(trunkConcurrentCalls), int32(trunkMaxMinutes)}
  fileName := "../test/jsonin/createTrunk" + randomName + ".json"
  b, err := json.Marshal(TrunkParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createCreateTrunkCliWithJsonIn(createTrunk, fileName, trunkName, trunkUri, int32(trunkConcurrentCalls), int32(trunkMaxMinutes))
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteTrunk, id);
  os.Remove(fileName)
}

type ReplaceTrunkJson struct {
  AccountId int32 `json:"account_id"`
  TrunkId int32 `json:"trunk_id"`
  Name string `json:"name"`
  Uri string `json:"uri"`
  MaxConcurrentCalls int32 `json:"max_concurrent_calls"`
  MaxMinutesPerMonth int32 `json:"max_minutes_per_month"`
}

func TestListReplaceTrunk(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listTrunks)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)

  randomName := randomString(12)
  trunkName := randomName
  trunkUri := "SIP/1234@phone.com:5060"
  trunkConcurrentCalls := 80
  trunkMaxMinutes := 800
  TrunkParamsJson := ReplaceTrunkJson{defaultAccountId, int32(firstId), trunkName, trunkUri, int32(trunkConcurrentCalls), int32(trunkMaxMinutes)}
  fileName := "../test/jsonin/createTrunk" + randomName + ".json"
  b, err := json.Marshal(TrunkParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createReplaceTrunkCliWithJsonIn(replaceTrunk, fileName, trunkName, trunkUri, int32(trunkConcurrentCalls), int32(trunkMaxMinutes), firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}
