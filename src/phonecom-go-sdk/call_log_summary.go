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

// The Call Log Summary Object is used to briefly represent a call. It consists of properties as below:
type CallLogSummary struct {

	// ID
	Id string `json:"id,omitempty"`

	// Internal system id, may be null
	Uuid string `json:"uuid,omitempty"`

	// Account extension
	Extension ExtensionSummary `json:"extension,omitempty"`

	// Call made from this phone number
	CallerId string `json:"caller_id,omitempty"`

	// Call made to this phone number
	CalledNumber string `json:"called_number,omitempty"`

	// Call start time
	StartTime string `json:"start_time,omitempty"`

	// Call log creation time. Same as call end time + time needed to create call log
	CreatedAt string `json:"created_at,omitempty"`

	// Call direction: in or out
	Direction string `json:"direction,omitempty"`

	// Call type: call, fax, audiogram ...
	Type_ string `json:"type,omitempty"`

	// Call duration in seconds
	CallDuration int32 `json:"call_duration,omitempty"`

	// Was call being monitored?
	IsMonitored string `json:"is_monitored,omitempty"`

	// Internal system call reference number
	CallNumber string `json:"call_number,omitempty"`

	// Last action of call flow
	FinalAction string `json:"final_action,omitempty"`

	// URL of call recording if available. Empty string if call recording does not exist
	CallRecording string `json:"call_recording,omitempty"`
}