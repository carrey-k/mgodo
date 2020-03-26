package mgodo

import (
	"time"

	//"gopkg.in/mgo.v2/bson"
	"github.com/globalsign/mgo/bson"
)

const (
	UPDATE = "UPDATE"
	DELETE = "DELETE" // soft delete
	ERASE  = "ERASE"  // hard delete
)

// BaseModel to be emmbered to other struct as audit trail perpurse
type BaseModel struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	CreatedAt time.Time     `bson:"CreatedAt,omitempty"`
	CreatedBy string        `bson:"CreatedBy,omitempty"`
	UpdatedAt time.Time     `bson:"UpdatedAt,omitempty"`
	UpdatedBy string        `bson:"UpdatedBy,omitempty"`
	IsRemoved bool          `bson:"IsRemoved,omitempty"`
	RemovedAt time.Time     `bson:"RemovedAt,omitempty"`
	RemovedBy string        `bson:"RemovedBy,omitempty"`
	IsLocked  bool          `bson:"IsLocked,omitempty"`
}

//ChangeLog
type ChangeLog struct {
	BaseModel    `bson:",inline"`
	ModelObjId   bson.ObjectId `bson:"ModelObjId,omitempty"`
	ModelName    string        `bson:"ModelName,omitempty"`
	ModelValue   interface{}   `bson:"ModelValue,omitempty"`
	Operation    string        `bson:"Operation,omitempty"`
	ChangeReason string        `bson:"ChangeReason,omitempty"`
}
