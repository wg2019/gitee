package main

import "github.com/go-ini/ini"

// Load 加载配置.
func Load() error {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		return err
	}
	if err = cfg.Section("accountInfo").MapTo(&Account); err != nil {
		return err
	}
	return nil
}
