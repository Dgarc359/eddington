// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/container/container.proto

package container

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum of ContainerStatus
type ContainerStatus int32

const (
	ContainerStatus_NOT_STARTED ContainerStatus = 0
	ContainerStatus_WAITING     ContainerStatus = 1
	ContainerStatus_BUILDING    ContainerStatus = 2
	ContainerStatus_BUILT       ContainerStatus = 3
	ContainerStatus_FAILED      ContainerStatus = 4
)

// Enum value maps for ContainerStatus.
var (
	ContainerStatus_name = map[int32]string{
		0: "NOT_STARTED",
		1: "WAITING",
		2: "BUILDING",
		3: "BUILT",
		4: "FAILED",
	}
	ContainerStatus_value = map[string]int32{
		"NOT_STARTED": 0,
		"WAITING":     1,
		"BUILDING":    2,
		"BUILT":       3,
		"FAILED":      4,
	}
)

func (x ContainerStatus) Enum() *ContainerStatus {
	p := new(ContainerStatus)
	*p = x
	return p
}

func (x ContainerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContainerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_container_container_proto_enumTypes[0].Descriptor()
}

func (ContainerStatus) Type() protoreflect.EnumType {
	return &file_proto_container_container_proto_enumTypes[0]
}

func (x ContainerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContainerStatus.Descriptor instead.
func (ContainerStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_container_container_proto_rawDescGZIP(), []int{0}
}

type CreateContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoURL    string `protobuf:"bytes,1,opt,name=repoURL,proto3" json:"repoURL,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	CustomerID string `protobuf:"bytes,3,opt,name=customerID,proto3" json:"customerID,omitempty"`
}

func (x *CreateContainerRequest) Reset() {
	*x = CreateContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_container_container_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContainerRequest) ProtoMessage() {}

func (x *CreateContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_container_container_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContainerRequest.ProtoReflect.Descriptor instead.
func (*CreateContainerRequest) Descriptor() ([]byte, []int) {
	return file_proto_container_container_proto_rawDescGZIP(), []int{0}
}

func (x *CreateContainerRequest) GetRepoURL() string {
	if x != nil {
		return x.RepoURL
	}
	return ""
}

func (x *CreateContainerRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateContainerRequest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

type CreateContainerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildID string `protobuf:"bytes,1,opt,name=buildID,proto3" json:"buildID,omitempty"`
}

func (x *CreateContainerReply) Reset() {
	*x = CreateContainerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_container_container_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContainerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContainerReply) ProtoMessage() {}

func (x *CreateContainerReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_container_container_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContainerReply.ProtoReflect.Descriptor instead.
func (*CreateContainerReply) Descriptor() ([]byte, []int) {
	return file_proto_container_container_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContainerReply) GetBuildID() string {
	if x != nil {
		return x.BuildID
	}
	return ""
}

type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_container_container_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_proto_container_container_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_proto_container_container_proto_rawDescGZIP(), []int{2}
}

func (x *Build) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ContainerImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ContainerStatus `protobuf:"varint,1,opt,name=status,proto3,enum=container.ContainerStatus" json:"status,omitempty"`
}

func (x *ContainerImage) Reset() {
	*x = ContainerImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_container_container_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerImage) ProtoMessage() {}

func (x *ContainerImage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_container_container_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerImage.ProtoReflect.Descriptor instead.
func (*ContainerImage) Descriptor() ([]byte, []int) {
	return file_proto_container_container_proto_rawDescGZIP(), []int{3}
}

func (x *ContainerImage) GetStatus() ContainerStatus {
	if x != nil {
		return x.Status
	}
	return ContainerStatus_NOT_STARTED
}

var File_proto_container_container_proto protoreflect.FileDescriptor

var file_proto_container_container_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x66, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x55, 0x52,
	0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x55, 0x52, 0x4c,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x30, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x22, 0x17, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x44, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x54, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x55, 0x49, 0x4c, 0x54, 0x10, 0x03, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x32, 0xa9, 0x01, 0x0a, 0x10,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x57, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x6b, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x17, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75,
	0x6c, 0x6c, 0x2d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x65, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_container_container_proto_rawDescOnce sync.Once
	file_proto_container_container_proto_rawDescData = file_proto_container_container_proto_rawDesc
)

func file_proto_container_container_proto_rawDescGZIP() []byte {
	file_proto_container_container_proto_rawDescOnce.Do(func() {
		file_proto_container_container_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_container_container_proto_rawDescData)
	})
	return file_proto_container_container_proto_rawDescData
}

var file_proto_container_container_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_container_container_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_container_container_proto_goTypes = []interface{}{
	(ContainerStatus)(0),           // 0: container.ContainerStatus
	(*CreateContainerRequest)(nil), // 1: container.CreateContainerRequest
	(*CreateContainerReply)(nil),   // 2: container.CreateContainerReply
	(*Build)(nil),                  // 3: container.Build
	(*ContainerImage)(nil),         // 4: container.ContainerImage
}
var file_proto_container_container_proto_depIdxs = []int32{
	0, // 0: container.ContainerImage.status:type_name -> container.ContainerStatus
	1, // 1: container.ContainerService.CreateContainer:input_type -> container.CreateContainerRequest
	3, // 2: container.ContainerService.ImageStatus:input_type -> container.Build
	2, // 3: container.ContainerService.CreateContainer:output_type -> container.CreateContainerReply
	4, // 4: container.ContainerService.ImageStatus:output_type -> container.ContainerImage
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_container_container_proto_init() }
func file_proto_container_container_proto_init() {
	if File_proto_container_container_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_container_container_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContainerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_container_container_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContainerReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_container_container_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_container_container_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerImage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_container_container_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_container_container_proto_goTypes,
		DependencyIndexes: file_proto_container_container_proto_depIdxs,
		EnumInfos:         file_proto_container_container_proto_enumTypes,
		MessageInfos:      file_proto_container_container_proto_msgTypes,
	}.Build()
	File_proto_container_container_proto = out.File
	file_proto_container_container_proto_rawDesc = nil
	file_proto_container_container_proto_goTypes = nil
	file_proto_container_container_proto_depIdxs = nil
}
