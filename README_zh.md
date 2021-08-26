# [ZeroTier Central Bot](https://github.com/X2OX/zerotier-central-bot)

[![JetBrains Open Source Licenses](https://img.shields.io/badge/-JetBrains%20Open%20Source%20License-000?style=flat-square&logo=JetBrains&logoColor=fff&labelColor=000)](https://www.jetbrains.com/?from=blackdatura)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## 介绍
ZeroTier 的 Telegram 节点管理机器人，可以管理设备，监控节点的上下线、新设备接入并通知。

项目使用 Vercel 的 Serverless 功能作为后端、Firebase 的 Realtime Database 为数据库、UptimeRobot 做定时 API，通过 Telegram Bot API 交互。

数据库用于存储离线设备和新设备，避免重复通知。Serverless 没有定时功能，所以使用 UptimeRobot 的定时监控调用通知接口。

安全第一，所以配置比较繁琐。Firebase 的 Realtime Database 需要设定规则，当然，裸奔也可以，但是可能会被刷流量导致被 Firebase 处理。
因为接口路径是固定的，如果项目是开源的，那么可能会被人恶意调用。所以 Telegram Webhook 的接口有 IP 校验。
用于 UptimeRobot 调用的接口有 Token 校验。

同样的，Telegram Bot 也做了防护措施。在没有填入 Telegram ID 的情况下，任何人都可以使用 Bot，但是没有通知功能。
只有填入了 Telegram ID 的情况下，才有通知功能。因为 Bot 发送消息需要知道 Chat ID，而目前也不支持群组使用。

## 部署
### 一键部署
[![](https://vercel.com/button)](https://vercel.com/new/project?template=https://github.com/X2OX/zerotier-central-bot&utm_source=x2ox&utm_campaign=oss)

### 手动部署
1. Fork Repo
2. 在 Vercel 中导入部署

### 环境变量获取
#### 必须
##### ZEROTIER_CENTRAL_TOKEN
在 [ZeroTier 账号中心](https://my.zerotier.com/account) 的 `API Access Tokens` 点击 `New Token` 获取

##### ZEROTIER_NETWORK_ID
在 [ZeroTier Network](https://my.zerotier.com/network) 处复制 `NETWORK ID`， 目前不支持多 `NETWORK` 管理，没账号测试，有需求的话请提供账号

##### TELEGRAM_BOT_TOKEN
在 [Telegram BotFather](https://t.me/BotFather?start=x2ox) 处获取

#### 可选
##### FIREBASE_URL
用于存储数据
1. 在 [Firebase Console](https://console.firebase.google.com/?x2ox) 新建项目
2. 新建 `Realtime Database` 复制数据库地址
3. 设置规则，如下。使用随机串代替 `X_SECRET`
```json
{
    "rules":{
        "X_SECRET":{
            ".read":true,
            ".write":true
        }
    }
}
```
4. 根据 `databaseURL+X_SECRET` 拼接地址。例，地址为 `https://vervel-default-rtdb.firebaseio.com/` 规则 `X_SECRET` 为 `x2ox`，最后拼接为 `https://vervel-default-rtdb.firebaseio.com/x2ox`

##### TELEGRAM_ID
用于验证及通知。可以通过 Bot 的 `/info` 命令或其他方式获取

##### TIMER_TOKEN
用于 UptimeRobot 自动检测状态，应当是一个随机串

### 环境变量配置
在 Vercel 项目中配置环境变量

### 开启 Webhook
1. 为项目绑定自定义域名，必须为单域名（Telegram 要求非通配符证书）
2. 通过 `curl -F "url=https://<DOMAIN>/api/zerotier" https://api.telegram.org/bot<YOURTOKEN>/setWebhook` 设置

### 开启通知
在 [UptimeRobot](https://uptimerobot.com/) 添加监控，`DOMAIN` 为配置的域名，`TIMER_TOKEN` 是环境变量中的随机串，最终地址为 `https://<DOMAIN>/api/timer?token=<TIMER_TOKEN>`

## 贡献
因为是临时写出来自己用的，代码并没有做什么组织，而这个项目也应当保持足够的小。所以，会有很多鼓励贡献的地方，也有不会被接受的地方。

### ~~开发计划~~
- FRP API 接入，ZeroTier 没有 FRP 稳，双线应该就差不多了

### 不一定会被接受的贡献
- 多用户功能
- 会影响 API 调用的测试，Serverless 应当足够的快，且不应当滥用 Vercel
- 引入非常重、未经维护的库的代码

### 鼓励的贡献
- 代码重构，提升可读性
- 注释
- 文档
- 翻译

## 鸣谢
- Vercel 提供的 Open Source 项目赞助
- JetBrains 提供的 Open Source Licenses
- ZeroTier
- Firebase
- UptimeRobot

[![Powered by Vercel](https://hitokoto.x2ox.com/powered-by-vercel.svg)](https://vercel.com?utm_source=x2ox&utm_campaign=oss)
