package corn

import "time"

func Run() {
	go func() {
		_ = SendSms("123456789")
		t:= time.NewTimer(time.Hour*6)
		<-t.C
	}()
}
