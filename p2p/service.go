package p2p

import (
	_ "fmt"
	"github.com/linkchain/common/util/log"
)

type Service struct {
}

func (s *Service) Init(i interface{}) bool {
	log.Info("p2p service init...")
	return true
}

func (s *Service) Start() bool {
	log.Info("p2p service start...")
	return true
}

func (s *Service) Stop() {
	log.Info("p2p service stop...")
}
