package db

import (
	"fmt"
	"strings"
	"github.com/jinzhu/gorm"
	"reflect"
)

// Define callbacks for updating
func init() {
	gorm.DefaultCallback.Update().Replace("gorm:assign_updating_attributes", assignUpdatingAttributesCallback)
	gorm.DefaultCallback.Update().Replace("gorm:begin_transaction", beginTransactionCallback)
	gorm.DefaultCallback.Update().Replace("gorm:before_update", beforeUpdateCallback)
	gorm.DefaultCallback.Update().Replace("gorm:save_before_associations", saveBeforeAssociationsCallback)
	gorm.DefaultCallback.Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	gorm.DefaultCallback.Update().Replace("gorm:update", updateCallback)
	gorm.DefaultCallback.Update().Replace("gorm:save_after_associations", saveAfterAssociationsCallback)
	gorm.DefaultCallback.Update().Replace("gorm:after_update", afterUpdateCallback)
	gorm.DefaultCallback.Update().Replace("gorm:commit_or_rollback_transaction", commitOrRollbackTransactionCallback)
}

// assignUpdatingAttributesCallback assign updating attributes to model
func assignUpdatingAttributesCallback(scope *gorm.Scope) {
	////logCodeLine() 
	if attrs, ok := scope.InstanceGet("gorm:update_interface"); ok {
		if updateMaps, hasUpdate := updatedAttrsWithValues(scope, attrs); hasUpdate {
			scope.InstanceSet("gorm:update_attrs", updateMaps)
		} else {
			scope.SkipLeft()
		}
	}
}

