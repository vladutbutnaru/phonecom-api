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

using System;
using System.Linq;
using System.IO;
using System.Text;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Runtime.Serialization;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;

namespace IO.Swagger.Model
{
    /// <summary>
    /// The Contact Summary Object is used to briefly represent a contact from your address book. It can be seen in several places throughout this API. Here are the properties:
    /// </summary>
    [DataContract]
    public partial class ContactSummary :  IEquatable<ContactSummary>
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ContactSummary" /> class.
        /// </summary>
        /// <param name="Id">Integer ID. Read-only..</param>
        /// <param name="Prefix">Salutation, such as Mr, Mrs, or Dr.</param>
        /// <param name="FirstName">First name or given name.</param>
        /// <param name="MiddleName">Middle or second name.</param>
        /// <param name="LastName">Last name or surname.</param>
        /// <param name="Suffix">Suffix, such as &#39;Jr.&#39;, &#39;Sr.&#39;, &#39;II&#39;, or &#39;III&#39;.</param>
        /// <param name="Nickname">Nickname, or a shortened informal version of the contact&#39;s name.</param>
        /// <param name="Company">Name of the contact&#39;s company.</param>
        public ContactSummary(int? Id = null, string Prefix = null, string FirstName = null, string MiddleName = null, string LastName = null, string Suffix = null, string Nickname = null, string Company = null)
        {
            this.Id = Id;
            this.Prefix = Prefix;
            this.FirstName = FirstName;
            this.MiddleName = MiddleName;
            this.LastName = LastName;
            this.Suffix = Suffix;
            this.Nickname = Nickname;
            this.Company = Company;
        }
        
        /// <summary>
        /// Integer ID. Read-only.
        /// </summary>
        /// <value>Integer ID. Read-only.</value>
        [DataMember(Name="id", EmitDefaultValue=false)]
        public int? Id { get; set; }
        /// <summary>
        /// Salutation, such as Mr, Mrs, or Dr
        /// </summary>
        /// <value>Salutation, such as Mr, Mrs, or Dr</value>
        [DataMember(Name="prefix", EmitDefaultValue=false)]
        public string Prefix { get; set; }
        /// <summary>
        /// First name or given name
        /// </summary>
        /// <value>First name or given name</value>
        [DataMember(Name="first_name", EmitDefaultValue=false)]
        public string FirstName { get; set; }
        /// <summary>
        /// Middle or second name
        /// </summary>
        /// <value>Middle or second name</value>
        [DataMember(Name="middle_name", EmitDefaultValue=false)]
        public string MiddleName { get; set; }
        /// <summary>
        /// Last name or surname
        /// </summary>
        /// <value>Last name or surname</value>
        [DataMember(Name="last_name", EmitDefaultValue=false)]
        public string LastName { get; set; }
        /// <summary>
        /// Suffix, such as &#39;Jr.&#39;, &#39;Sr.&#39;, &#39;II&#39;, or &#39;III&#39;
        /// </summary>
        /// <value>Suffix, such as &#39;Jr.&#39;, &#39;Sr.&#39;, &#39;II&#39;, or &#39;III&#39;</value>
        [DataMember(Name="suffix", EmitDefaultValue=false)]
        public string Suffix { get; set; }
        /// <summary>
        /// Nickname, or a shortened informal version of the contact&#39;s name
        /// </summary>
        /// <value>Nickname, or a shortened informal version of the contact&#39;s name</value>
        [DataMember(Name="nickname", EmitDefaultValue=false)]
        public string Nickname { get; set; }
        /// <summary>
        /// Name of the contact&#39;s company
        /// </summary>
        /// <value>Name of the contact&#39;s company</value>
        [DataMember(Name="company", EmitDefaultValue=false)]
        public string Company { get; set; }
        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class ContactSummary {\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Prefix: ").Append(Prefix).Append("\n");
            sb.Append("  FirstName: ").Append(FirstName).Append("\n");
            sb.Append("  MiddleName: ").Append(MiddleName).Append("\n");
            sb.Append("  LastName: ").Append(LastName).Append("\n");
            sb.Append("  Suffix: ").Append(Suffix).Append("\n");
            sb.Append("  Nickname: ").Append(Nickname).Append("\n");
            sb.Append("  Company: ").Append(Company).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }
  
        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public string ToJson()
        {
            return JsonConvert.SerializeObject(this, Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="obj">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object obj)
        {
            // credit: http://stackoverflow.com/a/10454552/677735
            return this.Equals(obj as ContactSummary);
        }

        /// <summary>
        /// Returns true if ContactSummary instances are equal
        /// </summary>
        /// <param name="other">Instance of ContactSummary to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ContactSummary other)
        {
            // credit: http://stackoverflow.com/a/10454552/677735
            if (other == null)
                return false;

            return 
                (
                    this.Id == other.Id ||
                    this.Id != null &&
                    this.Id.Equals(other.Id)
                ) && 
                (
                    this.Prefix == other.Prefix ||
                    this.Prefix != null &&
                    this.Prefix.Equals(other.Prefix)
                ) && 
                (
                    this.FirstName == other.FirstName ||
                    this.FirstName != null &&
                    this.FirstName.Equals(other.FirstName)
                ) && 
                (
                    this.MiddleName == other.MiddleName ||
                    this.MiddleName != null &&
                    this.MiddleName.Equals(other.MiddleName)
                ) && 
                (
                    this.LastName == other.LastName ||
                    this.LastName != null &&
                    this.LastName.Equals(other.LastName)
                ) && 
                (
                    this.Suffix == other.Suffix ||
                    this.Suffix != null &&
                    this.Suffix.Equals(other.Suffix)
                ) && 
                (
                    this.Nickname == other.Nickname ||
                    this.Nickname != null &&
                    this.Nickname.Equals(other.Nickname)
                ) && 
                (
                    this.Company == other.Company ||
                    this.Company != null &&
                    this.Company.Equals(other.Company)
                );
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            // credit: http://stackoverflow.com/a/263416/677735
            unchecked // Overflow is fine, just wrap
            {
                int hash = 41;
                // Suitable nullity checks etc, of course :)
                if (this.Id != null)
                    hash = hash * 59 + this.Id.GetHashCode();
                if (this.Prefix != null)
                    hash = hash * 59 + this.Prefix.GetHashCode();
                if (this.FirstName != null)
                    hash = hash * 59 + this.FirstName.GetHashCode();
                if (this.MiddleName != null)
                    hash = hash * 59 + this.MiddleName.GetHashCode();
                if (this.LastName != null)
                    hash = hash * 59 + this.LastName.GetHashCode();
                if (this.Suffix != null)
                    hash = hash * 59 + this.Suffix.GetHashCode();
                if (this.Nickname != null)
                    hash = hash * 59 + this.Nickname.GetHashCode();
                if (this.Company != null)
                    hash = hash * 59 + this.Company.GetHashCode();
                return hash;
            }
        }
    }

}