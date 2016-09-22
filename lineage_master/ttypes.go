// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package lineage_master

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

// Attributes:
//  - Command
//  - Conf
type CommandLineJobInfo struct {
	Command string       `thrift:"command,1" json:"command"`
	Conf    *JobConfInfo `thrift:"conf,2" json:"conf"`
}

func NewCommandLineJobInfo() *CommandLineJobInfo {
	return &CommandLineJobInfo{}
}

func (p *CommandLineJobInfo) GetCommand() string {
	return p.Command
}

var CommandLineJobInfo_Conf_DEFAULT JobConfInfo

func (p *CommandLineJobInfo) GetConf() JobConfInfo {
	if !p.IsSetConf() {
		return CommandLineJobInfo_Conf_DEFAULT
	}
	return *p.Conf
}
func (p *CommandLineJobInfo) IsSetConf() bool {
	return p.Conf != nil
}

func (p *CommandLineJobInfo) Read(iprot thrift.TProtocol) error {
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

func (p *CommandLineJobInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Command = v
	}
	return nil
}

func (p *CommandLineJobInfo) readField2(iprot thrift.TProtocol) error {
	p.Conf = &JobConfInfo{}
	if err := p.Conf.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Conf), err)
	}
	return nil
}

func (p *CommandLineJobInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CommandLineJobInfo"); err != nil {
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

func (p *CommandLineJobInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("command", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:command: ", p), err)
	}
	if err := oprot.WriteString(string(p.Command)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.command (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:command: ", p), err)
	}
	return err
}

func (p *CommandLineJobInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("conf", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:conf: ", p), err)
	}
	if err := p.Conf.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Conf), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:conf: ", p), err)
	}
	return err
}

func (p *CommandLineJobInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommandLineJobInfo(%+v)", *p)
}

type DependencyInfo struct {
}

func NewDependencyInfo() *DependencyInfo {
	return &DependencyInfo{}
}

func (p *DependencyInfo) Read(iprot thrift.TProtocol) error {
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

func (p *DependencyInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("DependencyInfo"); err != nil {
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

func (p *DependencyInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DependencyInfo(%+v)", *p)
}

// Attributes:
//  - OutputFile
type JobConfInfo struct {
	OutputFile string `thrift:"outputFile,1" json:"outputFile"`
}

func NewJobConfInfo() *JobConfInfo {
	return &JobConfInfo{}
}

func (p *JobConfInfo) GetOutputFile() string {
	return p.OutputFile
}
func (p *JobConfInfo) Read(iprot thrift.TProtocol) error {
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

func (p *JobConfInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.OutputFile = v
	}
	return nil
}

func (p *JobConfInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("JobConfInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
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

func (p *JobConfInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("outputFile", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:outputFile: ", p), err)
	}
	if err := oprot.WriteString(string(p.OutputFile)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.outputFile (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:outputFile: ", p), err)
	}
	return err
}

func (p *JobConfInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("JobConfInfo(%+v)", *p)
}

// Attributes:
//  - ID
//  - InputFiles
//  - OutputFiles
//  - Job
//  - CreationTimeMs
//  - Parents
//  - Children
type LineageInfo struct {
	ID             int64               `thrift:"id,1" json:"id"`
	InputFiles     []string            `thrift:"inputFiles,2" json:"inputFiles"`
	OutputFiles    []string            `thrift:"outputFiles,3" json:"outputFiles"`
	Job            *CommandLineJobInfo `thrift:"job,4" json:"job"`
	CreationTimeMs int64               `thrift:"creationTimeMs,5" json:"creationTimeMs"`
	Parents        []int64             `thrift:"parents,6" json:"parents"`
	Children       []int64             `thrift:"children,7" json:"children"`
}

func NewLineageInfo() *LineageInfo {
	return &LineageInfo{}
}

func (p *LineageInfo) GetID() int64 {
	return p.ID
}

func (p *LineageInfo) GetInputFiles() []string {
	return p.InputFiles
}

func (p *LineageInfo) GetOutputFiles() []string {
	return p.OutputFiles
}

var LineageInfo_Job_DEFAULT *CommandLineJobInfo

func (p *LineageInfo) GetJob() *CommandLineJobInfo {
	if !p.IsSetJob() {
		return LineageInfo_Job_DEFAULT
	}
	return p.Job
}

func (p *LineageInfo) GetCreationTimeMs() int64 {
	return p.CreationTimeMs
}

func (p *LineageInfo) GetParents() []int64 {
	return p.Parents
}

func (p *LineageInfo) GetChildren() []int64 {
	return p.Children
}
func (p *LineageInfo) IsSetJob() bool {
	return p.Job != nil
}

func (p *LineageInfo) Read(iprot thrift.TProtocol) error {
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
	return nil
}

func (p *LineageInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *LineageInfo) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.InputFiles = tSlice
	for i := 0; i < size; i++ {
		var _elem0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.InputFiles = append(p.InputFiles, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *LineageInfo) readField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.OutputFiles = tSlice
	for i := 0; i < size; i++ {
		var _elem1 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem1 = v
		}
		p.OutputFiles = append(p.OutputFiles, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *LineageInfo) readField4(iprot thrift.TProtocol) error {
	p.Job = &CommandLineJobInfo{}
	if err := p.Job.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Job), err)
	}
	return nil
}

func (p *LineageInfo) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.CreationTimeMs = v
	}
	return nil
}

func (p *LineageInfo) readField6(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]int64, 0, size)
	p.Parents = tSlice
	for i := 0; i < size; i++ {
		var _elem2 int64
		if v, err := iprot.ReadI64(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem2 = v
		}
		p.Parents = append(p.Parents, _elem2)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *LineageInfo) readField7(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]int64, 0, size)
	p.Children = tSlice
	for i := 0; i < size; i++ {
		var _elem3 int64
		if v, err := iprot.ReadI64(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem3 = v
		}
		p.Children = append(p.Children, _elem3)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *LineageInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("LineageInfo"); err != nil {
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

func (p *LineageInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return err
}

func (p *LineageInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("inputFiles", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:inputFiles: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.InputFiles)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.InputFiles {
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:inputFiles: ", p), err)
	}
	return err
}

func (p *LineageInfo) writeField3(oprot thrift.TProtocol) (err error) {
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
	return err
}

func (p *LineageInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("job", thrift.STRUCT, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:job: ", p), err)
	}
	if err := p.Job.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Job), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:job: ", p), err)
	}
	return err
}

func (p *LineageInfo) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("creationTimeMs", thrift.I64, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:creationTimeMs: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.CreationTimeMs)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.creationTimeMs (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:creationTimeMs: ", p), err)
	}
	return err
}

func (p *LineageInfo) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("parents", thrift.LIST, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:parents: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.I64, len(p.Parents)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Parents {
		if err := oprot.WriteI64(int64(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:parents: ", p), err)
	}
	return err
}

func (p *LineageInfo) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("children", thrift.LIST, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:children: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.I64, len(p.Children)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Children {
		if err := oprot.WriteI64(int64(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:children: ", p), err)
	}
	return err
}

func (p *LineageInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LineageInfo(%+v)", *p)
}
