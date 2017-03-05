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

# Unit tests for SwaggerClient::CalllogsApi
# Automatically generated by swagger-codegen (github.com/swagger-api/swagger-codegen)
# Please update as you see appropriate
describe 'CalllogsApi' do
  before do
    # run before each test
    @instance = SwaggerClient::CalllogsApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of CalllogsApi' do
    it 'should create an instact of CalllogsApi' do
      expect(@instance).to be_instance_of(SwaggerClient::CalllogsApi)
    end
  end

  # unit tests for get_account_call_log
  # Show details of an individual Call Log entry
  # This service shows the details of an individual Call Logs entry.
  # @param account_id Account ID
  # @param call_log_id Call Log ID
  # @param [Hash] opts the optional parameters
  # @return [CallLogFull]
  describe 'get_account_call_log test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_account_call_logs
  # Get a list of call details associated with your account
  # Fetch a list of call logs associated with your account.
  # @param account_id Account ID
  # @param [Hash] opts the optional parameters
  # @option opts [Array<String>] :filters_id ID filter
  # @option opts [Array<String>] :filters_start_time Call start time filter
  # @option opts [String] :filters_created_at Call log creation time filter
  # @option opts [String] :filters_direction Call direction filter: in or out
  # @option opts [String] :filters_called_number Called number
  # @option opts [String] :filters_type Call type, such as &#39;call&#39;, &#39;fax&#39;, &#39;audiogram&#39;
  # @option opts [Array<String>] :filters_extension Extension filter
  # @option opts [String] :sort_id ID sorting
  # @option opts [String] :sort_start_time Sorting by call start time: asc or desc
  # @option opts [String] :sort_created_at Sorting by call log creation time: asc or desc
  # @option opts [Integer] :limit Max results
  # @option opts [Integer] :offset Results to skip
  # @option opts [String] :fields Field set
  # @return [ListCallLogsFull]
  describe 'list_account_call_logs test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end