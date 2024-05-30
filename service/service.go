package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/seanrmurphy/cf-race-api/api"
	"github.com/seanrmurphy/cf-race-api/dbqueries"
	"github.com/syumai/workers/cloudflare/d1"
)

type ServerImplementation struct {
	DatabaseEnvVarName string
	initialized        bool
	db                 *sql.DB
}

func (p *ServerImplementation) initialize(ctx context.Context) error {
	// initialize DB.
	// D1 connector requires request's context to initialize DB.
	c, err := d1.OpenConnector(p.DatabaseEnvVarName)
	if err != nil {
		return fmt.Errorf("Internal server error: service unavailable: %w", err)
	}
	// use sql.OpenDB instead of sql.Open.
	p.db = sql.OpenDB(c)
	p.initialized = true
	return nil
}

// RacePost implements POST /race operation.
//
// Creates a new race.
//
// POST /race
func (p *ServerImplementation) RacePost(ctx context.Context, req *api.RaceInfo) (api.RacePostRes, error) {
	if p.initialized == false {
		err := p.initialize(ctx)
		if err != nil {
			return nil, fmt.Errorf("Internal server error: service unavailable: %w", err)
		}
	}
	queries := dbqueries.New(p.db)

	runTypesMarshalled, _ := json.Marshal(req.RunTypes)
	addRaceParams := dbqueries.AddRaceParams{
		Name:      req.Name,
		Location:  req.Location,
		RunTypes:  string(runTypesMarshalled),
		EventDate: req.EventDate.UnixMicro(),
		CreatedAt: time.Now().UnixMicro(),
	}

	dbRace, err := queries.AddRace(ctx, addRaceParams)
	if err != nil {
		if err == sql.ErrNoRows {
			return &api.RaceInfo{}, nil
		}
		// TODO: need to add error handling for the case in which the uniqueness constraint
		// prevents creation of the user; that generates an error, but I don't know how to
		// catch it precisely...
		log.Printf("error getting race info: %v\n", err.Error())
		response := api.Error{
			Message: err.Error(),
		}
		return &response, nil
	}

	// map user as per the db record to user as per the API.
	var runTypesUnmarshalled []string
	_ = json.Unmarshal([]byte(dbRace.RunTypes), &runTypesUnmarshalled)
	returnVal := api.RaceInfo{
		ID:        api.NewOptInt(int(dbRace.ID)),
		Name:      dbRace.Name,
		Location:  dbRace.Location,
		EventDate: time.UnixMicro(dbRace.EventDate),
		RunTypes:  runTypesUnmarshalled,
		CreatedAt: api.NewOptDateTime(time.UnixMicro(dbRace.CreatedAt)),
	}

	return &returnVal, nil
}

// RaceRaceIDGet implements GET /race/{race_id} operation.
//
// Returns basic info pertaining to a race.
//
// GET /race/{race_id}
func (p *ServerImplementation) RaceRaceIDGet(ctx context.Context, params api.RaceRaceIDGetParams) (api.RaceRaceIDGetRes, error) {
	if p.initialized == false {
		err := p.initialize(ctx)
		if err != nil {
			return nil, fmt.Errorf("Internal server error: service unavailable: %w", err)
		}
	}
	queries := dbqueries.New(p.db)

	dbRace, err := queries.GetRaceInfoById(ctx, int64(params.RaceID))
	if err != nil {
		if err == sql.ErrNoRows {
			return &api.RaceInfo{}, nil
		}
		// TODO: need to add error handling for the case in which the uniqueness constraint
		// prevents creation of the user; that generates an error, but I don't know how to
		// catch it precisely...
		log.Printf("error getting race info: %v\n", err.Error())
		response := api.Error{
			Message: err.Error(),
		}
		return &response, nil
	}

	// map user as per the db record to user as per the API.
	var runTypesUnmarshalled []string
	_ = json.Unmarshal([]byte(dbRace.RunTypes), &runTypesUnmarshalled)
	if err != nil {
		log.Printf("error unmarshalling run types: %v\n", err.Error())
		response := api.Error{
			Message: err.Error(),
		}
		return &response, nil
	}

	returnVal := api.RaceInfo{
		ID:        api.NewOptInt(int(dbRace.ID)),
		Name:      dbRace.Name,
		Location:  dbRace.Location,
		EventDate: time.UnixMicro(dbRace.EventDate),
		RunTypes:  runTypesUnmarshalled,
		CreatedAt: api.NewOptDateTime(time.UnixMicro(dbRace.CreatedAt)),
	}

	return &returnVal, nil
}

