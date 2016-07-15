package ctadashboard

import (
//"https://github.com/bluele/gcache"
)

type StationInfo struct {
}

type CtaApiService struct {
	stop    chan bool
	stopAck chan bool
}

func (s *CtaApiService) Serve() {
	s.stop = make(chan bool)
	s.stopAck = make(chan bool)
Loop:
	for {
		select {

		case <-s.stop:
			break Loop
		}
	}
	s.stopAck <- true
}

func (s *CtaApiService) GetStationInfo(stationId int) {

}
