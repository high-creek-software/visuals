package plausible

type SiteTokenPair struct {
	SiteID  string
	Token   string
	Favicon string
}

func NewSiteTokenPair(siteID, token string) SiteTokenPair {
	return SiteTokenPair{SiteID: siteID, Token: token}
}
