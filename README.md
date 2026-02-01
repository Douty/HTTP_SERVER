# Go-Raw-HTTP
A custom web server made purely with GO only using raw TCP packets


![Go Badge](https://img.shields.io/badge/go-1.21%2B-30d8be?style=for-the-badge&logo=go&logoSize=auto)

#### My inspiration for this project 

I wanted to gain a deeper understanding of how web servers like nginx and apache 
web server functioned at a low level and gain experience handling sockets 

#### Features To do list



- [x] **Manual TCP Handshake Handling**: Established using the `net` package without `net/http`.
- [x] **Custom HTTP Status Error Handling**
- [x] **HTTP Request Parser**: Decodes raw byte buffers into structured Go objects.
- [x] **HTTP Response Generator**: Manually constructs protocol-compliant response strings.
- [ ] **Dynamic Route Mapping**: API routing logic is currently in development.
- [ ] **Add CSS and JS support**: Expanding the `In-Memory Page Hashmap` to handle non-HTML assets.
- [ ] **Add TLS**: Implementing secure communication via `crypto/tls`.
- [ ] **Implement GitHub Actions CI/CD**: Automating testing and deployment workflows.
- [ ] **Containerize the server via Docker**: Creating a lightweight environment for deployment.
- [ ] **Implement LRU Cache**: Optimizing memory usage for high-traffic static assets.

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

```
#### Technical Challenges

