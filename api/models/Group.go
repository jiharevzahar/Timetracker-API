package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Group struct {
	ID    uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Title string `gorm:"size:255;not null;unique" json:"title"`
	Tasks []Task `json:"task,omitempty"`
}

func (u *Group) SaveGroup(db *gorm.DB) (*Group, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Group{}, err
	}
	return u, nil
}

func (p *Group) FindAllGroups(db *gorm.DB) (*[]Group, error) {
	var groups []Group
	rowsGroup,_ := db.Raw("SELECT id, title from groups").Rows()
	defer rowsGroup.Close()

	for rowsGroup.Next(){
		var group Group
		rowsGroup.Scan(&group.ID,&group.Title)

		group.Tasks = FindTaskByGroupID(db,group.ID)
		groups = append(groups,group)


	}
	return &groups, nil
}

func (u *Group) FindGroupByID(db *gorm.DB, uid uint32) (*Group, error) {
	var err error
	err = db.Debug().Model(Group{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Group{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Group{}, errors.New("Group Not Found")
	}
	return u, err
}

func (u *Group) UpdateAGroup(db *gorm.DB, uid uint32) (*Group, error) {
	db = db.Debug().Model(&Group{}).Where("id = ?", uid).Take(&Group{}).UpdateColumns(
		map[string]interface{}{
			"title":  u.Title,
			//"to_time": time.Now(),
		},
	)
	if db.Error != nil {
		return &Group{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&Group{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Group{}, err
	}
	return u, nil
}

func (u *Group) DeleteAGroup(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Group{}).Where("id = ?", uid).Take(&Group{}).Delete(&Group{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
