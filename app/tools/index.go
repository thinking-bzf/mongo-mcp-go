package tools

import (
	"context"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mcp/app/client"
	"mcp/app/model"
)

type IndexTool interface {
	// ListIndexes get indexes in mongodb
	ListIndexes() (mcp.Tool, server.ToolHandlerFunc)
	// CreateIndex create index in mongodb
	CreateIndex() (mcp.Tool, server.ToolHandlerFunc)
	// DropIndex drop index in mongodb
	DropIndex() (mcp.Tool, server.ToolHandlerFunc)
}

type indexTool struct{}

func NewIndexTool() IndexTool {
	return &indexTool{}
}

// ListIndexes List all indexes in mongodb
func (c indexTool) ListIndexes() (tool mcp.Tool, handler server.ToolHandlerFunc) {

	// MCP Tool
	tool = mcp.NewTool(
		"ListIndexes",
		mcp.WithDescription("List all indexes in mongodb"),
		mcp.WithString("collection",
			mcp.Required(),
			mcp.Description("Collection name"),
		),
	)
	// handler
	handler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		collectionName := request.Params.Arguments["collection"].(string)
		cur, err := client.DB.Collection(collectionName).Indexes().List(ctx)
		if err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}
		defer cur.Close(ctx)

		var indexes []bson.M
		if err = cur.All(ctx, &indexes); err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}
		var result string
		for _, doc := range indexes {
			result += fmt.Sprintf("%v\n", doc)
		}
		return mcp.NewToolResultText(result), nil

	}
	return
}

// CreateIndex create index in mongodb
func (c indexTool) CreateIndex() (tool mcp.Tool, handler server.ToolHandlerFunc) {

	// MCP Tool
	tool = mcp.NewTool(
		"CreateIndex",
		mcp.WithDescription("create index in mongodb"),
		mcp.WithString("collection",
			mcp.Required(),
			mcp.Description("Collection name"),
		),
		mcp.WithObject("index_spec",
			mcp.Required(),
			mcp.Description("Index specification (e.g., { field: 1 } for ascending index)"),
		),
	)
	// handler
	handler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.CreateIndexRequest
		err := mapstructure.Decode(request.Params.Arguments, &req)

		if err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}

		res, err := client.DB.Collection(req.Collection).Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys: req.IndexSpec,
		})
		if err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}

		return mcp.NewToolResultText(fmt.Sprintf("Index created, Name: %v", res)), nil
	}
	return
}

// DropIndex drop index in mongodb
func (c indexTool) DropIndex() (tool mcp.Tool, handler server.ToolHandlerFunc) {

	// MCP Tool
	tool = mcp.NewTool(
		"DropIndex",
		mcp.WithDescription("drop index in mongodb"),
		mcp.WithString("collection",
			mcp.Required(),
			mcp.Description("Collection name"),
		),
		mcp.WithObject("index_name",
			mcp.Required(),
			mcp.Description("Name of the index to drop"),
		),
	)
	// handler
	handler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.DropIndexRequest
		err := mapstructure.Decode(request.Params.Arguments, &req)

		if err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}

		res, err := client.DB.Collection(req.Collection).Indexes().DropOne(ctx, req.IndexName)
		if err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}
		return mcp.NewToolResultText(fmt.Sprintf("Index dropped Successfully, result: %v", res)), nil
	}
	return
}
