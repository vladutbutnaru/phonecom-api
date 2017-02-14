/* 
 * Phone.com API
 *
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type ReplaceExtensionParams struct {

	// Recording lookup object
	VoicemailGreetingAlternate interface{} `json:"voicemail[greeting][alternate],omitempty"`

	// Recording lookup object
	NameGreeting interface{} `json:"name_greeting,omitempty"`

	// Name (required)
	Name string `json:"name,omitempty"`

	// Timezone
	Timezone string `json:"timezone,omitempty"`

	// Include in dial-by-name directory
	IncludeInDirectory bool `json:"include_in_directory,omitempty"`

	// Extension number (required)
	Extension int32 `json:"extension,omitempty"`

	// Enable outgoing calls
	EnableOutboundCalls bool `json:"enable_outbound_calls,omitempty"`

	// Extension type
	UsageType string `json:"usage_type,omitempty"`

	// Voicemail password
	VoicemailPassword int32 `json:"voicemail[password],omitempty"`

	// Contact name
	FullName string `json:"full_name,omitempty"`

	// Enable Call Waiting
	EnableCallWaiting bool `json:"enable_call_waiting,omitempty"`

	// Recording lookup object
	VoicemailGreetingStandard interface{} `json:"voicemail[greeting][standard],omitempty"`

	// Voicemail greeting type
	VoicemailGreetingType string `json:"voicemail[greeting][type],omitempty"`

	// Caller ID
	CallerId string `json:"caller_id,omitempty"`

	// Local area code
	LocalAreaCode int32 `json:"local_area_code,omitempty"`

	// Voicemail enabled
	VoicemailEnabled bool `json:"voicemail[enabled],omitempty"`

	// Use leave message prompt after voicemail
	VoicemailGreetingEnableLeaveMessagePrompt bool `json:"voicemail[greeting][enable_leave_message_prompt],omitempty"`

	// Voicemail transcription type
	VoicemailTranscription string `json:"voicemail[transcription],omitempty"`

	// Email notifications for voicemails. Can be a single email or an array of emails
	VoicemailNotificationsEmails []string `json:"voicemail[notifications][emails],omitempty"`

	// SMS notifications for voicemails
	VoicemailNotificationsSms string `json:"voicemail[notifications][sms],omitempty"`

	// Email notifications for calls. Can be a single email or an array of emails
	CallNotificationsEmails []string `json:"call_notifications[emails],omitempty"`

	// SMS notifications for calls
	CallNotificationsSms string `json:"call_notifications[sms],omitempty"`

	// Route object lookup (must belong to this extension)
	Route []string `json:"route,omitempty"`
}
