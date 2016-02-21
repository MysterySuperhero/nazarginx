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
}

func (server *Server) init()  {
	flag.IntVar(&server.num_cpus, "cpus", runtime.NumCPU(), "")
	flag.Parse()
	runtime.GOMAXPROCS(server.num_cpus)
	fmt.Println("Running on " + strconv.Itoa(server.num_cpus) + " CPUs")
	server.host = config.Host
	server.port = config.Port
	server.full_address, _ = net.ResolveTCPAddr("tcp", server.host + ":" + server.port)
	fmt.Println("Server variables have been inited.")
}

func parseRequest(conn net.Conn) (request utils.Request, err error)  {
	return request, err
}

func makeResponse() utils.Response {
	response := new(utils.Response)
	return *response
}

func serveConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connected on " + conn.RemoteAddr().String())


	_, err := parseRequest(conn)

	if err != nil {
		return
	}

	_ = makeResponse()
	//conn.Write()
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
			go serveConnection(conn)
		}
	}

}
