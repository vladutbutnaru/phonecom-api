# SWGMenuFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**_id** | **NSNumber*** | Integer Menu ID. Read-only. | [optional] 
**name** | **NSString*** | Name. Required. Unique. | [optional] 
**allowExtensionDial** | **NSNumber*** | Boolean. Determines whether a caller can enter an extension number to bypass the menu. | [optional] 
**keypressWaitTime** | **NSNumber*** | Boolean. Determines whether a caller can enter an extension number to bypass the menu. | [optional] 
**greeting** | [**SWGMediaSummary***](SWGMediaSummary.md) | Greeting that is played when a caller enters a menu. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. | [optional] 
**keypressError** | [**SWGMediaSummary***](SWGMediaSummary.md) | Message that is played when the caller makes a keypress error. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. | [optional] 
**timeoutHandler** | [**SWGRouteSummary***](SWGRouteSummary.md) | Route that will be entered when the caller fails to choose a menu option within the allotted time. Output is a Route Summary Object if the route is named, otherwise the Full Route Object will be shown. Input must be a Route Lookup Object pointing to a named route. | [optional] 
**options** | [**SWGOptionsListMenus***](SWGOptionsListMenus.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


