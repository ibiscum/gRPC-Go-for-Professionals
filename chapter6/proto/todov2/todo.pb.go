// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: proto/todov2/todo.proto

package todov2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool                   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	DueDate     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todov2_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todov2_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_todov2_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *Task) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

type AddTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string                 `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	DueDate     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
}

func (x *AddTaskRequest) Reset() {
	*x = AddTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todov2_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskRequest) ProtoMessage() {}

func (x *AddTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todov2_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskRequest.ProtoReflect.Descriptor instead.
func (*AddTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_todov2_todo_proto_rawDescGZIP(), []int{1}
}

func (x *AddTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddTaskRequest) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

type AddTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddTaskResponse) Reset() {
	*x = AddTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todov2_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskResponse) ProtoMessage() {}

func (x *AddTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todov2_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskResponse.ProtoReflect.Descriptor instead.
func (*AddTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_todov2_todo_proto_rawDescGZIP(), []int{2}
}

func (x *AddTaskResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mask *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *ListTasksRequest) Reset() {
	*x = ListTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todov2_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksRequest) ProtoMessage() {}

func (x *ListTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todov2_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksRequest.ProtoReflect.Descriptor instead.
func (*ListTasksRequest) Descriptor() ([]byte, []int) {
	return file_proto_todov2_todo_proto_rawDescGZIP(), []int{3}
}

func (x *ListTasksRequest) GetMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Mask
	}
	return nil
}

type ListTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task    *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Overdue bool  `protobuf:"varint,2,opt,name=overdue,proto3" json:"overdue,omitempty"`
}

func (x *ListTasksResponse) Reset() {
	*x = ListTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todov2_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksResponse) ProtoMessage() {}

func (x *ListTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todov2_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksResponse.ProtoReflect.Descriptor instead.
func (*ListTasksResponse) Descriptor() ([]byte, []int) {
	return file_proto_todov2_todo_proto_rawDescGZIP(), []int{4}
}

func (x *ListTasksResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *ListTasksResponse) GetOverdue() bool {
	if x != nil {
		return x.Overdue
	}
	return false
}

type UpdateTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool                   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	DueDate     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
}

func (x *UpdateTasksRequest) Reset() {
	*x = UpdateTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todov2_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTasksRequest) ProtoMessage() {}

func (x *UpdateTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todov2_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTasksRequest.ProtoReflect.Descriptor instead.
func (*UpdateTasksRequest) Descriptor() ([]byte, []int) {
	return file_proto_todov2_todo_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTasksRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTasksRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateTasksRequest) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *UpdateTasksRequest) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

type UpdateTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTasksResponse) Reset() {
	*x = UpdateTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todov2_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTasksResponse) ProtoMessage() {}

func (x *UpdateTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todov2_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTasksResponse.ProtoReflect.Descriptor instead.
func (*UpdateTasksResponse) Descriptor() ([]byte, []int) {
	return file_proto_todov2_todo_proto_rawDescGZIP(), []int{6}
}

type DeleteTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTasksRequest) Reset() {
	*x = DeleteTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todov2_todo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTasksRequest) ProtoMessage() {}

func (x *DeleteTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todov2_todo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTasksRequest.ProtoReflect.Descriptor instead.
func (*DeleteTasksRequest) Descriptor() ([]byte, []int) {
	return file_proto_todov2_todo_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTasksRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTasksResponse) Reset() {
	*x = DeleteTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_todov2_todo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTasksResponse) ProtoMessage() {}

func (x *DeleteTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_todov2_todo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTasksResponse.ProtoReflect.Descriptor instead.
func (*DeleteTasksResponse) Descriptor() ([]byte, []int) {
	return file_proto_todov2_todo_proto_rawDescGZIP(), []int{8}
}

var File_proto_todov2_todo_proto protoreflect.FileDescriptor

var file_proto_todov2_todo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x76, 0x32, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x76,
	0x32, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64,
	0x6f, 0x6e, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x69, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35,
	0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x75,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0x4f, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x64, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6f, 0x76, 0x65, 0x72, 0x64, 0x75, 0x65, 0x22, 0x91, 0x01,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa3, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x16, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x76, 0x32, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x76,
	0x32, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x18,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x76,
	0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x48, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x12, 0x1a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x76, 0x32, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12,
	0x4a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1a,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x44, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x62, 0x69, 0x73, 0x63, 0x75,
	0x6d, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x47, 0x6f, 0x2d, 0x66, 0x6f, 0x72, 0x2d, 0x50, 0x72,
	0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x36, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x76,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_todov2_todo_proto_rawDescOnce sync.Once
	file_proto_todov2_todo_proto_rawDescData = file_proto_todov2_todo_proto_rawDesc
)

func file_proto_todov2_todo_proto_rawDescGZIP() []byte {
	file_proto_todov2_todo_proto_rawDescOnce.Do(func() {
		file_proto_todov2_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_todov2_todo_proto_rawDescData)
	})
	return file_proto_todov2_todo_proto_rawDescData
}

var file_proto_todov2_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_todov2_todo_proto_goTypes = []interface{}{
	(*Task)(nil),                  // 0: todov2.Task
	(*AddTaskRequest)(nil),        // 1: todov2.AddTaskRequest
	(*AddTaskResponse)(nil),       // 2: todov2.AddTaskResponse
	(*ListTasksRequest)(nil),      // 3: todov2.ListTasksRequest
	(*ListTasksResponse)(nil),     // 4: todov2.ListTasksResponse
	(*UpdateTasksRequest)(nil),    // 5: todov2.UpdateTasksRequest
	(*UpdateTasksResponse)(nil),   // 6: todov2.UpdateTasksResponse
	(*DeleteTasksRequest)(nil),    // 7: todov2.DeleteTasksRequest
	(*DeleteTasksResponse)(nil),   // 8: todov2.DeleteTasksResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 10: google.protobuf.FieldMask
}
var file_proto_todov2_todo_proto_depIdxs = []int32{
	9,  // 0: todov2.Task.due_date:type_name -> google.protobuf.Timestamp
	9,  // 1: todov2.AddTaskRequest.due_date:type_name -> google.protobuf.Timestamp
	10, // 2: todov2.ListTasksRequest.mask:type_name -> google.protobuf.FieldMask
	0,  // 3: todov2.ListTasksResponse.task:type_name -> todov2.Task
	9,  // 4: todov2.UpdateTasksRequest.due_date:type_name -> google.protobuf.Timestamp
	1,  // 5: todov2.TodoService.AddTask:input_type -> todov2.AddTaskRequest
	3,  // 6: todov2.TodoService.ListTasks:input_type -> todov2.ListTasksRequest
	5,  // 7: todov2.TodoService.UpdateTasks:input_type -> todov2.UpdateTasksRequest
	7,  // 8: todov2.TodoService.DeleteTasks:input_type -> todov2.DeleteTasksRequest
	2,  // 9: todov2.TodoService.AddTask:output_type -> todov2.AddTaskResponse
	4,  // 10: todov2.TodoService.ListTasks:output_type -> todov2.ListTasksResponse
	6,  // 11: todov2.TodoService.UpdateTasks:output_type -> todov2.UpdateTasksResponse
	8,  // 12: todov2.TodoService.DeleteTasks:output_type -> todov2.DeleteTasksResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_todov2_todo_proto_init() }
func file_proto_todov2_todo_proto_init() {
	if File_proto_todov2_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_todov2_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_proto_todov2_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskRequest); i {
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
		file_proto_todov2_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskResponse); i {
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
		file_proto_todov2_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTasksRequest); i {
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
		file_proto_todov2_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTasksResponse); i {
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
		file_proto_todov2_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTasksRequest); i {
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
		file_proto_todov2_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTasksResponse); i {
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
		file_proto_todov2_todo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTasksRequest); i {
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
		file_proto_todov2_todo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTasksResponse); i {
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
			RawDescriptor: file_proto_todov2_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_todov2_todo_proto_goTypes,
		DependencyIndexes: file_proto_todov2_todo_proto_depIdxs,
		MessageInfos:      file_proto_todov2_todo_proto_msgTypes,
	}.Build()
	File_proto_todov2_todo_proto = out.File
	file_proto_todov2_todo_proto_rawDesc = nil
	file_proto_todov2_todo_proto_goTypes = nil
	file_proto_todov2_todo_proto_depIdxs = nil
}
