package db

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	//"github.com/ChaosXu/nerv/lib/log"
)

// Define callbacks for deleting,support cascade deleting
func init() {
	gorm.DefaultCallback.Delete().Replace("gorm:begin_transaction", beginTransactionCallback)
	gorm.DefaultCallback.Delete().Replace("gorm:before_delete", beforeDeleteCallback)
	gorm.DefaultCallback.Delete().After("gorm:before_delete").Replace("ChaosXu:before_delete_associations", beforeDeleteAssociationsCallback)
	gorm.DefaultCallback.Delete().Replace("gorm:delete", deleteCallback)
	gorm.DefaultCallback.Delete().After("gorm:delete").Replace("ChaosXu:after_delete_associations", afterDeleteAssociationsCallback)
	gorm.DefaultCallback.Delete().Replace("gorm:after_delete", afterDeleteCallback)
	gorm.DefaultCallback.Delete().Replace("gorm:commit_or_rollback_transaction", commitOrRollbackTransactionCallback)
}

func beginTransactionCallback(scope *gorm.Scope) {
	scope.Begin()
}

// beforeDeleteCallback will invoke `BeforeDelete` method before deleting
func beforeDeleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		scope.CallMethod("BeforeDelete")
	}
}

// deleteCallback used to delete data from database or set deleted_at to current time (when using with soft delete)
func deleteCallback(scope *gorm.Scope) {
	//log.LogCodeLine()

	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		if !scope.Search.Unscoped && scope.HasColumn("DeletedAt") {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET deleted_at=%v%v%v",
				scope.QuotedTableName(),
				scope.AddToVars(gorm.NowFunc()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// afterDeleteCallback will invoke `AfterDelete` method after deleting
func afterDeleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		scope.CallMethod("AfterDelete")
	}
}

func commitOrRollbackTransactionCallback(scope *gorm.Scope) {
	scope.CommitOrRollback()
}

//cascade deleting
func beforeDeleteAssociationsCallback(scope *gorm.Scope) {
	//log.LogCodeLine()
	//TBD config  gorm:delete_associations
	//if !scope.shouldDeleteAssociations() {
	//	return
	//}
	for _, field := range scope.Fields() {
		if relationship := field.Relationship; relationship != nil && relationship.Kind == "has_many" {
			//TBD:Now only support one foreign field and unit type
			foreignValue := scope.IndirectValue().FieldByName(relationship.AssociationForeignFieldNames[0]).Uint()
			sql := fmt.Sprintf("%s = ?", relationship.ForeignDBNames[0])
			fieldType := field.Field.Type()
			elem := fieldType.Elem()
			class := ""
			if elem.Kind() == reflect.Ptr {
				class = elem.Elem().Name()
			} else {
				class = elem.Name()
			}
			if err := DB.Unscoped().Delete(Models[class].Type, sql, foreignValue).Error; err != nil {
				//log.LogCodeLine()
				scope.Err(err)
			}
		}
	}
}

func afterDeleteAssociationsCallback(scope *gorm.Scope) {
	//log.LogCodeLine()
}
