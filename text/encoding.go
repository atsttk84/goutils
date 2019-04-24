package text

import (
  "golang.org/x/text/encoding"
  "golang.org/x/text/transform"
  "log"
)

/*
  str: utf8 string
  enc: encoding var
    e.g.
      japanese.EUCJP
      japanese.ISO2022JP
      japanese.ShiftJIS
    cf. https://godoc.org/golang.org/x/text/encoding/japanese
*/
func EncodingString(str string, enc encoding.Encoding) string {
  result, _, err := transform.String(enc.NewEncoder(), str)
  if err != nil {
    log.Printf("Encoding error. str: %s, enc: %s", str, enc)
    return ""
  }
  return result
}
