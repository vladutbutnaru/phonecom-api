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
  DeviceParamsJson := CreateDeviceJson{1315091, randomName}
  fileName := "../test/jsonin/createDevice" + randomName + ".json"
  b, err := json.Marshal(DeviceParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createDevice, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)

  if (id == 0) {
    t.FailNow()
  }

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
  DeviceParamsJson := ReplaceDeviceJson{1315091, int32(firstId), randomName}
  fileName := "../test/jsonin/replaceDevice" + randomName + ".json"
  b, err := json.Marshal(DeviceParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  t.Log(firstId)
  err, result = createReplaceCliWithJsonIn(replaceDevice, fileName, firstId)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

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
  ExtensionParamsJson := CreateExtensionJson{1315091, "+12019570328", "unlimited", true, int32(randomNum), true, "The name", "The full name", "America/Los_Angeles", 619, true, false, true, false, 12345, "standard", "automated", "+18587741111", "+18587748888"}
  fileName := "../test/jsonin/createExtension" + randomName + ".json"
  b, err := json.Marshal(ExtensionParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createExtension, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

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
  ExtensionParamsJson := ReplaceExtensionJson{1315091, int32(firstId), randomName, "America/Los_Angeles", true, 111, true, "unlimited", 12344, "bobby McFerrin", true, "standard", "private", 619, true, true, "automated", "+18587741111", "+18587748888"}
  fileName := "../test/jsonin/replaceExtension" + randomName + ".json"
  b, err := json.Marshal(ExtensionParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)



  err, result = createReplaceCliWithJsonIn(replaceExtension, fileName, firstId)
  os.Remove(fileName)
  assertErrorNotNull(t, err)
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
  ContactParamsJson := CreateContactJson{1315091, 1764590, "Geordi", "middle name", "last name", "prefix", "phoneticFirstName", "phoneticMiddleName", "phoneticLastName", "suffix", "nickname", "company", "department", "jobTitle"}
  fileName := "../test/jsonin/createContact" + randomName + ".json"
  b, err := json.Marshal(ContactParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)



  err, result = createCliWithJsonIn(createContact, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteContact, id);
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
  GroupParamsJson := CreateGroupJson{1315091, 1764590, "Ferengi Traders"}
  b, err := json.Marshal(GroupParamsJson)
  //d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"extension_id\": 1764590,\n\t\"name\": \"Ferengi Traders\"\n}")
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createCliWithJsonIn(createGroup, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteGroup, id);

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
  MenuParamsJson := CreateMenuJson{1315091, randomName, true, 3}
  fileName := "../test/jsonin/createMenu" + randomName + ".json"
  b, err := json.Marshal(MenuParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createCliWithJsonIn(createMenu, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteMenu, id);

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
  MenuParamsJson := ReplaceMenuJson{1315091, int32(firstId), randomName, false, 5}
  fileName := "../test/jsonin/replaceMenu" + randomName + ".json"
  b, err := json.Marshal(MenuParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createReplaceCliWithJsonIn(replaceMenu, fileName, firstId)
  os.Remove(fileName)
  assertErrorNotNull(t, err)
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


  randomName := randomString(12)
  PhoneNumberParamsJson := CreatePhoneNumberJson{1315091, "+12019570329", "Phone Name Now", true, true, "Phone N", "business", "extension", "+18587740222"}
  fileName := "../test/jsonin/createPhoneNumber" + randomName + ".json"
  b, err := json.Marshal(PhoneNumberParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createCliWithJsonIn(createPhoneNumber, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

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
  PhoneNumberParamsJson := ReplacePhoneNumberJson{1315091, int32(firstId), "Robert", true, true, "Phone N", "business", "extension", "+18587740222"}
  fileName := "../test/jsonin/replacePhoneNumber" + randomName + ".json"
  b, err := json.Marshal(PhoneNumberParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)



  err, result = createReplaceCliWithJsonIn(replacePhoneNumber, fileName, firstId)
  os.Remove(fileName)
  assertErrorNotNull(t, err)
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
  QueueParamsJson := CreateQueueJson{1315091, randomName, 60, "called_number", 10}
  fileName := "../test/jsonin/createQueue" + randomName + ".json"
  b, err := json.Marshal(QueueParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)



  err, result = createCliWithJsonIn(createQueue, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteQueue, id);

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
  QueueParamsJson := ReplaceQueueJson{1315091, int32(firstId), randomName, 60, "called_number", 10}
  fileName := "../test/jsonin/replaceQueue" + randomName + ".json"
  b, err := json.Marshal(QueueParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)



  err, result = createReplaceCliWithJsonIn(replaceQueue, fileName, firstId)
  os.Remove(fileName)
  assertErrorNotNull(t, err)
}

//type CreateRouteJson struct {
//  TODO
//}

func TestCreateDeleteRoute(t *testing.T) {

  var result map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/createRoute" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"name\": \"API V3 Test\",\n\t\"rules\": [\n\t\t{\n\t\t\t\"actions\": [\n\t\t\t\t{\n\t\t\t\t\t\"action\": \"queue\",\n\t\t\t\t\t\"queue\": {\n\t\t\t\t\t\t\"id\": 21971,\n\t\t\t\t\t\t\"name\": \"sample queue\"\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t]\n\t\t}\n\t]\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, result = createCliWithJsonIn(createRoute, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteRoute, id);

}

//type ReplaceRouteJson struct {
//  TODO
//}

func TestListReplaceRoute(t *testing.T) {

  var result map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/replaceRoute" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"route_id\":12313,\n\t\"name\": \"API V3 Test\",\n\t\"rules\": [\n\t\t{\n\t\t\t\"actions\": [\n\t\t\t\t{\n\t\t\t\t\t\"action\": \"queue\",\n\t\t\t\t\t\"queue\": {\n\t\t\t\t\t\t\"id\": 21971,\n\t\t\t\t\t\t\"name\": \"sample queue\"\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t]\n\t\t}\n\t]\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, result = createCliWithJsonIn(createRoute, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

}

type CreateSmsJson struct {
  AccountId int32 `json:"account_id"`
  From string `json:"from"`
  To string `json:"to"`
  Text string `json:"text"`
  ExtensionId int32 `json:"extension_id"`
}

func TestCreateSms(t *testing.T) {

  var result map[string] interface{}
  var err error

  randomName := randomString(12)
  SmsParamsJson := CreateSmsJson{1315091, "+16309624775", "+12019570328", "Sms work too", 1764595}
  fileName := "../test/jsonin/createSms" + randomName + ".json"
  b, err := json.Marshal(SmsParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createSms, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

}

type CreateSubaccountJson struct {
  AccountId int32 `json:"account_id"`
  Username string `json:"username"`
  Password string `json:"password"`
}

func TestCreateSubaccount(t *testing.T) {

  var result map[string] interface{}
  var err error


  randomName := randomString(12)
  randomPassword := randomString(12)
  SubaccountParamsJson := CreateSubaccountJson{1315091, randomName, randomPassword}
  fileName := "../test/jsonin/createSubaccount" + randomName + ".json"
  b, err := json.Marshal(SubaccountParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createCliWithJsonIn(createSubaccount, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

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
  TrunkParamsJson := CreateTrunkJson{1315091, "The trunk name", "SIP/1234@phone.com:5060", 60, 800}
  fileName := "../test/jsonin/createTrunk" + randomName + ".json"
  b, err := json.Marshal(TrunkParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createCliWithJsonIn(createTrunk, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteTrunk, id);

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
  TrunkParamsJson := ReplaceTrunkJson{1315091, int32(firstId), randomName, "SIP/1234@phone.com:5060", 60, 800}
  fileName := "../test/jsonin/createTrunk" + randomName + ".json"
  b, err := json.Marshal(TrunkParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createReplaceCliWithJsonIn(replaceTrunk, fileName, firstId)
  os.Remove(fileName)
  assertErrorNotNull(t, err)
}