// beforeUpdateCallback will invoke `BeforeSave`, `BeforeUpdate` method before updating
func beforeUpdateCallback(scope *gorm.Scope) {
	//logCodeLine()
	if _, ok := scope.Get("gorm:update_column"); !ok {
		if !scope.HasError() {
			scope.CallMethod("BeforeSave")
		}
		if !scope.HasError() {
			scope.CallMethod("BeforeUpdate")
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdatedAt` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	//logCodeLine()
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", NowFunc())
	}
}

// updateCallback the callback used to update data to database
func updateCallback(scope *gorm.Scope) {
	//logCodeLine()
	if !scope.HasError() {
		var sqls []string

		if updateAttrs, ok := scope.InstanceGet("gorm:update_attrs"); ok {
			for column, value := range updateAttrs.(map[string]interface{}) {
				sqls = append(sqls, fmt.Sprintf("%v = %v", scope.Quote(column), scope.AddToVars(value)))
			}
		} else {
			for _, field := range scope.Fields() {
				if changeableField(scope, field) {
					if !field.IsPrimaryKey && field.IsNormal {
						sqls = append(sqls, fmt.Sprintf("%v = %v", scope.Quote(field.DBName), scope.AddToVars(field.Field.Interface())))
					} else if relationship := field.Relationship; relationship != nil && relationship.Kind == "belongs_to" {
						for _, foreignKey := range relationship.ForeignDBNames {
							if foreignField, ok := scope.FieldByName(foreignKey); ok && !changeableField(scope, foreignField) {
								sqls = append(sqls,
									fmt.Sprintf("%v = %v", scope.Quote(foreignField.DBName), scope.AddToVars(foreignField.Field.Interface())))
							}
						}
					}
				}
			}
		}

		var extraOption string
		if str, ok := scope.Get("gorm:update_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		if len(sqls) > 0 {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v%v%v",
				scope.QuotedTableName(),
				strings.Join(sqls, ", "),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// afterUpdateCallback will invoke `AfterUpdate`, `AfterSave` method after updating
func afterUpdateCallback(scope *gorm.Scope) {
	//logCodeLine()
	if _, ok := scope.Get("gorm:update_column"); !ok {
		if !scope.HasError() {
			scope.CallMethod("AfterUpdate")
		}
		if !scope.HasError() {
			scope.CallMethod("AfterSave")
		}
	}
}

func saveBeforeAssociationsCallback(scope *gorm.Scope) {
	//logCodeLine()
	if !shouldSaveAssociations(scope) {
		return
	}
	for _, field := range scope.Fields() {
		if changeableField(scope, field) && !field.IsBlank && !field.IsIgnored {
			if relationship := field.Relationship; relationship != nil && relationship.Kind == "belongs_to" {
				fieldValue := field.Field.Addr().Interface()
				scope.Err(scope.NewDB().Save(fieldValue).Error)
				if len(relationship.ForeignFieldNames) != 0 {
					// set value's foreign key
					for idx, fieldName := range relationship.ForeignFieldNames {
						associationForeignName := relationship.AssociationForeignDBNames[idx]
						if foreignField, ok := scope.New(fieldValue).FieldByName(associationForeignName); ok {
							scope.Err(scope.SetColumn(fieldName, foreignField.Field.Interface()))
						}
					}
				}
			}
		}
	}
}

func saveAfterAssociationsCallback(scope *gorm.Scope) {
	//logCodeLine()
	if !shouldSaveAssociations(scope) {
		return
	}
	for _, field := range scope.Fields() {
		if changeableField(scope, field) && !field.IsBlank && !field.IsIgnored {
			if relationship := field.Relationship; relationship != nil &&
					(relationship.Kind == "has_one" || relationship.Kind == "has_many" || relationship.Kind == "many_to_many") {
				value := field.Field

				switch value.Kind() {
				case reflect.Slice:
					for i := 0; i < value.Len(); i++ {
						newDB := scope.NewDB()
						elem := value.Index(i).Addr().Interface()
						newScope := newDB.NewScope(elem)

						if relationship.JoinTableHandler == nil && len(relationship.ForeignFieldNames) != 0 {
							for idx, fieldName := range relationship.ForeignFieldNames {
								associationForeignName := relationship.AssociationForeignDBNames[idx]
								if f, ok := scope.FieldByName(associationForeignName); ok {
									scope.Err(newScope.SetColumn(fieldName, f.Field.Interface()))
								}
							}
						}

						if relationship.PolymorphicType != "" {
							scope.Err(newScope.SetColumn(relationship.PolymorphicType, scope.TableName()))
						}

						scope.Err(newDB.Save(elem).Error)

						if joinTableHandler := relationship.JoinTableHandler; joinTableHandler != nil {
							scope.Err(joinTableHandler.Add(joinTableHandler, newDB, scope.Value, newScope.Value))
						}
					}
				default:
					elem := value.Addr().Interface()
					newScope := scope.New(elem)
					if len(relationship.ForeignFieldNames) != 0 {
						for idx, fieldName := range relationship.ForeignFieldNames {
							associationForeignName := relationship.AssociationForeignDBNames[idx]
							if f, ok := scope.FieldByName(associationForeignName); ok {
								scope.Err(newScope.SetColumn(fieldName, f.Field.Interface()))
							}
						}
					}

					if relationship.PolymorphicType != "" {
						scope.Err(newScope.SetColumn(relationship.PolymorphicType, scope.TableName()))
					}
					scope.Err(scope.NewDB().Save(elem).Error)
				}
			}
		}
	}
}

func updatedAttrsWithValues(scope *gorm.Scope, value interface{}) (results map[string]interface{}, hasUpdate bool) {
	//logCodeLine()
	if scope.IndirectValue().Kind() != reflect.Struct {
		return convertInterfaceToMap(value, false), true
	}

	results = map[string]interface{}{}

	for key, value := range convertInterfaceToMap(value, true) {
		if field, ok := scope.FieldByName(key); ok && changeableField(scope, field) {
			if _, ok := value.(*expr); ok {
				hasUpdate = true
				results[field.DBName] = value
			} else {
				err := field.Set(value)
				if field.IsNormal {
					hasUpdate = true
					if err == gorm.ErrUnaddressable {
						results[field.DBName] = value
					} else {
						results[field.DBName] = field.Field.Interface()
					}
				}
			}
		}
	}
	return
}

func convertInterfaceToMap(values interface{}, withIgnoredField bool) map[string]interface{} {
	var attrs = map[string]interface{}{}

	switch value := values.(type) {
	case map[string]interface{}:
		return value
	case []interface{}:
		for _, v := range value {
			for key, value := range convertInterfaceToMap(v, withIgnoredField) {
				attrs[key] = value
			}
		}
	case interface{}:
		reflectValue := reflect.ValueOf(values)

		switch reflectValue.Kind() {
		case reflect.Map:
			for _, key := range reflectValue.MapKeys() {
				attrs[ToDBName(key.Interface().(string))] = reflectValue.MapIndex(key).Interface()
			}
		default:
			for _, field := range (&gorm.Scope{Value: values}).Fields() {
				if !field.IsBlank && (withIgnoredField || !field.IsIgnored) {
					attrs[field.DBName] = field.Field.Interface()
				}
			}
		}
	}
	return attrs
}