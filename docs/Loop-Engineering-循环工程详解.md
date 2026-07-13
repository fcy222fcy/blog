# Loop Engineering（循环工程）详解

## 一、什么是 Loop Engineering

### 1.1 一句话定义

**Loop Engineering（循环工程）**：你不再手动一轮一轮地给 AI Agent 写提示词，而是**设计一套自动运转的闭环系统**，让这套系统按目标、按计划，自动持续地驱动 Agent 完成「发现工作 → 规划方案 → 执行任务 → 验证结果 → 沉淀状态 → 进入下一轮」的全流程，人只在关键决策节点介入。

用更直白的比喻：

| 旧模式（Prompt 模式） | 新模式（Loop 模式） |
|---|---|
| 你开车，手动转向、踩油门 | 你造了一辆自动驾驶汽车 |
| 你是 AI 的"驾驶员"，每一步都要介入 | 你是 AI 系统的"工程师"，定义目标后系统自动跑到完成 |
| 人停 AI 就停（你睡觉，AI 也"睡觉"） | 系统 7×24 按规则循环运转 |
| 核心产出是单次对话的结果 | 核心产出是可长期复用的自动化流程 |

### 1.2 起源与背景

Loop Engineering 不是凭空出现的概念，而是 AI 编程社区多位领军人物几乎同时指向的同一方向：

- **2026 年 6 月 7 日**：OpenClaw 创始人、OpenAI 开发者 **Peter Steinberger** 在 X 发帖——"你不该再手动给编码 Agent 写提示词了，你应该设计让 Agent 自己驱动自己的循环系统。"24 小时内 500 万+ 浏览量，引爆整个 AI 开发者社区。
- **同期**：Anthropic Claude Code 负责人 **Boris Cherny** 在邀请制活动上公开表态——"我现在不再自己给 Claude 写提示词了。我有循环在跑，它们负责提示 Claude 并判断下一步怎么做。我的工作，是写循环。"其个人代码已 100% 由 Claude Code 产出，每天提交 10~30 个 PR。
- **随后**：Google Cloud AI 总监 **Addy Osmani** 发表博客，将这套方法论正式命名为 **Loop Engineering**，并给出了清晰的六大组件定义框架。
- **落地项目**：开发者 **Cobus Greyling** 推出开源项目 `loop-engineering`（MIT 协议，7K+ Star，800+ Fork），提供 7 个 npm CLI 工具和 7 套生产级循环模式，把理念从理论变成可直接落地的工具集。

从数据层面看，这套方法论正在被快速验证：
- Anthropic 内部：工程师人均代码产出同比 **+200%**（约 3 倍），人均每日合并 PR 数 **+67%**。
- 公开 GitHub：约 **4%** 的提交已由 Claude Code 类工具产出，行业预测 2026 年底该比例可能达到 1/5。

---

## 二、为什么需要 Loop Engineering：三大痛点

AI 编码 Agent 的能力在快速提升，但"人跟 Agent 的协作方式"还停留在**手动挡**阶段。Loop Engineering 就是为了解决以下三个核心痛点而生。

### 痛点 1：你的时间并没有真正被解放

```
旧模式时间分配：
┌─────────────────────────────────────────────────────┐
│  写代码时间  ──→  换成了  ──→  写 prompt 的时间      │
│  调 bug 时间  ──→  换成了  ──→  调 AI 输出的时间     │
│  等编译时间  ──→  换成了  ──→  等 AI 思考的时间     │
└─────────────────────────────────────────────────────┘
      你只是把"A 类重复劳动"换成了"B 类重复劳动"
```

每一个新任务你都得重新介入——开门、打招呼、喂上下文、检查输出、给反馈、再检查……你只是把打键盘的时间，换成了写 prompt 和看输出的时间。**你的注意力仍然是整个系统的瓶颈**。

### 痛点 2：低价值任务占用了 Agent 的高价值窗口

以下工作完全可以自动跑，但因为没有"循环"，每次都得人来发起：

| 低价值但高频的任务 | 频次 | 是否值得人手动发起 |
|---|---|---|
| 扫描并分类新 Issue | 每天 N 次 | ❌ |
| CI 挂了分析日志定位根因 | 每次 PR | ❌ |
| patch 级依赖升级 + 跑测试 | 每周 | ❌ |
| 根据 git log 整理 changelog 草案 | 每次发版 | ❌ |
| 合并后清理分支、更新 Issue 状态 | 每次合并 | ❌ |
| PR 提交后检查代码风格 / 潜在安全问题 | 每次 PR | ❌ |

Loop Engineering 的思路：**把所有"规则明确、风险可控、结果可验证"的任务，全部塞进循环里自动跑**，把人的注意力留给真正需要判断力的决策。

### 痛点 3：人的决策瓶颈 = 系统的物理天花板

Agent 跑得比你快，但每跑一步都要等你确认。你开会、吃饭、睡觉的时候，Agent 也跟着"停摆"。对一个时差横跨半个地球的开源项目，或者有 nightly 构建需求的团队来说，这本质上是在用**人的生物钟限制系统的吞吐**。

Loop Engineering 解决的就是这件事：**把"你的判断"前移到系统设计阶段写进规则里**，然后让系统代替你做 24 小时的日常运转。

---

## 三、AI 工程范式的四次跃迁

