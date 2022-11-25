package dpfm_api_input_reader

import (
	"data-platform-api-business-user-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBusinessUser() *requests.BusinessUser {
	data := sdc.BusinessUser
	return &requests.BusinessUser{
		BusinessPartner:       data.BusinessPartner,
		BusinessUser:          data.BusinessUser,
		ValidityEndDate:       data.ValidityEndDate,
		ValidityStartDate:     data.ValidityStartDate,
		BusinessUserFirstName: data.BusinessUserFirstName,
		BusinessUserLastName:  data.BusinessUserLastName,
		BusinessUserFullName:  data.BusinessUserFullName,
		Language:              data.Language,
	}
}
