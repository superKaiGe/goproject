package module

import (
	"fmt"
	"goproject/haolin/webcrawler/errors"
	"sync"
)

// Registrar 代表组件注册器的接口。
type Registrar interface {
	// Register 用于注册组件实例。
	Register(module Module) (bool, error)
	// Unregister 用于注销组件实例。
	Unregister(mid MID) (bool, error)
	// Get 用于获取一个指定类型的组件的实例。
	// 本函数应该基于负载均衡策略返回实例。
	Get(moduleType Type) (Module, error)
	// GetAllByType 用于获取指定类型的所有组件实例。
	GetAllByType(moduleType Type) (map[Type]Module, error)
	// GetAll 用于获取所有组件实例。
	GetAll() map[MID]Module
	// Clear 会清除所有的组件注册记录。
	Clear()
}

// NewRegistrar 用于创建一个组件注册器的实例。
func NewRegistrar() Registrar {
	return &myRegistrar{
		moduleTypeMap: map[Type]map[MID]Module{},
		//rwlock:        sync.RWMutex{},
	}
}

// myRegistrar 代表组件注册器的实现类型。
type myRegistrar struct {
	// moduleTypeMap 代表组件类型与对应组件实例的映射。
	moduleTypeMap map[Type]map[MID]Module
	// rwlock 代表组件注册专用读写锁。
	rwlock sync.RWMutex
}

func (registrar *myRegistrar) Register(module Module) (bool, error) {
	//TODO implement me
	//panic("implement me")
	if module == nil {
		return false, errors.NewIllegalParameterError("nil module instance")
	}
	mid := module.ID()
	parts, err := SplitMID(mid)
	if err != nil {
		return false, err
	}
	moduleType := legalLetterTypeMap[parts[0]]
	if !CheckType(moduleType, module) {
		errMsg := fmt.Sprintf("incorrect module type: %s", moduleType)
		return false, errors.NewIllegalParameterError(errMsg)
	}
	registrar.rwlock.Lock()
	defer registrar.rwlock.Unlock()
	modules := registrar.moduleTypeMap[moduleType]
	if modules == nil {
		modules = map[MID]Module{}
	}
	if _, ok := modules[mid]; ok {
		return false, nil
	}
	modules[mid] = module
	registrar.moduleTypeMap[moduleType] = modules
	return true, nil
}

func (m myRegistrar) Unregister(mid MID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (registrar *myRegistrar) Get(moduleType Type) (Module, error) {
	//TODO implement me
	//panic("implement me")
	modules, err := registrar.GetAllByType(moduleType)
	if err != nil {
		return nil, err
	}
	minScore := uint64(0)
	var selectedModule Module
	for _, module := range modules {
		SetScore(module)
		score := module.Score()
		if minScore == 0 || score < minScore {
			selectedModule = module
			minScore = score
		}
	}
	return selectedModule, nil
}

func (registrar *myRegistrar) GetAllByType(moduleType Type) (map[Type]Module, error) {
	//TODO implement me
	//panic("implement me")
	if !legalSN(moduleType) {
		errMsg := fmt.Sprintf("illegal module type: %s", moduleType)
		return nil, errors.NewIllegalParameterError(errMsg)
	}
	registrar.rwlock.RLock()
	defer registrar.rwlock.RUnlock()
	modules := registrar.moduleTypeMap[moduleType]
	if len(modules) == 0 {
		return nil, ErrNotFoundModuleInstance
	}
	result := map[MID]Module{}
	for mid, module := range modules {
		result[mid] = module
	}
	return result, nil
}

func (registrar *myRegistrar) GetAll() map[MID]Module {
	//TODO implement me
	//panic("implement me")
	result := map[MID]Module{}
	registrar.rwlock.RLock()
	defer registrar.rwlock.RUnlock()
	for _, modules := range registrar.moduleTypeMap {
		for mid, module := range modules {
			result[mid] = module
		}
	}
	return result
}

func (registrar *myRegistrar) Clear() {
	//TODO implement me
	//panic("implement me")
	registrar.rwlock.Lock()
	defer registrar.rwlock.Unlock()
	registrar.moduleTypeMap = map[Type]map[MID]Module{}
}
