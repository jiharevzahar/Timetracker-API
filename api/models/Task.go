package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type Task struct {
	ID      uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title   string `gorm:"size:255;not null;unique" json:"title"`
	GroupID uint32 `gorm:"not null" json:"group_id"`
	Timeframes []Timeframe `json:"time_frame,omitempty"`
}

func (p *Task) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
}

func (p *Task) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.GroupID < 1 {
		return errors.New("Required Group")
	}
	return nil
}

func (p *Task) SaveTask(db *gorm.DB) (*Task, error) {
	var err error
	err = db.Debug().Model(&Task{}).Create(&p).Error
	if err != nil {
		return &Task{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Group{}).Where("id = ?", p.GroupID).Error
		if err != nil {
			return &Task{}, err
		}
	}
	return p, nil
}

func (p *Task) FindAllTasks(db *gorm.DB) (*[]Task, error) {
	var tasks []Task
	rowsTask,_ := db.Raw("SELECT id, title, group_id from tasks").Rows()
	defer rowsTask.Close()

	for rowsTask.Next(){
		var task Task
		rowsTask.Scan(&task.ID,&task.Title,&task.GroupID)
		task.Timeframes = FindTimerameByTaskID(db,task.GroupID)
		tasks = append(tasks, task)
	}
	return &tasks, nil
}

func (p *Task) FindTaskByID(db *gorm.DB, pid uint64) (*Task, error) {
	var err error
	err = db.Debug().Model(&Task{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Task{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Group{}).Where("id = ?", p.GroupID).Error
		if err != nil {
			return &Task{}, err
		}
	}
	return p, nil
}



func (p *Task) UpdateATask(db *gorm.DB) (*Task, error) {

	var err error

	err = db.Debug().Model(&Task{}).Where("id = ?", p.ID).Updates(Task{Title: p.Title}).Error
	if err != nil {
		return &Task{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Group{}).Where("id = ?", p.GroupID).Error
		if err != nil {
			return &Task{}, err
		}
	}
	return p, nil
}

func (p *Task) DeleteATask(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Task{}).Where("id = ?", pid).Take(&Task{}).Delete(&Task{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Task not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func FindTaskByGroupID(db *gorm.DB, GroupID uint32)([]Task){
	var tasks []Task
	rowsTask, err := db.Raw("SELECT id, title, group_id from tasks WHERE group_id = ?", GroupID).Rows()
	if err!=nil{
		panic(err)
	}
	defer rowsTask.Close()
	for rowsTask.Next(){
		var task Task
		rowsTask.Scan(&task.ID,&task.Title,&task.GroupID)
		task.Timeframes = FindTimerameByTaskID(db,uint32(task.ID))
		tasks = append(tasks, task)
	}
	return tasks
}