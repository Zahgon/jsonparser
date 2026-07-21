package benchmark

import (
	"errors"
	"runtime"
	"strconv"

	codec1978 "github.com/ugorji/go/codec"
)

const (
	codecSelferCcUTF86617 = 1
	codecSelferCcRAW6617  = 255

	codecSelferValueTypeArray6617  = 10
	codecSelferValueTypeMap6617    = 9
	codecSelferValueTypeString6617 = 6
	codecSelferValueTypeInt6617    = 2
	codecSelferValueTypeUint6617   = 3
	codecSelferValueTypeFloat6617  = 4
	codecSelferBitsize6617         = uint8(32 << (^uint(0) >> 63))
)

var (
	errCodecSelferOnlyMapOrArrayEncodeToStruct6617 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer6617 struct{}

func init() {
	if codec1978.GenVersion != 10 {
		_, file, _, _ := runtime.Caller(0)
		panic("codecgen version mismatch: current: 10, need " + strconv.FormatInt(int64(codec1978.GenVersion), 10) + ". Re-generate file: " + file)
	}
	if false {
		var _ byte = 0
	}
}

func (x *SmallPayload) CodecEncodeSelf(e *codec1978.Encoder) { _ = "STUB: not implemented"; return }

func (x *SmallPayload) CodecDecodeSelf(d *codec1978.Decoder) { _ = "STUB: not implemented"; return }

func (x *SmallPayload) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *SmallPayload) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *CBAvatar) CodecEncodeSelf(e *codec1978.Encoder) { _ = "STUB: not implemented"; return }

func (x *CBAvatar) CodecDecodeSelf(d *codec1978.Decoder) { _ = "STUB: not implemented"; return }

func (x *CBAvatar) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *CBAvatar) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *CBGravatar) CodecEncodeSelf(e *codec1978.Encoder) { _ = "STUB: not implemented"; return }

func (x *CBGravatar) CodecDecodeSelf(d *codec1978.Decoder) { _ = "STUB: not implemented"; return }

func (x *CBGravatar) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *CBGravatar) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *CBGithub) CodecEncodeSelf(e *codec1978.Encoder) { _ = "STUB: not implemented"; return }

func (x *CBGithub) CodecDecodeSelf(d *codec1978.Decoder) { _ = "STUB: not implemented"; return }

func (x *CBGithub) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *CBGithub) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *CBName) CodecEncodeSelf(e *codec1978.Encoder) { _ = "STUB: not implemented"; return }

func (x *CBName) CodecDecodeSelf(d *codec1978.Decoder) { _ = "STUB: not implemented"; return }

func (x *CBName) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *CBName) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *CBPerson) CodecEncodeSelf(e *codec1978.Encoder) { _ = "STUB: not implemented"; return }

func (x *CBPerson) CodecDecodeSelf(d *codec1978.Decoder) { _ = "STUB: not implemented"; return }

func (x *CBPerson) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *CBPerson) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *MediumPayload) CodecEncodeSelf(e *codec1978.Encoder) { _ = "STUB: not implemented"; return }

func (x *MediumPayload) CodecDecodeSelf(d *codec1978.Decoder) { _ = "STUB: not implemented"; return }

func (x *MediumPayload) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *MediumPayload) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *DSUser) CodecEncodeSelf(e *codec1978.Encoder) { _ = "STUB: not implemented"; return }

func (x *DSUser) CodecDecodeSelf(d *codec1978.Decoder) { _ = "STUB: not implemented"; return }

func (x *DSUser) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *DSUser) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *DSTopic) CodecEncodeSelf(e *codec1978.Encoder) { _ = "STUB: not implemented"; return }

func (x *DSTopic) CodecDecodeSelf(d *codec1978.Decoder) { _ = "STUB: not implemented"; return }

func (x *DSTopic) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *DSTopic) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *DSTopicsList) CodecEncodeSelf(e *codec1978.Encoder) { _ = "STUB: not implemented"; return }

func (x *DSTopicsList) CodecDecodeSelf(d *codec1978.Decoder) { _ = "STUB: not implemented"; return }

func (x *DSTopicsList) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *DSTopicsList) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *LargePayload) CodecEncodeSelf(e *codec1978.Encoder) { _ = "STUB: not implemented"; return }

func (x *LargePayload) CodecDecodeSelf(d *codec1978.Decoder) { _ = "STUB: not implemented"; return }

func (x *LargePayload) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x *LargePayload) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x codecSelfer6617) encSlicePtrtoCBAvatar(v []*CBAvatar, e *codec1978.Encoder) {
	_ = "STUB: not implemented"
	return
}

func (x codecSelfer6617) decSlicePtrtoCBAvatar(v *[]*CBAvatar, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x codecSelfer6617) encSlicePtrtoDSTopic(v []*DSTopic, e *codec1978.Encoder) {
	_ = "STUB: not implemented"
	return
}

func (x codecSelfer6617) decSlicePtrtoDSTopic(v *[]*DSTopic, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (x codecSelfer6617) encSlicePtrtoDSUser(v []*DSUser, e *codec1978.Encoder) {
	_ = "STUB: not implemented"
	return
}

func (x codecSelfer6617) decSlicePtrtoDSUser(v *[]*DSUser, d *codec1978.Decoder) {
	_ = "STUB: not implemented"
	return
}
