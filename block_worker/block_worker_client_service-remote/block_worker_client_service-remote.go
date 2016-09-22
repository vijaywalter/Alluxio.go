// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"block_worker"
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
	fmt.Fprintln(os.Stderr, "  void accessBlock(i64 blockId)")
	fmt.Fprintln(os.Stderr, "  bool asyncCheckpoint(i64 fileId)")
	fmt.Fprintln(os.Stderr, "  void cacheBlock(i64 sessionId, i64 blockId)")
	fmt.Fprintln(os.Stderr, "  void cancelBlock(i64 sessionId, i64 blockId)")
	fmt.Fprintln(os.Stderr, "  LockBlockResult lockBlock(i64 blockId, i64 sessionId)")
	fmt.Fprintln(os.Stderr, "  bool promoteBlock(i64 blockId)")
	fmt.Fprintln(os.Stderr, "  string requestBlockLocation(i64 sessionId, i64 blockId, i64 initialBytes)")
	fmt.Fprintln(os.Stderr, "  bool requestSpace(i64 sessionId, i64 blockId, i64 requestBytes)")
	fmt.Fprintln(os.Stderr, "  void sessionHeartbeat(i64 sessionId,  metrics)")
	fmt.Fprintln(os.Stderr, "  bool unlockBlock(i64 blockId, i64 sessionId)")
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
	client := block_worker.NewBlockWorkerClientServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "accessBlock":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AccessBlock requires 1 args")
			flag.Usage()
		}
		argvalue0, err22 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err22 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.AccessBlock(value0))
		fmt.Print("\n")
		break
	case "asyncCheckpoint":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AsyncCheckpoint requires 1 args")
			flag.Usage()
		}
		argvalue0, err23 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err23 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.AsyncCheckpoint(value0))
		fmt.Print("\n")
		break
	case "cacheBlock":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "CacheBlock requires 2 args")
			flag.Usage()
		}
		argvalue0, err24 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err24 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1, err25 := (strconv.ParseInt(flag.Arg(2), 10, 64))
		if err25 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.CacheBlock(value0, value1))
		fmt.Print("\n")
		break
	case "cancelBlock":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "CancelBlock requires 2 args")
			flag.Usage()
		}
		argvalue0, err26 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err26 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1, err27 := (strconv.ParseInt(flag.Arg(2), 10, 64))
		if err27 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.CancelBlock(value0, value1))
		fmt.Print("\n")
		break
	case "lockBlock":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "LockBlock requires 2 args")
			flag.Usage()
		}
		argvalue0, err28 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err28 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1, err29 := (strconv.ParseInt(flag.Arg(2), 10, 64))
		if err29 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.LockBlock(value0, value1))
		fmt.Print("\n")
		break
	case "promoteBlock":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "PromoteBlock requires 1 args")
			flag.Usage()
		}
		argvalue0, err30 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err30 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.PromoteBlock(value0))
		fmt.Print("\n")
		break
	case "requestBlockLocation":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "RequestBlockLocation requires 3 args")
			flag.Usage()
		}
		argvalue0, err31 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err31 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1, err32 := (strconv.ParseInt(flag.Arg(2), 10, 64))
		if err32 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		argvalue2, err33 := (strconv.ParseInt(flag.Arg(3), 10, 64))
		if err33 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		fmt.Print(client.RequestBlockLocation(value0, value1, value2))
		fmt.Print("\n")
		break
	case "requestSpace":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "RequestSpace requires 3 args")
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
		argvalue2, err36 := (strconv.ParseInt(flag.Arg(3), 10, 64))
		if err36 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		fmt.Print(client.RequestSpace(value0, value1, value2))
		fmt.Print("\n")
		break
	case "sessionHeartbeat":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SessionHeartbeat requires 2 args")
			flag.Usage()
		}
		argvalue0, err37 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err37 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg38 := flag.Arg(2)
		mbTrans39 := thrift.NewTMemoryBufferLen(len(arg38))
		defer mbTrans39.Close()
		_, err40 := mbTrans39.WriteString(arg38)
		if err40 != nil {
			Usage()
			return
		}
		factory41 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt42 := factory41.GetProtocol(mbTrans39)
		containerStruct1 := block_worker.NewBlockWorkerClientServiceSessionHeartbeatArgs()
		err43 := containerStruct1.ReadField2(jsProt42)
		if err43 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Metrics
		value1 := argvalue1
		fmt.Print(client.SessionHeartbeat(value0, value1))
		fmt.Print("\n")
		break
	case "unlockBlock":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "UnlockBlock requires 2 args")
			flag.Usage()
		}
		argvalue0, err44 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err44 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1, err45 := (strconv.ParseInt(flag.Arg(2), 10, 64))
		if err45 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.UnlockBlock(value0, value1))
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
