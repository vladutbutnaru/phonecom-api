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

type ListExpressServiceCodesSummary struct {

	Filters FilterIdArray `json:"filters,omitempty"`

	Sort SortIdArray `json:"sort,omitempty"`

	Total int32 `json:"total,omitempty"`

	Offset int32 `json:"offset,omitempty"`

	Limit int32 `json:"limit,omitempty"`

	Items ExpressServiceCodesSummary `json:"items,omitempty"`
}
