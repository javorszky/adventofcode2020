package day21

var foodcounter = map[string]int{
	"bcvmz":    30,
	"bjvcg":    30,
	"bkd":      27,
	"bnzq":     24,
	"bxv":      32,
	"cbbkfx":   27,
	"ckqq":     79,
	"clqk":     33,
	"cmnb":     37,
	"cnghsg":   36,
	"crnfzr":   23,
	"csfmx":    29,
	"cxfzhj":   13,
	"cxzkmr":   29,
	"dbx":      29,
	"dcbk":     26,
	"ddbmq":    30,
	"ddgdhg":   17,
	"ddrd":     14,
	"dfrg":     17,
	"dgrrm":    28,
	"djpsmd":   30,
	"dln":      38,
	"dmxhhd":   19,
	"dnpgcd":   28,
	"dpfphd":   30,
	"dsmc":     34,
	"dtlhh":    35,
	"dvcfx":    25,
	"dznd":     26,
	"fbcds":    29,
	"fdf":      25,
	"fhtsl":    25,
	"fjgxv":    22,
	"fkh":      36,
	"flhfddq":  34,
	"fmvvb":    24,
	"frld":     15,
	"fstgc":    25,
	"ghr":      20,
	"gjqfj":    27,
	"glrc":     26,
	"gmc":      40,
	"gsgmdn":   25,
	"gtgrcf":   27,
	"gzgvdh":   34,
	"hkqp":     30,
	"hkxcb":    28,
	"hlmvqh":   74,
	"hnmjfl":   33,
	"hrfmdk":   29,
	"htllnq":   15,
	"jdtzr":    22,
	"jfqqgtl":  46,
	"jgtrnm":   23,
	"jkbqk":    21,
	"jmdg":     73,
	"jmgt":     32,
	"jngghk":   17,
	"jnr":      25,
	"jvvmcq":   15,
	"jxjqp":    23,
	"kbtx":     21,
	"kfsfn":    18,
	"kglr":     30,
	"klmjmz":   30,
	"kltcm":    27,
	"klx":      33,
	"kmsh":     23,
	"knnmm":    27,
	"knqzf":    27,
	"kqzcj":    21,
	"kx":       20,
	"lbfgp":    27,
	"lbskg":    23,
	"lh":       31,
	"ljhn":     30,
	"ljmvpp":   23,
	"llgsg":    16,
	"lqmfgp":   27,
	"ltvr":     20,
	"lvjpp":    25,
	"lxr":      29,
	"mgxzl":    23,
	"mhrm":     33,
	"mjpt":     23,
	"mnvq":     33,
	"mppf":     29,
	"mrfxh":    74,
	"mrxg":     24,
	"mstc":     26,
	"mszc":     32,
	"mvnqdh":   21,
	"ncqdr":    28,
	"nfnzx":    23,
	"nggcjr":   31,
	"ngqh":     33,
	"nldzpc":   30,
	"npbnj":    12,
	"nqfc":     23,
	"nthhxn":   32,
	"nzlks":    16,
	"pbn":      28,
	"pcf":      15,
	"pfqxsxd":  31,
	"pgb":      28,
	"pgqxp":    29,
	"pjln":     37,
	"pjqdmpm":  34,
	"pmtfmv":   27,
	"pmvl":     24,
	"pnglkx":   25,
	"pptgt":    28,
	"pzxj":     26,
	"qchnn":    31,
	"qdpbx":    34,
	"qfkjsq":   17,
	"qfvfzg":   26,
	"qjsbk":    27,
	"qkgqv":    22,
	"qmmt":     35,
	"qmthj":    26,
	"qrftr":    35,
	"qrsczjv":  72,
	"qsgb":     17,
	"qspfqb":   20,
	"qszmzsh":  25,
	"rbvdt":    35,
	"rcr":      26,
	"rcvd":     28,
	"rczbvg":   32,
	"rfh":      26,
	"rgdx":     19,
	"rhrc":     26,
	"rknm":     27,
	"rkzhxmh":  24,
	"rpcdfph":  28,
	"rpmmv":    43,
	"rpmqq":    22,
	"rqqfnlc":  30,
	"rrbndl":   19,
	"rvchbn":   21,
	"rxr":      24,
	"rxrgtvs":  25,
	"sdccxkt":  30,
	"sfk":      23,
	"sjgx":     20,
	"srgnx":    21,
	"stcsp":    16,
	"szsbj":    22,
	"tbm":      27,
	"tfmgl":    20,
	"thcs":     20,
	"thvm":     80,
	"tjsdp":    24,
	"tphtz":    17,
	"tshn":     32,
	"ttxx":     31,
	"tvqbhv":   24,
	"tzjnvp":   25,
	"vbqbkt":   19,
	"vfkpj":    27,
	"vgjbgj":   23,
	"vgp":      32,
	"vhcnpg":   22,
	"vhqfz":    17,
	"vht":      26,
	"vnfmc":    30,
	"vnmfg":    24,
	"vpgvm":    21,
	"vpvj":     30,
	"vqrjn":    23,
	"vtljml":   23,
	"xbpb":     21,
	"xddkbd":   35,
	"xgbvk":    28,
	"xhmmt":    33,
	"xjdhk":    23,
	"xjrpr":    33,
	"xjzc":     34,
	"xlnn":     36,
	"xlzqfb":   32,
	"xmjlsn":   23,
	"xnfhq":    25,
	"xxfgvz":   42,
	"xzcdnr":   34,
	"zdntns":   17,
	"zfcvnj":   28,
	"zfml":     20,
	"zfsks":    26,
	"zhghprj":  34,
	"zkzdrn":   17,
	"zlgztsbd": 20,
	"zmb":      75,
	"zmcj":     20,
	"zntrfp":   30,
	"zqsc":     27,
	"zrgzf":    75,
	"zttx":     23,
	"zzldmh":   26,
}
