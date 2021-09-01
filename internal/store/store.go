package store

import (
	"encoding/json"
	"os"

	"github.com/ReeceRose/home-network-proxy/internal/consts"
	"github.com/ReeceRose/home-network-proxy/internal/types"
	"github.com/google/uuid"
)

// Store is an interface which provides method signatures for storing persistent information
type Store interface {
	// Core
	Get() ([]byte, error)
	Store([]byte) error

	// Custom
	GetReportingToolAgentInformation() types.ReportingToolAgent
}

var (
	_         Store = (*FileStore)(nil)
	fileStore *FileStore
)

// FileStore is a filestore implementation of a store
type FileStore struct {
}

// Instance returns the active instance of the file store
func Instance() Store {
	if fileStore != nil {
		return fileStore
	}

	fileStore = &FileStore{}
	fileStore.createFileIfNotExists(consts.REPORTING_TOOL_FILENAME)
	return fileStore
}

// Get reads a JSON file and returns the data
func (s *FileStore) Get() ([]byte, error) {
	return os.ReadFile(consts.REPORTING_TOOL_FILENAME)
}

// Store writes the desired JSON to a JSON file
func (s *FileStore) Store(data []byte) error {
	return os.WriteFile(consts.REPORTING_TOOL_FILENAME, data, 0644)
}

// createFileIfNotExists is a handy method which creates a given file if it does not exist
func (s *FileStore) createFileIfNotExists(fileName string) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

// GetReportingToolAgentInformation pulls reporting tool agent information out of the current store
func (s *FileStore) GetReportingToolAgentInformation() types.ReportingToolAgent {
	agentData, err := s.Get()
	if err != nil {
		return types.ReportingToolAgent{}
	}

	var reportingToolAgent types.ReportingToolAgent
	json.Unmarshal(agentData, &reportingToolAgent)
	if reportingToolAgent.ID.String() == "00000000-0000-0000-0000-000000000000" {
		reportingToolAgent = types.ReportingToolAgent{}
		reportingToolAgent.ID = uuid.New()
		data, _ := json.Marshal(reportingToolAgent)
		s.Store(data)
	}
	return reportingToolAgent
}
