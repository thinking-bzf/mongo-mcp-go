package app

import (
	"github.com/mark3labs/mcp-go/server"
	"mcp/app/tools"
)

func AddTools(s *server.MCPServer) {
	// Add Collection tools to MCP server
	collTool := tools.NewCollectionTool()
	AddCollectionTools(s, collTool)

	// Add Document tools to MCP server
	docTool := tools.NewDocumentTool()
	AddDocumentTools(s, docTool)

	// Add Index tools to MCP server
	indexTool := tools.NewIndexTool()
	AddIndexTools(s, indexTool)

}

// AddCollectionTools adds collection tools to the MCP server
func AddCollectionTools(s *server.MCPServer, collTool tools.CollectionTool) {
	s.AddTool(collTool.ListCollections())
}

// AddDocumentTools adds collection tools to the MCP server
func AddDocumentTools(s *server.MCPServer, docTool tools.DocumentTool) {
	s.AddTool(docTool.Find())
	s.AddTool(docTool.Count())
	s.AddTool(docTool.InsertOne())
	s.AddTool(docTool.DeleteOne())
	s.AddTool(docTool.UpdateOne())
}

// AddIndexTools adds collection tools to the MCP server
func AddIndexTools(s *server.MCPServer, indexTool tools.IndexTool) {
	s.AddTool(indexTool.CreateIndex())
	s.AddTool(indexTool.ListIndexes())
	s.AddTool(indexTool.DropIndex())
}
