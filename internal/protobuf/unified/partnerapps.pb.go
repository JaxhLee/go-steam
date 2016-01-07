// Code generated by protoc-gen-go.
// source: steammessages_partnerapps.steamclient.proto
// DO NOT EDIT!

package unified

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package unified is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type CPartnerApps_RequestUploadToken_Request struct {
	Filename         *string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Appid            *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_RequestUploadToken_Request) Reset() {
	*m = CPartnerApps_RequestUploadToken_Request{}
}
func (m *CPartnerApps_RequestUploadToken_Request) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_RequestUploadToken_Request) ProtoMessage()    {}
func (*CPartnerApps_RequestUploadToken_Request) Descriptor() ([]byte, []int) {
	return partnerapps_fileDescriptor0, []int{0}
}

func (m *CPartnerApps_RequestUploadToken_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CPartnerApps_RequestUploadToken_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type CPartnerApps_RequestUploadToken_Response struct {
	UploadToken      *uint64 `protobuf:"varint,1,opt,name=upload_token" json:"upload_token,omitempty"`
	Location         *string `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	RoutingId        *uint64 `protobuf:"varint,3,opt,name=routing_id" json:"routing_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_RequestUploadToken_Response) Reset() {
	*m = CPartnerApps_RequestUploadToken_Response{}
}
func (m *CPartnerApps_RequestUploadToken_Response) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_RequestUploadToken_Response) ProtoMessage()    {}
func (*CPartnerApps_RequestUploadToken_Response) Descriptor() ([]byte, []int) {
	return partnerapps_fileDescriptor0, []int{1}
}

func (m *CPartnerApps_RequestUploadToken_Response) GetUploadToken() uint64 {
	if m != nil && m.UploadToken != nil {
		return *m.UploadToken
	}
	return 0
}

func (m *CPartnerApps_RequestUploadToken_Response) GetLocation() string {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return ""
}

func (m *CPartnerApps_RequestUploadToken_Response) GetRoutingId() uint64 {
	if m != nil && m.RoutingId != nil {
		return *m.RoutingId
	}
	return 0
}

type CPartnerApps_FinishUpload_Request struct {
	UploadToken      *uint64 `protobuf:"varint,1,opt,name=upload_token" json:"upload_token,omitempty"`
	RoutingId        *uint64 `protobuf:"varint,2,opt,name=routing_id" json:"routing_id,omitempty"`
	AppId            *uint32 `protobuf:"varint,3,opt,name=app_id" json:"app_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_FinishUpload_Request) Reset()         { *m = CPartnerApps_FinishUpload_Request{} }
func (m *CPartnerApps_FinishUpload_Request) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_FinishUpload_Request) ProtoMessage()    {}
func (*CPartnerApps_FinishUpload_Request) Descriptor() ([]byte, []int) {
	return partnerapps_fileDescriptor0, []int{2}
}

func (m *CPartnerApps_FinishUpload_Request) GetUploadToken() uint64 {
	if m != nil && m.UploadToken != nil {
		return *m.UploadToken
	}
	return 0
}

func (m *CPartnerApps_FinishUpload_Request) GetRoutingId() uint64 {
	if m != nil && m.RoutingId != nil {
		return *m.RoutingId
	}
	return 0
}

func (m *CPartnerApps_FinishUpload_Request) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

type CPartnerApps_FinishUploadKVSign_Response struct {
	SignedInstallscript *string `protobuf:"bytes,1,opt,name=signed_installscript" json:"signed_installscript,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *CPartnerApps_FinishUploadKVSign_Response) Reset() {
	*m = CPartnerApps_FinishUploadKVSign_Response{}
}
func (m *CPartnerApps_FinishUploadKVSign_Response) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_FinishUploadKVSign_Response) ProtoMessage()    {}
func (*CPartnerApps_FinishUploadKVSign_Response) Descriptor() ([]byte, []int) {
	return partnerapps_fileDescriptor0, []int{3}
}

func (m *CPartnerApps_FinishUploadKVSign_Response) GetSignedInstallscript() string {
	if m != nil && m.SignedInstallscript != nil {
		return *m.SignedInstallscript
	}
	return ""
}

