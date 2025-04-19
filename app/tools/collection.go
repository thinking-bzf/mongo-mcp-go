package tools

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"go.mongodb.org/mongo-driver/bson"
	"mcp/app/client"
	"strings"
)

type CollectionTool interface {
	// ListCollections get collections in mongodb
	ListCollections() (mcp.Tool, server.ToolHandlerFunc)
}

type collectionTool struct{}

func NewCollectionTool() CollectionTool {
	return &collectionTool{}
}

// ListCollections List all collections in mongodb
func (c collectionTool) ListCollections() (tool mcp.Tool, handler server.ToolHandlerFunc) {

	// MCP Tool
	tool = mcp.NewTool(
		"ListCollections",
		mcp.WithDescription("List all collections in mongodb"),
	)
	// handler
	// request is empty, this tool will return all collections in mongodb
	handler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		collections, err := client.DB.ListCollectionNames(ctx, bson.D{})

		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		if len(collections) == 0 {
			return mcp.NewToolResultText("No collections found"), nil
		}
		result := mcp.NewToolResultText("Collections: " + strings.Join(collections, ", "))
		return result, nil
	}
	return
}
