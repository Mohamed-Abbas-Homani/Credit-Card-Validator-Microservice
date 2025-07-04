// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: pkg/proto/cardvalidator.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ValidateCardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CardNumber    string                 `protobuf:"bytes,1,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateCardRequest) Reset() {
	*x = ValidateCardRequest{}
	mi := &file_pkg_proto_cardvalidator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCardRequest) ProtoMessage() {}

func (x *ValidateCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cardvalidator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCardRequest.ProtoReflect.Descriptor instead.
func (*ValidateCardRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cardvalidator_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateCardRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

type ValidateCardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Valid         bool                   `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	CardType      string                 `protobuf:"bytes,2,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	CardNumber    string                 `protobuf:"bytes,3,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	Scheme        string                 `protobuf:"bytes,4,opt,name=scheme,proto3" json:"scheme,omitempty"`
	CardBrand     string                 `protobuf:"bytes,5,opt,name=card_brand,json=cardBrand,proto3" json:"card_brand,omitempty"`
	CardKind      string                 `protobuf:"bytes,6,opt,name=card_kind,json=cardKind,proto3" json:"card_kind,omitempty"`
	Country       *Country               `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`
	Bank          *Bank                  `protobuf:"bytes,8,opt,name=bank,proto3" json:"bank,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateCardResponse) Reset() {
	*x = ValidateCardResponse{}
	mi := &file_pkg_proto_cardvalidator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCardResponse) ProtoMessage() {}

func (x *ValidateCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cardvalidator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCardResponse.ProtoReflect.Descriptor instead.
func (*ValidateCardResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cardvalidator_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateCardResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ValidateCardResponse) GetCardType() string {
	if x != nil {
		return x.CardType
	}
	return ""
}

func (x *ValidateCardResponse) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *ValidateCardResponse) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *ValidateCardResponse) GetCardBrand() string {
	if x != nil {
		return x.CardBrand
	}
	return ""
}

func (x *ValidateCardResponse) GetCardKind() string {
	if x != nil {
		return x.CardKind
	}
	return ""
}

func (x *ValidateCardResponse) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

func (x *ValidateCardResponse) GetBank() *Bank {
	if x != nil {
		return x.Bank
	}
	return nil
}

type Country struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Alpha2        string                 `protobuf:"bytes,2,opt,name=alpha2,proto3" json:"alpha2,omitempty"`
	Currency      string                 `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Emoji         string                 `protobuf:"bytes,4,opt,name=emoji,proto3" json:"emoji,omitempty"`
	Latitude      int32                  `protobuf:"varint,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     int32                  `protobuf:"varint,6,opt,name=longitude,proto3" json:"longitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Country) Reset() {
	*x = Country{}
	mi := &file_pkg_proto_cardvalidator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cardvalidator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cardvalidator_proto_rawDescGZIP(), []int{2}
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Country) GetAlpha2() string {
	if x != nil {
		return x.Alpha2
	}
	return ""
}

func (x *Country) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Country) GetEmoji() string {
	if x != nil {
		return x.Emoji
	}
	return ""
}

func (x *Country) GetLatitude() int32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Country) GetLongitude() int32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Bank struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bank) Reset() {
	*x = Bank{}
	mi := &file_pkg_proto_cardvalidator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bank) ProtoMessage() {}

func (x *Bank) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_cardvalidator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bank.ProtoReflect.Descriptor instead.
func (*Bank) Descriptor() ([]byte, []int) {
	return file_pkg_proto_cardvalidator_proto_rawDescGZIP(), []int{3}
}

func (x *Bank) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bank) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Bank) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_pkg_proto_cardvalidator_proto protoreflect.FileDescriptor

