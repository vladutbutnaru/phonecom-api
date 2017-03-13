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

// The Contact Summary Object is used to briefly represent a contact from your address book. It can be seen in several places throughout this API. Here are the properties:
type ContactSummary struct {

	// Integer ID. Read-only.
	Id int32 `json:"id,omitempty"`

	// Salutation, such as Mr, Mrs, or Dr
	Prefix string `json:"prefix,omitempty"`

	// First name or given name
	FirstName string `json:"first_name,omitempty"`

	// Middle or second name
	MiddleName string `json:"middle_name,omitempty"`

	// Last name or surname
	LastName string `json:"last_name,omitempty"`

	// Suffix, such as 'Jr.', 'Sr.', 'II', or 'III'
	Suffix string `json:"suffix,omitempty"`

	// Nickname, or a shortened informal version of the contact's name
	Nickname string `json:"nickname,omitempty"`

	// Name of the contact's company
	Company string `json:"company,omitempty"`
}
