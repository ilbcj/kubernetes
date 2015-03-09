/*
Copyright 2015 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cadvisor

import (
	cadvisorApi "github.com/google/cadvisor/info/v1"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

// ContainerInfo is a mock implementation of CadvisorInterface.ContainerInfo.
func (c *Mock) ContainerInfo(name string, req *cadvisorApi.ContainerInfoRequest) (*cadvisorApi.ContainerInfo, error) {
	args := c.Called(name, req)
	return args.Get(0).(*cadvisorApi.ContainerInfo), args.Error(1)
}

// DockerContainer is a mock implementation of CadvisorInterface.DockerContainer.
func (c *Mock) DockerContainer(name string, req *cadvisorApi.ContainerInfoRequest) (cadvisorApi.ContainerInfo, error) {
	args := c.Called(name, req)
	return args.Get(0).(cadvisorApi.ContainerInfo), args.Error(1)
}

// MachineInfo is a mock implementation of CadvisorInterface.MachineInfo.
func (c *Mock) MachineInfo() (*cadvisorApi.MachineInfo, error) {
	args := c.Called()
	return args.Get(0).(*cadvisorApi.MachineInfo), args.Error(1)
}
