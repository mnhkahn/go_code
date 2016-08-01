package main

// https://github.com/mncaudill/ketama/blob/master/ketama_test.go

import (
	"crypto/sha1"
	"log"
	"sort"
	"strconv"
)

var (
	NODE_NUM = 600
	CatIds   = []string{"653", "654", "655", "659", "671", "672", "673", "674", "675", "677", "678", "679", "680", "681", "682", "683", "684", "686", "687", "688", "689", "690", "691", "692" /*, "693", "694", "695", "696", "697", "698", "699", "700", "701", "702", "703", "716", "717", "718", "719", "720", "721", "722", "723", "724", "725", "726", "727", "728", "729", "730", "731", "732", "733", "736", "738", "739", "740", "741", "742", "745", "747", "748", "749", "750", "751", "752", "753", "754", "755", "756", "757", "758", "759", "760", "761", "762", "782", "794", "795", "798", "801", "803", "806", "807", "808", "823", "825", "826", "828", "829", "830", "831", "832", "833", "834", "835", "836", "837", "838", "840", "841", "842", "843", "844", "845", "846", "847", "848", "851", "852", "854", "860", "861", "862", "863", "864", "866", "867", "868", "869", "870", "877", "878", "880", "881", "882", "897", "898", "899", "900", "901", "902", "934", "962", "963", "965", "966", "967", "968", "975", "981", "982", "983", "1047", "1048", "1049", "1050", "1051", "1052", "1098", "1099", "1105", "1195", "1199", "1203", "1219", "1225", "1229", "1261", "1263", "1274", "1276", "1277", "1278", "1279", "1283", "1287", "1289", "1290", "1291", "1292", "1293", "1294", "1295", "1299", "1300", "1301", "1342", "1343", "1345", "1346", "1348", "1349", "1350", "1354", "1355", "1356", "1362", "1364", "1365", "1368", "1369", "1371", "1376", "1378", "1381", "1383", "1384", "1385", "1386", "1387", "1389", "1390", "1391", "1392", "1393", "1394", "1395", "1396", "1397", "1398", "1400", "1401", "1402", "1403", "1404", "1405", "1406", "1407", "1408", "1409", "1410", "1411", "1412", "1416", "1419", "1420", "1421", "1422", "1423", "1424", "1425", "1426", "1427", "1428", "1429", "1440", "1442", "1443", "1444", "1445", "1446", "1449", "1455", "1462", "1463", "1466", "1472", "1473", "1474", "1475", "1476", "1477", "1478", "1479", "1480", "1483", "1484", "1487", "1502", "1503", "1504", "1505", "1506", "1508", "1509", "1510", "1514", "1516", "1517", "1518", "1519", "1523", "1524", "1525", "1526", "1527", "1528", "1533", "1534", "1536", "1537", "1538", "1539", "1546", "1548", "1550", "1551", "1552", "1553", "1555", "1556", "1557", "1558", "1559", "1560", "1562", "1563", "1564", "1565", "1566", "1567", "1568", "1569", "1581", "1583", "1584", "1585", "1589", "1590", "1591", "1592", "1593", "1594", "1595", "1601", "1602", "1621", "1623", "1624", "1625", "1626", "1627", "1628", "1629", "1630", "1631", "1632", "1633", "1634", "1647", "1648", "1649", "1650", "1651", "1653", "1654", "1655", "1656", "1657", "1658", "1659", "1660", "1661", "1662", "1663", "1667", "1669", "1670", "1671", "1694", "1695", "1696", "1697", "1698", "1699", "1700", "1701", "1706", "2562", "2575", "2576", "2577", "2580", "2584", "2587", "2588", "2589", "2599", "2600", "2601", "2603", "2606", "2615", "2628", "2629", "2631", "2641", "2642", "2643", "2644", "2647", "2648", "2653", "2656", "2669", "2670", "2675", "2676", "2677", "2678", "2679", "2680", "2687", "2688", "2691", "2694", "3977", "3979", "3982", "3983", "3986", "3997", "3998", "4000", "4693", "4698", "4699", "4702", "4833", "4834", "4835", "4836", "4837", "4838", "4839", "4840", "4934", "4935", "4937", "4939", "4940", "4941", "4942", "4943", "4944", "4945", "4946", "4947", "4948", "4950", "4952", "4954", "4955", "4956", "4979", "4991", "4996", "4997", "4998", "4999", "5000", "5001", "5002", "5003", "5004", "5005", "5006", "5008", "5009", "5010", "5011", "5012", "5014", "5019", "5020", "5021", "5022", "5023", "5024", "5026", "5027", "5028", "5029", "5030", "5064", "5065", "5066", "5069", "5070", "5071", "5074", "5087", "5089", "5092", "5094", "5096", "5104", "5106", "5107", "5112", "5146", "5147", "5148", "5149", "5150", "5152", "5153", "5155", "5156", "5160", "5161", "5162", "5163", "5164", "5256", "5257", "5258", "5259", "5260", "5261", "5262", "5265", "5266", "5270", "5271", "6145", "6146", "6147", "6148", "6149", "6150", "6151", "6152", "6155", "6156", "6157", "6158", "6159", "6160", "6161", "6162", "6163", "6164", "6165", "6167", "6168", "6169", "6170", "6172", "6173", "6174", "6175", "6176", "6177", "6178", "6179", "6180", "6181", "6182", "6183", "6184", "6185", "6186", "6187", "6188", "6189", "6190", "6191", "6192", "6193", "6194", "6195", "6197", "6198", "6199", "6200", "6201", "6202", "6203", "6204", "6205", "6206", "6207", "6209", "6210", "6211", "6212", "6214", "6215", "6216", "6218", "6219", "6220", "6221", "6222", "6223", "6224", "6225", "6227", "6228", "6229", "6230", "6231", "6232", "6234", "6235", "6236", "6237", "6239", "6240", "6241", "6242", "6243", "6244", "6245", "6246", "6247", "6248", "6249", "6253", "6254", "6255", "6257", "6258", "6259", "6260", "6261", "6262", "6263", "6264", "6265", "6266", "6268", "6269", "6271", "6272", "6273", "6274", "6275", "6276", "6277", "6278", "6279", "6280", "6281", "6282", "6283", "6284", "6286", "6287", "6288", "6289", "6291", "6292", "6293", "6294", "6296", "6297", "6298", "6299", "6300", "6301", "6302", "6303", "6304", "6305", "6306", "6307", "6308", "6309", "6313", "6314", "6315", "6316", "6317", "6319", "6727", "6729", "6734", "6736", "6737", "6738", "6739", "6740", "6742", "6743", "6745", "6747", "6749", "6752", "6753", "6756", "6757", "6762", "6766", "6767", "6768", "6769", "6770", "6774", "6779", "6782", "6784", "6785", "6786", "6788", "6791", "6792", "6795", "6796", "6798", "6801", "6804", "6806", "6807", "6850", "6854", "6855", "6860", "6861", "6863", "6864", "6865", "6868", "6877", "6879", "6880", "6881", "6887", "6891", "6896", "6906", "6907", "6908", "6909", "6910", "6911", "6912", "6913", "6914", "6915", "6916", "6917", "6918", "6920", "6922", "6962", "6963", "6964", "6965", "6967", "6971", "6972", "6973", "6974", "6975", "6976", "6977", "6978", "6979", "6980", "6989", "6990", "6992", "6995", "6996", "6997", "6998", "6999", "7000", "7001", "7002", "7003", "7004", "7005", "7006", "7007", "7008", "7009", "7011", "7012", "7013", "7015", "7016", "7017", "7018", "7019", "7020", "7021", "7022", "7023", "7024", "7025", "7026", "7027", "7028", "7029", "7030", "7031", "7032", "7035", "7036", "7037", "7039", "7041", "7052", "7054", "7055", "7057", "7058", "7060", "7061", "7062", "7063", "7066", "7067", "7068", "7069", "7070", "7071", "7072", "7076", "7089", "7094", "7095", "7121", "7122", "7123", "7124", "7127", "7130", "7169", "7170", "7172", "7174", "7185", "7186", "7189", "7190", "7371", "7372", "7373", "7374", "7375", "9173", "9178", "9182", "9183", "9184", "9185", "9186", "9187", "9188", "9189", "9190", "9191", "9193", "9194", "9195", "9196", "9197", "9200", "9201", "9202", "9203", "9205", "9207", "9208", "9209", "9211", "9212", "9214", "9215", "9216", "9221", "9222", "9224", "9225", "9226", "9227", "9229", "9230", "9246", "9248", "9249", "9251", "9384", "9388", "9389", "9392", "9393", "9394", "9398", "9399", "9409", "9412", "9419", "9427", "9434", "9435", "9436", "9437", "9438", "9439", "9448", "9471", "9472", "9506", "9539", "9541", "9550", "9551", "9555", "9660", "9667", "9705", "9706", "9707", "9708", "9710", "9711", "9712", "9713", "9714", "9715", "9716", "9717", "9718", "9719", "9720", "9721", "9722", "9723", "9724", "9725", "9726", "9728", "9729", "9730", "9731", "9732", "9733", "9734", "9735", "9736", "9737", "9738", "9739", "9740", "9741", "9742", "9743", "9744", "9745", "9747", "9748", "9749", "9751", "9753", "9754", "9755", "9756", "9757", "9758", "9759", "9760", "9761", "9762", "9763", "9764", "9765", "9766", "9767", "9768", "9769", "9770", "9772", "9774", "9775", "9776", "9777", "9778", "9779", "9780", "9781", "9782", "9783", "9789", "9790", "9792", "9793", "9794", "9835", "9848", "9849", "9850", "9851", "9852", "9853", "9854", "9856", "9857", "9858", "9859", "9860", "9861", "9862", "9863", "9864", "9865", "9866", "9867", "9868", "9869", "9870", "9871", "9872", "9873", "9874", "9875", "9876", "9877", "9878", "9879", "9880", "9881", "9882", "9883", "9884", "9885", "9886", "9887", "9888", "9889", "9890", "9891", "9892", "9893", "9894", "9895", "9896", "9897", "9898", "9899", "9900", "9901", "9902", "9903", "9904", "9905", "9906", "9907", "9908", "9909", "9910", "9911", "9912", "9913", "9914", "9915", "9916", "9917", "9918", "9919", "9920", "9921", "9922", "9923", "9924", "9925", "9926", "9927", "9928", "9929", "9930", "9931", "9932", "9933", "9934", "9935", "9936", "9937", "9938", "9939", "9940", "9941", "9942", "9943", "9944", "9945", "9946", "9947", "9948", "9956", "9959", "9961", "9962", "9964", "9971", "9974", "9976", "9985", "9988", "10001", "10010", "10011", "10969", "10971", "10972", "10975", "11069", "11100", "11142", "11143", "11148", "11149", "11150", "11151", "11152", "11153", "11154", "11155", "11156", "11158", "11159", "11160", "11161", "11162", "11163", "11164", "11165", "11166", "11167", "11221", "11222", "11223", "11224", "11225", "11226", "11227", "11228", "11229", "11230", "11231", "11232", "11233", "11234", "11235", "11238", "11301", "11302", "11303", "11304", "11730", "11731", "11760", "11762", "11842", "11843", "11849", "11850", "11852", "11859", "11867", "11875", "11878", "11879", "11880", "11881", "11882", "11883", "11886", "11887", "11888", "11889", "11898", "11922", "11923", "11924", "11925", "11928", "11929", "11930", "11932", "11934", "11935", "11936", "11937", "11949", "11950", "11951", "11953", "11954", "11955", "11957", "11958", "11959", "11960", "11961", "11962", "11963", "11964", "11965", "11966", "11967", "11968", "11969", "11970", "11971", "11972", "11973", "11974", "11975", "11976", "11977", "11978", "11979", "11980", "11981", "11982", "11983", "11984", "11985", "11986", "11987", "11988", "11989", "11991", "11993", "11996", "11998", "11999", "12000", "12001", "12002", "12003", "12004", "12005", "12006", "12008", "12009", "12010", "12011", "12012", "12013", "12014", "12015", "12017", "12018", "12019", "12020", "12021", "12022", "12023", "12024", "12025", "12026", "12027", "12028", "12029", "12030", "12033", "12034", "12035", "12036", "12037", "12038", "12039", "12040", "12041", "12042", "12043", "12044", "12045", "12046", "12047", "12048", "12049", "12050", "12051", "12052", "12053", "12054", "12055", "12056", "12057", "12058", "12059", "12060", "12061", "12062", "12063", "12064", "12066", "12067", "12068", "12069", "12070", "12071", "12072", "12073", "12074", "12076", "12078", "12079", "12080", "12089", "12091", "12092", "12093", "12094", "12095", "12099", "12100", "12101", "12102", "12103", "12104", "12105", "12106", "12107", "12108", "12109", "12110", "12111", "12112", "12113", "12114", "12115", "12116", "12117", "12118", "12119", "12120", "12121", "12122", "12123", "12124", "12125", "12126", "12127", "12128", "12129", "12130", "12131", "12132", "12133", "12134", "12135", "12136", "12137", "12138", "12139", "12140", "12141", "12142", "12143", "12144", "12145", "12146", "12147", "12148", "12149", "12150", "12151", "12152", "12153", "12154", "12155", "12156", "12157", "12158", "12159", "12160", "12161", "12162", "12163", "12164", "12165", "12166", "12167", "12168", "12169", "12170", "12171", "12172", "12173", "12174", "12175", "12176", "12177", "12178", "12179", "12180", "12181", "12182", "12183", "12184", "12185", "12186", "12187", "12188", "12189", "12190", "12191", "12192", "12193", "12194", "12195", "12196", "12197", "12198", "12199", "12200", "12201", "12202", "12203", "12204", "12205", "12206", "12207", "12208", "12209", "12210", "12211", "12212", "12213", "12214", "12215", "12216", "12217", "12219", "12220", "12221", "12222", "12223", "12224", "12225", "12226", "12227", "12228", "12229", "12230", "12231", "12232", "12233", "12234", "12235", "12236", "12237", "12238", "12239", "12240", "12241", "12242", "12243", "12244", "12245", "12246", "12247", "12248", "12249", "12250", "12251", "12252", "12253", "12254", "12255", "12256", "12257", "12258", "12260", "12261", "12273", "12275", "12276", "12277", "12281", "12282", "12286", "12287", "12293", "12296", "12299", "12300", "12301", "12302", "12316", "12317", "12318", "12319", "12320", "12321", "12322", "12341", "12342", "12343", "12344", "12345", "12346", "12347", "12348", "12349", "12350", "12351", "12352", "12353", "12354", "12355", "12356", "12357", "12358", "12359", "12360", "12370", "12372", "12376", "12377", "12380", "12381", "12382", "12383", "12384", "12385", "12386", "12387", "12388", "12389", "12392", "12394", "12395", "12396", "12397", "12398", "12399", "12400", "12401", "12402", "12403", "12404", "12405", "12406", "12407", "12408", "12409", "12415", "12417", "12420", "12421", "12422", "12423", "12424", "12428", "12429", "12471", "12479", "12480", "12514", "12515", "12516", "12517", "12518", "12524", "12525", "12526", "12527", "12529", "12530", "12531", "12532", "12533", "12558", "12560", "12561", "12562", "12565", "12566", "12567", "12568", "12569", "12570", "12571", "12572", "12573", "12574", "12576", "12577", "12578", "12579", "12580", "12582", "12583", "12584", "12585", "12586", "12587", "12588", "12589", "12590", "12591", "12593", "12594", "12595", "12596", "12597", "12599", "12600", "12602", "12603", "12604", "12605", "12606", "12607", "12608", "12609", "12610", "12611", "12612", "12613", "12614", "12615", "12616", "12617", "12618", "12619", "12623", "12624", "12625", "12626", "12627", "12630", "12631", "12632", "12633", "12634", "12635", "12636", "12637", "12638", "12639", "12640", "12641", "12642", "12643", "12644", "12645", "12646", "12647", "12773", "12798", "12799", "12800", "12801", "12802", "12803", "12804", "12805", "12806", "12807", "12808", "12809", "12810", "12811", "12854", "12856", "12857", "12858", "12859", "12861", "13046", "13047", "13050", "13051", "13052", "13053", "13054", "13055", "13056", "13057", "13058", "13059", "13060", "13061", "13062", "13063", "13064", "13065", "13066", "13067", "13068", "13069", "13070", "13071", "13072", "13073", "13074", "13075", "13076", "13077", "13078", "13080", "13081", "13101", "13104", "13109", "13110", "13111", "13112", "13116", "13117", "13118", "13121", "13229", "13231", "13234", "13240", "13241", "13242", "13243", "13244", "13245", "13246", "13247", "13248", "13249", "13250", "13251", "13252", "13253", "13254", "13255", "13256", "13257", "13258", "13259", "13260", "13261", "13262", "13263", "13264", "13265", "13266", "13267", "13269", "13270", "13286", "13287", "13288", "13289", "13290", "13291", "13292", "13293", "13295", "13296", "13297", "13298", "13299", "13302", "13303", "13304", "13305", "13306", "13307", "13308", "13309", "13310", "13311", "13312", "13313", "13323", "13329", "13357", "13435", "13436", "13439", "13502", "13529", "13530", "13531", "13533", "13534", "13535", "13536", "13537", "13538", "13539", "13540", "13541", "13542", "13543", "13544", "13546", "13547", "13548", "13549", "13550", "13551", "13553", "13554", "13555", "13556", "13557", "13558", "13559", "13560", "13561", "13562", "13563", "13564", "13565", "13566", "13567", "13568", "13569", "13570", "13571", "13572", "13573", "13574", "13575", "13576", "13577", "13578", "13579", "13580", "13581", "13582", "13583", "13585", "13586", "13587", "13588", "13590", "13591", "13592", "13593", "13594", "13595", "13596", "13597", "13598", "13601", "13603", "13604", "13657", "13658", "13659", "13660", "13661", "13662", "13664", "13665", "13666", "13667", "13668", "13669", "13670", "13671", "13672", "13673", "13674"*/}
	shards   = []*Shard{&Shard{Host: "10.191.1.100:8080", Cluster: "a"},
		&Shard{Host: "10.191.1.101:8080", Cluster: "a"},
		&Shard{Host: "10.191.1.102:8080", Cluster: "a"},
		&Shard{Host: "10.191.1.103:8080", Cluster: "a"},
		&Shard{Host: "10.191.1.104:8080", Cluster: "a"},
		&Shard{Host: "172.18.11.15:8080", Cluster: "a"},
		&Shard{Host: "172.18.11.16:8080", Cluster: "a"}}
)

