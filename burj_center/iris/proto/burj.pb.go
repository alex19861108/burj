// Code generated by protoc-gen-go. DO NOT EDIT.
// source: burj.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	burj.proto

It has these top-level messages:
	CPU
	Memory
	Device
	Addr
	Node
	User
	Tag
	Image
	Apk
	SubJob
	Job
	TriggerJobRequest
	TriggerJobResponse
	SearchNodeRequest
	NodeDeviceUnit
	SearchNodeResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	Status_PREPARE Status = 0
	Status_RUNNING Status = 1
	Status_SUCCESS Status = 2
	Status_ERROR   Status = 3
)

var Status_name = map[int32]string{
	0: "PREPARE",
	1: "RUNNING",
	2: "SUCCESS",
	3: "ERROR",
}
var Status_value = map[string]int32{
	"PREPARE": 0,
	"RUNNING": 1,
	"SUCCESS": 2,
	"ERROR":   3,
}

func (x Status) String() string {
	return proto1.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CPU struct {
	// @inject_tag: bson:"abi,omitempty"
	Abi string `protobuf:"bytes,1,opt,name=abi" json:"abi,omitempty" bson:"abi,omitempty"`
	// @inject_tag: bson:"abilist,omitempty"
	Abilist string `protobuf:"bytes,2,opt,name=abilist" json:"abilist,omitempty" bson:"abilist,omitempty"`
	// @inject_tag: bson:"abilist32,omitempty"
	Abilist32 string `protobuf:"bytes,3,opt,name=abilist32" json:"abilist32,omitempty" bson:"abilist32,omitempty"`
	// @inject_tag: bson:"abilist64,omitempty"
	Abilist64 string `protobuf:"bytes,4,opt,name=abilist64" json:"abilist64,omitempty" bson:"abilist64,omitempty"`
}

func (m *CPU) Reset()                    { *m = CPU{} }
func (m *CPU) String() string            { return proto1.CompactTextString(m) }
func (*CPU) ProtoMessage()               {}
func (*CPU) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CPU) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *CPU) GetAbilist() string {
	if m != nil {
		return m.Abilist
	}
	return ""
}

func (m *CPU) GetAbilist32() string {
	if m != nil {
		return m.Abilist32
	}
	return ""
}

func (m *CPU) GetAbilist64() string {
	if m != nil {
		return m.Abilist64
	}
	return ""
}

type Memory struct {
	// @inject_tag: bson:"total,omitempty"
	Total string `protobuf:"bytes,1,opt,name=total" json:"total,omitempty" bson:"total,omitempty"`
	// @inject_tag: bson:"used,omitempty"
	Used string `protobuf:"bytes,2,opt,name=used" json:"used,omitempty" bson:"used,omitempty"`
	// @inject_tag: bson:"available,omitempty"
	Available string `protobuf:"bytes,3,opt,name=available" json:"available,omitempty" bson:"available,omitempty"`
}

func (m *Memory) Reset()                    { *m = Memory{} }
func (m *Memory) String() string            { return proto1.CompactTextString(m) }
func (*Memory) ProtoMessage()               {}
func (*Memory) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Memory) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

func (m *Memory) GetUsed() string {
	if m != nil {
		return m.Used
	}
	return ""
}

func (m *Memory) GetAvailable() string {
	if m != nil {
		return m.Available
	}
	return ""
}

