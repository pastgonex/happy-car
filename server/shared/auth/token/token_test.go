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

func TestVerify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot parse public key: %v", err)
	}
	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	cases := []struct {
		name    string // 给人看的， 看一下是哪个测试用例出错了
		tkn     string
		now     time.Time // 什么时候运行这个测试
		want    string    // 用户ID
		wantErr bool      // 是否能够出错（token过期之类的）
	}{
		{
			name:    "valid_token",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiaGFwcHljYXIvYXV0aCIsInN1YiI6IjYyNjgwYjBhNTNhODg1MDZlZmQzNjRhZSJ9.f1eTtx24trvVFy5-a80zlSr6TBQUJgEF35WXUapOGFBSXs5F-SHagXdvNURySlNlY0qcwKrRfAvW9dbwXmh5pbkKtGgTufsyEYSx31V_ss3V38zm7P5aroXkheXwINWhvc-vIjDbwlHUeY6IcdGiI51NvxDNRc3AH9bhIJ8PNHN2XYLjA0sVJnQkqBtgXkS5qzidon76hFEuBmw_uMcFAJSq-4PBKlZk-uJVJYiWFTeMU5NeHlAf2Yjt_DI6jiTe4ekEiaPK4HmPrEL9wRc3KlfVIyFYIROclgWOp-vSk6hgI1xfFpDeCNwuWsMoaORkRxhq7XGJEAO83OrEJW-t9g",
			now:     time.Unix(1516239122, 0),
			want:    "62680b0a53a88506efd364ae",
			wantErr: false, // 不应该出错
		},
		{
			name:    "token_expired",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiaGFwcHljYXIvYXV0aCIsInN1YiI6IjYyNjgwYjBhNTNhODg1MDZlZmQzNjRhZSJ9.f1eTtx24trvVFy5-a80zlSr6TBQUJgEF35WXUapOGFBSXs5F-SHagXdvNURySlNlY0qcwKrRfAvW9dbwXmh5pbkKtGgTufsyEYSx31V_ss3V38zm7P5aroXkheXwINWhvc-vIjDbwlHUeY6IcdGiI51NvxDNRc3AH9bhIJ8PNHN2XYLjA0sVJnQkqBtgXkS5qzidon76hFEuBmw_uMcFAJSq-4PBKlZk-uJVJYiWFTeMU5NeHlAf2Yjt_DI6jiTe4ekEiaPK4HmPrEL9wRc3KlfVIyFYIROclgWOp-vSk6hgI1xfFpDeCNwuWsMoaORkRxhq7XGJEAO83OrEJW-t9g",
			now:     time.Unix(1517239122, 0),
			want:    "", //写空串或者直接不写
			wantErr: true,
		},
		{
			name:    "bad_token",
			tkn:     "bad_token",
			now:     time.Unix(1517239122, 0),
			want:    "", //写空串或者直接不写
			wantErr: true,
		},
		{
			name:    "user fake token", // 用户伪造token,
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiaGFwcHljYXIvYXV0aCIsInN1YiI6IjYyNjgwYjBhNTNhODg1MDZlZmQzNjRhZSJ9.f1eTtx24trvVFy5-a80zlSr6TBQUJgEF35WXUapOGFBSXs5F-SHagXdvNURySlNlY0qcwKrRfAvW9dbwXmh5pbkKtGgTufsyEYSx31V_ss3V38zm7P5aroXkheXwINWhvc-vIjDbwlHUeY6IcdGiI51NvxDNRc3AH9bhIJ8PNHN2XYLjA0sVJnQkqBtgXkS5qzidon76hFEuBmw_uMcFAJSq-4PBKlZk-uJVJYiWFTeMU5NeHlAf2Yjt_DI6jiTe4ekEiaPK4HmPrEL9wRc3KlfVIyFYIROclgWOp-vSk6hgI1xfFpDeCNwuWsMoaORkRxhq7XGJEAO83OrEJW-tgg",
			now:     time.Unix(1516239122, 0),
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// 使用jwt提供的方法，手动设置时间
			jwt.TimeFunc = func() time.Time {

				return c.now
			}
			accountID, err := v.Verify(c.tkn)
			if !c.wantErr && err != nil {
				t.Errorf("verification failed: %v", err)
			}
			if c.wantErr && err == nil {
				t.Errorf("want error; got no error")
			}
			if accountID != c.want {
				t.Errorf("wrong account id. want: %q; go: %q", c.want, accountID)
			}
		})
	}
}
