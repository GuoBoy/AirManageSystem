package corn

import "fmt"

func SendSms(phone string) error {
	fmt.Printf("给[%s]发送短信", phone)
	return nil
}