type Device struct {
	// @inject_tag: bson:"_id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" bson:"_id,omitempty"`
	// @inject_tag: bson:"addr,omitempty"
	Addr string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty" bson:"addr,omitempty"`
	// @inject_tag: bson:"label,omitempty"
	Label string `protobuf:"bytes,3,opt,name=label" json:"label,omitempty" bson:"label,omitempty"`
	// @inject_tag: bson:"cpu,omitempty"
	Cpu *CPU `protobuf:"bytes,4,opt,name=cpu" json:"cpu,omitempty" bson:"cpu,omitempty"`
	// @inject_tag: bson:"memory,omitempty"
	Memory *Memory `protobuf:"bytes,5,opt,name=memory" json:"memory,omitempty" bson:"memory,omitempty"`
	// @inject_tag: bson:"model,omitempty"
	Model string `protobuf:"bytes,6,opt,name=model" json:"model,omitempty" bson:"model,omitempty"`
	// @inject_tag: bson:"brand,omitempty"
	Brand string `protobuf:"bytes,7,opt,name=brand" json:"brand,omitempty" bson:"brand,omitempty"`
	// @inject_tag: bson:"name,omitempty"
	Name string `protobuf:"bytes,8,opt,name=name" json:"name,omitempty" bson:"name,omitempty"`
	// @inject_tag: bson:"device,omitempty"
	Device string `protobuf:"bytes,9,opt,name=device" json:"device,omitempty" bson:"device,omitempty"`
	// @inject_tag: bson:"board,omitempty"
	Board string `protobuf:"bytes,10,opt,name=board" json:"board,omitempty" bson:"board,omitempty"`
	// @inject_tag: bson:"locale,omitempty"
	Locale string `protobuf:"bytes,11,opt,name=locale" json:"locale,omitempty" bson:"locale,omitempty"`
	// @inject_tag: bson:"first_api_level,omitempty"
	FirstApiLevel string `protobuf:"bytes,12,opt,name=firstApiLevel" json:"firstApiLevel,omitempty" bson:"first_api_level,omitempty"`
	// @inject_tag: bson:"manufacturer,omitempty"
	Manufacturer string `protobuf:"bytes,13,opt,name=manufacturer" json:"manufacturer,omitempty" bson:"manufacturer,omitempty"`
	// @inject_tag: bson:"cuptsm,omitempty"
	Cuptsm string `protobuf:"bytes,14,opt,name=cuptsm" json:"cuptsm,omitempty" bson:"cuptsm,omitempty"`
	// @inject_tag: bson:"serial,omitempty"
	Serial string `protobuf:"bytes,15,opt,name=serial" json:"serial,omitempty" bson:"serial,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto1.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Device) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Device) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Device) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Device) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Device) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *Device) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Device) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *Device) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func (m *Device) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *Device) GetFirstApiLevel() string {
	if m != nil {
		return m.FirstApiLevel
	}
	return ""
}

func (m *Device) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *Device) GetCuptsm() string {
	if m != nil {
		return m.Cuptsm
	}
	return ""
}

func (m *Device) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

type Addr struct {
	// @inject_tag: bson:"http,omitempty"
	Http string `protobuf:"bytes,1,opt,name=http" json:"http,omitempty" bson:"http,omitempty"`
	// @inject_tag: bson:"rpc,omitempty"
	Rpc string `protobuf:"bytes,2,opt,name=rpc" json:"rpc,omitempty" bson:"rpc,omitempty"`
}

func (m *Addr) Reset()                    { *m = Addr{} }
func (m *Addr) String() string            { return proto1.CompactTextString(m) }
func (*Addr) ProtoMessage()               {}
func (*Addr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Addr) GetHttp() string {
	if m != nil {
		return m.Http
	}
	return ""
}

func (m *Addr) GetRpc() string {
	if m != nil {
		return m.Rpc
	}
	return ""
}

type Node struct {
	// @inject_tag: bson:"_id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" bson:"_id,omitempty"`
	// @inject_tag: bson:"addr,omitempty"
	Addr *Addr `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty" bson:"addr,omitempty"`
	// @inject_tag: bson:"labels,omitempty"
	Labels string `protobuf:"bytes,3,opt,name=labels" json:"labels,omitempty" bson:"labels,omitempty"`
	// @inject_tag: bson:"devices,omitempty"
	Devices []*Device `protobuf:"bytes,4,rep,name=devices" json:"devices,omitempty" bson:"devices,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto1.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddr() *Addr {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *Node) GetLabels() string {
	if m != nil {
		return m.Labels
	}
	return ""
}

func (m *Node) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type User struct {
	// @inject_tag: xorm:"int64 notnull pk autoincr"
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty" xorm:"int64 notnull pk autoincr"`
	// @inject_tag: xorm:"varchar(20) notnull"
	UserName string `protobuf:"bytes,2,opt,name=userName" json:"userName,omitempty" xorm:"varchar(20) notnull"`
	// @inject_tag: xorm:"varchar(50) notnull"
	PassWord string `protobuf:"bytes,3,opt,name=passWord" json:"passWord,omitempty" xorm:"varchar(50) notnull"`
	// @inject_tag: xorm:"varchar(20) notnull"
	Role string `protobuf:"bytes,4,opt,name=role" json:"role,omitempty" xorm:"varchar(20) notnull"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type Tag struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// @inject_tag: xorm:"varchar(100) notnull"
	Mark string `protobuf:"bytes,2,opt,name=mark" json:"mark,omitempty" xorm:"varchar(100) notnull"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto1.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Tag) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tag) GetMark() string {
	if m != nil {
		return m.Mark
	}
	return ""
}

