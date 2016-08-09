package main

import (
	"fmt"
	"idgo/config"
	"idgo/idgen"
	"os"
	"strconv"
	"thrift_datatype"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	NetworkAddr = "11.12.112.207:5800"
)

var idNum int

type IdService struct {
	idSeed int64
}

func (this *IdService) GetId(logIndex int64, caller string, ext string) (r *thrift_datatype.ResLong, err error) {
	fmt.Println("GetId")

	curTime := time.Now().Unix()
	this.idSeed++
	id := (curTime << 22) + (int64)(idNum << 10) + this.idSeed % 4096

	return &thrift_datatype.ResLong{200, id, ""}, nil
}

func (this *IdService) Echo(logIndex int64, caller string, srcStr string, ext string) (r *thrift_datatype.ResStr, err error) {
	fmt.Println("echo success!")
	return &thrift_datatype.ResStr{200, srcStr, ""}, err
}

func main() {
	c := config.NewCfg("test.conf")
	if err := c.Load(); err != nil {
		fmt.Println("load config fail...", err)
		os.Exit(1)
	}

	hostNum, err := c.ReadString("hostNum")
	if hostNum == "" {
		fmt.Println("get config hostNum error...")
		os.Exit(1)
	}
	fmt.Printf("hostNum:%s\n", hostNum)
	idNum, err = strconv.Atoi(hostNum)
	if err != nil {
		fmt.Println("hostNum Atoi fail...", err)
		os.Exit(1)
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocloFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Server start Fail", err)
		os.Exit(1)
	}

	handler := &IdService{}
	processor := idgen.NewIdGenServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocloFactory)
	fmt.Println("idgo start success!")
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}
