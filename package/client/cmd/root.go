package cmd

import (
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/client"
	"github.com/YangXingHao96/DistributedSystem/package/common"
	"github.com/manifoldco/promptui"
	"os"

	"github.com/spf13/cobra"
)

type cmdsPromptFormat struct {
	Prompt func() ([]byte, error)
	Fmt func(resp map[string]string)
}

var supportedCmds = map[string]cmdsPromptFormat{
	"Query flight identifier(s) by specifying the source and destination places": {
		Prompt: promptGetFlightIdBySourceDest,
		Fmt: fmtGetFlightIdBySourceDest,
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
	fmt.Println("Welcome to THE flight application!")
	conn := client.MustInit()
	defer conn.Close()

	var cmds []string
	for k := range supportedCmds {
		cmds = append(cmds, k)
	}
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

		// read and write to server
		if _, err := conn.Write(data); err != nil {
			panic(err)
		}
		buffer := make([]byte, 1024)
		mLen, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		resp := common.Deserialize(buffer[:mLen])
		supportedCmds[c].Fmt(resp)
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
