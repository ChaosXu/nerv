package lock

import (
	"github.com/ChaosXu/nerv/lib/db"
)

func init() {
	db.Models["Lock"] = lockDesc()
}

func lockDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Lock{},
		New: func() interface{} {
			return &Lock{}
		},
		NewSlice:func() interface{} {
			return &[]Lock{}
		},
	}
}

//Lock for mutex
type Lock struct {
	Type   string    `gorm:"unique_index:idx_lock_tol"`   //object type
	ObjID  uint        `gorm:"unique_index:idx_lock_tol"` //object id
	LockID uint        `gorm:"unique_index:idx_lock_tol"` //lock id
}

func GetLock(objType string, objID uint) *Lock {
	return &Lock{Type:objType, ObjID:objID, LockID:1}
}

//TryLock return true if the lock has been acquired
func (p *Lock) TryLock() bool {

	if err := db.DB.Create(p).Error; err != nil {
		return false
	}
	return true
}

//Unlock release the lock.Do nothing if no lock
func (p *Lock) Unlock() {
	db.DB.Where("type=? and obj_id=? and lock_id=?", p.Type, p.ObjID, p.LockID).Delete(p)
}


