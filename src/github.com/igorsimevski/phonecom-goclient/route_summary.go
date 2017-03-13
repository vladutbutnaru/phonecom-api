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

// The Route Summary Object is used to briefly represent a route. It can be seen in several places throughout this API. Here are the properties:
type RouteSummary struct {

	// Integer ID. Read-only.
	Id int32 `json:"id,omitempty"`

	// Name
	Name string `json:"name,omitempty"`
}
