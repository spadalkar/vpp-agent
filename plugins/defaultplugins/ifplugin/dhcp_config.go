// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate binapi-generator --input-file=/usr/share/vpp/api/dhcp.api.json --output-dir=../common/bin_api

package ifplugin

// DhcpConfigurator handles DHCP-related configuration for VPP interfaces
type DhcpConfigurator struct {
	// to be defined
}

func (plugin *DhcpConfigurator) Init() error {
	return nil
}

func (plugin *DhcpConfigurator) Close() error {
	return nil
}

