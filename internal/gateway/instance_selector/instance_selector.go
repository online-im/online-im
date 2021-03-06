package instance_selector

import (
	"fmt"
	"github.com/glory-go/glory/common"
	"github.com/glory-go/glory/config"
	"github.com/glory-go/glory/log"
	"github.com/glory-go/glory/plugin"
	"github.com/online-im/online-im/internal/constant"
	_ "github.com/online-im/online-im/internal/redis_client"
	perrors "github.com/pkg/errors"
	"go.uber.org/atomic"
	"sync"
)

var instanceSelector *InstanceSelector

type InstanceSelector struct {
	localCache map[string]bool
	counter    *atomic.Uint32
	length     uint32
	mapLock    *sync.RWMutex
}

func init() {
	instanceSelector = &InstanceSelector{
		localCache: make(map[string]bool),
		mapLock:    &sync.RWMutex{},
		counter:    atomic.NewUint32(0),
	}
	go instanceSelector.Watch()
}

func GetInstanceSelector() *InstanceSelector {
	return instanceSelector
}

// Watch is run in single gr
func (i *InstanceSelector) Watch() {
	reg := plugin.GetRegistry(config.GlobalServerConf.RegistryConfig[constant.RegistryKey])
	eventCh, err := reg.Subscribe(constant.IMInstanceServiceID)
	if err != nil {
		panic(err)
	}

	for v := range eventCh {
		addr := v.Addr.GetUrl()
		i.mapLock.Lock()
		switch v.Opt {
		case common.RegistryAddEvent, common.RegistryUpdateEvent:
			if _, ok := i.localCache[addr]; ok {
				continue
			}
			i.localCache[addr] = true
			i.length++
			log.Info("InstanceSelector start working")
			fmt.Println("InstanceSelector start working")
		case common.RegistryDeleteEvent:
			delete(i.localCache, addr)
			i.length--
		}
		i.mapLock.Unlock()
	}
}

// todo  select use round robin is not good!
// Select is called multi gr
func (i *InstanceSelector) Select() (string, error) {
	i.mapLock.RLock()
	idx := i.counter.Load() % i.length
	count := uint32(0)
	for k, _ := range i.localCache {
		if count == idx {
			return k, nil
		}
		count++
	}
	i.mapLock.RUnlock()
	return "", perrors.Errorf("no online-im instance existed")
}
