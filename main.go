package main

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"os"
	"strings"
	"unsafe"

	aw "github.com/deanishe/awgo"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
)

var algorithms = map[string]bool{
	"md4":       true,
	"md5":       true,
	"sha1":      true,
	"sha224":    true,
	"sha256":    true,
	"sha384":    true,
	"sha512":    true,
	"ripemd160": true,
	"base64":    true,
	"crc32":     true,
}

var wf *aw.Workflow

func main() {
	l := len(os.Args)
	if l == 1 {
		return
	}

	algorithm := ""
	arg := ""

	if l == 2 {
		arg = strings.Join(os.Args[1:], " ")
	} else if l > 2 {
		algorithm = os.Args[1]
		if _, ok := algorithms[algorithm]; !ok && !StringPrefixInArray(algorithm) {
			// if algorithm is not in the list of supported algorithms
			algorithm = ""
			arg = strings.Join(os.Args[1:], " ")
		} else {
			// if algorithm is in the list of supported algorithms
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
		Item{Name: "ripemd160", Value: calcRIPEMD160(arg)},
		Item{Name: "base64", Value: calcBase64(arg)},
		Item{Name: "crc32", Value: calcCRC32(arg)},
	)

	opt := aw.MaxResults(len(items))
	wf = aw.New(opt)

	fn := func() {
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
	wf.Run(fn)
}

func calcMD4(in string) string {
	hash := crypto.MD4.New()
	hash.Write([]byte(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcMD5(in string) string {
	hash := crypto.MD5.New()
	hash.Write(StringToByteSlice(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcSHA1(in string) string {
	hash := crypto.SHA1.New()
	hash.Write(StringToByteSlice(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcSHA224(in string) string {
	hash := crypto.SHA224.New()
	hash.Write(StringToByteSlice(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcSHA256(in string) string {
	hash := crypto.SHA256.New()
	hash.Write(StringToByteSlice(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcSHA384(in string) string {
	hash := crypto.SHA384.New()
	hash.Write(StringToByteSlice(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcSHA512(in string) string {
	hash := crypto.SHA512.New()
	hash.Write(StringToByteSlice(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcRIPEMD160(in string) string {
	hash := crypto.RIPEMD160.New()
	hash.Write(StringToByteSlice(in))
	data := hash.Sum(nil)
	return fmt.Sprintf("%x", data)
}

func calcBase64(in string) string {
	return base64.StdEncoding.EncodeToString(StringToByteSlice(in))
}

func calcCRC32(in string) string {
	return fmt.Sprintf("%x", crc32.ChecksumIEEE(StringToByteSlice(in)))
}

// StringToByteSlice converts a string to a byte slice.
func StringToByteSlice(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// StringPrefixInArray checks if a string is a prefix of any string in the array.
func StringPrefixInArray(prefix string) bool {
	for k := range algorithms {
		if strings.HasPrefix(k, prefix) {
			return true
		}
	}
	return false
}
