// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/infra_proxy/response/policyfiles.proto

package response

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

type Policyfiles struct {
	// Policyfiles list.
	Policies             []*PolicyfileListItem `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Policyfiles) Reset()         { *m = Policyfiles{} }
func (m *Policyfiles) String() string { return proto.CompactTextString(m) }
func (*Policyfiles) ProtoMessage()    {}
func (*Policyfiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{0}
}

func (m *Policyfiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policyfiles.Unmarshal(m, b)
}
func (m *Policyfiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policyfiles.Marshal(b, m, deterministic)
}
func (m *Policyfiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policyfiles.Merge(m, src)
}
func (m *Policyfiles) XXX_Size() int {
	return xxx_messageInfo_Policyfiles.Size(m)
}
func (m *Policyfiles) XXX_DiscardUnknown() {
	xxx_messageInfo_Policyfiles.DiscardUnknown(m)
}

var xxx_messageInfo_Policyfiles proto.InternalMessageInfo

func (m *Policyfiles) GetPolicies() []*PolicyfileListItem {
	if m != nil {
		return m.Policies
	}
	return nil
}

type PolicyfileListItem struct {
	// Policyfile name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Policyfile Revision ID.
	RevisionId string `protobuf:"bytes,3,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// Policyfile policy group.
	PolicyGroup          string   `protobuf:"bytes,2,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyfileListItem) Reset()         { *m = PolicyfileListItem{} }
func (m *PolicyfileListItem) String() string { return proto.CompactTextString(m) }
func (*PolicyfileListItem) ProtoMessage()    {}
func (*PolicyfileListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{1}
}

func (m *PolicyfileListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyfileListItem.Unmarshal(m, b)
}
func (m *PolicyfileListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyfileListItem.Marshal(b, m, deterministic)
}
func (m *PolicyfileListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyfileListItem.Merge(m, src)
}
func (m *PolicyfileListItem) XXX_Size() int {
	return xxx_messageInfo_PolicyfileListItem.Size(m)
}
func (m *PolicyfileListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyfileListItem.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyfileListItem proto.InternalMessageInfo

func (m *PolicyfileListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PolicyfileListItem) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

func (m *PolicyfileListItem) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

type Policyfile struct {
	// Policyfile name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Policy group name.
	PolicyGroup string `protobuf:"bytes,2,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	// Policy revision ID.
	RevisionId string `protobuf:"bytes,3,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// Run-list associated with the policy.
	RunList []string `protobuf:"bytes,4,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	// Named run-list associated with the policy.
	NamedRunList []*NamedRunList `protobuf:"bytes,5,rep,name=named_run_list,json=namedRunList,proto3" json:"named_run_list,omitempty"`
	// Included policy locks files.
	IncludedPolicyLocks []*IncludedPolicyLock `protobuf:"bytes,6,rep,name=included_policy_locks,json=includedPolicyLocks,proto3" json:"included_policy_locks,omitempty"`
	// List of cookbook locks under this policy.
	CookbookLocks []*CookbookLock `protobuf:"bytes,7,rep,name=cookbook_locks,json=cookbookLocks,proto3" json:"cookbook_locks,omitempty"`
	// Stringified JSON of the default attributes.
	DefaultAttributes string `protobuf:"bytes,8,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Stringified JSON of the override attributes.
	OverrideAttributes string `protobuf:"bytes,9,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
	// Expanded run-list associated with the policy.
	ExpandedRunList      []*ExpandedRunList `protobuf:"bytes,10,rep,name=expanded_run_list,json=expandedRunList,proto3" json:"expanded_run_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Policyfile) Reset()         { *m = Policyfile{} }
func (m *Policyfile) String() string { return proto.CompactTextString(m) }
func (*Policyfile) ProtoMessage()    {}
func (*Policyfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{2}
}

func (m *Policyfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policyfile.Unmarshal(m, b)
}
func (m *Policyfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policyfile.Marshal(b, m, deterministic)
}
func (m *Policyfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policyfile.Merge(m, src)
}
func (m *Policyfile) XXX_Size() int {
	return xxx_messageInfo_Policyfile.Size(m)
}
func (m *Policyfile) XXX_DiscardUnknown() {
	xxx_messageInfo_Policyfile.DiscardUnknown(m)
}

var xxx_messageInfo_Policyfile proto.InternalMessageInfo

func (m *Policyfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policyfile) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *Policyfile) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

func (m *Policyfile) GetRunList() []string {
	if m != nil {
		return m.RunList
	}
	return nil
}

func (m *Policyfile) GetNamedRunList() []*NamedRunList {
	if m != nil {
		return m.NamedRunList
	}
	return nil
}

func (m *Policyfile) GetIncludedPolicyLocks() []*IncludedPolicyLock {
	if m != nil {
		return m.IncludedPolicyLocks
	}
	return nil
}

func (m *Policyfile) GetCookbookLocks() []*CookbookLock {
	if m != nil {
		return m.CookbookLocks
	}
	return nil
}

func (m *Policyfile) GetDefaultAttributes() string {
	if m != nil {
		return m.DefaultAttributes
	}
	return ""
}

func (m *Policyfile) GetOverrideAttributes() string {
	if m != nil {
		return m.OverrideAttributes
	}
	return ""
}

func (m *Policyfile) GetExpandedRunList() []*ExpandedRunList {
	if m != nil {
		return m.ExpandedRunList
	}
	return nil
}

type IncludedPolicyLock struct {
	// Included Policyfile name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Policyfile revision ID.
	RevisionId string `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// Included policyfile source options.
	SourceOptions        *SourceOptions `protobuf:"bytes,3,opt,name=source_options,json=sourceOptions,proto3" json:"source_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IncludedPolicyLock) Reset()         { *m = IncludedPolicyLock{} }
func (m *IncludedPolicyLock) String() string { return proto.CompactTextString(m) }
func (*IncludedPolicyLock) ProtoMessage()    {}
func (*IncludedPolicyLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{3}
}

func (m *IncludedPolicyLock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncludedPolicyLock.Unmarshal(m, b)
}
func (m *IncludedPolicyLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncludedPolicyLock.Marshal(b, m, deterministic)
}
func (m *IncludedPolicyLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncludedPolicyLock.Merge(m, src)
}
func (m *IncludedPolicyLock) XXX_Size() int {
	return xxx_messageInfo_IncludedPolicyLock.Size(m)
}
func (m *IncludedPolicyLock) XXX_DiscardUnknown() {
	xxx_messageInfo_IncludedPolicyLock.DiscardUnknown(m)
}

var xxx_messageInfo_IncludedPolicyLock proto.InternalMessageInfo

func (m *IncludedPolicyLock) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IncludedPolicyLock) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

func (m *IncludedPolicyLock) GetSourceOptions() *SourceOptions {
	if m != nil {
		return m.SourceOptions
	}
	return nil
}

type CookbookLock struct {
	// Cookbook name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Cookbook version.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Cookbook identifier.
	Identifier string `protobuf:"bytes,3,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Cookbook decimal number identifier.
	DottedIdentifier string `protobuf:"bytes,4,opt,name=dotted_identifier,json=dottedIdentifier,proto3" json:"dotted_identifier,omitempty"`
	// Cookbook source.
	Source string `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	// Cookbook cache key.
	CacheKey string `protobuf:"bytes,6,opt,name=cache_key,json=cacheKey,proto3" json:"cache_key,omitempty"`
	// SCM detail.
	SCMDetail *SCMDetail `protobuf:"bytes,7,opt,name=SCMDetail,proto3" json:"SCMDetail,omitempty"`
	// Cookbook source path.
	SourceOptions        *SourceOptions `protobuf:"bytes,8,opt,name=source_options,json=sourceOptions,proto3" json:"source_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CookbookLock) Reset()         { *m = CookbookLock{} }
func (m *CookbookLock) String() string { return proto.CompactTextString(m) }
func (*CookbookLock) ProtoMessage()    {}
func (*CookbookLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{4}
}

func (m *CookbookLock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CookbookLock.Unmarshal(m, b)
}
func (m *CookbookLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CookbookLock.Marshal(b, m, deterministic)
}
func (m *CookbookLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CookbookLock.Merge(m, src)
}
func (m *CookbookLock) XXX_Size() int {
	return xxx_messageInfo_CookbookLock.Size(m)
}
func (m *CookbookLock) XXX_DiscardUnknown() {
	xxx_messageInfo_CookbookLock.DiscardUnknown(m)
}

var xxx_messageInfo_CookbookLock proto.InternalMessageInfo

func (m *CookbookLock) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CookbookLock) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CookbookLock) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *CookbookLock) GetDottedIdentifier() string {
	if m != nil {
		return m.DottedIdentifier
	}
	return ""
}