// RaceRaceIDResultsGet implements GET /race/{race_id}/results operation.
//
// Returns ordered set of results for a given race.
//
// GET /race/{race_id}/results
func (p *ServerImplementation) RaceRaceIDResultsGet(ctx context.Context, params api.RaceRaceIDResultsGetParams) (api.RaceRaceIDResultsGetRes, error) {
	if p.initialized == false {
		err := p.initialize(ctx)
		if err != nil {
			return nil, fmt.Errorf("Internal server error: service unavailable: %w", err)
		}
	}
	queries := dbqueries.New(p.db)

	raceResults, err := queries.GetRaceResults(ctx, int64(params.RaceID))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		// TODO: need to add error handling for the case in which the uniqueness constraint
		// prevents creation of the user; that generates an error, but I don't know how to
		// catch it precisely...
		log.Printf("error getting race info: %v\n", err.Error())
	}

	var returnVal []api.RaceResult
	for _, result := range raceResults {
		returnVal = append(returnVal, api.RaceResult{
			FirstName:  result.FirstName.String,
			LastName:   result.LastName.String,
			RaceNumber: int(result.RaceNumber.Int64),
			StartTime:  time.UnixMicro(result.StartTime),
			EndTime:    time.UnixMicro(result.EndTime),
			RunType:    result.RunType,
		})
	}

	returnValTyped := api.RaceRaceIDResultsGetOKApplicationJSON(returnVal)
	return &returnValTyped, nil
}

// RaceRaceIDResultsPost implements POST /race/{race_id}/results operation.
//
// Adds a set of results to a given race.
//
// POST /race/{race_id}/results
func (p *ServerImplementation) RaceRaceIDResultsPost(ctx context.Context, req []api.RaceResult, params api.RaceRaceIDResultsPostParams) error {
	if p.initialized == false {
		err := p.initialize(ctx)
		if err != nil {
			return fmt.Errorf("Internal server error: service unavailable: %w", err)
		}
	}
	queries := dbqueries.New(p.db)

	for _, result := range req {
		addRaceResultParams := dbqueries.AddRaceResultParams{
			FirstName: sql.NullString{String: result.FirstName},
			RaceID:    int64(params.RaceID),
			StartTime: result.StartTime.UnixMicro(),
			EndTime:   result.EndTime.UnixMicro(),
			RunType:   result.RunType,
		}
		_, err := queries.AddRaceResult(ctx, addRaceResultParams)
		if err != nil {
			if err == sql.ErrNoRows {
				return err
			}
			// TODO: need to add error handling for the case in which the uniqueness constraint
			// prevents creation of the user; that generates an error, but I don't know how to
			// catch it precisely...
			log.Printf("error getting race info: %v\n", err.Error())
		}

	}
	return nil
}

func (p *ServerImplementation) ResultsPost(ctx context.Context, req *api.ResultsFilter) (api.ResultsPostRes, error) {
	// the filter does not work as yet as it needs a little more
	if p.initialized == false {
		err := p.initialize(ctx)
		if err != nil {
			return nil, fmt.Errorf("Internal server error: service unavailable: %w", err)
		}
	}
	queries := dbqueries.New(p.db)

	res, _ := queries.GetFilteredRaceResults(ctx)
	var returnVal []api.RaceResult
	for _, result := range res {
		returnVal = append(returnVal, api.RaceResult{
			FirstName:  result.FirstName.String,
			LastName:   result.LastName.String,
			RaceNumber: int(result.RaceNumber.Int64),
			StartTime:  time.UnixMicro(result.StartTime),
			EndTime:    time.UnixMicro(result.EndTime),
			RunType:    result.RunType,
		})
	}
	returnValTyped := api.ResultsPostOKApplicationJSON(returnVal)
	return &returnValTyped, nil
}
