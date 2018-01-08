// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package request

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - Hash
//  - SizeBytes
type BazelDigest struct {
	Hash      string `thrift:"hash,1,required" json:"hash"`
	SizeBytes int64  `thrift:"sizeBytes,2,required" json:"sizeBytes"`
}

func NewBazelDigest() *BazelDigest {
	return &BazelDigest{}
}

func (p *BazelDigest) GetHash() string {
	return p.Hash
}

func (p *BazelDigest) GetSizeBytes() int64 {
	return p.SizeBytes
}
func (p *BazelDigest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetHash bool = false
	var issetSizeBytes bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetHash = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetSizeBytes = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetHash {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Hash is not set"))
	}
	if !issetSizeBytes {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field SizeBytes is not set"))
	}
	return nil
}

func (p *BazelDigest) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Hash = v
	}
	return nil
}

func (p *BazelDigest) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.SizeBytes = v
	}
	return nil
}

func (p *BazelDigest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("BazelDigest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BazelDigest) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("hash", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:hash: ", p), err)
	}
	if err := oprot.WriteString(string(p.Hash)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.hash (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:hash: ", p), err)
	}
	return err
}

func (p *BazelDigest) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sizeBytes", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:sizeBytes: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.SizeBytes)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sizeBytes (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:sizeBytes: ", p), err)
	}
	return err
}

func (p *BazelDigest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BazelDigest(%+v)", *p)
}

// Attributes:
//  - Name
//  - Value
type BazelProperty struct {
	Name  string `thrift:"name,1,required" json:"name"`
	Value string `thrift:"value,2,required" json:"value"`
}

func NewBazelProperty() *BazelProperty {
	return &BazelProperty{}
}

func (p *BazelProperty) GetName() string {
	return p.Name
}

func (p *BazelProperty) GetValue() string {
	return p.Value
}
func (p *BazelProperty) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetName bool = false
	var issetValue bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetName = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetValue = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetName {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Name is not set"))
	}
	if !issetValue {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Value is not set"))
	}
	return nil
}

func (p *BazelProperty) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *BazelProperty) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *BazelProperty) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("BazelProperty"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BazelProperty) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err)
	}
	return err
}

func (p *BazelProperty) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:value: ", p), err)
	}
	if err := oprot.WriteString(string(p.Value)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.value (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:value: ", p), err)
	}
	return err
}

func (p *BazelProperty) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BazelProperty(%+v)", *p)
}

// Attributes:
//  - CommandDigest
//  - InputDigest
//  - OutputFiles
//  - OutputDirs
//  - PlatformProperties
//  - TimeoutMs
//  - NoCache
type BazelAction struct {
	CommandDigest      *BazelDigest     `thrift:"commandDigest,1,required" json:"commandDigest"`
	InputDigest        *BazelDigest     `thrift:"inputDigest,2,required" json:"inputDigest"`
	OutputFiles        []string         `thrift:"outputFiles,3" json:"outputFiles,omitempty"`
	OutputDirs         []string         `thrift:"outputDirs,4" json:"outputDirs,omitempty"`
	PlatformProperties []*BazelProperty `thrift:"platformProperties,5" json:"platformProperties,omitempty"`
	TimeoutMs          *int64           `thrift:"timeoutMs,6" json:"timeoutMs,omitempty"`
	NoCache            *bool            `thrift:"noCache,7" json:"noCache,omitempty"`
}

func NewBazelAction() *BazelAction {
	return &BazelAction{}
}

var BazelAction_CommandDigest_DEFAULT *BazelDigest

func (p *BazelAction) GetCommandDigest() *BazelDigest {
	if !p.IsSetCommandDigest() {
		return BazelAction_CommandDigest_DEFAULT
	}
	return p.CommandDigest
}

var BazelAction_InputDigest_DEFAULT *BazelDigest

func (p *BazelAction) GetInputDigest() *BazelDigest {
	if !p.IsSetInputDigest() {
		return BazelAction_InputDigest_DEFAULT
	}
	return p.InputDigest
}

var BazelAction_OutputFiles_DEFAULT []string

