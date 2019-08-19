package blockdata

type Chain struct {
	Height int64
}

type Block struct {
	BlockHash         string
	Timestamp         int64
	Height            int64 `json:",string"`
	TransactionNumber int
	PreHash           string
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
	Asset                  Assets                 `json:"-"`
	Ast                    AstInfo                `json:"AstInfo"`
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
	Height                             int64    `json:"height"`
	Timestamp                          int64    `json:"timestamp"`
	TxID, Chaincode, Method, ChannelId string   `json:"txid"`
	CreatedFlag                        bool     `json:"createdFlag"`
	TxArgs                             [][]byte `json:"-"`
	Hash                               string   `json:"hash"`
	Time                               string   `json:"time"`
	PreHash                            string   `json:"preHash"`
}

//-------------------------------------------------------2019 08 16 新数据结构上链

/*
应收账款数据结构整理：
资产包：
		资产：（多个）
				1.合同
				2.发票
		融资信息：（单个）
		增信措施信息：（单个）
					1. 担保
					2. 质押
					3. 抵押
		发布人信息：（单个）

*/

//	发布人信息
type AstSendInfo struct {
	AstSenduuid      string `json:"astSenduuid" pigkey:""`
	AstSendfMasterID string `json:"astSendfMasterID`
	AstSendName      string `json:"astSendName`
	AstSendid        string `json:"astSendid`
	AstSendContact   string `json:"astSendContact`
	AstSendTime      string `json:"astSendTime`
}

//	质押品
type AstCreEnsureList struct {
	AstCreEnsuuid      string `json:"astCreEnsuuid piokey:""` // 质押品uuid
	AstCreEnsfMasterID string `json:"astCreEnsfMasterID`      //质押品fMasterID
	AstCreEnsType      string `json:"astCreEnsType`           //质押品类型
	AstCreEnsName      string `json:"astCreEnsName`           // 质押品名称
	AstCreEnsOwner     string `json:"astCreEnsOwner`          //质押品所有人
}

//	抵押品
type AstCrePledgeList struct {
	AstCrePleuuid      string `json:"astCrePleuuid pitkey:""` // 抵押品uuid
	AstCrePlefMasterID string `json:"astCrePlefMasterID `     // 抵押品fMasterID
	AstCrePleType      string `json:"astCrePleType `          // 抵押品类型
	AstCrePleName      string `json:"astCrePleName `          // 抵押品名称
	AstCrePleOwner     string `json:"astCrePleOwner `         // 抵押品所有人
}

// 担保 信息
type AstCreGuarantyList struct {
	AstCreGuauuid       string `json:"astCreGuauuid piakey:""` // 担保uuid
	AstCreGuafMasterID  string `json:"astCreGuafMasterID `     // 担保fMasterID
	AstCreGuaType       string `json:"astCreGuaType `          // 增信措施类型
	AstCreGuaName       string `json:"astCreGuaName `          // 保证人名称
	AstCreGuaManner     string `json:"astCreGuaManner `        // 担保方式
	AstCreGuaEnsureId   string `json:"astCreGuaEnsureId `      // 担保人ID
	AstCreGuaEnsureName string `json:"astCreGuaEnsureName `    // 担保人名称
}

// 担保列表
type AstCreditInfo struct {
	AstCreGuarantyList []AstCreGuarantyList `json:"astCreGuarantyList ` // 担保
	AstCrePledgeList   []AstCrePledgeList   `json:"astCrePledgeList `   // 抵押品
	AstCreEnsureList   []AstCreEnsureList   `json:"astCreEnsureList `   // 质押品

}

// 融资 信息
type AstFinancingInfo struct {
	AstFinuuid        string `json:"astFinuuid" pivkey:""` // 融资uuid
	AstFinfMasterID   string `json:"astFinfMasterID `      // 融资fMasterID
	AstFinPrice       string `json:"astFinPrice `          // 融资金额
	AstCreGuaName     string `json:"astFinLimit `          // 融资期限
	AstCreGuaManner   string `json:"astFinType `           // 融资方式
	AstCreGuaEnsureId string `json:"astFinUsefor `         // 资金用途
}

// 发票  信息
type AstInv struct {
	AstInvuuid           string         `json:"astInvuuid" pkey:""`   // 发票uuid
	AstInvfMasterID      string         `json:"astInvfMasterID `      // 发票fMasterID
	AstInvType           string         `json:"astInvType `           // 发票类型
	AstInvNum            string         `json:"astInvNum `            // 发票号码
	AstInvCode           string         `json:"astInvCode `           // 发票代码
	AstInvChecksum       string         `json:"astInvChecksum `       // 校验码
	AstInvPrice          string         `json:"astInvPrice `          // 发票金额
	AstInvUnit           string         `json:"astInvUnit `           // 金额单位
	AstInvTime           string         `json:"astInvTime `           // 发票开具日期
	AstInvBuyerTaxNum    string         `json:"astInvBuyerTaxNum `    // 购买方纳税识别号
	AstInvSellerTaxNum   string         `json:"astInvSellerTaxNum `   // 销售方纳税识别号
	AstInvCheckResult    string         `json:"astInvCheckResult `    // 发票验真结果
	AstInvAttachmentList AttachmentList `json:"astInvAttachmentList ` // 发票列表
}

