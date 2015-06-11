package toolkit

import (
  "encoding/hex"
  "crypto/sha256"
  "crypto/sha512"
  "crypto/sha1"
)
//bytes
func Sha512b( b []byte ) string {
  h := sha512.New()
  h.Write( b  )
  return hex.EncodeToString(h.Sum(nil))
}
func Sha1b( b []byte ) string {
  h := sha1.New()
  h.Write( b  )
  return hex.EncodeToString(h.Sum(nil))
}
func Sha256b( b []byte ) string {
  h := sha256.New()
  h.Write( b  )
  return hex.EncodeToString(h.Sum(nil))
}

//strings
func Sha256( s string ) string {
  h := sha256.New()
  h.Write( []byte(s) )
  return hex.EncodeToString(h.Sum(nil))
}
func Sha512( s string ) string {
  h := sha512.New()
  h.Write( []byte(s) )
  return hex.EncodeToString(h.Sum(nil))
}
func Sha1( s string ) string {
  h := sha1.New()
  h.Write( []byte(s) )
  return hex.EncodeToString(h.Sum(nil))
}

