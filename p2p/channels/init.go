/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-10-16
 */
package channels

import "v2ray.com/core/p2p/wire/pb/seedlist/types"

func init() {
	SeedlistNotice = make(chan types.HelloSeedList, CHANNEL_MAX_SIXE)
}
