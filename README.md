# SERVER_PROXY

**Seamless Connections, Limitless Possibilities**

Built with the tools and technologies you love.

---

## Table of Contents

- [Overview](#overview)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Testing](#testing)

---

## Overview

`server_proxy` is a powerful and flexible proxy server designed to handle HTTP, HTTPS, and TCP traffic efficiently. It integrates protocol detection, secure tunneling, and connection management to provide a seamless proxy experience for developers and network engineers.

### Why server_proxy?

This project simplifies the deployment of reliable and secure proxy services. The core features include:

- 🔀 **Protocol Detection**  
  Automatically identifies HTTP, CONNECT, and raw TCP requests for precise routing.

- 🌐 **HTTPS Tunneling Support**  
  Enables encrypted traffic forwarding through transparent proxying.

- ⚡ **TCP Forwarding**  
  Facilitates bidirectional TCP data transfer for high-performance tunneling.

- ⚙️ **Configurable CLI**  
  Easily customize network settings, timeouts, and verbosity for flexible deployments.

- 📄 **Automated Documentation**  
  Converts Markdown docs into man pages, enhancing accessibility and maintainability.

---

## Getting Started

### Prerequisites

This project requires the following dependencies:

- **Programming Language:** Go  
- **Package Manager:** Go modules

---

### Installation

Build `server_proxy` from the source and install dependencies:

1. **Clone the repository**
   ```bash
   git clone https://github.com/ask-elad/server_proxy
   cd server_proxy
   ```
2. **Install the dependencies**
   ```
   go mod tidy
   ```
3. **Run the proj**
   ```
   go run ./cmd/proxy
   ```

The flags 
```
--listen string
    Address to listen on (default :8080)

--dial-timeout duration
    Timeout for outbound connections (e.g. 10s)

--conn-timeout duration
    Maximum lifetime of a client connection (e.g. 5m)

--verbose
    Enable verbose logging

--blocked string
    Path to blocked domains/IPs file

-h, --help
    Show help for proxy

```

# Testing

The proxy was tested manually and programmatically to validate correctness, stability, and flag behavior across HTTP, HTTPS, and TCP traffic.

---

**Test**
1)Started the proxy with default settings
Sent HTTP requests through the proxy using `curl`

```bash
go run ./cmd/proxy --listen :8080
curl -x http://localhost:8080 http://example.com
```

Result

Requests were successfully forwarded to the upstream server
Responses were streamed back correctly
No request buffering or data corruption observed


2)Verified HTTPS support via the CONNECT method
Tested with multiple HTTPS endpoints
```bash
curl -x http://localhost:8080 https://example.com
```

Result

CONNECT tunnels were established correctly
TLS handshake occurred end-to-end without interception
Encrypted traffic passed transparently through the proxy
