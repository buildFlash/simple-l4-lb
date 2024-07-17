package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const CMD_StrategyChange = "strategy/change"

// const CMD_StrategyEdit = "strategy/edit"
const CMD_BackendAdd = "backend/add"
const CMD_TopologyList = "topology/list"
const CMD_TopologyTest = "topology/test"
const CMD_Exit = "exit"

var commands []string = []string{
	CMD_StrategyChange,
	CMD_BackendAdd,
	CMD_TopologyList,
	CMD_TopologyTest,
	CMD_Exit,
}

func cli() {
	for {
		var command string
		fmt.Print(">>>>>> ")
		fmt.Scanf("%s", &command)

		switch command {
		case CMD_Exit:
			lb.events <- Event{
				EventName: CMD_Exit,
			}
			// todo: this is not ideal. End this gracefully

		case CMD_BackendAdd:
			var host string
			var port int

			fmt.Print("       Host: ")
			fmt.Scanf("%s", &host)

			fmt.Print("       Port: ")
			fmt.Scanf("%d", &port)
			lb.events <- Event{
				EventName: CMD_BackendAdd,
				Data: Backend{
					Host: host,
					Port: port,
				},
			}

		case CMD_StrategyChange:
			var strategy string

			fmt.Print("       Name of strategy: ")
			fmt.Scanf("%s", &strategy)

			lb.events <- Event{
				EventName: CMD_StrategyChange,
				Data:      strategy,
			}

		case CMD_TopologyTest:
			var reqId string
			fmt.Print("       ReqId: ")
			fmt.Scanf("%s", &reqId)

			backend := lb.strategy.GetNextBackend(IncomingReq{reqId: reqId})
			fmt.Printf("req %s goes to backend %s\n", reqId, backend.String())
		case CMD_TopologyList:
			lb.strategy.PrintTopology()
		default:
			fmt.Printf("Available commands %s \n", strings.Join(commands, ","))
		}
	}
}

func setLogging() {
	f, err := os.OpenFile("lb.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file %v", err)
	}
	log.SetOutput(f)
}

func main() {
	setLogging()
	InitLB()
	go lb.Run()
	cli()
}
