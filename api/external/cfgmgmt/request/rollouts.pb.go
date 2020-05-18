// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cfgmgmt/request/rollouts.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SCMType int32

const (
	SCMType_UNKNOWN_SCM SCMType = 0
	SCMType_GIT         SCMType = 1
)

var SCMType_name = map[int32]string{
	0: "UNKNOWN_SCM",
	1: "GIT",
}

var SCMType_value = map[string]int32{
	"UNKNOWN_SCM": 0,
	"GIT":         1,
}

func (x SCMType) String() string {
	return proto.EnumName(SCMType_name, int32(x))
}

func (SCMType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{0}
}

type SCMWebType int32

const (
	SCMWebType_UNKNOWN_SCM_WEB SCMWebType = 0
	SCMWebType_GITHUB          SCMWebType = 1
)

var SCMWebType_name = map[int32]string{
	0: "UNKNOWN_SCM_WEB",
	1: "GITHUB",
}

var SCMWebType_value = map[string]int32{
	"UNKNOWN_SCM_WEB": 0,
	"GITHUB":          1,
}

func (x SCMWebType) String() string {
	return proto.EnumName(SCMWebType_name, int32(x))
}

func (SCMWebType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{1}
}

// CreateRollout is a request to create a new Rollout. All
// fields have the same meaning as with the response Rollout
// type.
type CreateRollout struct {
	PolicyName           string     `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	PolicyNodeGroup      string     `protobuf:"bytes,2,opt,name=policy_node_group,json=policyNodeGroup,proto3" json:"policy_node_group,omitempty"`
	PolicyRevisionId     string     `protobuf:"bytes,3,opt,name=policy_revision_id,json=policyRevisionId,proto3" json:"policy_revision_id,omitempty"`
	PolicyDomainUrl      string     `protobuf:"bytes,4,opt,name=policy_domain_url,json=policyDomainUrl,proto3" json:"policy_domain_url,omitempty"`
	ScmType              SCMType    `protobuf:"varint,5,opt,name=scm_type,json=scmType,proto3,enum=chef.automate.api.cfgmgmt.request.SCMType" json:"scm_type,omitempty"`
	ScmWebType           SCMWebType `protobuf:"varint,6,opt,name=scm_web_type,json=scmWebType,proto3,enum=chef.automate.api.cfgmgmt.request.SCMWebType" json:"scm_web_type,omitempty"`
	PolicyScmUrl         string     `protobuf:"bytes,7,opt,name=policy_scm_url,json=policyScmUrl,proto3" json:"policy_scm_url,omitempty"`
	PolicyScmWebUrl      string     `protobuf:"bytes,8,opt,name=policy_scm_web_url,json=policyScmWebUrl,proto3" json:"policy_scm_web_url,omitempty"`
	PolicyScmCommit      string     `protobuf:"bytes,9,opt,name=policy_scm_commit,json=policyScmCommit,proto3" json:"policy_scm_commit,omitempty"`
	Description          string     `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	CiJobUrl             string     `protobuf:"bytes,11,opt,name=ci_job_url,json=ciJobUrl,proto3" json:"ci_job_url,omitempty"`
	CiJobId              string     `protobuf:"bytes,12,opt,name=ci_job_id,json=ciJobId,proto3" json:"ci_job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateRollout) Reset()         { *m = CreateRollout{} }
func (m *CreateRollout) String() string { return proto.CompactTextString(m) }
func (*CreateRollout) ProtoMessage()    {}
func (*CreateRollout) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{0}
}

func (m *CreateRollout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRollout.Unmarshal(m, b)
}
func (m *CreateRollout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRollout.Marshal(b, m, deterministic)
}
func (m *CreateRollout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRollout.Merge(m, src)
}
func (m *CreateRollout) XXX_Size() int {
	return xxx_messageInfo_CreateRollout.Size(m)
}
func (m *CreateRollout) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRollout.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRollout proto.InternalMessageInfo

func (m *CreateRollout) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *CreateRollout) GetPolicyNodeGroup() string {
	if m != nil {
		return m.PolicyNodeGroup
	}
	return ""
}

func (m *CreateRollout) GetPolicyRevisionId() string {
	if m != nil {
		return m.PolicyRevisionId
	}
	return ""
}

func (m *CreateRollout) GetPolicyDomainUrl() string {
	if m != nil {
		return m.PolicyDomainUrl
	}
	return ""
}

func (m *CreateRollout) GetScmType() SCMType {
	if m != nil {
		return m.ScmType
	}
	return SCMType_UNKNOWN_SCM
}

func (m *CreateRollout) GetScmWebType() SCMWebType {
	if m != nil {
		return m.ScmWebType
	}
	return SCMWebType_UNKNOWN_SCM_WEB
}

func (m *CreateRollout) GetPolicyScmUrl() string {
	if m != nil {
		return m.PolicyScmUrl
	}
	return ""
}

func (m *CreateRollout) GetPolicyScmWebUrl() string {
	if m != nil {
		return m.PolicyScmWebUrl
	}
	return ""
}

func (m *CreateRollout) GetPolicyScmCommit() string {
	if m != nil {
		return m.PolicyScmCommit
	}
	return ""
}

func (m *CreateRollout) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateRollout) GetCiJobUrl() string {
	if m != nil {
		return m.CiJobUrl
	}
	return ""
}

func (m *CreateRollout) GetCiJobId() string {
	if m != nil {
		return m.CiJobId
	}
	return ""
}

type Rollouts struct {
	// Filters to apply to the request for the rollouts list.
	Filter               []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rollouts) Reset()         { *m = Rollouts{} }
func (m *Rollouts) String() string { return proto.CompactTextString(m) }
func (*Rollouts) ProtoMessage()    {}
func (*Rollouts) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{1}
}

func (m *Rollouts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rollouts.Unmarshal(m, b)
}
func (m *Rollouts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rollouts.Marshal(b, m, deterministic)
}
func (m *Rollouts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rollouts.Merge(m, src)
}
func (m *Rollouts) XXX_Size() int {
	return xxx_messageInfo_Rollouts.Size(m)
}
func (m *Rollouts) XXX_DiscardUnknown() {
	xxx_messageInfo_Rollouts.DiscardUnknown(m)
}

var xxx_messageInfo_Rollouts proto.InternalMessageInfo

func (m *Rollouts) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

type RolloutById struct {
	RolloutId            string   `protobuf:"bytes,1,opt,name=rollout_id,json=rolloutId,proto3" json:"rollout_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RolloutById) Reset()         { *m = RolloutById{} }
func (m *RolloutById) String() string { return proto.CompactTextString(m) }
func (*RolloutById) ProtoMessage()    {}
func (*RolloutById) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{2}
}

