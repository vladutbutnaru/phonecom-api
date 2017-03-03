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
    ///  Class for testing SubaccountsApi
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by Swagger Codegen.
    /// Please update the test case below to test the API endpoint.
    /// </remarks>
    [TestFixture]
    public class SubaccountsApiTests
    {
        private SubaccountsApi instance;

        /// <summary>
        /// Setup before each unit test
        /// </summary>
        [SetUp]
        public void Init()
        {
            instance = new SubaccountsApi();
        }

        /// <summary>
        /// Clean up after each unit test
        /// </summary>
        [TearDown]
        public void Cleanup()
        {

        }

        /// <summary>
        /// Test an instance of SubaccountsApi
        /// </summary>
        [Test]
        public void InstanceTest()
        {
            // TODO uncomment below to test 'IsInstanceOfType' SubaccountsApi
            //Assert.IsInstanceOfType(typeof(SubaccountsApi), instance, "instance is a SubaccountsApi");
        }

        
        /// <summary>
        /// Test CreateAccountSubaccount
        /// </summary>
        [Test]
        public void CreateAccountSubaccountTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //int? accountId = null;
            //CreateSubaccountParams data = null;
            //var response = instance.CreateAccountSubaccount(accountId, data);
            //Assert.IsInstanceOf<AccountFull> (response, "response is AccountFull");
        }
        
        /// <summary>
        /// Test ListAccountSubaccounts
        /// </summary>
        [Test]
        public void ListAccountSubaccountsTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //int? accountId = null;
            //List<string> filtersId = null;
            //string sortId = null;
            //int? limit = null;
            //int? offset = null;
            //string fields = null;
            //var response = instance.ListAccountSubaccounts(accountId, filtersId, sortId, limit, offset, fields);
            //Assert.IsInstanceOf<ListAccountsFull> (response, "response is ListAccountsFull");
        }
        
    }

}
