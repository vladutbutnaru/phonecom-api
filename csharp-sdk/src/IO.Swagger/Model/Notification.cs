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
    /// Notification
    /// </summary>
    [DataContract]
    public partial class Notification :  IEquatable<Notification>
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="Notification" /> class.
        /// </summary>
        /// <param name="Emails">Array of email addresses.</param>
        /// <param name="Sms">Phone number capable of receiving SMS messages.</param>
        public Notification(List<string> Emails = null, string Sms = null)
        {
            this.Emails = Emails;
            this.Sms = Sms;
        }
        
        /// <summary>
        /// Array of email addresses
        /// </summary>
        /// <value>Array of email addresses</value>
        [DataMember(Name="emails", EmitDefaultValue=false)]
        public List<string> Emails { get; set; }
        /// <summary>
        /// Phone number capable of receiving SMS messages
        /// </summary>
        /// <value>Phone number capable of receiving SMS messages</value>
        [DataMember(Name="sms", EmitDefaultValue=false)]
        public string Sms { get; set; }
        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class Notification {\n");
            sb.Append("  Emails: ").Append(Emails).Append("\n");
            sb.Append("  Sms: ").Append(Sms).Append("\n");
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
            return this.Equals(obj as Notification);
        }

        /// <summary>
        /// Returns true if Notification instances are equal
        /// </summary>
        /// <param name="other">Instance of Notification to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(Notification other)
        {
            // credit: http://stackoverflow.com/a/10454552/677735
            if (other == null)
                return false;

            return 
                (
                    this.Emails == other.Emails ||
                    this.Emails != null &&
                    this.Emails.SequenceEqual(other.Emails)
                ) && 
                (
                    this.Sms == other.Sms ||
                    this.Sms != null &&
                    this.Sms.Equals(other.Sms)
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
                if (this.Emails != null)
                    hash = hash * 59 + this.Emails.GetHashCode();
                if (this.Sms != null)
                    hash = hash * 59 + this.Sms.GetHashCode();
                return hash;
            }
        }
    }

}
