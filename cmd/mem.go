/**
 * @Author: guobob
 * @Description:
 * @File:  mem.go
 * @Version: 1.0.0
 * @Date: 2022/3/7 21:30
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func NewMemCommand() *cobra.Command {
	//Collect and analyze memory status
	var (

	)
	cmd := &cobra.Command{
		Use:   "mem",
		Short: "Collect and analyze memory status",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			ts := time.Now()
			fmt.Println(ts)
			return err
		},
	}
	return cmd
}