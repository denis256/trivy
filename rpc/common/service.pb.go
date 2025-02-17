// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/common/service.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Severity int32

const (
	Severity_UNKNOWN  Severity = 0
	Severity_LOW      Severity = 1
	Severity_MEDIUM   Severity = 2
	Severity_HIGH     Severity = 3
	Severity_CRITICAL Severity = 4
)

var Severity_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOW",
	2: "MEDIUM",
	3: "HIGH",
	4: "CRITICAL",
}

var Severity_value = map[string]int32{
	"UNKNOWN":  0,
	"LOW":      1,
	"MEDIUM":   2,
	"HIGH":     3,
	"CRITICAL": 4,
}

func (x Severity) String() string {
	return proto.EnumName(Severity_name, int32(x))
}

func (Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{0}
}

type OS struct {
	Family               string   `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Eosl                 bool     `protobuf:"varint,3,opt,name=eosl,proto3" json:"eosl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OS) Reset()         { *m = OS{} }
func (m *OS) String() string { return proto.CompactTextString(m) }
func (*OS) ProtoMessage()    {}
func (*OS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{0}
}

func (m *OS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OS.Unmarshal(m, b)
}
func (m *OS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OS.Marshal(b, m, deterministic)
}
func (m *OS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OS.Merge(m, src)
}
func (m *OS) XXX_Size() int {
	return xxx_messageInfo_OS.Size(m)
}
func (m *OS) XXX_DiscardUnknown() {
	xxx_messageInfo_OS.DiscardUnknown(m)
}

var xxx_messageInfo_OS proto.InternalMessageInfo

func (m *OS) GetFamily() string {
	if m != nil {
		return m.Family
	}
	return ""
}

func (m *OS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OS) GetEosl() bool {
	if m != nil {
		return m.Eosl
	}
	return false
}

type PackageInfo struct {
	FilePath             string     `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Packages             []*Package `protobuf:"bytes,2,rep,name=packages,proto3" json:"packages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PackageInfo) Reset()         { *m = PackageInfo{} }
func (m *PackageInfo) String() string { return proto.CompactTextString(m) }
func (*PackageInfo) ProtoMessage()    {}
func (*PackageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{1}
}

func (m *PackageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageInfo.Unmarshal(m, b)
}
func (m *PackageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageInfo.Marshal(b, m, deterministic)
}
func (m *PackageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageInfo.Merge(m, src)
}
func (m *PackageInfo) XXX_Size() int {
	return xxx_messageInfo_PackageInfo.Size(m)
}
func (m *PackageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PackageInfo proto.InternalMessageInfo

func (m *PackageInfo) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *PackageInfo) GetPackages() []*Package {
	if m != nil {
		return m.Packages
	}
	return nil
}

type Application struct {
	Type                 string     `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	FilePath             string     `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Libraries            []*Library `protobuf:"bytes,3,rep,name=libraries,proto3" json:"libraries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{2}
}

func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Application) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *Application) GetLibraries() []*Library {
	if m != nil {
		return m.Libraries
	}
	return nil
}

type Package struct {
	// binary package
	// e.g. bind-utils
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Release string `protobuf:"bytes,3,opt,name=release,proto3" json:"release,omitempty"`
	Epoch   int32  `protobuf:"varint,4,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Arch    string `protobuf:"bytes,5,opt,name=arch,proto3" json:"arch,omitempty"`
	// src package containing some binary packages
	// e.g. bind
	SrcName              string   `protobuf:"bytes,6,opt,name=src_name,json=srcName,proto3" json:"src_name,omitempty"`
	SrcVersion           string   `protobuf:"bytes,7,opt,name=src_version,json=srcVersion,proto3" json:"src_version,omitempty"`
	SrcRelease           string   `protobuf:"bytes,8,opt,name=src_release,json=srcRelease,proto3" json:"src_release,omitempty"`
	SrcEpoch             int32    `protobuf:"varint,9,opt,name=src_epoch,json=srcEpoch,proto3" json:"src_epoch,omitempty"`
	License              string   `protobuf:"bytes,10,opt,name=license,proto3" json:"license,omitempty"`
	Layer                *Layer   `protobuf:"bytes,11,opt,name=layer,proto3" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{3}
}

func (m *Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package.Unmarshal(m, b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package.Marshal(b, m, deterministic)
}
func (m *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(m, src)
}
func (m *Package) XXX_Size() int {
	return xxx_messageInfo_Package.Size(m)
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Package) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Package) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *Package) GetEpoch() int32 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Package) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

func (m *Package) GetSrcName() string {
	if m != nil {
		return m.SrcName
	}
	return ""
}

func (m *Package) GetSrcVersion() string {
	if m != nil {
		return m.SrcVersion
	}
	return ""
}

func (m *Package) GetSrcRelease() string {
	if m != nil {
		return m.SrcRelease
	}
	return ""
}

func (m *Package) GetSrcEpoch() int32 {
	if m != nil {
		return m.SrcEpoch
	}
	return 0
}

func (m *Package) GetLicense() string {
	if m != nil {
		return m.License
	}
	return ""
}

func (m *Package) GetLayer() *Layer {
	if m != nil {
		return m.Layer
	}
	return nil
}

type Library struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	License              string   `protobuf:"bytes,3,opt,name=license,proto3" json:"license,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Library) Reset()         { *m = Library{} }
func (m *Library) String() string { return proto.CompactTextString(m) }
func (*Library) ProtoMessage()    {}
func (*Library) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{4}
}

func (m *Library) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Library.Unmarshal(m, b)
}
func (m *Library) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Library.Marshal(b, m, deterministic)
}
func (m *Library) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Library.Merge(m, src)
}
func (m *Library) XXX_Size() int {
	return xxx_messageInfo_Library.Size(m)
}
func (m *Library) XXX_DiscardUnknown() {
	xxx_messageInfo_Library.DiscardUnknown(m)
}

