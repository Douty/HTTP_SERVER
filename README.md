# Go-Raw-HTTP
A custom web server made purely with GO only using raw TCP packets

#### My inspiration for this project 

I wanted to gain a deeper understanding of how web servers like nginx and apache 
web server functioned at a low level and gain experience handling sockets 

#### Features To do list

x = completed, / = in progress, [] = not started

[x] Manual TCP Handshake Handling 
[x] Custom HTTP Status Error Handling 
[x] HTTP Request Parser
[x] HTTP Response Generator  
[/] Dynamic Route Mapping (API routing not implemented yet)
[/] Add CSS and JS support
[]  Add TLS
[]  Implement Github Actions CI/CD 
[]  Containerize the server via Docker
[]  Implement LRU Cache 

#### Useage & Setup

1. Clone the repo
```git clone https://github.com/Douty/HTTP_SERVER```
2. Start the server 
``` go run main.go ```
3. On your browser type in the url 
"http://localhost/"




