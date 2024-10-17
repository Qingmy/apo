package log

import (
	"github.com/CloudDetail/apo/backend/pkg/model/request"
	"github.com/CloudDetail/apo/backend/pkg/model/response"
	"github.com/CloudDetail/apo/backend/pkg/repository/database"
	"github.com/CloudDetail/apo/backend/pkg/services/log/vector"
	"gopkg.in/yaml.v3"
)

func (s *service) DeleteLogParseRule(req *request.DeleteLogParseRequest) (*response.LogParseResponse, error) {
	logReq := &request.LogTableRequest{
		TableName: req.TableName,
		DataBase:  req.DataBase,
	}
	logReq.FillerValue()

	// 更新k8s configmap
	res := &response.LogParseResponse{
		ParseName: req.ParseName,
	}
	data, err := s.k8sApi.GetVectorConfigFile()
	if err != nil {
		return nil, err
	}
	var vectorCfg vector.VectorConfig
	err = yaml.Unmarshal([]byte(data["aggregator.yaml"]), &vectorCfg)
	if err != nil {
		return nil, err
	}
	p := vector.ParseInfo{
		ParseName: req.ParseName,
		TableName: req.TableName,
	}
	newData, err := p.DeleteParseRule(vectorCfg)
	if err != nil {
		return nil, err
	}
	err = s.k8sApi.UpdateVectorConfigFile(newData)
	if err != nil {
		return nil, err
	}
	_, err = s.chRepo.DropLogTable(logReq)
	if err != nil {
		return nil, err
	}

	// 更新sqlite表信息
	log := database.LogTableInfo{
		ParseName: req.ParseName,
		Table:     req.TableName,
		DataBase:  logReq.DataBase,
		Cluster:   logReq.Cluster,
	}
	err = s.dbRepo.OperateLogTableInfo(&log, database.DELETE)
	if err != nil {
		return nil, err
	}

	return res, nil
}