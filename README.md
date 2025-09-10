# Zapper

Zapper is an API for configuring and managing WhatsApp bots powered by Large Language Models (LLMs). It provides a flexible framework to build intelligent chatbots that can handle conversations, automate tasks, and integrate with various LLM providers.

## Features

### Implemented

### Upcoming
- [ ] WhatsApp integration using the whatsmeow library
- [ ] LLM client for generating responses and managing bot behavior
- [ ] RESTful API server for bot configuration and control
- [ ] CLI tool for local development and testing
- [ ] Structured logging system
- [ ] Advanced bot configuration endpoints (e.g., dynamic prompts, user sessions)
- [ ] Support for multiple LLM providers (e.g., OpenAI, Anthropic)
- [ ] Webhook integration for real-time event handling
- [ ] Persistent storage with database support (e.g., PostgreSQL)
- [ ] Authentication and authorization for API access

## Installation

1. Clone the repository:
   ```
   git clone &lt;repo-url&gt;
   cd zapper
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Build the project:
   ```
   go build -o bin/zapper ./cmd/cli
   go build -o bin/api ./cmd/api
   ```

## Usage

### CLI
Run the CLI tool for local bot management:
```
./bin/zapper --help
```

### API
Start the API server:
```
./bin/api
```
The server listens on `:8080` by default. Use endpoints like `/configure` to set up bots.

For detailed API documentation, refer to the [internal/handler](internal/handler) package.

## Contributing
Pull requests are welcome. For major changes, please open an issue first.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
