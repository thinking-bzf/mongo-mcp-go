package tools

import (
	"context"
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mcp/app/client"
	"mcp/app/model"
)

type DocumentTool interface {
	// Find get document in collection
	Find() (mcp.Tool, server.ToolHandlerFunc)
	// Count get documents count in collection
	Count() (mcp.Tool, server.ToolHandlerFunc)
	// InsertOne insert one document in collection
	InsertOne() (mcp.Tool, server.ToolHandlerFunc)
	// DeleteOne insert one document in collection
	DeleteOne() (mcp.Tool, server.ToolHandlerFunc)
	// UpdateOne insert one document in collection
	UpdateOne() (mcp.Tool, server.ToolHandlerFunc)
}

type documentTool struct{}

func NewDocumentTool() DocumentTool {
	return &documentTool{}
}

// Find get document in collection
func (c documentTool) Find() (tool mcp.Tool, handler server.ToolHandlerFunc) {

	// MCP Tool
	tool = mcp.NewTool(
		"Find",
		mcp.WithDescription("Query documents in a collection using MongoDB query syntax"),
		mcp.WithString("collection",
			mcp.Required(),
			mcp.Description("Collection name to query"),
		),
		mcp.WithObject("filter",
			mcp.Description("MongoDB query filter"),
			mcp.DefaultString("{}"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Limit the number of documents to return"),
			mcp.DefaultNumber(10),
			mcp.Min(1),
			mcp.Max(1000),
		),
		mcp.WithObject("projection",
			mcp.Description("MongoDB projection filter"),
			mcp.DefaultString("{}"),
		),
	)
	// handler
	handler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.FindDocumentRequest
		err := mapstructure.Decode(request.Params.Arguments, &req)
		if err != nil {
			return mcp.NewToolResultText("Parse request failed"), err
		}

		log.Printf(fmt.Sprintf("Find document in collection: %s, filter: %s, limit: %d", req.Collection, req.Filter, req.Limit))

		cur, err := client.DB.Collection(req.Collection).Find(ctx, req.Filter, &options.FindOptions{
			Limit:      &req.Limit,
			Projection: req.Projection,
		})
		defer cur.Close(ctx)

		if err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}

		var documents []bson.M
		if err = cur.All(ctx, &documents); err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}
		if len(documents) == 0 {
			return mcp.NewToolResultText("No documents found"), nil
		}
		var result string
		for _, doc := range documents {
			result += fmt.Sprintf("%v\n", doc)
		}
		return mcp.NewToolResultText(result), nil
	}
	return
}

// Count get documents count in collection
func (c documentTool) Count() (tool mcp.Tool, handler server.ToolHandlerFunc) {

	// MCP Tool
	tool = mcp.NewTool(
		"Count",
		mcp.WithDescription("Count documents in a collection using MongoDB query syntax"),
		mcp.WithString("collection",
			mcp.Required(),
			mcp.Description("Collection name to query"),
		),
		mcp.WithObject("filter",
			mcp.Description("MongoDB query filter"),
			mcp.DefaultString("{}"),
		),
		mcp.WithObject("projection",
			mcp.Description("MongoDB projection filter"),
			mcp.DefaultString("{}"),
		),
	)
	// handler
	handler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.CountDocumentRequest
		err := mapstructure.Decode(request.Params.Arguments, &req)
		if err != nil {
			return mcp.NewToolResultText("Parse request failed"), err
		}

		log.Printf(fmt.Sprintf("Count document in collection: %s, filter: %s", req.Collection, req.Filter))

		count, err := client.DB.Collection(req.Collection).CountDocuments(ctx, req.Filter, &options.CountOptions{
			Limit: &req.Limit,
		})
		if err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}

		return mcp.NewToolResultText(fmt.Sprintf("Count documents success, count: %d", count)), nil
	}
	return
}

