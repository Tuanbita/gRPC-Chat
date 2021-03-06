// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "openstars/core/bigset/generic"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  void needSplit(TContainerKey rootID, TNeedSplitInfo splitInfo)")
  fmt.Fprintln(os.Stderr, "  void splitInfoUpdated(TContainerKey rootID)")
  fmt.Fprintln(os.Stderr, "  SplitJob getJob()")
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
  client := generic.NewBSNotificationPoolClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "needSplit":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "NeedSplit requires 2 args")
      flag.Usage()
    }
    argvalue0, err264 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err264 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    arg265 := flag.Arg(2)
    mbTrans266 := thrift.NewTMemoryBufferLen(len(arg265))
    defer mbTrans266.Close()
    _, err267 := mbTrans266.WriteString(arg265)
    if err267 != nil {
      Usage()
      return
    }
    factory268 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt269 := factory268.GetProtocol(mbTrans266)
    argvalue1 := generic.NewTNeedSplitInfo()
    err270 := argvalue1.Read(jsProt269)
    if err270 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.NeedSplit(value0, value1))
    fmt.Print("\n")
    break
  case "splitInfoUpdated":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SplitInfoUpdated requires 1 args")
      flag.Usage()
    }
    argvalue0, err271 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err271 != nil {
      Usage()
      return
    }
    value0 := generic.TContainerKey(argvalue0)
    fmt.Print(client.SplitInfoUpdated(value0))
    fmt.Print("\n")
    break
  case "getJob":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetJob requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetJob())
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
