// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: car.proto

package carpb

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

type CarStatus int32

const (
	CarStatus_CS_NOT_SPECIFIED CarStatus = 0
	CarStatus_LOCKED           CarStatus = 1
	CarStatus_UNLOCKING        CarStatus = 2
	CarStatus_UNLOCKED         CarStatus = 3
	CarStatus_LOCKING          CarStatus = 4
)

// Enum value maps for CarStatus.
var (
	CarStatus_name = map[int32]string{
		0: "CS_NOT_SPECIFIED",
		1: "LOCKED",
		2: "UNLOCKING",
		3: "UNLOCKED",
		4: "LOCKING",
	}
	CarStatus_value = map[string]int32{
		"CS_NOT_SPECIFIED": 0,
		"LOCKED":           1,
		"UNLOCKING":        2,
		"UNLOCKED":         3,
		"LOCKING":          4,
	}
)

func (x CarStatus) Enum() *CarStatus {
	p := new(CarStatus)
	*p = x
	return p
}

func (x CarStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CarStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_car_proto_enumTypes[0].Descriptor()
}

func (CarStatus) Type() protoreflect.EnumType {
	return &file_car_proto_enumTypes[0]
}

func (x CarStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CarStatus.Descriptor instead.
func (CarStatus) EnumDescriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{0}
}

type CarEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Car *Car   `protobuf:"bytes,2,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *CarEntity) Reset() {
	*x = CarEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarEntity) ProtoMessage() {}

func (x *CarEntity) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarEntity.ProtoReflect.Descriptor instead.
func (*CarEntity) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{0}
}

func (x *CarEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CarEntity) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type Driver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AvatarUrl string `protobuf:"bytes,2,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
}

func (x *Driver) Reset() {
	*x = Driver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Driver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Driver) ProtoMessage() {}

func (x *Driver) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Driver.ProtoReflect.Descriptor instead.
func (*Driver) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{1}
}

func (x *Driver) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Driver) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{2}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   CarStatus `protobuf:"varint,1,opt,name=status,proto3,enum=car.v1.CarStatus" json:"status,omitempty"`
	Driver   *Driver   `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	Position *Location `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	TripId   string    `protobuf:"bytes,4,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{3}
}

func (x *Car) GetStatus() CarStatus {
	if x != nil {
		return x.Status
	}
	return CarStatus_CS_NOT_SPECIFIED
}

func (x *Car) GetDriver() *Driver {
	if x != nil {
		return x.Driver
	}
	return nil
}

func (x *Car) GetPosition() *Location {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Car) GetTripId() string {
	if x != nil {
		return x.TripId
	}
	return ""
}

type CreateCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateCarRequest) Reset() {
	*x = CreateCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarRequest) ProtoMessage() {}

func (x *CreateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarRequest.ProtoReflect.Descriptor instead.
func (*CreateCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{4}
}

type GetCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCarRequest) Reset() {
	*x = GetCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarRequest) ProtoMessage() {}

func (x *GetCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarRequest.ProtoReflect.Descriptor instead.
func (*GetCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{5}
}

func (x *GetCarRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCarsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCarsRequest) Reset() {
	*x = GetCarsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarsRequest) ProtoMessage() {}

func (x *GetCarsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarsRequest.ProtoReflect.Descriptor instead.
func (*GetCarsRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{6}
}

type GetCarsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cars []*CarEntity `protobuf:"bytes,1,rep,name=cars,proto3" json:"cars,omitempty"`
}

func (x *GetCarsResponse) Reset() {
	*x = GetCarsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarsResponse) ProtoMessage() {}

func (x *GetCarsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarsResponse.ProtoReflect.Descriptor instead.
func (*GetCarsResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{7}
}

func (x *GetCarsResponse) GetCars() []*CarEntity {
	if x != nil {
		return x.Cars
	}
	return nil
}

type LockCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LockCarRequest) Reset() {
	*x = LockCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockCarRequest) ProtoMessage() {}

func (x *LockCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockCarRequest.ProtoReflect.Descriptor instead.
func (*LockCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{8}
}

func (x *LockCarRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LockCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LockCarResponse) Reset() {
	*x = LockCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockCarResponse) ProtoMessage() {}

func (x *LockCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockCarResponse.ProtoReflect.Descriptor instead.
func (*LockCarResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{9}
}

type UnlockCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Driver *Driver `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	TripId string  `protobuf:"bytes,3,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
}

func (x *UnlockCarRequest) Reset() {
	*x = UnlockCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockCarRequest) ProtoMessage() {}

func (x *UnlockCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockCarRequest.ProtoReflect.Descriptor instead.
func (*UnlockCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{10}
}

func (x *UnlockCarRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UnlockCarRequest) GetDriver() *Driver {
	if x != nil {
		return x.Driver
	}
	return nil
}

func (x *UnlockCarRequest) GetTripId() string {
	if x != nil {
		return x.TripId
	}
	return ""
}

type UnlockCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnlockCarResponse) Reset() {
	*x = UnlockCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockCarResponse) ProtoMessage() {}

func (x *UnlockCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockCarResponse.ProtoReflect.Descriptor instead.
func (*UnlockCarResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{11}
}

type UpdateCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status   CarStatus `protobuf:"varint,2,opt,name=status,proto3,enum=car.v1.CarStatus" json:"status,omitempty"`
	Position *Location `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *UpdateCarRequest) Reset() {
	*x = UpdateCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarRequest) ProtoMessage() {}

func (x *UpdateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarRequest.ProtoReflect.Descriptor instead.
func (*UpdateCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateCarRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCarRequest) GetStatus() CarStatus {
	if x != nil {
		return x.Status
	}
	return CarStatus_CS_NOT_SPECIFIED
}

func (x *UpdateCarRequest) GetPosition() *Location {
	if x != nil {
		return x.Position
	}
	return nil
}

type UpdateCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCarResponse) Reset() {
	*x = UpdateCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarResponse) ProtoMessage() {}