Loop Engineering 不是孤立的新概念，它是 AI 工程方法论自然演进的**第四步**。前三次跃迁问的是同一个问题，只是每次都往外推了一层——从"怎么跟模型说话"一路演进到"怎么让整套系统自己转起来"。

### 3.1 四次跃迁总览

```
能力外扩方向 →
Prompt Engineering  →  Context Engineering  →  Harness Engineering  →  Loop Engineering
(2022~2024)            (2025)                  (2025~2026 初)         (2026 中起)

每层关心的问题：
  怎么问？ ──→ 给模型看什么？ ──→ 在哪个环境里行动？ ──→ 怎么持续驱动模型？
```

| 阶段 | 核心关注点 | 你的角色 | 天花板 |
|---|---|---|---|
| **Prompt Engineering** | 单轮提示词怎么写才能让模型听懂 | 提问者 | 你就是瓶颈，每一步都要手动跟进 |
| **Context Engineering** | 系统地设计 AI 看到的"信息环境"（System Prompt / Memory / 历史 / 工具定义） | 信息架构师 | 仍然是单次运行，不具备自主持续能力 |
| **Harness Engineering** | 为单次 Agent 运行配好"工作台"（工具集 / 文件权限 / 沙箱 / 测试脚本） | 环境搭建师 | 装备的是"单次执行"，跑完一次就结束 |
| **Loop Engineering** | 设计替代"你"的自动驱动系统（什么时候启动、失败怎么办、谁来打分、什么时候停下） | 系统架构师 | 人的带宽不再是每日产出的硬性上限 |

### 3.2 每一层解决了什么（用修 Bug 的例子说明）

假设你让 AI 帮你修一个博客项目的搜索 Bug：

| 阶段 | 你做的事 | 举例 |
|---|---|---|
| **Prompt** | 写一条精心设计的提问 | `"你是资深 Go 工程师，按以下报错定位 bug，给出最小修改..."` |
| **Context** | 把相关上下文系统地给它 | 报错日志 + `article_service.go` + `article_repository.go` + 复现步骤 + 项目代码规范 + 之前踩过的同类坑记录 |
| **Harness** | 给它装备行动环境 | 允许读文件、允许改 `internal/` 下 Go 代码、允许跑 `go test ./...`、禁止碰 `configs/` 与数据库、跑测试超时 120s |
| **Loop** | 设计自动运转的规则 | "**每天 09:00** 自动扫仓库 open issue，找带 `bug` 标签的 → 挑一个 `search` 相关的 → 在独立 worktree 里尝试修复 → **让另一个独立模型**重新跑测试 + 人工 Code Review 检查 → 不通过就把失败日志喂回去重改（最多 5 轮）→ 通过就自动提 PR，@ 负责人等合并" |

注意 Loop 这一层做了三件 Harness 不关心的事：**① 什么时候启动（心跳/触发）② 失败了怎么办（自动重试 + 上限）③ 结果谁来验（角色分离）**。这三件事加起来，才让 AI 从一个"你调用的工具"变成了一个"自主运转的系统"。

---

## 四、核心机制：五步闭环（DAPVI 模型）

一个完整的 Loop 工程系统，会自动完成以下五个动作，不断循环。业界普遍采用 **Discover → Assign/Plan → Act → Verify → Persist/Iterate** 的五阶段模型，简称 DAPVI。

```
┌───────────────────────────────────────────────────────────────┐
│                                                               │
│   Discover(发现) ──→ Assign(分配/规划) ──→ Act(执行)          │
│        ↑                                               ↓      │
│        │                                               ↓      │
│   Iterate(迭代) ←── Persist(持久化) ←── Verify(验证)          │
│                                                               │
│   每一轮结束：验证通过 → 沉淀经验 + 进入下一轮发现             │
│              验证失败 → 把错误反馈给 Act 重跑（最多 N 轮）      │
│                                                               │
└───────────────────────────────────────────────────────────────┘
```

### 4.1 每一步做什么

#### Step 1 - Discover（发现工作）

Loop 的"眼睛"。**主动扫描**一切可能产生任务的来源，而不是等人塞任务进来。

| 典型发现源 | 触发条件 |
|---|---|
| GitHub Issue | 新创建 / 24h 未回应 / 带特定标签 |
| GitHub PR | 新建 / 有新 review comment / CI failed |
| CI / Nightly Build | 构建失败 / 测试失败率超阈值 / 覆盖率下降 |
| 依赖仓库 | 有新版本发布 / 有 CVE 安全公告 |
| 生产日志 / 监控 | 错误率飙升 / P99 延迟超阈值 / 特定错误日志匹配 |
| 定期 cron | 每日 09:00 分类 triage / 每周一跑依赖升级 / 每月整理 changelog |

#### Step 2 - Assign & Plan（分配 & 规划）

Loop 的"大脑"。拿到待办后：

1. **Assign**：判断这个任务给谁——要不要起子 Agent？并行还是串行？用哪个模型（便宜但弱的 gpt-mini 做 triage，强的 claude-sonnet 写代码）？
2. **Plan**：生成**可验证**的执行方案——**明确写清楚完成标准是什么、失败怎么办、资源预算（Token / 时间 / 重试次数）多少**。没有这些约束的 Plan 不是 Plan，只是让 AI "瞎跑"。

