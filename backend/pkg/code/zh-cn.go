package code

var zhCnText = map[string]string{
	ServerError:    "内部服务器错误",
	ParamBindError: "参数信息错误",
	DbConnectError: "数据库连接失败",

	MockCreateError: "创建mock失败",
	MockListError:   "获取mock列表失败",
	MockDetailError: "获取mock详情失败",
	MockDeleteError: "删除mock失败",

	GetServiceUrlRelationError:     "获取服务调用关系失败",
	GetDescendantMetricsError:      "获取依赖视图的延时曲线失败",
	GetDescendantRelevanceError:    "获取依赖视图的关联度失败",
	GetPolarisInferError:           "获取北极星指标分析情况失败",
	GetErrorInstanceError:          "获取错误实例失败",
	GetErrorInstanceLogsError:      "获取错误实例故障现场日志失败",
	GetLogMetricsError:             "获取Log关键指标失败",
	GetLogLogsError:                "获取Log故障现场日志失败",
	GetTraceMetricsError:           "获取Trace关键指标失败",
	GetTraceLogsError:              "获取Trace故障现场日志失败",
	GetServiceListError:            "获取服务名列表失败",
	GetServiceInstanceOptionsError: "获取服务实例名列表失败",
	GetServiceEntryEndpointsError:  "获取服务入口Endpoint列表失败",
	GetK8sEventError:               "无法获取k8s事件",
	GetServiceEndPointListError:    "获取服务Endpoint列表失败",
	GetServiceRYGLightError:        "获取服务的红绿灯失败",

	GetFaultLogPageListError: "获取故障现场日志分页列表失败",
	GetFaultLogContentError:  "获取故障现场日志内容失败",

	QueryLogContextError: "查询日志上下文失败",
	QueryLogError:        "查询全量日志失败",
	GetLogChartError:     "获取全量日志图表数据失败",
	GetLogIndexError:     "获取全量日志索引失败",

	GetLogTableInfoError:    "获取日志表信息失败",
	GetLogParseRuleError:    "获取日志表解析规则失败",
	UpdateLogParseRuleError: "更新日志表解析规则失败",

	GetTracePageListError:    "获取Trace分页列表失败",
	GetTraceFiltersError:     "获取Trace过滤条件失败",
	GetTraceFilterValueError: "获取Trace过滤条件值失败",
	GetOnOffCPUError:         "获取Trace中span执行消耗失败",
	GetSingleTraceError:      "获取单条trace数据失败",

	GetOverviewServiceInstanceListError: "获取实例列表失败",
	GetServiceMoreUrlListError:          "获取更多服务端点失败",
	GetThresholdError:                   "获取阈值信息失败",
	GetTop3UrlListError:                 "获取应用下前三条异常服务端点信息失败",
	SetThresholdError:                   "设置阈值信息失败",
	GetServicesAlertError:               "获取服务告警信息失败",
	SetTTLError:                         "配置存储周期失败",
	GetTTLError:                         "获取存储周期失败",
	GetMonitorStatusError:               "获取监控服务状态失败",
	SetSingleTableTTLError:              "配置单个存储周期失败",

	GetAlertEventsError:       "获取告警事件失败",
	GetAlertEventsSampleError: "获取采样告警事件失败",

	GetSQLMetricError:    "获取SQL关键指标失败",
	GetAlertRuleError:    "获取告警规则失败",
	UpdateAlertRuleError: "更新告警规则失败",
	AddAlertRuleError:    "添加告警规则失败",
	DeleteAlertRuleError: "删除告警规则失败",

	UpdateAlertRuleValidateError: "验证告警规则失败,通常为规则中的expr非法",

	GetAMConfigReceiverError:    "获取告警通知对象失败",
	AddAMConfigReceiverError:    "添加告警通知对象失败",
	UpdateAMConfigReceiverError: "更新告警通知对象失败",
	DeleteConfigReceiverError:   "删除告警通知对象失败",

	AlertGroupAndLabelMismatchError: "组名和label中的组名不匹配",
	AlertKeepFiringForIllegalError:  "'keepFiringFor' 不合法",
	AlertForIllegalError:            "'for' 不合法",
	AlertOldGroupNotExistError:      "选择的告警规则所在组不存在",
	AlertAlertNotExistError:         "选择的告警规则不存在",
	AlertAlertAlreadyExistError:     "告警规则已存在",
	AlertConfigFileNotExistError:    "告警规则配置文件不存在",
	AlertTargetGroupNotExistError:   "目标组不存在",
	AlertCheckRuleError:             "查看告警规则名是否占用失败",

	AlertManagerReceiverAlreadyExistsError:  "告警通知对象名称已存在",
	AlertManagerReceiverNotExistsError:      "告警通知对象名称不存在",
	AlertManagerReceiverEmailHostMissing:    "告警通知对象 email 'smarthost' 配置缺失",
	AlertManagerReceiverEmailFromMissing:    "告警通知对象 email 'from' 配置缺失",
	AlertManagerEmptyReceiver:               "告警通知对象没有设置任何 webhook 或 email 配置",
	AlertManagerDefaultReceiverCannotDelete: "默认告警通知对象不能被删除",

	AlertEventImpactError:            "查询告警事件影响面失败",
	AlertEventImpactMissingTag:       "查询告警事件影响面失败, 事件标签中未找到任意可关联标签组 ",
	AlertEventImpactNoMatchedService: "查询告警事件影响面失败, 未找到告警事件匹配的服务 ",
	AlertEventIDMissing:              "查询告警事件影响面失败, 搜索时间范围内未找到告警事件ID对应的事件 ",

	AlertAnalyzeDescendantAnormalEventDeltaError: "分析告警事件失败: 查询下游异常事件失败",
	GetAnomalySpanError:                          "获取故障报告失败",
	GetDetectMutationExecListError:               "获取异常检测执行记录失败",
	GetDetectMutationRuleListError:               "获取取异常检测规则失败",
	GetQuickMutationMetricError:                  "获取预定义的快速异常检测指标失败",
	GetMetricPQLError:                            "获取预定义的快速告警规则指标失败",

	MutationPQLCheckFailed: "通过PQL检查指标突变失败",

	AlertAnalyzeDescendantAnormalEventError:   "获取下游异常事件失败",
	AlertAnalyzeDescendantAnormalContribution: "获取下游异常事件贡献失败",
	DetectDefectsError:                        "异常分析失败",
	DetectDefectsCreatAlertError:              "创建异常告警失败",
	AddExecRecordError:                        "添加执行记录失败",

	K8sGetResourceError: "获取k8s资源失败",
}
