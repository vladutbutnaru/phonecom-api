# IO.Swagger.Model.AvailableNumberFull
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | Phone number, in E.164 format | [optional] 
**Formatted** | **string** | Human-readable formatted version of the phone number | [optional] 
**Price** | **int?** | The one-time initial price for this number, in USD. Some numbers show REQUEST_QUOTE here. Please contact our sales department if you are interested in these numbers. | [optional] 
**IsTollFree** | **bool?** | Whether the number is toll-free | [optional] 
**CountryCode** | **string** | The international dialing prefix for this number. For US and Canadian numbers, for example, this will be \&quot;1\&quot;. | [optional] 
**Npa** | **string** | Area code (a.k.a. NPA). Included for North American numbers only. | [optional] 
**Nxx** | **string** | Second 3 digits (a.k.a. NXX). Included for North American numbers only. | [optional] 
**Xxxx** | **string** | Last 4 digits (a.k.a. XXXX). Included for North American numbers only. | [optional] 
**City** | **string** | City with which this number is associated, if known. Otherwise NULL. | [optional] 
**Province** | **string** | State or Province with which this number is associated, if known. Postal Code. Otherwise NULL. | [optional] 
**Country** | **string** | Country with which this number is associated, if known. Otherwise NULL. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
