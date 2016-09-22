// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package file_system_master

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

type FileSystemMasterWorkerService interface {
	common.AlluxioService
	//This interface contains file system master service endpoints for Alluxio workers.

	// Parameters:
	//  - FileId: the id of the file
	GetFileInfo(fileId int64) (r *FileInfo, err error)
	// Returns the set of pinned files.
	GetPinIdList() (r map[int64]bool, err error)
	// Periodic file system worker heartbeat. Returns the command for persisting
	// the blocks of a file.
	//
	// Parameters:
	//  - WorkerId: the id of the worker
	//  - PersistedFiles: the list of persisted files
	Heartbeat(workerId int64, persistedFiles []int64) (r *FileSystemCommand, err error)
}

//This interface contains file system master service endpoints for Alluxio workers.
type FileSystemMasterWorkerServiceClient struct {
	*common.AlluxioServiceClient
}

func NewFileSystemMasterWorkerServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *FileSystemMasterWorkerServiceClient {
	return &FileSystemMasterWorkerServiceClient{AlluxioServiceClient: common.NewAlluxioServiceClientFactory(t, f)}
}

func NewFileSystemMasterWorkerServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *FileSystemMasterWorkerServiceClient {
	return &FileSystemMasterWorkerServiceClient{AlluxioServiceClient: common.NewAlluxioServiceClientProtocol(t, iprot, oprot)}
}

// Parameters:
//  - FileId: the id of the file
func (p *FileSystemMasterWorkerServiceClient) GetFileInfo(fileId int64) (r *FileInfo, err error) {
	if err = p.sendGetFileInfo(fileId); err != nil {
		return
	}
	return p.recvGetFileInfo()
}

func (p *FileSystemMasterWorkerServiceClient) sendGetFileInfo(fileId int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getFileInfo", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := FileSystemMasterWorkerServiceGetFileInfoArgs{
		FileId: fileId,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *FileSystemMasterWorkerServiceClient) recvGetFileInfo() (value *FileInfo, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getFileInfo" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getFileInfo failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getFileInfo failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error102 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error103 error
		error103, err = error102.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error103
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getFileInfo failed: invalid message type")
		return
	}
	result := FileSystemMasterWorkerServiceGetFileInfoResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.E != nil {
		err = result.E
		return
	}
	value = result.GetSuccess()
	return
}

// Returns the set of pinned files.
func (p *FileSystemMasterWorkerServiceClient) GetPinIdList() (r map[int64]bool, err error) {
	if err = p.sendGetPinIdList(); err != nil {
		return
	}
	return p.recvGetPinIdList()
}

func (p *FileSystemMasterWorkerServiceClient) sendGetPinIdList() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getPinIdList", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := FileSystemMasterWorkerServiceGetPinIdListArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *FileSystemMasterWorkerServiceClient) recvGetPinIdList() (value map[int64]bool, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getPinIdList" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getPinIdList failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getPinIdList failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error104 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error105 error
		error105, err = error104.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error105
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getPinIdList failed: invalid message type")
		return
	}
	result := FileSystemMasterWorkerServiceGetPinIdListResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Periodic file system worker heartbeat. Returns the command for persisting
// the blocks of a file.
//
// Parameters:
//  - WorkerId: the id of the worker
//  - PersistedFiles: the list of persisted files
func (p *FileSystemMasterWorkerServiceClient) Heartbeat(workerId int64, persistedFiles []int64) (r *FileSystemCommand, err error) {
	if err = p.sendHeartbeat(workerId, persistedFiles); err != nil {
		return
	}
	return p.recvHeartbeat()
}

func (p *FileSystemMasterWorkerServiceClient) sendHeartbeat(workerId int64, persistedFiles []int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("heartbeat", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := FileSystemMasterWorkerServiceHeartbeatArgs{
		WorkerId:       workerId,
		PersistedFiles: persistedFiles,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *FileSystemMasterWorkerServiceClient) recvHeartbeat() (value *FileSystemCommand, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "heartbeat" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "heartbeat failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "heartbeat failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error106 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error107 error
		error107, err = error106.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error107
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "heartbeat failed: invalid message type")
		return
	}
	result := FileSystemMasterWorkerServiceHeartbeatResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.E != nil {
		err = result.E
		return
	}
	value = result.GetSuccess()
	return
}

