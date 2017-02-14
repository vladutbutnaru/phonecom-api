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

type CreateTrunkParams struct {

	// Name of Trunk
	Name string `json:"name,omitempty"`

	// URI of Trunk (in the form of SIP/user@host.com:port)
	Url string `json:"url,omitempty"`

	// Max concurrent calls
	MaxConcurrentCalls int32 `json:"max_concurrent_calls,omitempty"`

	// Max minutes per month
	MaxMinutesPerMonth int32 `json:"max_minutes_per_month,omitempty"`

	// Recording lookup object
	Greeting Object `json:"greeting,omitempty"`

	// Recording lookup object
	ErrorMessage Object `json:"error_message,omitempty"`

	// Custom audio codec configuration
	Codecs []Object `json:"codecs,omitempty"`
}