type Image struct {
	// @inject_tag: xorm:"varchar(50) notnull pk"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" xorm:"varchar(50) notnull pk"`
	// @inject_tag: xorm:"varchar(50)"
	Mark string `protobuf:"bytes,2,opt,name=mark" json:"mark,omitempty" xorm:"varchar(50)"`
	// @inject_tag: xorm:"varchar(100) notnull"
	Path string `protobuf:"bytes,3,opt,name=path" json:"path,omitempty" xorm:"varchar(100) notnull"`
	// @inject_tag: xorm:"varchar(40) notnull"
	LastUpdateBy string `protobuf:"bytes,4,opt,name=last_update_by,json=lastUpdateBy" json:"last_update_by,omitempty" xorm:"varchar(40) notnull"`
	// @inject_tag: xorm:"date"
	LastUpdate string `protobuf:"bytes,5,opt,name=last_update,json=lastUpdate" json:"last_update,omitempty" xorm:"date"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto1.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetMark() string {
	if m != nil {
		return m.Mark
	}
	return ""
}

func (m *Image) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Image) GetLastUpdateBy() string {
	if m != nil {
		return m.LastUpdateBy
	}
	return ""
}

func (m *Image) GetLastUpdate() string {
	if m != nil {
		return m.LastUpdate
	}
	return ""
}

type Apk struct {
	// @inject_tag: bson:"path,omitempty"
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty" bson:"path,omitempty"`
	// @inject_tag: bson:"pkg_name,omitempty"
	PkgName string `protobuf:"bytes,2,opt,name=pkgName" json:"pkgName,omitempty" bson:"pkg_name,omitempty"`
	// @inject_tag: bson:"class_name,omitempty"
	ClassName string `protobuf:"bytes,3,opt,name=className" json:"className,omitempty" bson:"class_name,omitempty"`
}

func (m *Apk) Reset()                    { *m = Apk{} }
func (m *Apk) String() string            { return proto1.CompactTextString(m) }
func (*Apk) ProtoMessage()               {}
func (*Apk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Apk) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Apk) GetPkgName() string {
	if m != nil {
		return m.PkgName
	}
	return ""
}

func (m *Apk) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

type SubJob struct {
	// @inject_tag: bson:"_id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" bson:"_id,omitempty"`
	// @inject_tag: bson:"images,omitempty"
	Images []*Image `protobuf:"bytes,2,rep,name=images" json:"images,omitempty" bson:"images,omitempty"`
	// @inject_tag: bson:"node,omitempty"
	Node *Node `protobuf:"bytes,3,opt,name=node" json:"node,omitempty" bson:"node,omitempty"`
	// @inject_tag: bson:"device,omitempty"
	Device *Device `protobuf:"bytes,4,opt,name=device" json:"device,omitempty" bson:"device,omitempty"`
	// @inject_tag: bson:"status,omitempty"
	Status Status `protobuf:"varint,5,opt,name=status,enum=proto.Status" json:"status,omitempty" bson:"status,omitempty"`
	// @inject_tag: bson:"update_time,omitempty"
	UpdateTime string `protobuf:"bytes,6,opt,name=updateTime" json:"updateTime,omitempty" bson:"update_time,omitempty"`
	// @inject_tag: bson:"log,omitempty"
	Log string `protobuf:"bytes,7,opt,name=log" json:"log,omitempty" bson:"log,omitempty"`
}

func (m *SubJob) Reset()                    { *m = SubJob{} }
func (m *SubJob) String() string            { return proto1.CompactTextString(m) }
func (*SubJob) ProtoMessage()               {}
func (*SubJob) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SubJob) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SubJob) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *SubJob) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *SubJob) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *SubJob) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PREPARE
}

