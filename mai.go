包装主

进口 (
	“encoding/json”
	“旗帜”
	“fmt”
	“数学/兰特”
	“os”
	“strconv”
	“时间”

	“github.com/FloatTech/ZeroBot-Plugin/kanban” // 在最前打印 banner

	 ---------以下插件均可通过前面加 // 注释，注释后停用并不加载插件--------- //
	 ----------------------插件优先级按顺序从高到低---------------------- //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	 ----------------------------高优先级区---------------------------- //
	vvvvvvvvv高优先级区vvvv
	vvvvvvvvvv高优先级区vvvvvvvv
	vvvvvvv高优先级区vvvvv //
	vvvvvvvvvvv //
	vvvv //

	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/chat” // 基础词库

	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/sleep_manage” // 统计睡眠时间

	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/atri” // ATRI词库

	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/manager” // 群管

	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/thesaurus” // 词典匹配回复

	_ “github.com/FloatTech/zbputils/job” // 定时指令触发器

	//                               ^^^^                               //
	//                          ^^^^^^^^^^^^^^                          //
	 ^^^^^^^高优先级区^^^^^^^ //
	 ^^^^^^^^^^^^^^高优先级区^^^^^^^^^^^^^^ //
	 ^^^^^^^^^^^^^^^^^^^^^^^^^^^^高优先级区^^^^^^^^^^^^^^^^^^^^^^^^^^^^ //
	 ----------------------------高优先级区---------------------------- //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	 ----------------------------中优先级区---------------------------- //
	vvvvvvvvvv
	vvvvvvvvvvv中优先级区vvvvvvvvv
	vvvvvvv中优先级区vvvvvvv //
	vvvvvvvvvvv //
	vvvv //

	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/ai_false” // 服务器监控
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/aiwife” // 随机老婆
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/b14” // base16384加解密
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/baidu” // 百度一下
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/bilibili” // 查询b站用户信息
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/bilibili_parse” // b站视频链接解析
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/book_review” // 哀伤雪刃吧推书记录
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/cangtoushi” // 藏头诗
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/choose” // 选择困难症帮手
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/chouxianghua” // 说抽象话
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/coser” // 三次元小姐姐
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/cpstory” // cp短打
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/danbooru” // DeepDanbooru二次元图标签识别
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/diana” // 嘉心糖发病
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/drift_bottle” // 漂流瓶
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/emojimix” // 合成表情符号
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/epidemic” // 城市疫情查询
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/font” // 渲染任意文字到图片
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/fortune” // 运势
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/funny” // 笑话
	_ “github.com/FloatTech/ZeroBot-Plugin/plugin/genshin” // 原神抽卡
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/gif"            // 制图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/github"         // 搜索GitHub仓库
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/guessmusic"     // 猜歌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/hs"             // 炉石
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/hyaku"          // 百人一首
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/image_finder"   // 关键字搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/inject"         // 注入指令
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/jandan"         // 煎蛋网无聊图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/juejuezi"       // 绝绝子生成器
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/lolicon"        // lolicon 随机图片
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/midicreate"     // 简易midi音乐制作
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/moyu"           // 摸鱼
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/moyu_calendar"  // 摸鱼人日历
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/music"          // 点歌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nativesetu"     // 本地涩图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nativewife"     // 本地老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nbnhhsh"        // 拼音首字母缩写释义工具
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/novel"          // 铅笔小说网搜索
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nsfw"           // nsfw图片识别
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/omikuji"        // 浅草寺求签
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/qqwife"         // 一群一天一夫一妻制群老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/reborn"         // 投胎
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/runcode"        // 在线运行代码
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/saucenao"       // 以图搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/scale"          // 叔叔的AI二次元图片放大
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/score"          // 分数
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/setutime"       // 来份涩图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/shadiao"        // 沙雕app
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/shindan"        // 测定
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/tarot"          // 抽塔罗牌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/tiangou"        // 舔狗日记
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/tracemoe"       // 搜番
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/translation"    // 翻译
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/vtb_quotation"  // vtb语录
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wangyiyun"      // 网易云音乐热评
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/word_count"     // 聊天热词
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wordle"         // 猜单词
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ymgal"          // 月幕galgame

	// _ "github.com/FloatTech/ZeroBot-Plugin/plugin/wtf"            // 鬼东西
	// _ "github.com/FloatTech/ZeroBot-Plugin/plugin/bilibili_push"  // b站推送

	//                               ^^^^                               //
	//                          ^^^^^^^^^^^^^^                          //
	//                      ^^^^^^^中优先级区^^^^^^^                      //
	//               ^^^^^^^^^^^^^^中优先级区^^^^^^^^^^^^^^               //
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^中优先级区^^^^^^^^^^^^^^^^^^^^^^^^^^^^ //
	// ----------------------------中优先级区---------------------------- //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	// ----------------------------低优先级区---------------------------- //
	// vvvvvvvvvvvvvvvvvvvvvvvvvvvv低优先级区vvvvvvvvvvvvvvvvvvvvvvvvvvvv //
	//               vvvvvvvvvvvvvv低优先级区vvvvvvvvvvvvvv               //
	//                      vvvvvvv低优先级区vvvvvvv                      //
	//                          vvvvvvvvvvvvvv                          //
	//                               vvvv                               //

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/curse" // 骂人

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ai_reply" // 人工智能回复

	//                               ^^^^                               //
	//                          ^^^^^^^^^^^^^^                          //
	//                      ^^^^^^^低优先级区^^^^^^^                      //
	//               ^^^^^^^^^^^^^^低优先级区^^^^^^^^^^^^^^               //
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^低优先级区^^^^^^^^^^^^^^^^^^^^^^^^^^^^ //
	// ----------------------------低优先级区---------------------------- //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	// -----------------------以下为内置依赖，勿动------------------------ //
	"github.com/FloatTech/zbputils/process"
	"github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"
	// -----------------------以上为内置依赖，勿动------------------------ //
)

