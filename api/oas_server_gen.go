// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// RacePost implements POST /race operation.
	//
	// Creates a new race.
	//
	// POST /race
	RacePost(ctx context.Context, req *RaceInfo) (RacePostRes, error)
	// RaceRaceIDGet implements GET /race/{race_id} operation.
	//
	// Returns basic info pertaining to a race.
	//
	// GET /race/{race_id}
	RaceRaceIDGet(ctx context.Context, params RaceRaceIDGetParams) (RaceRaceIDGetRes, error)
	// RaceRaceIDResultsGet implements GET /race/{race_id}/results operation.
	//
	// Returns ordered set of results for a given race.
	//
	// GET /race/{race_id}/results
	RaceRaceIDResultsGet(ctx context.Context, params RaceRaceIDResultsGetParams) (RaceRaceIDResultsGetRes, error)
	// RaceRaceIDResultsPost implements POST /race/{race_id}/results operation.
	//
	// Adds a set of results to a given race.
	//
	// POST /race/{race_id}/results
	RaceRaceIDResultsPost(ctx context.Context, req []RaceResult, params RaceRaceIDResultsPostParams) error
	// ResultsPost implements POST /results operation.
	//
	// Returns a set of results with the given filter.
	//
	// POST /results
	ResultsPost(ctx context.Context, req *ResultsFilter) (ResultsPostRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
