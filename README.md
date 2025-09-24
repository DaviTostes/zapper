# Zapper

Zapper is an open source web app for configuring and managing WhatsApp bots powered by Large Language Models (LLMs). 
It provides a flexible framework to build intelligent chatbots that can handle conversations, automate tasks, 
and integrate with various LLM providers.

## Features

### Implemented

### Upcoming
- [~] WhatsApp integration using the whatsmeow library
- [ ] LLM client for generating responses and managing bot behavior
- [ ] Structured logging system
- [ ] Support for multiple LLM providers (e.g., OpenAI, Anthropic)
- [ ] Webhook integration for real-time event handling

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/davitostes/zapper;
   cd zapper
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Build the project:
   ```
   go build -o bin/zapper ./cmd/server
   ```

## Usage

### Server
Run the server for local bot management:
```
./bin/zapper --help
```

For detailed API documentation, refer to the [internal/handler](internal/handler) package.

## Contributing
Pull requests are welcome. For major changes, please open an issue first.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
