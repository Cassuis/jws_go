# 用DynamoDB实现排行榜能力的研究
---------------------

## 1.Why DynamoDB？

相比Redis，DynamoDB节约运维成本、可以免去落地的问题，
如果排行榜可以用DynamoDB实现，可以省去很多麻烦。

Redis提供了有序集合这一数据结构，是实现排行榜的比较好的选择，但是因为不同的排行榜需求是不一样的，所以不见得一定要用有序集合。

下面分析几种常见的排行榜，研究其怎样用DynamoDB实现

## 2. 排位榜-固定打榜（典型例子：乱斗西游）

### 需求

- 玩家知道自己的排名
- 玩家知道前N名排名是谁
- 玩家知道自己排名的前P名（或者是前P名中固定相对位置的几名，例如300名的玩家可以看到300-1， 300-2，300-10，300-20这几位上的人）和后P名是谁
- 排位变化规则：玩家挑战成功，系统交换两个位置上的玩家

这种需求中对排名之间的关系没有要求，操作只针对特定排名，所以在DynamoDB上可以建立两个表.

表1 排位->玩家Id 通过这个可以查询玩家前N名后N名是谁

表2 玩家Id->排位 通过这个知道玩家排名

排位变更时更新表1和表2，这里需要注意操作是针对位置，所以先更新表1获取两个变化位置的玩家Id，再按照Id更新表2

表3 前N名的具体信息 --> 这个直接在服务器内存上做缓存即可，有变动时更新

## 3. 排位榜-随机打榜（典型例子：刀塔传奇竞技场）

### 需求

- 玩家知道自己的排名
- 玩家知道前N名排名是谁
- 假定玩家排名为N，对N1>N2>N3玩家需要从排名N-N1、排名N1-N2、排名N2-N3中随机各选择一个排位上的玩家进行挑战
- 排位变化规则：玩家挑战成功，系统交换两个位置上的玩家

表1 排位->玩家Id 通过这个交换玩家排位

表2 玩家Id->排位 通过这个知道玩家排名

表3 排位范围->排位范围中的玩家 用作随机池，随机时找一个或者几个池来随出玩家，
注意为了减少读取，一次随机应该随机出多个候选玩家，客户端不要每次都请求

排位变更时更新表1和表2，这里需要注意操作是针对位置，所以先更新表1获取两个变化位置的玩家Id，再按照Id更新表2

表3 前N名的具体信息 --> 这个直接在服务器内存上做缓存即可，有变动时更新

## 4. 排行榜-积分单调增长（典型例子：刀塔传奇各种其他榜）

### 需求

- 玩家知道自己的排名
- 玩家知道前N名排名是谁
- 排名依据的积分只增不减

表1 前N名的具体信息和前N名最小积分 --> 这个直接在服务器内存上做缓存即可，有变动时更新，当玩家变动后的积分超过最小积分时将其放入前N名缓存中。

表2 积分 -> 该积分的玩家的排名
具体算法见 http://blog.codingnow.com/2014/03/mmzb_db_2.html

## 5. 排行榜-积分可增可减（典型例子：Dota 2天梯积分榜）

### 需求

- 玩家知道自己的排名
- 玩家知道前N名排名是谁
- 排名依据的积分可增可减

这个似乎只能用redis的有序集合，非常遗憾，DynamoDB不适合实现这种榜

## 6. 其他需求 待补充

### 根据当前排名给玩家发奖
注意玩家不一定在线，所以需要遍历所有榜中的玩家，对于4中的榜，可以建一个类似于2、3中表1的表，进行遍历操作。

