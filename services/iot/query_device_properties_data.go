package iot

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

// QueryDevicePropertiesData invokes the iot.QueryDevicePropertiesData API synchronously
// api document: https://help.aliyun.com/api/iot/querydevicepropertiesdata.html
func (client *Client) QueryDevicePropertiesData(request *QueryDevicePropertiesDataRequest) (response *QueryDevicePropertiesDataResponse, err error) {
	response = CreateQueryDevicePropertiesDataResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDevicePropertiesDataWithChan invokes the iot.QueryDevicePropertiesData API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicepropertiesdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDevicePropertiesDataWithChan(request *QueryDevicePropertiesDataRequest) (<-chan *QueryDevicePropertiesDataResponse, <-chan error) {
	responseChan := make(chan *QueryDevicePropertiesDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDevicePropertiesData(request)
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

// QueryDevicePropertiesDataWithCallback invokes the iot.QueryDevicePropertiesData API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicepropertiesdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDevicePropertiesDataWithCallback(request *QueryDevicePropertiesDataRequest, callback func(response *QueryDevicePropertiesDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDevicePropertiesDataResponse
		var err error
		defer close(result)
		response, err = client.QueryDevicePropertiesData(request)
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

// QueryDevicePropertiesDataRequest is the request struct for api QueryDevicePropertiesData
type QueryDevicePropertiesDataRequest struct {
	*requests.RpcRequest
	Asc           requests.Integer `position:"Query" name:"Asc"`
	Identifier    *[]string        `position:"Query" name:"Identifier"  type:"Repeated"`
	IotId         string           `position:"Query" name:"IotId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
}

// QueryDevicePropertiesDataResponse is the response struct for api QueryDevicePropertiesData
type QueryDevicePropertiesDataResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Success           bool              `json:"Success" xml:"Success"`
	Code              string            `json:"Code" xml:"Code"`
	ErrorMessage      string            `json:"ErrorMessage" xml:"ErrorMessage"`
	NextValid         bool              `json:"NextValid" xml:"NextValid"`
	NextTime          int64             `json:"NextTime" xml:"NextTime"`
	PropertyDataInfos PropertyDataInfos `json:"PropertyDataInfos" xml:"PropertyDataInfos"`
}

// CreateQueryDevicePropertiesDataRequest creates a request to invoke QueryDevicePropertiesData API
func CreateQueryDevicePropertiesDataRequest() (request *QueryDevicePropertiesDataRequest) {
	request = &QueryDevicePropertiesDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDevicePropertiesData", "iot", "openAPI")
	return
}

// CreateQueryDevicePropertiesDataResponse creates a response to parse from QueryDevicePropertiesData response
func CreateQueryDevicePropertiesDataResponse() (response *QueryDevicePropertiesDataResponse) {
	response = &QueryDevicePropertiesDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