func (m *RolloutById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolloutById.Unmarshal(m, b)
}
func (m *RolloutById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolloutById.Marshal(b, m, deterministic)
}
func (m *RolloutById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolloutById.Merge(m, src)
}
func (m *RolloutById) XXX_Size() int {
	return xxx_messageInfo_RolloutById.Size(m)
}
func (m *RolloutById) XXX_DiscardUnknown() {
	xxx_messageInfo_RolloutById.DiscardUnknown(m)
}

var xxx_messageInfo_RolloutById proto.InternalMessageInfo

func (m *RolloutById) GetRolloutId() string {
	if m != nil {
		return m.RolloutId
	}
	return ""
}

type RolloutForChefRun struct {
	PolicyName           string   `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	PolicyGroup          string   `protobuf:"bytes,2,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	PolicyRevisionId     string   `protobuf:"bytes,3,opt,name=policy_revision_id,json=policyRevisionId,proto3" json:"policy_revision_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RolloutForChefRun) Reset()         { *m = RolloutForChefRun{} }
func (m *RolloutForChefRun) String() string { return proto.CompactTextString(m) }
func (*RolloutForChefRun) ProtoMessage()    {}
func (*RolloutForChefRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cacd55819479c9d, []int{3}
}

func (m *RolloutForChefRun) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolloutForChefRun.Unmarshal(m, b)
}
func (m *RolloutForChefRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolloutForChefRun.Marshal(b, m, deterministic)
}
func (m *RolloutForChefRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolloutForChefRun.Merge(m, src)
}
func (m *RolloutForChefRun) XXX_Size() int {
	return xxx_messageInfo_RolloutForChefRun.Size(m)
}
func (m *RolloutForChefRun) XXX_DiscardUnknown() {
	xxx_messageInfo_RolloutForChefRun.DiscardUnknown(m)
}

