package presenters

// New Link Endpoint
type (
	NewLinkRequest struct {
		Link string `json:"link"`
	}
	NewLinkResponse struct {
		Short string `json:"short"`
		Error string `json:"error"`
	}
)

// Redirect By Link Endpoint
type (
	RedirectByLinkResponse struct {
		Error string `json:"error"`
	}
)
