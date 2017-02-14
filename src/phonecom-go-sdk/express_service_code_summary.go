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

// The Express Service Code Summary Object is used to briefly represent a Express Service Code. It consists of the ID and code:
type ExpressServiceCodeSummary struct {

	// ID
	Id int32 `json:"id,omitempty"`

	// An 8-digit number in string format
	ExpressServiceCode string `json:"express_service_code,omitempty"`
}