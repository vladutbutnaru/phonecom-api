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

// The Full Recording Object includes all of the properties from the Recording Summary Object, along with the following:
type MediaFull struct {

	// Recording ID. Read-only.
	Id int32 `json:"id,omitempty"`

	// Name of recording
	Name string `json:"name,omitempty"`

	// Can be hold_music or greeting. Indicates the purpose of this recording and where it can be used.
	Type_ string `json:"type,omitempty"`
}
