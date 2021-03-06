package core

import (
	"fmt"
	"time"

	fqdn "github.com/ShowMax/go-fqdn"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/vjeantet/bitfan/core/webhook"
)

type Pipeline struct {
	Uuid               string
	Label              string
	agents             map[int]*Agent
	ConfigLocation     string
	ConfigHostLocation string
	StartedAt          time.Time

	Description string

	Webhooks   []webhook.Hook
	Schedulers []schedulerJob
}

func NewPipeline() *Pipeline {
	uid, _ := uuid.NewV4()

	return &Pipeline{
		Uuid:               uid.String(),
		Label:              uid.String(),
		Description:        "",
		ConfigLocation:     "",
		ConfigHostLocation: fqdn.Get(),
		agents:             map[int]*Agent{},
	}
}

func (p *Pipeline) AddAgent(a Agent) error {
	a.PipelineName = p.Label
	a.PipelineUUID = p.Uuid
	p.agents[a.ID] = &a
	return nil
}

// Start all agents, begin with last
func (p *Pipeline) Start() (string, error) {

	if _, ok := pipelines.Load(p.Uuid); ok {
		// a pipeline with same uuid is already running
		return "", fmt.Errorf("a pipeline with uuid %s is already running", p.Uuid)
	}

	//normalize
	for i, _ := range p.agents {
		p.agents[i].AgentRecipients = whoWaitForThisAgentID(p.agents[i].ID, p.agents)
	}

	orderedAgentConfList := Sort(p.agents, SortInputsFirst)
	for _, agentConf := range orderedAgentConfList {
		agentConf.PipelineUUID = p.Uuid
		agentConf.PipelineName = p.Label
		Log().Debugf("%s Agent '%-d' ", agentConf.Type, agentConf.ID)
		err := buildAgent(agentConf)
		if err != nil {
			Log().Errorf("%s Agent '%-d': %s", agentConf.Type, agentConf.ID, err.Error())
			return "", err
		}

		// register input chan for futur reference and connecting
		// for each sources
		for _, sourcePort := range agentConf.AgentSources {
			// find agent source.ID aSource
			aSource := p.agents[sourcePort.AgentID]
			// add a(in) to aSource outputs with port
			aSource.addOutput(agentConf.packetChan, sourcePort.PortNumber)
		}
		Log().Debugf("%s Agent '%-d' configured", agentConf.Type, agentConf.ID)
	}

	orderedAgentConfList = Sort(p.agents, SortOutputsFirst)
	for _, agentConf := range orderedAgentConfList {
		Log().Debugf("start %d - %s", agentConf.ID, p.agents[agentConf.ID].Label)
		p.agents[agentConf.ID].start()
	}
	p.StartedAt = time.Now()
	pipelines.Store(p.Uuid, p)
	return p.Uuid, nil
}

func (p *Pipeline) Stop() error {
	return StopPipeline(p.Uuid)
}

func (p *Pipeline) stop() error {
	orderedAgentConfList := Sort(p.agents, SortInputsFirst)
	for _, agentConf := range orderedAgentConfList {
		Log().Debugf("stop %d - %s", agentConf.ID, p.agents[agentConf.ID].Label)
		p.agents[agentConf.ID].stop()
	}
	return nil
}

func (p *Pipeline) Agents() map[int]*Agent {
	return p.agents
}