type FileSystemMasterWorkerServiceProcessor struct {
	*common.AlluxioServiceProcessor
}

func NewFileSystemMasterWorkerServiceProcessor(handler FileSystemMasterWorkerService) *FileSystemMasterWorkerServiceProcessor {
	self108 := &FileSystemMasterWorkerServiceProcessor{common.NewAlluxioServiceProcessor(handler)}
	self108.AddToProcessorMap("getFileInfo", &fileSystemMasterWorkerServiceProcessorGetFileInfo{handler: handler})
	self108.AddToProcessorMap("getPinIdList", &fileSystemMasterWorkerServiceProcessorGetPinIdList{handler: handler})
	self108.AddToProcessorMap("heartbeat", &fileSystemMasterWorkerServiceProcessorHeartbeat{handler: handler})
	return self108
}

type fileSystemMasterWorkerServiceProcessorGetFileInfo struct {
	handler FileSystemMasterWorkerService
}

func (p *fileSystemMasterWorkerServiceProcessorGetFileInfo) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FileSystemMasterWorkerServiceGetFileInfoArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getFileInfo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := FileSystemMasterWorkerServiceGetFileInfoResult{}
	var retval *FileInfo
	var err2 error
	if retval, err2 = p.handler.GetFileInfo(args.FileId); err2 != nil {
		switch v := err2.(type) {
		case *exception.AlluxioTException:
			result.E = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getFileInfo: "+err2.Error())
			oprot.WriteMessageBegin("getFileInfo", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getFileInfo", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type fileSystemMasterWorkerServiceProcessorGetPinIdList struct {
	handler FileSystemMasterWorkerService
}

func (p *fileSystemMasterWorkerServiceProcessorGetPinIdList) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FileSystemMasterWorkerServiceGetPinIdListArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getPinIdList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := FileSystemMasterWorkerServiceGetPinIdListResult{}
	var retval map[int64]bool
	var err2 error
	if retval, err2 = p.handler.GetPinIdList(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getPinIdList: "+err2.Error())
		oprot.WriteMessageBegin("getPinIdList", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getPinIdList", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type fileSystemMasterWorkerServiceProcessorHeartbeat struct {
	handler FileSystemMasterWorkerService
}

func (p *fileSystemMasterWorkerServiceProcessorHeartbeat) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FileSystemMasterWorkerServiceHeartbeatArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("heartbeat", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := FileSystemMasterWorkerServiceHeartbeatResult{}
	var retval *FileSystemCommand
	var err2 error
	if retval, err2 = p.handler.Heartbeat(args.WorkerId, args.PersistedFiles); err2 != nil {
		switch v := err2.(type) {
		case *exception.AlluxioTException:
			result.E = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing heartbeat: "+err2.Error())
			oprot.WriteMessageBegin("heartbeat", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("heartbeat", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - FileId: the id of the file
type FileSystemMasterWorkerServiceGetFileInfoArgs struct {
	FileId int64 `thrift:"fileId,1" json:"fileId"`
}

func NewFileSystemMasterWorkerServiceGetFileInfoArgs() *FileSystemMasterWorkerServiceGetFileInfoArgs {
	return &FileSystemMasterWorkerServiceGetFileInfoArgs{}
}

func (p *FileSystemMasterWorkerServiceGetFileInfoArgs) GetFileId() int64 {
	return p.FileId
}
func (p *FileSystemMasterWorkerServiceGetFileInfoArgs) Read(iprot thrift.TProtocol) error {
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

func (p *FileSystemMasterWorkerServiceGetFileInfoArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.FileId = v
	}
	return nil
}

func (p *FileSystemMasterWorkerServiceGetFileInfoArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getFileInfo_args"); err != nil {
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

func (p *FileSystemMasterWorkerServiceGetFileInfoArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("fileId", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:fileId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.FileId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.fileId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:fileId: ", p), err)
	}
	return err
}

func (p *FileSystemMasterWorkerServiceGetFileInfoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FileSystemMasterWorkerServiceGetFileInfoArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type FileSystemMasterWorkerServiceGetFileInfoResult struct {
	Success *FileInfo                    `thrift:"success,0" json:"success,omitempty"`
	E       *exception.AlluxioTException `thrift:"e,1" json:"e,omitempty"`
}

func NewFileSystemMasterWorkerServiceGetFileInfoResult() *FileSystemMasterWorkerServiceGetFileInfoResult {
	return &FileSystemMasterWorkerServiceGetFileInfoResult{}
}

var FileSystemMasterWorkerServiceGetFileInfoResult_Success_DEFAULT *FileInfo

func (p *FileSystemMasterWorkerServiceGetFileInfoResult) GetSuccess() *FileInfo {
	if !p.IsSetSuccess() {
		return FileSystemMasterWorkerServiceGetFileInfoResult_Success_DEFAULT
	}
	return p.Success
}

var FileSystemMasterWorkerServiceGetFileInfoResult_E_DEFAULT *exception.AlluxioTException

func (p *FileSystemMasterWorkerServiceGetFileInfoResult) GetE() *exception.AlluxioTException {
	if !p.IsSetE() {
		return FileSystemMasterWorkerServiceGetFileInfoResult_E_DEFAULT
	}
	return p.E
}
func (p *FileSystemMasterWorkerServiceGetFileInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FileSystemMasterWorkerServiceGetFileInfoResult) IsSetE() bool {
	return p.E != nil
}

func (p *FileSystemMasterWorkerServiceGetFileInfoResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *FileSystemMasterWorkerServiceGetFileInfoResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &FileInfo{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *FileSystemMasterWorkerServiceGetFileInfoResult) readField1(iprot thrift.TProtocol) error {
	p.E = &exception.AlluxioTException{}
	if err := p.E.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *FileSystemMasterWorkerServiceGetFileInfoResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getFileInfo_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *FileSystemMasterWorkerServiceGetFileInfoResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *FileSystemMasterWorkerServiceGetFileInfoResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin("e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *FileSystemMasterWorkerServiceGetFileInfoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FileSystemMasterWorkerServiceGetFileInfoResult(%+v)", *p)
}

type FileSystemMasterWorkerServiceGetPinIdListArgs struct {
}

func NewFileSystemMasterWorkerServiceGetPinIdListArgs() *FileSystemMasterWorkerServiceGetPinIdListArgs {
	return &FileSystemMasterWorkerServiceGetPinIdListArgs{}
}

func (p *FileSystemMasterWorkerServiceGetPinIdListArgs) Read(iprot thrift.TProtocol) error {
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

func (p *FileSystemMasterWorkerServiceGetPinIdListArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getPinIdList_args"); err != nil {
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

func (p *FileSystemMasterWorkerServiceGetPinIdListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FileSystemMasterWorkerServiceGetPinIdListArgs(%+v)", *p)
}

// Attributes:
//  - Success
type FileSystemMasterWorkerServiceGetPinIdListResult struct {
	Success map[int64]bool `thrift:"success,0" json:"success,omitempty"`
}

func NewFileSystemMasterWorkerServiceGetPinIdListResult() *FileSystemMasterWorkerServiceGetPinIdListResult {
	return &FileSystemMasterWorkerServiceGetPinIdListResult{}
}

var FileSystemMasterWorkerServiceGetPinIdListResult_Success_DEFAULT map[int64]bool

func (p *FileSystemMasterWorkerServiceGetPinIdListResult) GetSuccess() map[int64]bool {
	return p.Success
}
func (p *FileSystemMasterWorkerServiceGetPinIdListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FileSystemMasterWorkerServiceGetPinIdListResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
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

func (p *FileSystemMasterWorkerServiceGetPinIdListResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadSetBegin()
	if err != nil {
		return thrift.PrependError("error reading set begin: ", err)
	}
	tSet := make(map[int64]bool, size)
	p.Success = tSet
	for i := 0; i < size; i++ {
		var _elem109 int64
		if v, err := iprot.ReadI64(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem109 = v
		}
		p.Success[_elem109] = true
	}
	if err := iprot.ReadSetEnd(); err != nil {
		return thrift.PrependError("error reading set end: ", err)
	}
	return nil
}

func (p *FileSystemMasterWorkerServiceGetPinIdListResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getPinIdList_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *FileSystemMasterWorkerServiceGetPinIdListResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.SET, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteSetBegin(thrift.I64, len(p.Success)); err != nil {
			return thrift.PrependError("error writing set begin: ", err)
		}
		for v, _ := range p.Success {
			if err := oprot.WriteI64(int64(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteSetEnd(); err != nil {
			return thrift.PrependError("error writing set end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *FileSystemMasterWorkerServiceGetPinIdListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FileSystemMasterWorkerServiceGetPinIdListResult(%+v)", *p)
}

// Attributes:
//  - WorkerId: the id of the worker
//  - PersistedFiles: the list of persisted files
type FileSystemMasterWorkerServiceHeartbeatArgs struct {
	WorkerId       int64   `thrift:"workerId,1" json:"workerId"`
	PersistedFiles []int64 `thrift:"persistedFiles,2" json:"persistedFiles"`
}

func NewFileSystemMasterWorkerServiceHeartbeatArgs() *FileSystemMasterWorkerServiceHeartbeatArgs {
	return &FileSystemMasterWorkerServiceHeartbeatArgs{}
}

func (p *FileSystemMasterWorkerServiceHeartbeatArgs) GetWorkerId() int64 {
	return p.WorkerId
}

func (p *FileSystemMasterWorkerServiceHeartbeatArgs) GetPersistedFiles() []int64 {
	return p.PersistedFiles
}
func (p *FileSystemMasterWorkerServiceHeartbeatArgs) Read(iprot thrift.TProtocol) error {
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

func (p *FileSystemMasterWorkerServiceHeartbeatArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.WorkerId = v
	}
	return nil
}

func (p *FileSystemMasterWorkerServiceHeartbeatArgs) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]int64, 0, size)
	p.PersistedFiles = tSlice
	for i := 0; i < size; i++ {
		var _elem110 int64
		if v, err := iprot.ReadI64(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem110 = v
		}
		p.PersistedFiles = append(p.PersistedFiles, _elem110)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *FileSystemMasterWorkerServiceHeartbeatArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("heartbeat_args"); err != nil {
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

func (p *FileSystemMasterWorkerServiceHeartbeatArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("workerId", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:workerId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.WorkerId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.workerId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:workerId: ", p), err)
	}
	return err
}

func (p *FileSystemMasterWorkerServiceHeartbeatArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("persistedFiles", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:persistedFiles: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.I64, len(p.PersistedFiles)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.PersistedFiles {
		if err := oprot.WriteI64(int64(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:persistedFiles: ", p), err)
	}
	return err
}

func (p *FileSystemMasterWorkerServiceHeartbeatArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FileSystemMasterWorkerServiceHeartbeatArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type FileSystemMasterWorkerServiceHeartbeatResult struct {
	Success *FileSystemCommand           `thrift:"success,0" json:"success,omitempty"`
	E       *exception.AlluxioTException `thrift:"e,1" json:"e,omitempty"`
}

func NewFileSystemMasterWorkerServiceHeartbeatResult() *FileSystemMasterWorkerServiceHeartbeatResult {
	return &FileSystemMasterWorkerServiceHeartbeatResult{}
}

var FileSystemMasterWorkerServiceHeartbeatResult_Success_DEFAULT *FileSystemCommand

func (p *FileSystemMasterWorkerServiceHeartbeatResult) GetSuccess() *FileSystemCommand {
	if !p.IsSetSuccess() {
		return FileSystemMasterWorkerServiceHeartbeatResult_Success_DEFAULT
	}
	return p.Success
}

var FileSystemMasterWorkerServiceHeartbeatResult_E_DEFAULT *exception.AlluxioTException

func (p *FileSystemMasterWorkerServiceHeartbeatResult) GetE() *exception.AlluxioTException {
	if !p.IsSetE() {
		return FileSystemMasterWorkerServiceHeartbeatResult_E_DEFAULT
	}
	return p.E
}
func (p *FileSystemMasterWorkerServiceHeartbeatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FileSystemMasterWorkerServiceHeartbeatResult) IsSetE() bool {
	return p.E != nil
}

func (p *FileSystemMasterWorkerServiceHeartbeatResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
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

func (p *FileSystemMasterWorkerServiceHeartbeatResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &FileSystemCommand{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *FileSystemMasterWorkerServiceHeartbeatResult) readField1(iprot thrift.TProtocol) error {
	p.E = &exception.AlluxioTException{}
	if err := p.E.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *FileSystemMasterWorkerServiceHeartbeatResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("heartbeat_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
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

func (p *FileSystemMasterWorkerServiceHeartbeatResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *FileSystemMasterWorkerServiceHeartbeatResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin("e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *FileSystemMasterWorkerServiceHeartbeatResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FileSystemMasterWorkerServiceHeartbeatResult(%+v)", *p)
}