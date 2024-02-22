package m

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// Define the Policy structure
type Policy struct {
	SR SR
	AO AO
	SP SP  //1 is allow , 0 is deney
	AE AE
}

type SR struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	Group  string `json:"group"`
}

type AO struct {
	DeviceId string `json:"deviceId"`
	MAC string `json:"MAC"`
}

type SP struct {
	Permissions []string `json:"permissions"`
}

type AE struct {
	CreatedTime int64  `json:"createdTime"`
	EndTime     int64  `json:"endTime"`
	AllowedIP   string `json:"allowedIP"`
}

func (p *Policy) ToBytes() []byte {
	b, err := json.Marshal(*p)
	if err != nil {
		return nil
	}
	return b
}

func (p *Policy) GetID() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(p.SR.UserId+p.AO.DeviceId)))
}
