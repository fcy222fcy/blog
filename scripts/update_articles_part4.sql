-- 继续补充文章内容
USE blog;

-- 文章159: 我的 NAS 搭建历程
UPDATE articles SET content = '## 为什么需要 NAS？

作为一名开发者和数码爱好者，我有很多数据需要存储和管理：
- 工作项目代码和文档
- 个人照片和视频
- 电影和音乐收藏
- 各种备份数据

云存储虽然方便，但长期成本高，而且隐私无法保证。于是决定搭建自己的 NAS。

## 硬件选择

### 方案对比

| 方案 | 优点 | 缺点 |
|-----|------|------|
| 成品 NAS（群晖等） | 稳定、易用 | 价格高、配置固定 |
| DIY 组装 | 性价比高、可定制 | 需要技术能力 |
| 旧电脑改造 | 成本最低 | 功耗可能较高 |

### 我的选择
最终选择了 DIY 方案：
- 主板：ITX 主板
- CPU：Intel N100（低功耗）
- 内存：16GB
- 硬盘：4TB x 2（RAID 1）
- 机箱：NAS 专用机箱
- 电源：小功率电源

## 系统选择

### 常见 NAS 系统
- TrueNAS：功能强大，适合进阶用户
- Unraid：灵活易用，社区活跃
- OpenMediaVault：开源免费，插件丰富
- 黑群晖：功能最全，但有法律风险

### 我的选择
选择了 OpenMediaVault，原因：
- 完全免费开源
- 基于 Debian，稳定可靠
- Web 界面友好
- 插件丰富

## 搭建过程

### 1. 安装系统
```bash
# 下载镜像
wget https://downloads.sourceforge.net/project/openmediavault/install/images/omv6.img.xz

# 写入 U 盘
xzcat omv6.img.xz | sudo dd of=/dev/sdX bs=4M status=progress
```

### 2. 基础配置
- 设置 IP 地址
- 配置管理员账户
- 更新系统

### 3. 存储配置
- 创建文件系统
- 配置 RAID
- 设置共享文件夹

### 4. 服务配置
- SMB/NFS 共享
- Docker 容器
- 定时备份

## 常用服务

### 文件共享
通过 SMB 协议，可以像访问本地磁盘一样访问 NAS。

### Docker
NAS 上运行了多个 Docker 容器：
- Portainer：容器管理
- Plex：媒体服务
- Nextcloud：私有云盘
- Portainer：监控面板

### 自动备份
- 手机照片自动备份
- 电脑重要文件定时同步
- 云端数据定期下载

## 使用体验

### 优点
- 数据完全自己掌控
- 长期成本低
- 可以自由扩展功能

### 缺点
- 需要自己维护
- 断电断网时无法访问
- 初期投入较高

## 经验教训

1. 一定要做好备份（3-2-1 原则）
2. 选择合适的 RAID 方案
3. 注意散热和噪音
4. UPS 是必要的投资

## 总结

NAS 搭建是一个折腾但有趣的过程。现在我已经离不开它了，它已经成为我数字生活的中心。' WHERE id = 159;

-- 文章160: 程序员的健康生活
UPDATE articles SET content = '## 为什么关注健康？

程序员是健康问题的高发群体：久坐、熬夜、用眼过度、颈椎腰椎问题...是时候关注一下自己的健康了。

## 久坐问题

### 危害
- 肥胖
- 心血管疾病
- 颈椎腰椎问题
- 血液循环不畅

### 解决方案
1. 每小时起来活动 5-10 分钟
2. 使用站立式办公桌
3. 做简单的拉伸运动
4. 使用番茄工作法

### 我的做法
设置了每小时提醒，起来走动一下，做几个拉伸动作。

## 用眼健康

### 问题
- 长时间盯着屏幕
- 蓝光伤害
- 干眼症

### 解决方案
1. 20-20-20 法则：每 20 分钟看 20 英尺外 20 秒
2. 调整屏幕亮度和色温
3. 使用防蓝光眼镜
4. 保持适当的观看距离

## 颈椎腰椎保护

### 预防措施
1. 保持正确的坐姿
2. 使用人体工学椅
3. 调整显示器高度
4. 定时做颈椎操

### 锻炼方法
1. 颈部环绕运动
2. 肩部耸肩运动
3. 腰部扭转运动
4. 猫牛式伸展

## 运动建议

### 有氧运动
- 跑步
- 游泳
- 骑行
- 快走

### 力量训练
- 俯卧撑
- 深蹲
- 平板支撑
- 引体向上

### 运动频率
- 每周至少 3 次
- 每次 30 分钟以上
- 工作日午休时间也可以运动

## 饮食健康

