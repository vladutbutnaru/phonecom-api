/* 
 * Phone.com API
 *
 * This is a Phone.com api Swagger definition
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apisupport@phone.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

// Each item in the 'details' record may contain the following properties:
type CallDetails struct {

	// UNIX time stamp representing the UTC time that this call item starts
	StartTime int32 `json:"start_time,omitempty"`

	// Endpoint for this call item, such as call, voicemail, fax, menu, menu item, queue ...
	Type_ string `json:"type,omitempty"`

	// ID associated with this call item endpoint (type)
	IdValue int32 `json:"id_value,omitempty"`

	// Same as account id
	VoipId int32 `json:"voip_id,omitempty"`

	// Same as account extension id
	VoipPhoneId int32 `json:"voip_phone_id,omitempty"`
}
