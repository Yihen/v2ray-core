/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2020-04-08
 */
package rpc

import (
	"net/http"
	"strconv"

	"fmt"

	"v2ray.com/core/app/rpc/base"
)

func StartLocalRPCServer(port int) error {
	http.HandleFunc("/", base.Handle)
	//rpc.HandleFunc("getsysstatusscore", rpc.GetSysStatusScore)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		return fmt.Errorf("ListenAndServe error:%s", err)
	}
	return nil
}