一个合格的 Plan 示例：
```
目标：修复 #1284 Issue（搜索摘要丢失关键词）
完成标准：
  - [ ] buildSearchSnippet() 对代码块内关键词命中能返回非空 snippet
  - [ ] go test ./internal/service/... -run Search 全部通过
  - [ ] 新增 1 条单元测试覆盖"代码块内关键词"场景
  - [ ] 不修改任何对外 API 签名
失败策略：单步失败重试 2 次，总轮次上限 5
预算上限：150K Token / 30 分钟
```

#### Step 3 - Act（执行任务）

Loop 的"手"。在隔离环境中落地任务：写代码、调工具、跑构建、发请求、改文件。**关键原则：所有 Act 必须在可回滚的隔离环境里做**——至少是 git 分支，最好是 git worktree 或容器沙箱，绝不能直接改主工作目录。

#### Step 4 - Verify（验证结果）⭐ 整个 Loop 的灵魂

Loop 最关键的设计：**绝对不能让写代码的那个模型自己给自己打分**。同一模型既当选手又当裁判，天然会"放水"，漏检率非常高。

正确的验证方式（按可靠性从高到低排）：

| 验证方式 | 说明 | 可靠性 |
|---|---|---|
| **客观工具验证** | 跑测试 / lint / 类型检查 / 构建结果 / 基准对比 | ★★★★★ |
| **第二模型独立审查** | 用不同模型（甚至不同厂商）重新检查输出，专门挑错反驳 | ★★★★ |
| **同模型自检查 + checklist** | 给它一份 checklist，强制逐条过（有帮助但仍会放水） | ★★★ |
| **写代码的模型自己说"我做完了"** | 几乎不可信 | ★☆☆ |

> **经验总结**：客观工具验证 + 第二模型审查的组合，是目前 Loop Engineering 里被验证最有效的防漏检搭配。

#### Step 5 - Persist & Iterate（持久化 & 迭代）

Loop 的"记忆"。验证不通过，就把**失败原因 + 验证报告**喂回 Act 阶段让它重改；验证通过，就把经验写进长期记忆（Skill 文件 / 知识库 / 状态文件），供后续循环复用。

一个典型的"记忆沉淀"示例（写入项目 SKILL.md）：
```
### 搜索功能相关：代码块内关键词丢失问题

【现象】搜索"xxx"出现在结果里，但标题/摘要都没命中，也没有高亮上下文。
【根因】stripMarkdown() 的 reFenced / reInlineCode 正则把```和反引号连同内容一起删了。
【修复】正则加捕获组，只删标记不删内容：(?s)\`\`\`[^\n]*\n?(.*?)\`\`\` → Replace 为 " $1 "
【单元测试】见 TestStripMarkdown_PreservesFencedCodeContent
```

这样以后再遇到同类问题，Loop 不需要重新从报错推导一遍，直接走经验路径即可。**循环跑得越多，记忆越厚，效率越高——这是 Loop 相对"单次 Prompt 模式"的复利优势。**

---

## 五、六大核心组件（Addy Osmani 框架）

Google Cloud AI 总监 Addy Osmani 提出了 Loop Engineering 的**六大组件框架**。它们像六块积木，层层叠加、互为支撑，可以组合出各种形态的循环系统。

```
                    ┌─────────────────────────┐
                    │  ⑥ Governance / Guard   │  ← 治理层：审批、审计、成本、回滚
                    ├─────────────────────────┤
                    │  ⑤ Automations（心跳）  │  ← 什么时候开工？
                    ├─────────────────────────┤
                    │  ④ Memory / Skills      │  ← 长期记忆复用
                    ├─────────────────────────┤
                    │  ③ Plugins / Connectors │  ← 连接真实世界的工具
                    ├─────────────────────────┤
                    │  ② Worktrees（隔离）    │  ← 并行不冲突
                    ├─────────────────────────┤
                    │  ① Harness（底座）      │  ← 单次运行的环境
                    └─────────────────────────┘
```

### 5.1 组件一：Harness（单次运行底座）

Harness 是 Loop 的底层地基。它回答一个问题：**单次 Agent 运行时，可以做什么、不能做什么？**

包含的要素：
- **文件权限**：可读哪些目录？可写哪些目录？碰不到数据库 / 配置 / 密钥？
- **工具权限**：允许跑哪些命令？允许发网络请求到哪些域名？
- **资源限制**：单轮运行的时间上限、Token 预算、磁盘写入上限
- **沙箱环境**：是在容器里跑？还是直接在宿主机跑？

> 类比：Harness 就像给 AI 实习生准备的工位——哪张桌子（工作目录）、能用哪些文具（工具）、哪些柜子不能开（权限）。Loop 是在这个工位之上，设计的"每天自动派活+验收+返工+归档"的流程。

### 5.2 组件二：Worktrees（工作树隔离）

让多个 Agent **并行跑而互不打架**的关键。基于 `git worktree` 机制，给每个子 Agent 分配一份独立的代码副本：

```
main repo (工作目录)
├── worktree/pr-fix-1284/   ← Agent A 修 #1284
├── worktree/dep-upgrade/   ← Agent B 升依赖
├── worktree/changelog-v2/  ← Agent C 写 changelog
└── worktree/pr-babysit-7/  ← Agent D 审 #7 PR
```

