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

#### Usage & Setup

1. Clone the repo
```git clone https://github.com/Douty/HTTP_SERVER```
2. Start the server 
``` go run main.go ```
3. On your browser type in the url 
"http://localhost/"


## System Architecture & Request Lifecycle
```mermaid
sequenceDiagram
    autonumber
    participant Browser as Client (Browser)
    participant Main as Main (TCP Listener)
    participant Parser as Request Parser
    participant Resp as Response Builder
    participant Router as Router
    participant Pagemap as In-Memory Page Hashmap 

    Note over Browser, Main: 1. Connection & Data Receipt
    Browser->>Main: TCP Connection (Raw Bytes)
    Main->>Parser: Send Buffer 
    
    Note over Parser: 2. Protocol Parsing
    Parser-->>Main: Return Request Struct
    
    Note over Main, Resp: 3. Orchestration Phase
    Main->>Resp: HandleRequest(Request)
    
    Note over Resp, Pagemap: 4. Routing & Content Retrieval
    Resp->>Router: Check to see if Page/API Exists
    Router->>Pagemap: Check for valid pages or API Logic
    Note over Pagemap, Router: Status Package provides code
    Pagemap-->>Router: Content & Status Code
    
    Note over Router, Resp: 5. Response Generation
    Router->>Resp: Return Data
    Note over Resp: Build Status Line & Headers
    Resp-->>Main: Raw Response String
    
    Note over Main, Browser: 6. Transmission
    Main->>Browser: conn.Write(response)
    Main->>Browser: conn.Close()

