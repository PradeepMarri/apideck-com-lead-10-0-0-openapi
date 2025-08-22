package main

import (
	"github.com/lead-api/mcp-server/config"
	"github.com/lead-api/mcp-server/models"
	tools_leads "github.com/lead-api/mcp-server/tools/leads"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_leads.CreateLeadsallTool(cfg),
		tools_leads.CreateLeadsaddTool(cfg),
		tools_leads.CreateLeadsdeleteTool(cfg),
		tools_leads.CreateLeadsoneTool(cfg),
		tools_leads.CreateLeadsupdateTool(cfg),
	}
}
