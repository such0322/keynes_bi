package models

import (
	"fmt"
	"keynesTe/libs"
)

type Te struct {
	Id         int
	Title      string
	CreateTime libs.XTime `json:"create_time"`
}

func (m *Te) LoadTeById(id int) {

	err := db.Debug().First(&m, id).Error
	if err != nil {
		fmt.Println(err)
	}

}

func (m *Te) Save() {
	var err error
	if m.Id == 0 {
		err = db.Debug().Create(m).Error
	} else {
		err = db.Debug().Update(m).Error
	}
	if err != nil {
		fmt.Println(err)
	}
}
