// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: examples/busi.proto

package examples

import (
	dtmcli "github.com/yedf/dtm/dtmcli"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_examples_busi_proto protoreflect.FileDescriptor

var file_examples_busi_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x1a,
	0x13, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2f, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9a, 0x02, 0x0a, 0x04, 0x42, 0x75, 0x73, 0x69, 0x12, 0x30, 0x0a,
	0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2e, 0x42,
	0x75, 0x73, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x64, 0x74, 0x6d,
	0x63, 0x6c, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x12, 0x13, 0x2e, 0x64, 0x74, 0x6d,
	0x63, 0x6c, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x75, 0x74,
	0x12, 0x13, 0x2e, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2e, 0x42,
	0x75, 0x73, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x49, 0x6e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x64, 0x74,
	0x6d, 0x63, 0x6c, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4f, 0x75,
	0x74, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69,
	0x2e, 0x42, 0x75, 0x73, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x64,
	0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x65, 0x64, 0x66, 0x2f, 0x64, 0x74, 0x6d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_examples_busi_proto_goTypes = []interface{}{
	(*dtmcli.BusiRequest)(nil), // 0: dtmcli.BusiRequest
	(*dtmcli.BusiReply)(nil),   // 1: dtmcli.BusiReply
}
var file_examples_busi_proto_depIdxs = []int32{
	0, // 0: examples.Busi.Call:input_type -> dtmcli.BusiRequest
	0, // 1: examples.Busi.TransIn:input_type -> dtmcli.BusiRequest
	0, // 2: examples.Busi.TransOut:input_type -> dtmcli.BusiRequest
	0, // 3: examples.Busi.TransInRevert:input_type -> dtmcli.BusiRequest
	0, // 4: examples.Busi.TransOutRevert:input_type -> dtmcli.BusiRequest
	1, // 5: examples.Busi.Call:output_type -> dtmcli.BusiReply
	1, // 6: examples.Busi.TransIn:output_type -> dtmcli.BusiReply
	1, // 7: examples.Busi.TransOut:output_type -> dtmcli.BusiReply
	1, // 8: examples.Busi.TransInRevert:output_type -> dtmcli.BusiReply
	1, // 9: examples.Busi.TransOutRevert:output_type -> dtmcli.BusiReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_examples_busi_proto_init() }
func file_examples_busi_proto_init() {
	if File_examples_busi_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_examples_busi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_examples_busi_proto_goTypes,
		DependencyIndexes: file_examples_busi_proto_depIdxs,
	}.Build()
	File_examples_busi_proto = out.File
	file_examples_busi_proto_rawDesc = nil
	file_examples_busi_proto_goTypes = nil
	file_examples_busi_proto_depIdxs = nil
}
