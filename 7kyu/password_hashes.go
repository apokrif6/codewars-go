//When you sign up for an account somewhere, some websites do not actually store your password in their databases. Instead, they will transform your password into something else using a cryptographic hashing algorithm.

//After the password is transformed, it is then called a password hash. Whenever you try to login, the website will transform the password you tried using the same hashing algorithm and simply see if the password hashes are the same.

//Create the function that converts a given string into an md5 hash. The return value should be encoded in hexadecimal.
//Code Examples

//passHash("password") // --> "5f4dcc3b5aa765d61d8327deb882cf99"
//passHash("abc123") // --> "e99a18c428cb38d5f260853678922e03"


package kata

import (
    "crypto/md5"
    "encoding/hex"
)

func PassHash(str string) string {
  hash := md5.Sum([]byte(str))
  
  return hex.EncodeToString(hash[:])
}