type CPartnerApps_FinishUploadLegacyDRM_Request struct {
	UploadToken      *uint64 `protobuf:"varint,1,opt,name=upload_token" json:"upload_token,omitempty"`
	RoutingId        *uint64 `protobuf:"varint,2,opt,name=routing_id" json:"routing_id,omitempty"`
	AppId            *uint32 `protobuf:"varint,3,opt,name=app_id" json:"app_id,omitempty"`
	Flags            *uint32 `protobuf:"varint,4,opt,name=flags" json:"flags,omitempty"`
	ToolName         *string `protobuf:"bytes,5,opt,name=tool_name" json:"tool_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) Reset() {
	*m = CPartnerApps_FinishUploadLegacyDRM_Request{}
}
func (m *CPartnerApps_FinishUploadLegacyDRM_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CPartnerApps_FinishUploadLegacyDRM_Request) ProtoMessage() {}
func (*CPartnerApps_FinishUploadLegacyDRM_Request) Descriptor() ([]byte, []int) {
	return partnerapps_fileDescriptor0, []int{4}
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) GetUploadToken() uint64 {
	if m != nil && m.UploadToken != nil {
		return *m.UploadToken
	}
	return 0
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) GetRoutingId() uint64 {
	if m != nil && m.RoutingId != nil {
		return *m.RoutingId
	}
	return 0
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Request) GetToolName() string {
	if m != nil && m.ToolName != nil {
		return *m.ToolName
	}
	return ""
}

type CPartnerApps_FinishUploadLegacyDRM_Response struct {
	FileId           *string `protobuf:"bytes,1,opt,name=file_id" json:"file_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Response) Reset() {
	*m = CPartnerApps_FinishUploadLegacyDRM_Response{}
}
func (m *CPartnerApps_FinishUploadLegacyDRM_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CPartnerApps_FinishUploadLegacyDRM_Response) ProtoMessage() {}
func (*CPartnerApps_FinishUploadLegacyDRM_Response) Descriptor() ([]byte, []int) {
	return partnerapps_fileDescriptor0, []int{5}
}

func (m *CPartnerApps_FinishUploadLegacyDRM_Response) GetFileId() string {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return ""
}

type CPartnerApps_FinishUpload_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPartnerApps_FinishUpload_Response) Reset()         { *m = CPartnerApps_FinishUpload_Response{} }
func (m *CPartnerApps_FinishUpload_Response) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_FinishUpload_Response) ProtoMessage()    {}
func (*CPartnerApps_FinishUpload_Response) Descriptor() ([]byte, []int) {
	return partnerapps_fileDescriptor0, []int{6}
}

type CPartnerApps_FindDRMUploads_Request struct {
	AppId            *int32 `protobuf:"varint,1,opt,name=app_id" json:"app_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPartnerApps_FindDRMUploads_Request) Reset()         { *m = CPartnerApps_FindDRMUploads_Request{} }
func (m *CPartnerApps_FindDRMUploads_Request) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_FindDRMUploads_Request) ProtoMessage()    {}
func (*CPartnerApps_FindDRMUploads_Request) Descriptor() ([]byte, []int) {
	return partnerapps_fileDescriptor0, []int{7}
}

func (m *CPartnerApps_FindDRMUploads_Request) GetAppId() int32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

type CPartnerApps_ExistingDRMUpload struct {
	FileId           *string `protobuf:"bytes,1,opt,name=file_id" json:"file_id,omitempty"`
	AppId            *uint32 `protobuf:"varint,2,opt,name=app_id" json:"app_id,omitempty"`
	ActorId          *int32  `protobuf:"varint,3,opt,name=actor_id" json:"actor_id,omitempty"`
	SuppliedName     *string `protobuf:"bytes,5,opt,name=supplied_name" json:"supplied_name,omitempty"`
	Flags            *uint32 `protobuf:"varint,6,opt,name=flags" json:"flags,omitempty"`
	ModType          *string `protobuf:"bytes,7,opt,name=mod_type" json:"mod_type,omitempty"`
	Timestamp        *uint32 `protobuf:"fixed32,8,opt,name=timestamp" json:"timestamp,omitempty"`
	OrigFileId       *string `protobuf:"bytes,9,opt,name=orig_file_id" json:"orig_file_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_ExistingDRMUpload) Reset()                    { *m = CPartnerApps_ExistingDRMUpload{} }
func (m *CPartnerApps_ExistingDRMUpload) String() string            { return proto.CompactTextString(m) }
func (*CPartnerApps_ExistingDRMUpload) ProtoMessage()               {}
func (*CPartnerApps_ExistingDRMUpload) Descriptor() ([]byte, []int) { return partnerapps_fileDescriptor0, []int{8} }

func (m *CPartnerApps_ExistingDRMUpload) GetFileId() string {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return ""
}

func (m *CPartnerApps_ExistingDRMUpload) GetAppId() uint32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func (m *CPartnerApps_ExistingDRMUpload) GetActorId() int32 {
	if m != nil && m.ActorId != nil {
		return *m.ActorId
	}
	return 0
}

func (m *CPartnerApps_ExistingDRMUpload) GetSuppliedName() string {
	if m != nil && m.SuppliedName != nil {
		return *m.SuppliedName
	}
	return ""
}

func (m *CPartnerApps_ExistingDRMUpload) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *CPartnerApps_ExistingDRMUpload) GetModType() string {
	if m != nil && m.ModType != nil {
		return *m.ModType
	}
	return ""
}