func (p *BazelAction) GetOutputFiles() []string {
	return p.OutputFiles
}

var BazelAction_OutputDirs_DEFAULT []string

func (p *BazelAction) GetOutputDirs() []string {
	return p.OutputDirs
}

var BazelAction_PlatformProperties_DEFAULT []*BazelProperty

func (p *BazelAction) GetPlatformProperties() []*BazelProperty {
	return p.PlatformProperties
}

var BazelAction_TimeoutMs_DEFAULT int64

func (p *BazelAction) GetTimeoutMs() int64 {
	if !p.IsSetTimeoutMs() {
		return BazelAction_TimeoutMs_DEFAULT
	}
	return *p.TimeoutMs
}

var BazelAction_NoCache_DEFAULT bool

func (p *BazelAction) GetNoCache() bool {
	if !p.IsSetNoCache() {
		return BazelAction_NoCache_DEFAULT
	}
	return *p.NoCache
}
func (p *BazelAction) IsSetCommandDigest() bool {
	return p.CommandDigest != nil
}

func (p *BazelAction) IsSetInputDigest() bool {
	return p.InputDigest != nil
}

func (p *BazelAction) IsSetOutputFiles() bool {
	return p.OutputFiles != nil
}

func (p *BazelAction) IsSetOutputDirs() bool {
	return p.OutputDirs != nil
}

func (p *BazelAction) IsSetPlatformProperties() bool {
	return p.PlatformProperties != nil
}

func (p *BazelAction) IsSetTimeoutMs() bool {
	return p.TimeoutMs != nil
}

func (p *BazelAction) IsSetNoCache() bool {
	return p.NoCache != nil
}

func (p *BazelAction) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetCommandDigest bool = false
	var issetInputDigest bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetCommandDigest = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetInputDigest = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetCommandDigest {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field CommandDigest is not set"))
	}
	if !issetInputDigest {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field InputDigest is not set"))
	}
	return nil
}

func (p *BazelAction) readField1(iprot thrift.TProtocol) error {
	p.CommandDigest = &BazelDigest{}
	if err := p.CommandDigest.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.CommandDigest), err)
	}
	return nil
}

func (p *BazelAction) readField2(iprot thrift.TProtocol) error {
	p.InputDigest = &BazelDigest{}
	if err := p.InputDigest.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.InputDigest), err)
	}
	return nil
}

func (p *BazelAction) readField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.OutputFiles = tSlice
	for i := 0; i < size; i++ {
		var _elem0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.OutputFiles = append(p.OutputFiles, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *BazelAction) readField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.OutputDirs = tSlice
	for i := 0; i < size; i++ {
		var _elem1 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem1 = v
		}
		p.OutputDirs = append(p.OutputDirs, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *BazelAction) readField5(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*BazelProperty, 0, size)
	p.PlatformProperties = tSlice
	for i := 0; i < size; i++ {
		_elem2 := &BazelProperty{}
		if err := _elem2.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem2), err)
		}
		p.PlatformProperties = append(p.PlatformProperties, _elem2)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *BazelAction) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.TimeoutMs = &v
	}
	return nil
}

func (p *BazelAction) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.NoCache = &v
	}
	return nil
}

func (p *BazelAction) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("BazelAction"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BazelAction) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("commandDigest", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:commandDigest: ", p), err)
	}
	if err := p.CommandDigest.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.CommandDigest), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:commandDigest: ", p), err)
	}
	return err
}

func (p *BazelAction) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("inputDigest", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:inputDigest: ", p), err)
	}
	if err := p.InputDigest.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.InputDigest), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:inputDigest: ", p), err)
	}
	return err
}

func (p *BazelAction) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetOutputFiles() {
		if err := oprot.WriteFieldBegin("outputFiles", thrift.LIST, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:outputFiles: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.OutputFiles)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.OutputFiles {
			if err := oprot.WriteString(string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:outputFiles: ", p), err)
		}
	}
	return err
}