// InsertOne insert one document in collection
func (c documentTool) InsertOne() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	// MCP Tool
	tool = mcp.NewTool(
		"InertOne",
		mcp.WithDescription("Insert a single document into a collection"),
		mcp.WithString("collection",
			mcp.Required(),
			mcp.Description("Collection name to insert"),
		),
		mcp.WithString("document",
			mcp.Required(),
			mcp.Description("document to insert"),
			mcp.DefaultString("{}"),
		),
	)
	// handler
	handler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.InsertDocumentRequest

		err := mapstructure.Decode(request.Params.Arguments, &req)
		if err != nil {
			return mcp.NewToolResultText("Parse request failed"), err
		}

		log.Printf(fmt.Sprintf("Insert document in collection: %s, document: %s", req.Collection, req.Document))

		// Convert the document string to a BSON document
		var document bson.M
		err = bson.UnmarshalExtJSON([]byte(req.Document), true, &document)
		if err != nil {
			return mcp.NewToolResultText("Parse document failed"), err
		}
		res, err := client.DB.Collection(req.Collection).InsertOne(ctx, document)
		if err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}

		if res.InsertedID == nil {
			return mcp.NewToolResultText("Insert document failed"), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("Insert document success, id: %v", res.InsertedID)), nil
	}
	return
}

// DeleteOne insert one document in collection
func (c documentTool) DeleteOne() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	// MCP Tool
	tool = mcp.NewTool(
		"DeleteOne",
		mcp.WithDescription("Delete a single document into a collection"),
		mcp.WithString("collection",
			mcp.Required(),
			mcp.Description("Collection name to delete"),
		),
		mcp.WithObject("filter",
			mcp.Required(),
			mcp.Description("Filter to identify document"),
		),
	)
	// handler
	handler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.DeleteDocumentRequest
		err := mapstructure.Decode(request.Params.Arguments, &req)
		if err != nil {
			return mcp.NewToolResultText("Parse request failed"), err
		}

		log.Printf(fmt.Sprintf("Delete document in collection: %s, filter: %s", req.Collection, req.Filter))

		res, err := client.DB.Collection(req.Collection).DeleteOne(ctx, req.Filter)
		if err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}
		if res.DeletedCount == 0 {
			return mcp.NewToolResultText("No documents deleted"), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("Delete document success, deleted count: %d", res.DeletedCount)), nil
	}
	return
}

// UpdateOne insert one document in collection
func (c documentTool) UpdateOne() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	// MCP Tool
	tool = mcp.NewTool(
		"UpdateOne",
		mcp.WithDescription("update a single document into a collection"),
		mcp.WithString("collection",
			mcp.Required(),
			mcp.Description("Collection name to update"),
		),
		mcp.WithObject("filter",
			mcp.Required(),
			mcp.Description("Filter to identify document"),
		),
		mcp.WithObject("update",
			mcp.Required(),
			mcp.Description("Update operations to apply"),
		),
	)
	// handler
	handler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.UpdateDocumentRequest
		err := mapstructure.Decode(request.Params.Arguments, &req)
		if err != nil {
			return mcp.NewToolResultText("Parse request failed"), err
		}

		log.Printf(fmt.Sprintf("Update document in collection: %s, filter: %s", req.Collection, req.Filter))

		res, err := client.DB.Collection(req.Collection).UpdateOne(ctx, req.Filter, req.Update)
		if err != nil {
			return mcp.NewToolResultText(err.Error()), err
		}
		if res.MatchedCount == 0 {
			return mcp.NewToolResultText("No documents matched"), nil
		}
		if res.ModifiedCount == 0 {
			return mcp.NewToolResultText("No documents updated"), nil
		}
		return mcp.NewToolResultText(
			fmt.Sprintf("Update document success, matched count: %d, modified count: %d, upsertedId %d",
				res.MatchedCount, res.ModifiedCount, res.UpsertedID,
			),
		), nil
	}
	return
}
