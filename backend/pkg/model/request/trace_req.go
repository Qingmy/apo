package request

type GetTracePageListRequest struct {
	StartTime   int64  `json:"startTime" binding:"min=0"`                    // 查询开始时间
	EndTime     int64  `json:"endTime" binding:"required,gtfield=StartTime"` // 查询结束时间
	Service     string `json:"service"`                                      // 查询服务名
	EndPoint    string `json:"endpoint"`                                     // 查询Endpoint
	Instance    string `json:"instance"`                                     // 实例名
	NodeName    string `json:"nodeName"`                                     // 主机名
	ContainerId string `json:"containerId"`                                  // 容器名
	Pid         uint32 `json:"pid"`                                          // 进程号
	TraceId     string `json:"traceId"`                                      // TraceId
	PageNum     int    `json:"pageNum"`                                      // 第几页
	PageSize    int    `json:"pageSize"`

	Filters []*SpanTraceFilter `json:"filters"` // 过滤器
}

type GetOnOffCPURequest struct {
	PID       uint32 `form:"pid" binding:"required"`
	NodeName  string `form:"nodeName" binding:"required"`
	StartTime int64  `form:"startTime" binding:"required"`
	EndTime   int64  `form:"endTime" binding:"required"`
}

type GetSingleTraceInfoRequest struct {
	TraceID string `form:"traceId" binding:"required"`
}
