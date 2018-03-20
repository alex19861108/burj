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
	GetNodesRequest
	GetNodesResponse
	GetDevicesRequest
	GetDevicesResponse
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
	// @inject_tag: bson:"summary,omitempty"
	Summary string `protobuf:"bytes,7,opt,name=summary" json:"summary,omitempty" bson:"summary,omitempty"`
	// @inject_tag: bson:"log,omitempty"
	Log string `protobuf:"bytes,8,opt,name=log" json:"log,omitempty" bson:"log,omitempty"`
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

func (m *SubJob) GetSummary() string {
	if m != nil {
		return m.Summary
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

type GetNodesRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetNodesRequest) Reset()                    { *m = GetNodesRequest{} }
func (m *GetNodesRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetNodesRequest) ProtoMessage()               {}
func (*GetNodesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *GetNodesRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetNodesResponse struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *GetNodesResponse) Reset()                    { *m = GetNodesResponse{} }
func (m *GetNodesResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetNodesResponse) ProtoMessage()               {}
func (*GetNodesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *GetNodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type GetDevicesRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetDevicesRequest) Reset()                    { *m = GetDevicesRequest{} }
func (m *GetDevicesRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetDevicesRequest) ProtoMessage()               {}
func (*GetDevicesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *GetDevicesRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDevicesResponse struct {
	Devices []*Device `protobuf:"bytes,1,rep,name=devices" json:"devices,omitempty"`
}

func (m *GetDevicesResponse) Reset()                    { *m = GetDevicesResponse{} }
func (m *GetDevicesResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetDevicesResponse) ProtoMessage()               {}
func (*GetDevicesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *GetDevicesResponse) GetDevices() []*Device {
	if m != nil {
		return m.Devices
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
	proto1.RegisterType((*GetNodesRequest)(nil), "proto.GetNodesRequest")
	proto1.RegisterType((*GetNodesResponse)(nil), "proto.GetNodesResponse")
	proto1.RegisterType((*GetDevicesRequest)(nil), "proto.GetDevicesRequest")
	proto1.RegisterType((*GetDevicesResponse)(nil), "proto.GetDevicesResponse")
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
	GetNodes(ctx context.Context, in *GetNodesRequest, opts ...grpc.CallOption) (*GetNodesResponse, error)
	GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (*GetDevicesResponse, error)
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

func (c *nodeServiceClient) GetNodes(ctx context.Context, in *GetNodesRequest, opts ...grpc.CallOption) (*GetNodesResponse, error) {
	out := new(GetNodesResponse)
	err := grpc.Invoke(ctx, "/proto.NodeService/getNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (*GetDevicesResponse, error) {
	out := new(GetDevicesResponse)
	err := grpc.Invoke(ctx, "/proto.NodeService/getDevices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeService service

type NodeServiceServer interface {
	Search(context.Context, *SearchNodeRequest) (*SearchNodeResponse, error)
	GetNodes(context.Context, *GetNodesRequest) (*GetNodesResponse, error)
	GetDevices(context.Context, *GetDevicesRequest) (*GetDevicesResponse, error)
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

func _NodeService_GetNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/GetNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetNodes(ctx, req.(*GetNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/GetDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetDevices(ctx, req.(*GetDevicesRequest))
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
		{
			MethodName: "getNodes",
			Handler:    _NodeService_GetNodes_Handler,
		},
		{
			MethodName: "getDevices",
			Handler:    _NodeService_GetDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "burj.proto",
}

func init() { proto1.RegisterFile("burj.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1076 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5b, 0x73, 0xdb, 0x44,
	0x14, 0x8e, 0x2c, 0x5b, 0x8e, 0x8f, 0x73, 0x71, 0x76, 0xa0, 0x2c, 0x99, 0x0c, 0x69, 0x97, 0x74,
	0xca, 0x6d, 0xd2, 0x99, 0xb4, 0xf4, 0xa1, 0x33, 0x9d, 0x92, 0x86, 0x4c, 0xa6, 0x1d, 0x30, 0x41,
	0x8e, 0x87, 0xf2, 0xd4, 0x59, 0x59, 0x1b, 0x45, 0x8d, 0x64, 0x89, 0x5d, 0x29, 0x90, 0x37, 0x9e,
	0x80, 0x47, 0x7e, 0x1b, 0xff, 0x82, 0x9f, 0xc0, 0x1b, 0x73, 0x76, 0x57, 0xb2, 0x6c, 0x87, 0xf0,
	0xd2, 0x27, 0x9f, 0x9b, 0xcf, 0xed, 0xfb, 0xce, 0x0a, 0x20, 0x28, 0xe5, 0xdb, 0xfd, 0x5c, 0x66,
	0x45, 0x46, 0x3a, 0xfa, 0x67, 0x7b, 0x37, 0xca, 0xb2, 0x28, 0x11, 0x0f, 0xb5, 0x16, 0x94, 0xe7,
	0x0f, 0x8b, 0x38, 0x15, 0xaa, 0xe0, 0x69, 0x6e, 0xe2, 0x58, 0x0a, 0xee, 0xd1, 0xe9, 0x98, 0x0c,
	0xc0, 0xe5, 0x41, 0x4c, 0x9d, 0xbb, 0xce, 0x27, 0x3d, 0x1f, 0x45, 0x42, 0xa1, 0xcb, 0x83, 0x38,
	0x89, 0x55, 0x41, 0x5b, 0xda, 0x5a, 0xa9, 0x64, 0x07, 0x7a, 0x56, 0x7c, 0x74, 0x40, 0x5d, 0xed,
	0x9b, 0x19, 0x1a, 0xde, 0x27, 0x8f, 0x69, 0x7b, 0xce, 0xfb, 0xe4, 0x31, 0x3b, 0x05, 0xef, 0x5b,
	0x91, 0x66, 0xf2, 0x9a, 0xbc, 0x07, 0x9d, 0x22, 0x2b, 0x78, 0x62, 0x6b, 0x1a, 0x85, 0x10, 0x68,
	0x97, 0x4a, 0x84, 0xb6, 0xa4, 0x96, 0x75, 0xc6, 0x2b, 0x1e, 0x27, 0x3c, 0x48, 0x44, 0x5d, 0xaf,
	0x32, 0xb0, 0xdf, 0x5d, 0xf0, 0xbe, 0x16, 0x57, 0xf1, 0x44, 0x90, 0x0d, 0x68, 0xc5, 0xa1, 0xcd,
	0xd7, 0x8a, 0x43, 0x4c, 0xc6, 0xc3, 0x50, 0x56, 0xc9, 0x50, 0xc6, 0xb2, 0x09, 0x0f, 0x44, 0x62,
	0x13, 0x19, 0x85, 0xec, 0x80, 0x3b, 0xc9, 0x4b, 0xdd, 0x6e, 0xff, 0x00, 0xcc, 0x6a, 0xf6, 0x8f,
	0x4e, 0xc7, 0x3e, 0x9a, 0xc9, 0x7d, 0xf0, 0x52, 0xdd, 0x34, 0xed, 0xe8, 0x80, 0x75, 0x1b, 0x60,
	0x26, 0xf1, 0xad, 0x13, 0x53, 0xa7, 0x59, 0x28, 0x12, 0xea, 0x99, 0xd4, 0x5a, 0x41, 0x6b, 0x20,
	0xf9, 0x34, 0xa4, 0x5d, 0x63, 0xd5, 0x0a, 0xb6, 0x36, 0xe5, 0xa9, 0xa0, 0xab, 0xa6, 0x35, 0x94,
	0xc9, 0x1d, 0xf0, 0x42, 0x3d, 0x08, 0xed, 0x69, 0xab, 0xd5, 0x74, 0x86, 0x8c, 0xcb, 0x90, 0x82,
	0xcd, 0x80, 0x0a, 0x46, 0x27, 0xd9, 0x84, 0x27, 0x82, 0xf6, 0x4d, 0xb4, 0xd1, 0xc8, 0x1e, 0xac,
	0x9f, 0xc7, 0x52, 0x15, 0x87, 0x79, 0xfc, 0x8d, 0xb8, 0x12, 0x09, 0x5d, 0xd3, 0xee, 0x79, 0x23,
	0x61, 0xb0, 0x96, 0xf2, 0x69, 0x79, 0xce, 0x27, 0x45, 0x29, 0x85, 0xa4, 0xeb, 0x3a, 0x68, 0xce,
	0x86, 0x15, 0x26, 0x65, 0x5e, 0xa8, 0x94, 0x6e, 0x98, 0x0a, 0x46, 0x43, 0xbb, 0x12, 0x32, 0xe6,
	0x09, 0xdd, 0x34, 0x76, 0xa3, 0xb1, 0x2f, 0xa0, 0x7d, 0x88, 0x2b, 0x26, 0xd0, 0xbe, 0x28, 0x8a,
	0xdc, 0x02, 0xa1, 0x65, 0xe4, 0x97, 0xcc, 0x27, 0x16, 0x09, 0x14, 0xd9, 0x2f, 0xd0, 0x1e, 0x66,
	0xe1, 0x32, 0x68, 0xbb, 0x0d, 0xd0, 0xfa, 0x07, 0x7d, 0xbb, 0x6a, 0x4c, 0x6c, 0x11, 0xc4, 0xc1,
	0x11, 0x34, 0x65, 0x21, 0xb4, 0x1a, 0x79, 0x00, 0x5d, 0xb3, 0x30, 0x45, 0xdb, 0x77, 0xdd, 0x06,
	0x4c, 0x86, 0x1d, 0x7e, 0xe5, 0x65, 0x01, 0xb4, 0xc7, 0x4a, 0xc8, 0x46, 0x65, 0x57, 0x57, 0xde,
	0x86, 0xd5, 0x52, 0x09, 0x39, 0x44, 0x5c, 0x4c, 0xa3, 0xb5, 0x8e, 0xbe, 0x9c, 0x2b, 0xf5, 0x43,
	0x26, 0x43, 0x5b, 0xb6, 0xd6, 0x71, 0x5e, 0x99, 0x25, 0xc2, 0x92, 0x5d, 0xcb, 0xec, 0x53, 0x70,
	0xcf, 0x78, 0xb4, 0x54, 0x82, 0x40, 0x3b, 0xe5, 0xf2, 0xb2, 0x62, 0x24, 0xca, 0xec, 0x37, 0x07,
	0x3a, 0x2f, 0x53, 0x1e, 0xdd, 0xc8, 0xdf, 0xc5, 0x68, 0xb4, 0xe5, 0xbc, 0xb8, 0xb0, 0x4d, 0x68,
	0x99, 0xec, 0xc1, 0x46, 0xc2, 0x55, 0xf1, 0xa6, 0xcc, 0x43, 0x5e, 0x88, 0x37, 0xc1, 0xb5, 0x6d,
	0x65, 0x0d, 0xad, 0x63, 0x6d, 0x7c, 0x71, 0x4d, 0x76, 0xa1, 0xdf, 0x88, 0xd2, 0x54, 0xee, 0xf9,
	0x30, 0x0b, 0x61, 0xdf, 0x83, 0x7b, 0x98, 0xcf, 0x2a, 0x38, 0x8d, 0x0a, 0x14, 0xba, 0xf9, 0x65,
	0xd4, 0xd8, 0x4c, 0xa5, 0xe2, 0x71, 0x4e, 0x12, 0xae, 0x94, 0xf6, 0xd9, 0xe3, 0xac, 0x0d, 0xec,
	0x1f, 0x07, 0xbc, 0x51, 0x19, 0xbc, 0xca, 0x82, 0xa5, 0xe1, 0xf6, 0xc0, 0x8b, 0x71, 0x6a, 0x45,
	0x5b, 0x1a, 0xad, 0x35, 0x8b, 0x96, 0x5e, 0x85, 0x6f, 0x7d, 0xc8, 0x86, 0x69, 0x16, 0x9a, 0xcc,
	0x33, 0x36, 0x20, 0x71, 0x7c, 0xed, 0xc0, 0xdb, 0xb4, 0x47, 0xd3, 0x9e, 0xbb, 0x4d, 0x0b, 0x7a,
	0x75, 0x43, 0xf7, 0xc1, 0x53, 0x05, 0x2f, 0x4a, 0xa5, 0xe7, 0xde, 0xa8, 0xc3, 0x46, 0xda, 0xe8,
	0x5b, 0x27, 0xf9, 0x08, 0xc0, 0xac, 0xe7, 0x2c, 0x4e, 0x85, 0xbd, 0xe3, 0x86, 0x05, 0xf7, 0xa0,
	0xca, 0x34, 0xe5, 0xf2, 0xda, 0x9e, 0x73, 0xa5, 0x22, 0xc1, 0x93, 0x2c, 0xb2, 0xf7, 0x8c, 0x22,
	0xfb, 0xdb, 0x01, 0xf7, 0xa6, 0xc1, 0xab, 0xd3, 0x6f, 0x35, 0x4e, 0x7f, 0x47, 0xaf, 0xde, 0x4e,
	0x59, 0xbd, 0x3f, 0x87, 0xf9, 0xa5, 0xaf, 0x11, 0x79, 0x80, 0x55, 0x71, 0x89, 0x8b, 0xcc, 0x36,
	0xab, 0xf5, 0x2b, 0x2f, 0xbe, 0x14, 0xd9, 0xcf, 0x53, 0x21, 0x2d, 0xb8, 0x46, 0x69, 0xcc, 0xee,
	0xdd, 0x36, 0xfb, 0xd3, 0xb9, 0xd9, 0xbb, 0xba, 0x95, 0xed, 0x7d, 0xf3, 0xfd, 0xd8, 0xaf, 0xbe,
	0x1f, 0xfb, 0x67, 0xd5, 0xf7, 0xa3, 0xb9, 0x17, 0xf6, 0x1a, 0xb6, 0xce, 0x64, 0x1c, 0x45, 0x42,
	0x62, 0x3f, 0xe2, 0xa7, 0x52, 0xe8, 0xef, 0x84, 0xfb, 0x36, 0x0b, 0xf4, 0xe4, 0xb3, 0xa1, 0xd0,
	0x8f, 0x66, 0xdd, 0x95, 0x6e, 0xdb, 0x5e, 0xfa, 0xc2, 0x4c, 0xd6, 0xc9, 0x7e, 0x04, 0xd2, 0xcc,
	0xac, 0xf2, 0x6c, 0xaa, 0xc4, 0xbb, 0x49, 0xfd, 0x39, 0x6c, 0x8d, 0x04, 0x97, 0x93, 0x0b, 0x4d,
	0x27, 0xdb, 0xf4, 0xec, 0x71, 0x73, 0xe6, 0x1e, 0xb7, 0xd7, 0xb0, 0x81, 0x61, 0x86, 0x56, 0xe3,
	0x69, 0x5c, 0xd4, 0xd4, 0x74, 0xfe, 0x9f, 0x9a, 0xad, 0x5b, 0xa8, 0xc9, 0xc6, 0x40, 0x9a, 0x6d,
	0xd8, 0x09, 0x9f, 0xc3, 0xe6, 0x74, 0xae, 0x9e, 0xa2, 0x8e, 0xc6, 0xfe, 0xfd, 0x46, 0xa1, 0x99,
	0xd7, 0x5f, 0x8c, 0x66, 0xf7, 0x60, 0xf3, 0x44, 0x14, 0x18, 0xa5, 0xaa, 0xd9, 0x16, 0x98, 0xc8,
	0xbe, 0x84, 0xc1, 0x2c, 0xc4, 0xd6, 0xbd, 0x07, 0x1d, 0xcc, 0x54, 0x55, 0x9b, 0x1b, 0xcb, 0x78,
	0xd8, 0xc7, 0xb0, 0x75, 0x22, 0x0a, 0x53, 0xeb, 0x3f, 0x73, 0x3f, 0x03, 0xd2, 0x0c, 0xb2, 0xd9,
	0x1b, 0x6f, 0xb4, 0x73, 0xdb, 0x1b, 0xfd, 0xd9, 0x53, 0xf0, 0x0c, 0x3d, 0x49, 0x1f, 0xba, 0xa7,
	0xfe, 0xf1, 0xe9, 0xa1, 0x7f, 0x3c, 0x58, 0x41, 0xc5, 0x1f, 0x0f, 0x87, 0x2f, 0x87, 0x27, 0x03,
	0x07, 0x95, 0xd1, 0xf8, 0xe8, 0xe8, 0x78, 0x34, 0x1a, 0xb4, 0x48, 0x0f, 0x3a, 0xc7, 0xbe, 0xff,
	0x9d, 0x3f, 0x70, 0x0f, 0x86, 0x00, 0xaf, 0xb2, 0x60, 0x24, 0xa4, 0xbe, 0xfc, 0xaf, 0xa0, 0x5b,
	0x18, 0x02, 0x11, 0x6a, 0x8b, 0x2d, 0x51, 0x75, 0xfb, 0xc3, 0x1b, 0x3c, 0xa6, 0x65, 0xb6, 0x72,
	0xf0, 0x97, 0x03, 0x7d, 0x9c, 0xbf, 0xca, 0xf8, 0x1c, 0x29, 0x82, 0x80, 0xd5, 0x09, 0x97, 0x68,
	0x54, 0x27, 0x5c, 0x46, 0x96, 0xad, 0x90, 0x67, 0xb0, 0x1a, 0xd9, 0xbd, 0x93, 0x3b, 0x36, 0x70,
	0x01, 0xab, 0xed, 0x0f, 0x96, 0xec, 0xf5, 0xdf, 0x8f, 0x00, 0xa2, 0x7a, 0xb5, 0x75, 0x0f, 0x4b,
	0x90, 0xd4, 0x3d, 0x2c, 0xe3, 0xc0, 0x56, 0x5e, 0xf4, 0x7e, 0x75, 0x9c, 0x3f, 0x1c, 0xe7, 0x4f,
	0xc7, 0x09, 0x3c, 0x1d, 0xf6, 0xe8, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x2b, 0x55, 0x3c,
	0x3f, 0x0a, 0x00, 0x00,
}
