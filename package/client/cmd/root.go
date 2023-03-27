package cmd

import (
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/client"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/YangXingHao96/DistributedSystem/package/common/constant"
	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"math/rand"
	"net"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type cmdsPromptFormat struct {
	Prompt func() ([]byte, error)
	Fmt    func(resp map[string]interface{}) string
}

const (
	queryFlightPromptStr     = "Query flight identifier(s) by specifying the source and destination places"
	queryFlightByIdStr       = "Query flight details by its ID"
	addFlightPromptStr       = "Add a flight"
	makeResPromptStr         = "Make a flight reservation"
	cancelResPromptStr       = "Cancel a flight reservation"
	checkResPromptStr        = "Check your reservation for a flight"
	registerMonitorPromptStr = "Subscribe to a flight for live update on seat availability"
)

var supportedCmds = map[string]cmdsPromptFormat{
	queryFlightPromptStr: {
		Prompt: promptGetFlightIdBySourceDest,
		Fmt:    fmtGetFlightIdBySourceDest,
	},
	queryFlightByIdStr: {
		Prompt: promptGetFlightDetail,
		Fmt:    fmtGetFlightDetail,
	},
	addFlightPromptStr: {
		Prompt: promptAddFlight,
		Fmt:    fmtSimpleAck,
	},
	makeResPromptStr: {
		Prompt: promptMakeReservation,
		Fmt:    fmtSimpleAck,
	},
	cancelResPromptStr: {
		Prompt: promptCancelReservation,
		Fmt:    fmtSimpleAck,
	},
	checkResPromptStr: {
		Prompt: promptGetReservationForFlight,
		Fmt:    fmtGetReservationForFlight,
	},
	registerMonitorPromptStr: {
		Prompt: promptRegisterMonitorReq,
		Fmt:    fmtMonitorFlightResp,
	},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Client side for a UDP Flight Application",
	Long:  "An interactive CLI acting as a client side for a UDP Flight Application",
	Run:   start,
}

func start(cmd *cobra.Command, args []string) {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to THE flight application!")
	conn := client.MustInit()
	defer conn.Close()

	var cmds []string
	for k := range supportedCmds {
		cmds = append(cmds, k)
	}
	sort.Strings(cmds)
	for {
		prompt := promptui.Select{
			Label: "Select an action",
			Items: cmds,
		}

		_, c, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if _, exist := supportedCmds[c]; !exist {
			fmt.Printf("chosen command %s is not supported", c)
			continue
		}
		data, err := supportedCmds[c].Prompt()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		s := spinner.New(spinner.CharSets[rand.Intn(44)], 100*time.Millisecond)
		s.Prefix = "working on it... "
		s.Color("blue")
		// measure execution time
		start := time.Now()
		readDuration := client.ReadTimeoutMs
		if c == registerMonitorPromptStr {
			monitorDurationSec := common.Deserialize(data)[constant.MonitorIntervalSec].(int)
			if monitorDurationSec*1000 > readDuration { // client allowed to read longer than default
				readDuration = monitorDurationSec * 1000
			}
		}
		readDeadline := time.Now().Add(time.Duration(readDuration) * time.Millisecond)
		s.Start()
		// read and write to server
		if _, err := conn.Write(data); err != nil {
			panic(err)
		}
		keepReading := true
		for keepReading {
			keepReading = false
			buffer := make([]byte, 1024)
			mLen, err := conn.Read(buffer, readDeadline)
			if err != nil {
				if isTimeoutError(err) {
					s.FinalMSG = fmt.Sprintf("\nConnection to server ended, exceeded read deadline: %d (ms).\n", readDuration)
					break
				}
				fmt.Println("Error reading:", err.Error())
				return
			}
			resp := common.Deserialize(buffer[:mLen])
			msgType := resp[constant.MsgType].(int)
			if msgType == constant.GeneralErrResp {
				s.FinalMSG = fmtGeneralErrResp(resp)
				continue
			}

			respFmtOutput := supportedCmds[c].Fmt(resp)

			s.FinalMSG = fmt.Sprintf("\ndone.\n============ Output ============\n%v\n", respFmtOutput)
			if msgType == constant.MonitorUpdateResp {
				keepReading = true
				fmt.Println(respFmtOutput)
				s.FinalMSG = ""
			}
		}
		executionTime := time.Since(start)
		s.Stop()
		fmt.Printf("🧰 Total execution time: %d ms\n", executionTime.Milliseconds())
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&client.ServerHost, "host", "localhost", "the server host (defaults to localhost)")
	rootCmd.PersistentFlags().StringVar(&client.ServerPort, "port", "2222", "the server port (defaults to 2222)")
	rootCmd.PersistentFlags().IntVar(&client.ReadTimeoutMs, "readTimeout", 60000, "the server port (defaults to 2222)")
}

func isTimeoutError(err error) bool {
	e, ok := err.(net.Error)
	return ok && e.Timeout() || strings.Contains(err.Error(), "i/o timeout")
}
