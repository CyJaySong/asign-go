package model

// 重新发送认证验证码 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=captchaResend
type AuthCaptchaResendReqBody struct {
	SerialNo string `json:"serialNo" dc:"认证流水号"`
}

// 重新发送认证验证码 Res
type AuthCaptchaResendResBody struct{}

// 认证验证码校验 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=captchaVerify
type AuthCaptchaVerifyReqBody struct {
	SerialNo string `json:"serialNo" dc:"认证流水号"`
	Captcha  string `json:"captcha" dc:"短信验证码"`
}

// 认证验证码校验 Res
type AuthCaptchaVerifyResBody struct {
	Result   int    `json:"result" dc:"认证结果：1.成功 2.失败"`
	SerialNo string `json:"serialNo" dc:"认证流水号"`
	Type     string `json:"type" dc:"认证类型"`
}

// 实名认证流水号查询 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=getSerialNo
type GetAuthSerialNoReqBody struct {
	UserType int    `json:"userType" dc:"认证主体类型：1.企业，2.个人"`
	UserName string `json:"userName,omitempty" dc:"姓名 或 企业名称"`
	IdNo     string `json:"idNo,omitempty" dc:"身份证 或 企业证件号"`
}

// 实名认证流水号查询 Res
type GetAuthSerialNoResBody []GetAuthSerialNoResBodyItem

type GetAuthSerialNoResBodyItem struct {
	SerialNo     string `json:"serialNo" dc:"认证流水号"`
	Result       int    `json:"result" dc:"认证结果：1.成功 2.失败"`
	UserType     int    `json:"userType" dc:"认证主体类型：1.企业，2.个人"`
	AuthTypeCode int    `json:"authTypeCode" dc:"认证类型码"`
	AuthTypeName string `json:"authTypeName" dc:"认证类型"`
}

// 实名认证信息查询 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=getAuthRecordInfo
type GetAuthRecordInfoReqBody struct {
	SerialNo string `json:"serialNo" dc:"认证流水号"`
}

// 实名认证信息查询 Res
type GetAuthRecordInfoResBody struct {
	SerialNo          string `json:"serialNo" dc:"认证流水号"`
	Result            int    `json:"result" dc:"认证结果：1.成功 2.失败"`
	ResultMsg         string `json:"resultMsg,omitempty" dc:"当结果是失败时，返回失败原因"`
	UserType          int    `json:"userType" dc:"认证主体类型：1.企业，2.个人"`
	RealName          string `json:"realName,omitempty" dc:"个人姓名 / 法人姓名"`
	IdCardNo          string `json:"idCardNo,omitempty" dc:"个人身份证 / 法人身份证"`
	Mobile            string `json:"mobile,omitempty" dc:"个人手机号 / 法人手机号"`
	CompanyName       string `json:"companyName,omitempty" dc:"企业名称"`
	CreditCode        string `json:"creditCode,omitempty" dc:"社会统一信用代码"`
	BankCard          string `json:"bankCard,omitempty" dc:"用户银行卡号"`
	AuthTypeCode      string `json:"authTypeCode,omitempty" dc:"认证类型码"`
	AuthTypeName      string `json:"authTypeName,omitempty" dc:"认证类型"`
	AgentResult       int    `json:"agentResult,omitempty" dc:"办理人认证结果：0.认证中 1.认证成功 2.认证失败"`
	AgentName         string `json:"agentName,omitempty" dc:"办理人姓名"`
	AgentIdCardNo     string `json:"agentIdCardNo,omitempty" dc:"办理人身份证"`
	AgentMobile       string `json:"agentMobile,omitempty" dc:"办理人手机号"`
	AgentAuthTypeCode int    `json:"agentAuthTypeCode,omitempty" dc:"办理人认证类型码"`
	AgentAuthTypeName string `json:"agentAuthTypeName,omitempty" dc:"办理人认证类型"`
}

// 查询人脸核身结果 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=faceresult
type GetFaceResultReqBody struct {
	BizId string `json:"bizId" dc:"接入商自定义业务ID，最长32位"`
}

// 查询人脸核身结果 Res
type GetFaceResultResBody struct {
	Status      int     `json:"status" dc:"人脸结果状态; 1:等待认证;2:认证成功;3:认证失败"`
	Msg         string  `json:"msg" dc:"认证结果消息list数据详解"`
	LivingScore float64 `json:"livingScore" dc:"活体分数"`
	VerifyScore float64 `json:"verifyScore" dc:"比对分数"`
}

