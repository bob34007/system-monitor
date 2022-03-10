/**
 * @Author: guobob
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2022/3/7 14:58
 */

package main

import (
	"os"

	"github.com/soccer-crawler/cmd"
	"go.uber.org/zap"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		zap.L().Error("error exit: "+err.Error(), zap.Error(err))
		os.Exit(1)
	}
}
