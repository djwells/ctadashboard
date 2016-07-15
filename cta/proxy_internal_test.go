package cta

import (
	"github.com/djwells/ctadashboard/cta/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var ARRIVAL_RESPONSE_XML_VALID []byte = testdata.MustAsset("arrivals_api_valid_response.xml")
var ARRIVAL_RESPONSE_XML_INVALID []byte = []byte(`<?xml version="1.0" encoding="utf-8"?><ctattbad><tmst>20160711 20:49:57</tmst><errCd>0</errCd><errNm /><eta><staId>40380</staId><stpId>30075</stpId><staNm>Clark/Lake</staNm><stpDe>Service at Outer Loop platform</stpDe><rn>621</rn><rt>G</rt><destSt>30004</destSt><destNm>Harlem/Lake</destNm><trDr>1</trDr><prdt>20160711 20:49:38</prdt><arrT>20160711 20:50:38</arrT><isApp>1</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.88574</lat><lon>-87.62706</lon><heading>270</heading></eta><eta><staId>40380</staId><stpId>30075</stpId><staNm>Clark/Lake</staNm><stpDe>Service at Outer Loop platform</stpDe><rn>430</rn><rt>Brn</rt><destSt>30249</destSt><destNm>Kimball</destNm><trDr>1</trDr><prdt>20160711 20:49:34</prdt><arrT>20160711 20:51:34</arrT><isApp>0</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.88261</lat><lon>-87.62617</lon><heading>358</heading></eta><eta><staId>40380</staId><stpId>30074</stpId><staNm>Clark/Lake</staNm><stpDe>Service at Inner Loop platform</stpDe><rn>612</rn><rt>G</rt><destSt>30057</destSt><destNm>Ashland/63rd</destNm><trDr>5</trDr><prdt>20160711 20:49:37</prdt><arrT>20160711 20:50:37</arrT><isApp>1</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.88574</lat><lon>-87.63181</lon><heading>89</heading></eta><eta><staId>40380</staId><stpId>30375</stpId><staNm>Clark/Lake</staNm><stpDe>Subway service toward O'Hare</stpDe><rn>224</rn><rt>Blue</rt><destSt>30171</destSt><destNm>O'Hare</destNm><trDr>1</trDr><prdt>20160711 20:49:13</prdt><arrT>20160711 20:52:13</arrT><isApp>0</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.87818</lat><lon>-87.6293</lon><heading>358</heading></eta><eta><staId>40380</staId><stpId>30075</stpId><staNm>Clark/Lake</staNm><stpDe>Service at Outer Loop platform</stpDe><rn>424</rn><rt>Brn</rt><destSt>30249</destSt><destNm>Kimball</destNm><trDr>1</trDr><prdt>20160711 20:49:26</prdt><arrT>20160711 20:56:26</arrT><isApp>0</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.87685</lat><lon>-87.63236</lon><heading>89</heading></eta><eta><staId>40380</staId><stpId>30074</stpId><staNm>Clark/Lake</staNm><stpDe>Service at Inner Loop platform</stpDe><rn>313</rn><rt>Pink</rt><destSt>30114</destSt><destNm>54th/Cermak</destNm><trDr>5</trDr><prdt>20160711 20:49:15</prdt><arrT>20160711 20:56:15</arrT><isApp>0</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.88574</lat><lon>-87.63089</lon><heading>89</heading></eta><eta><staId>40380</staId><stpId>30074</stpId><staNm>Clark/Lake</staNm><stpDe>Service at Inner Loop platform</stpDe><rn>726</rn><rt>Org</rt><destSt>30182</destSt><destNm>Midway</destNm><trDr>5</trDr><prdt>20160711 20:49:39</prdt><arrT>20160711 20:58:39</arrT><isApp>0</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.86683</lat><lon>-87.62658</lon><heading>359</heading></eta><eta><staId>40380</staId><stpId>30375</stpId><staNm>Clark/Lake</staNm><stpDe>Subway service toward O'Hare</stpDe><rn>137</rn><rt>Blue</rt><destSt>30171</destSt><destNm>O'Hare</destNm><trDr>1</trDr><prdt>20160711 20:49:32</prdt><arrT>20160711 21:00:32</arrT><isApp>0</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.87573</lat><lon>-87.66925</lon><heading>88</heading></eta><eta><staId>40380</staId><stpId>30374</stpId><staNm>Clark/Lake</staNm><stpDe>Subway service toward Forest Park</stpDe><rn>140</rn><rt>Blue</rt><destSt>30077</destSt><destNm>Forest Park</destNm><trDr>5</trDr><prdt>20160711 20:49:27</prdt><arrT>20160711 21:01:27</arrT><isApp>0</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.9224</lat><lon>-87.69762</lon><heading>121</heading></eta><eta><staId>40380</staId><stpId>30074</stpId><staNm>Clark/Lake</staNm><stpDe>Service at Inner Loop platform</stpDe><rn>015</rn><rt>G</rt><destSt>30139</destSt><destNm>Cottage Grove</destNm><trDr>5</trDr><prdt>20160711 20:49:17</prdt><arrT>20160711 21:02:17</arrT><isApp>0</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.8849</lat><lon>-87.71652</lon><heading>93</heading></eta><eta><staId>40380</staId><stpId>30075</stpId><staNm>Clark/Lake</staNm><stpDe>Service at Outer Loop platform</stpDe><rn>429</rn><rt>Brn</rt><destSt>30249</destSt><destNm>Kimball</destNm><trDr>1</trDr><prdt>20160711 20:49:27</prdt><arrT>20160711 21:03:27</arrT><isApp>0</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.90129</lat><lon>-87.6367</lon><heading>182</heading></eta><eta><staId>40380</staId><stpId>30374</stpId><staNm>Clark/Lake</staNm><stpDe>Subway service toward Forest Park</stpDe><rn>218</rn><rt>Blue</rt><destSt>30077</destSt><destNm>Forest Park</destNm><trDr>5</trDr><prdt>20160711 20:49:16</prdt><arrT>20160711 21:08:16</arrT><isApp>0</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.9526</lat><lon>-87.72857</lon><heading>119</heading></eta><eta><staId>40380</staId><stpId>30074</stpId><staNm>Clark/Lake</staNm><stpDe>Service at Inner Loop platform</stpDe><rn>317</rn><rt>Pink</rt><destSt>30114</destSt><destNm>54th/Cermak</destNm><trDr>5</trDr><prdt>20160711 20:49:38</prdt><arrT>20160711 21:08:38</arrT><isApp>0</isApp><isSch>0</isSch><isDly>0</isDly><isFlt>0</isFlt><flags /><lat>41.85394</lat><lon>-87.71733</lon><heading>89</heading></eta></ctattbad>`)

func Test_parse_ValidData_ResultIsCorrect(t *testing.T) {
	root, err := parse([]byte(ARRIVAL_RESPONSE_XML_VALID))
	assert.Nil(t, err, "Precondition failed")

	assert.NotNil(t, root, "root should not be nil")

	numArrivals := len(root.Arrivals)
	assert.Equal(t, 13, numArrivals, "Unexpected number of arrival nodes")

	firstArr := root.Arrivals[0]
	assert.Equal(t, 40380, firstArr.StationId)
	assert.Equal(t, 30075, firstArr.StopId)
	assert.Equal(t, "Clark/Lake", firstArr.StationName)
	assert.Equal(t, "Service at Outer Loop platform", firstArr.StationDescription)
	assert.Equal(t, 621, firstArr.RouteNumber)
	assert.Equal(t, "G", firstArr.RouteName)
	assert.Equal(t, 30004, firstArr.DestStationId)
	assert.Equal(t, "Harlem/Lake", firstArr.DestStationName)
	assert.Equal(t, 1, firstArr.Direction)
}

func Test_parse_BadData_ResultIsCorrect(t *testing.T) {
	badDataList := [][]byte{nil, ARRIVAL_RESPONSE_XML_INVALID}
	for _, badData := range badDataList {
		root, err := parse([]byte(badData))

		assert.Nil(t, root)
		assert.NotNil(t, err)
	}
}

func Test_hydrate_ResultIsCorrect(t *testing.T) {
	root, parseErr := parse([]byte(ARRIVAL_RESPONSE_XML_VALID))
	assert.Nil(t, parseErr, "Precondition failed")

	stationId := 40380
	info, hydrErr := hydrate(stationId, root)

	assert.NotNil(t, info)
	assert.Nil(t, hydrErr)
	assert.Equal(t, 40380, info.StationId)
	assert.Equal(t, "Clark/Lake", info.StationName)
	expectedTimestamp, _ := time.Parse(
		time.RFC3339, "2016-07-11T20:49:57Z")
	assert.Equal(t, expectedTimestamp, info.Timestamp)
}

func Test_hydrate_NilArg_ResultIsCorrect(t *testing.T) {
	info, hydrErr := hydrate(12345, nil)

	assert.Nil(t, info)
	assert.NotNil(t, hydrErr)
}
