# YamatoMCP

You are using your own personal Yamato token, see line 61 in main.go

To get started, you need to have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).
# Installation
```bash
make run
```
# Usage Inspector
```bash
make help
```

# Usage Claude Code CLI
```
claude mcp add --transport http yamato-mcp-streamable http://localhost:8080/mcp
Added HTTP MCP server yamato-mcp-streamable with URL: http://localhost:8080/mcp to local config
File modified: /Users/johan.eliasson/.claude.json [project: /Users/johan.eliasson/GolandProjects/YamatoMCP]
~/GolandProjects/YamatoMCP git:[main]
claude mcp list
Checking MCP server health...

yamato-mcp-streamable: http://localhost:8080/mcp (HTTP) - ✓ Connected
```
## Example
```
> can you use the same server to get job definition history for .yamato/utr.yml#build_utr_win ?
> can you use the yamato mcp and fetch data for the job id 51154833?
```