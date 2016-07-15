package cta

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/hashicorp/errwrap"
	"io/ioutil"
	"net/http"
	"time"
)

const ARRIVAL_API_URL_TEMPLATE string = "http://lapi.transitchicago.com/api/1.0/ttarrivals.aspx?key=%v&mapid=%v"
const TIME_FORMAT string = "20060102 15:04:05"
const HYDRATE_ERR_MSG string = "Error hydrating ArrivalInfo"

type CtaProxy struct {
	url    string
	apiKey string
}

func NewCtaProxy(url string, apiKey string) *CtaProxy {
	// TODO: Validate parameters
	s := CtaProxy{url: ARRIVAL_API_URL_TEMPLATE, apiKey: apiKey}
	return &s
}

func (p *CtaProxy) GetArrivalInfo(stationId int) (*ArrivalInfo, error) {
	url := fmt.Sprintf(p.url, p.apiKey, stationId)
	response, httpErr := http.Get(url)
	if httpErr != nil {
		return nil, errwrap.Wrap(
			errors.New("Error retrieving data from CTA Arrivals API"), httpErr)
	}

	defer response.Body.Close()
	data, bodyErr := ioutil.ReadAll(response.Body)
	if bodyErr != nil {
		return nil, errwrap.Wrap(
			errors.New("Failed to read response body"), bodyErr)
	}

	root, parseErr := parse(data)
	if parseErr != nil {
		return nil, errwrap.Wrap(
			errors.New("Invalid XML response"), parseErr)
	}

	info, hydrErr := hydrate(stationId, root)
	if hydrErr != nil {
		return nil, errwrap.Wrap(
			errors.New("Unable to hydrate Station object"), hydrErr)
	}
	return info, nil
}

type stationNode struct {
	StationId int
	XMLName   struct{} `xml:"ctatt"`
	Timestamp string   `xml:"tmst"`
	ErrorCode int      `xml:"errCd"`
	ErrorMsg  string   `xml:"errNm"`
	Arrivals  []eta    `xml:"eta"`
}

type eta struct {
	StationId          int    `xml:"staId"`
	StopId             int    `xml:"stpId"`
	StationName        string `xml:"staNm"`
	StationDescription string `xml:"stpDe"`
	RouteNumber        int    `xml:"rn"`
	RouteName          string `xml:"rt"`
	DestStationId      int    `xml:"destSt"`
	DestStationName    string `xml:"destNm"`
	Direction          int    `xml:"trDr"`
}

func parse(data []byte) (*stationNode, error) {
	if data == nil {
		return nil, errors.New("Nil argument: error")
	}
	root := &stationNode{}
	err := xml.Unmarshal(data, root)
	if err != nil {
		root = nil // Discard the parse result
	}
	return root, err
}

func hydrate(stationId int, node *stationNode) (*ArrivalInfo, error) {
	if node == nil {
		return nil, errors.New("stationNode cannot be nil")
	}
	ts, tsErr := time.Parse(TIME_FORMAT, node.Timestamp)
	if tsErr != nil {
		return nil, errwrap.Wrapf(HYDRATE_ERR_MSG, tsErr)
	}
	info := &ArrivalInfo{
		StationId:   stationId,
		StationName: "Clark/Lake",
		Timestamp:   ts,
	}
	return info, nil
}
