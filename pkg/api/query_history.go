package api

import (
	"net/http"

	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/web"
)

func (hs *HTTPServer) addToQueryHistory(c *models.ReqContext) response.Response {
	cmd := dtos.AddToQueryHistoryCmd{}
	if err := web.Bind(c.Req, &cmd); err != nil {
		return response.Error(http.StatusBadRequest, "bad request data", err)
	}

	hs.log.Debug("Received request to add query to query history", "query", cmd.Queries, "datasource", cmd.DataSourceUid)

	_, err := hs.QueryHistoryService.AddToQueryHistory(c.Req.Context(), c.SignedInUser, cmd.Queries, cmd.DataSourceUid)
	if err != nil {
		return response.Error(500, "Failed to create query history", err)
	}

	return response.Success("Query successfully added to query history")
}

func (hs *HTTPServer) searchInQueryHistory(c *models.ReqContext) response.Response {
	dataSourceUIDs := c.QueryStrings("dataSourceUid")
	query := c.Query("query")
	sort := c.Query("sort")

	queryHistory, err := hs.QueryHistoryService.ListQueryHistory(c.Req.Context(), c.SignedInUser, dataSourceUIDs, query, sort)
	if err != nil {
		return response.Error(500, "Failed to get query history", err)
	}

	return response.JSON(200, queryHistory)
}

func (hs *HTTPServer) deleteQueryFromQueryHistory(c *models.ReqContext) response.Response {
	queryUid := web.Params(c.Req)[":uid"]

	err := hs.QueryHistoryService.DeleteQuery(c.Req.Context(), c.SignedInUser, queryUid)
	if err != nil {
		return response.Error(500, "Failed to delete query from history", err)
	}

	return response.Success("Query successfully deleted from query history")
}

func (hs *HTTPServer) updateQueryInQueryHistory(c *models.ReqContext) response.Response {
	cmd := dtos.UpdateQueryInQueryHistoryCmd{}
	queryUid := web.Params(c.Req)[":uid"]

	if err := web.Bind(c.Req, &cmd); err != nil {
		return response.Error(http.StatusBadRequest, "bad request data", err)
	}

	query, err := hs.QueryHistoryService.GetQueryByUid(c.Req.Context(), c.SignedInUser, queryUid)
	err = hs.QueryHistoryService.UpdateComment(c.Req.Context(), c.SignedInUser, query, cmd.Comment)
	if err != nil {
		return response.Error(500, "Failed to update comment in query history", err)
	}

	return response.Success("Query comment successfully updated in query history")
}

func (hs *HTTPServer) starQueryInQueryHistory(c *models.ReqContext) response.Response {
	queryUid := web.Params(c.Req)[":uid"]

	err := hs.QueryHistoryService.StarQuery(c.Req.Context(), c.SignedInUser, queryUid)
	if err != nil {
		return response.Error(500, "Failed to star query in query history", err)
	}

	return response.Success("Query successfully starred")
}

func (hs *HTTPServer) unstarQueryInQueryHistory(c *models.ReqContext) response.Response {
	queryUid := web.Params(c.Req)[":uid"]

	err := hs.QueryHistoryService.UnstarQuery(c.Req.Context(), c.SignedInUser, queryUid)
	if err != nil {
		return response.Error(500, "Failed to unstar query in query history", err)
	}

	return response.Success("Query successfully unstarred")
}