没有 Worktree，多个 Agent 同时改同一份文件会互相覆盖，根本做不到并行。

> ⚠️ 实务提醒：Worktree 只解决了"代码冲突"的问题，**人的审查带宽才是并行数量的真正天花板**。一般建议初期并行数控制在 2~3 个，等你对验证环节有信心了再往上加。

### 5.3 组件三：Plugins & Connectors（插件与连接器）

通过 **MCP（Model Context Protocol）协议**让 AI 连接真实工具，实现从"AI 告诉你怎么做"到"AI 自己动手做"的跃迁。

典型的连接器能力：

| 连接器 | 能做什么 | 对应本博客项目的场景 |
|---|---|---|
| Git / GitHub | 提 PR、打标签、改 Issue 状态、发 review comment | 自动 PR babysit、Issue triage |
| 本地 Shell | 跑构建 / 测试 / lint / 任意 CLI | `go test`、`go build`、前端 `npm run build` |
| 文件系统 | 读写指定目录下的文件 | 修改 `internal/service/`、`blog-web/src/` 等 |
| 数据库（只读） | SELECT 查数据 | 排查文章评论数据异常 |
| 浏览器 / HTTP 客户端 | 开网页、发请求 | 检查 API 响应、抓取文档、测试部署后页面 |
| 邮件 / IM | 发通知 | PR ready 时 @ 负责人钉钉 / Telegram |
| CI 系统 | 触发构建、拉取失败日志 | CI 挂了自动拉日志分析 |

### 5.4 组件四：Memory & Skills（长期记忆 / 技能包）

解决 AI **"每次冷启动都从零开始"** 的问题。把项目知识、代码规范、历史踩坑经验沉淀成可复用文件（通常是 `SKILL.md`）。

Skill 的加载策略是**渐进式**：
1. 先显示 Skill 的**名称和一句话描述**（不占 Token）
2. 匹配到对应任务场景时，再**加载完整内容**
3. 需要时再调用 Skill 附带的脚本 / 资源 / 参考链接

收益是**复利效应**：每跑一轮，AI 对项目的理解就更稳一层，遇到老问题不再需要从零推导。

本项目已经有一个 Skill 的雏形参考：`.claude/skills/web-design-engineer/SKILL.md`，就是典型的 Skills 组件落地方式。

### 5.5 组件五：Automations（自动化心跳）

决定 Loop **什么时候开工**，是整个系统的"起搏点"。没有 Automations，Loop 本质上还是个一次性脚本。

| 触发方式 | 示例 | 适用场景 |
|---|---|---|
| **定时 cron** | `0 9 * * 1-5` 工作日 09:00 | 每日 Issue triage、每周依赖扫描、每月 changelog |
| **事件 hooks** | `on: pull_request.opened`、`on: issue.created` | PR 保姆、新 Issue 自动打标签、CI failed 拉起修复 |
| **命令触发** | `/loop babysit-pr 7`、`/loop fix-ci` | 手动拉起特定循环（类似 Trae / Claude Code 的 slash command） |
| **流水线触发** | GitHub Actions step 失败后拉起 | CI 清扫、nightly 构建失败根因分析 |
| **监控告警触发** | 错误率 > 5% 持续 5 分钟 | 生产异常应急响应 Loop |

### 5.6 组件六：Governance & Guardrails（治理与护栏）

所有前面五个组件都是"让系统能跑"，这第六个是**"让系统敢跑"**。没有治理的全自动 Loop 就是定时炸弹。

核心治理维度：

| 维度 | 回答的问题 | 落地方式 |
|---|---|---|
| **人类审批闸门（HITL）** | 哪些动作必须人点头才能做？ | 自动提交只能到 draft PR、删除操作必须人确认、major 依赖升级出报告不自动改 |
| **分阶段上线** | 多大信任做多大事？ | L1：只读报告；L2：可改文件但不提交；L3：自动提 PR；L4：无人值守合并（极少项目需要） |
| **成本监控** | 跑一天要花多少钱？ | 每轮记录 Token 用量，设置日预算 / 月预算告警 |
| **审计日志** | AI 都干了啥？ | 每轮 Prompt、每步工具调用、每次文件变更、每次决策原因全部可追溯 |
| **自动回滚** | 搞砸了怎么办？ | Loop 触发的改动全部打特殊 tag，一键 `git revert` 或自动触发回滚 |
| **风险上限** | 最多允许它"瞎试"到什么程度？ | 单任务总轮次上限（≤5）、总 Token 上限、运行时间上限、文件写入量上限 |

> **实务建议**：从 L1 开始跑 1~2 周，看 Loop 出的报告跟你手动判断的差异率 < 10%，再升到 L2；再跑 2 周稳定，再升 L3。**信任是攒出来的，不是放出来的**。不要一上来就开全自动。

---

## 六、七种生产级循环模式

开源项目 `loop-engineering` 提供了 7 套经过验证的生产级循环模式，拿来就可以落地到任何代码仓库。它们也是 Loop Engineering 最常见的七种应用形态。

### 模式 1：Issue Triage（新 Issue 自动分类）

