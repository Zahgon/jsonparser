package benchmark

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *CBAvatar) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mj *CBAvatar) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffj_t_CBAvatarbase = iota
	ffj_t_CBAvatarno_such_key

	ffj_t_CBAvatar_Url
)

var ffj_key_CBAvatar_Url = []byte("Url")

func (uj *CBAvatar) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (uj *CBAvatar) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}

func (mj *CBGithub) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mj *CBGithub) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffj_t_CBGithubbase = iota
	ffj_t_CBGithubno_such_key

	ffj_t_CBGithub_Followers
)

var ffj_key_CBGithub_Followers = []byte("Followers")

func (uj *CBGithub) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (uj *CBGithub) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}

func (mj *CBGravatar) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mj *CBGravatar) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffj_t_CBGravatarbase = iota
	ffj_t_CBGravatarno_such_key

	ffj_t_CBGravatar_Avatars
)

var ffj_key_CBGravatar_Avatars = []byte("Avatars")

func (uj *CBGravatar) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (uj *CBGravatar) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}

func (mj *CBName) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mj *CBName) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffj_t_CBNamebase = iota
	ffj_t_CBNameno_such_key

	ffj_t_CBName_FullName
)

var ffj_key_CBName_FullName = []byte("FullName")

func (uj *CBName) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (uj *CBName) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}

func (mj *CBPerson) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mj *CBPerson) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffj_t_CBPersonbase = iota
	ffj_t_CBPersonno_such_key

	ffj_t_CBPerson_Name

	ffj_t_CBPerson_Github

	ffj_t_CBPerson_Gravatar
)

var ffj_key_CBPerson_Name = []byte("Name")

var ffj_key_CBPerson_Github = []byte("Github")

var ffj_key_CBPerson_Gravatar = []byte("Gravatar")

func (uj *CBPerson) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (uj *CBPerson) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}

func (mj *DSTopic) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mj *DSTopic) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffj_t_DSTopicbase = iota
	ffj_t_DSTopicno_such_key

	ffj_t_DSTopic_Id

	ffj_t_DSTopic_Slug
)

var ffj_key_DSTopic_Id = []byte("Id")

var ffj_key_DSTopic_Slug = []byte("Slug")

func (uj *DSTopic) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (uj *DSTopic) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}

func (mj *DSTopicsList) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mj *DSTopicsList) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffj_t_DSTopicsListbase = iota
	ffj_t_DSTopicsListno_such_key

	ffj_t_DSTopicsList_Topics

	ffj_t_DSTopicsList_MoreTopicsUrl
)

var ffj_key_DSTopicsList_Topics = []byte("Topics")

var ffj_key_DSTopicsList_MoreTopicsUrl = []byte("MoreTopicsUrl")

func (uj *DSTopicsList) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (uj *DSTopicsList) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}

func (mj *DSUser) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mj *DSUser) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffj_t_DSUserbase = iota
	ffj_t_DSUserno_such_key

	ffj_t_DSUser_Username
)

var ffj_key_DSUser_Username = []byte("Username")

func (uj *DSUser) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (uj *DSUser) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}

func (mj *LargePayload) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mj *LargePayload) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffj_t_LargePayloadbase = iota
	ffj_t_LargePayloadno_such_key

	ffj_t_LargePayload_Users

	ffj_t_LargePayload_Topics
)

var ffj_key_LargePayload_Users = []byte("Users")

var ffj_key_LargePayload_Topics = []byte("Topics")

func (uj *LargePayload) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (uj *LargePayload) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}

func (mj *MediumPayload) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mj *MediumPayload) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffj_t_MediumPayloadbase = iota
	ffj_t_MediumPayloadno_such_key

	ffj_t_MediumPayload_Person

	ffj_t_MediumPayload_Company
)

var ffj_key_MediumPayload_Person = []byte("Person")

var ffj_key_MediumPayload_Company = []byte("Company")

func (uj *MediumPayload) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (uj *MediumPayload) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}

func (mj *SmallPayload) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mj *SmallPayload) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

const (
	ffj_t_SmallPayloadbase = iota
	ffj_t_SmallPayloadno_such_key

	ffj_t_SmallPayload_St

	ffj_t_SmallPayload_Sid

	ffj_t_SmallPayload_Tt

	ffj_t_SmallPayload_Gr

	ffj_t_SmallPayload_Uuid

	ffj_t_SmallPayload_Ip

	ffj_t_SmallPayload_Ua

	ffj_t_SmallPayload_Tz

	ffj_t_SmallPayload_V
)

var ffj_key_SmallPayload_St = []byte("St")

var ffj_key_SmallPayload_Sid = []byte("Sid")

var ffj_key_SmallPayload_Tt = []byte("Tt")

var ffj_key_SmallPayload_Gr = []byte("Gr")

var ffj_key_SmallPayload_Uuid = []byte("Uuid")

var ffj_key_SmallPayload_Ip = []byte("Ip")

var ffj_key_SmallPayload_Ua = []byte("Ua")

var ffj_key_SmallPayload_Tz = []byte("Tz")

var ffj_key_SmallPayload_V = []byte("V")

func (uj *SmallPayload) UnmarshalJSON(input []byte) error { _ = "STUB: not implemented"; return nil }

func (uj *SmallPayload) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	_ = "STUB: not implemented"
	return nil
}
