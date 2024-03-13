package main

type Conditions struct {
    AgeStart int      `json:"ageStart"`
    AgeEnd   int      `json:"ageEnd"`
    Gender   string   `json:"gender"`
    Country  []string `json:"country"`
    Platform []string `json:"platform"`
}
type Ad struct {
    Title      string     `json:"title"`
    StartAt    string     `json:"startAt"`
    EndAt      string     `json:"endAt"`
    Conditions Conditions `json:"conditions"`
}
type QueryCondition struct {
    Offset   int    `json:"offset"`
    Limit    int    `json:"limit"`
    Age      int    `json:"age"`
    Gender   string `json:"gender"`
    Country  string `json:"country"`
    Platform string `json:"platform"`
}
