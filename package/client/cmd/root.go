package cmd

import (
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/client"
	"github.com/manifoldco/promptui"
	"os"

	"github.com/spf13/cobra"
)

var supportedCmds = map[string]func(conn client.UdpConn) error{
	"Query flight identifier(s) by specifying the source and destination places": getFlightIdBySourceDest,
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

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if _, exist := supportedCmds[result]; !exist {
			fmt.Printf("chosen command %s is not supported", result)
			continue
		}
		if err := supportedCmds[result](conn); err != nil {
			fmt.Println(err)
			continue
		}
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
