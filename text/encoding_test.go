package text

import (
  "fmt"
  "github.com/stretchr/testify/assert"
  "golang.org/x/text/encoding/japanese"
  "testing"
)

func TestEncodingString(t *testing.T) {
  fmt.Println(EncodingString("テスト", japanese.EUCJP))
  fmt.Println(EncodingString("テスト", japanese.ISO2022JP))
  fmt.Println(EncodingString("テスト", japanese.ShiftJIS))
  assert.Equal(t, []int32{65533, 421, 65533, 65533, 65533},                []rune(EncodingString("テスト", japanese.EUCJP)))
  assert.Equal(t, []int32{27, 36, 66, 37, 70, 37, 57, 37, 72, 27, 40, 66}, []rune(EncodingString("テスト", japanese.ISO2022JP)))
  assert.Equal(t, []int32{65533, 101, 65533, 88, 65533, 103},              []rune(EncodingString("テスト", japanese.ShiftJIS)))
}
