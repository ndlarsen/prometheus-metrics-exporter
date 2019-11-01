package scrapetarget

// scrapeTarget related
type ErrorScrapeTargetUnmarshal struct {
	Err string
}

func (e ErrorScrapeTargetUnmarshal) Error() string {
	return e.Err
}
