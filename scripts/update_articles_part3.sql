-- 继续补充文章内容
USE blog;

-- 文章155: 周末徒步：秦岭深处的宁静
UPDATE articles SET content = '## 出发

周末天气晴朗，决定去秦岭徒步。远离城市的喧嚣，去山里呼吸新鲜空气，放松一下心情。

## 路线选择

选择了一条约 10 公里的环形路线，难度适中，适合周末半日游。

### 路线信息
- 起点：秦岭北麓某停车场
- 终点：同起点（环形路线）
- 距离：约 10 公里
- 海拔：600-1200 米
- 用时：约 4-5 小时

## 沿途风景

### 山间小路
蜿蜒的山间小路，两旁是郁郁葱葱的树木。阳光透过树叶洒下斑驳的光影，走在其中格外惬意。

### 溪流
途中遇到一条清澈的溪流，水声潺潺，忍不住停下来休息一会儿，听听大自然的声音。

### 山顶
登上山顶，视野豁然开朗。远处的山峦层层叠叠，蓝天白云下格外壮观。

## 途中感悟

### 关于节奏
徒步最重要的不是速度，而是节奏。保持自己的节奏，不急不躁，才能走得更远。

### 关于坚持
中间有几段比较陡的坡，确实有点累。但咬牙坚持过去后，看到的风景更加美丽。

### 关于放下
在山里，手机信号时有时无。反而因此放下了手机，专注于脚下的路和身边的风景。

## 装备清单

- 登山鞋（必备）
- 双肩包
- 水和干粮
- 防晒霜和帽子
- 登山杖（可选）
- 急救包

## 安全提示

1. 告知家人朋友你的行程
2. 不要独自进入未开发区域
3. 注意天气变化
4. 保持手机电量充足
5. 携带足够的水和食物

## 结语

秦岭的宁静让人难忘。在忙碌的工作之余，偶尔走进大自然，感受山水之美，是对心灵最好的滋养。

下次准备挑战更高难度的路线！' WHERE id = 155;

-- 文章156: AI 编程助手使用心得
UPDATE articles SET content = '## 引言

AI 编程助手已经成为开发者工具箱中的重要成员。本文分享使用 AI 编程助手的一些心得和技巧。

## 常用工具

### GitHub Copilot
- 代码补全
- 函数生成
- 注释生成代码

### Claude / ChatGPT
- 代码解释
- Bug 修复建议
- 方案设计讨论

### Cursor
- 代码编辑
- 智能重构
- 多文件编辑

## 高效使用技巧

### 1. 提供足够的上下文

```go
// 不好的提示
"写一个函数"

// 好的提示
"在 Go 语言中，写一个函数，接收 []User 切片，返回年龄大于 18 的用户列表，使用泛型"
```

### 2. 分步骤提问

不要一次问太复杂的问题，可以分步骤：
1. 先讨论方案设计
2. 再讨论具体实现
3. 最后优化和改进

### 3. 验证生成的代码

AI 生成的代码不一定完全正确，需要：
- 检查逻辑是否正确
- 运行测试验证
- 检查边界情况

### 4. 利用 AI 学习

当遇到不熟悉的代码时，可以让 AI 解释：
- 这段代码的作用是什么？
- 为什么要这样写？
- 有什么需要注意的地方？

## 实际应用场景

### 快速原型开发
用 AI 快速生成基础代码，然后手动优化。

### 代码审查
让 AI 帮忙检查代码，发现潜在问题。

### 文档生成
根据代码自动生成注释和文档。

### 单元测试
让 AI 帮忙生成测试用例。

## 注意事项

### 不要过度依赖
AI 是工具，不是替代品。核心逻辑还是需要自己思考。

### 保护敏感信息
不要把公司的敏感代码或数据发给 AI 服务。

### 保持学习
使用 AI 的同时，也要理解背后的原理。

## 效率对比

| 任务 | 手动开发 | AI 辅助 |
|-----|---------|--------|
| 写 CRUD | 30 分钟 | 5 分钟 |
| 写测试 | 20 分钟 | 5 分钟 |
| 调试 Bug | 30 分钟 | 10 分钟 |
| 学习新 API | 1 小时 | 15 分钟 |

## 总结

AI 编程助手是提高效率的利器，但要正确使用。把 AI 当作结对编程的伙伴，而不是代码生成机器。

善用 AI，让编程更高效、更有趣。' WHERE id = 156;

-- 文章157: 从零开始学习 Kubernetes
UPDATE articles SET content = '## 什么是 Kubernetes？

Kubernetes（简称 K8s）是一个开源的容器编排系统，用于自动化容器化应用的部署、扩展和管理。

### 核心功能
- 自动化部署和回滚
- 服务发现和负载均衡
- 自动装箱（调度）
- 自我修复
- 密钥和配置管理

## 基本概念

