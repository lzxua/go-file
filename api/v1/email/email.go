package email

import "strconv"
import "gopkg.in/gomail.v2"

// 可以通过 text/html 处理文本格式进行特殊处理，如换行、缩进、加粗等等
// m.SetBody("text/html", fmt.Sprintf(message, "testUser"))
// text/plain的意思是将文件设置为纯文本的形式，浏览器在获取到这种文件时并不会对其进行处理
// m.SetBody("text/plain", "纯文本")
// m.Attach("test.sh")   // 附件文件，可以是文件，照片，视频等等
// m.Attach("lolcatVideo.mp4") // 视频
// m.Attach("lolcat.jpg") // 照片

// SendMail 发送邮件
// mailTo是收件人
// username 是发件人的邮箱
// sendName 是名字
// subject 邮件主题
// body 内容
func SendMail(userName, authCode, host, portStr, mailTo, sendName string, subject, body string) error {
	port, _ := strconv.Atoi(portStr)
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, authCode)
	err := d.DialAndSend(m)
	return err
}
