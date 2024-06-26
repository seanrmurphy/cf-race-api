// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// RacePost implements POST /race operation.
//
// Creates a new race.
//
// POST /race
func (UnimplementedHandler) RacePost(ctx context.Context, req *RaceInfo) (r RacePostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// RaceRaceIDGet implements GET /race/{race_id} operation.
//
// Returns basic info pertaining to a race.
//
// GET /race/{race_id}
func (UnimplementedHandler) RaceRaceIDGet(ctx context.Context, params RaceRaceIDGetParams) (r RaceRaceIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// RaceRaceIDResultsGet implements GET /race/{race_id}/results operation.
//
// Returns ordered set of results for a given race.
//
// GET /race/{race_id}/results
func (UnimplementedHandler) RaceRaceIDResultsGet(ctx context.Context, params RaceRaceIDResultsGetParams) (r RaceRaceIDResultsGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// RaceRaceIDResultsPost implements POST /race/{race_id}/results operation.
//
// Adds a set of results to a given race.
//
// POST /race/{race_id}/results
func (UnimplementedHandler) RaceRaceIDResultsPost(ctx context.Context, req []RaceResult, params RaceRaceIDResultsPostParams) error {
	return ht.ErrNotImplemented
}

// ResultsPost implements POST /results operation.
//
// Returns a set of results with the given filter.
//
// POST /results
func (UnimplementedHandler) ResultsPost(ctx context.Context, req *ResultsFilter) (r ResultsPostRes, _ error) {
	return r, ht.ErrNotImplemented
}