type Shard struct {
	Host    string
	Cluster string
}

func main() {
	ring := NewRing(NODE_NUM)

	for _, shard := range shards {
		ring.AddNode(shard.Host, 1)
	}

	ring.Bake()

	// // 统计每台机器是否均匀
	// m := make(map[string]int)
	// for _, catId := range CatIds {
	// 	m[ring.Hash(catId)]++
	// }

	// for _, shard := range shards {
	// 	fmt.Println(shard.Host, m[shard.Host])
	// }

	for _, catId := range CatIds {
		log.Println(catId, ring.Hash(catId))
	}

	log.Println("========================")
	ring.AddNode("10.191.1.99:8080", 1)
	ring.Bake()

	for _, catId := range CatIds {
		log.Println(catId, ring.Hash(catId))
	}

}

type node struct {
	node string // node name
	hash uint
}

type tickArray []node

func (p tickArray) Len() int           { return len(p) }
func (p tickArray) Less(i, j int) bool { return p[i].hash < p[j].hash }
func (p tickArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p tickArray) Sort()              { sort.Sort(p) }

type hashRing struct {
	defaultSpots int
	ticks        tickArray
	length       int
}

func NewRing(n int) (h *hashRing) {
	h = new(hashRing)
	h.defaultSpots = n
	return
}

// Adds a new node to a hash ring
// n: name of the server
// s: multiplier for default number of ticks (useful when one cache node has more resources, like RAM, than another)
func (h *hashRing) AddNode(n string, s int) {
	tSpots := h.defaultSpots * s
	hash := sha1.New()
	for i := 1; i <= tSpots; i++ {
		hash.Write([]byte(n + ":" + strconv.Itoa(i)))
		hashBytes := hash.Sum(nil)

		n := &node{
			node: n,
			hash: uint(hashBytes[19]) | uint(hashBytes[18])<<8 | uint(hashBytes[17])<<16 | uint(hashBytes[16])<<24,
		}

		h.ticks = append(h.ticks, *n)
		hash.Reset()
	}
}

func (h *hashRing) Bake() {
	h.ticks.Sort()
	h.length = len(h.ticks)
}

func (h *hashRing) Hash(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)
	v := uint(hashBytes[19]) | uint(hashBytes[18])<<8 | uint(hashBytes[17])<<16 | uint(hashBytes[16])<<24
	i := sort.Search(h.length, func(i int) bool { return h.ticks[i].hash >= v })

	if i == h.length {
		i = 0
	}

	return h.ticks[i].node
}
