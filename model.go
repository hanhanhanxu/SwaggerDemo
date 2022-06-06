package main

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	// example binding是swagger里用到的标签
	Title  string `gorm:"type:varchar(100);" json:"title" example:"title" binding:"required"`
	Des    string `gorm:"type:varchar(100);" json:"des" example:"desc" binding:"required"`
	Status string `gorm:"type:varchar(200);" json:"status" example:"Active"`
}

/*
CREATE TABLE `post`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `des` varchar(100) NOT NULL,
	`status` varchar(200),
  `create_time` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB;
*/

type Response struct {
	Msg  string
	Data interface{}
}