const file_pkg_proto_cardvalidator_proto_rawDesc = "" +
	"\n" +
	"\x1dpkg/proto/cardvalidator.proto\x12\rcardvalidator\"6\n" +
	"\x13ValidateCardRequest\x12\x1f\n" +
	"\vcard_number\x18\x01 \x01(\tR\n" +
	"cardNumber\"\x99\x02\n" +
	"\x14ValidateCardResponse\x12\x14\n" +
	"\x05valid\x18\x01 \x01(\bR\x05valid\x12\x1b\n" +
	"\tcard_type\x18\x02 \x01(\tR\bcardType\x12\x1f\n" +
	"\vcard_number\x18\x03 \x01(\tR\n" +
	"cardNumber\x12\x16\n" +
	"\x06scheme\x18\x04 \x01(\tR\x06scheme\x12\x1d\n" +
	"\n" +
	"card_brand\x18\x05 \x01(\tR\tcardBrand\x12\x1b\n" +
	"\tcard_kind\x18\x06 \x01(\tR\bcardKind\x120\n" +
	"\acountry\x18\a \x01(\v2\x16.cardvalidator.CountryR\acountry\x12'\n" +
	"\x04bank\x18\b \x01(\v2\x13.cardvalidator.BankR\x04bank\"\xa1\x01\n" +
	"\aCountry\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x16\n" +
	"\x06alpha2\x18\x02 \x01(\tR\x06alpha2\x12\x1a\n" +
	"\bcurrency\x18\x03 \x01(\tR\bcurrency\x12\x14\n" +
	"\x05emoji\x18\x04 \x01(\tR\x05emoji\x12\x1a\n" +
	"\blatitude\x18\x05 \x01(\x05R\blatitude\x12\x1c\n" +
	"\tlongitude\x18\x06 \x01(\x05R\tlongitude\"B\n" +
	"\x04Bank\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x10\n" +
	"\x03url\x18\x02 \x01(\tR\x03url\x12\x14\n" +
	"\x05phone\x18\x03 \x01(\tR\x05phone2h\n" +
	"\rCardValidator\x12W\n" +
	"\fValidateCard\x12\".cardvalidator.ValidateCardRequest\x1a#.cardvalidator.ValidateCardResponseB!Z\x1fcredit-card-validator/pkg/protob\x06proto3"

var (
	file_pkg_proto_cardvalidator_proto_rawDescOnce sync.Once
	file_pkg_proto_cardvalidator_proto_rawDescData []byte
)

func file_pkg_proto_cardvalidator_proto_rawDescGZIP() []byte {
	file_pkg_proto_cardvalidator_proto_rawDescOnce.Do(func() {
		file_pkg_proto_cardvalidator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_proto_cardvalidator_proto_rawDesc), len(file_pkg_proto_cardvalidator_proto_rawDesc)))
	})
	return file_pkg_proto_cardvalidator_proto_rawDescData
}

var file_pkg_proto_cardvalidator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_proto_cardvalidator_proto_goTypes = []any{
	(*ValidateCardRequest)(nil),  // 0: cardvalidator.ValidateCardRequest
	(*ValidateCardResponse)(nil), // 1: cardvalidator.ValidateCardResponse
	(*Country)(nil),              // 2: cardvalidator.Country
	(*Bank)(nil),                 // 3: cardvalidator.Bank
}
var file_pkg_proto_cardvalidator_proto_depIdxs = []int32{
	2, // 0: cardvalidator.ValidateCardResponse.country:type_name -> cardvalidator.Country
	3, // 1: cardvalidator.ValidateCardResponse.bank:type_name -> cardvalidator.Bank
	0, // 2: cardvalidator.CardValidator.ValidateCard:input_type -> cardvalidator.ValidateCardRequest
	1, // 3: cardvalidator.CardValidator.ValidateCard:output_type -> cardvalidator.ValidateCardResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_cardvalidator_proto_init() }
func file_pkg_proto_cardvalidator_proto_init() {
	if File_pkg_proto_cardvalidator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_proto_cardvalidator_proto_rawDesc), len(file_pkg_proto_cardvalidator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_cardvalidator_proto_goTypes,
		DependencyIndexes: file_pkg_proto_cardvalidator_proto_depIdxs,
		MessageInfos:      file_pkg_proto_cardvalidator_proto_msgTypes,
	}.Build()
	File_pkg_proto_cardvalidator_proto = out.File
	file_pkg_proto_cardvalidator_proto_goTypes = nil
	file_pkg_proto_cardvalidator_proto_depIdxs = nil
}
