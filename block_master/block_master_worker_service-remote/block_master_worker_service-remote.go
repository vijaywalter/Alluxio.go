// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"block_master"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  void commitBlock(i64 workerId, i64 usedBytesOnTier, string tierAlias, i64 blockId, i64 length)")
	fmt.Fprintln(os.Stderr, "  i64 getWorkerId(WorkerNetAddress workerNetAddress)")
	fmt.Fprintln(os.Stderr, "  Command heartbeat(i64 workerId,  usedBytesOnTiers,  removedBlockIds,  addedBlocksOnTiers)")
	fmt.Fprintln(os.Stderr, "  void registerWorker(i64 workerId,  storageTiers,  totalBytesOnTiers,  usedBytesOnTiers,  currentBlocksOnTiers)")
	fmt.Fprintln(os.Stderr, "  i64 getServiceVersion()")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := block_master.NewBlockMasterWorkerServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "commitBlock":
		if flag.NArg()-1 != 5 {
			fmt.Fprintln(os.Stderr, "CommitBlock requires 5 args")
			flag.Usage()
		}
		argvalue0, err34 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err34 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1, err35 := (strconv.ParseInt(flag.Arg(2), 10, 64))
		if err35 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		argvalue3, err37 := (strconv.ParseInt(flag.Arg(4), 10, 64))
		if err37 != nil {
			Usage()
			return
		}
		value3 := argvalue3
		argvalue4, err38 := (strconv.ParseInt(flag.Arg(5), 10, 64))
		if err38 != nil {
			Usage()
			return
		}
		value4 := argvalue4
		fmt.Print(client.CommitBlock(value0, value1, value2, value3, value4))
		fmt.Print("\n")
		break
	case "getWorkerId":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetWorkerId requires 1 args")
			flag.Usage()
		}
		arg39 := flag.Arg(1)
		mbTrans40 := thrift.NewTMemoryBufferLen(len(arg39))
		defer mbTrans40.Close()
		_, err41 := mbTrans40.WriteString(arg39)
		if err41 != nil {
			Usage()
			return
		}
		factory42 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt43 := factory42.GetProtocol(mbTrans40)
		argvalue0 := block_master.NewWorkerNetAddress()
		err44 := argvalue0.Read(jsProt43)
		if err44 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetWorkerId(value0))
		fmt.Print("\n")
		break
	case "heartbeat":
		if flag.NArg()-1 != 4 {
			fmt.Fprintln(os.Stderr, "Heartbeat requires 4 args")
			flag.Usage()
		}
		argvalue0, err45 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err45 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg46 := flag.Arg(2)
		mbTrans47 := thrift.NewTMemoryBufferLen(len(arg46))
		defer mbTrans47.Close()
		_, err48 := mbTrans47.WriteString(arg46)
		if err48 != nil {
			Usage()
			return
		}
		factory49 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt50 := factory49.GetProtocol(mbTrans47)
		containerStruct1 := block_master.NewBlockMasterWorkerServiceHeartbeatArgs()
		err51 := containerStruct1.ReadField2(jsProt50)
		if err51 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.UsedBytesOnTiers
		value1 := argvalue1
		arg52 := flag.Arg(3)
		mbTrans53 := thrift.NewTMemoryBufferLen(len(arg52))
		defer mbTrans53.Close()
		_, err54 := mbTrans53.WriteString(arg52)
		if err54 != nil {
			Usage()
			return
		}
		factory55 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt56 := factory55.GetProtocol(mbTrans53)
		containerStruct2 := block_master.NewBlockMasterWorkerServiceHeartbeatArgs()
		err57 := containerStruct2.ReadField3(jsProt56)
		if err57 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.RemovedBlockIds
		value2 := argvalue2
		arg58 := flag.Arg(4)
		mbTrans59 := thrift.NewTMemoryBufferLen(len(arg58))
		defer mbTrans59.Close()
		_, err60 := mbTrans59.WriteString(arg58)
		if err60 != nil {
			Usage()
			return
		}
		factory61 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt62 := factory61.GetProtocol(mbTrans59)
		containerStruct3 := block_master.NewBlockMasterWorkerServiceHeartbeatArgs()
		err63 := containerStruct3.ReadField4(jsProt62)
		if err63 != nil {
			Usage()
			return
		}
		argvalue3 := containerStruct3.AddedBlocksOnTiers
		value3 := argvalue3
		fmt.Print(client.Heartbeat(value0, value1, value2, value3))
		fmt.Print("\n")
		break
	case "registerWorker":
		if flag.NArg()-1 != 5 {
			fmt.Fprintln(os.Stderr, "RegisterWorker requires 5 args")
			flag.Usage()
		}
		argvalue0, err64 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err64 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg65 := flag.Arg(2)
		mbTrans66 := thrift.NewTMemoryBufferLen(len(arg65))
		defer mbTrans66.Close()
		_, err67 := mbTrans66.WriteString(arg65)
		if err67 != nil {
			Usage()
			return
		}
		factory68 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt69 := factory68.GetProtocol(mbTrans66)
		containerStruct1 := block_master.NewBlockMasterWorkerServiceRegisterWorkerArgs()
		err70 := containerStruct1.ReadField2(jsProt69)
		if err70 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.StorageTiers
		value1 := argvalue1
		arg71 := flag.Arg(3)
		mbTrans72 := thrift.NewTMemoryBufferLen(len(arg71))
		defer mbTrans72.Close()
		_, err73 := mbTrans72.WriteString(arg71)
		if err73 != nil {
			Usage()
			return
		}
		factory74 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt75 := factory74.GetProtocol(mbTrans72)
		containerStruct2 := block_master.NewBlockMasterWorkerServiceRegisterWorkerArgs()
		err76 := containerStruct2.ReadField3(jsProt75)
		if err76 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.TotalBytesOnTiers
		value2 := argvalue2
		arg77 := flag.Arg(4)
		mbTrans78 := thrift.NewTMemoryBufferLen(len(arg77))
		defer mbTrans78.Close()
		_, err79 := mbTrans78.WriteString(arg77)
		if err79 != nil {
			Usage()
			return
		}
		factory80 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt81 := factory80.GetProtocol(mbTrans78)
		containerStruct3 := block_master.NewBlockMasterWorkerServiceRegisterWorkerArgs()
		err82 := containerStruct3.ReadField4(jsProt81)
		if err82 != nil {
			Usage()
			return
		}
		argvalue3 := containerStruct3.UsedBytesOnTiers
		value3 := argvalue3
		arg83 := flag.Arg(5)
		mbTrans84 := thrift.NewTMemoryBufferLen(len(arg83))
		defer mbTrans84.Close()
		_, err85 := mbTrans84.WriteString(arg83)
		if err85 != nil {
			Usage()
			return
		}
		factory86 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt87 := factory86.GetProtocol(mbTrans84)
		containerStruct4 := block_master.NewBlockMasterWorkerServiceRegisterWorkerArgs()
		err88 := containerStruct4.ReadField5(jsProt87)
		if err88 != nil {
			Usage()
			return
		}
		argvalue4 := containerStruct4.CurrentBlocksOnTiers
		value4 := argvalue4
		fmt.Print(client.RegisterWorker(value0, value1, value2, value3, value4))
		fmt.Print("\n")
		break
	case "getServiceVersion":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetServiceVersion requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetServiceVersion())
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
