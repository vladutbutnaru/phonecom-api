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
    /// Line
    /// </summary>
    [DataContract]
    public partial class Line :  IEquatable<Line>
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="Line" /> class.
        /// </summary>
        /// <param name="_Line">Line number.</param>
        /// <param name="Extension">Extensions object that this line number is mapped to..</param>
        public Line(int? _Line = null, ExtensionSummary Extension = null)
        {
            this._Line = _Line;
            this.Extension = Extension;
        }
        
        /// <summary>
        /// Line number
        /// </summary>
        /// <value>Line number</value>
        [DataMember(Name="line", EmitDefaultValue=false)]
        public int? _Line { get; set; }
        /// <summary>
        /// Extensions object that this line number is mapped to.
        /// </summary>
        /// <value>Extensions object that this line number is mapped to.</value>
        [DataMember(Name="extension", EmitDefaultValue=false)]
        public ExtensionSummary Extension { get; set; }
        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class Line {\n");
            sb.Append("  _Line: ").Append(_Line).Append("\n");
            sb.Append("  Extension: ").Append(Extension).Append("\n");
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
            return this.Equals(obj as Line);
        }

        /// <summary>
        /// Returns true if Line instances are equal
        /// </summary>
        /// <param name="other">Instance of Line to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(Line other)
        {
            // credit: http://stackoverflow.com/a/10454552/677735
            if (other == null)
                return false;

            return 
                (
                    this._Line == other._Line ||
                    this._Line != null &&
                    this._Line.Equals(other._Line)
                ) && 
                (
                    this.Extension == other.Extension ||
                    this.Extension != null &&
                    this.Extension.Equals(other.Extension)
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
                if (this._Line != null)
                    hash = hash * 59 + this._Line.GetHashCode();
                if (this.Extension != null)
                    hash = hash * 59 + this.Extension.GetHashCode();
                return hash;
            }
        }
    }

}
