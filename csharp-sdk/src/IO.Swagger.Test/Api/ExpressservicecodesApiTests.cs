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
    ///  Class for testing ExpressservicecodesApi
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by Swagger Codegen.
    /// Please update the test case below to test the API endpoint.
    /// </remarks>
    [TestFixture]
    public class ExpressservicecodesApiTests
    {
        private ExpressservicecodesApi instance;

        /// <summary>
        /// Setup before each unit test
        /// </summary>
        [SetUp]
        public void Init()
        {
            instance = new ExpressservicecodesApi();
        }

        /// <summary>
        /// Clean up after each unit test
        /// </summary>
        [TearDown]
        public void Cleanup()
        {

        }

        /// <summary>
        /// Test an instance of ExpressservicecodesApi
        /// </summary>
        [Test]
        public void InstanceTest()
        {
            // TODO uncomment below to test 'IsInstanceOfType' ExpressservicecodesApi
            //Assert.IsInstanceOfType(typeof(ExpressservicecodesApi), instance, "instance is a ExpressservicecodesApi");
        }

        
        /// <summary>
        /// Test GetAccountExpressSrvCode
        /// </summary>
        [Test]
        public void GetAccountExpressSrvCodeTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //int? accountId = null;
            //int? codeId = null;
            //var response = instance.GetAccountExpressSrvCode(accountId, codeId);
            //Assert.IsInstanceOf<ExpressServiceCodeFull> (response, "response is ExpressServiceCodeFull");
        }
        
        /// <summary>
        /// Test ListAccountExpressSrvCodes
        /// </summary>
        [Test]
        public void ListAccountExpressSrvCodesTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //int? accountId = null;
            //List<string> filtersId = null;
            //var response = instance.ListAccountExpressSrvCodes(accountId, filtersId);
            //Assert.IsInstanceOf<ListExpressServiceCodesFull> (response, "response is ListExpressServiceCodesFull");
        }
        
    }

}
