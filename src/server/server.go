package server

import (
	"net"
	"config"
	"flag"
	"runtime"
	"fmt"
	"utils"
	"strconv"
	"errors"
)

type Server struct {
	host		string
	port 		string
	num_cpus 	int
	full_address	*net.TCPAddr
	listener 	*net.TCPListener
	document_root 	string
}

func (server *Server) init()  {
	utils.InitLog()

	flag.IntVar(&server.num_cpus, "c", runtime.NumCPU(), "")
	flag.StringVar(&server.document_root, "r", "../httptest", "")
	flag.Parse()

	runtime.GOMAXPROCS(server.num_cpus)
	utils.LogInfo("Running on " + strconv.Itoa(server.num_cpus) + " CPUs")

	server.host = config.Host

	server.port = config.Port

	server.full_address, _ = net.ResolveTCPAddr("tcp", server.host + ":" + server.port)
}

// func which returns parsed request into utils.Request
func (server *Server) handleRequest(conn net.Conn) (*utils.Request, error)  {
	buffer := make([]byte, 2048)

	_, read_err := conn.Read(buffer)
	if read_err != nil {
		fmt.Println(read_err)
		return nil, errors.New("Bad request")
	}

	parsed_request, err := utils.ParseRequest(string(buffer))

	if err != nil {
		return nil, err
	}

	return parsed_request, nil
}

func (server *Server) serveConnection(conn net.Conn) {
	defer conn.Close()

	utils.LogError("New connection from " + conn.RemoteAddr().String())

	// parse input request to Request{}
	request, err := server.handleRequest(conn)

	if err != nil {
		response := new(utils.Response)
		response.CreateResponseForBadRequest()
		conn.Write(response.Byte())
		return
	}

	// make a Response{}
	response := new(utils.Response)
	response.CreateResponse(request.Method, request.Path, server.document_root)
	// give it back in []byte form
	conn.Write(response.Byte())
}

func (server *Server) Run() {
	server.init()

	server.listener, _ = net.ListenTCP("tcp", server.full_address)

	for {
		conn, err := server.listener.Accept()

		if err != nil {
			utils.LogError(err)
		}

		if conn != nil {
			go server.serveConnection(conn)
		}
	}

}