var xxx_messageInfo_RolloutForChefRun proto.InternalMessageInfo

func (m *RolloutForChefRun) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *RolloutForChefRun) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *RolloutForChefRun) GetPolicyRevisionId() string {
	if m != nil {
		return m.PolicyRevisionId
	}
	return ""
}

func init() {
	proto.RegisterEnum("chef.automate.api.cfgmgmt.request.SCMType", SCMType_name, SCMType_value)
	proto.RegisterEnum("chef.automate.api.cfgmgmt.request.SCMWebType", SCMWebType_name, SCMWebType_value)
	proto.RegisterType((*CreateRollout)(nil), "chef.automate.api.cfgmgmt.request.CreateRollout")
	proto.RegisterType((*Rollouts)(nil), "chef.automate.api.cfgmgmt.request.Rollouts")
	proto.RegisterType((*RolloutById)(nil), "chef.automate.api.cfgmgmt.request.RolloutById")
	proto.RegisterType((*RolloutForChefRun)(nil), "chef.automate.api.cfgmgmt.request.RolloutForChefRun")
}

func init() {
	proto.RegisterFile("api/external/cfgmgmt/request/rollouts.proto", fileDescriptor_8cacd55819479c9d)
}

var fileDescriptor_8cacd55819479c9d = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0xc6, 0x97, 0xb7, 0x7b, 0xfb, 0xe7, 0xb4, 0xac, 0x9d, 0x91, 0x50, 0x84, 0x40, 0x74, 0x85,
	0x8b, 0xaa, 0xdb, 0x12, 0x09, 0x84, 0xb8, 0x6f, 0x19, 0x25, 0xa0, 0x75, 0x52, 0xda, 0xaa, 0x12,
	0x37, 0x51, 0x6a, 0xbb, 0xad, 0x51, 0x1c, 0x07, 0xc7, 0x01, 0xfa, 0x01, 0xf8, 0x30, 0x7c, 0x4b,
	0x64, 0xc7, 0x65, 0x81, 0x0b, 0x86, 0xb8, 0x73, 0x9f, 0xf3, 0x3b, 0x8f, 0x7b, 0x4e, 0x1e, 0xc3,
	0x79, 0x9c, 0x31, 0x9f, 0x7e, 0x55, 0x54, 0xa6, 0x71, 0xe2, 0xe3, 0xcd, 0x96, 0x6f, 0xb9, 0xf2,
	0x25, 0xfd, 0x54, 0xd0, 0x5c, 0xf9, 0x52, 0x24, 0x89, 0x28, 0x54, 0xee, 0x65, 0x52, 0x28, 0x81,
	0xce, 0xf0, 0x8e, 0x6e, 0xbc, 0xb8, 0x50, 0x82, 0xc7, 0x8a, 0x7a, 0x71, 0xc6, 0x3c, 0xdb, 0xe1,
	0xd9, 0x8e, 0xc1, 0xf7, 0x63, 0xb8, 0x37, 0x91, 0x34, 0x56, 0x34, 0x2c, 0x7b, 0xd1, 0x13, 0x68,
	0x67, 0x22, 0x61, 0x78, 0x1f, 0xa5, 0x31, 0xa7, 0xae, 0xd3, 0x77, 0x86, 0xad, 0x10, 0x4a, 0x69,
	0x16, 0x73, 0x8a, 0x46, 0x70, 0x7a, 0x00, 0x04, 0xa1, 0xd1, 0x56, 0x8a, 0x22, 0x73, 0xff, 0x33,
	0x58, 0xd7, 0x62, 0x82, 0xd0, 0xa9, 0x96, 0xd1, 0x05, 0x20, 0xcb, 0x4a, 0xfa, 0x99, 0xe5, 0x4c,
	0xa4, 0x11, 0x23, 0x6e, 0xcd, 0xc0, 0xbd, 0xb2, 0x12, 0xda, 0x42, 0x40, 0x2a, 0xce, 0x44, 0xf0,
	0x98, 0xa5, 0x51, 0x21, 0x13, 0xf7, 0xb8, 0xea, 0xfc, 0xda, 0xe8, 0x4b, 0x99, 0xa0, 0x2b, 0x68,
	0xe6, 0x98, 0x47, 0x6a, 0x9f, 0x51, 0xf7, 0xff, 0xbe, 0x33, 0x3c, 0x79, 0x3e, 0xf2, 0xee, 0x1c,
	0xd7, 0x9b, 0x4f, 0xae, 0x17, 0xfb, 0x8c, 0x86, 0x8d, 0x1c, 0x73, 0x7d, 0x40, 0x37, 0xd0, 0xd1,
	0x36, 0x5f, 0xe8, 0xba, 0xb4, 0xaa, 0x1b, 0xab, 0xcb, 0xbf, 0xb3, 0x5a, 0xd1, 0xb5, 0x71, 0x83,
	0x1c, 0x73, 0x7b, 0x46, 0xcf, 0xe0, 0xc4, 0xce, 0xa0, 0x7d, 0xf5, 0x00, 0x0d, 0x33, 0x40, 0xa7,
	0x54, 0xe7, 0x98, 0xeb, 0x7f, 0x7f, 0xfe, 0x73, 0x2f, 0x87, 0xdb, 0x35, 0xd9, 0xac, 0x8e, 0x3a,
	0x37, 0x9e, 0x1a, 0xbe, 0x5d, 0x8b, 0x86, 0xb1, 0xe0, 0x9c, 0x29, 0xb7, 0xf5, 0x1b, 0x3b, 0x31,
	0x32, 0xea, 0x43, 0x9b, 0xd0, 0x1c, 0x4b, 0x96, 0x29, 0x26, 0x52, 0x17, 0x0c, 0x55, 0x95, 0xd0,
	0x23, 0x00, 0xcc, 0xa2, 0x8f, 0xa2, 0xbc, 0xb2, 0x6d, 0x80, 0x26, 0x66, 0xef, 0x84, 0xb9, 0xeb,
	0x21, 0xb4, 0x6c, 0x95, 0x11, 0xb7, 0x63, 0x8a, 0x0d, 0x53, 0x0c, 0xc8, 0x60, 0x00, 0x4d, 0x1b,
	0x92, 0x1c, 0x3d, 0x80, 0xfa, 0x86, 0x25, 0x8a, 0x4a, 0xd7, 0xe9, 0xd7, 0x86, 0xad, 0xd0, 0xfe,
	0x1a, 0x5c, 0x40, 0xdb, 0x32, 0xe3, 0x7d, 0x40, 0xd0, 0x63, 0x00, 0x9b, 0x49, 0xed, 0x57, 0x66,
	0xa9, 0x65, 0x95, 0x80, 0x0c, 0xbe, 0x39, 0x70, 0x6a, 0xf1, 0x37, 0x42, 0x4e, 0x76, 0x74, 0x13,
	0x16, 0xe9, 0xdd, 0x09, 0x3c, 0x03, 0xbb, 0xcd, 0x5f, 0xc2, 0x67, 0x9b, 0xfe, 0x21, 0x78, 0xa3,
	0xa7, 0xd0, 0xb0, 0xc9, 0x40, 0x5d, 0x68, 0x2f, 0x67, 0xef, 0x67, 0x37, 0xab, 0x59, 0x34, 0x9f,
	0x5c, 0xf7, 0x8e, 0x50, 0x03, 0x6a, 0xd3, 0x60, 0xd1, 0x73, 0x46, 0x97, 0x00, 0xb7, 0xdf, 0x1c,
	0xdd, 0x87, 0x6e, 0x85, 0x8b, 0x56, 0x57, 0xe3, 0xde, 0x11, 0x02, 0xa8, 0x4f, 0x83, 0xc5, 0xdb,
	0xe5, 0xb8, 0xe7, 0x8c, 0x5f, 0x7d, 0x78, 0xb9, 0x65, 0x6a, 0x57, 0xac, 0x3d, 0x2c, 0xb8, 0xaf,
	0xf3, 0xe4, 0x1f, 0xf2, 0xe4, 0xff, 0xe9, 0x11, 0xaf, 0xeb, 0xe6, 0xf1, 0xbe, 0xf8, 0x11, 0x00,
	0x00, 0xff, 0xff, 0x2a, 0xee, 0xf0, 0xae, 0xeb, 0x03, 0x00, 0x00,
}
