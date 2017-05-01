// Code generated by protoc-gen-go.
// source: ProfileDTO.proto
// DO NOT EDIT!

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This file lists all the objects related to Service Entity profiles
// created by user in environment or in VMTurbo
type EntityProfileDTO struct {
	// id of the entity profile.  This should allow access to the object
	// in the system it was discovered from and it should be unique in VMT also.
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// display name to be displayed to the user
	DisplayName *string `protobuf:"bytes,2,req,name=displayName" json:"displayName,omitempty"`
	// Type of entity this profile applies to
	EntityType *EntityDTO_EntityType `protobuf:"varint,3,req,name=entityType,enum=common_dto.EntityDTO_EntityType" json:"entityType,omitempty"`
	// The profile should contain multiple related commodity profiles for example
	// profile for MEM, CPU, VSTORAGE...
	CommodityProfile []*CommodityProfileDTO `protobuf:"bytes,4,rep,name=commodityProfile" json:"commodityProfile,omitempty"`
	// Model related to the profile
	Model *string `protobuf:"bytes,5,opt,name=model" json:"model,omitempty"`
	// Vendor related to the profile
	Vendor *string `protobuf:"bytes,6,opt,name=vendor" json:"vendor,omitempty"`
	// Description of the profile
	Description *string `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	// If this is a profile for VMs, vmProfileDTO must be specified
	// If this is a profile for PMs, pmProfileDTO must be specified
	//
	// Types that are valid to be assigned to VMOrPMProfileData:
	//	*EntityProfileDTO_VmProfileDTO
	//	*EntityProfileDTO_PmProfileDTO
	VMOrPMProfileData isEntityProfileDTO_VMOrPMProfileData `protobuf_oneof:"VMOrPMProfileData"`
	// This flag indicates where existing entities can be matched against this profile
	EnableProvisionMatch *bool `protobuf:"varint,10,opt,name=enableProvisionMatch" json:"enableProvisionMatch,omitempty"`
	// This flag indicates whether a resize action may use this profile to resize to
	EnableResizeMatch *bool `protobuf:"varint,11,opt,name=enableResizeMatch" json:"enableResizeMatch,omitempty"`
	// Allow entity properties to be specified related to the entity profile dto.
	// Entity properties are a list of <string, string, string> namespace, key, value triplets
	EntityProperties []*EntityDTO_EntityProperty `protobuf:"bytes,12,rep,name=entityProperties" json:"entityProperties,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *EntityProfileDTO) Reset()                    { *m = EntityProfileDTO{} }
func (m *EntityProfileDTO) String() string            { return proto.CompactTextString(m) }
func (*EntityProfileDTO) ProtoMessage()               {}
func (*EntityProfileDTO) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type isEntityProfileDTO_VMOrPMProfileData interface {
	isEntityProfileDTO_VMOrPMProfileData()
}

type EntityProfileDTO_VmProfileDTO struct {
	VmProfileDTO *EntityProfileDTO_VMProfileDTO `protobuf:"bytes,8,opt,name=vmProfileDTO,oneof"`
}
type EntityProfileDTO_PmProfileDTO struct {
	PmProfileDTO *EntityProfileDTO_PMProfileDTO `protobuf:"bytes,9,opt,name=pmProfileDTO,oneof"`
}

func (*EntityProfileDTO_VmProfileDTO) isEntityProfileDTO_VMOrPMProfileData() {}
func (*EntityProfileDTO_PmProfileDTO) isEntityProfileDTO_VMOrPMProfileData() {}

func (m *EntityProfileDTO) GetVMOrPMProfileData() isEntityProfileDTO_VMOrPMProfileData {
	if m != nil {
		return m.VMOrPMProfileData
	}
	return nil
}

func (m *EntityProfileDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *EntityProfileDTO) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *EntityProfileDTO) GetEntityType() EntityDTO_EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityDTO_SWITCH
}

func (m *EntityProfileDTO) GetCommodityProfile() []*CommodityProfileDTO {
	if m != nil {
		return m.CommodityProfile
	}
	return nil
}

func (m *EntityProfileDTO) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *EntityProfileDTO) GetVendor() string {
	if m != nil && m.Vendor != nil {
		return *m.Vendor
	}
	return ""
}

func (m *EntityProfileDTO) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *EntityProfileDTO) GetVmProfileDTO() *EntityProfileDTO_VMProfileDTO {
	if x, ok := m.GetVMOrPMProfileData().(*EntityProfileDTO_VmProfileDTO); ok {
		return x.VmProfileDTO
	}
	return nil
}

