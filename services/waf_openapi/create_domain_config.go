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

// CreateDomainConfig invokes the waf_openapi.CreateDomainConfig API synchronously
// api document: https://help.aliyun.com/api/waf-openapi/createdomainconfig.html
func (client *Client) CreateDomainConfig(request *CreateDomainConfigRequest) (response *CreateDomainConfigResponse, err error) {
	response = CreateCreateDomainConfigResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDomainConfigWithChan invokes the waf_openapi.CreateDomainConfig API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/createdomainconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDomainConfigWithChan(request *CreateDomainConfigRequest) (<-chan *CreateDomainConfigResponse, <-chan error) {
	responseChan := make(chan *CreateDomainConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDomainConfig(request)
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

// CreateDomainConfigWithCallback invokes the waf_openapi.CreateDomainConfig API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/createdomainconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDomainConfigWithCallback(request *CreateDomainConfigRequest, callback func(response *CreateDomainConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDomainConfigResponse
		var err error
		defer close(result)
		response, err = client.CreateDomainConfig(request)
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

// CreateDomainConfigRequest is the request struct for api CreateDomainConfig
type CreateDomainConfigRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	HttpPort        string           `position:"Query" name:"HttpPort"`
	Lang            string           `position:"Query" name:"Lang"`
	Protocols       string           `position:"Query" name:"Protocols"`
	RsType          requests.Integer `position:"Query" name:"RsType"`
	HttpsRedirect   requests.Integer `position:"Query" name:"HttpsRedirect"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	SourceIps       string           `position:"Query" name:"SourceIps"`
	Domain          string           `position:"Query" name:"Domain"`
	IsAccessProduct requests.Integer `position:"Query" name:"IsAccessProduct"`
	HttpsPort       string           `position:"Query" name:"HttpsPort"`
	Region          string           `position:"Query" name:"Region"`
	LoadBalancing   requests.Integer `position:"Query" name:"LoadBalancing"`
	HttpToUserIp    requests.Integer `position:"Query" name:"HttpToUserIp"`
}

// CreateDomainConfigResponse is the response struct for api CreateDomainConfig
type CreateDomainConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateCreateDomainConfigRequest creates a request to invoke CreateDomainConfig API
func CreateCreateDomainConfigRequest() (request *CreateDomainConfigRequest) {
	request = &CreateDomainConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2018-01-17", "CreateDomainConfig", "waf", "openAPI")
	return
}

// CreateCreateDomainConfigResponse creates a response to parse from CreateDomainConfig response
func CreateCreateDomainConfigResponse() (response *CreateDomainConfigResponse) {
	response = &CreateDomainConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
