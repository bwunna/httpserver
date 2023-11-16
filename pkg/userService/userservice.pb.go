// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: userservice.proto

package userService

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Basic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Basic) Reset() {
	*x = Basic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Basic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Basic) ProtoMessage() {}

func (x *Basic) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Basic.ProtoReflect.Descriptor instead.
func (*Basic) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{0}
}

func (x *Basic) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Basic) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type EmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *EmailRequest) Reset() {
	*x = EmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequest) ProtoMessage() {}

func (x *EmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequest.ProtoReflect.Descriptor instead.
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{1}
}

func (x *EmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email          string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Status         string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Salary         int32  `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	CompanyName    string `protobuf:"bytes,5,opt,name=companyName,proto3" json:"companyName,omitempty"`
	DepartmentName string `protobuf:"bytes,6,opt,name=departmentName,proto3" json:"departmentName,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{2}
}

func (x *Employee) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Employee) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Employee) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Employee) GetSalary() int32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *Employee) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *Employee) GetDepartmentName() string {
	if x != nil {
		return x.DepartmentName
	}
	return ""
}

var File_userservice_proto protoreflect.FileDescriptor

var file_userservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x35, 0x0a, 0x05, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xae,
	0x01, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0x9c, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x42, 0x79,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0d, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22,
	0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2d, 0x62, 0x79, 0x2d, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x48, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12,
	0x09, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x1a, 0x06, 0x2e, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1e, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64,
	0x64, 0x2d, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x0d, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x06, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2c, 0x2a, 0x2a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2d, 0x62, 0x79, 0x2d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x13,
	0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userservice_proto_rawDescOnce sync.Once
	file_userservice_proto_rawDescData = file_userservice_proto_rawDesc
)

func file_userservice_proto_rawDescGZIP() []byte {
	file_userservice_proto_rawDescOnce.Do(func() {
		file_userservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_userservice_proto_rawDescData)
	})
	return file_userservice_proto_rawDescData
}

var file_userservice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_userservice_proto_goTypes = []interface{}{
	(*Basic)(nil),        // 0: Basic
	(*EmailRequest)(nil), // 1: EmailRequest
	(*Employee)(nil),     // 2: Employee
}
var file_userservice_proto_depIdxs = []int32{
	1, // 0: UserService.GetEmployeeByEmail:input_type -> EmailRequest
	2, // 1: UserService.AddEmployee:input_type -> Employee
	1, // 2: UserService.DeleteEmployeeByEmail:input_type -> EmailRequest
	2, // 3: UserService.GetEmployeeByEmail:output_type -> Employee
	0, // 4: UserService.AddEmployee:output_type -> Basic
	0, // 5: UserService.DeleteEmployeeByEmail:output_type -> Basic
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userservice_proto_init() }
func file_userservice_proto_init() {
	if File_userservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Basic); i {
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
		file_userservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailRequest); i {
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
		file_userservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
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
			RawDescriptor: file_userservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userservice_proto_goTypes,
		DependencyIndexes: file_userservice_proto_depIdxs,
		MessageInfos:      file_userservice_proto_msgTypes,
	}.Build()
	File_userservice_proto = out.File
	file_userservice_proto_rawDesc = nil
	file_userservice_proto_goTypes = nil
	file_userservice_proto_depIdxs = nil
}
