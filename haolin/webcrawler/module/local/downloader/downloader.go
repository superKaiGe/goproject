package downloader

import (
	"goproject/haolin/loadgen/helper/log"
	"goproject/haolin/webcrawler/module"
	"goproject/haolin/webcrawler/module/stub"
	"net/http"
)

//logger 代表日志记录器
var logger = log.DLogger()

// New 用于创建一个下载器实例。
func New(mid module.MID, client *http.Client, scoreCalculator module.CalculateScore) (module.Downloader, error) {
	moduleBase, err := stub.NewModuleInternal(mid, scoreCalculator)
	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, genParameterError("nil http client")
	}
	return &myDownloader{
		ModuleInternal: moduleBase,
		httpClient:     *client,
	}, nil
}

// myDownloader 代表下载器的实现类型。
type myDownloader struct {
	// stub.ModuleInternal 代表组件基础实例。 匿名字段
	stub.ModuleInternal
	// httpClient 代表下载用的HTTP客户端。
	httpClient http.Client
}

func (downloader *myDownloader) Download(req *module.Request) (*module.Response, error) {
	//TODO implement me
	//panic("implement me")
	downloader.ModuleInternal.IncrHandlingNumber()
	defer downloader.ModuleInternal.IncrHandlingNumber()
	downloader.ModuleInternal.IncrCalledCount()
	if req == nil {
		return nil, genParameterError("nil request")
	}
	httReq := req.HTTPReq()
	if httReq == nil {
		return nil, genParameterError("nil HTTP request")
	}
	downloader.ModuleInternal.IncrAcceptedCount()
	logger.Infof("Do the request (URL: %s, depth: %d)... \n", httReq.URL, req.Depth())
	httpResp, err := downloader.httpClient.Do(httReq)
	if err != nil {
		return nil, err
	}
	downloader.ModuleInternal.IncrCompletedCount()
	return module.NewResponse(httpResp, req.Depth()), nil
}
