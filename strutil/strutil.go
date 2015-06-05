package strutil

import (
  "strconv"
)

func ParseUint(str string, defaultVal uint64, base int, bitSize int) uint64{
  i, err := strconv.ParseUint(str, base, bitSize)
  if err != nil{
    i = defaultVal
  }
  return i;
}
func ParseInt(str string, defaultVal int64, base int, bitSize int) int64{
  i, err := strconv.ParseInt(str, base, bitSize)
  if err != nil{
    i = defaultVal
  }
  return i;
}

func ToUint64(str string, defaultVal uint64) uint64{
  return ParseUint(str, defaultVal, 10, 64)
}
func ToInt64(str string, defaultVal int64) int64{
  return ParseInt(str, defaultVal, 10, 64)
}

func ToUint32(str string, defaultVal uint32) uint32{
  return uint32(ParseUint(str, uint64(defaultVal), 10, 32))
}
func ToInt32(str string, defaultVal int32) int32{
  return int32(ParseInt(str, int64(defaultVal), 10, 32))
}