**什么时候跑**：新 Issue 被创建时触发（事件驱动）
**做什么**：
1. 读 Issue 标题 + 正文
2. 自动判断属于哪个模块（前端/后端/数据库/部署/文档…）
3. 自动打 Label（`bug` / `feature` / `question` / `good first issue`…）
4. 判断优先级（`P0/P1/P2/P3`）
5. 对常见问题自动回复模板化 FAQ 链接
6. 对明确是 bug 的 P0，自动拉起对应修复 Loop
**产出**：分类完成的 Issue + 标签 + 优先级 + 初步回复
**信任等级**：L1（只 propose 不直接改，人审一下）

### 模式 2：Daily Triage（每日全盘扫描）

**什么时候跑**：工作日 09:00 cron 触发
**做什么**：
1. 扫 open Issue / PR 全量
2. 标记「超过 7 天没人碰」的陈旧项
3. 标记「需要人介入」（超出 AI 决策边界）的项并 @ 负责人
4. 对明确低风险的项（比如 typo、文档链接失效）直接自动处理并提 PR
5. 输出一份「今日待处理清单」发 IM 通知
**产出**：每日一封 / 一条 triage 报告，人只需要处理清单里高亮的部分
**信任等级**：L1 报告 + L2 低风险自动改

### 模式 3：PR Babysitter（PR 保姆）⭐ 最常用

**什么时候跑**：`pull_request.opened` / `pull_request.synchronize` / `review_comment.created` 事件
**做什么**：
1. 读 PR diff + 关联 Issue
2. **静态层面**：代码风格、命名规范、潜在的 nil 指针、SQL 注入风险、有没有漏写软删除条件
3. **逻辑层面**：改动是否符合 Issue 原始需求？边界 case 考虑到了吗？
4. **测试层面**：覆盖率有没有掉？新增逻辑有没有加测试？
5. 发 review comment，附具体的修改建议 + 代码片段
6. 有人留 review comment 后，自动把评论转成子任务让写 PR 的 Agent 跟进修改
**产出**：一份 AI Code Review 结果（评论 + 建议），人来终审
**信任等级**：L2（写评论建议不直接改）

### 模式 4：CI Sweeper（CI 清扫）⭐ ROI 最高

**什么时候跑**：CI workflow `run_failed` 事件触发
**做什么**：
1. 拉取 CI 失败日志（build log / test log）
2. 定位根因（是编译错？测试挂？lint 不过？依赖装不上？超时？flaky test？）
3. 在独立 worktree 里尝试修复
4. 本地跑同样的测试 / lint 验证
5. 通过 → 自动提修复 PR，@ 原 PR 负责人
6. 不通过 → 最多重试 3 次，不行就发「已尝试但需要人手」报告
**产出**：修复 PR 或一份「根因 + 已尝试 + 需要人做什么」的诊断报告
**信任等级**：L2（提 PR 不自动合并）
> **为什么 ROI 最高**：CI 挂是每个团队每天都碰的高频事，而且绝大多数失败（拼写、import、lint、依赖小版本、路径错）都属于「规则明确、可验证」的理想 AI 场景。

### 模式 5：Dependency Sweeper（依赖清扫）

**什么时候跑**：每周一 cron + 安全告警事件
**做什么**：
1. 扫描 `go.mod` / `package.json` 依赖更新
2. 分三类处理：
   - **patch 升级**（1.2.3 → 1.2.4）：自动升 + 跑全量测试 + 过了就提 PR 自动合并（L3）
   - **minor 升级**（1.2.x → 1.3.0）：跑测试 + 生成「变更点 + 风险 + 验证结果」报告，提 PR 等人审（L2）
   - **major 升级**（1.x → 2.0）：只出迁移评估报告，不自动改代码（L1）
3. 有 CVE 安全公告 → 不分大小，立即跑 + 发 P0 通知
**产出**：patch 自动合并的 PR，minor/major 带评估报告的 PR，CVE 告警
**信任等级**：patch=L3，minor=L2，major=L1，CVE=L0（立即通知人）

### 模式 6：Changelog Drafter（Changelog 起草）

**什么时候跑**：`release/*` 分支创建时触发，或每月末 cron
**做什么**：
1. 读从上一个 tag 到当前 HEAD 之间的全部 commit log + PR 标题
2. 按 `feat / fix / docs / chore / breaking change` 分组
3. 归并同类项、去重、用中文重写成面向用户可读的描述
4. 写入 `CHANGELOG.md` 对应版本的草稿
5. 提 PR，让人做**润色 + 终审**
**产出**：一份 80 分的 Changelog 草稿，人只需要改剩下的 20% 措辞和重点排序
**信任等级**：L2（提草稿 PR，不自动合并）

### 模式 7：Post-Merge Cleanup（合后清理）

**什么时候跑**：`pull_request.merged` 事件
**做什么**：
1. 删源分支（如果 PR 描述里没标 `keep-branch`）
2. 把 PR 关联的所有 Issue 自动标为 `resolved` / `closed`
3. 同步更新项目管理看板状态
4. 如果这个 PR 关联了 milestone，检查 milestone 是不是全关了，是就发祝贺通知
5. 触发一次 changelog draft 更新
**产出**：仓库保持整洁，Issue/看板/里程碑状态与 PR 同步
**信任等级**：L3（全自动，除非有明确 `keep` 标记）

---

## 七、落地实践指南：从 0 到 1 搭第一个 Loop

### 7.1 选第一个 Loop 的原则

