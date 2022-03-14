/**
 * @Author: guobob
 * @Description:
 * @File:  cache.go
 * @Version: 1.0.0
 * @Date: 2022/3/12 17:29
 */

package cache

import (
	//"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func InitCache() *cache.Cache{
	return cache.New(5*time.Minute, 10*time.Minute)
}
