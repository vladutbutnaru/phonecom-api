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

type ListQueues struct {

	Filters FilterIdNameArray `json:"filters,omitempty"`

	Sort SortIdName `json:"sort,omitempty"`

	Total int32 `json:"total,omitempty"`

	Offset int32 `json:"offset,omitempty"`

	Limit int32 `json:"limit,omitempty"`

	Items []QueueFull `json:"items,omitempty"`
}
