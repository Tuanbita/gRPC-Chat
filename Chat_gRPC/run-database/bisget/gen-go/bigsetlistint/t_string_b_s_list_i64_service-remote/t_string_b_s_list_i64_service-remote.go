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
        "bigsetlistint"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TBigSetInfoResult createStringBigSet(TStringKey bsName)")
  fmt.Fprintln(os.Stderr, "  TBigSetInfoResult getBigSetInfoByName(TStringKey bsName)")
  fmt.Fprintln(os.Stderr, "  TBigSetInfoResult assignBigSetName(TStringKey bsName, TContainerKey bigsetID)")
  fmt.Fprintln(os.Stderr, "  TPutItemResult bsPutItem(TStringKey bsName, TItem item)")
  fmt.Fprintln(os.Stderr, "  TChildItemResult addChildItem(TStringKey bsName, TItemKey itemKey, TItemChild aChild, TChildItemOptions opOption)")
  fmt.Fprintln(os.Stderr, "  TChildItemResult addChildrenItem(TStringKey bsName, TItemKey itemKey,  aChild, TChildItemOptions opOption)")
  fmt.Fprintln(os.Stderr, "  TChildItemResult removeChildItem(TStringKey bsName, TItemKey itemKey, TItemChild aChild, TChildItemOptions opOption)")
  fmt.Fprintln(os.Stderr, "  bool bsRemoveItem(TStringKey bsName, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TExistedResult bsExisted(TStringKey bsName, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TItemResult bsGetItem(TStringKey bsName, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsGetSlice(TStringKey bsName, i32 fromPos, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsGetSliceFromItem(TStringKey bsName, TItemKey fromKey, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsGetSliceR(TStringKey bsName, i32 fromPos, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsGetSliceFromItemR(TStringKey bsName, TItemKey fromKey, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsRangeQuery(TStringKey bsName, TItemKey startKey, TItemKey endKey)")
  fmt.Fprintln(os.Stderr, "  bool bsBulkLoad(TStringKey bsName, TItemSet setData)")
  fmt.Fprintln(os.Stderr, "  TMultiPutItemResult bsMultiPut(TStringKey bsName, TItemSet setData, bool getAddedItems, bool getReplacedItems)")
  fmt.Fprintln(os.Stderr, "  i64 getTotalCount(TStringKey bsName)")
  fmt.Fprintln(os.Stderr, "  i64 removeAll(TStringKey bsName)")
  fmt.Fprintln(os.Stderr, "  i64 totalStringKeyCount()")
  fmt.Fprintln(os.Stderr, "   getListKey(i64 fromIndex, i32 count)")
  fmt.Fprintln(os.Stderr, "   getListKeyFrom(TStringKey keyFrom, i32 count)")
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
  client := bigsetlistint.NewTStringBSListI64ServiceClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "createStringBigSet":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CreateStringBigSet requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    fmt.Print(client.CreateStringBigSet(value0))
    fmt.Print("\n")
    break
  case "getBigSetInfoByName":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetBigSetInfoByName requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    fmt.Print(client.GetBigSetInfoByName(value0))
    fmt.Print("\n")
    break
  case "assignBigSetName":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "AssignBigSetName requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    argvalue1, err160 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err160 != nil {
      Usage()
      return
    }
    value1 := bigsetlistint.TContainerKey(argvalue1)
    fmt.Print(client.AssignBigSetName(value0, value1))
    fmt.Print("\n")
    break
  case "bsPutItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsPutItem requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    arg162 := flag.Arg(2)
    mbTrans163 := thrift.NewTMemoryBufferLen(len(arg162))
    defer mbTrans163.Close()
    _, err164 := mbTrans163.WriteString(arg162)
    if err164 != nil {
      Usage()
      return
    }
    factory165 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt166 := factory165.GetProtocol(mbTrans163)
    argvalue1 := bigsetlistint.NewTItem()
    err167 := argvalue1.Read(jsProt166)
    if err167 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.BsPutItem(value0, value1))
    fmt.Print("\n")
    break
  case "addChildItem":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "AddChildItem requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := bigsetlistint.TItemKey(argvalue1)
    argvalue2, err170 := (strconv.ParseInt(flag.Arg(3), 10, 64))
    if err170 != nil {
      Usage()
      return
    }
    value2 := bigsetlistint.TItemChild(argvalue2)
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := bigsetlistint.TChildItemOptions(tmp3)
    value3 := argvalue3
    fmt.Print(client.AddChildItem(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "addChildrenItem":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "AddChildrenItem requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := bigsetlistint.TItemKey(argvalue1)
    arg173 := flag.Arg(3)
    mbTrans174 := thrift.NewTMemoryBufferLen(len(arg173))
    defer mbTrans174.Close()
    _, err175 := mbTrans174.WriteString(arg173)
    if err175 != nil { 
      Usage()
      return
    }
    factory176 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt177 := factory176.GetProtocol(mbTrans174)
    containerStruct2 := bigsetlistint.NewTStringBSListI64ServiceAddChildrenItemArgs()
    err178 := containerStruct2.ReadField3(jsProt177)
    if err178 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.AChild
    value2 := argvalue2
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := bigsetlistint.TChildItemOptions(tmp3)
    value3 := argvalue3
    fmt.Print(client.AddChildrenItem(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "removeChildItem":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "RemoveChildItem requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := bigsetlistint.TItemKey(argvalue1)
    argvalue2, err181 := (strconv.ParseInt(flag.Arg(3), 10, 64))
    if err181 != nil {
      Usage()
      return
    }
    value2 := bigsetlistint.TItemChild(argvalue2)
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := bigsetlistint.TChildItemOptions(tmp3)
    value3 := argvalue3
    fmt.Print(client.RemoveChildItem(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "bsRemoveItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsRemoveItem requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := bigsetlistint.TItemKey(argvalue1)
    fmt.Print(client.BsRemoveItem(value0, value1))
    fmt.Print("\n")
    break
  case "bsExisted":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsExisted requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := bigsetlistint.TItemKey(argvalue1)
    fmt.Print(client.BsExisted(value0, value1))
    fmt.Print("\n")
    break
  case "bsGetItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsGetItem requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := bigsetlistint.TItemKey(argvalue1)
    fmt.Print(client.BsGetItem(value0, value1))
    fmt.Print("\n")
    break
  case "bsGetSlice":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsGetSlice requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    tmp1, err189 := (strconv.Atoi(flag.Arg(2)))
    if err189 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err190 := (strconv.Atoi(flag.Arg(3)))
    if err190 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsGetSlice(value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsGetSliceFromItem":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsGetSliceFromItem requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := bigsetlistint.TItemKey(argvalue1)
    tmp2, err193 := (strconv.Atoi(flag.Arg(3)))
    if err193 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsGetSliceFromItem(value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsGetSliceR":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsGetSliceR requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    tmp1, err195 := (strconv.Atoi(flag.Arg(2)))
    if err195 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err196 := (strconv.Atoi(flag.Arg(3)))
    if err196 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsGetSliceR(value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsGetSliceFromItemR":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsGetSliceFromItemR requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := bigsetlistint.TItemKey(argvalue1)
    tmp2, err199 := (strconv.Atoi(flag.Arg(3)))
    if err199 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsGetSliceFromItemR(value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsRangeQuery":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsRangeQuery requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := bigsetlistint.TItemKey(argvalue1)
    argvalue2 := []byte(flag.Arg(3))
    value2 := bigsetlistint.TItemKey(argvalue2)
    fmt.Print(client.BsRangeQuery(value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsBulkLoad":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsBulkLoad requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    arg204 := flag.Arg(2)
    mbTrans205 := thrift.NewTMemoryBufferLen(len(arg204))
    defer mbTrans205.Close()
    _, err206 := mbTrans205.WriteString(arg204)
    if err206 != nil {
      Usage()
      return
    }
    factory207 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt208 := factory207.GetProtocol(mbTrans205)
    argvalue1 := bigsetlistint.NewTItemSet()
    err209 := argvalue1.Read(jsProt208)
    if err209 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.BsBulkLoad(value0, value1))
    fmt.Print("\n")
    break
  case "bsMultiPut":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "BsMultiPut requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    arg211 := flag.Arg(2)
    mbTrans212 := thrift.NewTMemoryBufferLen(len(arg211))
    defer mbTrans212.Close()
    _, err213 := mbTrans212.WriteString(arg211)
    if err213 != nil {
      Usage()
      return
    }
    factory214 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt215 := factory214.GetProtocol(mbTrans212)
    argvalue1 := bigsetlistint.NewTItemSet()
    err216 := argvalue1.Read(jsProt215)
    if err216 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    argvalue2 := flag.Arg(3) == "true"
    value2 := argvalue2
    argvalue3 := flag.Arg(4) == "true"
    value3 := argvalue3
    fmt.Print(client.BsMultiPut(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "getTotalCount":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTotalCount requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    fmt.Print(client.GetTotalCount(value0))
    fmt.Print("\n")
    break
  case "removeAll":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RemoveAll requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    fmt.Print(client.RemoveAll(value0))
    fmt.Print("\n")
    break
  case "totalStringKeyCount":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "TotalStringKeyCount requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.TotalStringKeyCount())
    fmt.Print("\n")
    break
  case "getListKey":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetListKey requires 2 args")
      flag.Usage()
    }
    argvalue0, err221 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err221 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err222 := (strconv.Atoi(flag.Arg(2)))
    if err222 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.GetListKey(value0, value1))
    fmt.Print("\n")
    break
  case "getListKeyFrom":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetListKeyFrom requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := bigsetlistint.TStringKey(argvalue0)
    tmp1, err224 := (strconv.Atoi(flag.Arg(2)))
    if err224 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.GetListKeyFrom(value0, value1))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