func (m *SubJob) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *SubJob) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

type Job struct {
	// @inject_tag: bson:"_id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty" bson:"_id,omitempty"`
	// @inject_tag: bson:"name,omitempty"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" bson:"name,omitempty"`
	// @inject_tag: bson:"apk,omitempty"
	Apk *Apk `protobuf:"bytes,3,opt,name=Apk" json:"Apk,omitempty" bson:"apk,omitempty"`
	// @inject_tag: bson:"subjobs,omitempty"
	SubJobs []*SubJob `protobuf:"bytes,4,rep,name=subJobs" json:"subJobs,omitempty" bson:"subjobs,omitempty"`
	// @inject_tag: bson:"owner,omitempty"
	Owner string `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty" bson:"owner,omitempty"`
	// @inject_tag: bson:"status,omitempty"
	Status Status `protobuf:"varint,6,opt,name=status,enum=proto.Status" json:"status,omitempty" bson:"status,omitempty"`
	// @inject_tag: bson:"update_time,omitempty"
	UpdateTime *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=updateTime" json:"updateTime,omitempty" bson:"update_time,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto1.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Job) GetApk() *Apk {
	if m != nil {
		return m.Apk
	}
	return nil
}

func (m *Job) GetSubJobs() []*SubJob {
	if m != nil {
		return m.SubJobs
	}
	return nil
}

func (m *Job) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Job) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PREPARE
}

func (m *Job) GetUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type TriggerJobRequest struct {
	Job    *Job    `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	SubJob *SubJob `protobuf:"bytes,2,opt,name=subJob" json:"subJob,omitempty"`
}

func (m *TriggerJobRequest) Reset()                    { *m = TriggerJobRequest{} }
func (m *TriggerJobRequest) String() string            { return proto1.CompactTextString(m) }
func (*TriggerJobRequest) ProtoMessage()               {}
func (*TriggerJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *TriggerJobRequest) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *TriggerJobRequest) GetSubJob() *SubJob {
	if m != nil {
		return m.SubJob
	}
	return nil
}

type TriggerJobResponse struct {
	Job    *Job    `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	SubJob *SubJob `protobuf:"bytes,2,opt,name=subJob" json:"subJob,omitempty"`
}

func (m *TriggerJobResponse) Reset()                    { *m = TriggerJobResponse{} }
func (m *TriggerJobResponse) String() string            { return proto1.CompactTextString(m) }
func (*TriggerJobResponse) ProtoMessage()               {}
func (*TriggerJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *TriggerJobResponse) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *TriggerJobResponse) GetSubJob() *SubJob {
	if m != nil {
		return m.SubJob
	}
	return nil
}

type SearchNodeRequest struct {
	Serial string `protobuf:"bytes,1,opt,name=serial" json:"serial,omitempty"`
}

func (m *SearchNodeRequest) Reset()                    { *m = SearchNodeRequest{} }
func (m *SearchNodeRequest) String() string            { return proto1.CompactTextString(m) }
func (*SearchNodeRequest) ProtoMessage()               {}
func (*SearchNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SearchNodeRequest) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

