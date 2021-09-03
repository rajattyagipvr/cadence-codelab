package common

import (
	"fmt"
	"io/ioutil"

	"go.uber.org/cadence/internal"
	//"go.uber.org/cadence/.gen/go/cadence"
	s "go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/cadence/common"
	"go.uber.org/zap"
	"github.com/uber-go/tally"
	"gopkg.in/yaml.v2"
)

const (
	configFile = "config/development.yaml"
)

type (
	// .
	Runtime struct {
		Service m.TChanWorkflowService
		Scope   tally.Scope
		Logger  *zap.Logger
		Config  Configuration
		Builder *WorkflowClientBuilder
	}

	// Configuration for running samples.
	Configuration struct {
		DomainName      string `yaml:"domain"`
		ServiceName     string `yaml:"service"`
		HostNameAndPort string `yaml:"host"`
	}
)

var domainCreated bool

func NewRuntime() *Runtime {
	c := &Runtime{}
	c.doInit()
	return c
}

// SetupServiceConfig setup the config for the sample code run
func (h *Runtime) doInit() {
	if h.Service != nil {
		return
	}

	// Initialize developer config for running samples
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to log config file: %v, Error: %v", configFile, err))
	}

	if err := yaml.Unmarshal(configData, &h.Config); err != nil {
		panic(fmt.Sprintf("Error initializing configuration: %v", err))
	}

	// Initialize logger for running samples
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.Info("Logger created.")
	h.Logger = logger
	h.Scope = tally.NoopScope
	h.Builder = NewBuilder().
		SetHostPort(h.Config.HostNameAndPort).
		SetDomain(h.Config.DomainName).
		SetMetricsScope(h.Scope)
	service, err := h.Builder.BuildServiceClient()
	if err != nil {
		panic(err)
	}
	h.Service = service

	if domainCreated {
		return
	}
	domainClient, _ := h.Builder.BuildCadenceDomainClient()
	request := &s.RegisterDomainRequest{
		Name:                                   common.StringPtr(h.Config.DomainName),
		Description:                            common.StringPtr("domain for cadence sample code"),
		WorkflowExecutionRetentionPeriodInDays: common.Int32Ptr(3)}
	err = domainClient.Register(request)
	if err != nil {
		if _, ok := err.(*s.DomainAlreadyExistsError); !ok {
			panic(err)
		}
		logger.Info("Domain already registered.", zap.String("Domain", h.Config.DomainName))
	} else {
		logger.Info("Domain succeesfully registered.", zap.String("Domain", h.Config.DomainName))
	}
	domainCreated = true
}

// StartWorkflow starts a workflow
func (h *Runtime) StartWorkflow(options internal.StartWorkflowOptions, workflow interface{}, args ...interface{}) {
	workflowClient, err := h.Builder.BuildCadenceClient()
	if err != nil {
		h.Logger.Error("Failed to build cadence client.", zap.Error(err))
		panic(err)
	}

	we, err := workflowClient.StartWorkflow(options, workflow, args...)
	if err != nil {
		h.Logger.Error("Failed to create workflow", zap.Error(err))
		panic("Failed to create workflow.")

	} else {
		h.Logger.Info("Started Workflow", zap.String("WorkflowID", we.ID), zap.String("RunID", we.RunID))
	}
}

// StartWorkers starts workflow worker and activity worker based on configured options.
func (h *Runtime) StartWorkers(domainName, groupName string, options internal.WorkerOptions) {
	worker := internal.NewWorker(h.Service, domainName, groupName, options)
	err := worker.Start()
	if err != nil {
		h.Logger.Error("Failed to start workers.", zap.Error(err))
		panic("Failed to start workers")
	}
}
