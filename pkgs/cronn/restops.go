package cronn

import (
	"fmt"
	"log"
	"time"

	"github.com/qobbysam/socketnotify/pkgs/locdb"
)

func (s *State) LockUpdate() {

	s.CanUpdate = false
}

func (s *State) UnlockUpdate() {
	s.CanUpdate = true
}

func (s *State) LockNewMsg() {
	s.NewMsgReceived = false
}

func (s *State) UnlockNewMsg() {
	s.NewMsgReceived = true
}

func (s *State) CanSave() bool {

	if s.CanUpdate {
		return s.CanUpdate
	}

	return false
}

func (s *State) UpdateLastReport() {

	now := time.Now()

	s.LastMsgReportTime = now
}

func (s *State) AddToBuffer(data locdb.NotificationRequest) {

	s.Bufferwait = append(s.Bufferwait, data)
}

func (s *State) CleanBuffer() {

	for _, v := range s.Bufferwait {

		err := s.DB.SaveRequest(v)

		if err != nil {
			fmt.Println("failed to save data")
			fmt.Println(err)
		}

		log.Println("save success full")

	}
	s.Bufferwait = make([]locdb.NotificationRequest, 0)
}
