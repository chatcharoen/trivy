// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/common/service.proto

package common

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

type Library struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
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

type Vulnerability struct {
	VulnerabilityId      string   `protobuf:"bytes,1,opt,name=vulnerability_id,json=vulnerabilityId,proto3" json:"vulnerability_id,omitempty"`
	PkgName              string   `protobuf:"bytes,2,opt,name=pkg_name,json=pkgName,proto3" json:"pkg_name,omitempty"`
	InstalledVersion     string   `protobuf:"bytes,3,opt,name=installed_version,json=installedVersion,proto3" json:"installed_version,omitempty"`
	FixedVersion         string   `protobuf:"bytes,4,opt,name=fixed_version,json=fixedVersion,proto3" json:"fixed_version,omitempty"`
	Title                string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Severity             Severity `protobuf:"varint,7,opt,name=severity,proto3,enum=trivy.common.Severity" json:"severity,omitempty"`
	References           []string `protobuf:"bytes,8,rep,name=references,proto3" json:"references,omitempty"`
	Layer                *Layer   `protobuf:"bytes,10,opt,name=layer,proto3" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vulnerability) Reset()         { *m = Vulnerability{} }
func (m *Vulnerability) String() string { return proto.CompactTextString(m) }
func (*Vulnerability) ProtoMessage()    {}
func (*Vulnerability) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e749acacaaabfff, []int{5}
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
	return fileDescriptor_6e749acacaaabfff, []int{6}
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

func init() {
	proto.RegisterEnum("trivy.common.Severity", Severity_name, Severity_value)
	proto.RegisterType((*OS)(nil), "trivy.common.OS")
	proto.RegisterType((*PackageInfo)(nil), "trivy.common.PackageInfo")
	proto.RegisterType((*Application)(nil), "trivy.common.Application")
	proto.RegisterType((*Package)(nil), "trivy.common.Package")
	proto.RegisterType((*Library)(nil), "trivy.common.Library")
	proto.RegisterType((*Vulnerability)(nil), "trivy.common.Vulnerability")
	proto.RegisterType((*Layer)(nil), "trivy.common.Layer")
}

func init() { proto.RegisterFile("rpc/common/service.proto", fileDescriptor_6e749acacaaabfff) }

var fileDescriptor_6e749acacaaabfff = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xce, 0x87, 0x9d, 0x71, 0x0b, 0x66, 0x29, 0xc5, 0xa8, 0x12, 0x44, 0xe6, 0x92, 0x82,
	0x94, 0x42, 0x7a, 0x80, 0x6b, 0x69, 0x2b, 0x6a, 0xd1, 0xa6, 0x95, 0x4b, 0x5b, 0x09, 0x09, 0x45,
	0x5b, 0x7b, 0x9c, 0xac, 0xea, 0xd8, 0xd6, 0xae, 0x89, 0xf0, 0x9f, 0xe5, 0x5f, 0x70, 0x47, 0xbb,
	0x5e, 0xa7, 0x49, 0xc5, 0x85, 0xdb, 0xbe, 0x79, 0x2f, 0xf3, 0xde, 0x4c, 0x76, 0x0d, 0x1e, 0x2f,
	0xa2, 0xbd, 0x28, 0x9f, 0xcf, 0xf3, 0x6c, 0x4f, 0x20, 0x5f, 0xb0, 0x08, 0x87, 0x05, 0xcf, 0xcb,
	0x9c, 0x6c, 0x94, 0x9c, 0x2d, 0xaa, 0x61, 0xcd, 0xf9, 0xef, 0xc1, 0x3c, 0xbf, 0x24, 0xdb, 0xd0,
	0x4d, 0xe8, 0x9c, 0xa5, 0x95, 0x67, 0xf4, 0x8d, 0x41, 0x2f, 0xd4, 0x88, 0x10, 0x68, 0x67, 0x74,
	0x8e, 0x9e, 0xa9, 0xaa, 0xea, 0xec, 0xff, 0x00, 0xe7, 0x82, 0x46, 0x77, 0x74, 0x8a, 0x41, 0x96,
	0xe4, 0x64, 0x07, 0x7a, 0x09, 0x4b, 0x71, 0x52, 0xd0, 0x72, 0xa6, 0x7f, 0x6d, 0xcb, 0xc2, 0x05,
	0x2d, 0x67, 0xe4, 0x03, 0xd8, 0x45, 0xad, 0x15, 0x9e, 0xd9, 0x6f, 0x0d, 0x9c, 0xd1, 0xf3, 0xe1,
	0xaa, 0xfd, 0x50, 0x77, 0x0a, 0x97, 0x32, 0x5f, 0x80, 0x73, 0x50, 0x14, 0x29, 0x8b, 0x68, 0xc9,
	0xf2, 0x4c, 0x26, 0x28, 0xab, 0x02, 0x75, 0x67, 0x75, 0x5e, 0xb7, 0x34, 0x1f, 0x58, 0xee, 0x43,
	0x2f, 0x65, 0xb7, 0x9c, 0x72, 0x86, 0xc2, 0x6b, 0xfd, 0xcb, 0xf3, 0x54, 0xd1, 0x55, 0x78, 0xaf,
	0xf3, 0xff, 0x18, 0x60, 0xe9, 0x28, 0xcb, 0x99, 0x8d, 0xfb, 0x99, 0x89, 0x07, 0xd6, 0x02, 0xb9,
	0x60, 0x79, 0xa6, 0xfd, 0x1a, 0x28, 0x19, 0x8e, 0x29, 0x52, 0x81, 0x5e, 0xab, 0x66, 0x34, 0x24,
	0x5b, 0xd0, 0xc1, 0x22, 0x8f, 0x66, 0x5e, 0xbb, 0x6f, 0x0c, 0x3a, 0x61, 0x0d, 0x64, 0x77, 0xca,
	0xa3, 0x99, 0xd7, 0xa9, 0xbb, 0xcb, 0x33, 0x79, 0x09, 0xb6, 0xe0, 0xd1, 0x44, 0xb9, 0x76, 0xeb,
	0x26, 0x82, 0x47, 0x63, 0x69, 0xfc, 0x1a, 0x1c, 0x49, 0x35, 0xe6, 0x96, 0x62, 0x41, 0xf0, 0xe8,
	0x5a, 0xfb, 0x6b, 0x41, 0x93, 0xc1, 0x5e, 0x0a, 0x42, 0x1d, 0x63, 0x07, 0x7a, 0x52, 0x50, 0x47,
	0xe9, 0xa9, 0x28, 0xd2, 0xed, 0x58, 0x62, 0xff, 0x23, 0x58, 0x7a, 0x1b, 0xff, 0x37, 0xb6, 0xff,
	0xdb, 0x84, 0xcd, 0xeb, 0x9f, 0x69, 0x86, 0x9c, 0xde, 0xb2, 0x94, 0x95, 0x15, 0xd9, 0x05, 0x77,
	0xb1, 0x5a, 0x98, 0xb0, 0x58, 0xf7, 0x7a, 0xb2, 0x56, 0x0f, 0x62, 0x39, 0x6f, 0x71, 0x37, 0x9d,
	0xac, 0xdc, 0x2c, 0xab, 0xb8, 0x9b, 0xaa, 0x79, 0xdf, 0xc1, 0x53, 0x96, 0x89, 0x92, 0xa6, 0x29,
	0xc6, 0xcb, 0xa9, 0xeb, 0xc5, 0xba, 0x4b, 0xa2, 0x99, 0xfd, 0x0d, 0x6c, 0x26, 0xec, 0xd7, 0x8a,
	0xb0, 0xad, 0x84, 0x1b, 0xaa, 0xd8, 0x88, 0xb6, 0xa0, 0x53, 0xb2, 0x32, 0x45, 0xbd, 0xf1, 0x1a,
	0x90, 0x3e, 0x38, 0x31, 0x8a, 0x88, 0xb3, 0x42, 0xde, 0x32, 0xbd, 0xf5, 0xd5, 0x12, 0x19, 0x81,
	0x2d, 0x70, 0x81, 0x9c, 0x95, 0x95, 0x5a, 0xfb, 0xe3, 0xd1, 0xf6, 0xfa, 0x35, 0xba, 0xd4, 0x6c,
	0xb8, 0xd4, 0x91, 0x57, 0x00, 0x1c, 0x13, 0xe4, 0x98, 0x45, 0x28, 0x3c, 0xbb, 0xdf, 0x92, 0xff,
	0xc5, 0x7d, 0x85, 0xec, 0x42, 0x27, 0xa5, 0x15, 0x72, 0x0f, 0xfa, 0xc6, 0xc0, 0x19, 0x3d, 0x7b,
	0x70, 0x2f, 0x25, 0x15, 0xd6, 0x0a, 0xff, 0x13, 0x74, 0x14, 0x96, 0x4f, 0x33, 0x66, 0x53, 0x14,
	0x65, 0xf3, 0x34, 0x6b, 0x44, 0x5e, 0x80, 0x15, 0xb3, 0x24, 0x91, 0x6b, 0x36, 0x1b, 0x22, 0x49,
	0x82, 0xf8, 0xed, 0x11, 0xd8, 0x4d, 0x34, 0xe2, 0x80, 0x75, 0x35, 0xfe, 0x3a, 0x3e, 0xbf, 0x19,
	0xbb, 0x8f, 0x88, 0x05, 0xad, 0xd3, 0xf3, 0x1b, 0xd7, 0x20, 0x00, 0xdd, 0xb3, 0xe3, 0xa3, 0xe0,
	0xea, 0xcc, 0x35, 0x89, 0x0d, 0xed, 0x93, 0xe0, 0xcb, 0x89, 0xdb, 0x22, 0x1b, 0x60, 0x1f, 0x86,
	0xc1, 0xb7, 0xe0, 0xf0, 0xe0, 0xd4, 0x6d, 0x7f, 0xb6, 0xbf, 0x77, 0xeb, 0x58, 0xb7, 0x5d, 0xf5,
	0xd9, 0xd8, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xec, 0xc8, 0x65, 0x39, 0x52, 0x04, 0x00, 0x00,
}