package model

import (
	"loopvector_server_management/controller/helper"
)

type AdminConfig struct {
	SMTPHost                      string   `yaml:"smtp_host,omitempty"`
	SMTPPort                      uint16   `yaml:"smtp_port,omitempty"`
	SMTPUser                      string   `yaml:"smtp_user,omitempty"`
	SMTPPassword                  string   `yaml:"smtp_password,omitempty"`
	SignupDomainWhitelist         []string `yaml:"signup_domain_whitelist,omitempty"`
	UserEmailVerificationRequired bool     `yaml:"user_email_verification_required,omitempty"`
}

const (
	kAdminSettingFilePath = "./config/admin_settings.yaml"
)

func GenerateAdminSetting(config AdminConfig) error {
	return helper.GenerateConfig[AdminConfig](kAdminSettingFilePath, config)
}

func UpdateAdminSetting(updates AdminConfig) error {
	return helper.UpdateConfig[AdminConfig](kAdminSettingFilePath, updates)
}

func LoadAdminSetting() (AdminConfig, error) {
	return helper.LoadConfig[AdminConfig](kAdminSettingFilePath)
}
