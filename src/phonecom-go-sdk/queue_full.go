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

// The Full QueueObject has the same properties as the Queue Summary Object, along with the following:
type QueueFull struct {

	// Integer ID. Read-only.
	Id int32 `json:"id,omitempty"`

	// Name. Required.
	Name string `json:"name,omitempty"`

	// Greeting to be played when caller first connects. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. Can be set to NULL to disable the greeting.
	Greeting MediaSummary `json:"greeting,omitempty"`

	HoldMusic HoldMusic `json:"hold_music,omitempty"`

	// Maximum hold time in seconds. If provided, must equal one of: 60, 120, 180, 240, 300, 600, 900, 1200, 1800, 2700, 3600. Default is 300.
	MaxHoldTime int32 `json:"max_hold_time,omitempty"`

	// Caller id type to show members. If provided, must equal one of: 'called_number', 'calling_number'. Default is 'calling_number'.
	CallerIdType string `json:"caller_id_type,omitempty"`

	// Number of seconds to ring a member before cycling to the next member. If provided, must equal one of: 5, 10, 15, 20, 25, 30. Default is 5.
	RingTime int32 `json:"ring_time,omitempty"`

	Members Members `json:"members,omitempty"`
}
