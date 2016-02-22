package server

import (
	"net"
	"config"
	"flag"
	"runtime"
	"fmt"
	"utils"
	"strconv"
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
	flag.IntVar(&server.num_cpus, "c", runtime.NumCPU(), "")
	flag.StringVar(&server.document_root, "r", "../httptest", "")
	flag.Parse()
	runtime.GOMAXPROCS(server.num_cpus)
	fmt.Println("Running on " + strconv.Itoa(server.num_cpus) + " CPUs")
	server.host = config.Host
	server.port = config.Port
	server.full_address, _ = net.ResolveTCPAddr("tcp", server.host + ":" + server.port)
	fmt.Println("Server variables have been inited.")
}

// func which returns parsed request into utils.Request
func (server *Server) handleRequest(conn net.Conn) (*utils.Request)  {
	buffer := make([]byte, 2048)

	_, read_err := conn.Read(buffer)
	if read_err != nil {
		fmt.Println(read_err)
		return nil
	}

	fmt.Println("Request:")
	fmt.Println(string(buffer))

	return utils.ParseRequest(string(buffer))
}

func (server *Server) serveConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connected on " + conn.RemoteAddr().String())


	request := server.handleRequest(conn)

	if request == nil {
		return
	}

	response := new(utils.Response)
	response.CreateResponse(request.Method, request.Path, server.document_root)
	conn.Write(response.Byte())
}

func (server *Server) Run() {
	fmt.Println("Starting server...")
	server.init()

	fmt.Println("Starting on " + server.full_address.String())
	server.listener, _ = net.ListenTCP("tcp", server.full_address)
	fmt.Println(server.listener)

	for {
		conn, err := server.listener.Accept()

		if err != nil {
			fmt.Println(err)
		}

		if conn != nil {
			go server.serveConnection(conn)
		}
	}

}