var xxx_messageInfo_Library proto.InternalMessageInfo

func (m *Library) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Library) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Library) GetLicense() string {
	if m != nil {
		return m.License
	}
	return ""
}

type Misconfiguration struct {
	FileType             string           `protobuf:"bytes,1,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	FilePath             string           `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Successes            []*MisconfResult `protobuf:"bytes,3,rep,name=successes,proto3" json:"successes,omitempty"`
	Warnings             []*MisconfResult `protobuf:"bytes,4,rep,name=warnings,proto3" json:"warnings,omitempty"`
	Failures             []*MisconfResult `protobuf:"bytes,5,rep,name=failures,proto3" json:"failures,omitempty"`
	Exceptions           []*MisconfResult `protobuf:"bytes,6,rep,name=exceptions,proto3" json:"exceptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Misconfiguration) Reset()         { *m = Misconfiguration{} }
func (m *Misconfiguration) String() string { return proto.CompactTextString(m) }
func (*Misconfiguration) ProtoMessage()    {}
func (*Misconfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{5}
}

func (m *Misconfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Misconfiguration.Unmarshal(m, b)
}
func (m *Misconfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Misconfiguration.Marshal(b, m, deterministic)
}
func (m *Misconfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Misconfiguration.Merge(m, src)
}
func (m *Misconfiguration) XXX_Size() int {
	return xxx_messageInfo_Misconfiguration.Size(m)
}
func (m *Misconfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_Misconfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_Misconfiguration proto.InternalMessageInfo

func (m *Misconfiguration) GetFileType() string {
	if m != nil {
		return m.FileType
	}
	return ""
}

func (m *Misconfiguration) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *Misconfiguration) GetSuccesses() []*MisconfResult {
	if m != nil {
		return m.Successes
	}
	return nil
}

func (m *Misconfiguration) GetWarnings() []*MisconfResult {
	if m != nil {
		return m.Warnings
	}
	return nil
}

func (m *Misconfiguration) GetFailures() []*MisconfResult {
	if m != nil {
		return m.Failures
	}
	return nil
}

func (m *Misconfiguration) GetExceptions() []*MisconfResult {
	if m != nil {
		return m.Exceptions
	}
	return nil
}

type MisconfResult struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Severity             string   `protobuf:"bytes,6,opt,name=severity,proto3" json:"severity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MisconfResult) Reset()         { *m = MisconfResult{} }
func (m *MisconfResult) String() string { return proto.CompactTextString(m) }
func (*MisconfResult) ProtoMessage()    {}
func (*MisconfResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{6}
}

func (m *MisconfResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MisconfResult.Unmarshal(m, b)
}
func (m *MisconfResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MisconfResult.Marshal(b, m, deterministic)
}
func (m *MisconfResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MisconfResult.Merge(m, src)
}
func (m *MisconfResult) XXX_Size() int {
	return xxx_messageInfo_MisconfResult.Size(m)
}
func (m *MisconfResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MisconfResult.DiscardUnknown(m)
}

var xxx_messageInfo_MisconfResult proto.InternalMessageInfo

func (m *MisconfResult) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *MisconfResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MisconfResult) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MisconfResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MisconfResult) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MisconfResult) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

type DetectedMisconfiguration struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Namespace            string   `protobuf:"bytes,6,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Resolution           string   `protobuf:"bytes,7,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Severity             Severity `protobuf:"varint,8,opt,name=severity,proto3,enum=trivy.common.Severity" json:"severity,omitempty"`
	PrimaryUrl           string   `protobuf:"bytes,9,opt,name=primary_url,json=primaryUrl,proto3" json:"primary_url,omitempty"`
	References           []string `protobuf:"bytes,10,rep,name=references,proto3" json:"references,omitempty"`
	Status               string   `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Layer                *Layer   `protobuf:"bytes,12,opt,name=layer,proto3" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetectedMisconfiguration) Reset()         { *m = DetectedMisconfiguration{} }
