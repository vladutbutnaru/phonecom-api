=begin
#Phone.com API

#This is a Phone.com api Swagger definition

OpenAPI spec version: 1.0.0
Contact: apisupport@phone.com
Generated by: https://github.com/swagger-api/swagger-codegen.git

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

=end

require 'spec_helper'
require 'json'

# Unit tests for SwaggerClient::SubaccountsApi
# Automatically generated by swagger-codegen (github.com/swagger-api/swagger-codegen)
# Please update as you see appropriate
describe 'SubaccountsApi' do
  before do
    # run before each test
    @instance = SwaggerClient::SubaccountsApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of SubaccountsApi' do
    it 'should create an instact of SubaccountsApi' do
      expect(@instance).to be_instance_of(SwaggerClient::SubaccountsApi)
    end
  end

  # unit tests for create_account_subaccount
  # Add a subaccount for the authenticated user or client
  # This service shows the details of an individual Subaccount.
  # @param account_id Account ID
  # @param data SMS data
  # @param [Hash] opts the optional parameters
  # @return [AccountFull]
  describe 'create_account_subaccount test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_account_subaccounts
  # Get a list of subaccounts for the authenticated user or client
  # This service lists the Subaccount of the authenticated client. In most cases, there will not be any.
  # @param account_id Account ID
  # @param [Hash] opts the optional parameters
  # @option opts [Array<String>] :filters_id ID filter
  # @option opts [String] :sort_id ID sorting
  # @option opts [Integer] :limit Max results
  # @option opts [Integer] :offset Results to skip
  # @option opts [String] :fields Field set
  # @return [ListAccountsFull]
  describe 'list_account_subaccounts test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end