// 个人运营商三要素认证 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=mobile3
type AuthPersonMobile3ReqBody struct {
	RealName string `json:"realName" dc:"真实姓名"`
	IdCardNo string `json:"idCardNo" dc:"身份证号"`
	Mobile   string `json:"mobile" dc:"手机号码（限中国大陆11位手机号）"`
}

// 个人运营商三要素认证 Res
type AuthPersonMobile3ResBody struct {
	Result   int    `json:"result" dc:"认证结果：0.暂无结果/认证中 1.成功 2.失败"`
	SerialNo string `json:"serialNo" dc:"认证流水号"`
	Type     string `json:"type" dc:"认证类型"`
}

// 企业法人运营商三要素认证 Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=companyMobile3
type AuthCompanyMobile3ReqBody struct {
	CompanyName string `json:"companyName" dc:"企业名称"`
	CreditCode  string `json:"creditCode" dc:"社会统一信用代码"`
	RealName    string `json:"realName" dc:"真实姓名"`
	IdCardNo    string `json:"idCardNo" dc:"身份证号"`
	Mobile      string `json:"mobile" dc:"手机号码（限中国大陆11位手机号）"`
}

// 企业法人运营商三要素认证 Res
type AuthCompanyMobile3ResBody struct {
	Result   int    `json:"result" dc:"认证结果：0.暂无结果/认证中 1.成功 2.失败"`
	SerialNo string `json:"serialNo" dc:"认证流水号"`
	Type     string `json:"type" dc:"认证类型"`
}

// 个体实名认证网页版（新规版） Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=personIdentifyUrl
type AuthPersonIdentifyUrlReqBody struct {
	Account      string `json:"account,omitempty" dc:"用户唯一识别码，可传值需要重新实名认证的用户account"`
	AuthTypeList []int  `json:"authTypeList,omitempty" dc:"选择显示给办理人的页面认证方式"`
	// 选择显示给办理人的页面认证方式
	// 151:运营商三要素认证;152:银行卡四要素认证;153:人脸活体认证;154:个人视频双录认证(需联系商务人员开启权限后使用)
	RealName     string `json:"realName" dc:"用户姓名（传值后认证页面上不可修改，传空表示可以由用户输入）"`
	IdCardNo     string `json:"idCardNo,omitempty" dc:"身份证号（传值后认证页面上不可修改，传空表示可以由用户输入）"`
	Mobile       string `json:"mobile,omitempty" dc:"手机号码（传值后认证页面上不可修改，传空表示可以由用户输入）"`
	BankCard     string `json:"bankCard,omitempty" dc:"银行卡号（传值后认证页面上不可修改，传空表示可以由用户输入）"`
	RedirectUrl  string `json:"redirectUrl,omitempty" dc:"认证完成后同步重定向跳转H5页面地址"`
	NotifyUrl    string `json:"notifyUrl,omitempty" dc:"认证完成后异步通知地址"`
	BizId        string `json:"bizId,omitempty" dc:"用户自定义业务编号，在异步通知时会发送回业务方（最长60位）"`
	FaceAuthMode int    `json:"faceAuthMode,omitempty" dc:"人脸识别方式"`
	// 人脸识别方式 1:支付宝（不可在支付宝小程序中接入）2:H5（默认）4:微信小程序 5:支付宝小程序 9:支付宝H5版
	Resubmit       int    `json:"resubmit,omitempty" dc:"认证失败后是否允许重新提交;0否(默认); 1:是"`
	ShowResult     int    `json:"showResult,omitempty" dc:"认证结束后是否展示结果页;0否(默认); 1:是"`
	MobileRequired int    `json:"mobileRequired,omitempty" dc:"签约手机号采集标记;0否(默认); 1:是"`
	IsIframe       int    `json:"isIframe,omitempty" dc:"H5人脸需要开启无Cookie模式,或使用IFrame框架;0否(默认); 1:是"`
	IsNotice       int    `json:"isNotice,omitempty" dc:"是否将认证链接以短信形式通知用户;0:否(默认); 1:是"`
	NoticeMobile   string `json:"noticeMobile,omitempty" dc:"通知手机号"`
	NeedSignPwd    int    `json:"needSignPwd,omitempty" dc:"实名认证时是否采集签约密码参数;0:否(默认); 1:是"`
}

