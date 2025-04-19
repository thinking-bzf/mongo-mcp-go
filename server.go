package main

import (
	"fmt"
	"github.com/mark3labs/mcp-go/server"
	"mcp/app"
	"mcp/app/client"
	"mcp/app/configs"
)

func main() {
	// Connect mongo client
	config := configs.LoadConfig(".", "")
	client.ConnectMongo(config.Mongo)

	MCPConfig := config.MCP

	// Create a new MCP server
	s := server.NewMCPServer(
		MCPConfig.Name,
		MCPConfig.Version,
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)
	// 添加工具到 MCP 服务器中
	app.AddTools(s)
	// Start the server
	if MCPConfig.SSE {
		// SSE server
		sse := server.NewSSEServer(s, server.WithBaseURL(MCPConfig.BaseUrl))
		if err := sse.Start(MCPConfig.Address); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	} else {
		// stdio server
		if err := server.ServeStdio(s); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}

}
