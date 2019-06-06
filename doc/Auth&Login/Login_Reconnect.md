**玩家掉线：**

玩家有很多情况下都可能掉线，比如3G信号不稳定或者接电话是我们最需要考虑的情况。根据这个情况，划分为临时掉线和彻底掉线。
假设所有掉线都进入临时掉线状态，然后经过指定之间计时后自动进入彻底掉线模式。
彻底掉线后，游戏服务器应该清理掉 {user_id}@{shard}@{gate}@subid:handshakeid:randomkey，若客户端回来则会触发客户端进行Auth阶段认证和Login登录.

如果客户端想重复和游戏服务器建立连接，它不需要再次去登陆服务器登陆。只需要把上次的 handshakeid 递增，并重新生成一个 randomkey ，去和游戏服务器（或者Gate）握手即可。
小于当前handshakeid的请求都不进行处理。

如何判断临时掉线状态发生？玩家心跳超时是玩家掉线的唯一标志！
玩家主动退出怎么处理？如果客户端有主动退出功能，应该由客户端主动发出logout消息给游戏服务器。



**客户端链接握手和断线链接的过程** TODO


**基础方案**

客户端和服务器建立链接后，第一条信息发送文本"loginToken\n"，否则断开链接。

**进阶方案**

参考：
 - http://blog.codingnow.com/2014/02/connection_reuse.html
 - 代码实现参考https://github.com/xjdrew/goscon.git

连接建立过程如下：

1. 客户端向服务器发起一个 TCP 连接，并发送第一个握手信号：包含一个 0 和按 Diffie-Hellman 密钥交换算法产生一串随机量 A 。
2. 服务器收到第一个握手信号后，检查第一个字段，若不为 0 则进入连接修复阶段 2.2，否则继续创建连接过程 2.1
    - 2.1 此时服务器生成另一个随机串 B ，并通过 DH 密钥交换算法得到了一个 secret 。此时回应 DH 算法需要的 B ，以及一个新的随机串 E （用于校验）。
    - 2.2 当第一个字段不为 0 则认为是需要修复一个已有连接，这个数字表示在旧连接上客户端已收到过服务器发过来的数据包数量。此时，第二个字段应理解为旧连接上已收到数据包的指纹。服务器根据包数量和指纹可以核对所有保持的有效连接，如果不能找到匹配的连接（指纹相同），就断开客户端。否则回应客户端在旧连接上一共收到客户端发送的数据包数量，以及一个新的随机串 E ，用于确认客户端是否知道旧连接的 secret 。（这个校验是有必要的，否则会有人监听到链接重建过程，而重复发送这个握手包来踢掉合法用户刚修复的连接）
3. 客户端收到随机串 E 后，和 secret 连接在一起做一次 hash (可以使用 md5 算法) H，回应服务器。这可以让服务器校验客户端是否真的拥有 secret 。
4. 服务器收到二次握手信号 H 后，用同样的 hash 算法做一次 secret 校验，确认是合法的客户端后继续通讯；若是非法连接则立即断开。
5. 链接处理
    - 5.1 如果是新连接，那么服务器利用得到的 secret 初始化 RC4 加密算法需要的 s-box ，之后的通讯利用 RC4 算法加密。
    - 5.2 如果是旧连接修复，那么服务器将客户端未收到的数据包重发一次。并从旧通道上复制 RC4 所用的 s-box 以及 secret 用于后续通讯。
6. 此后的每次数据通讯，在数据打包后，都利用 RC4 算法做一次加密，并利用数据更新数据指纹（可以用加密后的数据流的 CRC 值）。每个数据包都记录当前的指纹，并 cache 最近发送的 128 个数据包用于事后的连接修复。
7. 设定一个超时时间，定期清理没有数据来往的 TCP 连接。

这个协议的好处是，客户端在握手完成后，任何时间都可以向服务器发起一个新的 TCP 连接取代旧的连接（无须利用旧连接是否还有效），而对应用层来说，连接重来没有中断过。

