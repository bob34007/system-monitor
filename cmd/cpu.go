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
			return err
		},
	}
	return cmd
}