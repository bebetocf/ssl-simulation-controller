// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: ssl_simulation_config.proto

package simctl

import (
	referee "github.com/RoboCup-SSL/ssl-simulation-controller/internal/referee"
	vision "github.com/RoboCup-SSL/ssl-simulation-controller/internal/vision"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Movement limits for a robot
type RobotLimits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Max absolute speed-up acceleration [m/s^2]
	AccSpeedupAbsoluteMax *float32 `protobuf:"fixed32,1,opt,name=acc_speedup_absolute_max,json=accSpeedupAbsoluteMax" json:"acc_speedup_absolute_max,omitempty"`
	// Max angular speed-up acceleration [rad/s^2]
	AccSpeedupAngularMax *float32 `protobuf:"fixed32,2,opt,name=acc_speedup_angular_max,json=accSpeedupAngularMax" json:"acc_speedup_angular_max,omitempty"`
	// Max absolute brake acceleration [m/s^2]
	AccBrakeAbsoluteMax *float32 `protobuf:"fixed32,3,opt,name=acc_brake_absolute_max,json=accBrakeAbsoluteMax" json:"acc_brake_absolute_max,omitempty"`
	// Max angular brake acceleration [rad/s^2]
	AccBrakeAngularMax *float32 `protobuf:"fixed32,4,opt,name=acc_brake_angular_max,json=accBrakeAngularMax" json:"acc_brake_angular_max,omitempty"`
	// Max absolute velocity [m/s]
	VelAbsoluteMax *float32 `protobuf:"fixed32,5,opt,name=vel_absolute_max,json=velAbsoluteMax" json:"vel_absolute_max,omitempty"`
	// Max angular velocity [rad/s]
	VelAngularMax *float32 `protobuf:"fixed32,6,opt,name=vel_angular_max,json=velAngularMax" json:"vel_angular_max,omitempty"`
}

func (x *RobotLimits) Reset() {
	*x = RobotLimits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_simulation_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotLimits) ProtoMessage() {}

func (x *RobotLimits) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotLimits.ProtoReflect.Descriptor instead.
func (*RobotLimits) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_config_proto_rawDescGZIP(), []int{0}
}

func (x *RobotLimits) GetAccSpeedupAbsoluteMax() float32 {
	if x != nil && x.AccSpeedupAbsoluteMax != nil {
		return *x.AccSpeedupAbsoluteMax
	}
	return 0
}

func (x *RobotLimits) GetAccSpeedupAngularMax() float32 {
	if x != nil && x.AccSpeedupAngularMax != nil {
		return *x.AccSpeedupAngularMax
	}
	return 0
}

func (x *RobotLimits) GetAccBrakeAbsoluteMax() float32 {
	if x != nil && x.AccBrakeAbsoluteMax != nil {
		return *x.AccBrakeAbsoluteMax
	}
	return 0
}

func (x *RobotLimits) GetAccBrakeAngularMax() float32 {
	if x != nil && x.AccBrakeAngularMax != nil {
		return *x.AccBrakeAngularMax
	}
	return 0
}

func (x *RobotLimits) GetVelAbsoluteMax() float32 {
	if x != nil && x.VelAbsoluteMax != nil {
		return *x.VelAbsoluteMax
	}
	return 0
}

func (x *RobotLimits) GetVelAngularMax() float32 {
	if x != nil && x.VelAngularMax != nil {
		return *x.VelAngularMax
	}
	return 0
}

// Robot wheel angle configuration
// all angles are relative to looking forward,
// all wheels / angles are clockwise
type RobotWheelAngles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Angle front right [rad]
	FrontRight *float32 `protobuf:"fixed32,1,req,name=front_right,json=frontRight" json:"front_right,omitempty"`
	// Angle back right [rad]
	BackRight *float32 `protobuf:"fixed32,2,req,name=back_right,json=backRight" json:"back_right,omitempty"`
	// Angle back left [rad]
	BackLeft *float32 `protobuf:"fixed32,3,req,name=back_left,json=backLeft" json:"back_left,omitempty"`
	// Angle front left [rad]
	FrontLeft *float32 `protobuf:"fixed32,4,req,name=front_left,json=frontLeft" json:"front_left,omitempty"`
}

