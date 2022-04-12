package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/project"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TargetSearch struct{
    project.Target
    request.PageInfo
    OrderKey string
    Desc bool
    ProjectId primitive.ObjectID `json:"project_id" bson:"project_id"`
}