func (x *UpdateCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarResponse.ProtoReflect.Descriptor instead.
func (*UpdateCarResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{13}
}

var File_car_proto protoreflect.FileDescriptor

var file_car_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x61, 0x72,
	0x2e, 0x76, 0x31, 0x22, 0x3a, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22,
	0x37, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x9f,
	0x01, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64,
	0x22, 0x12, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x61,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x63, 0x61, 0x72,
	0x73, 0x22, 0x20, 0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x0a, 0x10, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x7b, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2c, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x13, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2a, 0x57, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x10, 0x43, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b,
	0x0a, 0x07, 0x4c, 0x4f, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x32, 0xf0, 0x02, 0x0a, 0x0a,
	0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x12, 0x15,
	0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x72, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x12, 0x16, 0x2e,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x07, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x43,
	0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f,
	0x5a, 0x1d, 0x68, 0x61, 0x70, 0x70, 0x79, 0x63, 0x61, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x61, 0x72, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_car_proto_rawDescOnce sync.Once
	file_car_proto_rawDescData = file_car_proto_rawDesc
)

func file_car_proto_rawDescGZIP() []byte {
	file_car_proto_rawDescOnce.Do(func() {
		file_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_proto_rawDescData)
	})
	return file_car_proto_rawDescData
}

var file_car_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_car_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_car_proto_goTypes = []interface{}{
	(CarStatus)(0),            // 0: car.v1.CarStatus
	(*CarEntity)(nil),         // 1: car.v1.CarEntity
	(*Driver)(nil),            // 2: car.v1.Driver
	(*Location)(nil),          // 3: car.v1.Location
	(*Car)(nil),               // 4: car.v1.Car
	(*CreateCarRequest)(nil),  // 5: car.v1.CreateCarRequest
	(*GetCarRequest)(nil),     // 6: car.v1.GetCarRequest
	(*GetCarsRequest)(nil),    // 7: car.v1.GetCarsRequest
	(*GetCarsResponse)(nil),   // 8: car.v1.GetCarsResponse
	(*LockCarRequest)(nil),    // 9: car.v1.LockCarRequest
	(*LockCarResponse)(nil),   // 10: car.v1.LockCarResponse
	(*UnlockCarRequest)(nil),  // 11: car.v1.UnlockCarRequest
	(*UnlockCarResponse)(nil), // 12: car.v1.UnlockCarResponse
	(*UpdateCarRequest)(nil),  // 13: car.v1.UpdateCarRequest
	(*UpdateCarResponse)(nil), // 14: car.v1.UpdateCarResponse
}
var file_car_proto_depIdxs = []int32{
	4,  // 0: car.v1.CarEntity.car:type_name -> car.v1.Car
	0,  // 1: car.v1.Car.status:type_name -> car.v1.CarStatus
	2,  // 2: car.v1.Car.driver:type_name -> car.v1.Driver
	3,  // 3: car.v1.Car.position:type_name -> car.v1.Location
	1,  // 4: car.v1.GetCarsResponse.cars:type_name -> car.v1.CarEntity
	2,  // 5: car.v1.UnlockCarRequest.driver:type_name -> car.v1.Driver
	0,  // 6: car.v1.UpdateCarRequest.status:type_name -> car.v1.CarStatus
	3,  // 7: car.v1.UpdateCarRequest.position:type_name -> car.v1.Location
	5,  // 8: car.v1.CarService.CreateCar:input_type -> car.v1.CreateCarRequest
	6,  // 9: car.v1.CarService.GetCar:input_type -> car.v1.GetCarRequest
	7,  // 10: car.v1.CarService.GetCars:input_type -> car.v1.GetCarsRequest
	9,  // 11: car.v1.CarService.LockCar:input_type -> car.v1.LockCarRequest
	11, // 12: car.v1.CarService.UnlockCar:input_type -> car.v1.UnlockCarRequest
	13, // 13: car.v1.CarService.UpdateCar:input_type -> car.v1.UpdateCarRequest
	1,  // 14: car.v1.CarService.CreateCar:output_type -> car.v1.CarEntity
	4,  // 15: car.v1.CarService.GetCar:output_type -> car.v1.Car
	8,  // 16: car.v1.CarService.GetCars:output_type -> car.v1.GetCarsResponse
	10, // 17: car.v1.CarService.LockCar:output_type -> car.v1.LockCarResponse
	12, // 18: car.v1.CarService.UnlockCar:output_type -> car.v1.UnlockCarResponse
	14, // 19: car.v1.CarService.UpdateCar:output_type -> car.v1.UpdateCarResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_car_proto_init() }
func file_car_proto_init() {
	if File_car_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_car_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarEntity); i {
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
		file_car_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Driver); i {
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
		file_car_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_car_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_car_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarRequest); i {
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
		file_car_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarRequest); i {
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
		file_car_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarsRequest); i {
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
		file_car_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarsResponse); i {
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
		file_car_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockCarRequest); i {
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
		file_car_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockCarResponse); i {
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
		file_car_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockCarRequest); i {
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
		file_car_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockCarResponse); i {
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
		file_car_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarRequest); i {
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
		file_car_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarResponse); i {
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
			RawDescriptor: file_car_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_proto_goTypes,
		DependencyIndexes: file_car_proto_depIdxs,
		EnumInfos:         file_car_proto_enumTypes,
		MessageInfos:      file_car_proto_msgTypes,
	}.Build()
	File_car_proto = out.File
	file_car_proto_rawDesc = nil
	file_car_proto_goTypes = nil
	file_car_proto_depIdxs = nil
}
