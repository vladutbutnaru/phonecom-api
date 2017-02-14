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

// The Full Express Service Code Object includes the properties in the Express Service Code Summary Object, along with the following:
type ExpressServiceCodeFull struct {

	// ID
	Id int32 `json:"id,omitempty"`

	// An 8-digit number in string format
	ExpressServiceCode string `json:"express_service_code,omitempty"`

	// UNIX time stamp representing the UTC time that the Express Service Code expires. Please note that every time this service is executed, the expire_date is set to now + 24 hours.
	ExpireDate int32 `json:"expire_date,omitempty"`
}
