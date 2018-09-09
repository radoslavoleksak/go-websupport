package websupport

import "fmt"

type UserService interface {
	ListAllUsers() (ListAllUsersResponse, error)
	GetUserDetail(userId int) (UserDetail, error)
}

type UserServiceImpl struct {
	client *Client
}

func (s *UserServiceImpl) ListAllUsers() (ListAllUsersResponse, error) {
	path := "/v1/user"

	req, err := s.client.newRequest("GET", path, nil)

	var responseBody ListAllUsersResponse
	_, err = s.client.do(req, &responseBody)
	return responseBody, err
}

func (s *UserServiceImpl) GetUserDetail(userId int) (UserDetail, error) {
	path := fmt.Sprintf("/v1/user/%v", userId)

	req, err := s.client.newRequest("GET", path, nil)

	var userDetail UserDetail
	_, err = s.client.do(req, &userDetail)
	return userDetail, err
}

type ListAllUsersResponse struct {
	Items		[]User	`json:"items"`
}

type User struct {
	Id   		int     `json:"id"`
	Login 		string  `json:"login"`
	ParentId 	int 	`json:"parentId"`
	Active 		bool    `json:"active"`
	CreateTime 	int64   `json:"createTime"`
	Group 		string  `json:"group"`
}

type Market struct {
	Name 		string  `json:"name"`
	Identifier 	string  `json:"identifier"`
	Currency 	string  `json:"currency"`
}

type Billing struct {
	Id   		 int     `json:"id"`
	Profile 	 string  `json:"profile"`
	IsDefault 	 bool    `json:"isDefault"`
	Name 		 string  `json:"name"`
	City 		 string  `json:"city"`
	Street 		 string  `json:"street"`
	CompanyRegId string  `json:"companyRegId"`
	TaxId 		 string  `json:"taxId"`
	VatId 		 string  `json:"vatId"`
	ZIP 		 string  `json:"zip"`
	Country 	 string  `json:"country"`
	ISIC 		 string  `json:"isic"`
}

type UserDetail struct {
	*User
	Email   	  string  `json:"email"`
	Phone   	  string  `json:"phone"`
	SKNICHandle   string  `json:"sknicHandle"`
	ContactPerson string  `json:"contactPerson"`
	ResellerToken string  `json:"resellerToken"`
	Credit   	  int  	  `json:"credit"`
	VerifyUrl     string  `json:"verifyUrl"`
	Billing		  []Billing `json:"billing"`
	Market		  Market  `json:"market"`
}
