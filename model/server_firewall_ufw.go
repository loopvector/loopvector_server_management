package model

type ServerFirewallUfw struct {
	ID                           uint64                `gorm:"primaryKey;autoIncrement"`
	ServerID                     uint64                `gorm:"not null;index:uk_server_group_idx_server_id_name,unique;"`
	ServerFirewallUfwRuleName    string                `gorm:"not_null;type:VARCHAR(255)"`
	ServerFirewallUfwRule        ServerFirewallUfwRule `gorm:"not_null;references:Name;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	Ip                           *string
	Port                         string                    `gorm:"not_null"`
	ServerConnectionProtocolName *string                   `gorm:"type:VARCHAR(255)"`
	ServerIpActiveState          *ServerConnectionProtocol `gorm:"references:Name;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (ServerFirewallUfw) Initialize() {
	DB.AutoMigrate(&ServerFirewallUfw{})
}

func (s ServerFirewallUfw) Update() error {
	if err := DB.FirstOrCreate(&s, &ServerFirewallUfw{
		ServerID:                  s.ServerID,
		ServerFirewallUfwRuleName: s.ServerFirewallUfwRuleName,
	}).Assign(&s).Error; err != nil {
		return err
	}
	return nil
}