func (m *EntityProfileDTO) GetPmProfileDTO() *EntityProfileDTO_PMProfileDTO {
	if x, ok := m.GetVMOrPMProfileData().(*EntityProfileDTO_PmProfileDTO); ok {
		return x.PmProfileDTO
	}
	return nil
}

func (m *EntityProfileDTO) GetEnableProvisionMatch() bool {
	if m != nil && m.EnableProvisionMatch != nil {
		return *m.EnableProvisionMatch
	}
	return false
}

func (m *EntityProfileDTO) GetEnableResizeMatch() bool {
	if m != nil && m.EnableResizeMatch != nil {
		return *m.EnableResizeMatch
	}
	return false
}

func (m *EntityProfileDTO) GetEntityProperties() []*EntityDTO_EntityProperty {
	if m != nil {
		return m.EntityProperties
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EntityProfileDTO) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EntityProfileDTO_OneofMarshaler, _EntityProfileDTO_OneofUnmarshaler, _EntityProfileDTO_OneofSizer, []interface{}{
		(*EntityProfileDTO_VmProfileDTO)(nil),
		(*EntityProfileDTO_PmProfileDTO)(nil),
	}
}

func _EntityProfileDTO_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EntityProfileDTO)
	// VMOrPMProfileData
	switch x := m.VMOrPMProfileData.(type) {
	case *EntityProfileDTO_VmProfileDTO:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VmProfileDTO); err != nil {
			return err
		}
	case *EntityProfileDTO_PmProfileDTO:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PmProfileDTO); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EntityProfileDTO.VMOrPMProfileData has unexpected type %T", x)
	}
	return nil
}

func _EntityProfileDTO_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EntityProfileDTO)
	switch tag {
	case 8: // VMOrPMProfileData.vmProfileDTO
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EntityProfileDTO_VMProfileDTO)
		err := b.DecodeMessage(msg)
		m.VMOrPMProfileData = &EntityProfileDTO_VmProfileDTO{msg}
		return true, err
	case 9: // VMOrPMProfileData.pmProfileDTO
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EntityProfileDTO_PMProfileDTO)
		err := b.DecodeMessage(msg)
		m.VMOrPMProfileData = &EntityProfileDTO_PmProfileDTO{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EntityProfileDTO_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EntityProfileDTO)
	// VMOrPMProfileData
	switch x := m.VMOrPMProfileData.(type) {
	case *EntityProfileDTO_VmProfileDTO:
		s := proto.Size(x.VmProfileDTO)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EntityProfileDTO_PmProfileDTO:
		s := proto.Size(x.PmProfileDTO)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Specific data related to a vm profile
type EntityProfileDTO_VMProfileDTO struct {
	// At least one of numVCPUs and vCPUSpeed should be specified.
	// One of the two can be derived from the other using the capacity
	// from the commodityDTO
	// number of VCPUs
	NumVCPUs *int32 `protobuf:"varint,1,opt,name=numVCPUs" json:"numVCPUs,omitempty"`
	// VCPU speed
	VCPUSpeed *float32 `protobuf:"fixed32,2,opt,name=vCPUSpeed" json:"vCPUSpeed,omitempty"`
	// Number of storage entities that this VM will use storage from
	NumStorageConsumed *int32 `protobuf:"varint,3,opt,name=numStorageConsumed" json:"numStorageConsumed,omitempty"`
	XXX_unrecognized   []byte `json:"-"`
}

func (m *EntityProfileDTO_VMProfileDTO) Reset()         { *m = EntityProfileDTO_VMProfileDTO{} }
func (m *EntityProfileDTO_VMProfileDTO) String() string { return proto.CompactTextString(m) }
func (*EntityProfileDTO_VMProfileDTO) ProtoMessage()    {}
func (*EntityProfileDTO_VMProfileDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0, 0}
}

func (m *EntityProfileDTO_VMProfileDTO) GetNumVCPUs() int32 {
	if m != nil && m.NumVCPUs != nil {
		return *m.NumVCPUs
	}
	return 0
}

func (m *EntityProfileDTO_VMProfileDTO) GetVCPUSpeed() float32 {
	if m != nil && m.VCPUSpeed != nil {
		return *m.VCPUSpeed
	}
	return 0
}

func (m *EntityProfileDTO_VMProfileDTO) GetNumStorageConsumed() int32 {
	if m != nil && m.NumStorageConsumed != nil {
		return *m.NumStorageConsumed
	}
	return 0
}

