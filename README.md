# YamatoMCP

You are using your own personal Yamato token, see line 61 in main.go

# Prerequisites
To run this project, you need to have the following prerequisites installed:
- Go - you can download it from [the official Go website](https://golang.org/dl/)
- Make (optional, for running commands)
- Claude Code CLI (for interacting with the MCP server)
- Yamato https://internaldocs.unity.com/yamato_continuous_integration/cli/

You'll be using your own personal Yamato token, see line 61 in `main.go`. Long-lived token will be added later.

# Installation
```bash
make run
```
# Usage Inspector
```bash
make help
```
You can uesa the inspector to see the available commands and their usage.

# Usage Claude Code CLI Streamable
```
claude mcp add --transport http yamato-mcp-streamable http://localhost:8080/mcp
Added HTTP MCP server yamato-mcp-streamable with URL: http://localhost:8080/mcp to local config
File modified: /Users/johan.eliasson/.claude.json [project: /Users/johan.eliasson/GolandProjects/YamatoMCP]
~/GolandProjects/YamatoMCP git:[main]
claude mcp list
Checking MCP server health...

yamato-mcp-streamable: http://localhost:8080/mcp (HTTP) - ✓ Connected
```

# Usage Claude Code CLI stdin/stdout
```
claude mcp add yamato-mcp-stdio go run  main.go
Added stdio MCP server yamato-mcp-stdio with command: go run main.go to local config
File modified: /Users/johan.eliasson/.claude.json [project: /Users/johan.eliasson/GolandProjects/YamatoMCP]
~/GolandProjectsclaude mcp list
Checking MCP server health...

yamato-mcp-streamable: http://localhost:8080/mcp (HTTP) - ✓ Connected
yamato-mcp-stdio: go run main.go - ✓ Connectedyamato-mcp-streamable: http://localhost:8080/mcp (HTTP) - ✓ Connected
```

## Example
```
> can you use the same server to get job definition history for .yamato/utr.yml#build_utr_win ?
> can you use the yamato mcp and fetch data for the job id 51154833?
```