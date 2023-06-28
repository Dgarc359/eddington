// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/github-watcher/watchRepo.proto

package github_watcher

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

type WatchRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Branch     string `protobuf:"bytes,2,opt,name=branch,proto3" json:"branch,omitempty"`
	SHA        string `protobuf:"bytes,3,opt,name=SHA,proto3" json:"SHA,omitempty"`
	Owner      string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *WatchRepoRequest) Reset() {
	*x = WatchRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_github_watcher_watchRepo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRepoRequest) ProtoMessage() {}

func (x *WatchRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_github_watcher_watchRepo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRepoRequest.ProtoReflect.Descriptor instead.
func (*WatchRepoRequest) Descriptor() ([]byte, []int) {
	return file_proto_github_watcher_watchRepo_proto_rawDescGZIP(), []int{0}
}

func (x *WatchRepoRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *WatchRepoRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *WatchRepoRequest) GetSHA() string {
	if x != nil {
		return x.SHA
	}
	return ""
}

func (x *WatchRepoRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

type WatchRepoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsUpdated bool `protobuf:"varint,1,opt,name=isUpdated,proto3" json:"isUpdated,omitempty"`
}

func (x *WatchRepoResponse) Reset() {
	*x = WatchRepoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_github_watcher_watchRepo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRepoResponse) ProtoMessage() {}

func (x *WatchRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_github_watcher_watchRepo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRepoResponse.ProtoReflect.Descriptor instead.
func (*WatchRepoResponse) Descriptor() ([]byte, []int) {
	return file_proto_github_watcher_watchRepo_proto_rawDescGZIP(), []int{1}
}

func (x *WatchRepoResponse) GetIsUpdated() bool {
	if x != nil {
		return x.IsUpdated
	}
	return false
}

var File_proto_github_watcher_watchRepo_proto protoreflect.FileDescriptor

var file_proto_github_watcher_watchRepo_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2d, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0x72, 0x0a, 0x10, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x48, 0x41, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x53, 0x48, 0x41, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x11, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x32, 0x64, 0x0a, 0x10,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x50, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x1f, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x75, 0x6c, 0x6c, 0x2d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x65, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2d, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_github_watcher_watchRepo_proto_rawDescOnce sync.Once
	file_proto_github_watcher_watchRepo_proto_rawDescData = file_proto_github_watcher_watchRepo_proto_rawDesc
)

func file_proto_github_watcher_watchRepo_proto_rawDescGZIP() []byte {
	file_proto_github_watcher_watchRepo_proto_rawDescOnce.Do(func() {
		file_proto_github_watcher_watchRepo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_github_watcher_watchRepo_proto_rawDescData)
	})
	return file_proto_github_watcher_watchRepo_proto_rawDescData
}

var file_proto_github_watcher_watchRepo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_github_watcher_watchRepo_proto_goTypes = []interface{}{
	(*WatchRepoRequest)(nil),  // 0: githubWatcher.WatchRepoRequest
	(*WatchRepoResponse)(nil), // 1: githubWatcher.WatchRepoResponse
}
var file_proto_github_watcher_watchRepo_proto_depIdxs = []int32{
	0, // 0: githubWatcher.WatchRepoService.WatchRepo:input_type -> githubWatcher.WatchRepoRequest
	1, // 1: githubWatcher.WatchRepoService.WatchRepo:output_type -> githubWatcher.WatchRepoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_github_watcher_watchRepo_proto_init() }
func file_proto_github_watcher_watchRepo_proto_init() {
	if File_proto_github_watcher_watchRepo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_github_watcher_watchRepo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRepoRequest); i {
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
		file_proto_github_watcher_watchRepo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRepoResponse); i {
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
			RawDescriptor: file_proto_github_watcher_watchRepo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_github_watcher_watchRepo_proto_goTypes,
		DependencyIndexes: file_proto_github_watcher_watchRepo_proto_depIdxs,
		MessageInfos:      file_proto_github_watcher_watchRepo_proto_msgTypes,
	}.Build()
	File_proto_github_watcher_watchRepo_proto = out.File
	file_proto_github_watcher_watchRepo_proto_rawDesc = nil
	file_proto_github_watcher_watchRepo_proto_goTypes = nil
	file_proto_github_watcher_watchRepo_proto_depIdxs = nil
}
