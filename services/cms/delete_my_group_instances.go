package cms

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

// DeleteMyGroupInstances invokes the cms.DeleteMyGroupInstances API synchronously
// api document: https://help.aliyun.com/api/cms/deletemygroupinstances.html
func (client *Client) DeleteMyGroupInstances(request *DeleteMyGroupInstancesRequest) (response *DeleteMyGroupInstancesResponse, err error) {
	response = CreateDeleteMyGroupInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMyGroupInstancesWithChan invokes the cms.DeleteMyGroupInstances API asynchronously
// api document: https://help.aliyun.com/api/cms/deletemygroupinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMyGroupInstancesWithChan(request *DeleteMyGroupInstancesRequest) (<-chan *DeleteMyGroupInstancesResponse, <-chan error) {
	responseChan := make(chan *DeleteMyGroupInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMyGroupInstances(request)
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

// DeleteMyGroupInstancesWithCallback invokes the cms.DeleteMyGroupInstances API asynchronously
// api document: https://help.aliyun.com/api/cms/deletemygroupinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMyGroupInstancesWithCallback(request *DeleteMyGroupInstancesRequest, callback func(response *DeleteMyGroupInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMyGroupInstancesResponse
		var err error
		defer close(result)
		response, err = client.DeleteMyGroupInstances(request)
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

// DeleteMyGroupInstancesRequest is the request struct for api DeleteMyGroupInstances
type DeleteMyGroupInstancesRequest struct {
	*requests.RpcRequest
	InstanceIds string           `position:"Query" name:"InstanceIds"`
	GroupId     requests.Integer `position:"Query" name:"GroupId"`
}

// DeleteMyGroupInstancesResponse is the response struct for api DeleteMyGroupInstances
type DeleteMyGroupInstancesResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteMyGroupInstancesRequest creates a request to invoke DeleteMyGroupInstances API
func CreateDeleteMyGroupInstancesRequest() (request *DeleteMyGroupInstancesRequest) {
	request = &DeleteMyGroupInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "DeleteMyGroupInstances", "cms", "openAPI")
	return
}

// CreateDeleteMyGroupInstancesResponse creates a response to parse from DeleteMyGroupInstances response
func CreateDeleteMyGroupInstancesResponse() (response *DeleteMyGroupInstancesResponse) {
	response = &DeleteMyGroupInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