func (m *DetectedMisconfiguration) String() string { return proto.CompactTextString(m) }
func (*DetectedMisconfiguration) ProtoMessage()    {}
func (*DetectedMisconfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{7}
}

func (m *DetectedMisconfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectedMisconfiguration.Unmarshal(m, b)
}
func (m *DetectedMisconfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectedMisconfiguration.Marshal(b, m, deterministic)
}
func (m *DetectedMisconfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectedMisconfiguration.Merge(m, src)
}
func (m *DetectedMisconfiguration) XXX_Size() int {
	return xxx_messageInfo_DetectedMisconfiguration.Size(m)
}
func (m *DetectedMisconfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectedMisconfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_DetectedMisconfiguration proto.InternalMessageInfo

func (m *DetectedMisconfiguration) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DetectedMisconfiguration) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DetectedMisconfiguration) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DetectedMisconfiguration) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DetectedMisconfiguration) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DetectedMisconfiguration) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DetectedMisconfiguration) GetResolution() string {
	if m != nil {
		return m.Resolution
	}
	return ""
}

func (m *DetectedMisconfiguration) GetSeverity() Severity {
	if m != nil {
		return m.Severity
	}
	return Severity_UNKNOWN
}

func (m *DetectedMisconfiguration) GetPrimaryUrl() string {
	if m != nil {
		return m.PrimaryUrl
	}
	return ""
}

func (m *DetectedMisconfiguration) GetReferences() []string {
	if m != nil {
		return m.References
	}
	return nil
}

func (m *DetectedMisconfiguration) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DetectedMisconfiguration) GetLayer() *Layer {
	if m != nil {
		return m.Layer
	}
	return nil
}