type NodeDeviceUnit struct {
	Node   *Node   `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	Device *Device `protobuf:"bytes,2,opt,name=device" json:"device,omitempty"`
}

func (m *NodeDeviceUnit) Reset()                    { *m = NodeDeviceUnit{} }
func (m *NodeDeviceUnit) String() string            { return proto1.CompactTextString(m) }
func (*NodeDeviceUnit) ProtoMessage()               {}
func (*NodeDeviceUnit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *NodeDeviceUnit) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *NodeDeviceUnit) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type SearchNodeResponse struct {
	NodeDeviceUnits []*NodeDeviceUnit `protobuf:"bytes,1,rep,name=nodeDeviceUnits" json:"nodeDeviceUnits,omitempty"`
}

func (m *SearchNodeResponse) Reset()                    { *m = SearchNodeResponse{} }
func (m *SearchNodeResponse) String() string            { return proto1.CompactTextString(m) }
func (*SearchNodeResponse) ProtoMessage()               {}
func (*SearchNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *SearchNodeResponse) GetNodeDeviceUnits() []*NodeDeviceUnit {
	if m != nil {
		return m.NodeDeviceUnits
	}
	return nil
}

func init() {
	proto1.RegisterType((*CPU)(nil), "proto.CPU")
	proto1.RegisterType((*Memory)(nil), "proto.Memory")
	proto1.RegisterType((*Device)(nil), "proto.Device")
	proto1.RegisterType((*Addr)(nil), "proto.Addr")
	proto1.RegisterType((*Node)(nil), "proto.Node")
	proto1.RegisterType((*User)(nil), "proto.User")
	proto1.RegisterType((*Tag)(nil), "proto.Tag")
	proto1.RegisterType((*Image)(nil), "proto.Image")
	proto1.RegisterType((*Apk)(nil), "proto.Apk")
	proto1.RegisterType((*SubJob)(nil), "proto.SubJob")
	proto1.RegisterType((*Job)(nil), "proto.Job")
	proto1.RegisterType((*TriggerJobRequest)(nil), "proto.TriggerJobRequest")
	proto1.RegisterType((*TriggerJobResponse)(nil), "proto.TriggerJobResponse")
	proto1.RegisterType((*SearchNodeRequest)(nil), "proto.SearchNodeRequest")
	proto1.RegisterType((*NodeDeviceUnit)(nil), "proto.NodeDeviceUnit")
	proto1.RegisterType((*SearchNodeResponse)(nil), "proto.SearchNodeResponse")
	proto1.RegisterEnum("proto.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for JobService service

type JobServiceClient interface {
	Trigger(ctx context.Context, in *TriggerJobRequest, opts ...grpc.CallOption) (*TriggerJobResponse, error)
}

type jobServiceClient struct {
	cc *grpc.ClientConn
}

func NewJobServiceClient(cc *grpc.ClientConn) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) Trigger(ctx context.Context, in *TriggerJobRequest, opts ...grpc.CallOption) (*TriggerJobResponse, error) {
	out := new(TriggerJobResponse)
	err := grpc.Invoke(ctx, "/proto.JobService/trigger", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JobService service

type JobServiceServer interface {
	Trigger(context.Context, *TriggerJobRequest) (*TriggerJobResponse, error)
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_Trigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).Trigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.JobService/Trigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).Trigger(ctx, req.(*TriggerJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "trigger",
			Handler:    _JobService_Trigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "burj.proto",
}

// Client API for NodeService service

type NodeServiceClient interface {
	Search(ctx context.Context, in *SearchNodeRequest, opts ...grpc.CallOption) (*SearchNodeResponse, error)
}

type nodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodeServiceClient(cc *grpc.ClientConn) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Search(ctx context.Context, in *SearchNodeRequest, opts ...grpc.CallOption) (*SearchNodeResponse, error) {
	out := new(SearchNodeResponse)
	err := grpc.Invoke(ctx, "/proto.NodeService/search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeService service

type NodeServiceServer interface {
	Search(context.Context, *SearchNodeRequest) (*SearchNodeResponse, error)
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Search(ctx, req.(*SearchNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "search",
			Handler:    _NodeService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "burj.proto",
}

func init() { proto1.RegisterFile("burj.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 976 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0x45, 0x99, 0x8a, 0x46, 0xb6, 0xa3, 0x2c, 0xda, 0x62, 0x6b, 0x18, 0x75, 0x40, 0x38,
	0x48, 0xff, 0xa0, 0x00, 0x4a, 0x90, 0x43, 0x2e, 0xa9, 0xe2, 0x1a, 0x45, 0x82, 0x56, 0x55, 0x29,
	0x0b, 0x4d, 0x4f, 0xc1, 0x52, 0x5c, 0xcb, 0x8c, 0x49, 0x2d, 0xbb, 0x4b, 0xba, 0xf5, 0xad, 0xa7,
	0xb6, 0xc7, 0xbe, 0x61, 0xcf, 0x7d, 0x8b, 0x62, 0x66, 0x97, 0x12, 0x65, 0x19, 0xe9, 0xa5, 0x27,
	0xce, 0x1f, 0x66, 0xe6, 0x9b, 0x6f, 0x86, 0x0b, 0x10, 0x57, 0xfa, 0xdd, 0xa0, 0xd0, 0xaa, 0x54,
	0x6c, 0x87, 0x3e, 0x07, 0x47, 0x0b, 0xa5, 0x16, 0x99, 0x7c, 0x4c, 0x5a, 0x5c, 0x9d, 0x3f, 0x2e,
	0xd3, 0x5c, 0x9a, 0x52, 0xe4, 0x85, 0x8d, 0x0b, 0x73, 0xf0, 0x4f, 0x26, 0x33, 0xd6, 0x07, 0x5f,
	0xc4, 0x29, 0xf7, 0x1e, 0x78, 0x9f, 0x76, 0x23, 0x14, 0x19, 0x87, 0x8e, 0x88, 0xd3, 0x2c, 0x35,
	0x25, 0x6f, 0x91, 0xb5, 0x56, 0xd9, 0x21, 0x74, 0x9d, 0xf8, 0x64, 0xc8, 0x7d, 0xf2, 0xad, 0x0d,
	0x0d, 0xef, 0xb3, 0xa7, 0xbc, 0xbd, 0xe1, 0x7d, 0xf6, 0x34, 0x9c, 0x40, 0xf0, 0x9d, 0xcc, 0x95,
	0xbe, 0x66, 0x1f, 0xc0, 0x4e, 0xa9, 0x4a, 0x91, 0xb9, 0x9a, 0x56, 0x61, 0x0c, 0xda, 0x95, 0x91,
	0x89, 0x2b, 0x49, 0x32, 0x65, 0xbc, 0x12, 0x69, 0x26, 0xe2, 0x4c, 0xae, 0xea, 0xd5, 0x86, 0xf0,
	0x0f, 0x1f, 0x82, 0xaf, 0xe5, 0x55, 0x3a, 0x97, 0x6c, 0x1f, 0x5a, 0x69, 0xe2, 0xf2, 0xb5, 0xd2,
	0x04, 0x93, 0x89, 0x24, 0xd1, 0x75, 0x32, 0x94, 0xb1, 0x6c, 0x26, 0x62, 0x99, 0xb9, 0x44, 0x56,
	0x61, 0x87, 0xe0, 0xcf, 0x8b, 0x8a, 0xda, 0xed, 0x0d, 0xc1, 0x8e, 0x66, 0x70, 0x32, 0x99, 0x45,
	0x68, 0x66, 0x0f, 0x21, 0xc8, 0xa9, 0x69, 0xbe, 0x43, 0x01, 0x7b, 0x2e, 0xc0, 0x22, 0x89, 0x9c,
	0x13, 0x53, 0xe7, 0x2a, 0x91, 0x19, 0x0f, 0x6c, 0x6a, 0x52, 0xd0, 0x1a, 0x6b, 0xb1, 0x4c, 0x78,
	0xc7, 0x5a, 0x49, 0xc1, 0xd6, 0x96, 0x22, 0x97, 0xfc, 0xae, 0x6d, 0x0d, 0x65, 0xf6, 0x11, 0x04,
	0x09, 0x01, 0xe1, 0x5d, 0xb2, 0x3a, 0x8d, 0x32, 0x28, 0xa1, 0x13, 0x0e, 0x2e, 0x03, 0x2a, 0x18,
	0x9d, 0xa9, 0xb9, 0xc8, 0x24, 0xef, 0xd9, 0x68, 0xab, 0xb1, 0x63, 0xd8, 0x3b, 0x4f, 0xb5, 0x29,
	0x47, 0x45, 0xfa, 0xad, 0xbc, 0x92, 0x19, 0xdf, 0x25, 0xf7, 0xa6, 0x91, 0x85, 0xb0, 0x9b, 0x8b,
	0x65, 0x75, 0x2e, 0xe6, 0x65, 0xa5, 0xa5, 0xe6, 0x7b, 0x14, 0xb4, 0x61, 0xc3, 0x0a, 0xf3, 0xaa,
	0x28, 0x4d, 0xce, 0xf7, 0x6d, 0x05, 0xab, 0xa1, 0xdd, 0x48, 0x9d, 0x8a, 0x8c, 0xdf, 0xb3, 0x76,
	0xab, 0x85, 0x5f, 0x42, 0x7b, 0x84, 0x23, 0x66, 0xd0, 0xbe, 0x28, 0xcb, 0xc2, 0x11, 0x41, 0x32,
	0xee, 0x97, 0x2e, 0xe6, 0x8e, 0x09, 0x14, 0xc3, 0x5f, 0xa1, 0x3d, 0x56, 0xc9, 0x36, 0x69, 0x47,
	0x0d, 0xd2, 0x7a, 0xc3, 0x9e, 0x1b, 0x35, 0x26, 0x76, 0x0c, 0x22, 0x70, 0x24, 0xcd, 0x38, 0x0a,
	0x9d, 0xc6, 0x1e, 0x41, 0xc7, 0x0e, 0xcc, 0xf0, 0xf6, 0x03, 0xbf, 0x41, 0x93, 0xdd, 0x8e, 0xa8,
	0xf6, 0x86, 0x31, 0xb4, 0x67, 0x46, 0xea, 0x46, 0x65, 0x9f, 0x2a, 0x1f, 0xc0, 0xdd, 0xca, 0x48,
	0x3d, 0x46, 0x5e, 0x6c, 0xa3, 0x2b, 0x1d, 0x7d, 0x85, 0x30, 0xe6, 0x47, 0xa5, 0x13, 0x57, 0x76,
	0xa5, 0x23, 0x5e, 0xad, 0x32, 0xe9, 0x96, 0x9d, 0xe4, 0xf0, 0x33, 0xf0, 0xcf, 0xc4, 0x62, 0xab,
	0x04, 0x83, 0x76, 0x2e, 0xf4, 0x65, 0xbd, 0x91, 0x28, 0x87, 0xbf, 0x7b, 0xb0, 0xf3, 0x2a, 0x17,
	0x8b, 0x5b, 0xf7, 0xf7, 0x66, 0x34, 0xda, 0x0a, 0x51, 0x5e, 0xb8, 0x26, 0x48, 0x66, 0xc7, 0xb0,
	0x9f, 0x09, 0x53, 0xbe, 0xad, 0x8a, 0x44, 0x94, 0xf2, 0x6d, 0x7c, 0xed, 0x5a, 0xd9, 0x45, 0xeb,
	0x8c, 0x8c, 0x2f, 0xaf, 0xd9, 0x11, 0xf4, 0x1a, 0x51, 0xb4, 0xca, 0xdd, 0x08, 0xd6, 0x21, 0xe1,
	0x0f, 0xe0, 0x8f, 0x8a, 0x75, 0x05, 0xaf, 0x51, 0x81, 0x43, 0xa7, 0xb8, 0x5c, 0x34, 0x26, 0x53,
	0xab, 0x78, 0x9c, 0xf3, 0x4c, 0x18, 0x43, 0x3e, 0x77, 0x9c, 0x2b, 0x43, 0xf8, 0xb7, 0x07, 0xc1,
	0xb4, 0x8a, 0x5f, 0xab, 0x78, 0x0b, 0xdc, 0x31, 0x04, 0x29, 0xa2, 0x36, 0xbc, 0x45, 0x6c, 0xed,
	0x3a, 0xb6, 0x68, 0x14, 0x91, 0xf3, 0xe1, 0x36, 0x2c, 0x55, 0x62, 0x33, 0xaf, 0xb7, 0x01, 0x17,
	0x27, 0x22, 0x07, 0xde, 0xa6, 0x3b, 0x9a, 0xf6, 0xc6, 0x6d, 0x3a, 0xd2, 0xeb, 0x1b, 0x7a, 0x08,
	0x81, 0x29, 0x45, 0x59, 0x19, 0xc2, 0xbd, 0xbf, 0x0a, 0x9b, 0x92, 0x31, 0x72, 0x4e, 0xf6, 0x09,
	0x80, 0x1d, 0xcf, 0x59, 0x9a, 0x4b, 0x77, 0xc7, 0x0d, 0x0b, 0xae, 0x71, 0xa6, 0x16, 0xee, 0x94,
	0x51, 0x0c, 0xff, 0xf1, 0xc0, 0xbf, 0x0d, 0x5e, 0x7d, 0xe0, 0xad, 0xc6, 0x81, 0x1f, 0xd2, 0x80,
	0x1d, 0x96, 0xfa, 0x2f, 0x33, 0x2a, 0x2e, 0x23, 0x9a, 0xfb, 0x23, 0xe8, 0x18, 0x1a, 0xd5, 0xcd,
	0xfd, 0xb5, 0x03, 0x8c, 0x6a, 0x2f, 0xfe, 0x0f, 0xd4, 0x2f, 0x4b, 0xa9, 0x1d, 0x85, 0x56, 0x69,
	0x20, 0x0c, 0xde, 0x87, 0xf0, 0xf9, 0x06, 0xc2, 0x0e, 0xb5, 0x72, 0x30, 0xb0, 0xaf, 0xc4, 0xa0,
	0x7e, 0x25, 0x06, 0x67, 0xf5, 0x2b, 0xd1, 0x44, 0x1f, 0xbe, 0x81, 0xfb, 0x67, 0x3a, 0x5d, 0x2c,
	0xa4, 0xc6, 0x7e, 0xe4, 0xcf, 0x95, 0xa4, 0xd7, 0xc0, 0x7f, 0xa7, 0x62, 0x42, 0xbe, 0x06, 0x85,
	0x7e, 0x34, 0x53, 0x57, 0xd4, 0xb6, 0xbb, 0xe7, 0x1b, 0x98, 0x9c, 0x33, 0xfc, 0x09, 0x58, 0x33,
	0xb3, 0x29, 0xd4, 0xd2, 0xc8, 0xff, 0x27, 0xf5, 0x17, 0x70, 0x7f, 0x2a, 0x85, 0x9e, 0x5f, 0xd0,
	0xd2, 0xb8, 0xa6, 0xd7, 0xbf, 0x30, 0x6f, 0xe3, 0x17, 0xf6, 0x06, 0xf6, 0x31, 0xcc, 0x2e, 0xcf,
	0x6c, 0x99, 0x96, 0xab, 0x05, 0xf4, 0xfe, 0x7b, 0x01, 0x5b, 0xef, 0x59, 0xc0, 0x70, 0x06, 0xac,
	0xd9, 0x86, 0x43, 0xf8, 0x02, 0xee, 0x2d, 0x37, 0xea, 0x19, 0xee, 0x11, 0xf7, 0x1f, 0x36, 0x0a,
	0xad, 0xbd, 0xd1, 0xcd, 0xe8, 0xcf, 0x9f, 0x43, 0x60, 0x09, 0x66, 0x3d, 0xe8, 0x4c, 0xa2, 0xd3,
	0xc9, 0x28, 0x3a, 0xed, 0xdf, 0x41, 0x25, 0x9a, 0x8d, 0xc7, 0xaf, 0xc6, 0xdf, 0xf4, 0x3d, 0x54,
	0xa6, 0xb3, 0x93, 0x93, 0xd3, 0xe9, 0xb4, 0xdf, 0x62, 0x5d, 0xd8, 0x39, 0x8d, 0xa2, 0xef, 0xa3,
	0xbe, 0x3f, 0x1c, 0x03, 0xbc, 0x56, 0xf1, 0x54, 0x6a, 0xba, 0x90, 0xaf, 0xa0, 0x53, 0x5a, 0x0a,
	0x18, 0x77, 0xc5, 0xb7, 0xc8, 0x3e, 0xf8, 0xf8, 0x16, 0x8f, 0x85, 0x12, 0xde, 0x19, 0x8e, 0xa1,
	0x87, 0xed, 0xd6, 0x09, 0x5f, 0xe0, 0x8c, 0x11, 0xf1, 0x2a, 0xdf, 0x16, 0x0f, 0xab, 0x7c, 0xdb,
	0xa3, 0x09, 0xef, 0xbc, 0xec, 0xfe, 0xe6, 0x79, 0x7f, 0x7a, 0xde, 0x5f, 0x9e, 0x17, 0x07, 0x14,
	0xf6, 0xe4, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x8c, 0x83, 0x2a, 0xe2, 0x08, 0x00, 0x00,
}