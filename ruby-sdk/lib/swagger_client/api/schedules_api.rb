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

require "uri"

module SwaggerClient
  class SchedulesApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end

    # Show details of an individual schedule
    # This service shows the details of an individual schedule.
    # @param account_id Account ID
    # @param schedule_id Schedule ID
    # @param [Hash] opts the optional parameters
    # @return [ScheduleFull]
    def get_account_schedule(account_id, schedule_id, opts = {})
      data, _status_code, _headers = get_account_schedule_with_http_info(account_id, schedule_id, opts)
      return data
    end

    # Show details of an individual schedule
    # This service shows the details of an individual schedule.
    # @param account_id Account ID
    # @param schedule_id Schedule ID
    # @param [Hash] opts the optional parameters
    # @return [Array<(ScheduleFull, Fixnum, Hash)>] ScheduleFull data, response status code and response headers
    def get_account_schedule_with_http_info(account_id, schedule_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug "Calling API: SchedulesApi.get_account_schedule ..."
      end
      # verify the required parameter 'account_id' is set
      fail ArgumentError, "Missing the required parameter 'account_id' when calling SchedulesApi.get_account_schedule" if account_id.nil?
      # verify the required parameter 'schedule_id' is set
      fail ArgumentError, "Missing the required parameter 'schedule_id' when calling SchedulesApi.get_account_schedule" if schedule_id.nil?
      # resource path
      local_var_path = "/accounts/{account_id}/schedules/{schedule_id}".sub('{format}','json').sub('{' + 'account_id' + '}', account_id.to_s).sub('{' + 'schedule_id' + '}', schedule_id.to_s)

      # query parameters
      query_params = {}

      # header parameters
      header_params = {}

      # HTTP header 'Accept' (if needed)
      local_header_accept = ['application/json']
      local_header_accept_result = @api_client.select_header_accept(local_header_accept) and header_params['Accept'] = local_header_accept_result

      # HTTP header 'Content-Type'
      local_header_content_type = ['application/json']
      header_params['Content-Type'] = @api_client.select_header_content_type(local_header_content_type)

      # form parameters
      form_params = {}

      # http body (model)
      post_body = nil
      auth_names = ['apiKey']
      data, status_code, headers = @api_client.call_api(:GET, local_var_path,
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => 'ScheduleFull')
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: SchedulesApi#get_account_schedule\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get a list of schedules for an account
    # See Intro to Schedules for more info on the properties.
    # @param account_id Account ID
    # @param [Hash] opts the optional parameters
    # @option opts [Array<String>] :filters_id ID filter
    # @option opts [Array<String>] :filters_name Name filter
    # @option opts [String] :sort_id ID sorting
    # @option opts [String] :sort_name Name sorting
    # @option opts [Integer] :limit Max results
    # @option opts [Integer] :offset Results to skip
    # @option opts [String] :fields Field set
    # @return [ListSchedulesFull]
    def list_account_schedules(account_id, opts = {})
      data, _status_code, _headers = list_account_schedules_with_http_info(account_id, opts)
      return data
    end

    # Get a list of schedules for an account
    # See Intro to Schedules for more info on the properties.
    # @param account_id Account ID
    # @param [Hash] opts the optional parameters
    # @option opts [Array<String>] :filters_id ID filter
    # @option opts [Array<String>] :filters_name Name filter
    # @option opts [String] :sort_id ID sorting
    # @option opts [String] :sort_name Name sorting
    # @option opts [Integer] :limit Max results
    # @option opts [Integer] :offset Results to skip
    # @option opts [String] :fields Field set
    # @return [Array<(ListSchedulesFull, Fixnum, Hash)>] ListSchedulesFull data, response status code and response headers
    def list_account_schedules_with_http_info(account_id, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug "Calling API: SchedulesApi.list_account_schedules ..."
      end
      # verify the required parameter 'account_id' is set
      fail ArgumentError, "Missing the required parameter 'account_id' when calling SchedulesApi.list_account_schedules" if account_id.nil?
      if !opts[:'sort_id'].nil? && opts[:'sort_id'] !~ Regexp.new(asc|desc)
        fail ArgumentError, 'invalid value for "opts[:"sort_id"]" when calling SchedulesApi.list_account_schedules, must conform to the pattern asc|desc.'
      end

      if !opts[:'sort_name'].nil? && opts[:'sort_name'] !~ Regexp.new(asc|desc)
        fail ArgumentError, 'invalid value for "opts[:"sort_name"]" when calling SchedulesApi.list_account_schedules, must conform to the pattern asc|desc.'
      end

      # resource path
      local_var_path = "/accounts/{account_id}/schedules".sub('{format}','json').sub('{' + 'account_id' + '}', account_id.to_s)

      # query parameters
      query_params = {}
      query_params[:'filters[id]'] = @api_client.build_collection_param(opts[:'filters_id'], :multi) if !opts[:'filters_id'].nil?
      query_params[:'filters[name]'] = @api_client.build_collection_param(opts[:'filters_name'], :multi) if !opts[:'filters_name'].nil?
      query_params[:'sort[id]'] = opts[:'sort_id'] if !opts[:'sort_id'].nil?
      query_params[:'sort[name]'] = opts[:'sort_name'] if !opts[:'sort_name'].nil?
      query_params[:'limit'] = opts[:'limit'] if !opts[:'limit'].nil?
      query_params[:'offset'] = opts[:'offset'] if !opts[:'offset'].nil?
      query_params[:'fields'] = opts[:'fields'] if !opts[:'fields'].nil?

      # header parameters
      header_params = {}

      # HTTP header 'Accept' (if needed)
      local_header_accept = ['application/json']
      local_header_accept_result = @api_client.select_header_accept(local_header_accept) and header_params['Accept'] = local_header_accept_result

      # HTTP header 'Content-Type'
      local_header_content_type = ['application/json']
      header_params['Content-Type'] = @api_client.select_header_content_type(local_header_content_type)

      # form parameters
      form_params = {}

      # http body (model)
      post_body = nil
      auth_names = ['apiKey']
      data, status_code, headers = @api_client.call_api(:GET, local_var_path,
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => 'ListSchedulesFull')
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: SchedulesApi#list_account_schedules\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
