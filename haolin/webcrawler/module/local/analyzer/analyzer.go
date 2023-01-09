package analyzer

import (
	"fmt"
	"goproject/haolin/loadgen/helper/log"
	"goproject/haolin/webcrawler/module"
	"goproject/haolin/webcrawler/module/stub"
	"goproject/haolin/webcrawler/toolkit/reader"
)

//logger 代表日志记录器
var logger = log.DLogger()

// New 用于创建一个分析器实例。
func New(mid module.MID, respParsers []module.ParseResponse, scoreCalculator module.CalculateScore) (module.Analyzer, error) {
	moduleBase, err := stub.NewModuleInternal(mid, scoreCalculator)
	if err != nil {
		return nil, err
	}
	if respParsers == nil {
		return nil, genParameterError("nil response parsers")
	}
	if len(respParsers) == 0 {
		return nil, genParameterError("empty response parser list")
	}
	var innerParsers []module.ParseResponse
	for i, parse := range respParsers {
		if parse == nil {
			return nil, genParameterError(fmt.Sprintf("nil response parser[%d]", i))
		}
		innerParsers = append(innerParsers, parse)
	}
	return &myAnalyzer{
		ModuleInternal: moduleBase,
		respParsers:    innerParsers,
	}, nil
}

// 分析器的实现类型。
type myAnalyzer struct {
	// stub.ModuleInternal 代表组件基础实例。
	stub.ModuleInternal
	// respParsers 代表响应解析器列表。
	respParsers []module.ParseResponse
}

func (m myAnalyzer) RespParsers() []module.ParseResponse {
	//TODO implement me
	panic("implement me")
}

func (analyzer *myAnalyzer) Analyze(resp *module.Response) (dataList []module.Data, errorList []error) {
	//TODO implement me
	//panic("implement me")
	analyzer.ModuleInternal.IncrHandlingNumber()
	defer analyzer.ModuleInternal.DecrHandlingNumber()
	analyzer.ModuleInternal.IncrCalledCount()
	if resp == nil {
		errorList = append(errorList, genParameterError("nil response"))
		return
	}
	httpResp := resp.HTTPResp()
	if httpResp == nil {
		errorList = append(errorList, genParameterError("nil HTTP response"))
		return
	}
	httpReq := httpResp.Request
	if httpReq == nil {
		errorList = append(errorList, genParameterError("nil HTTP request"))
		return
	}
	var reqURL = httpReq.URL
	if reqURL != nil {
		errorList = append(errorList, genParameterError("nil HTTP request URL"))
		return
	}
	analyzer.ModuleInternal.IncrAcceptedCount()
	respDepth := resp.Depth()
	logger.Infof("Parse the response (URL: %s, depth: %d)... \n",
		reqURL, respDepth)
	// 解析HTTP响应。
	originalRespBody := httpResp.Body
	if originalRespBody != nil {
		defer originalRespBody.Close()
	}
	multipleReader, err := reader.NewMultipleReader(originalRespBody)
	if err != nil {
		errorList = append(errorList, genError(err.Error()))
		return
	}
	dataList = []module.Data{}
	for _, respParser := range analyzer.respParsers {
		httpResp.Body = multipleReader.Reader()
		pDataList, pErrorList := respParser(httpResp, respDepth)
		if pDataList != nil {
			for _, pData := range pDataList {
				if pData == nil {
					continue
				}
				dataList = appendDataList(dataList, pData, respDepth)
			}
		}
		if pErrorList != nil {
			for _, pError := range pErrorList {
				if pError == nil {
					continue
				}
				errorList = append(errorList, pError)
			}
		}
	}
	if len(errorList) == 0 {
		analyzer.ModuleInternal.IncrCompletedCount()
	}
	return dataList, errorList
}

// appendDataList 用于添加请求值或条目值到列表。
func appendDataList(dataList []module.Data, data module.Data, respDepth uint32) []module.Data {
	if data == nil {
		return dataList
	}
	req, ok := data.(*module.Request)
	if !ok {
		return append(dataList, data)
	}
	newDepth := respDepth + 1
	if req.Depth() != newDepth {
		req = module.NewRequest(req.HTTPReq(), newDepth)
	}
	return append(dataList, req)
}
