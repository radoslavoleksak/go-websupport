package websupport

import (
	"fmt"
)

type DNSService interface {
	ListAllDNSZones(userId int) (ListAllDNSZonesResponse, error)
	GetDNSZoneDetail(userId int, domainName string) (DNSZone, error)
	ListAllDNSRecords(userId int, domainName string) (ListAllDNSRecordsResponse, error)
	GetDNSRecordDetail(userId int, domainName string, recordId int) (DNSRecord, error)
	CreateDNSRecord(userId int, domainName string, dnsRecord *DNSRecord) (CreateDNSRecordResponse, error)
	UpdateDNSRecord(userId int, domainName string, recordId int, dnsRecord *DNSRecord) (CreateDNSRecordResponse, error)
	DeleteDNSRecord(userId int, domainName string, recordId int) (CreateDNSRecordResponse, error)
}

type DNSServiceImpl struct {
	client *Client
}

func (s *DNSServiceImpl) ListAllDNSZones(userId int) (ListAllDNSZonesResponse, error) {
	path := fmt.Sprintf("/v1/user/%v/zone", userId)

    req, err := s.client.newRequest("GET", path, nil)

	var responseBody ListAllDNSZonesResponse
    _, err = s.client.do(req, &responseBody)
	return responseBody, err
}

func (s *DNSServiceImpl) GetDNSZoneDetail(userId int, domainName string) (DNSZone, error) {
	path := fmt.Sprintf("/v1/user/%v/zone/%v", userId, domainName)

	req, err := s.client.newRequest("GET", path, nil)

	var dnsZone DNSZone
	_, err = s.client.do(req, &dnsZone)
	return dnsZone, err
}

func (s *DNSServiceImpl) ListAllDNSRecords(userId int, domainName string) (ListAllDNSRecordsResponse, error) {
	path := fmt.Sprintf("/v1/user/%v/zone/%v/record", userId, domainName)

	req, err := s.client.newRequest("GET", path, nil)

	var responseBody ListAllDNSRecordsResponse
	_, err = s.client.do(req, &responseBody)
	return responseBody, err
}

func (s *DNSServiceImpl) GetDNSRecordDetail(userId int, domainName string, recordId int) (DNSRecord, error) {
	path := fmt.Sprintf("/v1/user/%v/zone/%v/record/%v", userId, domainName, recordId)

	req, err := s.client.newRequest("GET", path, nil)

	var dnsRecord DNSRecord
	_, err = s.client.do(req, &dnsRecord)
	return dnsRecord, err
}

func (s *DNSServiceImpl) CreateDNSRecord(userId int, domainName string, dnsRecord *DNSRecord) (CreateDNSRecordResponse, error) {
	path := fmt.Sprintf("/v1/user/%v/zone/%v/record", userId, domainName)

	req, err := s.client.newRequest("POST", path, dnsRecord)

	var createDNSRecordResponse CreateDNSRecordResponse
	_, err = s.client.do(req, &createDNSRecordResponse)
	return createDNSRecordResponse, err
}

func (s *DNSServiceImpl) UpdateDNSRecord(userId int, domainName string, recordId int, dnsRecord *DNSRecord) (CreateDNSRecordResponse, error) {
	path := fmt.Sprintf("/v1/user/%v/zone/%v/record/%v", userId, domainName, recordId)

	req, err := s.client.newRequest("PUT", path, dnsRecord)

	var createDNSRecordResponse CreateDNSRecordResponse
	_, err = s.client.do(req, &createDNSRecordResponse)
	return createDNSRecordResponse, err
}

func (s *DNSServiceImpl) DeleteDNSRecord(userId int, domainName string, recordId int) (CreateDNSRecordResponse, error) {
	path := fmt.Sprintf("/v1/user/%v/zone/%v/record/%v", userId, domainName, recordId)

	req, err := s.client.newRequest("DELETE", path, nil)

	var createDNSRecordResponse CreateDNSRecordResponse
	_, err = s.client.do(req, &createDNSRecordResponse)
	return createDNSRecordResponse, err
}

type ListAllDNSZonesResponse struct {
	Items		[]DNSZone	`json:"items"`
}

type DNSZone struct {
	Id   		int  		`json:"id"`
	Name 		string  	`json:"name"`
	UpdateTime	int64		`json:"updateTime"`
}

type ListAllDNSRecordsResponse struct {
	Items		[]DNSRecord	`json:"items"`
}

type DNSRecord struct {
	Id   		int  		`json:"id"`
	Type 		string  	`json:"type"`
	Name 		string  	`json:"name"`
	Content		string		`json:"content"`
	TTL   		int  		`json:"ttl"`
	Prio   		string  	`json:"prio"`
	Weight   	int  		`json:"weight"`
	Port   		int  		`json:"port"`
	Zone		DNSZone		`json:"zone"`
}

type CreateDNSRecordResponse struct {
	Status		string  				`json:"status"`
	Item		DNSRecord				`json:"item"`
	Errors		map[string]interface{}  `json:"errors"`
}