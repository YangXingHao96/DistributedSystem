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
	Fmt func(resp map[string]interface{}) string
}

var supportedCmds = map[string]cmdsPromptFormat{
	"Query flight identifier(s) by specifying the source and destination places": {
		Prompt: promptGetFlightIdBySourceDest,
		Fmt: fmtGetFlightIdBySourceDest,
	},
	"Query flight details by its ID": {
		Prompt: promptGetFlightDetail,
		Fmt: fmtGetFlightDetail,
	},
	"Add a flight": {
		Prompt: promptAddFlight,
		Fmt: fmtSimpleAck,
	},
	"Make a flight reservation": {
		Prompt: promptMakeReservation,
		Fmt: fmtSimpleAck,
	},
	"Cancel a flight reservation": {
		Prompt: promptCancelReservation,
		Fmt: fmtSimpleAck,
	},
	"Check your reservation for a flight": {
		Prompt: promptGetReservationForFlight,
		Fmt: fmtGetReservationForFlight,
	},
	"Subscribe to a flight for live update on seat availability": {
		Prompt: promptRegisterMonitorReq,
		Fmt: fmtMonitorFlightResp,
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
		s.Start()
		// read and write to server
		if _, err := conn.Write(data); err != nil {
			panic(err)
		}
		keepReading := true
		for keepReading {
			keepReading = false
			buffer := make([]byte, 1024)
			mLen, err := conn.Read(buffer)
			if err != nil {
				if isTimeoutError(err) {
					s.FinalMSG = fmt.Sprintf("\nwarning: connection to server timed out, exceeded threshold: %d (ms).\n", client.ReadTimeoutMs)
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
}

func isTimeoutError(err error) bool {
	e, ok := err.(net.Error)
	return ok && e.Timeout() || strings.Contains(err.Error(), "i/o timeout")
}
