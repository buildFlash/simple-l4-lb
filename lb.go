package main

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net"
)

type Backend struct {
	Host        string
	Port        int
	IsHealthy   bool
	numRequests int
}

func (b Backend) String() string {
	return fmt.Sprintf("%s:%d", b.Host, b.Port)
}

type Event struct {
	EventName string
	Data      interface{}
}

type LB struct {
	backends []*Backend
	events   chan Event
	strategy BalancingStrategy
}

type IncomingReq struct {
	srcConn net.Conn
	reqId   string
}

var lb *LB

func InitLB() {
	backends := []*Backend{
		{Host: "localhost", Port: 8080, IsHealthy: true},
		{Host: "localhost", Port: 8081, IsHealthy: true},
		{Host: "localhost", Port: 8082, IsHealthy: true},
		{Host: "localhost", Port: 8083, IsHealthy: true},
	}
	lb = &LB{
		backends: backends,
		events:   make(chan Event),
		strategy: NewConsistentHashingStrategy(backends, 3),
	}
}

func (lb *LB) Run() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	go handleEvents()

	defer listener.Close()
	log.Println("LB listening on 9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go lb.proxy(IncomingReq{
			srcConn: conn,
			reqId:   uuid.New().String(),
		})
	}

}

func (lb *LB) proxy(req IncomingReq) {
	backend := lb.strategy.GetNextBackend(req)
	log.Printf("[in-req] %s out-req %s \n", req.reqId, backend.String())

	backendConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", backend.Host, backend.Port))
	if err != nil {
		req.srcConn.Write([]byte("Backend not available"))
		req.srcConn.Close()
		panic(err)
	}

	backend.numRequests++
	go io.Copy(req.srcConn, backendConn)
	go io.Copy(backendConn, req.srcConn)
}

func handleEvents() {
	for event := range lb.events {
		switch event.EventName {
		case CMD_StrategyChange:
			if strategyStr, ok := event.Data.(string); ok {
				if strategyStr == "CH" {
					if _, ok := lb.strategy.(*ConsistentHashingStrategy); !ok {
						lb.strategy = NewConsistentHashingStrategy(lb.backends, 3)
					}
				} else if strategyStr == "RR" {
					if _, ok := lb.strategy.(*RRBalancingStrategy); !ok {
						lb.strategy = NewRRBalancingStrategy(lb.backends)
					}
				}
			}
			break
		case CMD_BackendAdd:
			if backend, ok := event.Data.(Backend); ok {
				lb.backends = append(lb.backends, &backend)
				lb.strategy.RegisterBackend(&backend)
			}
			break
		case CMD_Exit:
			// handle exit here
			return
		}
	}
}
