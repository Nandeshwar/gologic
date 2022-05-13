package models

/*
{
  "response": {
    "header": {
      "next_id": 1069850178
    }
  },
  "schedules": [
    {
      "sourceId": 76950,
      "startTime": "21:00:00",
      "programId": "SH022888590000",
      "startDate": "2022-05-09",
*/
type CosmosApiModel struct {
	Response  Response   `json:"response"`
	Schedules []Schedule `json:"schedules"`
}

type Response struct {
	Header Header `json:"header"`
}

type Schedule struct {
	SourceId      int64  `json:"sourceId"`
	StartTime     string `json:"startTime"`
	StartDate     string `json:"startDate"`
	DurationInSec int64  `json:"duration"`
}

type Header struct {
	NextId int64 `json:"next_id"`
}
