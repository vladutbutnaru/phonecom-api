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

// The Full Contact Object includes all of the properties in the Contact Summary Object, along with several more:
type ContactFull struct {

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

	// Suffix, such as \"Jr.\", \"Sr.\", \"II\", or \"III\"
	Suffix string `json:"suffix,omitempty"`

	// Nickname, or a shortened informal version of the contact's name
	Nickname string `json:"nickname,omitempty"`

	// Name of the contact's company
	Company string `json:"company,omitempty"`

	// Phonetic first name. Useful for remembering how to pronounce the contact's name.
	PhoneticFirstName string `json:"phonetic_first_name,omitempty"`

	// Phonetic middle name. Useful for remembering how to pronounce the contact's name.
	PhoneticMiddleName string `json:"phonetic_middle_name,omitempty"`

	// Phonetic last name. Useful for remembering how to pronounce the contact's name.
	PhoneticLastName string `json:"phonetic_last_name,omitempty"`

	// Name of the contact's department
	Department string `json:"department,omitempty"`

	// Contact's job title
	JobTitle string `json:"job_title,omitempty"`

	Emails Emails `json:"emails,omitempty"`

	PhoneNumbers []PhoneNumberContact `json:"phone_numbers,omitempty"`

	Addresses Addresses `json:"addresses,omitempty"`

	Group GroupListContacts `json:"group,omitempty"`

	// Integer UNIX timestamp when the contact was created. Read-only.
	CreatedAt int32 `json:"created_at,omitempty"`

	// Integer UNIX timestamp when the contact was created. Read-only.
	UpdatedAt int32 `json:"updated_at,omitempty"`
}