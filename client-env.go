package date_agent

import (
	"k8s.io/klog/v2"
	"os"
	"time"
)

const DateAgentRegisterService = "DATE_AGENT_REGISTER_SERVICE"

func NewClientByEnv() {
	if svc := os.Getenv(DateAgentRegisterService); svc != "" {
		for {
			client, err := NewClient(svc)
			if err != nil {
				time.Sleep(time.Second * 10)
				klog.V(5).Infof("agent.NewClient err:%v", err)
				continue
			}
			select {
			case <-client.DoneSignal():
				klog.V(6).Info("client.DoneSignal shutdown due to error")
			}
		}
	} else {
		klog.Infof("%s was not set", DateAgentRegisterService)
		return
	}
}
