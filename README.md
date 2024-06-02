# HTTP Requester Tool

This tool allows you to send custom malformed HTTP requests for web application testing and inspect the responses. It supports both HTTP and HTTPS protocols, and provides options to customize various parts of the request.

## Features

- Send HTTP/HTTPS requests to specified host and port.
- Include custom headers.
- Match a string in the response body.
- Optionally include the port number in the `Host` header.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/http-requester.git
   ```
2. Navigate to the project directory:
    ```sh
    cd http-requester
    ```
3. Build the project:
    ```sh
    go build -o requester
    ```

## Usage

### Options
    -https (bool): Use HTTPS (default: false)
    -host_port (bool): Include port after host in header (default: false)
    -host (string): Host name (default: "localhost")
    -port (int): Port number (default: 8000)
    -path (string): Request path (default: "/")
    -prefix (string): Host prefix (default: "none")
    -headers (string): Custom headers (default: "none")
    -mb (string): String to match in response body (default: "none")
