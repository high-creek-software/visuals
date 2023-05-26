package plausible

import (
	"encoding/json"
	"strconv"

	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/slog"
)

const (
    host = "https://plausible.io/api/v1/"
    
    paramSiteId = "site_id"
    paramProperty = "property"
    paramPeriod = "period"
    
    routeBreakdown = "stats/breakdown"
    routeTimeseries = "stats/timeseries"
    routeRealtimeVisitors = "stats/realtime/visitors"
    routeAggregate = "stats/aggregate"
    
    propertyEventPage = "event:page"
    propertyVisitSource = "visit:source"
    
    PropertyPeriodSevenDay = "7d"
    PropertyPeriodThirtyDay = "30d"
    PropertyPeriodSixMo = "6mo"
    PropertyPeriodTwelveMo = "12mo"
    PropertyPeriodMonth = "month"
    PropertyPeridoDay = "day"
)

type StatsRepo interface {
    LoadTopSources(period string) ([]Source, error)
    LoadTopPages(period string) ([]Page, error)
    LoadTimeSeries(period string) ([]DaySnapshot, error)
    LoadAggregate(period string) (Aggregate, error)
    LoadRealtime() (int, error)
}

type StatsRepoResty struct {
    pair SiteTokenPair
    
    client *resty.Client
}

func NewStatsRepoResty(pair SiteTokenPair) StatsRepo {
    sr := &StatsRepoResty{pair: pair}
    
    sr.client = resty.New().SetHostURL(host).SetAuthToken(pair.Token).SetHeader("User-Agent", "io.highcreeksoftware.visuals").SetDebug(true)
    
    return sr
}

func (r *StatsRepoResty) LoadTopSources(period string) ([]Source, error) {
    
    req :=  r.client.R().EnableTrace().
        SetQueryParam(paramSiteId, r.pair.SiteID).
        SetQueryParam(paramProperty, propertyVisitSource)
        
    if period != "" {
        req = req.SetQueryParam(paramPeriod, period)
    }
    
    resp, err := req.Get(routeBreakdown)
        
    if err != nil {
        slog.Error("error requesting top sources", "error", err)
        return nil, err
    }
    
    res := &Results[[]Source]{}
    err = json.Unmarshal(resp.Body(), res)
    
    return res.Results, err
}

func (r *StatsRepoResty) LoadTopPages(period string) ([]Page, error) {
    
    req := r.client.R().EnableTrace().
        SetQueryParam(paramSiteId, r.pair.SiteID).
        SetQueryParam(paramProperty, propertyEventPage)
        
    if period != "" {
        req = req.SetQueryParam(paramPeriod, period)
    }
    
    resp, err := req.Get(routeBreakdown)
        
    if err != nil {
        slog.Error("error requestion top pages", "error", err)
        return nil, err
    }
    
    res := &Results[[]Page]{}
    err = json.Unmarshal(resp.Body(), res)
    
    return res.Results, err
}

func (r *StatsRepoResty) LoadTimeSeries(period string) ([]DaySnapshot, error) {
    
    req := r.client.R().EnableTrace().
        SetQueryParam(paramSiteId, r.pair.SiteID)
        
    if period != "" {
        req = req.SetQueryParam(paramPeriod, period)
    }
    
    resp, err := req.Get(routeTimeseries)
    
    if err != nil {
        slog.Error("error requesting time series", "error", err)
        return nil, err
    }
    
    res := &Results[[]DaySnapshot]{}
    err = json.Unmarshal(resp.Body(), res)
    
    return res.Results, err
}

func (r *StatsRepoResty) LoadAggregate(period string) (Aggregate, error) {
    
    req := r.client.R().EnableTrace().
        SetQueryParam(paramSiteId, r.pair.SiteID).
        SetQueryParam("metrics", "visitors,pageviews,bounce_rate,visit_duration")
        
    if period != "" {
        req = req.SetQueryParam(paramPeriod, period)
    }
    
    resp, err := req.Get(routeAggregate)
        
    if err != nil {
        return Aggregate{}, err
    }
    
    res := &Results[Aggregate]{}
    err = json.Unmarshal(resp.Body(), res)
    
    return res.Results, err
}

func (r *StatsRepoResty) LoadRealtime() (int, error) {
    resp, err := r.client.R().EnableTrace().
        SetQueryParam(paramSiteId, r.pair.SiteID).
        Get(routeRealtimeVisitors)
        
    if err != nil {
        return -1, err
    }
    
    return strconv.Atoi(string(resp.Body()))
}