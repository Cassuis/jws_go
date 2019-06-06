# 掉落TODO
--------------
## 1.掉落算法优化 [T623](http://wiki.taiyouxi.net/T623?workflow=534)
作为服务器开发，需要对掉落算法进行优化

###1.1. 建立一个敌兵掉落表的池，避免每次都重新计算掉落池
现在的逻辑中当玩家开始进入一个副本时，服务器都会计算一个掉落表，当结束副本时
服务器根据客户端的报告，根据掉落表给玩家发送掉落物品，这个过程当中为了应对玩家掉线问题，
服务器会在Redis中缓冲掉落表信息。

实际上，对于玩家来说，并不需要每次都重建掉落表，可以建立一个玩家间共享的掉落表池，
每次玩家请求时，从掉落池中随机返回一个掉落表

###1.2. 建立好掉落池后，tmp信息中只需要记录池中的索引即可
对于玩家来说，掉落表池是不变的，所以Redis缓存中只需要存储掉落表池的索引即可

###1.3. 掉落池应该足够大，因为考虑到玩家副本过程中已经显示掉落数据，所以掉落池本身不能经常更新
客户端在获取掉落表之后，会在副本结算之前提前告知玩家掉落内容，这使得掉落表池必须在玩家
整个副本期间必须不变，考虑到玩家掉线问题，服务器很难知道什么时间可以修改掉落表，所以
只能在服务器重启时重建掉落表，为了保证玩家感受上相信副本掉落确实是随机的，这个掉落表池应该足够大
具体多大要根据玩家每天副本最大次数来确认

###1.4. 需要整理生成掉落表的算法实现的结构，现有代码太乱了


###1.5. 客户端反序列化掉落表算法也需要整理