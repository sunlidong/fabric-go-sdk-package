package blockdata

type Chain struct {
	Height int64
}

type Block struct {
	BlockHash         string
	Timestamp         int64
	Height            int64 `json:",string"`
	TransactionNumber int
	PreHash 		string
	Transaction       []*ChainTransaction `json:"Transaction"`
}

type ChainTransaction struct {
	Height                             int64
	Timestamp                          int64
	TxID, Chaincode, Method, ChannelId string
	CreatedFlag                        bool
	TxArgs                             [][]byte `json:"-"`
	Hash                               string
	Time                               string
	PreHash                            string


}
type ChainTransaction1 struct {
	Height                             int64       `json:"height"`
	Timestamp                          int64       `json:"timeStamp"`
	TxID, Chaincode, Method, ChannelId string      `json:"txid"`
	CreatedFlag                        bool        `json:"createdFlag"`
	TxArgs                             [][]byte    `json:"-"`
	Hash                               string      `json:"hash"`
	Time                               string      `json:"time"`
	PreHash                            string      `json:"preHash"`
	Data                               []string    `json:"-"`
	Datalist                           []*Datalist `json:"-"`
	Asset                              Assets      `json:"assetData"`
}
type ChainTransactionQuery struct {
	ChainTransactionConfig ChainTransactionConfig `json:"config"`
	Data                   []string               `json:"-"`
	Datalist               []*Datalist            `json:"-"`
	Asset                  Assets                 `json:"assetData"`
}

const (
	TxStatus_Success = 0
	TxStatus_Fail    = 1
)

type ChainTxEvents struct {
	TxID, Chaincode, Name string
	Status                int
	Payload               []byte `json:"-"`
}

type ChainBlock struct {
	Height       int64 `json:",string"`
	Hash         string
	TimeStamp    string              `json:",omitempty"`
	Transactions []*ChainTransaction `json:"-"`
	TxEvents     []*ChainTxEvents    `json:"-"`
}

type Datalist struct {
	Arrset []Assets
	Proset []ProAttachmentlist
}

type Assets struct {
	BaseAsset
	ProUUID                   string   `json:"prouuid"`
	ProNo                     string   `json:"proNo" pkey:""`
	ProPlatForm               string   `json:"proPlatForm"`
	ProName                   string   `json:"proName"`
	ProAddress                string   `json:"proAddress"`
	ProOwnerName              string   `json:"proOwnerName"`
	ProOwnerId                string   `json:"proOwnerId"`
	ProDisposalName           string   `json:"proDisposalName"`
	ProDisposalId             string   `json:"proDisposalId"`
	ProType                   string   `json:"proType"`
	ProSendTime               string   `json:"proSendTime"`
	ProFinaningNum            string   `json:"proFinaningNum"`
	ProFinancingUnit          string   `json:"proFinancingUnit"`
	ProFinancingEndLine       string   `json:"proFinancingEndLine"`
	ProFinancingEndUnit       string   `json:"proFinancingEndUnit"`
	ProFinancingWay           string   `json:"proFinancingWay"`
	ProFinancingPurpose       string   `json:"proFinancingPurpose"`
	ProEcreditWay             string   `json:"proEcreditWay"`
	ProEcreditAssure          string   `json:"proEcreditAssure"`
	ProEcreditGuaranteeName   string   `json:"proEcreditGuaranteeName"`
	ProEcreditGuaranteeId     string   `json:"proEcreditGuaranteeId"`
	ProEcreditGuaranteeType   string   `json:"proEcreditGuaranteeType"`
	ProReportDuediligenceHash string   `json:"proReportDuediligenceHash"`
	ProReportDuediligenceId   string   `json:"proReportDuediligenceId"`
	ProReportThirdHash        string   `json:"proReportThirdHash"`
	ProReportThirdId          string   `json:"proReportThirdId"`
	ProState                  string   `json:"proState"`
	ProNote                   []string `json:"proNote"`
	ProAttachmentlist         []struct {
		IpfsHash      string `json:"ipfsHash"`
		ProAttachName string `json:"proAttachName"`
		ProAttachType string `json:"proAttachType"`
		ProAttachNo   string `json:"proAttachNo" `
		ProAttachMD   string `json:"proAttachMD"`
		ProAttachAddr string `json:"proAttachAddr"`
	}
}

// 其他附件
type ProAttachmentlist struct {
	BaseAsset
	IpfsHash      string `json:"ipfsHash"`
	ProAttachName string `json:"proAttachName"`
	ProAttachType string `json:"proAttachType"`
	// 唯一标识
	ProAttachNo   string `json:"proAttachNo" pkey:""`
	ProAttachMD   string `json:"proAttachMD"`
	ProAttachAddr string `json:"proAttachAddr"`
}

// 资产基本字段
type BaseAsset struct {
	// IsDeleted  bool   `json:"isDeleted"`
	UpdateDate string `json:"updateDate"`
	// 仅用在底层资产中、其Parent有可能为产品、保理
	ParentType string `json:"parentType"`
	ParentID   string `json:"parentID"`
}

type ChainTransactionConfig struct {
	Height                             int64       `json:"height"`
	Timestamp                          int64       `json:"timestamp"`
	TxID, Chaincode, Method, ChannelId string      `json:"txid"`
	CreatedFlag                        bool        `json:"createdFlag"`
	TxArgs                             [][]byte    `json:"-"`
	Hash                               string      `json:"hash"`
	Time                               string      `json:"time"`
	PreHash                            string      `json:"preHash"`
}