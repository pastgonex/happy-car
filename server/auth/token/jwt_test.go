package token

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

// 用于测试的密钥对
const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAlTyqaNi2fEEe8xT5K69rTjYgjgrY1KwwPJavyT331TBAdm3x
gRTZxcQOgm10GePb1OqLTNGWTAYxupavnVWORuburg6L/elId5jnI0zrwGU4kwgN
NtFn7EmB3TBcvYmFZ7MbKSvFBImB8NOhr2dhIrzxI+RBzdwMAhGiSJY03rTKaze3
zIwkDQQX4B8ZmqlYzo/sAeJicAREQOjZdJT40D+6u42NYb54nliaE/K2Gzgt++dp
frvnbKlSy4W29XjcrsJH/ZVzxfD+yx0AVJJWDdPX9m/E0mjr5aQD5zH2yLwWtZRN
dzX74fN1PMEnG2IA7p6rcLWdtxdaf5O9eEEMAQIDAQABAoIBADu2pX0EUahQKkEb
gEPxkd8em1Ialv4p07c6mCXOzM6Z2wFIJpz+qdCPDTs07oK5gKmmG4zmQ9wxkk5V
ovkyVZabKp+spGk5ocxw+vNOAPrfxY5XZ2oqyglqtDK7+i/AygdfTBviLpgNKVgi
ZgyYMe7DaT1yRL5JMOjXA3dF2uvNDcRK2JQjfDIcrsKxYpW0sVSebT8oj4Ycfi3/
S3U0S3FgVoSFmps/0CVcwKdX81VQzwuVyIt03njbmY+m0pk0XgBhkO6TwWJoVS6J
DFCRGkjzmZAK3KrP/fBco0XSVqVliJkJ+IUr/InEZWlI+mhv7LaGrxem5MtcAbGU
4u/hz30CgYEA7hZdaiUfcCFc+XxgtcBrvtUX4jR/YE4PPK3TnVw5ODUKyl2tvQZz
mO7hwb2KyWmNWYC6+nE/AbV8g7vO+U3nE59/Z0YV1qRW3jK1stMMyirBg6OBmPNf
AQgQGkCclxg/Ckf2fLUR4E0+slZvmfAvMOx8xlyzQZb78N0LfgAwgccCgYEAoHcD
5vuV+Z5Oayhbg3Zby9kAXzBAr88w8FDPnU3yXta9fKnxHxUMOXV8d9NFbWYxUEjK
l/H0LOBhqxMbfxFB/r7htZhDBND8mtn3OMs5UCjzv5tQh9YhBIwJUbQyM9lVK0qX
1dzbGPALQTCCWNaOFveLyqNj2lJgMDniFiCdg/cCgYBknp0SC+hSajcx1QfGKDEg
8EvstIUUfUjhOxFQ0rX39CrARYD0fvKBsotEZsdtwacUKVxcFVosbbfWsJuTLwI2
f9THH46BLOjtdP7nOVyRYCpyaLqPmmMPO4oani3PiVazEKCyKZAJfHu/wNnfc+tt
wLQm3OyBDr8hFzoRISFe5wKBgF5RNjj5XdDH7P3tTcT6t8Acv6wzl4H+/ZvzU+T5
IOH/xIbW+MQ1QecwGEXyJ5fC/m9bKcGf1M7f4GiGpZ3NjgXnOpHbemEFWcTHIxn9
0aU9PPA9oVsGUIf0q7GbgfqZ3wbAJHjvBNUmmubpVWRUUFZNkJw953429xTBoRMy
foNHAoGAbJdgpToUCfAA5d1qSF+KK1qczBT2txQwXs3vLrPBlALnT7HiWuW/XNep
SeAd7OFCKvKe7+NZ/wPurlgqajQrTefyJ6GeqVU2JAvkU3q2n6EjsmHTGgSZ5QqW
4u/i7K/SCB6aovIPd3kbsPORM7pR18U62hF+0BDKEIBK2xuuib0=
-----END RSA PRIVATE KEY-----`

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlTyqaNi2fEEe8xT5K69r
TjYgjgrY1KwwPJavyT331TBAdm3xgRTZxcQOgm10GePb1OqLTNGWTAYxupavnVWO
Ruburg6L/elId5jnI0zrwGU4kwgNNtFn7EmB3TBcvYmFZ7MbKSvFBImB8NOhr2dh
IrzxI+RBzdwMAhGiSJY03rTKaze3zIwkDQQX4B8ZmqlYzo/sAeJicAREQOjZdJT4
0D+6u42NYb54nliaE/K2Gzgt++dpfrvnbKlSy4W29XjcrsJH/ZVzxfD+yx0AVJJW
DdPX9m/E0mjr5aQD5zH2yLwWtZRNdzX74fN1PMEnG2IA7p6rcLWdtxdaf5O9eEEM
AQIDAQAB
-----END PUBLIC KEY-----`

func TestGenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse private key: %v", err)
	}
	g := NewJWTTokenGen("happycar/auth", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	token, err := g.GenerateToken("62680b0a53a88506efd364ae", 2*time.Hour)
	if err != nil {
		// 错了后面还会继续跑
		t.Errorf("cannot generate token: %v", err)
	}
	want := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiaGFwcHljYXIvYXV0aCIsInN1YiI6IjYyNjgwYjBhNTNhODg1MDZlZmQzNjRhZSJ9.f1eTtx24trvVFy5-a80zlSr6TBQUJgEF35WXUapOGFBSXs5F-SHagXdvNURySlNlY0qcwKrRfAvW9dbwXmh5pbkKtGgTufsyEYSx31V_ss3V38zm7P5aroXkheXwINWhvc-vIjDbwlHUeY6IcdGiI51NvxDNRc3AH9bhIJ8PNHN2XYLjA0sVJnQkqBtgXkS5qzidon76hFEuBmw_uMcFAJSq-4PBKlZk-uJVJYiWFTeMU5NeHlAf2Yjt_DI6jiTe4ekEiaPK4HmPrEL9wRc3KlfVIyFYIROclgWOp-vSk6hgI1xfFpDeCNwuWsMoaORkRxhq7XGJEAO83OrEJW-t9g"
	if token != want {
		t.Errorf("wrong token generated. want: %q; got: %q", want, token)
	}
}
