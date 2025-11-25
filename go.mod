module github.com/jianlu8023/go-tools/v2

go 1.22

require (
	// chainmaker.org/chainmaker/pb-go/v2 v2.3.0
	// chainmaker.org/chainmaker/sdk-go/v2 v2.3.0
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/aryann/difflib v0.0.0-20210328193216-ff5ff6dc229b
	github.com/bytedance/sonic v1.14.1
	github.com/cespare/xxhash/v2 v2.1.2
	github.com/cloudwego/base64x v0.1.6
	github.com/elgs/gojq v0.0.0-20230628214826-df5c4045598e
	github.com/facebookgo/atomicfile v0.0.0-20151019160806-2de1f203e7d5
	github.com/fogleman/gg v1.3.0
	github.com/gin-gonic/gin v1.10.1
	github.com/go-ego/gse v0.80.3
	github.com/go-resty/resty/v2 v2.13.1
	github.com/gocarina/gocsv v0.0.0-20240520201108-78e41c74b4b1
	github.com/google/uuid v1.6.0
	github.com/json-iterator/go v1.1.12
	github.com/mattn/go-colorable v0.1.13
	github.com/mattn/go-isatty v0.0.20
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/schollz/progressbar/v3 v3.18.0
	github.com/shirou/gopsutil/v4 v4.24.10
	github.com/stretchr/testify v1.10.0
	github.com/tjfoc/gmsm v1.4.1
	golang.org/x/crypto v0.33.0
	golang.org/x/sys v0.30.0
)

require (
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic/loader v0.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ebitengine/purego v0.8.1 // indirect
	github.com/elgs/gosplitargs v0.0.0-20230310130726-7d16e488436a // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	github.com/vcaesar/cedar v0.20.2 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/image v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/term v0.29.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	// chainmaker.org/chainmaker/pb-go/v2 => chainmaker.org/chainmaker/pb-go/v2 v2.3.0
	// chainmaker.org/chainmaker/sdk-go/v2 => chainmaker.org/chainmaker/sdk-go/v2 v2.3.0
	github.com/araddon/dateparse => github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/aryann/difflib => github.com/aryann/difflib v0.0.0-20210328193216-ff5ff6dc229b
	github.com/bytedance/sonic => github.com/bytedance/sonic v1.14.1
	github.com/cespare/xxhash/v2 => github.com/cespare/xxhash/v2 v2.3.0
	github.com/cloudwego/base64x => github.com/cloudwego/base64x v0.1.6
	github.com/elgs/gojq => github.com/elgs/gojq v0.0.0-20230628214826-df5c4045598e
	github.com/facebookgo/atomicfile => github.com/facebookgo/atomicfile v0.0.0-20151019160806-2de1f203e7d5
	github.com/fogleman/gg => github.com/fogleman/gg v1.3.0
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.10.1
	github.com/go-ego/gse => github.com/go-ego/gse v0.80.3
	github.com/go-resty/resty/v2 v2.13.1 => github.com/go-resty/resty/v2 v2.16.5
	github.com/gocarina/gocsv => github.com/gocarina/gocsv v0.0.0-20240520201108-78e41c74b4b1
	github.com/json-iterator/go => github.com/json-iterator/go v1.1.12
	github.com/mattn/go-colorable => github.com/mattn/go-colorable v0.1.14
	github.com/mitchellh/go-homedir => github.com/mitchellh/go-homedir v1.1.0
	github.com/shirou/gopsutil/v4 => github.com/shirou/gopsutil/v4 v4.24.10
	github.com/stretchr/testify v1.10.0 => github.com/stretchr/testify v1.11.1
	github.com/tjfoc/gmsm => github.com/tjfoc/gmsm v1.4.1
	golang.org/x/arch => golang.org/x/arch v0.14.0
	golang.org/x/crypto => golang.org/x/crypto v0.33.0
	golang.org/x/image => golang.org/x/image v0.24.0
	golang.org/x/net => golang.org/x/net v0.35.0
	golang.org/x/sys => golang.org/x/sys v0.30.0
	golang.org/x/text => golang.org/x/text v0.22.0
)
