package db

import (
	"gorm.io/gorm"
	"time"
)

type FlightInfo struct {
	gorm.Model
	FlightName     string    `json:"flightname,omitempty"`
	FlightStartTime     time.Time `json:"flightstarttime"`
	FlightEndTime time.Time `json:"flightendtime"`
	TotalTicketNum int       `json:"totalticketnum,omitempty"`
	DiscountNum    int       `json:"discountnum,omitempty"`
	DiscountSize   float64   `json:"discountsize,omitempty"`
	TicketPrice    float64   `json:"ticketprice,omitempty"`
	HadDisNum int `json:"haddisnum,omitempty"`
}

func (f *FlightInfo) AddFlight()  {
	DB.Create(f)
}


func DelFlight(ID string)  {
	var f FlightInfo
	DB.Model(&FlightInfo{}).Where("ID=?", ID).Find(&f)
	DB.Delete(&f)
}

func (f *FlightInfo) UpdateFlightInfo()  {
	DB.Updates(f)
}

func GetAllFlight() ([]*FlightInfo, int64) {
	var flights []*FlightInfo
	tx := DB.Find(&flights)
	if tx.Error != nil{
		return nil, 0
	}
	return flights, tx.RowsAffected
}

type TicketInfo struct {
	gorm.Model
	FlightFid int `json:"flightfid"`
	UserIDCard string `json:"useridcard,omitempty"`
}

func (t *TicketInfo) AddTicket()  {
	DB.Create(t)
	var f FlightInfo
	DB.Model(&FlightInfo{}).Where("ID=?", t.FlightFid).First(&f)
	f.HadDisNum += 1
	f.UpdateFlightInfo()
}

func DelTicket(ID string)  {
	var t TicketInfo
	DB.Model(&TicketInfo{}).Where("ID=?", ID).First(&t)
	DB.Delete(&t)
	var f FlightInfo
	DB.Model(&FlightInfo{}).Where("ID=?", t.FlightFid).First(&f)
	f.HadDisNum -= 1
	f.UpdateFlightInfo()
}

func GetTicketByUid(uid string) ([]*TicketInfo, int64) {
	var tickets []*TicketInfo
	tx := DB.Model(&TicketInfo{}).Where("user_id_card=?", uid).Find(&tickets)
	if tx.Error != nil{
		return nil, 0
	}
	return tickets, tx.RowsAffected
}

type Order struct {
	gorm.Model
	Tid        uint    `json:"tid,omitempty"`
	UserName   string  `json:"username,omitempty"`
	FlightName string  `json:"flightname,omitempty"`
	Price      float64 `json:"price,omitempty"`
}

func GetNameByIDCard(ID string) string {
	var u User
	if err:= DB.Model(&User{}).Where("id_card=?", ID).Find(&u).Error; err!=nil{
		return ""
	}
	return u.Username
}

func GetFLightByFid(ID int) *FlightInfo {
	var f FlightInfo
	if err:= DB.Model(&FlightInfo{}).Where("id=?", ID).Find(&f).Error; err!=nil{
		return nil
	}
	return &f
}

func GetTickets() ([]*Order, int64) {
	var ts []*TicketInfo
	var ol []*Order
	tx:= DB.Model(&TicketInfo{}).Find(&ts)
	if tx.Error!=nil{
		return ol, 0
	}
	for _, t := range ts {
		flight := GetFLightByFid(t.FlightFid)
		ol = append(ol, &Order{
			Model:      t.Model,
			Tid:        t.ID,
			UserName:   GetNameByIDCard(t.UserIDCard),
			FlightName: flight.FlightName,
			Price:      flight.TicketPrice,
		})
	}
	return ol, tx.RowsAffected
}
