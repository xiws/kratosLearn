package external

import (
	"github.com/hashicorp/consul/api"
	"log"
)

func RegisterConsul() {
	consulConfig := &api.Config{
		Address: "https://consul.test.shijizhongyun.com",
		Token:   "16bf2708-f93e-7278-b29f-c542699480c1",
	}

	client, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("Consul 客户端创建失败: %v", err)
	}

	serviceID := "workflowDemo-1" // 确保 serviceID 唯一
	registration := &api.AgentServiceRegistration{
		ID:      serviceID,             // 服务唯一ID
		Name:    "workflowDemo",        // Consul 中的服务名称
		Address: "192.168.1.100",       // 服务运行的 IP
		Port:    8080,                  // 监听的端口
		Tags:    []string{"go", "api"}, // 标签
		Check: &api.AgentServiceCheck{
			HTTP:     "http://192.168.1.100:8080/health", // 健康检查地址
			Interval: "10s",                              // 每 10 秒检查一次
			Timeout:  "5s",                               // 超时 5 秒
		},
	}

	// 3️⃣ 注册到 Consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("服务注册失败: %v", err)
	}

	log.Println("✅ workflowDemo 服务注册成功！")

}
