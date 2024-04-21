package table

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"regexp"
)

const (
	password_self = "_self_"
)

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) SetPassword(ps string) error {
	expr := `^(?![0-9a-zA-Z]+$)(?![a-zA-Z!@#$%^&*]+$)(?![0-9!@#$%^&*]+$)[0-9A-Za-z!@#$%^&*]{8,16}$`
	if regexp.MustCompile(expr).MatchString(ps) {
		h := md5.New()
		io.WriteString(h, password_self+ps)
		u.Password = hex.EncodeToString(h.Sum(nil))
		return nil
	} else {
		return fmt.Errorf("密码包含至少一位数字，字母和特殊字符,且长度8-16")
	}
}

func (u *User) SetEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