func init() {
	sus := make([]int64, 0, 16)
	// 解析命令行参数
	d := flag.Bool("d", false, "Enable debug level log and higher.")
	w := flag.Bool("w", false, "Enable warning level log and higher.")
	h := flag.Bool("h", false, "Display this help.")
	// 直接写死 AccessToken 时，请更改下面第二个参数
	token := flag.String("t", "", "Set AccessToken of WSClient.")
	// 直接写死 URL 时，请更改下面第二个参数
	url := flag.String("u", "ws://127.0.0.1:6700", "Set Url of WSClient.")
	// 默认昵称
	adana := flag.String("n", "椛椛", "Set default nickname.")
	prefix := flag.String("p", "/", "Set command prefix.")
	runcfg := flag.String("c", "", "Run from config file.")
	save := flag.String("s", "", "Save default config to file and exit.")

	flag.Parse()

	if *h {
		kanban.PrintBanner()
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	} else {
		if *d && !*w {
			logrus.SetLevel(logrus.DebugLevel)
		}
		if *w {
			logrus.SetLevel(logrus.WarnLevel)
		}
	}

	for _, s := range flag.Args() {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			continue
		}
		sus = append(sus, i)
	}

	// 通过代码写死的方式添加主人账号
	// sus = append(sus, 12345678)
	// sus = append(sus, 87654321)

	if *runcfg != "" {
		f, err := os.Open(*runcfg)
		if err != nil {
			panic(err)
		}
		config.W = make([]*driver.WSClient, 0, 2)
		err = json.NewDecoder(f).Decode(&config)
		f.Close()
		if err != nil {
			panic(err)
		}
		config.Z.Driver = make([]zero.Driver, len(config.W))
		for i, w := range config.W {
			config.Z.Driver[i] = w
		}
		logrus.Infoln（“[main] 从”， *runcfg， “读取配置文件”)
		返回
	}

	配置。W = []*驱动程序。WSClient{driver.NewWebSocketClient（*url， *token)}
	配置。Z = 零。配置{
		昵称： append（[]string{*adana}， “ATRI”， “atri”， “亚托莉”， “アトリ”），
		命令前缀： *前缀，
		超级用户： sus，
		驱动程序：[]零。驱动程序{配置。W[0]}，
	}

	如果 *保存 ！= "" {
		f， err ：= os.创建（*保存)
		如果错误 ！= 无 {
			panic（err）)
		}
		err = json.NewEncoder（f）.Encode（&config)
		六.关闭()
		如果错误 ！= 无 {
			panic（err）)
		}
		logrus.Infoln（“[main] 配置文件已保存到”， *保存)
		os.出口（0)
	}
}

功能主要() {
	兰特。种子（时间。Now（）.UnixNano（）） // 全局 seed，其他插件无需再 seed
	 帮助
	零。OnFullMatchGroup（[]string{“/help”， “.help”， “菜单”}， zero.OnlyToMe）。SetBlock（true）.
		手柄（func（ctx *zero.断续器) {
			腾讯网.发送链（消息。文本（看板。Banner， “\n可发送\”/服务列表\“查看 bot 功能”))
		})
	零。OnFullMatch（“查看zbp公告”， zero.OnlyToMe，零。管理员权限）。SetBlock（true）.
		手柄（func（ctx *zero.断续器) {
			腾讯网.发送链（消息。文本（看板。看板()))
		})
	零。RunAndBlock（config.Z，过程。GlobalInitMutex.开锁)
}
