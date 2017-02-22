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

import (
	"time"
)

// The Full SMS Object is identical to the SMS Summary Object. See above for details.
type SmsFull struct {

	// Unique SMS ID. Read-only.
	Id string `json:"id,omitempty"`

	// Caller ID number to display on the incoming/outgoing SMS message. For an outgoing message, it must be a phone number associated with your Phone.com account.
	From string `json:"from,omitempty"`

	// An array of SMS recipients.
	To []Recipient `json:"to,omitempty"`

	// Direction of SMS. 'in' for Incoming messages, 'out' for Outgoing messages.
	Direction string `json:"direction,omitempty"`

	// Unix time stamp representing the UTC time that the object was created in the Phone.com API system.
	CreatedEpoch int32 `json:"created_epoch,omitempty"`

	// Date string representing the UTC time that the object was created in the Phone.com API system.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Body of the SMS text
	Text string `json:"text,omitempty"`
}
