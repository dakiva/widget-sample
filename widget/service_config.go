package widget

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/emicklei/go-restful"
	"github.com/op/go-logging"
)

var ApiVersion string = "1.0.0"
var ServiceVersion string
var ServiceName string = "Sample service"
var Container *restful.Container
var logger = logging.MustGetLogger("sample")

const (
	SCHEMA_VERSION = 1
)

type ServiceConfig struct {
	Hostname string   `json:"hostname"`
	Port     int      `json:"port"`
	WidgetDB DBConfig `json:"widget_db"`
}

// Returns the host address host:port. If the host is empty, returns a leading ':'.
func (this *ServiceConfig) GetHostAddress() string {
	return fmt.Sprintf("%v:%d", this.Hostname, this.Port)
}

func (this *ServiceConfig) Validate() error {
	if this.Port <= 0 || this.Port > 65535 {
		return errors.New("Port value must a specified valid number between 0 and 65535.")
	}
	err := this.WidgetDB.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (this *ServiceConfig) Initialize() error {
	err := this.WidgetDB.Initialize(SCHEMA_VERSION)
	if err != nil {
		return err
	}

	Container = newContainer()
	return nil
}

func newContainer() *restful.Container {
	restful.DefaultRequestContentType(restful.MIME_JSON)
	container := restful.NewContainer()
	new(WidgetController).register(container)
	return container
}

// Loads a configuration from a file name into the structure specified as input, returning an error if an error occurs
func LoadServiceConfig(fileName string) (*ServiceConfig, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configData := &ServiceConfig{}
	if err := decoder.Decode(configData); err != nil {
		return nil, err
	}
	return configData, nil
}
