package models

import "time"

type URL struct {
	ID               int       `db:"id" json:"id"`
	URL              string    `db:"url" json:"url"`
	Status           string    `db:"status" json:"status"`
	HTMLVersion      string    `db:"html_version" json:"html_version"`
	Title            string    `db:"title" json:"title"`
	H1Count          int       `db:"h1_count" json:"h1_count"`
	H2Count          int       `db:"h2_count" json:"h2_count"`
	H3Count          int       `db:"h3_count" json:"h3_count"`
	InternalLinks    int       `db:"internal_links" json:"internal_links"`
	ExternalLinks    int       `db:"external_links" json:"external_links"`
	BrokenLinksCount int       `db:"broken_links_count" json:"broken_links_count"`
	HasLoginForm     bool      `db:"has_login_form" json:"has_login_form"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
}
