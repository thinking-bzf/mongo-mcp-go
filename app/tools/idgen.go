package tools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mcp/app/client"
)

type IdGenerateTool interface {
	Generate() (mcp.Tool, server.ToolHandlerFunc)
}

type idGenerateTool struct {
	EntityPrefixMap map[string]interface{}
}

func NewIdGenerateTool() IdGenerateTool {
	return &idGenerateTool{
		EntityPrefixMap: map[string]interface{}{
			"task":       "TSK",
			"lesson":     "LSN",
			"course":     "CRS",
			"experiment": "EXP",
		},
	}
}

func (i idGenerateTool) Generate() (tool mcp.Tool, handler server.ToolHandlerFunc) {

	var prefixNameSet []string
	for key := range i.EntityPrefixMap {
		prefixNameSet = append(prefixNameSet, key)
	}
	tool = mcp.NewTool(
		"entity_id_generator",
		mcp.WithDescription("generate id for different type entity"),
		mcp.WithString("entity_type",
			mcp.Required(),
			mcp.Description(fmt.Sprintf("type of entity to generate id, it could be one of %s", prefixNameSet)),
		),
	)
	handler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		entityType := request.Params.Arguments["entity_type"].(string)
		log.Printf("Generate id for entity: %s", entityType)
		var entityPrefix string
		if value, ok := i.EntityPrefixMap[entityType]; ok {
			entityPrefix = value.(string)
			//return mcp.NewToolResultText(prefix.(string) + "-" + mcp.GenerateId()), nil
		} else {
			return mcp.NewToolResultText(fmt.Sprintf("entity type %s not found, please use one of %s", entityType, prefixNameSet)), nil
		}
		counterID, err := i.getNextSequence(ctx, entityPrefix)
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("failed to get next sequence: %v", err)), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("%s-%04d", entityPrefix, counterID)), nil
	}
	return
}

func (i idGenerateTool) getNextSequence(ctx context.Context, counterIDType string) (int, error) {
	filter := bson.M{"id_type": counterIDType}
	update := bson.M{"$inc": bson.M{"sequence": 1}}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var result struct {
		Sequence int `bson:"sequence"`
	}
	err := client.DB.Collection("counters").FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.Sequence, nil
}
