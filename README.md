# Installation and launching
1. Clone project
2. cd in terminal to /chat
3. run in terminal(Docker engine should be running and Makefile.exe should be in your Path)
   ```
   make buildbuilder
   ```
4. run in terminal
    ```
   make start
   ```
# Available requests
I suggest you to use desktop version of [postman](https://web.postman.co) OR run files from /test folder

IP is localhost:8082
### HTTP Requests
- POST /register

`Registers new user. Requiers json body with login and password parameters. Response contains headers X-Access-Token and X-Refresh-Token`
- POST /login

`Login user. Requiers json body with login and password parameters. Response contains headers X-Access-Token and X-Refresh-Token`
- GET /refresh

`Returns new tokens. Requiers headers X-Access-Token and X-Refresh-Token with corresponding values. Response contains headers X-Access-Token and X-Refresh-Token`
- GET /getchat?user=

`Returns chat with given user. Response contains text with ALL messages between requester and user`
- GET /ping

`Pings server. Response contain text`
### Websocket
- /chat?access_token=

`Establishes websocket connection. Requiers url parameter access_token. Messages should contain json body with reciever and text parameters`
