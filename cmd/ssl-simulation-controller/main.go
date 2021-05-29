package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RoboCup-SSL/ssl-simulation-controller/internal/simctl"
)

var visionAddress = flag.String("visionAddress", "224.5.23.2:10020", "The address (ip+port) from which vision packages are received")
var trackerAddress = flag.String("trackerAddress", "224.5.23.2:10010", "The address (ip+port) from which tracker packages are received")
var refereeAddress = flag.String("refereeAddress", "224.5.23.1:10003", "The address (ip+port) from which referee packages are received")
var simControlPort = flag.String("simControlPort", "10300", "The port to which simulation control packets are send")
var robotSpecConfig = flag.String("robotSpecConfig", "config/robot-specs.yaml", "The robot specs config file")

func main() {
	log.Println("Baguncinha")
	flag.Parse()
	ctl := simctl.NewSimulationController(*visionAddress, *refereeAddress, *trackerAddress, *simControlPort, *robotSpecConfig)
	ctl.Start()
	log.Println("Done")

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("e", text) == 0 {
			ctl.Stop()
			break
		}

		words := strings.Fields(text)
		if strings.Compare("b", words[0]) == 0 {
			x, _ := strconv.ParseFloat(words[1], 32)
			y, _ := strconv.ParseFloat(words[2], 32)
			ctl.ReplaceBall(float32(x), float32(y))
		} else if strings.Compare("r", words[0]) == 0 {
			color := 0
			if strings.Compare("y", words[1]) == 0 {
				color = 1
			} else if strings.Compare("b", words[1]) == 0 {
				color = 2
			} else {
				color = 0
			}
			id, _ := strconv.ParseUint(words[2], 10, 32)
			x, _ := strconv.ParseFloat(words[3], 32)
			y, _ := strconv.ParseFloat(words[4], 32)
			ctl.ReplaceRobot(int32(color), uint32(id), float32(x), float32(y))
			// ctl.ReplaceBall(float32(x), float32(y))
		}
	}
}
