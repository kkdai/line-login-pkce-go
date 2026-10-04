LINE login in Go using PKCE: Sample code for LINE login with PKCE in Go
==============

 [![GoDoc](https://godoc.org/github.com/kkdai/line-login-pkce-go.svg?status.svg)](https://godoc.org/github.com/kkdai/line-login-pkce-go)[![goreportcard.com](https://goreportcard.com/badge/github.com/kkdai/line-login-pkce-go)](https://goreportcard.com/report/github.com/kkdai/line-login-pkce-go)
 ![Go](https://github.com/kkdai/line-login-pkce-go/workflows/Go/badge.svg)


![](https://developers.line.biz/media/line-login/integrate-login-web/login-flow-web-0bc4c99d.png)

Refer LINE Developer Document "[Integrating LINE Login with your web app](https://developers.line.biz/en/docs/line-login/web/integrate-line-login/)" for more detail.


This sample code implement how to integrate LINE login to your website in Go. You can use this sample code to integrate LINE login in your web application. Also this also provide link a chatbot service when user use LINE login. Refer "[Linking a bot with your LINE Login channel](https://developers.line.biz/en/docs/line-login/web/link-a-bot/)".

Deploy on Heroku
=============

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

Before deploy this to your Heroku, you will need complete as follows:

- Create a LINE login channel. Remember its channel ID and channel secret.
- Create a LINE Message API channel. Remember its channel secret and token.
- Link the chatbot to the LINE login channel

Configuration
=============

Set the following environment variables (also defined in `app.json`):

| Variable | Description |
| --- | --- |
| `LINECORP_PLATFORM_CHANNEL_CHANNELID` | LINE Login - channel ID |
| `LINECORP_PLATFORM_CHANNEL_CHANNELSECRET` | LINE Login - channel secret |
| `LINECORP_PLATFORM_CHATBOT_CHANNELSECRET` | LINE Messaging API - channel secret |
| `LINECORP_PLATFORM_CHATBOT_CHANNELTOKEN` | LINE Messaging API - channel access token |
| `LINECORP_PLATFORM_SERVERURL` | Public URL of this server (no trailing slash); `<SERVERURL>/auth` must be registered as the callback URL of the LINE Login channel |
| `PORT` | HTTP port to listen on |

Run
=============

```
go run .
```

Or build with the included `Dockerfile`. The server must be started from the repository root, since it loads `login.tmpl`, `login_success.tmpl` and `static/`.

Endpoints
=============

| Path | Description |
| --- | --- |
| `/` | Login page (`login.tmpl`) |
| `/gotoauthpage` | Starts PKCE login with scope `profile`; optional `chatbot` form value sets `bot_prompt` |
| `/gotoauthOpenIDpage` | Starts PKCE login with scope `profile openid` |
| `/auth` | Callback: verifies `state`, exchanges code for token (PKCE), verifies/refreshes token, then shows the profile (from the profile API, or from the ID token with `nonce` and `exp` checks) in `login_success.tmpl` |
| `/logout` | Revokes the stored access token and redirects to `/` |
| `/callback` | Webhook of the linked chatbot; replies to text messages |
| `/static/` | Static assets (CSS, images) |

Note: state, nonce, code verifier and access token are kept in global variables, so this sample supports a single user at a time and is not suitable for production as is.

Dependencies
=============

- [line-login-sdk-go](https://github.com/kkdai/line-login-sdk-go) `v0.9.0` (requires Go 1.23+)
- [line-bot-sdk-go v7](https://github.com/line/line-bot-sdk-go)

Notes on SDK v0.9.0 usage:

- `GenerateNonce`, `GenerateCodeVerifier` and `GetPKCEWebLoginURL` now return an `error` that must be handled.
- `GetPKCEWebLoinURL` (misspelled) is deprecated; use `GetPKCEWebLoginURL`.
- ID token payload type is now `social.BasicPayload`, and `DecodePayloadWithOptions` is used to also verify `nonce` and `exp`.

License
=============

MIT License
