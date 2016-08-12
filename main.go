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
	"github.com/kataras/iris"
)

const (
	NetworkAddr = "11.12.112.87:5800"
)

var idNum int

var idSeed int64

type IdService struct {
}

func (this *IdService) GetId(logIndex int64, caller string, ext string) (r *thrift_datatype.ResLong, err error) {
	fmt.Println("GetId")

	curTime := time.Now().Unix()
	idSeed++
	id := (curTime << 22) + (int64)(idNum<<10) + idSeed%4096

	return &thrift_datatype.ResLong{200, id, ""}, nil
}

func (this *IdService) Echo(logIndex int64, caller string, srcStr string, ext string) (r *thrift_datatype.ResStr, err error) {
	fmt.Println("echo success!")
	return &thrift_datatype.ResStr{200, srcStr, ""}, err
}

type IdAPI struct {
	*iris.Context
}

//GET /id
func (idApi IdAPI) Get() {

	curTime := time.Now().Unix()
	idSeed++
	id := (curTime << 22) + (int64)(idNum<<10) + idSeed%4096

	fmt.Println("get id ", id)

	idApi.Write(strconv.FormatInt(id, 10))
}

func startListen() {
	iris.API("/getid", IdAPI{})
	iris.Listen(":8080")
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

	go startListen()

	server.Serve()
}
