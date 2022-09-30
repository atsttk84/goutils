package crypto

import (
	"reflect"
	"unsafe"
)

const (
	TAG_NAME    = "crypto"
	ALG_AES_CBC = "aes_cbc"
	ALG_SHA256  = "sha256"
)

func Encrypt(data interface{}) {
	if data == nil {
		return
	}
	t := reflect.TypeOf(reflect.ValueOf(data).Elem().Interface())
	v := reflect.ValueOf(data).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := f.Name
		pv := v.FieldByName(name)
		ps := (*string)(unsafe.Pointer(pv.UnsafeAddr()))
		switch f.Tag.Get(TAG_NAME) {
		case ALG_AES_CBC:
			*ps = EncryptAesCbc(*ps)
		case ALG_SHA256:
			*ps = SumSha256(*ps)
		}
	}
}

func Decrypt(data interface{}) {
	if data == nil {
		return
	}
	t := reflect.TypeOf(reflect.ValueOf(data).Elem().Interface())
	v := reflect.ValueOf(data).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name := f.Name
		pv := v.FieldByName(name)
		ps := (*string)(unsafe.Pointer(pv.UnsafeAddr()))
		switch f.Tag.Get(TAG_NAME) {
		case ALG_AES_CBC:
			*ps = DecryptAesCbc(*ps)
		}
	}
}
