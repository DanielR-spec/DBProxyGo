# GoLang Proxy App

This repository contains a simple proxy application built in GoLang that facilitates routing and forwarding HTTP requests.

## Features

- **HTTP Proxy**: Handles HTTP requests and forwards them to the specified destination.
- **Customizable Routing**: Easily configure routing rules to direct traffic as needed.
- **Logging**: Keeps track of incoming requests and responses for monitoring purposes.

## Installation

To get started with the proxy app, follow these steps:

1. Clone the repository:

    ```
    git clone https://github.com/DanielR-spec/DBProxyGo.git
    ```

2. Navigate to the project directory:

    ```
    cd golang-proxy
    ```
3. Build the application:

    ```
    go build
    ```

4. Run the proxy app:

    ```
    ./golang-proxy
    ```

## Configuration

The proxy can be configured through the `config.json` file. Modify the file to specify routing rules and proxy settings.

```json
{
  "routes": [
    {
      "path": "/example",
      "destination": "http://example.com"
    },
    {
      "path": "/api",
      "destination": "http://api.example.com"
    }
  ],
  "port": 8080,
  "logLevel": "info"
}
```    
5. Run the proxy app:
