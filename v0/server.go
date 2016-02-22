package colonycore

import (
	"github.com/eaciit/orm/v1"
)

type Server struct {
	orm.ModelBase
	ID       string `json:"_id", bson:"_id"`
	OS       string `json:"os", bson:"os"`
	AppPath  string `json:"appPath",bson:"appPath"`
	DataPath string `json:"dataPath",bson:"dataPath"`
	Host     string `json:"host", bson:"host"`
	SSHType  string `json:"sshtype", bson:"sshtype"`
	SSHFile  string `json:"sshfile", bson:"sshfile"`
	SSHUser  string `json:"sshuser", bson:"sshuser"`
	SSHPass  string `json:"sshpass", bson:"sshpass"`
}

func (s *Server) TableName() string {
	return "servers"
}

func (s *Server) RecordID() interface{} {
	return s.ID
}