func (m *CookbookLock) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *CookbookLock) GetCacheKey() string {
	if m != nil {
		return m.CacheKey
	}
	return ""
}

func (m *CookbookLock) GetSCMDetail() *SCMDetail {
	if m != nil {
		return m.SCMDetail
	}
	return nil
}

func (m *CookbookLock) GetSourceOptions() *SourceOptions {
	if m != nil {
		return m.SourceOptions
	}
	return nil
}

type SCMDetail struct {
	// SCM name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// SCM remote location.
	Remote string `protobuf:"bytes,2,opt,name=remote,proto3" json:"remote,omitempty"`
	// SCM revision detail.
	Revision string `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// Boolean that denotes if the working tree is clean or not.
	WorkingTreeClean bool `protobuf:"varint,4,opt,name=working_tree_clean,json=workingTreeClean,proto3" json:"working_tree_clean,omitempty"`
	// Source's published information.
	Published bool `protobuf:"varint,5,opt,name=published,proto3" json:"published,omitempty"`
	// Synchronized remote branches list.
	SynchronizedRemoteBranches []string `protobuf:"bytes,6,rep,name=synchronized_remote_branches,json=synchronizedRemoteBranches,proto3" json:"synchronized_remote_branches,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *SCMDetail) Reset()         { *m = SCMDetail{} }
func (m *SCMDetail) String() string { return proto.CompactTextString(m) }
func (*SCMDetail) ProtoMessage()    {}
func (*SCMDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{5}
}

func (m *SCMDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SCMDetail.Unmarshal(m, b)
}
func (m *SCMDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SCMDetail.Marshal(b, m, deterministic)
}
func (m *SCMDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SCMDetail.Merge(m, src)
}
func (m *SCMDetail) XXX_Size() int {
	return xxx_messageInfo_SCMDetail.Size(m)
}
func (m *SCMDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_SCMDetail.DiscardUnknown(m)
}

var xxx_messageInfo_SCMDetail proto.InternalMessageInfo

func (m *SCMDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SCMDetail) GetRemote() string {
	if m != nil {
		return m.Remote
	}
	return ""
}

func (m *SCMDetail) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *SCMDetail) GetWorkingTreeClean() bool {
	if m != nil {
		return m.WorkingTreeClean
	}
	return false
}

