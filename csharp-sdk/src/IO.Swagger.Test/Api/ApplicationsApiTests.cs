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
    ///  Class for testing ApplicationsApi
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by Swagger Codegen.
    /// Please update the test case below to test the API endpoint.
    /// </remarks>
    [TestFixture]
    public class ApplicationsApiTests
    {
        private ApplicationsApi instance;

        /// <summary>
        /// Setup before each unit test
        /// </summary>
        [SetUp]
        public void Init()
        {
            instance = new ApplicationsApi();
        }

        /// <summary>
        /// Clean up after each unit test
        /// </summary>
        [TearDown]
        public void Cleanup()
        {

        }

        /// <summary>
        /// Test an instance of ApplicationsApi
        /// </summary>
        [Test]
        public void InstanceTest()
        {
            // TODO uncomment below to test 'IsInstanceOfType' ApplicationsApi
            //Assert.IsInstanceOfType(typeof(ApplicationsApi), instance, "instance is a ApplicationsApi");
        }

        
        /// <summary>
        /// Test GetAccountApplication
        /// </summary>
        [Test]
        public void GetAccountApplicationTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //int? accountId = null;
            //int? applicationId = null;
            //var response = instance.GetAccountApplication(accountId, applicationId);
            //Assert.IsInstanceOf<ApplicationFull> (response, "response is ApplicationFull");
        }
        
        /// <summary>
        /// Test ListAccountApplications
        /// </summary>
        [Test]
        public void ListAccountApplicationsTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //int? accountId = null;
            //List<string> filtersId = null;
            //List<string> filtersName = null;
            //string sortId = null;
            //string sortName = null;
            //int? limit = null;
            //int? offset = null;
            //string fields = null;
            //var response = instance.ListAccountApplications(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
            //Assert.IsInstanceOf<ListApplicationsFull> (response, "response is ListApplicationsFull");
        }
        
    }

}
