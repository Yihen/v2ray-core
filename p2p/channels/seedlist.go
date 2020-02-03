/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2019-10-16
 */
package channels

import "v2ray.com/core/p2p/wire/pb/seedlist/types"

const CHANNEL_MAX_SIXE uint16 = 256

var SeedlistNotice chan types.HelloSeedList