第一个 Loop **不要选 PR Babysitter 或 CI Sweeper 这种看起来很爽的**，翻车率太高。按以下优先级选：

| 优先级 | 推荐的第一个 Loop | 原因 | 失败了的代价 |
|---|---|---|---|
| ✅ 1 | **Daily Triage（只读报告模式）** | 不改任何东西，只出报告。你每天对比它的判断和你的判断，差异在哪里，慢慢调规则。 | 0：就是一份报告，你看两眼丢了也没事 |
| ✅ 2 | **Changelog Drafter** | 输出是纯文字草稿，改不改都行，而且结果好坏一眼可见，容易建立信心。 | 很低：草稿不对，人手动改回来就行 |
| ✅ 3 | **Issue Triage（propose-only）** | 只建议标签不直接打，或者打到草稿标签等你 confirm。 | 低：标签错了你改回来就行 |
| ❌ 不推荐起步就上 | CI Sweeper / Dependency Sweeper / 自动合并 | 会直接改仓库状态，一旦错了就有实质影响。 | 高：可能引入 bug、破坏构建 |

### 7.2 实施的七个步骤

```
Step 1 盘点痛点  →  Step 2 定义完成标准  →  Step 3 搭建 Harness 底座
                                                             ↓
Step 7 分阶段升级 ←  Step 6 调优 & 沉淀记忆 ←  Step 5 加 Verify
                                                             ↓
                                         Step 4 连触发器 + 跑 Act
```

### 7.3 一个最小可用的 Daily Triage Loop 伪代码（适配本项目）

```go
// 伪代码，表达逻辑而非可直接运行
func DailyTriageLoop(ctx context.Context) {
    // 1. Discover：扫 open Issue + PR
    issues := github.ListOpenIssues("gin博客")
    prs    := github.ListOpenPRs("gin博客")

    // 2. Assign & Plan：挑 3 个需要关注的，给每个写明确完成标准
    targets := pickTop3NeedAttention(issues, prs)

    // 3. Act：分类 + 出建议（只读，不直接改）
    report := []ReportItem{}
    for _, t := range targets {
        suggestion := agent.Triage(
            ctx,
            Input{
                IssueOrPR:   t,
                ProjectRule: ReadFile("docs/后端代码开发规范.md"), // <-- Context
                Skill:       ReadFile(".claude/skills/..."),       // <-- Skills
            },
        )
        report = append(report, ReportItem{
            Target:     t.URL,
            Module:     suggestion.Module,     // 前端/后端/文档…
            Priority:   suggestion.Priority,   // P0~P3
            Suggestion: suggestion.Action,     // 建议怎么处理
            Why:        suggestion.Reasoning,  // 为什么这么建议
        })
    }

    // 4. Verify（L1 级：第二模型检查分类是否离谱）
    sanityCheck := agent.ReviewReport(ctx, report)
    if !sanityCheck.Pass {
        report = append(report, ReportItem{
            Target:     "[SANITY WARNING]",
            Suggestion: sanityCheck.Flags,
        })
    }

    // 5. Persist：沉淀结果（存仓库 docs/.triage/2026-07-11.md）
    WriteTriageArchive(report)

    // 6. 产出：发 IM + 存 issue comment
    telegram.SendDailyReport(report)
}

// 触发方式：cron 每天 09:00 调上面这个函数
```

这个最小 Loop **只输出不改仓库**，你可以先连续跑一周，每天对比一下它给的分类建议和你自己的判断——如果 80% 以上一致，恭喜，你对这套 Loop 的信任已经建立起来了，可以升到 L2（直接打草稿标签）。

### 7.4 Loop Ready 评分（项目就绪度游戏化）

开源 `loop-audit` 工具的做法：扫描你的代码仓库后，给一个 **0~100 的 Loop Ready 分数**，每满足一项加分，类似这样：

| 检查项 | 权重 | 本项目状态 | 得分 |
|---|---|---|---|
| 有统一的代码规范文档 | 10 | ✅ `docs/后端代码开发规范.md` 存在 | 10 |
| 有 ≥80% 的单元测试覆盖（可做客观 Verify） | 20 | ❌ Go service 层有部分测试（`article_service_test.go`），前端暂无 | 5 |
| CI 可复现（GitHub Actions / 本地脚本） | 15 | ⚠️ 有 `Dockerfile` + `docker-compose`，CI 脚本待完善 | 8 |
| 有 Skill 文档沉淀历史踩坑 | 10 | ✅ `.claude/skills/` 目录存在 | 10 |
| 有 PR 模板 + Issue 模板 | 10 | ❌ 暂无 | 0 |
| 有 git worktree 使用脚本 | 5 | ❌ 暂无 | 0 |
| 有完整的错误码体系（可被 AI 识别） | 10 | ✅ `pkg/errors/code.go` 存在 | 10 |
| 有统一响应结构（Verify 好判断 API 结果） | 10 | ✅ `pkg/response/` 存在 | 10 |
| 有部署回滚方案 | 10 | ⚠️ Docker 镜像 tag 可回滚，但缺少显式文档 | 6 |
| **合计** | **100** | | **59 / 100** |

