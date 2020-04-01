/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-07-23
 */
package common

import (
	"time"

	"v2ray.com/core/app/porter/config"
	"v2ray.com/core/app/porter/src/utils/log"
)

func CheckLogFileSize() {
	ti := time.NewTicker(time.Minute)
	for {
		select {
		case <-ti.C:
			isNeedNewFile := log.CheckIfNeedNewFile()
			if isNeedNewFile {
				log.ClosePrintLog()
				log.InitLog(int(config.Parameters.LogLevel), GetLogDir(), log.Stdout)
			}
		}
	}
}
