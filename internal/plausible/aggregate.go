package plausible

type Aggregate struct {
    BounceRate Valuer `json:"bounce_rage"`
    PageViews Valuer `json:"pageviews"`
    VisitDuration Valuer `json:"visit_duration"`
    Visitors Valuer `json:"visitors"`
}

type Valuer struct {
    Value float64 `json:"value"`
}