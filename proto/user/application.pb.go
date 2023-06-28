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
// 	protoc        v3.21.12
// source: proto/user/application.proto

package user

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

type GetUserContextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserContextRequest) Reset() {
	*x = GetUserContextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserContextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserContextRequest) ProtoMessage() {}

func (x *GetUserContextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserContextRequest.ProtoReflect.Descriptor instead.
func (*GetUserContextRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_application_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserContextRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// The response message containing the greetings
type GetUserContextReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceGroups []*ResourceGroup `protobuf:"bytes,1,rep,name=resourceGroups,proto3" json:"resourceGroups,omitempty"`
}

func (x *GetUserContextReply) Reset() {
	*x = GetUserContextReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserContextReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserContextReply) ProtoMessage() {}

func (x *GetUserContextReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserContextReply.ProtoReflect.Descriptor instead.
func (*GetUserContextReply) Descriptor() ([]byte, []int) {
	return file_proto_user_application_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserContextReply) GetResourceGroups() []*ResourceGroup {
	if x != nil {
		return x.ResourceGroups
	}
	return nil
}

type ResourceGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrgId     int64       `protobuf:"varint,2,opt,name=orgId,proto3" json:"orgId,omitempty"`
	Name      string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Resources []*Resource `protobuf:"bytes,4,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *ResourceGroup) Reset() {
	*x = ResourceGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceGroup) ProtoMessage() {}

func (x *ResourceGroup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceGroup.ProtoReflect.Descriptor instead.
func (*ResourceGroup) Descriptor() ([]byte, []int) {
	return file_proto_user_application_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceGroup) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResourceGroup) GetOrgId() int64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *ResourceGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceGroup) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt int64  `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_proto_user_application_proto_rawDescGZIP(), []int{3}
}

func (x *Resource) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Resource) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Resource) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_proto_user_application_proto protoreflect.FileDescriptor

var file_proto_user_application_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x59, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x42, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x7e, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x32, 0x67, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x56, 0x0a, 0x12, 0x69, 0x6f, 0x2e, 0x6e, 0x75, 0x6c, 0x6c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75, 0x6c, 0x6c, 0x2d, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x2f, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_application_proto_rawDescOnce sync.Once
	file_proto_user_application_proto_rawDescData = file_proto_user_application_proto_rawDesc
)

func file_proto_user_application_proto_rawDescGZIP() []byte {
	file_proto_user_application_proto_rawDescOnce.Do(func() {
		file_proto_user_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_application_proto_rawDescData)
	})
	return file_proto_user_application_proto_rawDescData
}

var file_proto_user_application_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_user_application_proto_goTypes = []interface{}{
	(*GetUserContextRequest)(nil), // 0: application.GetUserContextRequest
	(*GetUserContextReply)(nil),   // 1: application.GetUserContextReply
	(*ResourceGroup)(nil),         // 2: application.ResourceGroup
	(*Resource)(nil),              // 3: application.Resource
}
var file_proto_user_application_proto_depIdxs = []int32{
	2, // 0: application.GetUserContextReply.resourceGroups:type_name -> application.ResourceGroup
	3, // 1: application.ResourceGroup.resources:type_name -> application.Resource
	0, // 2: application.UserService.GetUserContext:input_type -> application.GetUserContextRequest
	1, // 3: application.UserService.GetUserContext:output_type -> application.GetUserContextReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_user_application_proto_init() }
func file_proto_user_application_proto_init() {
	if File_proto_user_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserContextRequest); i {
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
		file_proto_user_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserContextReply); i {
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
		file_proto_user_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceGroup); i {
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
		file_proto_user_application_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
			RawDescriptor: file_proto_user_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_application_proto_goTypes,
		DependencyIndexes: file_proto_user_application_proto_depIdxs,
		MessageInfos:      file_proto_user_application_proto_msgTypes,
	}.Build()
	File_proto_user_application_proto = out.File
	file_proto_user_application_proto_rawDesc = nil
	file_proto_user_application_proto_goTypes = nil
	file_proto_user_application_proto_depIdxs = nil
}
