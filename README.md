# go-tools

## 介绍

本人自用go工具集合

### crypto

#### rsa

* 使用demo

```go
package rsa

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/crypto/rsa"
	"github.com/jianlu8023/go-tools/pkg/encoding/base64"
)

// TestRsaEncryptDecrypt
// @Description: TestRsaEncryptDecrypt
// @author ght
// @date 2023-10-07 17:02:45
// @param t:
func TestRsaEncryptDecrypt(t *testing.T) {
	data := []byte("hello world")
	fmt.Println(data)
	encrypt := rsa.RsaEncrypt(data)
	fmt.Println(encrypt)
	encode := base64.Base64Encode(encrypt)
	fmt.Println(encode)
	decrypt := rsa.RsaDecrypt(encrypt)
	fmt.Println(decrypt)

}


```

#### md5

使用demo

```go
package md5

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/crypto/md5"
)

func TestMd5(t *testing.T) {

	str := md5.MD5("hello world")
	fmt.Println(str)
}
```

### encoding

#### base64

使用demo

```go
package base64

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/encoding/base64"
)

// TestBase64
// @Description: TestBase64
// @author ght
// @date 2023-10-07 17:08:51
// @param t:
func TestBase64(t *testing.T) {
	encode := base64.ByteToBase64([]byte("hello world"))
	fmt.Println(encode)
	decode, err := base64.Base64ToByte(encode)
	if err != nil {
		fmt.Println(fmt.Sprintf("err:%s", err))
	}
	fmt.Println(decode)
}


```

#### base62

使用demo

```go
package base62

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/encoding/base62"
)

// TestBase62
// @Description: TestBase62
// @author ght
// @date 2023-10-07 17:15:04
// @param t:
func TestBase62(t *testing.T) {
	toBase62 := base62.IntToBase62(10)
	fmt.Println(toBase62)
}
```

### format

#### json

使用demo

```go
package json

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/format/json"
)

type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// TestPrettyJSON
// @Description: TestPrettyJSON
// @author ght
// @date 2023-10-07 17:15:59
// @param t:
func TestPrettyJSON(t *testing.T) {
	animals := []Animal{
		{Name: "兔子", Age: 10},
		{Name: "猫", Age: 11},
	}
	prettyJSON := json.PrettyJSON(animals)
	fmt.Println(prettyJSON)
}

```

### replace

使用demo

```go
package replace

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/replace"
)

// TestReplace
// @Description: TestReplace
// @author ght
// @date 2023-10-07 17:22:06
// @param t:
func TestReplace(t *testing.T) {
	template := "https://api.github.com/repo/{{owner}}/{{repo}}"

	str := replace.Replace(template, map[string]string{
		"{{owner}}": "ceshi",
		"{{repo}}":  "ceshi",
	})
	fmt.Println(template)
	fmt.Println(str)
}

```

### uuid

使用demo

```go
package uuid

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/uuid"
)

// TestUUID
// @Description: TestUUID
// @author ght
// @date 2023-10-07 17:19:36
// @param t:
func TestUUID(t *testing.T) {
	genUUID := uuid.GenUUID()
	fmt.Println(genUUID)
}

```