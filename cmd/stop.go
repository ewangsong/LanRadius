package cmd

import (
	"io/ioutil"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Lanradius",
	Run: func(cmd *cobra.Command, args []string) {
		strb, _ := ioutil.ReadFile("/Users/waso/go/src/radiusweb/cmd/Lanradius-admin.pid")
		strb1, _ := ioutil.ReadFile("/Users/waso/go/src/radiusweb/cmd/Lanradius-radiusct.pid")
		command := exec.Command("kill", string(strb), string(strb1))
		command.Start()
		println("lanradius stop")
	},
}
