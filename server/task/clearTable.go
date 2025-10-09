package task

import (
	"errors"
	"fmt"
	"time"
	"goLinker-admin/server/global"
	"goLinker-admin/server/model/common"

	"gorm.io/gorm"
)

//@function: ClearTable
//@description: 清理数据库表数据
//@param: db(数据库对象) *gorm.DB, tableName(表名) string, compareField(比较字段) string, interval(间隔) string
//@return: error

func ClearTable(db *gorm.DB) error {
	var ClearTableDetail []common.ClearDB

	ClearTableDetail = append(ClearTableDetail, common.ClearDB{
		TableName:    "sys_operation_records",
		CompareField: "created_at",
		Interval:     "2160h",
	})

	ClearTableDetail = append(ClearTableDetail, common.ClearDB{
		TableName:    "jwt_blacklists",
		CompareField: "created_at",
		Interval:     "168h",
	})

	if db == nil {
		return errors.New("db Cannot be empty")
	}

	var totalDeleted int64 = 0
	var results []string

	for _, detail := range ClearTableDetail {
		duration, err := time.ParseDuration(detail.Interval)
		if err != nil {
			return err
		}
		if duration < 0 {
			return errors.New("parse duration < 0")
		}
		result := db.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", detail.TableName, detail.CompareField), time.Now().Add(-duration))
		if result.Error != nil {
			return result.Error
		}

		deletedCount := result.RowsAffected
		totalDeleted += deletedCount
		results = append(results, fmt.Sprintf("%s: 删除 %d 条记录", detail.TableName, deletedCount))
	}

	// 更新历史记录，包含删除统计信息
	global.UpdateCronHistory(global.CronClearDbName, &global.CronHistory{
		LastResult: fmt.Sprintf("清理完毕，共删除 %d 条记录。详情：%s", totalDeleted, fmt.Sprintf("%v", results)),
	})

	return nil
}
