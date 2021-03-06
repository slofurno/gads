package v201705

// https://developers.google.com/adwords/api/docs/reference/v201705/AdGroupExtensionSettingService.Page
// Contains the results from a get call.
type Page struct {
	TotalNumEntries int    `xml:"https://adwords.google.com/api/adwords/cm/v201705 totalNumEntries,omitempty"`
	PageType        string `xml:"https://adwords.google.com/api/adwords/cm/v201705 Page.Type,omitempty"`
}