// Specific data related to a pm profile
type EntityProfileDTO_PMProfileDTO struct {
	// At least one of numCores and cpuCoreSpeed should be specified
	// The other can be derived from the cpu capacity in
	// the commodity dto.
	NumCores         *int32   `protobuf:"varint,1,opt,name=numCores" json:"numCores,omitempty"`
	CpuCoreSpeed     *float32 `protobuf:"fixed32,2,opt,name=cpuCoreSpeed" json:"cpuCoreSpeed,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *EntityProfileDTO_PMProfileDTO) Reset()         { *m = EntityProfileDTO_PMProfileDTO{} }
func (m *EntityProfileDTO_PMProfileDTO) String() string { return proto.CompactTextString(m) }
func (*EntityProfileDTO_PMProfileDTO) ProtoMessage()    {}
func (*EntityProfileDTO_PMProfileDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor5, []int{0, 1}
}

func (m *EntityProfileDTO_PMProfileDTO) GetNumCores() int32 {
	if m != nil && m.NumCores != nil {
		return *m.NumCores
	}
	return 0
}

func (m *EntityProfileDTO_PMProfileDTO) GetCpuCoreSpeed() float32 {
	if m != nil && m.CpuCoreSpeed != nil {
		return *m.CpuCoreSpeed
	}
	return 0
}

// Data related to a commodity profile used in an entity profile
// Note typically only a subset of these fields may be specified in a profile for
// each commmodity.
type CommodityProfileDTO struct {
	// The type of commodity such as MEM, CPU, STORAGE
	CommodityType *CommodityDTO_CommodityType `protobuf:"varint,1,req,name=commodityType,enum=common_dto.CommodityDTO_CommodityType" json:"commodityType,omitempty"`
	// The capacity of the commodity
	Capacity *float32 `protobuf:"fixed32,2,opt,name=capacity" json:"capacity,omitempty"`
	// Consumed factor may be used to calculate consumed based on capacity
	ConsumedFactor *float32 `protobuf:"fixed32,3,opt,name=consumedFactor" json:"consumedFactor,omitempty"`
	// Consumed value to be used in the profile
	Consumed *float32 `protobuf:"fixed32,4,opt,name=consumed" json:"consumed,omitempty"`
	// A reservation related to this commodity
	Reservation *float32 `protobuf:"fixed32,5,opt,name=reservation" json:"reservation,omitempty"`
	// Overhead related to this commodity - for example overheadMem
	Overhead         *float32 `protobuf:"fixed32,6,opt,name=overhead" json:"overhead,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CommodityProfileDTO) Reset()                    { *m = CommodityProfileDTO{} }
func (m *CommodityProfileDTO) String() string            { return proto.CompactTextString(m) }
func (*CommodityProfileDTO) ProtoMessage()               {}
func (*CommodityProfileDTO) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *CommodityProfileDTO) GetCommodityType() CommodityDTO_CommodityType {
	if m != nil && m.CommodityType != nil {
		return *m.CommodityType
	}
	return CommodityDTO_CLUSTER
}

func (m *CommodityProfileDTO) GetCapacity() float32 {
	if m != nil && m.Capacity != nil {
		return *m.Capacity
	}
	return 0
}

func (m *CommodityProfileDTO) GetConsumedFactor() float32 {
	if m != nil && m.ConsumedFactor != nil {
		return *m.ConsumedFactor
	}
	return 0
}

func (m *CommodityProfileDTO) GetConsumed() float32 {
	if m != nil && m.Consumed != nil {
		return *m.Consumed
	}
	return 0
}

func (m *CommodityProfileDTO) GetReservation() float32 {
	if m != nil && m.Reservation != nil {
		return *m.Reservation
	}
	return 0
}

func (m *CommodityProfileDTO) GetOverhead() float32 {
	if m != nil && m.Overhead != nil {
		return *m.Overhead
	}
	return 0
}

