package promptpay

// Bank is a Thai bank identified by its 3-digit Bank of Thailand (BOT) code,
// with English and Thai names.
//
// The codes are the BOT sending-bank codes carried in a bank slip's 00>01
// field (see DecodeSlip). NameTh may be empty for banks with no common Thai
// name.
type Bank struct {
	// Code is the 3-digit BOT bank code (e.g. "014").
	Code string
	// NameEn is the English bank name (e.g. "Siam Commercial Bank").
	NameEn string
	// NameTh is the Thai bank name (may be empty).
	NameTh string
}

// thaiBanks holds all known BOT 3-digit sending-bank codes, keyed by code.
//
// Source: Bank of Thailand bank codes, cross-referenced with the 2C2P
// payout-bank-code table.
var thaiBanks = map[string]Bank{
	"002": {"002", "Bangkok Bank", "ธนาคารกรุงเทพ"},
	"004": {"004", "Kasikornbank", "ธนาคารกสิกรไทย"},
	"006": {"006", "Krung Thai Bank", "ธนาคารกรุงไทย"},
	"011": {"011", "TMBThanachart Bank (ttb)", "ธนาคารทหารไทยธนชาต"},
	"014": {"014", "Siam Commercial Bank", "ธนาคารไทยพาณิชย์"},
	"017": {"017", "Citibank", ""},
	"020": {"020", "Standard Chartered Bank (Thai)", ""},
	"022": {"022", "CIMB Thai Bank", "ธนาคารซีไอเอ็มบีไทย"},
	"024": {"024", "United Overseas Bank (Thai)", "ธนาคารยูโอบี"},
	"025": {"025", "Bank of Ayudhya (Krungsri)", "ธนาคารกรุงศรีอยุธยา"},
	"030": {"030", "Government Savings Bank", "ธนาคารออมสิน"},
	"031": {"031", "HSBC", ""},
	"033": {"033", "Government Housing Bank", "ธนาคารอาคารสงเคราะห์"},
	"034": {"034", "Bank for Agriculture and Agricultural Cooperatives (BAAC)", "ธนาคารเพื่อการเกษตรและสหกรณ์การเกษตร"},
	"035": {"035", "Export-Import Bank of Thailand (EXIM)", "ธนาคารเพื่อการส่งออกและนำเข้าแห่งประเทศไทย"},
	"065": {"065", "Thanachart Bank", "ธนาคารธนชาต"},
	"066": {"066", "Islamic Bank of Thailand", "ธนาคารอิสลามแห่งประเทศไทย"},
	"067": {"067", "Tisco Bank", "ธนาคารทิสโก้"},
	"069": {"069", "Kiatnakin Phatra Bank (KKP)", "ธนาคารเกียรตินาคินภัทร"},
	"070": {"070", "ICBC (Thai)", ""},
	"071": {"071", "Thai Credit Bank", "ธนาคารไทยเครดิต"},
	"073": {"073", "Land and Houses Bank (LH Bank)", "ธนาคารแลนด์ แอนด์ เฮ้าส์"},
}

// ThaiBankByCode returns the Bank for a 3-digit BOT code. ok is false when the
// code is unknown or malformed (it never returns a partial result). Lookup is
// pure and offline.
func ThaiBankByCode(code string) (*Bank, bool) {
	b, ok := thaiBanks[code]
	if !ok {
		return nil, false
	}
	return &b, true
}
