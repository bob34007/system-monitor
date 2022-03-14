/**
 * @Author: guobob
 * @Description:
 * @File:  comand.go
 * @Version: 1.0.0
 * @Date: 2022/3/12 18:24
 */

package utils

import (
	"context"
	"strings"
)


var invoke Invoker = Invoke{}

func addItemCountToMap(m map[string]int,itemType string ){
	if v ,ok := m[itemType];!ok {
		m[itemType]=1
	} else{
		m[itemType]=v+1
	}
}


//GetProcessStatusCount Run the ps axo status command to obtain the
//number of processes corresponding to each system status
func GetProcessStatusCount(m map[string]int) error {
	out, err := invoke.CommandWithContext(context.Background(), "ps", "axo", "state")
	if err != nil {
		return  err
	}
	lines := strings.Split(string(out), "\n")

	for _, l := range lines {
		addItemCountToMap(m,strings.TrimSpace(l))
	}
	return nil
}


