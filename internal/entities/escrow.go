package entities

type EscrowType string

const (
	EscrowTypeSingleRelease EscrowType = "single_release"
	EscrowTypeMultiRelease  EscrowType = "multi_release"
)

type Escrow struct {
	ContractID       string
	EscrowType       EscrowType
	Deployer         string
	FactoryContract  string
	DeployerSalt     string
	WasmHash         string
	InitFunction     string
	Amount           uint64      // Only for single release (at escrow level)
	Description      string
	EngagementID     string
	Title            string
	PlatformFee      uint32
	ReceiverMemo     string
	Flags            EscrowFlags // Only for single release (at escrow level)
	Roles            EscrowRoles
	Milestones       []Milestone
	TrustlineAddress string
}

type EscrowFlags struct {
	Approved bool // Only used in multi-release milestone flags
	Disputed bool
	Released bool
	Resolved bool
}

type EscrowRoles struct {
	ServiceProvider string
	Receiver        string
	Approver        string
	ReleaseSigner   string
	DisputeResolver string
	PlatformAddress string
}

type Milestone struct {
	Description string
	Status      string
	Approved    bool      // Only for single release
	Evidence    string
	Amount      uint64    // Only for multi release
	Flags       *EscrowFlags // Only for multi release (per milestone)
	Receiver    string    // Only for multi release (per milestone)
}
