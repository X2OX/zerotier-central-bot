# [ZeroTier Central Bot](https://github.com/X2OX/zerotier-central-bot)

[![JetBrains Open Source Licenses](https://img.shields.io/badge/-JetBrains%20Open%20Source%20License-000?style=flat-square&logo=JetBrains&logoColor=fff&labelColor=000)](https://www.jetbrains.com/?from=blackdatura)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Introduce
ZeroTier's Telegram node management robot, can manage devices, monitor and notification node online and offline,
new device access.

The project uses Vercel's Serverless function as the backend, Firebase's Realtime Database as the database, 
and UptimeRobot as the timed task, and interacts through the Telegram Bot API.

Safety first, so the configuration is more cumbersome. Firebase's Realtime database requires rules to be set. 
But it’s not necessary, it’s just that you may be maliciously attacked.

Because the interface path is fixed, if the project is open source, it may be called maliciously. 
So the interface of Telegram Webhook has IP verify. The interface used for UptimeRobot call has Token verify.

Similarly, Telegram Bot has also taken protective measures. Without filling in the Telegram ID, anyone can use the bot, 
but there is no notification function. Only when the Telegram ID is filled in, the notification function is available. 
Because bots need to know the Chat ID to send messages, and group use is not currently supported.

## Deployment
### One-click deployment
[![](https://vercel.com/button)](https://vercel.com/new/project?template=https://github.com/X2OX/zerotier-central-bot&utm_source=x2ox&utm_campaign=oss)

### Manual deployment
1. Fork Repo
2. Import deployment in Vercel

### Get environment variable
#### Must
##### ZEROTIER_CENTRAL_TOKEN
Open [ZeroTier Account Center](https://my.zerotier.com/account)  click `New Token` generate

##### ZEROTIER_NETWORK_ID
Open [ZeroTier Network](https://my.zerotier.com/network) copy `NETWORK ID`

##### TELEGRAM_BOT_TOKEN
Open [Telegram BotFather](https://t.me/BotFather?start=x2ox) new bot

#### Optional
##### FIREBASE_URL
Used to store data
1. Open [Firebase Console](https://console.firebase.google.com/?x2ox) new project
2. New `Realtime Database`, copy database address
3. Set rules, use random string instead of `X_SECRET`
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
4. Splice `databaseURL+X_SECRET`, example: `https://vervel-default-rtdb.firebaseio.com/x2ox`

##### TELEGRAM_ID
Used for authentication and notification. 
Can be obtained through Bot's `/info` command or other methods

##### TIMER_TOKEN
used to cooperate with UptimeRobot to check status regularly. It's a random string.

### Configure environment variable

### Used Webhook
1. Bind a custom domain name for the project, it must be a single domain name (Telegram requires a non-wildcard certificate).
2. Set by `curl -F "url=https://<DOMAIN>/api/zerotier" https://api.telegram.org/bot<YOURTOKEN>/setWebhook`

### Turn on notifications
Open [UptimeRobot](https://uptimerobot.com/) Add monitoring, `DOMAIN` is your configured domain name, 
`TIMER_TOKEN` is a random string in the environment variable, the final address is `https://<DOMAIN>/api/timer?token=<TIMER_TOKEN>`

## Contribute
### ~~TODO~~
- Add FRP API

## Thanks
- Vercel supply Open Source project sponsorship
- JetBrains supply Open Source Licenses
- ZeroTier
- Firebase
- UptimeRobot

[![Powered by Vercel](https://hitokoto.x2ox.com/powered-by-vercel.svg)](https://vercel.com?utm_source=x2ox&utm_campaign=oss)