func (m *CPartnerApps_ExistingDRMUpload) GetTimestamp() uint32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *CPartnerApps_ExistingDRMUpload) GetOrigFileId() string {
	if m != nil && m.OrigFileId != nil {
		return *m.OrigFileId
	}
	return ""
}

type CPartnerApps_FindDRMUploads_Response struct {
	Uploads          []*CPartnerApps_ExistingDRMUpload `protobuf:"bytes,1,rep,name=uploads" json:"uploads,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *CPartnerApps_FindDRMUploads_Response) Reset()         { *m = CPartnerApps_FindDRMUploads_Response{} }
func (m *CPartnerApps_FindDRMUploads_Response) String() string { return proto.CompactTextString(m) }
func (*CPartnerApps_FindDRMUploads_Response) ProtoMessage()    {}
func (*CPartnerApps_FindDRMUploads_Response) Descriptor() ([]byte, []int) {
	return partnerapps_fileDescriptor0, []int{9}
}

func (m *CPartnerApps_FindDRMUploads_Response) GetUploads() []*CPartnerApps_ExistingDRMUpload {
	if m != nil {
		return m.Uploads
	}
	return nil
}

type CPartnerApps_Download_Request struct {
	FileId           *string `protobuf:"bytes,1,opt,name=file_id" json:"file_id,omitempty"`
	AppId            *int32  `protobuf:"varint,2,opt,name=app_id" json:"app_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_Download_Request) Reset()                    { *m = CPartnerApps_Download_Request{} }
func (m *CPartnerApps_Download_Request) String() string            { return proto.CompactTextString(m) }
func (*CPartnerApps_Download_Request) ProtoMessage()               {}
func (*CPartnerApps_Download_Request) Descriptor() ([]byte, []int) { return partnerapps_fileDescriptor0, []int{10} }

func (m *CPartnerApps_Download_Request) GetFileId() string {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return ""
}

func (m *CPartnerApps_Download_Request) GetAppId() int32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

type CPartnerApps_Download_Response struct {
	DownloadUrl      *string `protobuf:"bytes,1,opt,name=download_url" json:"download_url,omitempty"`
	AppId            *int32  `protobuf:"varint,2,opt,name=app_id" json:"app_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPartnerApps_Download_Response) Reset()                    { *m = CPartnerApps_Download_Response{} }
func (m *CPartnerApps_Download_Response) String() string            { return proto.CompactTextString(m) }
func (*CPartnerApps_Download_Response) ProtoMessage()               {}
func (*CPartnerApps_Download_Response) Descriptor() ([]byte, []int) { return partnerapps_fileDescriptor0, []int{11} }

func (m *CPartnerApps_Download_Response) GetDownloadUrl() string {
	if m != nil && m.DownloadUrl != nil {
		return *m.DownloadUrl
	}
	return ""
}

func (m *CPartnerApps_Download_Response) GetAppId() int32 {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return 0
}

func init() {
	proto.RegisterType((*CPartnerApps_RequestUploadToken_Request)(nil), "CPartnerApps_RequestUploadToken_Request")
	proto.RegisterType((*CPartnerApps_RequestUploadToken_Response)(nil), "CPartnerApps_RequestUploadToken_Response")
	proto.RegisterType((*CPartnerApps_FinishUpload_Request)(nil), "CPartnerApps_FinishUpload_Request")
	proto.RegisterType((*CPartnerApps_FinishUploadKVSign_Response)(nil), "CPartnerApps_FinishUploadKVSign_Response")
	proto.RegisterType((*CPartnerApps_FinishUploadLegacyDRM_Request)(nil), "CPartnerApps_FinishUploadLegacyDRM_Request")
	proto.RegisterType((*CPartnerApps_FinishUploadLegacyDRM_Response)(nil), "CPartnerApps_FinishUploadLegacyDRM_Response")
	proto.RegisterType((*CPartnerApps_FinishUpload_Response)(nil), "CPartnerApps_FinishUpload_Response")
	proto.RegisterType((*CPartnerApps_FindDRMUploads_Request)(nil), "CPartnerApps_FindDRMUploads_Request")
	proto.RegisterType((*CPartnerApps_ExistingDRMUpload)(nil), "CPartnerApps_ExistingDRMUpload")
	proto.RegisterType((*CPartnerApps_FindDRMUploads_Response)(nil), "CPartnerApps_FindDRMUploads_Response")
	proto.RegisterType((*CPartnerApps_Download_Request)(nil), "CPartnerApps_Download_Request")
	proto.RegisterType((*CPartnerApps_Download_Response)(nil), "CPartnerApps_Download_Response")
}

var partnerapps_fileDescriptor0 = []byte{
	// 809 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x96, 0x4f, 0x6f, 0xd3, 0x4a,
	0x14, 0xc5, 0xe5, 0xb6, 0x69, 0x9a, 0xe9, 0x4b, 0xdf, 0x7b, 0x56, 0x2b, 0x45, 0x11, 0xb4, 0xc6,
	0x2d, 0x22, 0xb4, 0xd5, 0x50, 0x15, 0x21, 0x36, 0x08, 0x95, 0xa6, 0x7f, 0x10, 0x05, 0x84, 0x5a,
	0x40, 0x6c, 0x50, 0x34, 0x75, 0x26, 0xe9, 0x08, 0xdb, 0x63, 0x3c, 0x63, 0x68, 0x77, 0x88, 0x15,
	0x1b, 0x3e, 0x00, 0x7b, 0x76, 0x20, 0x24, 0x24, 0xfa, 0xfd, 0xb8, 0x33, 0xb6, 0x13, 0x3b, 0x69,
	0x53, 0x23, 0x84, 0x58, 0x66, 0x7c, 0x7d, 0xee, 0x6f, 0xce, 0xdc, 0x39, 0x31, 0x5a, 0x11, 0x92,
	0x12, 0xcf, 0xa3, 0x42, 0x90, 0x2e, 0x15, 0xad, 0x80, 0x84, 0xd2, 0xa7, 0x21, 0x09, 0x02, 0x81,
	0xf5, 0x13, 0xc7, 0x65, 0xd4, 0x97, 0x38, 0x08, 0xb9, 0xe4, 0xf5, 0xd5, 0x7c, 0x71, 0xe4, 0xb3,
	0x0e, 0xa3, 0xed, 0xd6, 0x21, 0x11, 0x74, 0xb8, 0xda, 0x7e, 0x80, 0xae, 0x35, 0x9f, 0xc4, 0x7a,
	0xf7, 0x40, 0xaf, 0xb5, 0x4f, 0x5f, 0x47, 0x54, 0xc8, 0x67, 0x81, 0xcb, 0x49, 0xfb, 0x29, 0x7f,
	0x45, 0xfd, 0x74, 0xc9, 0xfc, 0x0f, 0x4d, 0x75, 0x98, 0x4b, 0x7d, 0xe2, 0xd1, 0x9a, 0x61, 0x19,
	0x8d, 0x8a, 0x59, 0x45, 0x25, 0x80, 0x60, 0xed, 0xda, 0x18, 0xfc, 0xac, 0xda, 0x1d, 0xd4, 0xb8,
	0x58, 0x4b, 0x04, 0xdc, 0x17, 0xd4, 0x9c, 0x45, 0xff, 0x44, 0x7a, 0xbd, 0x25, 0xd5, 0x03, 0x2d,
	0x38, 0xa1, 0x5a, 0xb8, 0xdc, 0x21, 0x92, 0x71, 0x5f, 0x6b, 0x56, 0x4c, 0x13, 0xa1, 0x90, 0x47,
	0x92, 0xf9, 0xdd, 0x16, 0xf4, 0x19, 0x57, 0x55, 0xf6, 0x4b, 0x74, 0x25, 0xd7, 0x67, 0x87, 0xf9,
	0x4c, 0x1c, 0xc5, 0x6d, 0x7a, 0xb4, 0x67, 0x37, 0xc8, 0xcb, 0x8d, 0xe9, 0xb5, 0x19, 0x34, 0x09,
	0xbb, 0x48, 0xe5, 0xab, 0xf6, 0xfd, 0x81, 0x6d, 0x64, 0xe5, 0xf7, 0x9e, 0x1f, 0xb0, 0x6e, 0x66,
	0x1b, 0x97, 0xd0, 0xac, 0x80, 0x05, 0xf0, 0x97, 0xf9, 0x42, 0x12, 0xd7, 0x15, 0x4e, 0xc8, 0x02,
	0x19, 0xfb, 0x63, 0x7f, 0x30, 0xd0, 0xf2, 0xb9, 0x52, 0x0f, 0x69, 0x97, 0x38, 0x27, 0x5b, 0xfb,
	0x8f, 0x7e, 0x1f, 0x59, 0x1d, 0x44, 0xc7, 0x25, 0x5d, 0x51, 0x9b, 0xd0, 0x3f, 0xff, 0x47, 0x15,
	0xc9, 0xb9, 0xdb, 0xd2, 0x47, 0x55, 0xd2, 0x28, 0x77, 0xd1, 0x4a, 0x21, 0x92, 0x64, 0x5f, 0xff,
	0xa2, 0xb2, 0x3a, 0x6b, 0xd5, 0x21, 0xde, 0xca, 0x12, 0xb2, 0x47, 0x79, 0x1e, 0xbf, 0x66, 0xdf,
	0x42, 0x8b, 0x83, 0x55, 0x6d, 0xd0, 0x8d, 0xcb, 0x7a, 0x03, 0x91, 0xc1, 0x57, 0xe2, 0x25, 0xfb,
	0x87, 0x81, 0xe6, 0x73, 0xef, 0x6d, 0x1f, 0x33, 0xa1, 0x76, 0xdc, 0x7b, 0x77, 0x08, 0x28, 0xa3,
	0xa1, 0x87, 0x4f, 0x8d, 0x0e, 0x71, 0x24, 0x0f, 0x53, 0x53, 0x4a, 0xe6, 0x1c, 0xaa, 0x8a, 0x28,
	0x08, 0x5c, 0x35, 0xfd, 0x7d, 0x27, 0xfa, 0x5e, 0x4d, 0xa6, 0xef, 0x79, 0x1c, 0x1c, 0x3f, 0x09,
	0x68, 0xad, 0xac, 0x0b, 0x94, 0x7b, 0x0c, 0xee, 0x8f, 0x24, 0x5e, 0x50, 0x9b, 0x82, 0xa5, 0xb2,
	0x3a, 0x19, 0x1e, 0xb2, 0x6e, 0x2b, 0x45, 0xa8, 0x68, 0x4f, 0x5e, 0xa0, 0xa5, 0xd1, 0xbb, 0x4d,
	0xcc, 0x5c, 0x43, 0xe5, 0xf8, 0x5c, 0x05, 0xb0, 0x8f, 0x37, 0xa6, 0xd7, 0x17, 0xf0, 0xe8, 0xdd,
	0xda, 0x1b, 0xe8, 0x72, 0xae, 0x62, 0x8b, 0xbf, 0xf5, 0x73, 0xd3, 0x7d, 0x81, 0x1d, 0x25, 0x7b,
	0x67, 0xc0, 0xd1, 0x8c, 0x42, 0xff, 0x06, 0xb6, 0xd3, 0xc5, 0x28, 0x74, 0xcf, 0xd6, 0x59, 0xff,
	0x8a, 0xd0, 0x74, 0x46, 0xc7, 0xfc, 0x6e, 0xa0, 0x5a, 0x02, 0x11, 0xdf, 0x85, 0xcc, 0xed, 0x36,
	0x1b, 0xb8, 0x60, 0x96, 0xd4, 0xaf, 0xe3, 0xa2, 0x49, 0x61, 0x6f, 0xbc, 0x3f, 0xad, 0xdd, 0x49,
	0x0a, 0xac, 0xd8, 0x47, 0x4b, 0xdf, 0x0f, 0xab, 0xc3, 0x43, 0x2b, 0x77, 0xed, 0x2c, 0xe5, 0x49,
	0x5a, 0x73, 0xc3, 0x52, 0x37, 0x13, 0x6c, 0x35, 0xbf, 0x19, 0x68, 0x2e, 0x11, 0xe8, 0x59, 0xfc,
	0xd7, 0x80, 0xe9, 0x31, 0x75, 0x22, 0x49, 0x0e, 0x01, 0x34, 0x4f, 0x0b, 0x79, 0xec, 0x40, 0x68,
	0x2b, 0xe0, 0xd3, 0x3e, 0x70, 0x73, 0x7b, 0xf7, 0x8f, 0x03, 0xef, 0x02, 0x70, 0xf3, 0x5c, 0x60,
	0x27, 0x12, 0x92, 0x7b, 0x85, 0xb8, 0xbf, 0x18, 0xc8, 0x1c, 0x4e, 0x4b, 0xd3, 0xc6, 0x17, 0xc6,
	0xf5, 0x20, 0xee, 0x88, 0xcc, 0xb5, 0x77, 0x00, 0x77, 0xb3, 0xc9, 0x3d, 0x8f, 0x49, 0xcb, 0xa3,
	0xf2, 0x88, 0x2b, 0x5a, 0x7d, 0xd8, 0x16, 0xf1, 0x07, 0x06, 0x82, 0x74, 0x24, 0x0d, 0xd5, 0x72,
	0x82, 0xcb, 0x84, 0xe5, 0x70, 0x2f, 0x70, 0xa9, 0xa4, 0xe6, 0x67, 0x70, 0x39, 0xdb, 0xa7, 0x1f,
	0x36, 0x2b, 0xb8, 0x78, 0x6a, 0xd7, 0x57, 0xf1, 0x2f, 0x04, 0xab, 0xbd, 0x06, 0xf0, 0xab, 0x79,
	0xf8, 0xb3, 0xa6, 0x22, 0x63, 0xea, 0xa7, 0x01, 0xcc, 0xde, 0x44, 0x14, 0xf2, 0x75, 0x11, 0x17,
	0x88, 0xed, 0xdb, 0x00, 0x75, 0x73, 0x18, 0xea, 0x9c, 0x93, 0xcf, 0xb0, 0x7d, 0x34, 0xd0, 0x4c,
	0x3e, 0xf5, 0xcc, 0x25, 0x5c, 0xe0, 0x1f, 0xa0, 0x7e, 0x15, 0x17, 0x49, 0x4e, 0x1b, 0x03, 0xd8,
	0xb2, 0x7a, 0x28, 0xac, 0xc7, 0x96, 0xc7, 0x61, 0x3c, 0x43, 0xea, 0xc0, 0xc7, 0x0b, 0x30, 0x30,
	0x00, 0x4c, 0x52, 0xd5, 0x3a, 0x3c, 0xb1, 0xf4, 0x77, 0x88, 0xb9, 0x87, 0xa6, 0xd2, 0xa0, 0x33,
	0xe7, 0xf1, 0xc8, 0x08, 0xad, 0x2f, 0xe0, 0xd1, 0x01, 0x59, 0x5f, 0x87, 0xe6, 0xf8, 0x80, 0x86,
	0x6f, 0x98, 0x43, 0x13, 0x5b, 0x84, 0xf6, 0x05, 0x7a, 0x59, 0x1e, 0xf1, 0xe1, 0xf3, 0xca, 0x53,
	0x2c, 0xd0, 0x3d, 0xf9, 0x22, 0x13, 0x9b, 0xe3, 0xef, 0x0c, 0xe3, 0x67, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xec, 0x20, 0xc0, 0x4b, 0xaf, 0x09, 0x00, 0x00,
}
