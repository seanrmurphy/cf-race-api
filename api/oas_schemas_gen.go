// Code generated by ogen, DO NOT EDIT.

package api

import (
	"time"
)

// Ref: #/components/schemas/Error
type Error struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

func (*Error) racePostRes()             {}
func (*Error) raceRaceIDGetRes()        {}
func (*Error) raceRaceIDResultsGetRes() {}
func (*Error) resultsPostRes()          {}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/RaceInfo
type RaceInfo struct {
	ID        OptInt      `json:"id"`
	Location  string      `json:"location"`
	CreatedAt OptDateTime `json:"created_at"`
	Name      string      `json:"name"`
	EventDate time.Time   `json:"event_date"`
	RunTypes  []string    `json:"run_types"`
}

// GetID returns the value of ID.
func (s *RaceInfo) GetID() OptInt {
	return s.ID
}

// GetLocation returns the value of Location.
func (s *RaceInfo) GetLocation() string {
	return s.Location
}

// GetCreatedAt returns the value of CreatedAt.
func (s *RaceInfo) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetName returns the value of Name.
func (s *RaceInfo) GetName() string {
	return s.Name
}

// GetEventDate returns the value of EventDate.
func (s *RaceInfo) GetEventDate() time.Time {
	return s.EventDate
}

// GetRunTypes returns the value of RunTypes.
func (s *RaceInfo) GetRunTypes() []string {
	return s.RunTypes
}

// SetID sets the value of ID.
func (s *RaceInfo) SetID(val OptInt) {
	s.ID = val
}

// SetLocation sets the value of Location.
func (s *RaceInfo) SetLocation(val string) {
	s.Location = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *RaceInfo) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetName sets the value of Name.
func (s *RaceInfo) SetName(val string) {
	s.Name = val
}

// SetEventDate sets the value of EventDate.
func (s *RaceInfo) SetEventDate(val time.Time) {
	s.EventDate = val
}

// SetRunTypes sets the value of RunTypes.
func (s *RaceInfo) SetRunTypes(val []string) {
	s.RunTypes = val
}

func (*RaceInfo) racePostRes()      {}
func (*RaceInfo) raceRaceIDGetRes() {}

type RaceRaceIDResultsGetOKApplicationJSON []RaceResult

func (*RaceRaceIDResultsGetOKApplicationJSON) raceRaceIDResultsGetRes() {}

// RaceRaceIDResultsPostCreated is response for RaceRaceIDResultsPost operation.
type RaceRaceIDResultsPostCreated struct{}

// Ref: #/components/schemas/RaceResult
type RaceResult struct {
	ID           OptInt    `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	RaceNumber   int       `json:"race_number"`
	RaceID       OptInt    `json:"race_id"`
	RaceLocation OptString `json:"race_location"`
	RunType      string    `json:"run_type"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
}

// GetID returns the value of ID.
func (s *RaceResult) GetID() OptInt {
	return s.ID
}

// GetFirstName returns the value of FirstName.
func (s *RaceResult) GetFirstName() string {
	return s.FirstName
}

// GetLastName returns the value of LastName.
func (s *RaceResult) GetLastName() string {
	return s.LastName
}

// GetRaceNumber returns the value of RaceNumber.
func (s *RaceResult) GetRaceNumber() int {
	return s.RaceNumber
}

// GetRaceID returns the value of RaceID.
func (s *RaceResult) GetRaceID() OptInt {
	return s.RaceID
}

// GetRaceLocation returns the value of RaceLocation.
func (s *RaceResult) GetRaceLocation() OptString {
	return s.RaceLocation
}

// GetRunType returns the value of RunType.
func (s *RaceResult) GetRunType() string {
	return s.RunType
}

// GetStartTime returns the value of StartTime.
func (s *RaceResult) GetStartTime() time.Time {
	return s.StartTime
}

// GetEndTime returns the value of EndTime.
func (s *RaceResult) GetEndTime() time.Time {
	return s.EndTime
}

// SetID sets the value of ID.
func (s *RaceResult) SetID(val OptInt) {
	s.ID = val
}

// SetFirstName sets the value of FirstName.
func (s *RaceResult) SetFirstName(val string) {
	s.FirstName = val
}

// SetLastName sets the value of LastName.
func (s *RaceResult) SetLastName(val string) {
	s.LastName = val
}

// SetRaceNumber sets the value of RaceNumber.
func (s *RaceResult) SetRaceNumber(val int) {
	s.RaceNumber = val
}

// SetRaceID sets the value of RaceID.
func (s *RaceResult) SetRaceID(val OptInt) {
	s.RaceID = val
}

// SetRaceLocation sets the value of RaceLocation.
func (s *RaceResult) SetRaceLocation(val OptString) {
	s.RaceLocation = val
}

// SetRunType sets the value of RunType.
func (s *RaceResult) SetRunType(val string) {
	s.RunType = val
}

// SetStartTime sets the value of StartTime.
func (s *RaceResult) SetStartTime(val time.Time) {
	s.StartTime = val
}

// SetEndTime sets the value of EndTime.
func (s *RaceResult) SetEndTime(val time.Time) {
	s.EndTime = val
}

// Ref: #/components/schemas/ResultsFilter
type ResultsFilter struct {
	Locations    []string    `json:"locations"`
	RaceType     []string    `json:"race_type"`
	TimeRangeMin OptDateTime `json:"time_range_min"`
	TimeRangeMax OptDateTime `json:"time_range_max"`
}

// GetLocations returns the value of Locations.
func (s *ResultsFilter) GetLocations() []string {
	return s.Locations
}

// GetRaceType returns the value of RaceType.
func (s *ResultsFilter) GetRaceType() []string {
	return s.RaceType
}

// GetTimeRangeMin returns the value of TimeRangeMin.
func (s *ResultsFilter) GetTimeRangeMin() OptDateTime {
	return s.TimeRangeMin
}

// GetTimeRangeMax returns the value of TimeRangeMax.
func (s *ResultsFilter) GetTimeRangeMax() OptDateTime {
	return s.TimeRangeMax
}

// SetLocations sets the value of Locations.
func (s *ResultsFilter) SetLocations(val []string) {
	s.Locations = val
}

// SetRaceType sets the value of RaceType.
func (s *ResultsFilter) SetRaceType(val []string) {
	s.RaceType = val
}

// SetTimeRangeMin sets the value of TimeRangeMin.
func (s *ResultsFilter) SetTimeRangeMin(val OptDateTime) {
	s.TimeRangeMin = val
}

// SetTimeRangeMax sets the value of TimeRangeMax.
func (s *ResultsFilter) SetTimeRangeMax(val OptDateTime) {
	s.TimeRangeMax = val
}

type ResultsPostOKApplicationJSON []RaceResult

func (*ResultsPostOKApplicationJSON) resultsPostRes() {}