// This represents a deployment profile (service catalog item) which is related
// to a service entity profile (template)
// This DTO ties image information with scope and a profile to allow for
// the deployment of an entity related to a profile
type DeploymentProfileDTO struct {
	// id related to this.  This may be an id from the system where this was discovered
	// it must be unique in VMTurbo
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// Display name for the deployment profile
	ProfileName *string `protobuf:"bytes,2,req,name=profileName" json:"profileName,omitempty"`
	// Context data needed to actually invoke deploy - such as URIs
	ContextData []*ContextData `protobuf:"bytes,3,rep,name=contextData" json:"contextData,omitempty"`
	// related service entity profiles (templates)
	RelatedEntityProfileId []string `protobuf:"bytes,4,rep,name=relatedEntityProfileId" json:"relatedEntityProfileId,omitempty"`
	// scopes in which this can be used for example clusters
	RelatedScopeId   []string `protobuf:"bytes,5,rep,name=relatedScopeId" json:"relatedScopeId,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DeploymentProfileDTO) Reset()                    { *m = DeploymentProfileDTO{} }
func (m *DeploymentProfileDTO) String() string            { return proto.CompactTextString(m) }
func (*DeploymentProfileDTO) ProtoMessage()               {}
func (*DeploymentProfileDTO) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *DeploymentProfileDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *DeploymentProfileDTO) GetProfileName() string {
	if m != nil && m.ProfileName != nil {
		return *m.ProfileName
	}
	return ""
}

func (m *DeploymentProfileDTO) GetContextData() []*ContextData {
	if m != nil {
		return m.ContextData
	}
	return nil
}

func (m *DeploymentProfileDTO) GetRelatedEntityProfileId() []string {
	if m != nil {
		return m.RelatedEntityProfileId
	}
	return nil
}

func (m *DeploymentProfileDTO) GetRelatedScopeId() []string {
	if m != nil {
		return m.RelatedScopeId
	}
	return nil
}

func init() {
	proto.RegisterType((*EntityProfileDTO)(nil), "common_dto.EntityProfileDTO")
	proto.RegisterType((*EntityProfileDTO_VMProfileDTO)(nil), "common_dto.EntityProfileDTO.VMProfileDTO")
	proto.RegisterType((*EntityProfileDTO_PMProfileDTO)(nil), "common_dto.EntityProfileDTO.PMProfileDTO")
	proto.RegisterType((*CommodityProfileDTO)(nil), "common_dto.CommodityProfileDTO")
	proto.RegisterType((*DeploymentProfileDTO)(nil), "common_dto.DeploymentProfileDTO")
}

func init() { proto.RegisterFile("ProfileDTO.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xfd, 0x25, 0x5d, 0xf7, 0x5b, 0x6f, 0xba, 0xd1, 0xa5, 0xd3, 0x30, 0x15, 0x62, 0xd1, 0x84,
	0x50, 0x90, 0x20, 0x0f, 0x13, 0x42, 0xe2, 0x01, 0x24, 0x96, 0x0d, 0xc1, 0x43, 0x69, 0x44, 0xb7,
	0xbe, 0x22, 0x2f, 0xbe, 0x63, 0x91, 0xe2, 0xd8, 0x72, 0xdc, 0x88, 0xf0, 0x5d, 0xf8, 0x1c, 0x7c,
	0x32, 0xde, 0x91, 0xdd, 0x3f, 0x49, 0xd9, 0x40, 0xbc, 0x25, 0xf7, 0x1c, 0x9f, 0x7b, 0xcf, 0xf5,
	0x31, 0x0c, 0x12, 0x25, 0xae, 0xb3, 0x1c, 0xcf, 0x2e, 0x26, 0x91, 0x54, 0x42, 0x0b, 0x1f, 0x52,
	0xc1, 0xb9, 0x28, 0x3e, 0x33, 0x2d, 0x46, 0xf7, 0x62, 0xfb, 0xbd, 0x06, 0x8f, 0x7f, 0x6e, 0xc1,
	0xe0, 0xbc, 0xd0, 0x99, 0xae, 0x9b, 0x73, 0x3e, 0x80, 0x9b, 0x31, 0xe2, 0x04, 0x6e, 0xd8, 0xf3,
	0x87, 0xe0, 0xb1, 0xac, 0x94, 0x39, 0xad, 0x3f, 0x52, 0x8e, 0xc4, 0xb5, 0xc5, 0x17, 0x00, 0x68,
	0x0f, 0x5d, 0xd4, 0x12, 0x49, 0x27, 0x70, 0xc3, 0xbd, 0x93, 0x20, 0x6a, 0xfa, 0x44, 0x0b, 0x49,
	0xd3, 0xe6, 0x7c, 0xcd, 0xf3, 0x5f, 0xc1, 0xc0, 0x52, 0x58, 0xd3, 0x8d, 0x6c, 0x05, 0x9d, 0xd0,
	0x3b, 0x39, 0x6a, 0x9f, 0x8d, 0x7f, 0xe3, 0x98, 0x89, 0x76, 0xa1, 0xcb, 0x05, 0xc3, 0x9c, 0x74,
	0x03, 0x27, 0xec, 0xf9, 0x7b, 0xb0, 0x5d, 0x61, 0xc1, 0x84, 0x22, 0xdb, 0xf6, 0xdf, 0x0c, 0x89,
	0x65, 0xaa, 0x32, 0xa9, 0x33, 0x51, 0x90, 0xff, 0x6d, 0xf1, 0x2d, 0xf4, 0x2b, 0xde, 0x68, 0x90,
	0x9d, 0xc0, 0x09, 0xbd, 0x93, 0xa7, 0xb7, 0xc7, 0x6c, 0x6d, 0x6c, 0x36, 0x6e, 0x7e, 0xde, 0xff,
	0x67, 0x24, 0x64, 0x5b, 0xa2, 0xf7, 0x0f, 0x12, 0xc9, 0xa6, 0xc4, 0x43, 0x38, 0xc0, 0x82, 0x5e,
	0xe5, 0x98, 0x28, 0x51, 0x65, 0x65, 0x26, 0x8a, 0x31, 0xd5, 0xe9, 0x0d, 0x81, 0xc0, 0x09, 0x77,
	0xfc, 0x07, 0xb0, 0xbf, 0x40, 0x3f, 0x61, 0x99, 0x7d, 0xc3, 0x05, 0xe4, 0x59, 0xe8, 0x0d, 0x0c,
	0x70, 0xa5, 0x2d, 0x51, 0xe9, 0x0c, 0x4b, 0xd2, 0xb7, 0xdb, 0x7a, 0xfc, 0xb7, 0x4d, 0x2f, 0xd9,
	0xf5, 0x68, 0x02, 0xfd, 0xb6, 0x1b, 0x7f, 0x00, 0x3b, 0xc5, 0x9c, 0xcf, 0xe2, 0xe4, 0xb2, 0x24,
	0x4e, 0xe0, 0x84, 0x5d, 0x7f, 0x1f, 0x7a, 0x55, 0x9c, 0x5c, 0x4e, 0x25, 0x22, 0x23, 0x6e, 0xe0,
	0x84, 0xae, 0x3f, 0x02, 0xbf, 0x98, 0xf3, 0xa9, 0x16, 0x8a, 0x7e, 0xc1, 0x58, 0x14, 0xe5, 0x9c,
	0x23, 0x23, 0x1d, 0x43, 0x1f, 0xbd, 0x84, 0x7e, 0x72, 0x5b, 0x30, 0x16, 0x0a, 0x57, 0x82, 0x07,
	0xd0, 0x4f, 0xe5, 0xdc, 0x54, 0x5a, 0x9a, 0xa7, 0x43, 0xd8, 0x9f, 0x8d, 0x27, 0xaa, 0x39, 0x4b,
	0x35, 0x3d, 0xfe, 0xe1, 0xc0, 0xf0, 0xae, 0x8b, 0x7e, 0x0d, 0xbb, 0xeb, 0x8c, 0xd8, 0x70, 0x39,
	0x36, 0x5c, 0x4f, 0xee, 0x0c, 0x88, 0x71, 0x1d, 0xb7, 0xd9, 0x66, 0xa6, 0x94, 0x4a, 0x9a, 0x66,
	0xba, 0x5e, 0x3a, 0x3a, 0x84, 0xbd, 0x74, 0xe9, 0xe3, 0x1d, 0x4d, 0xb5, 0x50, 0xd6, 0x8d, 0x6b,
	0x99, 0x2b, 0x7f, 0x5b, 0xb6, 0x32, 0x04, 0x4f, 0x61, 0x89, 0xaa, 0xa2, 0x36, 0x44, 0xdd, 0x15,
	0x4d, 0x54, 0xa8, 0x6e, 0x90, 0x32, 0x9b, 0x35, 0xf7, 0xf8, 0xbb, 0x03, 0x07, 0x67, 0x28, 0x73,
	0x51, 0x73, 0x2c, 0xf4, 0x9f, 0x5f, 0x8d, 0x5c, 0x20, 0xad, 0x57, 0xf3, 0x0c, 0xbc, 0x54, 0x14,
	0x1a, 0xbf, 0x6a, 0xb3, 0x02, 0xd2, 0xb1, 0x97, 0x79, 0x7f, 0xd3, 0xd9, 0x1a, 0xf6, 0x1f, 0xc1,
	0xa1, 0xc2, 0x9c, 0x6a, 0x64, 0x1b, 0x11, 0xfb, 0xc0, 0xec, 0x9b, 0xe9, 0x19, 0x63, 0x4b, 0x7c,
	0x9a, 0x0a, 0x69, 0xea, 0x5d, 0x53, 0x3f, 0x7d, 0x0e, 0x47, 0xa9, 0xe0, 0x51, 0xc5, 0xf5, 0x5c,
	0x5d, 0x89, 0x48, 0xe6, 0x54, 0x5f, 0x0b, 0xc5, 0x97, 0x6d, 0x22, 0xa6, 0xc5, 0x29, 0x34, 0x53,
	0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xc4, 0x52, 0x8a, 0x30, 0x04, 0x00, 0x00,
}