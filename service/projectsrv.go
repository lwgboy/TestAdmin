package service

import (
	"fmt"
	"salotto/model"
	"salotto/utils"
	"salotto/utils/qjson"
)

var ProjectSrv = &projectService{}

type projectService struct {
}

func (ps *projectService) AddProject(projectInfo *model.ProjectInfo) {
	projectInfo.ProjectId = utils.GenerateUUID()
	if err := DB.Create(projectInfo); err != nil {
		fmt.Println(err)
		return
	}

}

func (ps *projectService) GetProjectList(qj *qjson.QJson) []*model.ProjectInfo {
	var ret []*model.ProjectInfo
	if err := DB.Order("created_at desc").Find(&ret).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return ret
}
