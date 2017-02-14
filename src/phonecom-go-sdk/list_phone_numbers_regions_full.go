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

type ListPhoneNumbersRegionsFull struct {

	GroupBy GroupByListPhoneNumbersRegions `json:"group_by,omitempty"`

	Filters FilterListPhoneNumbersRegions `json:"filters,omitempty"`

	Sort SortListPhoneNumbersRegions `json:"sort,omitempty"`

	Total int32 `json:"total,omitempty"`

	Offset int32 `json:"offset,omitempty"`

	Limit int32 `json:"limit,omitempty"`

	Items PhoneNumbersRegionsFull `json:"items,omitempty"`
}
