package loadgen

import (
	"bytes"
	"errors"
	"fmt"
	"goproject/haolin/loadgen/lib"
	"strings"
	"time"
)

type ParamSet struct {
	Caller     lib.Caller           //调用器
	TimeoutNS  time.Duration        //响应超时时间 单位 纳秒
	LPS        uint32               //每秒载荷量
	DurationNS time.Duration        //负载持续时间 单位纳秒
	ResultCh   chan *lib.CallResult //调用结果通道
}

func (pset *ParamSet) Check() error {
	var errMsgs []string
	if pset.Caller == nil {
		errMsgs = append(errMsgs, "Invalid caller")
	}
	if pset.TimeoutNS == 0 {
		errMsgs = append(errMsgs, "Invalid timeoutNS")
	}
	if pset.LPS == 0 {
		errMsgs = append(errMsgs, "Invalid lps(load per second)!")
	}
	if pset.DurationNS == 0 {
		errMsgs = append(errMsgs, "Invalid durationNS!")
	}
	if pset.ResultCh == nil {
		errMsgs = append(errMsgs, "Invalid result channel!")
	}
	var buf bytes.Buffer
	buf.WriteString("Checking the parameters...")
	if errMsgs != nil {
		errMsg := strings.Join(errMsgs, " ")
		buf.WriteString(fmt.Sprintf("NOT passed!(%s)", errMsgs))
		logger.Infoln(buf.String())
		return errors.New(errMsg)
	}
	buf.WriteString(fmt.Sprintf("Passed. (timeoutNS=%s, lps=%d, durationNS=%s)",
		pset.TimeoutNS, pset.LPS, pset.DurationNS))
	logger.Infoln(buf.String())
	return nil
}
