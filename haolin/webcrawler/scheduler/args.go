package scheduler

import "goproject/haolin/webcrawler/module"

// Args 代表参数容器的接口类型。
type Args interface {
	// Check 用于自检参数的有效性。
	// 若结果值为nil，则说明未发现问题，否则就意味着自检未通过。
	check() error
}

// RequestArgs 代表请求相关的参数容器的类型。
type RequestArgs struct {
	// AcceptedDomains 代表可以接受的URL的主域名的列表。
	// URL主域名不在列表中的请求都会被忽略，
	AcceptedDomains []string `json:"accepted_primary_domains"`
	// maxDepth 代表了需要被爬取的最大深度。
	// 实际深度大于此值的请求都会被忽略。
	maxDepth uint32 `json:"max_depth"`
}

// DataArgs 代表数据相关的参数容器的类型。
type DataArgs struct {
	// ReqBufferCap 代表请求缓冲器的容量。
	ReqBufferCap uint32 `json:"req_buffer_cap"`
	// ReqMaxBufferNumber 代表请求缓冲器的最大数量。
	ReqMaxBufferNumber uint32 `json:"req_max_buffer_number"`
	// RespBufferCap 代表响应缓冲器的容量。
	RespBufferCap uint32 `json:"resp_buffer_cap"`
	// RespMaxBufferNumber 代表响应缓冲器的最大数量。
	RespMaxBufferNumber uint32 `json:"resp_max_buffer_number"`
	// ItemBufferCap 代表条目缓冲器的容量。
	ItemBufferCap uint32 `json:"item_buffer_cap"`
	// ItemMaxBufferNumber 代表条目缓冲器的最大数量。
	ItemMaxBufferNumber uint32 `json:"item_max_buffer_number"`
	// ErrorBufferCap 代表错误缓冲器的容量。
	ErrorBufferCap uint32 `json:"error_buffer_cap"`
	// ErrorMaxBufferNumber 代表错误缓冲器的最大数量。
	ErrorMaxBufferNumber uint32 `json:"error_max_buffer_number"`
}

// ModuleArgs 代表组件相关的参数容器的类型。
type ModuleArgs struct {
	// Downloaders 代表下载器列表。
	Downloaders []module.Downloader
	// Analyzers 代表分析器列表。
	Analyzers []module.Analyzer
	// Pipelines 代表条目处理管道管道列表。
	Pipelines []module.Pipeline
}

// ModuleArgsSummary 代表组件相关的参数容器的摘要类型。
type ModuleArgsSummary struct {
	DownloaderListSize int `json:"downloader_list_size"`
	AnalyzerListSize   int `json:"analyzer_list_size"`
	PipelineListSize   int `json:"pipeline_list_size"`
}

func (args *RequestArgs) Check() error {
	if args.AcceptedDomains == nil {
		return genError("nil accepted primary domain list")
	}
	return nil
}
