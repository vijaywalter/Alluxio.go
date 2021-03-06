// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package file_system_worker

import (
	"bytes"
	"common"
	"exception"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = common.GoUnusedProtection__
var _ = exception.GoUnusedProtection__
var GoUnusedProtection__ int

type CancelUfsFileTOptions struct {
}

func NewCancelUfsFileTOptions() *CancelUfsFileTOptions {
	return &CancelUfsFileTOptions{}
}

func (p *CancelUfsFileTOptions) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CancelUfsFileTOptions) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CancelUfsFileTOptions"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CancelUfsFileTOptions) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CancelUfsFileTOptions(%+v)", *p)
}

type CloseUfsFileTOptions struct {
}

func NewCloseUfsFileTOptions() *CloseUfsFileTOptions {
	return &CloseUfsFileTOptions{}
}

func (p *CloseUfsFileTOptions) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CloseUfsFileTOptions) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CloseUfsFileTOptions"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CloseUfsFileTOptions) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CloseUfsFileTOptions(%+v)", *p)
}

// Attributes:
//  - Owner
//  - Group
//  - Mode
type CompleteUfsFileTOptions struct {
	Owner *string `thrift:"owner,1" json:"owner,omitempty"`
	Group *string `thrift:"group,2" json:"group,omitempty"`
	Mode  *int16  `thrift:"mode,3" json:"mode,omitempty"`
}

func NewCompleteUfsFileTOptions() *CompleteUfsFileTOptions {
	return &CompleteUfsFileTOptions{}
}

var CompleteUfsFileTOptions_Owner_DEFAULT string

func (p *CompleteUfsFileTOptions) GetOwner() string {
	if !p.IsSetOwner() {
		return CompleteUfsFileTOptions_Owner_DEFAULT
	}
	return *p.Owner
}

var CompleteUfsFileTOptions_Group_DEFAULT string

func (p *CompleteUfsFileTOptions) GetGroup() string {
	if !p.IsSetGroup() {
		return CompleteUfsFileTOptions_Group_DEFAULT
	}
	return *p.Group
}

var CompleteUfsFileTOptions_Mode_DEFAULT int16

func (p *CompleteUfsFileTOptions) GetMode() int16 {
	if !p.IsSetMode() {
		return CompleteUfsFileTOptions_Mode_DEFAULT
	}
	return *p.Mode
}
func (p *CompleteUfsFileTOptions) IsSetOwner() bool {
	return p.Owner != nil
}

func (p *CompleteUfsFileTOptions) IsSetGroup() bool {
	return p.Group != nil
}

func (p *CompleteUfsFileTOptions) IsSetMode() bool {
	return p.Mode != nil
}

func (p *CompleteUfsFileTOptions) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

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
	return nil
}

func (p *CompleteUfsFileTOptions) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Owner = &v
	}
	return nil
}

func (p *CompleteUfsFileTOptions) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Group = &v
	}
	return nil
}

func (p *CompleteUfsFileTOptions) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Mode = &v
	}
	return nil
}

func (p *CompleteUfsFileTOptions) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CompleteUfsFileTOptions"); err != nil {
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

func (p *CompleteUfsFileTOptions) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetOwner() {
		if err := oprot.WriteFieldBegin("owner", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:owner: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Owner)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.owner (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:owner: ", p), err)
		}
	}
	return err
}

func (p *CompleteUfsFileTOptions) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetGroup() {
		if err := oprot.WriteFieldBegin("group", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:group: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Group)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.group (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:group: ", p), err)
		}
	}
	return err
}

func (p *CompleteUfsFileTOptions) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetMode() {
		if err := oprot.WriteFieldBegin("mode", thrift.I16, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:mode: ", p), err)
		}
		if err := oprot.WriteI16(int16(*p.Mode)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.mode (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:mode: ", p), err)
		}
	}
	return err
}

func (p *CompleteUfsFileTOptions) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CompleteUfsFileTOptions(%+v)", *p)
}

// Attributes:
//  - Owner
//  - Group
//  - Mode
type CreateUfsFileTOptions struct {
	Owner *string `thrift:"owner,1" json:"owner,omitempty"`
	Group *string `thrift:"group,2" json:"group,omitempty"`
	Mode  *int16  `thrift:"mode,3" json:"mode,omitempty"`
}

func NewCreateUfsFileTOptions() *CreateUfsFileTOptions {
	return &CreateUfsFileTOptions{}
}

var CreateUfsFileTOptions_Owner_DEFAULT string

func (p *CreateUfsFileTOptions) GetOwner() string {
	if !p.IsSetOwner() {
		return CreateUfsFileTOptions_Owner_DEFAULT
	}
	return *p.Owner
}

var CreateUfsFileTOptions_Group_DEFAULT string

func (p *CreateUfsFileTOptions) GetGroup() string {
	if !p.IsSetGroup() {
		return CreateUfsFileTOptions_Group_DEFAULT
	}
	return *p.Group
}

var CreateUfsFileTOptions_Mode_DEFAULT int16

func (p *CreateUfsFileTOptions) GetMode() int16 {
	if !p.IsSetMode() {
		return CreateUfsFileTOptions_Mode_DEFAULT
	}
	return *p.Mode
}
func (p *CreateUfsFileTOptions) IsSetOwner() bool {
	return p.Owner != nil
}

func (p *CreateUfsFileTOptions) IsSetGroup() bool {
	return p.Group != nil
}

func (p *CreateUfsFileTOptions) IsSetMode() bool {
	return p.Mode != nil
}

func (p *CreateUfsFileTOptions) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

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
	return nil
}

func (p *CreateUfsFileTOptions) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Owner = &v
	}
	return nil
}

func (p *CreateUfsFileTOptions) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Group = &v
	}
	return nil
}

func (p *CreateUfsFileTOptions) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Mode = &v
	}
	return nil
}

func (p *CreateUfsFileTOptions) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CreateUfsFileTOptions"); err != nil {
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

func (p *CreateUfsFileTOptions) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetOwner() {
		if err := oprot.WriteFieldBegin("owner", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:owner: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Owner)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.owner (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:owner: ", p), err)
		}
	}
	return err
}

func (p *CreateUfsFileTOptions) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetGroup() {
		if err := oprot.WriteFieldBegin("group", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:group: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Group)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.group (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:group: ", p), err)
		}
	}
	return err
}

func (p *CreateUfsFileTOptions) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetMode() {
		if err := oprot.WriteFieldBegin("mode", thrift.I16, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:mode: ", p), err)
		}
		if err := oprot.WriteI16(int16(*p.Mode)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.mode (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:mode: ", p), err)
		}
	}
	return err
}

func (p *CreateUfsFileTOptions) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateUfsFileTOptions(%+v)", *p)
}

type OpenUfsFileTOptions struct {
}

func NewOpenUfsFileTOptions() *OpenUfsFileTOptions {
	return &OpenUfsFileTOptions{}
}

func (p *OpenUfsFileTOptions) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *OpenUfsFileTOptions) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("OpenUfsFileTOptions"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *OpenUfsFileTOptions) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OpenUfsFileTOptions(%+v)", *p)
}
