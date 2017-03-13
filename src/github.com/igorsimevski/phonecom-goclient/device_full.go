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

// The Full Device Object includes the properties in the Device Summary Object, along with the following:
type DeviceFull struct {

	// ID
	Id int32 `json:"id,omitempty"`

	// User-supplied name, otherwise NULL
	Name string `json:"name,omitempty"`

	SipAuthentication SipAuthentication `json:"sip_authentication,omitempty"`

	// Array of Line Objects showing which extensions are attached to this device, and their assigned line numbers. See below for details.
	Lines []Line `json:"lines,omitempty"`
}
