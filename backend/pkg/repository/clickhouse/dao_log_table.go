package clickhouse

import (
	"context"

	"github.com/CloudDetail/apo/backend/pkg/model/request"
	"github.com/CloudDetail/apo/backend/pkg/repository/clickhouse/factory"
)

func (ch *chRepo) CreateLogTable(params *request.LogTableRequest) ([]string, error) {
	sqls := factory.GetCreateTableSQL(params)
	for _, sql := range sqls {
		err := ch.conn.Exec(context.Background(), sql)
		if err != nil {
			return nil, err
		}
	}
	return sqls, nil
}

func (ch *chRepo) DropLogTable(params *request.LogTableRequest) ([]string, error) {
	sqls := factory.GetDropTableSQL(params)
	for _, sql := range sqls {
		err := ch.conn.Exec(context.Background(), sql)
		if err != nil {
			return nil, err
		}
	}
	return sqls, nil
}

func (ch *chRepo) UpdateLogTable(params *request.LogTableRequest, new, old []request.Field) ([]string, error) {
	sqls := factory.GetUpdateTableSQLByFields(params, new, old)
	for _, sql := range sqls {
		err := ch.conn.Exec(context.Background(), sql)
		if err != nil {
			return nil, err
		}
	}
	return sqls, nil
}