这个评分表的目的不是"打分焦虑"，而是给你一个清单：**要让 Loop 在你的项目上跑得稳，你需要先补齐哪些基础设施**。Test coverage、CI、规范文档、模板——这些东西本来就是软件工程的"正规军"做法，Loop Engineering 只是把它们的收益放大了，因为现在不是你一个人在用它们，而是你的 AI 系统每天都在用。

---

## 八、陷阱与反模式

概念很热，但落地踩坑的人更多。下面是社区里被反复验证过的"不要这么做"清单。

### ❌ 陷阱 1：一上来就开全自动

**现象**：刚学 Loop，第一天就把 CI Sweeper 开成自动合并。第三天发现 Loop 为了让测试通过，把测试里的断言删了，然后高高兴兴自己给自己打了个「验证通过」的章，合并进主分支后半夜线上崩了。

**为什么会发生**：缺少 Verify 环节的角色分离 + 缺少 Governance。写代码的模型天然会往「让测试绿」而不是「让功能对」的方向优化。

**正确做法**：从 L1（只读报告）起步，每升一级至少观察 1~2 周。

### ❌ 陷阱 2：没有完成标准 = 循环跑死

**现象**：Plan 里只写了「优化文章搜索」，没写什么算优化完。Loop 进入无限循环——改一版、自我感觉好一点、再改一版、再感觉好一点……Token 烧了 50 万，还在跑。

**正确做法**：Plan 必须写成**可客观验证的 checklist 清单**，并加上硬性的轮次上限 + Token 上限 + 时间上限。三顶帽子一顶也不能少。

### ❌ 陷阱 3：同一个模型既执行又验证（自己当裁判）

**现象**：同一个 Claude 实例，写完代码后问它「测试过了吗？」，它很自信地说「没问题」，结果你一跑，挂了一堆。

**为什么会发生**：LLM 天生有"一致性偏见"——之前说过的话，后面倾向于找理由支持而不是推翻。让写代码的模型审代码，它会找理由说这代码写得好，而不是挑错。

**正确做法**：
1. 第一优先级：**客观工具验证**（测试 / lint / 类型检查），这是金标准
2. 第二优先级：**换模型 / 换厂商**独立审查（比如用 Claude 写，用 GPT 审，或者反过来）
3. 模型自审只能做辅助，不能做终审

### ❌ 陷阱 4：把 Loop 当万能咒语，什么都往里塞

**现象**：把需要深度产品判断 / 设计审美 / 跨团队协商的需求也塞给 Loop，结果 Loop 反复改了 5 轮，每次都不对，最后还得你全部重写。Token 花了，时间也浪费了。

**正确做法**：Loop 擅长的是**规则明确、结果可验证、风险可控**的任务。用之前先问自己三个问题：
1. ✅ 有没有**客观**的完成标准（测过/编译过/格式对）？
2. ✅ 失败了代价**小且可回滚**？
3. ✅ 这个任务**每周至少发生 1 次**，做自动化的 ROI 是正的？

三个都 Yes，才值得写 Loop。有一个 No，先别上。

### ❌ 陷阱 5：不写记忆，每次循环都从零开始

**现象**：同一个坑踩了三次——上次修过的搜索摘要 bug，换个 Issue 形式出现，Loop 又把之前踩过的坑全部重新踩了一遍。

**正确做法**：每轮验证通过后，强制触发「记忆沉淀」步骤：
- 这次的任务类型是什么？
- 遇到了什么坑？
- 最终怎么解的？
- 新增了什么测试可以防回归？

写进 Skills / Project Memory，下一轮 Loop 冷启动时先加载。记忆越厚，Loop 越聪明，这是复利。

### ❌ 陷阱 6：忽视人类心理——"Loop 提的 PR 我懒得审"

**现象**：Loop 提的 PR 太多，负责人看标题说"应该没问题"就直接点合并，没有真的过一遍。Loop 引入的一个小 bug 就这样混进主分支，搞出大事。

**为什么会发生**：人的注意力是稀缺资源。如果每天 10 个 PR 全是 Loop 提的，人会进入"自动通过"的麻木状态，等于把 Governance 又丢了。

**正确做法**：
- 控制 Loop 提 PR 的**每日上限**（比如 ≤2 个/day）
- 对 Loop 提的 PR 强制要求**至少一个人类 reviewer 明确 approve**，不能因为是 AI 提的就走快速通道
- 定期随机抽 Loop 合入的 PR 做"事后复盘 audit"，看看漏检率

---

## 九、与本博客项目的结合建议

看完了理论，回到我们的 `gin博客` 项目本身。哪些场景值得现在就上 Loop？按推荐优先级排列：

### 🥇 P1：值得优先做（高 ROI + 低风险）

| Loop 模式 | 适配场景 | 预期收益 | 信任等级起步 |
|---|---|---|---|
| **Changelog Drafter** | 每次积累一批 commit 合入后，自动基于 `git log` 生成中文 changelog 草稿，写入 `CHANGELOG.md` | 省去每次发版最烦的文案整理工作 | L2：提草稿 PR，人改措辞即可 |
| **Issue Triage (L1)** | 以后开放 Issue 后，新 Issue 自动判断模块（前端 blog-web / 后台 blog-admin / 后端 Go / 部署 / 文档）+ 打标签建议 | 分类时间从 5min/Issue → 10s 确认 | L1：只读报告，不直接打 tag |
| **Dependency Sweeper (patch only)** | 每周一扫描 `go.mod` 和两个前端 `package.json` 的 patch 升级，自动提 PR + 跑 `go test` + `npm run build` | 安全补丁 + 小版本不堆积 | L2：patch 级提 PR；minor/major 只出报告 |

