package requests

type BusinessUser struct {
	BusinessPartner       *int    `json:"BusinessPartner"`
	BusinessUser          *int    `json:"BusinessUser"`
	ValidityEndDate       *string `json:"ValidityEndDate"`
	ValidityStartDate     *string `json:"ValidityStartDate"`
	BusinessUserFirstName string  `json:"BusinessUserFirstName"`
	BusinessUserLastName  string  `json:"BusinessUserLastName"`
	BusinessUserFullName  string  `json:"BusinessUserFullName"`
	Language              string  `json:"Language"`
}
