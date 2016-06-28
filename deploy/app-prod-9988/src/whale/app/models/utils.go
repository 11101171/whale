//****************************************************/
// Addres: https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package models

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"

	"github.com/revel/revel"
)

// There is util list
var desUtil *DESUtil = &DESUtil{}

type DESUtil struct{}

// var str string = desUtil.Encrypt("nihao")
// fmt.Println("str", str)
// var str2 string = desUtil.Decrypt(str)
// fmt.Println("str2", str2)

// var key []byte = []byte("z0oI3d94")
var key []byte = []byte("!o#i_dcws294el#unvmx6ze[")

// 加密
// return string
func (du *DESUtil) Encrypt(data string) string {
	result, err := du.TripleDesEncrypt([]byte(data), key)
	if err != nil {
		revel.ERROR.Println("加密失败")
		return ""
	} else {
		return base64.StdEncoding.EncodeToString(result)
	}
}

// 解密
// return string
func (du *DESUtil) Decrypt(data string) string {
	dataStr, _ := base64.StdEncoding.DecodeString(data)
	origData, err := du.TripleDesDecrypt(dataStr, key)
	if err != nil {
		revel.ERROR.Println("解密失败")
		return ""
	} else {
		return string(origData)
	}
}

func (du *DESUtil) DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = du.PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (du *DESUtil) DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

// 3DES加密
func (du *DESUtil) TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = du.PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func (du *DESUtil) TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func (du *DESUtil) ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func (du *DESUtil) ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func (du *DESUtil) PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
