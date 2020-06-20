package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jiharevzahar/fullstack/api/models"
)

var groups = []models.Group{
	models.Group{
		Title:    "final task EPAM",
	},
	models.Group{
		Title:    "oop BSUIR",
	},
}

var tasks = []models.Task{
	models.Task{
		Title:   "researches",
	},
	models.Task{
		Title:   "lab7",
	},
}

var timeframes = []models.Timeframe{
	models.Timeframe{
		TaskID: 1,
	},
	models.Timeframe{
		TaskID: 1,
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Timeframe{},&models.Task{}, &models.Group{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Group{}, &models.Task{}, &models.Timeframe{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Task{}).AddForeignKey("group_id", "groups(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Timeframe{}).AddForeignKey("task_id", "tasks(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range groups {
		err = db.Debug().Model(&models.Group{}).Create(&groups[i]).Error
		if err != nil {
			log.Fatalf("cannot seed groups table: %v", err)
		}
		tasks[i].GroupID = groups[i].ID

		err = db.Debug().Model(&models.Task{}).Create(&tasks[i]).Error
		if err != nil {
			log.Fatalf("cannot seed tasks table: %v", err)
		}

		err = db.Debug().Model(&models.Timeframe{}).Create(&timeframes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed timeframes table: %v", err)
		}
	}
}