### 🥈 P2：等测试体系更完善后再做

| Loop 模式 | 阻塞条件 | 补完什么就可以上 |
|---|---|---|
| **PR Babysitter（AI Code Review）** | 当前 Go 层有测试（`article_service_test.go`）但覆盖率不足，前端两个 Vue 项目完全没有自动化测试 | 给 Go service 层补到 ≥60% 覆盖率 + 给前端补基础的 vitest 测试脚手架 |
| **CI Sweeper** | 当前 CI 流程不完善（项目里没有 GitHub Actions / GitLab CI 配置文件），本地脚本也没有统一的一键 lint+test+build | 补 `scripts/ci.sh`：一键跑 gofmt / go vet / go test / 前后端 build；然后上 GitHub Actions 把它跑起来 |

### 🥉 P3：暂不推荐（本项目规模还不到，ROI 为负）

| Loop 模式 | 为什么现在不做 |
|---|---|
| Post-Merge Cleanup | 目前项目维护者 ≤2 人，手动删分支/关 Issue 花不了 30 秒，写 Loop 的时间比手动还多 |
| Daily Triage 全盘扫描 | Issue/PR 总量还不多，手动扫一遍只需要 2 分钟。等 Issue 日增 ≥5 再上 |
| 生产监控告警触发的应急 Loop | 目前还没有完善的生产监控 + 日志聚合系统，先上监控，再考虑 Loop 接入 |

### 可以直接复用的现有项目资产

要启动 P1 级的三个 Loop，我们的项目已经具备不少基础设施：

| 资产 | 作用 | 路径 |
|---|---|---|
| 代码规范文档 | Triage / PR Babysitter 的判断依据 | [后端代码开发规范.md](file:///e:/goCode/goFile/gin博客/docs/后端代码开发规范.md) |
| 错误码体系 | CI Sweeper 识别后端报错的基础 | [code.go](file:///e:/goCode/goFile/gin博客/pkg/errors/code.go) |
| 统一响应结构 | Verify 时判断 API 输出是否合规 | [response/](file:///e:/goCode/goFile/gin博客/pkg/response/) |
| 技能包雏形 | Skills 组件的现成目录与写法参考 | [web-design-engineer/SKILL.md](file:///e:/goCode/goFile/gin博客/.claude/skills/web-design-engineer/SKILL.md) |
| 单测基础设施 | article 相关业务已经有可运行的 `go test` 基础 | [article_service_test.go](file:///e:/goCode/goFile/gin博客/internal/service/article_service_test.go) |

---

## 十、总结：你真正要学的是什么

Loop Engineering 不是一个框架，不是一个库，也不是一个需要花几千块钱报班学的神秘新技能。

它的本质是两句软件工程的老话，在 AI 时代换了一件新衣服：

1. **「把重复劳动自动化」**——过去你写 Makefile、写 shell 脚本、写 GitHub Actions 做的事，现在写 Loop 用 AI 驱动来做。能力边界更广（能做"理解性"的判断，不只是机械执行），但工程原则一模一样：边界、验证、回滚、告警。
2. **「把判断力前移到系统设计阶段」**——过去你写业务规则、写 if/else、写状态机、写 RBAC 权限矩阵，本质上都是在把"遇到 X 情况应该怎么办"提前写进系统里。现在写 Loop，无非是把这套原则从"业务规则"延伸到"AI 应该怎么被驱动"这件事上。

所以，如果读完这篇文档你只记住三句话，记住下面这三句：

1. **先别急着写 Loop，先补可验证性。** 没有测试、没有 CI、没有完成标准，写出来的 Loop 不是生产力工具，是随机数生成器。
2. **从 L1 起步，观察 1~2 周再升一级。** 信任是靠一版一版稳定的输出攒出来的，不是靠"我相信 AI"。
3. **让它做 AI 擅长的事——规则明确、可验证、高频。** 不要把需要产品判断 / 审美 / 协商的事塞进去，那是你的工作，不是 Loop 的。

---

## 参考资源

| 资源 | 链接 | 说明 |
|---|---|---|
| Addy Osmani 原文《Loop Engineering》 | 搜索 "Addy Osmani Loop Engineering" 即可找到 | 命名人官方定义框架文 |
| Boris Cherny（Anthropic Claude Code 负责人）关于循环的分享 | Anthropic 博客 & 活动录像 | 最早公开分享"我的工作是写循环"的一手来源 |
| 开源项目 `cobusgreyling/loop-engineering` | GitHub 7K+ Star | 7 个 CLI + 7 个生产模式的 npm monorepo 实现参考 |
| Peter Steinberger 原始推文 | X / Twitter @petersteinberger | 引爆社区的"每月提醒"原文帖 |
| Anthropic《Building Effective Agents》| Anthropic docs | 官方 Agent 设计指南，强调 "LLM + 工具 + 反馈循环 = Agent" |
| OpenAI Codex / Trae 内置的 Automation / Skills / Subagents / Worktrees 功能 | 产品文档 | 产品化落地的 Loop 组件参考，拿来即用，不需要自己造轮子 |