func (m *SCMDetail) GetPublished() bool {
	if m != nil {
		return m.Published
	}
	return false
}

func (m *SCMDetail) GetSynchronizedRemoteBranches() []string {
	if m != nil {
		return m.SynchronizedRemoteBranches
	}
	return nil
}

type SourceOptions struct {
	// Source options path.
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceOptions) Reset()         { *m = SourceOptions{} }
func (m *SourceOptions) String() string { return proto.CompactTextString(m) }
func (*SourceOptions) ProtoMessage()    {}
func (*SourceOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{6}
}

func (m *SourceOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceOptions.Unmarshal(m, b)
}
func (m *SourceOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceOptions.Marshal(b, m, deterministic)
}
func (m *SourceOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceOptions.Merge(m, src)
}
func (m *SourceOptions) XXX_Size() int {
	return xxx_messageInfo_SourceOptions.Size(m)
}
func (m *SourceOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SourceOptions proto.InternalMessageInfo

func (m *SourceOptions) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type NamedRunList struct {
	// Run list name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Run list associated with the policy.
	RunList              []string `protobuf:"bytes,2,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamedRunList) Reset()         { *m = NamedRunList{} }
func (m *NamedRunList) String() string { return proto.CompactTextString(m) }
func (*NamedRunList) ProtoMessage()    {}
func (*NamedRunList) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fca7030dbc5441, []int{7}
}

func (m *NamedRunList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedRunList.Unmarshal(m, b)
}
func (m *NamedRunList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedRunList.Marshal(b, m, deterministic)
}
func (m *NamedRunList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedRunList.Merge(m, src)
}
func (m *NamedRunList) XXX_Size() int {
	return xxx_messageInfo_NamedRunList.Size(m)
}
func (m *NamedRunList) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedRunList.DiscardUnknown(m)
}

var xxx_messageInfo_NamedRunList proto.InternalMessageInfo

func (m *NamedRunList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedRunList) GetRunList() []string {
	if m != nil {
		return m.RunList
	}
	return nil
}

func init() {
	proto.RegisterType((*Policyfiles)(nil), "chef.automate.api.infra_proxy.response.Policyfiles")
	proto.RegisterType((*PolicyfileListItem)(nil), "chef.automate.api.infra_proxy.response.PolicyfileListItem")
	proto.RegisterType((*Policyfile)(nil), "chef.automate.api.infra_proxy.response.Policyfile")
	proto.RegisterType((*IncludedPolicyLock)(nil), "chef.automate.api.infra_proxy.response.IncludedPolicyLock")
	proto.RegisterType((*CookbookLock)(nil), "chef.automate.api.infra_proxy.response.CookbookLock")
	proto.RegisterType((*SCMDetail)(nil), "chef.automate.api.infra_proxy.response.SCMDetail")
	proto.RegisterType((*SourceOptions)(nil), "chef.automate.api.infra_proxy.response.SourceOptions")
	proto.RegisterType((*NamedRunList)(nil), "chef.automate.api.infra_proxy.response.NamedRunList")
}

func init() {
	proto.RegisterFile("api/external/infra_proxy/response/policyfiles.proto", fileDescriptor_64fca7030dbc5441)
}

var fileDescriptor_64fca7030dbc5441 = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xe1, 0x6a, 0xdb, 0x48,
	0x10, 0xc6, 0x8e, 0x63, 0xcb, 0x63, 0x27, 0x97, 0x6c, 0xb8, 0xa0, 0xcb, 0x85, 0xbb, 0x9c, 0x0f,
	0x4a, 0xa0, 0x8d, 0x44, 0x93, 0x96, 0x42, 0x68, 0xa1, 0x4d, 0x5a, 0x8a, 0x69, 0xda, 0x14, 0xa5,
	0xf4, 0x47, 0x5a, 0x10, 0xf2, 0x6a, 0x1c, 0x2f, 0x96, 0x77, 0xc5, 0xee, 0x2a, 0x8d, 0xfb, 0x42,
	0x7d, 0xae, 0x3e, 0x40, 0xde, 0xa1, 0x68, 0x25, 0xd9, 0x0a, 0x36, 0xd8, 0x81, 0xfe, 0xd3, 0xcc,
	0x7c, 0x33, 0xf3, 0xcd, 0x37, 0xab, 0x5d, 0x38, 0x0a, 0x62, 0xe6, 0xe2, 0x8d, 0x46, 0xc9, 0x83,
	0xc8, 0x65, 0xbc, 0x2f, 0x03, 0x3f, 0x96, 0xe2, 0x66, 0xec, 0x4a, 0x54, 0xb1, 0xe0, 0x0a, 0xdd,
	0x58, 0x44, 0x8c, 0x8e, 0xfb, 0x2c, 0x42, 0xe5, 0xc4, 0x52, 0x68, 0x41, 0x1e, 0xd0, 0x01, 0xf6,
	0x9d, 0x20, 0xd1, 0x62, 0x14, 0x68, 0x74, 0x82, 0x98, 0x39, 0xa5, 0x4c, 0xa7, 0xc8, 0xdc, 0x39,
	0x58, 0x5c, 0x5c, 0x8a, 0x49, 0xd9, 0x0e, 0x42, 0xeb, 0xe3, 0xb4, 0x17, 0xf9, 0x0c, 0x96, 0x69,
	0xcd, 0x50, 0xd9, 0x95, 0xbd, 0x95, 0xfd, 0xd6, 0xe1, 0xb1, 0xb3, 0x5c, 0x63, 0x67, 0x5a, 0xe6,
	0x8c, 0x29, 0xdd, 0xd5, 0x38, 0xf2, 0x26, 0xb5, 0x3a, 0x11, 0x90, 0xd9, 0x38, 0x21, 0x50, 0xe3,
	0xc1, 0x08, 0xed, 0xca, 0x5e, 0x65, 0xbf, 0xe9, 0x99, 0x6f, 0xf2, 0x2f, 0xb4, 0x24, 0x5e, 0x33,
	0xc5, 0x04, 0xf7, 0x59, 0x68, 0xaf, 0x98, 0x10, 0x14, 0xae, 0x6e, 0x48, 0xfe, 0x83, 0x76, 0xa6,
	0x8e, 0x7f, 0x25, 0x45, 0x12, 0xdb, 0x55, 0x83, 0x68, 0x65, 0xbe, 0xb7, 0xa9, 0xab, 0x73, 0x5b,
	0x03, 0x98, 0xb6, 0x9b, 0xdb, 0x66, 0x71, 0x95, 0xc5, 0x4c, 0xfe, 0x02, 0x4b, 0x26, 0xdc, 0x8f,
	0x98, 0xd2, 0x76, 0x6d, 0x6f, 0x65, 0xbf, 0xe9, 0x35, 0x64, 0xc2, 0xd3, 0xe9, 0xc8, 0x25, 0xac,
	0xa7, 0x6d, 0x42, 0x7f, 0x02, 0x58, 0x35, 0x6a, 0x3e, 0x59, 0x56, 0xcd, 0x0f, 0x69, 0xb6, 0x97,
	0x55, 0xf3, 0xda, 0xbc, 0x64, 0x11, 0x0e, 0x7f, 0x32, 0x4e, 0xa3, 0x24, 0xc4, 0xd0, 0xcf, 0x67,
	0x88, 0x04, 0x1d, 0x2a, 0xbb, 0x7e, 0xbf, 0x85, 0x75, 0xf3, 0x22, 0x99, 0x52, 0x67, 0x82, 0x0e,
	0xbd, 0x2d, 0x36, 0xe3, 0x53, 0xe4, 0x0b, 0xac, 0x53, 0x21, 0x86, 0x3d, 0x21, 0x86, 0x79, 0xa3,
	0xc6, 0xfd, 0x66, 0x39, 0xcd, 0xb3, 0x4d, 0x8b, 0x35, 0x5a, 0xb2, 0x14, 0x39, 0x00, 0x12, 0x62,
	0x3f, 0x48, 0x22, 0xed, 0x07, 0x5a, 0x4b, 0xd6, 0x4b, 0x34, 0x2a, 0xdb, 0x32, 0x5a, 0x6f, 0xe6,
	0x91, 0x57, 0x93, 0x00, 0x71, 0x61, 0x4b, 0x5c, 0xa3, 0x94, 0x2c, 0xc4, 0x32, 0xbe, 0x69, 0xf0,
	0xa4, 0x08, 0x95, 0x12, 0x28, 0x6c, 0xe2, 0x4d, 0x1c, 0xf0, 0xb0, 0xbc, 0x0b, 0x30, 0xfc, 0x9f,
	0x2d, 0xcb, 0xff, 0x4d, 0x5e, 0xa0, 0x58, 0xc7, 0x1f, 0x78, 0xd7, 0xd1, 0xf9, 0x51, 0x01, 0x32,
	0xab, 0xe6, 0x32, 0xc7, 0xbb, 0x3a, 0x73, 0xa8, 0xbe, 0xc2, 0xba, 0x12, 0x89, 0xa4, 0xe8, 0x8b,
	0x58, 0x33, 0xc1, 0x95, 0x39, 0x78, 0xad, 0xc3, 0xa7, 0xcb, 0xb2, 0xbd, 0x30, 0xd9, 0xe7, 0x59,
	0xb2, 0xb7, 0xa6, 0xca, 0x66, 0xe7, 0xb6, 0x0a, 0xed, 0xf2, 0x3a, 0xe6, 0x72, 0xb4, 0xa1, 0x71,
	0x8d, 0x32, 0xe5, 0x93, 0xf3, 0x2b, 0x4c, 0xf2, 0x0f, 0x00, 0x0b, 0x91, 0x6b, 0xd6, 0x67, 0x28,
	0x8b, 0x3f, 0x62, 0xea, 0x21, 0x0f, 0x61, 0x33, 0x14, 0x5a, 0x63, 0xe8, 0x97, 0x60, 0x35, 0x03,
	0xdb, 0xc8, 0x02, 0xdd, 0x29, 0x78, 0x1b, 0xea, 0x19, 0x39, 0x7b, 0xd5, 0x20, 0x72, 0x8b, 0xfc,
	0x0d, 0x4d, 0x1a, 0xd0, 0x01, 0xfa, 0x43, 0x1c, 0xdb, 0x75, 0x13, 0xb2, 0x8c, 0xe3, 0x1d, 0x8e,
	0xc9, 0x39, 0x34, 0x2f, 0x4e, 0xdf, 0xbf, 0x46, 0x1d, 0xb0, 0xc8, 0x6e, 0x18, 0x65, 0x1e, 0x2f,
	0xad, 0x4c, 0x91, 0xe8, 0x4d, 0x6b, 0xcc, 0xd1, 0xdb, 0xfa, 0x8d, 0x7a, 0xff, 0xac, 0x94, 0xf8,
	0xce, 0x15, 0x7b, 0x1b, 0xea, 0x12, 0x47, 0x42, 0x63, 0xae, 0x75, 0x6e, 0x91, 0x1d, 0xb0, 0x8a,
	0x53, 0x91, 0x0b, 0x3d, 0xb1, 0xc9, 0x23, 0x20, 0xdf, 0x84, 0x1c, 0x32, 0x7e, 0xe5, 0x6b, 0x89,
	0xe8, 0xd3, 0x08, 0x03, 0x6e, 0x74, 0xb6, 0xbc, 0x8d, 0x3c, 0xf2, 0x49, 0x22, 0x9e, 0xa6, 0x7e,
	0xb2, 0x0b, 0xcd, 0x38, 0xe9, 0x45, 0x4c, 0x0d, 0x30, 0x34, 0x52, 0x5b, 0xde, 0xd4, 0x41, 0x5e,
	0xc2, 0xae, 0x1a, 0x73, 0x3a, 0x90, 0x82, 0xb3, 0xef, 0xe9, 0x4f, 0x62, 0xda, 0xfb, 0x3d, 0x19,
	0x70, 0x3a, 0xc0, 0xec, 0x52, 0x69, 0x7a, 0x3b, 0x65, 0x8c, 0x67, 0x20, 0x27, 0x39, 0xa2, 0xf3,
	0x3f, 0xac, 0xdd, 0xd1, 0x20, 0x1d, 0x33, 0x0e, 0xf4, 0xa0, 0x18, 0x33, 0xfd, 0xee, 0xbc, 0x80,
	0x76, 0xf9, 0x4a, 0x9b, 0x2b, 0x45, 0xf9, 0x3e, 0xad, 0xde, 0xb9, 0x4f, 0x4f, 0x9e, 0x5f, 0x1e,
	0x5f, 0x31, 0x3d, 0x48, 0x7a, 0x0e, 0x15, 0x23, 0x37, 0xdd, 0x8c, 0x5b, 0x6c, 0xc6, 0x5d, 0xf8,
	0xe0, 0xf5, 0xea, 0xe6, 0xad, 0x3b, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0xae, 0xe0, 0x8a, 0x63,
	0x79, 0x07, 0x00, 0x00,
}