// Package config 存放程序所有的配置信息
package config

import "gohub/pkg/config"

func init() {
	config.Add("chataikey", func() map[string]interface{} {
		return map[string]interface{}{

			// 智谱AI配置
			"chatglmkey": config.Env("CHAT_GLM_KEY", ""),
		}
	})
}