### Pod
Pod 是 Kubernetes 中最小的可部署单元，包含一个或多个容器。

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - name: my-container
    image: nginx
    ports:
    - containerPort: 80
```

### Deployment
Deployment 用于管理 Pod 的副本集。

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-container
        image: nginx
```

### Service
Service 为一组 Pod 提供稳定的网络访问。

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: my-app
  ports:
  - port: 80
    targetPort: 80
  type: LoadBalancer
```

### Ingress
Ingress 管理外部访问集群内服务的规则。

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
spec:
  rules:
  - host: my-domain.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: my-service
            port:
              number: 80
```

## 安装 Kubernetes

### 使用 Minikube（本地学习）

```bash
# 安装 minikube
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube

# 启动集群
minikube start

# 查看状态
minikube status
```

### 使用 kubeadm（生产环境）

```bash
# 初始化主节点
sudo kubeadm init

# 配置 kubectl
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config

# 安装网络插件
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```

## 常用命令

```bash
# 查看集群信息
kubectl cluster-info

# 查看所有 Pod
kubectl get pods --all-namespaces

# 查看 Pod 详情
kubectl describe pod <pod-name>

# 查看日志
kubectl logs <pod-name>

# 进入 Pod
kubectl exec -it <pod-name> -- /bin/bash

# 删除 Pod
kubectl delete pod <pod-name>

# 扩缩容
kubectl scale deployment my-deployment --replicas=5

# 滚动更新
kubectl set image deployment/my-deployment my-container=new-image:latest

# 回滚
kubectl rollout undo deployment/my-deployment
```

## 实战：部署应用

### 1. 创建 Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-go-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-go-app
  template:
    metadata:
      labels:
        app: my-go-app
    spec:
      containers:
      - name: my-go-app
        image: my-registry/my-go-app:v1
        ports:
        - containerPort: 8080
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "128Mi"
            cpu: "500m"
```

### 2. 创建 Service

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-go-app-service
spec:
  selector:
    app: my-go-app
  ports:
  - port: 80
    targetPort: 8080
  type: ClusterIP
```

### 3. 部署应用

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

# 查看状态
kubectl get pods
kubectl get services
```

## 最佳实践

1. 使用命名空间隔离环境
2. 设置资源限制
3. 使用 ConfigMap 和 Secret 管理配置
4. 实现健康检查（liveness 和 readiness）
5. 使用 RBAC 控制访问权限

## 总结

Kubernetes 功能强大，但学习曲线较陡。建议从 Minikube 开始学习，逐步掌握核心概念，再深入生产环境的实践。' WHERE id = 157;

-- 文章158: 读书笔记：《深入理解计算机系统》
UPDATE articles SET content = '## 书籍介绍

《深入理解计算机系统》（CSAPP）是一本经典的计算机科学教材，由 Randal E. Bryant 和 David R. O''Hallaron 合著。

## 主要内容

### 第一部分：程序结构和执行
- 信息的表示和处理
- 程序的机器级表示
- 处理器体系结构
- 优化程序性能
- 存储器层次结构

### 第二部分：在系统上运行程序
- 链接
- 异常控制流
- 虚拟内存

### 第三部分：程序间的交互和通信
- 系统级 I/O
- 网络编程
- 并发编程

## 重点笔记

### 信息的表示

计算机用二进制表示信息。理解补码、浮点数的表示方式很重要。

```
补码表示：
- 最高位是符号位
- 正数：原码
- 负数：取反加一

浮点数表示：
- 符号位 + 指数 + 尾数
- IEEE 754 标准
```

### 处理器体系结构

学习了简单的 Y86-64 处理器设计，理解了指令的执行过程：
- 取指
- 译码
- 执行
- 访存
- 写回

### 虚拟内存

虚拟内存是现代操作系统的核心概念：
- 每个进程有自己的地址空间
- 通过页表映射到物理内存
- 支持内存保护和共享

### 并发编程

三种实现并发的方式：
1. 进程：独立地址空间，通信复杂
2. I/O 多路复用：单进程处理多连接
3. 线程：共享地址空间，需要同步

## 学习心得

### 理论与实践结合
书中的实验非常重要，通过做实验可以加深理解。

### 系统思维
学习这本书培养了系统级的思维方式，理解了从源码到执行的完整过程。

### 优化意识
了解了编译器优化、缓存友好等概念，对写高性能代码有帮助。

## 推荐理由

- 内容全面，涵盖计算机系统的核心知识
- 讲解深入，不仅告诉你"是什么"，还告诉你"为什么"
- 实验丰富，动手实践加深理解

## 适合人群

- 计算机专业学生
- 想深入理解计算机系统的开发者
- 准备系统编程面试的工程师

## 总结

CSAPP 是一本值得反复阅读的经典书籍。虽然内容有些难度，但收获巨大。推荐每个计算机专业的学生都读一读。' WHERE id = 158;
