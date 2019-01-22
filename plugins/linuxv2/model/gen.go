//  Copyright (c) 2018 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

//go:generate protoc --proto_path=namespace --gogo_out=namespace namespace.proto
//go:generate protoc --proto_path=interfaces --proto_path=${GOPATH}/src --gogo_out=interfaces interfaces.proto
//go:generate protoc --proto_path=l3 --proto_path=${GOPATH}/src --gogo_out=l3 l3.proto
//go:generate protoc --proto_path=punt --proto_path=${GOPATH}/src --gogo_out=punt punt/punt.proto

package model