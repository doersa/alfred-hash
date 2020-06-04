package main

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"os"
	"strings"

	aw "github.com/deanishe/awgo"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
)

func main() {
	l := len(os.Args)
	if l == 1 {
		return
	}

	algorithms := make(map[string]bool, 0)
	algorithms["md"] = true
	algorithms["sha"] = true
	algorithms["base"] = true
	algorithms["crc"] = true
	algorithms["ripemd"] = true

	algorithm := ""
	arg := ""

	arg = strings.Join(os.Args[1:]," ")

	if l > 2 {
		algorithm = os.Args[1]
		if _, ok := algorithms[algorithm]; !ok {
			arg = strings.Join(os.Args[1:], " ")
		} else {
			arg = strings.Join(os.Args[2:], " ")
		}
	}

	type Item struct {
		Name  string
		Value string
	}

	items := make([]Item, 0)
	items = append(items,
		Item{Name: "md4", Value: calcMD4(arg)},
		Item{Name: "md5", Value: calcMD5(arg)},
		Item{Name: "sha1", Value: calcSHA1(arg)},
		Item{Name: "sha224", Value: calcSHA224(arg)},
		Item{Name: "sha256", Value: calcSHA256(arg)},
		Item{Name: "sha384", Value: calcSHA384(arg)},
		Item{Name: "sha512", Value: calcSHA512(arg)},
		Item{Name: "sha512", Value: calcSHA512(arg)},
		Item{Name: "ripemd160", Value: calcRIPEMD160(arg)},
		Item{Name: "base64", Value: calcBase64(arg)},
		Item{Name: "crc32", Value: calcCRC32(arg)},
	)

	opt := aw.MaxResults(len(items))
	wf := aw.New(opt)

	for _, v := range items {
		if algorithm != "" && !strings.HasPrefix(v.Name, algorithm) {
			continue
		}
		item := wf.NewItem(v.Name)
		item.Arg(v.Value)
		item.Subtitle(v.Value)
		item.Valid(true)
	}

	wf.SendFeedback()
}

func calcMD4(in string) string {
	hash := crypto.MD4.New()
	hash.Write([]byte(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcMD5(in string) string {
	hash := crypto.MD5.New()
	hash.Write([]byte(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcSHA1(in string) string {
	hash := crypto.SHA1.New()
	hash.Write([]byte(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcSHA224(in string) string {
	hash := crypto.SHA224.New()
	hash.Write([]byte(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcSHA256(in string) string {
	hash := crypto.SHA256.New()
	hash.Write([]byte(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcSHA384(in string) string {
	hash := crypto.SHA384.New()
	hash.Write([]byte(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcSHA512(in string) string {
	hash := crypto.SHA512.New()
	hash.Write([]byte(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcRIPEMD160(in string) string {
	hash := crypto.RIPEMD160.New()
	hash.Write([]byte(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcBase64(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}

func calcCRC32(in string) string {
	return fmt.Sprintf("%x", crc32.ChecksumIEEE([]byte(in)))
}