func (p *BazelAction) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetOutputDirs() {
		if err := oprot.WriteFieldBegin("outputDirs", thrift.LIST, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:outputDirs: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.OutputDirs)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.OutputDirs {
			if err := oprot.WriteString(string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:outputDirs: ", p), err)
		}
	}
	return err
}

func (p *BazelAction) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetPlatformProperties() {
		if err := oprot.WriteFieldBegin("platformProperties", thrift.LIST, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:platformProperties: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.PlatformProperties)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.PlatformProperties {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:platformProperties: ", p), err)
		}
	}
	return err
}

func (p *BazelAction) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimeoutMs() {
		if err := oprot.WriteFieldBegin("timeoutMs", thrift.I64, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:timeoutMs: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.TimeoutMs)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.timeoutMs (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:timeoutMs: ", p), err)
		}
	}
	return err
}

func (p *BazelAction) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetNoCache() {
		if err := oprot.WriteFieldBegin("noCache", thrift.BOOL, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:noCache: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.NoCache)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.noCache (7) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:noCache: ", p), err)
		}
	}
	return err
}

func (p *BazelAction) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BazelAction(%+v)", *p)
}

// Attributes:
//  - Action
//  - InstanceName
//  - SkipCache
type BazelExecuteRequest struct {
	Action       *BazelAction `thrift:"action,1,required" json:"action"`
	InstanceName *string      `thrift:"instanceName,2" json:"instanceName,omitempty"`
	SkipCache    *bool        `thrift:"skipCache,3" json:"skipCache,omitempty"`
}

func NewBazelExecuteRequest() *BazelExecuteRequest {
	return &BazelExecuteRequest{}
}

var BazelExecuteRequest_Action_DEFAULT *BazelAction

func (p *BazelExecuteRequest) GetAction() *BazelAction {
	if !p.IsSetAction() {
		return BazelExecuteRequest_Action_DEFAULT
	}
	return p.Action
}

var BazelExecuteRequest_InstanceName_DEFAULT string

func (p *BazelExecuteRequest) GetInstanceName() string {
	if !p.IsSetInstanceName() {
		return BazelExecuteRequest_InstanceName_DEFAULT
	}
	return *p.InstanceName
}

var BazelExecuteRequest_SkipCache_DEFAULT bool

func (p *BazelExecuteRequest) GetSkipCache() bool {
	if !p.IsSetSkipCache() {
		return BazelExecuteRequest_SkipCache_DEFAULT
	}
	return *p.SkipCache
}
func (p *BazelExecuteRequest) IsSetAction() bool {
	return p.Action != nil
}

func (p *BazelExecuteRequest) IsSetInstanceName() bool {
	return p.InstanceName != nil
}

func (p *BazelExecuteRequest) IsSetSkipCache() bool {
	return p.SkipCache != nil
}

func (p *BazelExecuteRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetAction bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetAction = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetAction {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Action is not set"))
	}
	return nil
}

func (p *BazelExecuteRequest) readField1(iprot thrift.TProtocol) error {
	p.Action = &BazelAction{}
	if err := p.Action.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Action), err)
	}
	return nil
}

func (p *BazelExecuteRequest) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.InstanceName = &v
	}
	return nil
}

func (p *BazelExecuteRequest) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.SkipCache = &v
	}
	return nil
}

func (p *BazelExecuteRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("BazelExecuteRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *BazelExecuteRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("action", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:action: ", p), err)
	}
	if err := p.Action.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Action), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:action: ", p), err)
	}
	return err
}

func (p *BazelExecuteRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetInstanceName() {
		if err := oprot.WriteFieldBegin("instanceName", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:instanceName: ", p), err)
		}
		if err := oprot.WriteString(string(*p.InstanceName)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.instanceName (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:instanceName: ", p), err)
		}
	}
	return err
}

func (p *BazelExecuteRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetSkipCache() {
		if err := oprot.WriteFieldBegin("skipCache", thrift.BOOL, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:skipCache: ", p), err)
		}
		if err := oprot.WriteBool(bool(*p.SkipCache)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.skipCache (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:skipCache: ", p), err)
		}
	}
	return err
}

func (p *BazelExecuteRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BazelExecuteRequest(%+v)", *p)
}
