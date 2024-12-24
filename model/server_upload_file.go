package model

type ServerUploadFile struct {
	ID                                    uint64 `gorm:"primaryKey;autoIncrement"`
	ServerID                              uint64 `gorm:"not null;index:uk_server_ipv4_idx_server_id_ip_subnet,unique;"`
	SourceFilePath                        string
	DestinationFilePath                   string
	DestinationFolderPermission           string
	SourceFileName                        string
	DestinationFileName                   string
	OldDestinationFileBackupDirectoryPath string
	OldDestinationFileBackupFileName      string
	ShouldOverwrite                       bool
	ShouldBackupOldDestinationFile        bool
	SourceFileContent                     *string
	OldDestinationFileContent             *string
}

func (ServerUploadFile) Initialize() {
	DB.AutoMigrate(&ServerUploadFile{})
}

func (s ServerUploadFile) AddNew() error {
	if err := DB.Create(&s).Error; err != nil {
		return err
	}
	return nil
}
