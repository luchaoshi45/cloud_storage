package handler

import (
	"cloud_storage/db/mysql"
	"cloud_storage/service/account/proto"
	"context"
	"encoding/json"
)

type UserFile struct {
}

// UserFiles : 查询批量的文件元信息
func (uf *User) UserFiles(ctx context.Context, req *proto.ReqUserFile, resp *proto.RespUserFile) error {
	limitCnt := int(req.Limit)
	username := req.Username
	//fileMetas, _ := meta.GetLastFileMetasDB(limitCnt)
	userFile := mysql.NewUserFile()
	userFiles, err := userFile.QueryUserFileMetas(username, limitCnt)
	//fmt.Println(userFiles)
	if err != nil {
		resp.Code = -1
		resp.Message = "userFile.QueryUserFileMetas(username, limitCnt)"
		return nil
	}

	jdata, err := json.Marshal(userFiles)
	if err != nil {
		resp.Code = -2
		resp.Message = "json.Marshal(userFiles)"
		return nil
	}
	resp.FileData = jdata
	return nil
}

// UserFileRename : 用户文件重命名
func (uf *User) UserFileRename(ctx context.Context, req *proto.ReqUserFileRename, res *proto.RespUserFileRename) error {
	//TODO: 实现 UserFileRename
	res.FileData = nil
	return nil
}
