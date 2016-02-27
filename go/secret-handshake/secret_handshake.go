package secret

var code = []string {
  "wink",
  "double blink",
  "close your eyes",
  "jump",
}

func Handshake(n int) []string {
  if n <= 0 {
    return nil
  }

  var ans []string
  reverse := false
  if n & 16 > 0 {
    reverse = true
  }

  for i := 0; i < 4; i++ {
    if n & (1 << uint(i)) > 0 {
      if reverse {
        ans = append([]string{code[i]}, ans...)
      } else {
        ans = append(ans, code[i])
      }
    }
  }
  return ans
}
