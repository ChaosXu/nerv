package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/chaosxu/nerv/lib/model"
)

// Define callbacks for deleting,support cascade deleting
func init() {
	gorm.DefaultCallback.Delete().Replace("gorm:begin_transaction", beginTransactionCallback)
	gorm.DefaultCallback.Delete().Replace("gorm:before_delete", beforeDeleteCallback)
	gorm.DefaultCallback.Delete().After("gorm:before_delete").Replace("chaosxu:before_delete_associations", beforeDeleteAssociationsCallback)
	gorm.DefaultCallback.Delete().Replace("gorm:delete", deleteCallback)
	gorm.DefaultCallback.Delete().After("gorm:delete").Replace("chaosxu:after_delete_associations", afterDeleteAssociationsCallback)
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
	logCodeLine()

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
	logCodeLine()
	//TBD config  gorm:delete_associations
	//if !scope.shouldDeleteAssociations() {
	//	return
	//}
	for _, field := range scope.Fields() {

		if relationship := field.Relationship; relationship != nil && relationship.Kind == "has_many" {
			//TBD:Now only support one foreign field and unit type
			foreignValue := scope.IndirectValue().FieldByName(relationship.AssociationForeignFieldNames[0]).Uint()
			//fmt.Println(foreignValue)
			sql := fmt.Sprintf("%s = ?", relationship.ForeignDBNames[0])
			//fmt.Println(sql)
			class := field.Field.Type().Elem()
			fmt.Println(class.Name())
			if err := DB.Unscoped().Delete(model.Models[class.Name()].Type, sql, foreignValue).Error; err != nil {
				scope.Err(err)
			}
		}
	}
}

func afterDeleteAssociationsCallback(scope *gorm.Scope) {
	logCodeLine()
	//if !scope.shouldSaveAssociations() {
	//	return
	//}
	//for _, field := range scope.Fields() {
	//	if scope.changeableField(field) && !field.IsBlank && !field.IsIgnored {
	//		if relationship := field.Relationship; relationship != nil &&
	//				(relationship.Kind == "has_one" || relationship.Kind == "has_many" || relationship.Kind == "many_to_many") {
	//			value := field.Field
	//
	//			switch value.Kind() {
	//			case reflect.Slice:
	//				for i := 0; i < value.Len(); i++ {
	//					newDB := scope.NewDB()
	//					elem := value.Index(i).Addr().Interface()
	//					newScope := newDB.NewScope(elem)
	//
	//					if relationship.JoinTableHandler == nil && len(relationship.ForeignFieldNames) != 0 {
	//						for idx, fieldName := range relationship.ForeignFieldNames {
	//							associationForeignName := relationship.AssociationForeignDBNames[idx]
	//							if f, ok := scope.FieldByName(associationForeignName); ok {
	//								scope.Err(newScope.SetColumn(fieldName, f.Field.Interface()))
	//							}
	//						}
	//					}
	//
	//					if relationship.PolymorphicType != "" {
	//						scope.Err(newScope.SetColumn(relationship.PolymorphicType, scope.TableName()))
	//					}
	//
	//					scope.Err(newDB.Save(elem).Error)
	//
	//					if joinTableHandler := relationship.JoinTableHandler; joinTableHandler != nil {
	//						scope.Err(joinTableHandler.Add(joinTableHandler, newDB, scope.Value, newScope.Value))
	//					}
	//				}
	//			default:
	//				elem := value.Addr().Interface()
	//				newScope := scope.New(elem)
	//				if len(relationship.ForeignFieldNames) != 0 {
	//					for idx, fieldName := range relationship.ForeignFieldNames {
	//						associationForeignName := relationship.AssociationForeignDBNames[idx]
	//						if f, ok := scope.FieldByName(associationForeignName); ok {
	//							scope.Err(newScope.SetColumn(fieldName, f.Field.Interface()))
	//						}
	//					}
	//				}
	//
	//				if relationship.PolymorphicType != "" {
	//					scope.Err(newScope.SetColumn(relationship.PolymorphicType, scope.TableName()))
	//				}
	//				scope.Err(scope.NewDB().Save(elem).Error)
	//			}
	//		}
	//	}
}
