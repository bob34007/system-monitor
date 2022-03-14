/**
 * @Author: guobob
 * @Description:
 * @File:  cpu.go
 * @Version: 1.0.0
 * @Date: 2022/3/7 21:28
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/system-monitor/cpu"
	"time"
)

func NewCpuCommand() *cobra.Command {
	//Collect and analyze cpu status
	var (

	)
	cmd := &cobra.Command{
		Use:   "cpu",
		Short: "Collect and analyze cpu status ",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			ts := time.Now()
			fmt.Println(ts)
			c := cpu.NewCPUS()
			c.SetLoadAvg()
			c.SetInfo()
			c.SetProcessStatusCount()

			fmt.Println(c.Print())
			return err
		},
	}
	return cmd
}