package main

type Result struct {
    Sum int `json:"sum"`
}

type Conditions struct {
    AgeStart int      `json:"ageStart"`
    AgeEnd   int      `json:"ageEnd"`
    Country  []string `json:"country"`
    Platform []string `json:"platform"` 
}

type Ad struct {
    Title      string     `json:"title"`
    StartAt    string     `json:"startAt"`
    EndAt      string     `json:"endAt"`
    Conditions Conditions `json:"conditions"`
}