func (x *RobotWheelAngles) Reset() {
	*x = RobotWheelAngles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_simulation_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotWheelAngles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotWheelAngles) ProtoMessage() {}

func (x *RobotWheelAngles) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotWheelAngles.ProtoReflect.Descriptor instead.
func (*RobotWheelAngles) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_config_proto_rawDescGZIP(), []int{1}
}

func (x *RobotWheelAngles) GetFrontRight() float32 {
	if x != nil && x.FrontRight != nil {
		return *x.FrontRight
	}
	return 0
}

func (x *RobotWheelAngles) GetBackRight() float32 {
	if x != nil && x.BackRight != nil {
		return *x.BackRight
	}
	return 0
}

func (x *RobotWheelAngles) GetBackLeft() float32 {
	if x != nil && x.BackLeft != nil {
		return *x.BackLeft
	}
	return 0
}

func (x *RobotWheelAngles) GetFrontLeft() float32 {
	if x != nil && x.FrontLeft != nil {
		return *x.FrontLeft
	}
	return 0
}

// Specs of a robot
type RobotSpecs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the robot
	Id *referee.RobotId `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// Robot radius [m]
	Radius *float32 `protobuf:"fixed32,2,opt,name=radius,def=0.09" json:"radius,omitempty"`
	// Robot height [m]
	Height *float32 `protobuf:"fixed32,3,opt,name=height,def=0.15" json:"height,omitempty"`
	// Robot mass [kg]
	Mass *float32 `protobuf:"fixed32,4,opt,name=mass" json:"mass,omitempty"`
	// Max linear kick speed [m/s] (unset = unlimited)
	MaxLinearKickSpeed *float32 `protobuf:"fixed32,7,opt,name=max_linear_kick_speed,json=maxLinearKickSpeed" json:"max_linear_kick_speed,omitempty"`
	// Max chip kick speed [m/s] (unset = unlimited)
	MaxChipKickSpeed *float32 `protobuf:"fixed32,8,opt,name=max_chip_kick_speed,json=maxChipKickSpeed" json:"max_chip_kick_speed,omitempty"`
	// Width of the dribbler [m] (implicitly defines the distance from robot center to dribbler and opening angle)
	DribblerWidth *float32 `protobuf:"fixed32,9,opt,name=dribbler_width,json=dribblerWidth" json:"dribbler_width,omitempty"`
	// Movement limits
	Limits *RobotLimits `protobuf:"bytes,10,opt,name=limits" json:"limits,omitempty"`
	// Wheel angle configuration
	WheelAngles *RobotWheelAngles `protobuf:"bytes,13,opt,name=wheel_angles,json=wheelAngles" json:"wheel_angles,omitempty"`
	// Custom robot spec for specific simulators (the protobuf files are managed by the simulators)
	Custom *any.Any `protobuf:"bytes,14,opt,name=custom" json:"custom,omitempty"`
}

// Default values for RobotSpecs fields.
const (
	Default_RobotSpecs_Radius = float32(0.09000000357627869)
	Default_RobotSpecs_Height = float32(0.15000000596046448)
)

func (x *RobotSpecs) Reset() {
	*x = RobotSpecs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_simulation_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotSpecs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotSpecs) ProtoMessage() {}

func (x *RobotSpecs) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotSpecs.ProtoReflect.Descriptor instead.
func (*RobotSpecs) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_config_proto_rawDescGZIP(), []int{2}
}

func (x *RobotSpecs) GetId() *referee.RobotId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *RobotSpecs) GetRadius() float32 {
	if x != nil && x.Radius != nil {
		return *x.Radius
	}
	return Default_RobotSpecs_Radius
}

func (x *RobotSpecs) GetHeight() float32 {
	if x != nil && x.Height != nil {
		return *x.Height
	}
	return Default_RobotSpecs_Height
}

func (x *RobotSpecs) GetMass() float32 {
	if x != nil && x.Mass != nil {
		return *x.Mass
	}
	return 0
}

func (x *RobotSpecs) GetMaxLinearKickSpeed() float32 {
	if x != nil && x.MaxLinearKickSpeed != nil {
		return *x.MaxLinearKickSpeed
	}
	return 0
}

func (x *RobotSpecs) GetMaxChipKickSpeed() float32 {
	if x != nil && x.MaxChipKickSpeed != nil {
		return *x.MaxChipKickSpeed
	}
	return 0
}

func (x *RobotSpecs) GetDribblerWidth() float32 {
	if x != nil && x.DribblerWidth != nil {
		return *x.DribblerWidth
	}
	return 0
}

func (x *RobotSpecs) GetLimits() *RobotLimits {
	if x != nil {
		return x.Limits
	}
	return nil
}

func (x *RobotSpecs) GetWheelAngles() *RobotWheelAngles {
	if x != nil {
		return x.WheelAngles
	}
	return nil
}

func (x *RobotSpecs) GetCustom() *any.Any {
	if x != nil {
		return x.Custom
	}
	return nil
}

type RealismConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Custom config for specific simulators (the protobuf files are managed by the simulators)
	Custom *any.Any `protobuf:"bytes,1,opt,name=custom" json:"custom,omitempty"`
}

func (x *RealismConfig) Reset() {
	*x = RealismConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_simulation_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RealismConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RealismConfig) ProtoMessage() {}

func (x *RealismConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RealismConfig.ProtoReflect.Descriptor instead.
func (*RealismConfig) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_config_proto_rawDescGZIP(), []int{3}
}

func (x *RealismConfig) GetCustom() *any.Any {
	if x != nil {
		return x.Custom
	}
	return nil
}

// Change the simulator configuration
type SimulatorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Update the geometry
	Geometry *vision.SSL_GeometryData `protobuf:"bytes,1,opt,name=geometry" json:"geometry,omitempty"`
	// Update the robot specs
	RobotSpecs []*RobotSpecs `protobuf:"bytes,2,rep,name=robot_specs,json=robotSpecs" json:"robot_specs,omitempty"`
	// Update realism configuration
	RealismConfig *RealismConfig `protobuf:"bytes,3,opt,name=realism_config,json=realismConfig" json:"realism_config,omitempty"`
	// Change the vision publish port
	VisionPort *uint32 `protobuf:"varint,4,opt,name=vision_port,json=visionPort" json:"vision_port,omitempty"`
}

func (x *SimulatorConfig) Reset() {
	*x = SimulatorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_simulation_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulatorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulatorConfig) ProtoMessage() {}

func (x *SimulatorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_simulation_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulatorConfig.ProtoReflect.Descriptor instead.
func (*SimulatorConfig) Descriptor() ([]byte, []int) {
	return file_ssl_simulation_config_proto_rawDescGZIP(), []int{4}
}

func (x *SimulatorConfig) GetGeometry() *vision.SSL_GeometryData {
	if x != nil {
		return x.Geometry
	}
	return nil
}

func (x *SimulatorConfig) GetRobotSpecs() []*RobotSpecs {
	if x != nil {
		return x.RobotSpecs
	}
	return nil
}

func (x *SimulatorConfig) GetRealismConfig() *RealismConfig {
	if x != nil {
		return x.RealismConfig
	}
	return nil
}

func (x *SimulatorConfig) GetVisionPort() uint32 {
	if x != nil && x.VisionPort != nil {
		return *x.VisionPort
	}
	return 0
}

var File_ssl_simulation_config_proto protoreflect.FileDescriptor

var file_ssl_simulation_config_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x73, 0x6c, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x65,
	0x64, 0x75, 0x2e, 0x74, 0x69, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x75, 0x6d, 0x61, 0x74, 0x72,
	0x61, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x73, 0x73,
	0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x73, 0x73, 0x6c, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65,
	0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x02, 0x0a, 0x0b, 0x52, 0x6f, 0x62, 0x6f,
	0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x61, 0x63, 0x63, 0x5f, 0x73,
	0x70, 0x65, 0x65, 0x64, 0x75, 0x70, 0x5f, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x15, 0x61, 0x63, 0x63, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x75, 0x70, 0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x78,
	0x12, 0x35, 0x0a, 0x17, 0x61, 0x63, 0x63, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x75, 0x70, 0x5f,
	0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x14, 0x61, 0x63, 0x63, 0x53, 0x70, 0x65, 0x65, 0x64, 0x75, 0x70, 0x41, 0x6e, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x4d, 0x61, 0x78, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x63, 0x63, 0x5f, 0x62,
	0x72, 0x61, 0x6b, 0x65, 0x5f, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x61, 0x63, 0x63, 0x42, 0x72, 0x61, 0x6b,
	0x65, 0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x78, 0x12, 0x31, 0x0a, 0x15,
	0x61, 0x63, 0x63, 0x5f, 0x62, 0x72, 0x61, 0x6b, 0x65, 0x5f, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61,
	0x72, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x61, 0x63, 0x63,
	0x42, 0x72, 0x61, 0x6b, 0x65, 0x41, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x61, 0x78, 0x12,
	0x28, 0x0a, 0x10, 0x76, 0x65, 0x6c, 0x5f, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x76, 0x65, 0x6c, 0x41, 0x62,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x78, 0x12, 0x26, 0x0a, 0x0f, 0x76, 0x65, 0x6c,
	0x5f, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0d, 0x76, 0x65, 0x6c, 0x41, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x61,
	0x78, 0x22, 0x8e, 0x01, 0x0a, 0x10, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x57, 0x68, 0x65, 0x65, 0x6c,
	0x41, 0x6e, 0x67, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f,
	0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x5f,
	0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x02, 0x52, 0x09, 0x62, 0x61, 0x63,
	0x6b, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x6c,
	0x65, 0x66, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x02, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x4c,
	0x65, 0x66, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x66,
	0x74, 0x18, 0x04, 0x20, 0x02, 0x28, 0x02, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x4c, 0x65,
	0x66, 0x74, 0x22, 0xc5, 0x03, 0x0a, 0x0a, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x73, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x52, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x06, 0x72,
	0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x04, 0x30, 0x2e, 0x30,
	0x39, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x04, 0x30, 0x2e, 0x31, 0x35, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x6d,
	0x61, 0x78, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x5f, 0x6b, 0x69, 0x63, 0x6b, 0x5f, 0x73,
	0x70, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x4c,
	0x69, 0x6e, 0x65, 0x61, 0x72, 0x4b, 0x69, 0x63, 0x6b, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x2d,
	0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x68, 0x69, 0x70, 0x5f, 0x6b, 0x69, 0x63, 0x6b, 0x5f,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x6d, 0x61, 0x78,
	0x43, 0x68, 0x69, 0x70, 0x4b, 0x69, 0x63, 0x6b, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x72, 0x69, 0x62, 0x62, 0x6c, 0x65, 0x72, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x64, 0x72, 0x69, 0x62, 0x62, 0x6c, 0x65, 0x72, 0x57,
	0x69, 0x64, 0x74, 0x68, 0x12, 0x42, 0x0a, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x64, 0x75, 0x2e, 0x74, 0x69, 0x67, 0x65, 0x72,
	0x73, 0x2e, 0x73, 0x75, 0x6d, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x52, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x52, 0x0a, 0x0c, 0x77, 0x68, 0x65, 0x65,
	0x6c, 0x5f, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x65, 0x64, 0x75, 0x2e, 0x74, 0x69, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x75, 0x6d, 0x61,
	0x74, 0x72, 0x61, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x6f, 0x62, 0x6f, 0x74, 0x57, 0x68, 0x65, 0x65, 0x6c, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x73, 0x52,
	0x0b, 0x77, 0x68, 0x65, 0x65, 0x6c, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x06,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x22, 0x3d, 0x0a, 0x0d, 0x52, 0x65,
	0x61, 0x6c, 0x69, 0x73, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2c, 0x0a, 0x06, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x22, 0x82, 0x02, 0x0a, 0x0f, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x0a,
	0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x53, 0x53, 0x4c, 0x5f, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x4a, 0x0a, 0x0b,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x65, 0x64, 0x75, 0x2e, 0x74, 0x69, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x73,
	0x75, 0x6d, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x53, 0x70, 0x65, 0x63, 0x73, 0x52, 0x0a, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x53, 0x70, 0x65, 0x63, 0x73, 0x12, 0x53, 0x0a, 0x0e, 0x72, 0x65, 0x61, 0x6c,
	0x69, 0x73, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x65, 0x64, 0x75, 0x2e, 0x74, 0x69, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x75,
	0x6d, 0x61, 0x74, 0x72, 0x61, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d,
	0x72, 0x65, 0x61, 0x6c, 0x69, 0x73, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a,
	0x0b, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x42,
	0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62,
	0x6f, 0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x73, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x69, 0x6d, 0x63,
	0x74, 0x6c,
}

var (
	file_ssl_simulation_config_proto_rawDescOnce sync.Once
	file_ssl_simulation_config_proto_rawDescData = file_ssl_simulation_config_proto_rawDesc
)

func file_ssl_simulation_config_proto_rawDescGZIP() []byte {
	file_ssl_simulation_config_proto_rawDescOnce.Do(func() {
		file_ssl_simulation_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_simulation_config_proto_rawDescData)
	})
	return file_ssl_simulation_config_proto_rawDescData
}

var file_ssl_simulation_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ssl_simulation_config_proto_goTypes = []interface{}{
	(*RobotLimits)(nil),             // 0: edu.tigers.sumatra.simulation.RobotLimits
	(*RobotWheelAngles)(nil),        // 1: edu.tigers.sumatra.simulation.RobotWheelAngles
	(*RobotSpecs)(nil),              // 2: edu.tigers.sumatra.simulation.RobotSpecs
	(*RealismConfig)(nil),           // 3: edu.tigers.sumatra.simulation.RealismConfig
	(*SimulatorConfig)(nil),         // 4: edu.tigers.sumatra.simulation.SimulatorConfig
	(*referee.RobotId)(nil),         // 5: RobotId
	(*any.Any)(nil),                 // 6: google.protobuf.Any
	(*vision.SSL_GeometryData)(nil), // 7: SSL_GeometryData
}
var file_ssl_simulation_config_proto_depIdxs = []int32{
	5, // 0: edu.tigers.sumatra.simulation.RobotSpecs.id:type_name -> RobotId
	0, // 1: edu.tigers.sumatra.simulation.RobotSpecs.limits:type_name -> edu.tigers.sumatra.simulation.RobotLimits
	1, // 2: edu.tigers.sumatra.simulation.RobotSpecs.wheel_angles:type_name -> edu.tigers.sumatra.simulation.RobotWheelAngles
	6, // 3: edu.tigers.sumatra.simulation.RobotSpecs.custom:type_name -> google.protobuf.Any
	6, // 4: edu.tigers.sumatra.simulation.RealismConfig.custom:type_name -> google.protobuf.Any
	7, // 5: edu.tigers.sumatra.simulation.SimulatorConfig.geometry:type_name -> SSL_GeometryData
	2, // 6: edu.tigers.sumatra.simulation.SimulatorConfig.robot_specs:type_name -> edu.tigers.sumatra.simulation.RobotSpecs
	3, // 7: edu.tigers.sumatra.simulation.SimulatorConfig.realism_config:type_name -> edu.tigers.sumatra.simulation.RealismConfig
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_ssl_simulation_config_proto_init() }
func file_ssl_simulation_config_proto_init() {
	if File_ssl_simulation_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ssl_simulation_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotLimits); i {
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
		file_ssl_simulation_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotWheelAngles); i {
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
		file_ssl_simulation_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotSpecs); i {
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
		file_ssl_simulation_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RealismConfig); i {
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
		file_ssl_simulation_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulatorConfig); i {
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
			RawDescriptor: file_ssl_simulation_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_simulation_config_proto_goTypes,
		DependencyIndexes: file_ssl_simulation_config_proto_depIdxs,
		MessageInfos:      file_ssl_simulation_config_proto_msgTypes,
	}.Build()
	File_ssl_simulation_config_proto = out.File
	file_ssl_simulation_config_proto_rawDesc = nil
	file_ssl_simulation_config_proto_goTypes = nil
	file_ssl_simulation_config_proto_depIdxs = nil
}
