package ui

import "github.com/high-creek-software/visuals/internal/plausible"

type SiteTokenRepo interface {
    List() ([]plausible.SiteTokenPair, error)
    Store(pair plausible.SiteTokenPair) error
    Delete(pair plausible.SiteTokenPair) error
}