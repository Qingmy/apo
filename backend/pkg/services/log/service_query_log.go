package log

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/CloudDetail/apo/backend/pkg/model/request"
	"github.com/CloudDetail/apo/backend/pkg/model/response"
	"github.com/CloudDetail/apo/backend/pkg/repository/database"
)

func (s *service) QueryLog(req *request.LogQueryRequest) (*response.LogQueryResponse, error) {
	logs, sql, err := s.chRepo.QueryAllLogs(req)
	res := &response.LogQueryResponse{Query: sql}
	if err != nil {
		res.Err = err.Error()
		return res, nil
	}

	rows, err := s.chRepo.OtherLogTableInfo(&request.OtherTableInfoRequest{
		DataBase:  req.DataBase,
		TableName: req.TableName,
	})
	if err != nil {
		res.Err = err.Error()
		return res, nil
	}
	allFileds := []string{}
	for _, row := range rows {
		allFileds = append(allFileds, row["name"].(string))
	}
	res.DefaultFields = allFileds

	hiddenFields := []string{}
	model := &database.LogTableInfo{
		DataBase: req.DataBase,
		Table:    req.TableName,
	}
	s.dbRepo.OperateLogTableInfo(model, database.QUERY)
	var fields []request.Field
	_ = json.Unmarshal([]byte(model.Fields), &fields)
	// if err != nil {
	// 	return nil, err
	// }

	for _, field := range fields {
		hiddenFields = append(hiddenFields, field.Name)
	}

	hMap := make(map[string]struct{})
	for _, item := range hiddenFields {
		hMap[item] = struct{}{}
	}

	var defaultFields []string
	for _, item := range allFileds {
		if _, exists := hMap[item]; !exists {
			if item == req.TimeField || item == req.LogField {
				continue
			}
			defaultFields = append(defaultFields, item)
		}
	}
	res.Limited = req.PageSize
	res.HiddenFields = hiddenFields
	res.DefaultFields = defaultFields

	if len(logs) == 0 {
		res.Err = "未查询到任何日志数据"
		return res, nil
	}

	var timestamp int64
	logitems := make([]response.LogItem, len(logs))
	for i, log := range logs {
		content := log[req.LogField]
		delete(log, req.LogField)

		for k, v := range log {
			if k == req.TimeField {
				ts, ok := v.(time.Time)
				if ok {
					timestamp = ts.UnixMicro()
				} else {
					return nil, errors.New("timestamp type error")
				}
				delete(log, k)
			}
			vMap, ok := v.(map[string]string)
			if ok {
				for k2, v2 := range vMap {
					log[k+"."+k2] = v2
				}
				delete(log, k)
			}
		}

		logitems[i] = response.LogItem{
			Content: content,
			Tags:    log,
			Time:    timestamp,
		}
	}

	res.Logs = logitems
	res.Query = sql
	return res, nil
}