// 企业实名认证网页版（新规版） Req https://preweb.asign.cn/platform/openDoc/docDetail?mid=companyIdentifyUrl
type AuthCompanyIdentifyUrlReqBody struct {
	Account      string `json:"account,omitempty" dc:"用户唯一识别码，可传值需要重新实名认证的用户account"`
	AuthTypeList []int  `json:"authTypeList,omitempty" dc:"选择显示给办理人的页面认证方式"`
	// 选择显示给办理人的页面认证方式
	// 151:运营商三要素认证;152:银行卡四要素认证;153:人脸活体认证;154:个人视频双录认证(需联系商务人员开启权限后使用)
	AgentSignPower      bool  `json:"agentSignPower" dc:"是否授权经办人签署"`
	CompanyAuthTypeList []int `json:"companyAuthTypeList,omitempty" dc:"选择显示的公司认证方式"`
	// 选择显示的公司认证方式 351:法定代表人认证; 369:企业线下盖章授权认证（默认）;
	// 370:企业线下盖章授权认证+对公账户打款（反向打款）（需联系商务人员开启权限后使用）;
	// 371:企业线下盖章授权认证+对公账户打款（正向打款）（需联系商务人员开启权限后使用）
	CompanyName         string   `json:"companyName,omitempty" dc:"预填企业名称"`
	CreditCode          string   `json:"creditCode,omitempty" dc:"预填社会统一信用代码号"`
	LegalPersonName     string   `json:"legalPersonName,omitempty" dc:"预填法人姓名"`
	LegalPersonIdCardNo string   `json:"legalPersonIdCardNo,omitempty" dc:"预填法人身份证号"`
	LegalPersonMobile   string   `json:"legalPersonMobile,omitempty" dc:"预填法人手机号"`
	AgentName           string   `json:"agentName,omitempty" dc:"预填办理人姓名"`
	AgentIdCardNo       string   `json:"agentIdCardNo,omitempty" dc:"预填办理人身份证"`
	AgentMobile         string   `json:"agentMobile,omitempty" dc:"预填办理人手机号"`
	ImmutableInfoList   []string `json:"immutableInfoList,omitempty" dc:"指定页面不可修改参数"` //
	// 指定页面不可修改参数 companyName:公司名称;creditCode:社会统一信用代码;legalPersonName:法人名称;legalPersonIdCardNo:法人身份证
	// legalPersonMobile:法人手机号;agentName:代理人名称;agentIdCardNo:代理人身份证;agentMobile:代理人手机号
	RedirectUrl  string `json:"redirectUrl,omitempty" dc:"认证完成后同步重定向跳转H5页面地址"`
	NotifyUrl    string `json:"notifyUrl,omitempty" dc:"认证完成后异步通知地址"`
	AuthMode     int    `json:"authMode,omitempty" dc:"是否开启支付宝快捷认证,1:不开启(默认),2:开启支付宝企业认证"`
	ZhiMaScheme  string `json:"zhimaScheme,omitempty" dc:"是否开启支付宝快捷认证,1:不开启(默认),2:开启支付宝企业认证"`
	BizId        string `json:"bizId,omitempty" dc:"用户自定义业务编号，在异步通知时会发送回业务方（最长60位）"`
	FaceAuthMode int    `json:"faceAuthMode,omitempty" dc:"人脸识别方式"`
	// 人脸识别方式 1:支付宝（不可在支付宝小程序中接入）2:H5（默认）4:微信小程序 5:支付宝小程序 9:支付宝H5版
	ShowResult     int    `json:"showResult,omitempty" dc:"认证结束后是否展示结果页;0否(默认); 1:是"`
	MobileRequired int    `json:"mobileRequired,omitempty" dc:"签约手机号采集标记;0否(默认); 1:是"`
	IsIframe       int    `json:"isIframe,omitempty" dc:"H5人脸需要开启无Cookie模式,或使用IFrame框架;0否(默认); 1:是"`
	IsNotice       int    `json:"isNotice,omitempty" dc:"是否将认证链接以短信形式通知用户;0:否(默认); 1:是"`
	NoticeMobile   string `json:"noticeMobile,omitempty" dc:"通知手机号"`
	NeedSignPwd    int    `json:"needSignPwd,omitempty" dc:"实名认证时是否采集签约密码参数;0:否(默认); 1:是"`
}

// 实名认证网页版（新规版）
type AuthIdentifyUrlResBody struct {
	Result      int    `json:"result" dc:"认证结果：0.暂无结果/认证中 1.成功 2.失败"`
	SerialNo    string `json:"serialNo" dc:"认证流水号"`
	Type        string `json:"type" dc:"认证类型"`
	IdentifyUrl string `json:"identifyUrl" dc:"认证链接"`
	ShortUrl    string `json:"shortUrl" dc:"认证链接"`
}
