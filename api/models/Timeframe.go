package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Timeframe struct {
	ID       uint64    `gorm:"primary_key;auto_increment" json:"id"`
	TaskID   uint32    `gorm:"not null" json:"task_id"`
	FromTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"from"`
	ToTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"to"`
}

func (p *Timeframe) Prepare() {
	p.ID = 0
	p.FromTime = time.Now()
	p.ToTime = time.Now()
}

func (p *Timeframe) Validate() error {
	if p.TaskID < 1 {
		return errors.New("Required Timeframe")
	}
	return nil
}

func (p *Timeframe) SaveTimeframe(db *gorm.DB) (*Timeframe, error) {
	var err error
	err = db.Debug().Model(&Timeframe{}).Create(&p).Error
	if err != nil {
		return &Timeframe{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Timeframe{}).Where("id = ?", p.TaskID).Error
		if err != nil {
			return &Timeframe{}, err
		}
	}
	return p, nil
}

func (p *Timeframe) FindAllTimeframes(db *gorm.DB) (*[]Timeframe, error) {
	var err error
	tasks := []Timeframe{}
	err = db.Debug().Model(&Timeframe{}).Limit(100).Find(&tasks).Error
	if err != nil {
		return &[]Timeframe{}, err
	}
	if len(tasks) > 0 {
		for i, _ := range tasks {
			err := db.Debug().Model(&Timeframe{}).Where("id = ?", tasks[i].TaskID).Error
			if err != nil {
				return &[]Timeframe{}, err
			}
		}
	}
	return &tasks, nil
}

func (p *Timeframe) FindTimeframeByID(db *gorm.DB, pid uint64) (*Timeframe, error) {
	var err error
	err = db.Debug().Model(&Timeframe{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Timeframe{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Timeframe{}).Where("id = ?", p.TaskID).Error
		if err != nil {
			return &Timeframe{}, err
		}
	}
	return p, nil
}

func (p *Timeframe) DeleteATimeframe(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Timeframe{}).Where("id = ?", pid).Take(&Timeframe{}).Delete(&Timeframe{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Timeframe not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func FindTimerameByTaskID(db *gorm.DB, TaskID uint32)([]Timeframe){
	var timeframes []Timeframe
	rowsTimeframe, err := db.Raw("SELECT id, task_id, from_time, to_time from timeframes WHERE task_id = ?", TaskID ).Rows()
	if err!=nil{
		panic(err)
	}
	defer rowsTimeframe.Close()
	for rowsTimeframe.Next(){
		var timeframe Timeframe
		rowsTimeframe.Scan(&timeframe.ID,&timeframe.TaskID,&timeframe.FromTime,&timeframe.ToTime)

		timeframes = append(timeframes, timeframe)
	}
	return timeframes
}