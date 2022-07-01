package authz

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AccessRule struct {
	Path string     `json:"path"`
	Verb AccessVerb `json:"verb"`
	Kind string     `json:"kind"` // * or single kind
	Who  string     `json:"who"`  // group or userid
}

type AccessVerb int32

const (
	// Each permission implies the previous
	AccessUnknown AccessVerb = 0
	AccessNone    AccessVerb = 1    // block / deny
	AccessRead    AccessVerb = 100  // read
	AccessExec    AccessVerb = 200  // query, execute, etc
	AccessWrite   AccessVerb = 300  // read+write
	AccessManage  AccessVerb = 400  // read+write+delete
	AccessAdmin   AccessVerb = 1000 // and change permissions
)

func (p AccessVerb) String() string {
	switch p {
	case AccessUnknown:
		return "UNKNOWN"
	case AccessNone:
		return "NONE"
	case AccessRead:
		return "READ"
	case AccessExec:
		return "EXE"
	case AccessWrite:
		return "WRITE"
	case AccessManage:
		return "MANAGE"
	case AccessAdmin:
		return "ADMIN"
	}
	return fmt.Sprintf("%d", p)
}

// MarshalJSON marshals the enum as a quoted json string
func (p AccessVerb) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(p.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmarshals a quoted json string to the enum value
func (p *AccessVerb) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	f := AccessVerbFrom(j)
	*p = f
	return nil
}

// AccessVerbFor returns a concrete type for a given interface or unknown if not known
func AccessVerbFrom(str string) AccessVerb {
	switch str {
	case "NONE":
		return AccessNone
	case "READ":
		return AccessRead
	case "EXE":
		return AccessExec
	case "WRITE":
		return AccessWrite
	case "MANAGE":
		return AccessManage
	case "ADMIN":
		return AccessAdmin
	}
	return AccessUnknown
}