type Vulnerability struct {
	VulnerabilityId      string                 `protobuf:"bytes,1,opt,name=vulnerability_id,json=vulnerabilityId,proto3" json:"vulnerability_id,omitempty"`
	PkgName              string                 `protobuf:"bytes,2,opt,name=pkg_name,json=pkgName,proto3" json:"pkg_name,omitempty"`
	InstalledVersion     string                 `protobuf:"bytes,3,opt,name=installed_version,json=installedVersion,proto3" json:"installed_version,omitempty"`
	FixedVersion         string                 `protobuf:"bytes,4,opt,name=fixed_version,json=fixedVersion,proto3" json:"fixed_version,omitempty"`
	Title                string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description          string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Severity             Severity               `protobuf:"varint,7,opt,name=severity,proto3,enum=trivy.common.Severity" json:"severity,omitempty"`
	References           []string               `protobuf:"bytes,8,rep,name=references,proto3" json:"references,omitempty"`
	Layer                *Layer                 `protobuf:"bytes,10,opt,name=layer,proto3" json:"layer,omitempty"`
	SeveritySource       string                 `protobuf:"bytes,11,opt,name=severity_source,json=severitySource,proto3" json:"severity_source,omitempty"`
	Cvss                 map[string]*CVSS       `protobuf:"bytes,12,rep,name=cvss,proto3" json:"cvss,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CweIds               []string               `protobuf:"bytes,13,rep,name=cwe_ids,json=cweIds,proto3" json:"cwe_ids,omitempty"`
	PrimaryUrl           string                 `protobuf:"bytes,14,opt,name=primary_url,json=primaryUrl,proto3" json:"primary_url,omitempty"`
	PublishedDate        *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=published_date,json=publishedDate,proto3" json:"published_date,omitempty"`
	LastModifiedDate     *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=last_modified_date,json=lastModifiedDate,proto3" json:"last_modified_date,omitempty"`
	CustomAdvisoryData   *structpb.Value        `protobuf:"bytes,17,opt,name=custom_advisory_data,json=customAdvisoryData,proto3" json:"custom_advisory_data,omitempty"`
	CustomVulnData       *structpb.Value        `protobuf:"bytes,18,opt,name=custom_vuln_data,json=customVulnData,proto3" json:"custom_vuln_data,omitempty"`
	VendorIds            []string               `protobuf:"bytes,19,rep,name=vendor_ids,json=vendorIds,proto3" json:"vendor_ids,omitempty"`
	DataSource           *DataSource            `protobuf:"bytes,20,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Vulnerability) Reset()         { *m = Vulnerability{} }
func (m *Vulnerability) String() string { return proto.CompactTextString(m) }
func (*Vulnerability) ProtoMessage()    {}
func (*Vulnerability) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{8}
}

func (m *Vulnerability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vulnerability.Unmarshal(m, b)
}
func (m *Vulnerability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vulnerability.Marshal(b, m, deterministic)
}
func (m *Vulnerability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vulnerability.Merge(m, src)
}
func (m *Vulnerability) XXX_Size() int {
	return xxx_messageInfo_Vulnerability.Size(m)
}
func (m *Vulnerability) XXX_DiscardUnknown() {
	xxx_messageInfo_Vulnerability.DiscardUnknown(m)
}

var xxx_messageInfo_Vulnerability proto.InternalMessageInfo

func (m *Vulnerability) GetVulnerabilityId() string {
	if m != nil {
		return m.VulnerabilityId
	}
	return ""
}

func (m *Vulnerability) GetPkgName() string {
	if m != nil {
		return m.PkgName
	}
	return ""
}

func (m *Vulnerability) GetInstalledVersion() string {
	if m != nil {
		return m.InstalledVersion
	}
	return ""
}

func (m *Vulnerability) GetFixedVersion() string {
	if m != nil {
		return m.FixedVersion
	}
	return ""
}

func (m *Vulnerability) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Vulnerability) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Vulnerability) GetSeverity() Severity {
	if m != nil {
		return m.Severity
	}
	return Severity_UNKNOWN
}

func (m *Vulnerability) GetReferences() []string {
	if m != nil {
		return m.References
	}
	return nil
}

func (m *Vulnerability) GetLayer() *Layer {
	if m != nil {
		return m.Layer
	}
	return nil
}

func (m *Vulnerability) GetSeveritySource() string {
	if m != nil {
		return m.SeveritySource
	}
	return ""
}

func (m *Vulnerability) GetCvss() map[string]*CVSS {
	if m != nil {
		return m.Cvss
	}
	return nil
}

func (m *Vulnerability) GetCweIds() []string {
	if m != nil {
		return m.CweIds
	}
	return nil
}

func (m *Vulnerability) GetPrimaryUrl() string {
	if m != nil {
		return m.PrimaryUrl
	}
	return ""
}

func (m *Vulnerability) GetPublishedDate() *timestamppb.Timestamp {
	if m != nil {
		return m.PublishedDate
	}
	return nil
}

func (m *Vulnerability) GetLastModifiedDate() *timestamppb.Timestamp {
	if m != nil {
		return m.LastModifiedDate
	}
	return nil
}

func (m *Vulnerability) GetCustomAdvisoryData() *structpb.Value {
	if m != nil {
		return m.CustomAdvisoryData
	}
	return nil
}

func (m *Vulnerability) GetCustomVulnData() *structpb.Value {
	if m != nil {
		return m.CustomVulnData
	}
	return nil
}

func (m *Vulnerability) GetVendorIds() []string {
	if m != nil {
		return m.VendorIds
	}
	return nil
}

func (m *Vulnerability) GetDataSource() *DataSource {
	if m != nil {
		return m.DataSource
	}
	return nil
}

type DataSource struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataSource) Reset()         { *m = DataSource{} }
func (m *DataSource) String() string { return proto.CompactTextString(m) }
func (*DataSource) ProtoMessage()    {}
func (*DataSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{9}
}

func (m *DataSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataSource.Unmarshal(m, b)
}
func (m *DataSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataSource.Marshal(b, m, deterministic)
}
func (m *DataSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSource.Merge(m, src)
}
func (m *DataSource) XXX_Size() int {
	return xxx_messageInfo_DataSource.Size(m)
}
func (m *DataSource) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSource.DiscardUnknown(m)
}

var xxx_messageInfo_DataSource proto.InternalMessageInfo

func (m *DataSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataSource) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Layer struct {
	Digest               string   `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	DiffId               string   `protobuf:"bytes,2,opt,name=diff_id,json=diffId,proto3" json:"diff_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Layer) Reset()         { *m = Layer{} }
func (m *Layer) String() string { return proto.CompactTextString(m) }
func (*Layer) ProtoMessage()    {}
func (*Layer) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{10}
}

func (m *Layer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer.Unmarshal(m, b)
}
func (m *Layer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer.Marshal(b, m, deterministic)
}
func (m *Layer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer.Merge(m, src)
}
func (m *Layer) XXX_Size() int {
	return xxx_messageInfo_Layer.Size(m)
}
func (m *Layer) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer.DiscardUnknown(m)
}

var xxx_messageInfo_Layer proto.InternalMessageInfo

func (m *Layer) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *Layer) GetDiffId() string {
	if m != nil {
		return m.DiffId
	}
	return ""
}

type CVSS struct {
	V2Vector             string   `protobuf:"bytes,1,opt,name=v2_vector,json=v2Vector,proto3" json:"v2_vector,omitempty"`
	V3Vector             string   `protobuf:"bytes,2,opt,name=v3_vector,json=v3Vector,proto3" json:"v3_vector,omitempty"`
	V2Score              float64  `protobuf:"fixed64,3,opt,name=v2_score,json=v2Score,proto3" json:"v2_score,omitempty"`
	V3Score              float64  `protobuf:"fixed64,4,opt,name=v3_score,json=v3Score,proto3" json:"v3_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CVSS) Reset()         { *m = CVSS{} }
func (m *CVSS) String() string { return proto.CompactTextString(m) }
func (*CVSS) ProtoMessage()    {}
func (*CVSS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{11}
}

func (m *CVSS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CVSS.Unmarshal(m, b)
}
func (m *CVSS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CVSS.Marshal(b, m, deterministic)
}
func (m *CVSS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CVSS.Merge(m, src)
}
func (m *CVSS) XXX_Size() int {
	return xxx_messageInfo_CVSS.Size(m)
}
func (m *CVSS) XXX_DiscardUnknown() {
	xxx_messageInfo_CVSS.DiscardUnknown(m)
}

var xxx_messageInfo_CVSS proto.InternalMessageInfo

func (m *CVSS) GetV2Vector() string {
	if m != nil {
		return m.V2Vector
	}
	return ""
}

func (m *CVSS) GetV3Vector() string {
	if m != nil {
		return m.V3Vector
	}
	return ""
}

func (m *CVSS) GetV2Score() float64 {
	if m != nil {
		return m.V2Score
	}
	return 0
}

func (m *CVSS) GetV3Score() float64 {
	if m != nil {
		return m.V3Score
	}
	return 0
}

func init() {
	proto.RegisterEnum("trivy.common.Severity", Severity_name, Severity_value)
	proto.RegisterType((*OS)(nil), "trivy.common.OS")
	proto.RegisterType((*PackageInfo)(nil), "trivy.common.PackageInfo")
	proto.RegisterType((*Application)(nil), "trivy.common.Application")
	proto.RegisterType((*Package)(nil), "trivy.common.Package")
	proto.RegisterType((*Library)(nil), "trivy.common.Library")
	proto.RegisterType((*Misconfiguration)(nil), "trivy.common.Misconfiguration")
	proto.RegisterType((*MisconfResult)(nil), "trivy.common.MisconfResult")
	proto.RegisterType((*DetectedMisconfiguration)(nil), "trivy.common.DetectedMisconfiguration")
	proto.RegisterType((*Vulnerability)(nil), "trivy.common.Vulnerability")
	proto.RegisterMapType((map[string]*CVSS)(nil), "trivy.common.Vulnerability.CvssEntry")
	proto.RegisterType((*DataSource)(nil), "trivy.common.DataSource")
	proto.RegisterType((*Layer)(nil), "trivy.common.Layer")
	proto.RegisterType((*CVSS)(nil), "trivy.common.CVSS")
}

func init() { proto.RegisterFile("rpc/common/service.proto", fileDescriptor_6e749acacaaabfff) }

var fileDescriptor_6e749acacaaabfff = []byte{
	// 1253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x6e, 0xdb, 0xb6,
	0x17, 0xff, 0xfb, 0xdb, 0x3a, 0x4e, 0x52, 0x95, 0xed, 0xbf, 0x53, 0xd3, 0x6e, 0x35, 0x3c, 0x0c,
	0x4b, 0x37, 0xc0, 0x46, 0x9d, 0x8b, 0xb5, 0xeb, 0xcd, 0xb2, 0xa4, 0x58, 0x8d, 0x36, 0x69, 0xa7,
	0xb4, 0x29, 0x30, 0x60, 0x10, 0x18, 0x8a, 0x76, 0x88, 0xc8, 0x92, 0x46, 0x52, 0x4a, 0xfd, 0x02,
	0x7b, 0x8b, 0xbd, 0xc6, 0x2e, 0x86, 0x3d, 0xdc, 0xc0, 0x0f, 0xc9, 0x52, 0x12, 0x2c, 0xdd, 0x95,
	0x78, 0x3e, 0xf8, 0x3b, 0x87, 0xbf, 0x73, 0x0e, 0x45, 0xf0, 0x78, 0x4a, 0x26, 0x24, 0x59, 0x2e,
	0x93, 0x78, 0x22, 0x28, 0xcf, 0x19, 0xa1, 0xe3, 0x94, 0x27, 0x32, 0x41, 0x1b, 0x92, 0xb3, 0x7c,
	0x35, 0x36, 0xb6, 0xed, 0x47, 0x8b, 0x24, 0x59, 0x44, 0x74, 0xa2, 0x6d, 0xa7, 0xd9, 0x7c, 0x22,
	0xd9, 0x92, 0x0a, 0x89, 0x97, 0xa9, 0x71, 0xdf, 0x7e, 0x78, 0xd9, 0x41, 0x48, 0x9e, 0x11, 0x69,
	0xac, 0xa3, 0x03, 0x68, 0xbe, 0x39, 0x46, 0xf7, 0xa0, 0x3b, 0xc7, 0x4b, 0x16, 0xad, 0xbc, 0xc6,
	0xb0, 0xb1, 0xe3, 0xf8, 0x56, 0x42, 0x08, 0xda, 0x31, 0x5e, 0x52, 0xaf, 0xa9, 0xb5, 0x7a, 0xad,
	0x74, 0x34, 0x11, 0x91, 0xd7, 0x1a, 0x36, 0x76, 0xfa, 0xbe, 0x5e, 0x8f, 0x7e, 0x85, 0xc1, 0x5b,
	0x4c, 0xce, 0xf1, 0x82, 0xce, 0xe2, 0x79, 0x82, 0x1e, 0x80, 0x33, 0x67, 0x11, 0x0d, 0x52, 0x2c,
	0xcf, 0x2c, 0x62, 0x5f, 0x29, 0xde, 0x62, 0x79, 0x86, 0x9e, 0x40, 0x3f, 0x35, 0xbe, 0xc2, 0x6b,
	0x0e, 0x5b, 0x3b, 0x83, 0xe9, 0xff, 0xc7, 0xd5, 0x13, 0x8d, 0x2d, 0x92, 0x5f, 0xba, 0x8d, 0x04,
	0x0c, 0xf6, 0xd2, 0x34, 0x62, 0x04, 0x4b, 0x96, 0xc4, 0x2a, 0x03, 0xb9, 0x4a, 0xa9, 0x45, 0xd6,
	0xeb, 0x7a, 0xc8, 0xe6, 0xa5, 0x90, 0xbb, 0xe0, 0x44, 0xec, 0x94, 0x63, 0xce, 0xa8, 0xf0, 0x5a,
	0xd7, 0xc5, 0x7c, 0xad, 0xcd, 0x2b, 0x7f, 0xed, 0x37, 0xfa, 0xab, 0x09, 0x3d, 0x9b, 0x4a, 0xc9,
	0x43, 0xa3, 0xc2, 0x83, 0x07, 0xbd, 0x9c, 0x72, 0xc1, 0x92, 0xd8, 0xc6, 0x2b, 0x44, 0x65, 0xe1,
	0x34, 0xa2, 0x58, 0x50, 0x4d, 0x92, 0xe3, 0x17, 0x22, 0xba, 0x0b, 0x1d, 0x9a, 0x26, 0xe4, 0xcc,
	0x6b, 0x0f, 0x1b, 0x3b, 0x1d, 0xdf, 0x08, 0x0a, 0x1d, 0x73, 0x72, 0xe6, 0x75, 0x0c, 0xba, 0x5a,
	0xa3, 0xfb, 0xd0, 0x17, 0x9c, 0x04, 0x3a, 0x6a, 0xd7, 0x80, 0x08, 0x4e, 0x8e, 0x54, 0xe0, 0x47,
	0x30, 0x50, 0xa6, 0x22, 0x78, 0x4f, 0x5b, 0x41, 0x70, 0x72, 0x62, 0xe3, 0x5b, 0x87, 0x22, 0x87,
	0x7e, 0xe9, 0xe0, 0xdb, 0x34, 0x1e, 0x80, 0xa3, 0x1c, 0x4c, 0x2a, 0x8e, 0x4e, 0x45, 0x45, 0x7b,
	0xa1, 0xb3, 0xf1, 0xa0, 0x17, 0x31, 0x42, 0x63, 0x41, 0x3d, 0x30, 0x81, 0xad, 0x88, 0x1e, 0x43,
	0x27, 0xc2, 0x2b, 0xca, 0xbd, 0xc1, 0xb0, 0xb1, 0x33, 0x98, 0xde, 0xb9, 0x44, 0xa1, 0x32, 0xf9,
	0xc6, 0x63, 0xf4, 0x33, 0xf4, 0x2c, 0xa5, 0xff, 0x9d, 0xbb, 0x22, 0x7a, 0xab, 0x16, 0x7d, 0xf4,
	0x67, 0x13, 0xdc, 0x43, 0x26, 0x48, 0x12, 0xcf, 0xd9, 0x22, 0xe3, 0xa6, 0x15, 0x8a, 0xb2, 0x57,
	0xfa, 0x41, 0x97, 0xfd, 0xdd, 0x8d, 0x3d, 0xf1, 0x0c, 0x1c, 0x91, 0x11, 0x42, 0x85, 0x28, 0x7b,
	0xe2, 0x41, 0xfd, 0x40, 0x36, 0x98, 0x4f, 0x45, 0x16, 0x49, 0x7f, 0xed, 0x8d, 0xbe, 0x83, 0xfe,
	0x05, 0xe6, 0x31, 0x8b, 0x17, 0xc2, 0x6b, 0xdf, 0xbc, 0xb3, 0x74, 0x56, 0x1b, 0xe7, 0x98, 0x45,
	0x19, 0xa7, 0xc2, 0xeb, 0x7c, 0xc2, 0xc6, 0xc2, 0x19, 0x3d, 0x07, 0xa0, 0x1f, 0x09, 0x4d, 0xd5,
	0x99, 0x85, 0xd7, 0xbd, 0x79, 0x6b, 0xc5, 0x7d, 0xf4, 0x47, 0x03, 0x36, 0x6b, 0x56, 0xf4, 0x10,
	0x1c, 0x55, 0x06, 0x91, 0x62, 0x52, 0xb0, 0xb6, 0x56, 0xa8, 0x12, 0x2c, 0xa9, 0x10, 0x78, 0x51,
	0xcc, 0x7d, 0x21, 0x96, 0x83, 0xd7, 0xaa, 0x0c, 0xde, 0x16, 0x34, 0x59, 0xa8, 0xfb, 0xd9, 0xf1,
	0x9b, 0x2c, 0x54, 0x2d, 0x2e, 0x99, 0x8c, 0xa8, 0xed, 0x66, 0x23, 0xa0, 0x6d, 0xe8, 0x0b, 0x9a,
	0x53, 0xce, 0xe4, 0xca, 0xb6, 0x73, 0x29, 0x8f, 0x7e, 0x6f, 0x81, 0x77, 0x40, 0x25, 0x25, 0x92,
	0x86, 0x57, 0x0a, 0x7c, 0xdd, 0xac, 0x9b, 0x90, 0xcd, 0xab, 0x21, 0x5b, 0xd5, 0x90, 0x43, 0x18,
	0x84, 0x54, 0x10, 0xce, 0x34, 0x0d, 0x36, 0xc3, 0xaa, 0xaa, 0x7a, 0xd0, 0x4e, 0xfd, 0xa0, 0x35,
	0x82, 0xba, 0x97, 0x09, 0xfa, 0x02, 0x80, 0x53, 0x91, 0x44, 0x99, 0xac, 0xcc, 0xdf, 0x5a, 0x83,
	0xa6, 0x95, 0xc3, 0xaa, 0xe1, 0xdb, 0x9a, 0xde, 0xab, 0xd7, 0xea, 0xd8, 0x5a, 0xd7, 0x24, 0xa8,
	0x99, 0x4d, 0x39, 0x5b, 0x62, 0xbe, 0x0a, 0x32, 0x1e, 0xe9, 0xa1, 0x74, 0x7c, 0xb0, 0xaa, 0xf7,
	0x3c, 0x32, 0x41, 0xe7, 0x94, 0xd3, 0x98, 0x50, 0xe1, 0xc1, 0xb0, 0x65, 0x82, 0x16, 0x1a, 0x75,
	0x85, 0x0b, 0x89, 0x65, 0x26, 0xf4, 0x74, 0x3a, 0xbe, 0x95, 0xd6, 0x43, 0xbb, 0x71, 0xe3, 0xd0,
	0xfe, 0xdd, 0x83, 0xcd, 0x93, 0x2c, 0x8a, 0x29, 0xc7, 0xa7, 0x2c, 0x52, 0x59, 0x3d, 0x06, 0x37,
	0xaf, 0x2a, 0x02, 0x16, 0xda, 0x4a, 0xdc, 0xaa, 0xe9, 0x67, 0xa1, 0xba, 0xb0, 0xd2, 0xf3, 0x45,
	0x50, 0xf9, 0x5d, 0xf4, 0xd2, 0xf3, 0x85, 0xbe, 0xb0, 0xbe, 0x85, 0xdb, 0x2c, 0x16, 0x12, 0x47,
	0x11, 0x0d, 0xcb, 0x6b, 0xcb, 0xd4, 0xca, 0x2d, 0x0d, 0xc5, 0xe5, 0xf5, 0x25, 0x6c, 0xce, 0xd9,
	0xc7, 0x8a, 0xa3, 0x29, 0xdc, 0x86, 0x56, 0x16, 0x4e, 0xd7, 0x37, 0xd9, 0xa5, 0x8a, 0x77, 0xaf,
	0x56, 0xbc, 0x5a, 0x99, 0xde, 0x27, 0x56, 0xa6, 0x4e, 0x7c, 0xff, 0x0a, 0xf1, 0x25, 0xc1, 0x70,
	0x13, 0xc1, 0xe8, 0x6b, 0xb8, 0x55, 0xc0, 0x06, 0x22, 0xc9, 0x38, 0xa1, 0xb6, 0x58, 0x5b, 0x85,
	0xfa, 0x58, 0x6b, 0xd1, 0x33, 0x68, 0x93, 0x5c, 0x08, 0x6f, 0x43, 0x4f, 0xfa, 0x57, 0x75, 0xc8,
	0x5a, 0x89, 0xc6, 0xfb, 0xb9, 0x10, 0x2f, 0x62, 0xc9, 0x57, 0xbe, 0xde, 0x82, 0x3e, 0x83, 0x1e,
	0xb9, 0xa0, 0x01, 0x0b, 0x85, 0xb7, 0xa9, 0x73, 0xed, 0x92, 0x0b, 0x3a, 0x0b, 0xc5, 0xe5, 0x0e,
	0xdb, 0xba, 0xd2, 0x61, 0x7b, 0xb0, 0x95, 0x66, 0xa7, 0x11, 0x13, 0x67, 0x34, 0x0c, 0x42, 0x2c,
	0xa9, 0x77, 0x4b, 0x9f, 0x68, 0x7b, 0x6c, 0x5e, 0x10, 0xe3, 0xe2, 0x05, 0x31, 0x7e, 0x57, 0x3c,
	0x31, 0xfc, 0xcd, 0x72, 0xc7, 0x01, 0x96, 0x14, 0xbd, 0x04, 0x14, 0x61, 0x21, 0x83, 0x65, 0x12,
	0xb2, 0x39, 0x2b, 0x60, 0xdc, 0x1b, 0x61, 0x5c, 0xb5, 0xeb, 0xd0, 0x6e, 0xb2, 0x48, 0x77, 0x49,
	0x26, 0x64, 0xb2, 0x0c, 0x70, 0x98, 0x33, 0x91, 0xf0, 0x95, 0xc2, 0xc2, 0xde, 0x6d, 0x8d, 0x75,
	0xef, 0x0a, 0xd6, 0x09, 0x8e, 0x32, 0xea, 0x23, 0xb3, 0x67, 0xcf, 0x6e, 0x39, 0xc0, 0x12, 0xa3,
	0x1f, 0xc0, 0xb5, 0x48, 0xaa, 0x65, 0x0d, 0x0a, 0xfa, 0x57, 0x94, 0x2d, 0xe3, 0xaf, 0x78, 0xd6,
	0x08, 0x9f, 0x03, 0xe4, 0x34, 0x0e, 0x13, 0xae, 0x59, 0xbd, 0xa3, 0x59, 0x75, 0x8c, 0x46, 0x11,
	0xfb, 0x0c, 0x06, 0x0a, 0xb4, 0xa8, 0xe8, 0x5d, 0x8d, 0xed, 0xd5, 0x6b, 0xa6, 0x70, 0x4c, 0x6d,
	0x7d, 0x08, 0xcb, 0xf5, 0xf6, 0x2b, 0x70, 0xca, 0xfa, 0x21, 0x17, 0x5a, 0xe7, 0xb4, 0x78, 0x81,
	0xa9, 0x25, 0xda, 0x81, 0x4e, 0xae, 0x32, 0xd2, 0x03, 0x35, 0x98, 0xa2, 0x3a, 0xe6, 0xfe, 0xc9,
	0xf1, 0xb1, 0x6f, 0x1c, 0xbe, 0x6f, 0x3e, 0x6d, 0x8c, 0xa6, 0x00, 0xeb, 0x30, 0xd7, 0xfe, 0x76,
	0x5d, 0x68, 0xa9, 0xd2, 0x9b, 0xf1, 0x54, 0xcb, 0xd1, 0x53, 0xe8, 0xe8, 0x0e, 0x55, 0xd7, 0x47,
	0xc8, 0x16, 0x54, 0xc8, 0xe2, 0x05, 0x68, 0x24, 0xd5, 0x4e, 0x21, 0x9b, 0xcf, 0x83, 0xf2, 0xc2,
	0xed, 0x2a, 0x71, 0x16, 0x8e, 0x72, 0x68, 0xab, 0x04, 0xd4, 0x4f, 0x36, 0x9f, 0x06, 0x39, 0x25,
	0x32, 0xe1, 0xc5, 0x1f, 0x38, 0x9f, 0x9e, 0x68, 0x59, 0x1b, 0x77, 0x0b, 0xa3, 0xfd, 0x03, 0xe7,
	0xbb, 0xd6, 0x78, 0x1f, 0xfa, 0xf9, 0x34, 0x10, 0x24, 0xe1, 0xe6, 0xe6, 0x6e, 0xf8, 0xbd, 0x7c,
	0x7a, 0xac, 0x44, 0x6d, 0xda, 0xb5, 0xa6, 0xb6, 0x35, 0xed, 0x6a, 0xd3, 0x37, 0x07, 0xd0, 0x2f,
	0x86, 0x14, 0x0d, 0xa0, 0xf7, 0xfe, 0xe8, 0xd5, 0xd1, 0x9b, 0x0f, 0x47, 0xee, 0xff, 0x50, 0x0f,
	0x5a, 0xaf, 0xdf, 0x7c, 0x70, 0x1b, 0x08, 0xa0, 0x7b, 0xf8, 0xe2, 0x60, 0xf6, 0xfe, 0xd0, 0x6d,
	0xa2, 0x3e, 0xb4, 0x5f, 0xce, 0x7e, 0x7a, 0xe9, 0xb6, 0xd0, 0x06, 0xf4, 0xf7, 0xfd, 0xd9, 0xbb,
	0xd9, 0xfe, 0xde, 0x6b, 0xb7, 0xfd, 0xe3, 0x93, 0x5f, 0x26, 0x0b, 0x26, 0xcf, 0xb2, 0x53, 0x45,
	0xe6, 0x04, 0xff, 0x96, 0x61, 0x41, 0x49, 0xa6, 0x40, 0x27, 0x9a, 0xe3, 0xc9, 0xfa, 0xf5, 0xfd,
	0xdc, 0x7c, 0x4e, 0xbb, 0xba, 0x4b, 0x76, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x8e, 0x21,
	0xdb, 0x99, 0x0b, 0x00, 0x00,
}
