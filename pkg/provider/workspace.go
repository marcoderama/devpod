package provider

import (
	"github.com/loft-sh/devpod/pkg/config"
	"github.com/loft-sh/devpod/pkg/types"
)

type Workspace struct {
	// ID is the workspace id to use
	ID string `json:"id,omitempty"`

	// Folder is the local folder where workspace related contents will be stored
	Folder string `json:"folder,omitempty"`

	// Provider is the provider used to create this workspace
	Provider WorkspaceProviderConfig `json:"provider,omitempty"`

	// Server is the server to use for this workspace
	Server WorkspaceServerConfig `json:"server,omitempty"`

	// IDE holds IDE specific settings
	IDE WorkspaceIDEConfig `json:"ide,omitempty"`

	// Source is the source where this workspace will be created from
	Source WorkspaceSource `json:"source,omitempty"`

	// CreationTimestamp is the timestamp when this workspace was created
	CreationTimestamp types.Time `json:"creationTimestamp,omitempty"`

	// Context is the context where this config file was loaded from
	Context string `json:"context,omitempty"`

	// Origin is the place where this config file was loaded from
	Origin string `json:"-"`
}

type WorkspaceServerConfig struct {
	// ID is the server ID to use for this workspace
	ID string `json:"serverId,omitempty"`

	// AutoDelete specifies if the server should get destroyed when
	// the workspace is destroyed
	AutoDelete bool `json:"autoDelete,omitempty"`
}

type WorkspaceIDEConfig struct {
	// IDE is the name of the ide to use
	IDE IDE `json:"ide,omitempty"`

	// Options are additional ide options
	Options map[string]string `json:"options,omitempty"`
}

type IDE string

const (
	IDENone       IDE = "none"
	IDEVSCode     IDE = "vscode"
	IDEOpenVSCode IDE = "openvscode"
	IDEGoland     IDE = "goland"
)

type WorkspaceIDEVSCode struct {
	// Browser determines if the vscode should be opened in the browser
	Browser bool `json:"browser,omitempty"`
}

type WorkspaceProviderConfig struct {
	// Name is the provider name
	Name string `json:"name,omitempty"`

	// Options are the provider options used to create the workspace
	Options map[string]config.OptionValue `json:"options,omitempty"`

	// Agent is the config from the provider
	Agent ProviderAgentConfig `json:"agent,omitempty"`
}

type WorkspaceSource struct {
	// GitRepository is the repository to clone
	GitRepository string `json:"gitRepository,omitempty"`

	// GitBranch is the branch to use
	GitBranch string `json:"gitBranch,omitempty"`

	// GitCommit is the commit to use
	GitCommit string `json:"gitCommit,omitempty"`

	// LocalFolder is the local folder to use
	LocalFolder string `json:"localFolder,omitempty"`

	// Image is the docker image to use
	Image string `json:"image,omitempty"`
}

type AgentWorkspaceInfo struct {
	// Workspace holds the workspace info
	Workspace Workspace `json:"workspace,omitempty"`

	// Folder holds the workspace folder on the remote server
	Folder string `json:"-"`
}

func (w WorkspaceSource) String() string {
	if w.GitRepository != "" {
		if w.GitBranch != "" {
			return w.GitRepository + "@" + w.GitBranch
		}
		if w.GitCommit != "" {
			return w.GitRepository + "@" + w.GitCommit
		}

		return w.GitRepository
	}

	if w.LocalFolder != "" {
		return w.LocalFolder
	}

	return w.Image
}
