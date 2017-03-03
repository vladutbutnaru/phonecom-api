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
using System.IO;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Reflection;
using RestSharp;
using NUnit.Framework;

using IO.Swagger.Client;
using IO.Swagger.Api;
using IO.Swagger.Model;

namespace IO.Swagger.Test
{
    /// <summary>
    ///  Class for testing ContactsApi
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by Swagger Codegen.
    /// Please update the test case below to test the API endpoint.
    /// </remarks>
    [TestFixture]
    public class ContactsApiTests
    {
        private ContactsApi instance;

        /// <summary>
        /// Setup before each unit test
        /// </summary>
        [SetUp]
        public void Init()
        {
            instance = new ContactsApi();
        }

        /// <summary>
        /// Clean up after each unit test
        /// </summary>
        [TearDown]
        public void Cleanup()
        {

        }

        /// <summary>
        /// Test an instance of ContactsApi
        /// </summary>
        [Test]
        public void InstanceTest()
        {
            // TODO uncomment below to test 'IsInstanceOfType' ContactsApi
            //Assert.IsInstanceOfType(typeof(ContactsApi), instance, "instance is a ContactsApi");
        }

        
        /// <summary>
        /// Test CreateAccountExtensionContact
        /// </summary>
        [Test]
        public void CreateAccountExtensionContactTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //int? accountId = null;
            //int? extensionId = null;
            //CreateContactParams data = null;
            //var response = instance.CreateAccountExtensionContact(accountId, extensionId, data);
            //Assert.IsInstanceOf<ContactFull> (response, "response is ContactFull");
        }
        
        /// <summary>
        /// Test DeleteAccountExtensionContact
        /// </summary>
        [Test]
        public void DeleteAccountExtensionContactTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //int? accountId = null;
            //int? extensionId = null;
            //int? contactId = null;
            //var response = instance.DeleteAccountExtensionContact(accountId, extensionId, contactId);
            //Assert.IsInstanceOf<DeleteContact> (response, "response is DeleteContact");
        }
        
        /// <summary>
        /// Test GetAccountExtensionContact
        /// </summary>
        [Test]
        public void GetAccountExtensionContactTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //int? accountId = null;
            //int? extensionId = null;
            //int? contactId = null;
            //var response = instance.GetAccountExtensionContact(accountId, extensionId, contactId);
            //Assert.IsInstanceOf<ContactFull> (response, "response is ContactFull");
        }
        
        /// <summary>
        /// Test ListAccountExtensionContacts
        /// </summary>
        [Test]
        public void ListAccountExtensionContactsTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //int? accountId = null;
            //int? extensionId = null;
            //List<string> filtersId = null;
            //List<string> filtersGroupId = null;
            //List<string> filtersUpdatedAt = null;
            //string sortId = null;
            //string sortUpdatedAt = null;
            //int? limit = null;
            //int? offset = null;
            //string fields = null;
            //var response = instance.ListAccountExtensionContacts(accountId, extensionId, filtersId, filtersGroupId, filtersUpdatedAt, sortId, sortUpdatedAt, limit, offset, fields);
            //Assert.IsInstanceOf<ListContactsFull> (response, "response is ListContactsFull");
        }
        
        /// <summary>
        /// Test ReplaceAccountExtensionContact
        /// </summary>
        [Test]
        public void ReplaceAccountExtensionContactTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //int? accountId = null;
            //int? extensionId = null;
            //CreateContactParams data = null;
            //var response = instance.ReplaceAccountExtensionContact(accountId, extensionId, data);
            //Assert.IsInstanceOf<ContactFull> (response, "response is ContactFull");
        }
        
    }

}
