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

type CreateSmsParams struct {

	// Phone number of sender
	From string `json:"from,omitempty"`

	// Outgoing destination numbers
	To string `json:"to,omitempty"`

	// Text body of the outgoing SMS message. Maximum 160 characters per message.
	Text string `json:"text,omitempty"`

	Extension int32 `json:"extension,omitempty"`
}