// 发票  信息
type AstCon struct {
	AstConuuid           string         `json:"astConuuid" pkey:""`   // 合同uuid
	AstConfMasterID      string         `json:"astConfMasterID `      // 合同fMasterID
	AstConNo             string         `json:"astConNo `             // 合同编号
	AstConName           string         `json:"astConName `           // 合同名称
	AstConPrice          string         `json:"astConPrice `          // 合同金额
	AstConPriceUnit      string         `json:"astConPriceUnit `      // 金额单位
	AstConType           string         `json:"astConType `           // 币种
	AstConPayerTaxNum    string         `json:"astConPayerTaxNum `    // 付款方纳税识别号
	AstConPayerName      string         `json:"astConPayerName `      // 付款方名称
	AstConPayeeTaxNum    string         `json:"astConPayeeTaxNum `    // 收款方纳税识别号
	AstConPayeeName      string         `json:"astConPayeeName `      // 收款方名称
	AstConTime           string         `json:"astConTime `           // 合同签署时间
	AstConCount          string         `json:"astConCount `          // 单次/多次
	AstConDays           string         `json:"astConDays `           // 账期
	AstConAttachmentList AttachmentList `json:"astConAttachmentList ` // 附件列表
}

//	附件 信息
type AttachmentList struct {
	AstAttachNo        string `json:"astAttachNo" pkey:""` //附件附加编号
	IpfsHash           string `json:"ipfsHash `            //IPFS-Hash
	AstAttachfMasterID string `json:"astAttachfMasterID `  //附件附加fMasterID
	AstAttachName      string `json:"astAttachName `       //附件附加名称
	AstAttachType      string `json:"astAttachType `       //附件附加类型
	AstAttachMD        string `json:"astAttachMD `         //附件附MD
	AstAttachAddr      string `json:"astAttachAddr `       //附件附加地址
}

// 资产信息  应收账款信息
type AstAssetsList struct {
	//
	AstAssetsuuid           string   `json:"astAssetsuuid" pkey:""`   //资产
	AstAssetsfMasterID      string   `json:"astAssetsfMasterID `      //外键
	AstAssetsType           string   `json:"astAssetsType `           //资产类型
	AstAssetsIntroduce      string   `json:"astAssetsIntroduce `      //资产简介
	AstAssetsCreditorName   string   `json:"astAssetsCreditorName `   //债权人名称
	AstAssetsCreditorTaxNum string   `json:"astAssetsCreditorTaxNum ` //债权人纳税识别号
	AstAssetsDebtorName     string   `json:"astAssetsDebtorName `     //债务人名称
	AstAssetsDebtorTaxNum   string   `json:"astAssetsDebtorTaxNum `   //债务人纳税识别号
	AstAssetsValuation      string   `json:"astAssetsValuation `      //资产估值
	AstAssetsHonour         string   `json:"astAssetsHonour `         //资产兑付日
	AstAssetsState          string   `json:"astAssetsState `          //当前状态
	AstAssetsPrimeval       string   `json:"astAssetsPrimeval `       //是否为原始资产
	AstContractInfoList     []AstCon `json:"astContractInfoList `     //合同列表
	AstInvoiceInfoList      []AstInv `json:"astInvoiceInfoList `      //发票列表
}

//  资产包 信息
type AstAssetsInfo struct {
	//
	BaseAsset
	AstPackageuuid       string           `json:"astPackageuuid" pidkey:""` //资产包UUID
	AstPackageNo         string           `json:"astPackageNo `             //资产包编号
	AstPackagePlatForm   string           `json:"astPackagePlatForm `       //平台信息-鼎链名称
	AstPackagePlatFormID string           `json:"astPackagePlatFormID `     //平台信息-鼎链ID
	AstPackageName       string           `json:"astPackageName `           //资产包名称
	AstPackageNumber     string           `json:"astPackageNumber `         //资产数量
	AstPackageOwnerName  string           `json:"astPackageOwnerName `      //资产所有人名称
	AstPackageOwnerId    string           `json:"astPackageOwnerId `        //资产所有人ID
	AstPackageSplit      string           `json:"astPackageSplit `          //是否允许拆分
	AstPackageEvaluation string           `json:"astPackageEvaluation `     //综合评价
	AstAssetsList        []AstAssetsList  `json:"astAssetsList `            //合同发票列表
	AstFinancingInfo     AstFinancingInfo `json:"astFinancingInfo `         //融资信息
	AstCreditInfo        AstCreditInfo    `json:"astCreditInfo `            //增信措施信息
	AstSendInfo          AstSendInfo      `json:"astSendInfo `              //发布人信息
}

// 应收账款   主体包
type AstInfo struct {
	BaseAsset
	AstAssetsInfo AstAssetsInfo `json:"astAssetsInfo `

	// 单个应收账款的UUID
	AstAssetsInfoUUID string `json:"uuid" pkey:""` //资产包UUID
}
