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

// AclRule is a nested struct in waf_openapi response
type AclRule struct {
	ContinueWaf             int        `json:"ContinueWaf" xml:"ContinueWaf"`
	ContinueBlockGeo        int        `json:"ContinueBlockGeo" xml:"ContinueBlockGeo"`
	ContinueSA              int        `json:"ContinueSA" xml:"ContinueSA"`
	ContinueSdk             int        `json:"ContinueSdk" xml:"ContinueSdk"`
	ContinueDataRiskControl int        `json:"ContinueDataRiskControl" xml:"ContinueDataRiskControl"`
	ContinueCc              int        `json:"ContinueCc" xml:"ContinueCc"`
	IsDefault               int        `json:"IsDefault" xml:"IsDefault"`
	Name                    string     `json:"Name" xml:"Name"`
	Action                  int        `json:"Action" xml:"Action"`
	Id                      int64      `json:"Id" xml:"Id"`
	Order                   int        `json:"Order" xml:"Order"`
	Conditions              Conditions `json:"Conditions" xml:"Conditions"`
}
