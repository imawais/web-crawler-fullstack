package crawler

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/imawais/web-crawler-fullstack/backend/database"
)

// CrawlAndStore crawls a given URL and stores results into DB
func CrawlAndStore(targetURL string, id int) error {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Get(targetURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return errors.New("page returned non-OK status")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	// ---- 1. HTML version detection ----
	htmlVersion := "HTML5"
	if strings.Contains(strings.ToLower(resp.Proto), "1.0") {
		htmlVersion = "HTML 4.01"
	}

	// ---- 2. Page title ----
	title := strings.TrimSpace(doc.Find("title").Text())

	// ---- 3. Heading tag counts ----
	hCounts := make(map[string]int)
	for _, h := range []string{"h1", "h2", "h3"} {
		hCounts[h] = doc.Find(h).Length()
	}

	// ---- 4. Internal and external links ----
	parsedBase, _ := url.Parse(targetURL)
	internalCount := 0
	externalCount := 0
	brokenLinks := 0
	loginFormFound := false

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if href == "" || strings.HasPrefix(href, "javascript:") {
			return
		}
		linkURL, err := url.Parse(href)
		if err != nil {
			return
		}
		fullURL := parsedBase.ResolveReference(linkURL).String()

		if parsedBase.Host == linkURL.Host || linkURL.Host == "" {
			internalCount++
		} else {
			externalCount++
		}

		// Check link status
		go func(link string) {
			if statusCode := checkLink(link); statusCode >= 400 {
				_, _ = database.DB.Exec("INSERT INTO broken_links (url_id, link, status_code) VALUES (?, ?, ?)", id, link, statusCode)
			}
		}(fullURL)
	})

	// ---- 5. Login form detection ----
	doc.Find("form").Each(func(i int, form *goquery.Selection) {
		if form.Find("input[type='password']").Length() > 0 {
			loginFormFound = true
		}
	})

	// ---- 6. Save to DB ----
	_, err = database.DB.Exec(`
		UPDATE urls SET
			html_version = ?, title = ?,
			h1_count = ?, h2_count = ?, h3_count = ?,
			internal_links = ?, external_links = ?,
			has_login_form = ?, status = 'done',
			updated_at = NOW()
		WHERE id = ?
	`, htmlVersion, title, hCounts["h1"], hCounts["h2"], hCounts["h3"],
		internalCount, externalCount, loginFormFound, id)

	return err
}

// Check link status code
func checkLink(link string) int {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Head(link)
	if err != nil {
		return 500
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
