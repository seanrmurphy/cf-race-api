// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// RaceRaceIDGetParams is parameters of GET /race/{race_id} operation.
type RaceRaceIDGetParams struct {
	// ID of Race.
	RaceID int
}

func unpackRaceRaceIDGetParams(packed middleware.Parameters) (params RaceRaceIDGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "race_id",
			In:   "path",
		}
		params.RaceID = packed[key].(int)
	}
	return params
}

func decodeRaceRaceIDGetParams(args [1]string, argsEscaped bool, r *http.Request) (params RaceRaceIDGetParams, _ error) {
	// Decode path: race_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "race_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.RaceID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "race_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RaceRaceIDResultsGetParams is parameters of GET /race/{race_id}/results operation.
type RaceRaceIDResultsGetParams struct {
	// ID of Race.
	RaceID int
}

func unpackRaceRaceIDResultsGetParams(packed middleware.Parameters) (params RaceRaceIDResultsGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "race_id",
			In:   "path",
		}
		params.RaceID = packed[key].(int)
	}
	return params
}

func decodeRaceRaceIDResultsGetParams(args [1]string, argsEscaped bool, r *http.Request) (params RaceRaceIDResultsGetParams, _ error) {
	// Decode path: race_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "race_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.RaceID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "race_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RaceRaceIDResultsPostParams is parameters of POST /race/{race_id}/results operation.
type RaceRaceIDResultsPostParams struct {
	// ID of Race.
	RaceID int
}

func unpackRaceRaceIDResultsPostParams(packed middleware.Parameters) (params RaceRaceIDResultsPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "race_id",
			In:   "path",
		}
		params.RaceID = packed[key].(int)
	}
	return params
}

func decodeRaceRaceIDResultsPostParams(args [1]string, argsEscaped bool, r *http.Request) (params RaceRaceIDResultsPostParams, _ error) {
	// Decode path: race_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "race_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.RaceID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "race_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}