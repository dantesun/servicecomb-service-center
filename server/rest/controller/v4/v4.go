//Copyright 2017 Huawei Technologies Co., Ltd
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
package v4

import (
	roa "github.com/ServiceComb/service-center/pkg/rest"
)

func init() {
	initRouter()
}

func initRouter() {
	roa.RegisterServent(&MainService{})
	roa.RegisterServent(&MicroServiceService{})
	roa.RegisterServent(&SchemaService{})
	roa.RegisterServent(&DependencyService{})
	roa.RegisterServent(&TagService{})
	roa.RegisterServent(&RuleService{})
	roa.RegisterServent(&MicroServiceInstanceService{})
	roa.RegisterServent(&WatchService{})
}