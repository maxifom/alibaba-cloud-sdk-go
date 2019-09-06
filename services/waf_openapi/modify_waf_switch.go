package waf_openapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyWafSwitch invokes the waf_openapi.ModifyWafSwitch API synchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifywafswitch.html
func (client *Client) ModifyWafSwitch(request *ModifyWafSwitchRequest) (response *ModifyWafSwitchResponse, err error) {
	response = CreateModifyWafSwitchResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyWafSwitchWithChan invokes the waf_openapi.ModifyWafSwitch API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifywafswitch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWafSwitchWithChan(request *ModifyWafSwitchRequest) (<-chan *ModifyWafSwitchResponse, <-chan error) {
	responseChan := make(chan *ModifyWafSwitchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyWafSwitch(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyWafSwitchWithCallback invokes the waf_openapi.ModifyWafSwitch API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifywafswitch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyWafSwitchWithCallback(request *ModifyWafSwitchRequest, callback func(response *ModifyWafSwitchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyWafSwitchResponse
		var err error
		defer close(result)
		response, err = client.ModifyWafSwitch(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyWafSwitchRequest is the request struct for api ModifyWafSwitch
type ModifyWafSwitchRequest struct {
	*requests.RpcRequest
	InstanceId string           `position:"Query" name:"InstanceId"`
	SourceIp   string           `position:"Query" name:"SourceIp"`
	Domain     string           `position:"Query" name:"Domain"`
	ServiceOn  requests.Integer `position:"Query" name:"ServiceOn"`
	Lang       string           `position:"Query" name:"Lang"`
	Region     string           `position:"Query" name:"Region"`
}

// ModifyWafSwitchResponse is the response struct for api ModifyWafSwitch
type ModifyWafSwitchResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateModifyWafSwitchRequest creates a request to invoke ModifyWafSwitch API
func CreateModifyWafSwitchRequest() (request *ModifyWafSwitchRequest) {
	request = &ModifyWafSwitchRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2018-01-17", "ModifyWafSwitch", "waf", "openAPI")
	return
}

// CreateModifyWafSwitchResponse creates a response to parse from ModifyWafSwitch response
func CreateModifyWafSwitchResponse() (response *ModifyWafSwitchResponse) {
	response = &ModifyWafSwitchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
