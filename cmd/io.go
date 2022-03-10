/**
 * @Author: guobob
 * @Description:
 * @File:  io.go
 * @Version: 1.0.0
 * @Date: 2022/3/7 21:29
 */

package cmd


import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)
func NewIOCommand() *cobra.Command {
	//Collect and analyze io status
	var (

	)
	cmd := &cobra.Command{
		Use:   "io",
		Short: "Collect and analyze io status ",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			ts := time.Now()
			fmt.Println(ts)
			return err
		},
	}
	return cmd
}
