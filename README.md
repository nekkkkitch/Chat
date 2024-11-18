# Installation and launching
1. Clone project
2. cd in terminal to /chat
3. run in terminal
   ```
   make buildbuilder
   ```
4. run in terminal
    ```
   make start
   ```
# Available requests
I suggest you to use desktop version of [postman](https://web.postman.co)
IP is localhost:8082
### HTTP Requests
- /register

`Registers new user. Requiers json body with login and password parameters. Response contains headers X-Access-Token and X-Refresh-Token`
- /login

`Login user. Requiers json body with login and password parameters. Response contains headers X-Access-Token and X-Refresh-Token`
- /refresh

`Returns new tokens. Requiers headers X-Access-Token and X-Refresh-Token with corresponding values. Response contains headers X-Access-Token and X-Refresh-Token`
- /ping

`Pings server. Response contain text`
### Websocket
- /chat

`Establishes websocket connection. Requiers url parameter access_token. Messages should contain json body with reciever and text parameters`