### 注意事项
1. 规律饮食，不要不吃早餐
2. 少吃外卖，多自己做饭
3. 多喝水，少喝饮料
4. 控制咖啡因摄入
5. 多吃蔬菜水果

### 我的改变
- 开始自己做早餐
- 带午餐去公司
- 下午茶用水果代替零食

## 睡眠质量

### 提高睡眠质量的方法
1. 保持规律的作息时间
2. 睡前不要玩手机
3. 保持卧室黑暗安静
4. 避免睡前剧烈运动

### 睡眠时间
- 保证 7-8 小时睡眠
- 午休 20-30 分钟

## 心理健康

### 压力管理
- 学会说不
- 培养工作外的兴趣
- 与朋友家人交流
- 必要时寻求专业帮助

### 保持积极
- 设定合理的目标
- 庆祝小成就
- 保持学习和成长

## 总结

健康是 1，其他都是 0。没有健康，一切都没有意义。

从今天开始，关注自己的健康。' WHERE id = 160;

-- 文章161: Redis 缓存实战：从入门到精通
UPDATE articles SET content = '## 什么是 Redis？

Redis（Remote Dictionary Server）是一个开源的内存数据结构存储系统，可用作数据库、缓存和消息中间件。

### 核心特点
- 基于内存，读写速度快
- 支持多种数据结构
- 支持持久化
- 支持集群和主从复制

## 数据结构

### String
最基本的类型，可以存储字符串、数字。

```bash
SET key value
GET key
INCR counter
```

### Hash
适合存储对象。

```bash
HSET user:1 name "Tom" age 25
HGET user:1 name
HGETALL user:1
```

### List
有序的字符串列表。

```bash
LPUSH queue task1
RPUSH queue task2
LPOP queue
LRANGE queue 0 -1
```

### Set
无序的字符串集合。

```bash
SADD tags "go" "redis" "docker"
SMEMBERS tags
SISMEMBER tags "go"
```

### Sorted Set
有序的字符串集合，每个元素关联一个分数。

```bash
ZADD leaderboard 100 player1
ZADD leaderboard 200 player2
ZRANGE leaderboard 0 -1 WITHSCORES
```

## 实际应用场景

### 1. 缓存

```go
func GetUser(id int) (*User, error) {
    key := fmt.Sprintf("user:%d", id)

    // 先查缓存
    cached, err := redis.Get(key).Result()
    if err == nil {
        var user User
        json.Unmarshal([]byte(cached), &user)
        return &user, nil
    }

    // 缓存未命中，查数据库
    user, err := db.GetUser(id)
    if err != nil {
        return nil, err
    }

    // 写入缓存
    data, _ := json.Marshal(user)
    redis.Set(key, string(data), time.Hour)

    return user, nil
}
```

### 2. 分布式锁

```go
func AcquireLock(lockKey string, timeout time.Duration) (bool, error) {
    result, err := redis.SetNX(lockKey, "1", timeout).Result()
    return result, err
}

func ReleaseLock(lockKey string) error {
    return redis.Del(lockKey).Err()
}
```

### 3. 计数器

```go
func IncrementPageViews(articleID int) error {
    key := fmt.Sprintf("article:%d:views", articleID)
    return redis.Incr(key).Err()
}

func GetPageViews(articleID int) (int64, error) {
    key := fmt.Sprintf("article:%d:views", articleID)
    return redis.Get(key).Int64()
}
```

### 4. 排行榜

```go
func UpdateScore(userID string, score float64) error {
    return redis.ZAdd("leaderboard", &redis.Z{
        Score:  score,
        Member: userID,
    }).Err()
}

func GetTopUsers(top int) ([]redis.Z, error) {
    return redis.ZRevRangeWithScores("leaderboard", 0, int64(top-1)).Result()
}
```

## 缓存策略

### 缓存穿透
问题：查询不存在的数据，每次都会查数据库。
解决：
1. 缓存空值
2. 使用布隆过滤器

### 缓存击穿
问题：热点 key 过期，大量请求同时查数据库。
解决：
1. 使用互斥锁
2. 热点 key 永不过期

### 缓存雪崩
问题：大量 key 同时过期，数据库压力骤增。
解决：
1. 过期时间加随机值
2. 使用集群提高可用性

## 性能优化

### 1. 使用 Pipeline

```go
pipe := redis.Pipeline()
for i := 0; i < 100; i++ {
    pipe.Set(fmt.Sprintf("key:%d", i), "value", 0)
}
pipe.Exec()
```

### 2. 合理使用数据结构
选择合适的数据结构可以提高效率。

### 3. 控制 key 大小
key 越小，占用内存越少，查找越快。

## 总结

Redis 功能强大，但要根据实际场景合理使用。理解数据结构、掌握缓存策略、注意性能优化，才能发挥 Redis 的最大价值。' WHERE id = 161;
