package director_test

import (
	"crypto/tls"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	. "github.com/cloudfoundry/bosh-init/director"
)

func BuildServer() (Director, *ghttp.Server) {
	server := ghttp.NewUnstartedServer()

	cert, err := tls.X509KeyPair(validCert, validKey)
	Expect(err).ToNot(HaveOccurred())

	server.HTTPTestServer.TLS = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	server.HTTPTestServer.StartTLS()

	config, err := NewConfigFromURL(server.URL())
	Expect(err).ToNot(HaveOccurred())

	config.Username = "username"
	config.Password = "password"
	config.CACert = validCACert

	logger := boshlog.NewLogger(boshlog.LevelNone)
	taskReporter := NewNoopTaskReporter()
	fileReporter := NewNoopFileReporter()

	director, err := NewFactory(logger).New(config, taskReporter, fileReporter)
	Expect(err).ToNot(HaveOccurred())

	return director, server
}

var validCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDDTCCAfWgAwIBAgIJAOYPl1HNpMPsMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwIBcNMTYwMTE2MDY0NTA0WhgPMjI4OTEwMzAwNjQ1MDRa
MDAxCzAJBgNVBAYTAlVTMQ0wCwYDVQQKDARCT1NIMRIwEAYDVQQDDAkxMjcuMC4w
LjEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDlptk3/IXbiBgJO7DO
dSc9MASV7FSBATumxQcXvKzUuaBJECD/S/QdevoBtIXQhtyNdSNu8GN6cD550xs2
3DYibgPD+At1IxRHfGu0Hxn2ZbU4yP9SqUchJHOa7Rix6T2cnauYhh+FhilO0Elm
kOyOtAshnv70ZWUDez8ybExgSK2kCiq3tmFotNHpxN6gNJ9IQfYz1U3thX/kyjag
MrOTTzluGGgpyS7o+4eD5rL/pWTylkgufhqUm4CJkRbXlJ8Dd/bwuBtRTumO6C4q
sYU6/OGQT/HM+sYDzrUd2pe36dQ41oeWZhKn2DyixnLLqlcH3QxnHTeg139sIQfy
rIMPAgMBAAGjEzARMA8GA1UdEQQIMAaHBH8AAAEwDQYJKoZIhvcNAQEFBQADggEB
AKj2aCf1tLQFZLq+TYa/THoVP7Pmwnt49ViQO8nMnfCM3459d52vCdIodGocVg9T
x8N4/pIG3S0VCzQt+4+UJej6SyecrYpcMCtWhZ73zxTJ7lQUmknsqZCvC5BcjYgF
McML3CeFsHuHvwb7uH5h8VO6UWyFTj7xNsH4E3XZT3I92fdS11pfrBSJDGfkiAQ/
j3N1QevrxTlEuKLQFfFSbnA3XZGpkDzg/sqYiOHnVgbn84IIZ3lGXs+qzC5kTFfM
SC0K79vs7peS+FdzPUAuG7uyy0W0s5hFTRIlcvBO5w9QrwEnBEv7WrZ6oSZ5F3Ku
/M/AnjGop4LUFIbJQR0ns7U=
-----END CERTIFICATE-----`)

var validKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA5abZN/yF24gYCTuwznUnPTAElexUgQE7psUHF7ys1LmgSRAg
/0v0HXr6AbSF0IbcjXUjbvBjenA+edMbNtw2Im4Dw/gLdSMUR3xrtB8Z9mW1OMj/
UqlHISRzmu0Ysek9nJ2rmIYfhYYpTtBJZpDsjrQLIZ7+9GVlA3s/MmxMYEitpAoq
t7ZhaLTR6cTeoDSfSEH2M9VN7YV/5Mo2oDKzk085bhhoKcku6PuHg+ay/6Vk8pZI
Ln4alJuAiZEW15SfA3f28LgbUU7pjuguKrGFOvzhkE/xzPrGA861HdqXt+nUONaH
lmYSp9g8osZyy6pXB90MZx03oNd/bCEH8qyDDwIDAQABAoIBAH82/f1VlZEWwrna
pwa3PxVWFDQ4xlbwJ+sqGdO8YME2UuQmWyERIhlylit7pOTu0B5MVWSPJYwdwX4a
w2iQdCx+ZPeZ4D7zP7iZ48/Tqr4jeVALh+RygUSKvL+Ft7hWTBsF/JhxM+TzfM57
8y0t+tzSP5hQS0t3H43eKBP2ihiLHVQwSV8F3GTNh2/yc3+Y+usO4n5T38Q6nC8K
OblPkR1riLMReMZRdhDvIox1OwZC00PH/2tJP/vLAHXxxGuD7Bo7BjuOND3aZ9Du
xi998w1B4LrRI/W9X53Q0q+GEGwGbuvEl7GDlihiucse7GzO6+Q0qk5bNvlmd0jL
EW3hGQECgYEA+5tznyQYxZQOhSpuISY3MbZ+SL1fZzmpnQviegWJg8dZTrUxLOmK
ku4Gyr+S+kA/tgK6ys4qRlzF2UCrytXGwoQuOxkK81rbhHqmZEZupLd5PXyAtZpz
AySUK/YLrtmXZP+gvOhO2ss9jMD8Nwy8nMa3hZDBomE63sHhFwgHhMcCgYEA6alE
ieOdMJJZ8ZpDMnYtwXuU31ne9Gj7lkirbSZ/l/SZ69xCMp8O6Oj66w5l8lH42lP3
cqJ+n35F0/TwKvft9nXCrc5H1zUw5qKYgDqn80bHk2cPqrenikSK3RhKHW4gJm+R
J9SY6bPPz/CCMqMX+cC/bjhdH/2gkldra2GoN3kCgYEA3bM9LwXsefQa00Xu4nC9
A6XtIoUTEm7hwIrfVWuZny9BxzOrEAr82ri37WDeznlcajF/jAIbiAJpJyRv+3tg
9rbn0ZUgbAwsD1DPWt4g0i0EvKP++YYNP8C0ewQDiV8bopgId0wvZ2TcaDEITC2B
6JbE0QEbTcxkxjGJ9/RQQ7MCgYEA3B681ZWaoIZOuz8S7LfONQah4aM9WUyJLjN5
YxMwgktIsZxGtH+JQTsyHjvrKFO2tp8BbnnMBZ6kU5/cnO4Bu/uGEcxRe1i9n5gv
SCV50MGuA5vEc5Qd/jDCDLT0JTN4kBzsRvSNtSPSstalIOTqEjtVW5U3jYqWOSan
qHpQSSkCgYBkXiIhH7rPd5oNYccNEzEr+0ew6HqzP8AkDSLTP323JK9kGw+dvEb/
dEG/RBqYUo0MiCCNXOVOsri1tL5cKZEfWgcTyzbkX/7BgHMWHkAD5QnhXaik5NZN
nLpUTgSaa9Cd6yjEW4wGyls8DxPHonM3XDSFc15VX1VFQiwbZBxQiw==
-----END RSA PRIVATE KEY-----`)

var validCACert = `-----BEGIN CERTIFICATE-----
MIIDXzCCAkegAwIBAgIJAPerMgLAne5vMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwIBcNMTYwMTE2MDY0NTA0WhgPMjI4OTEwMzAwNjQ1MDRa
MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJ
bnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw
ggEKAoIBAQCtSo3KPjnVPzodb6+mNwbCdcpzVop8OmfwJ3ynQtyBEzGaKsAn4tlz
/wfQQrKFHgxqVpqcoxAlWPNMs5+iO2Jst3Gz2+oLcaDyz/EWorw0iF5q1F6+WYHp
EijY20MzaWYMyu4UhhlbJCkSGZSjujh5SFOAXQwWYJXsqjyxA9KaTD6OdH5Kpger
B9D4zogX0We00eouyvvz/sAeDbTshk9sJRGWHNFJr+TjVx2D01alU49liAL94yF6
1eEOEbE50OAhv9RNsRh6O58idaHg30bbMf1yAzcgBvh8CzIHH0BPofoF2pRfztoY
uudZ0ftJjTz4fA2h/7GOVzxemrTjx88vAgMBAAGjUDBOMB0GA1UdDgQWBBQjz5Q2
YW2kBTb4XLqKFZMSBLpi6zAfBgNVHSMEGDAWgBQjz5Q2YW2kBTb4XLqKFZMSBLpi
6zAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQA/s94M/mSGELHJWIb1
oE0IKHWajBd3Pc8+O1TZRE+ke3q+rZRfcxd2dAjq6zQHJUs2+fs0B3DyT9Wtyyoq
UrRdsgprOdf2Cuw8bMIsCQOvqWKhhdlLTnCi2xaGJawGsIkheuD1n+Il9gRQ2WGy
lACxVngPwjNYxjOE+CUnSZCuAmAfQYzqto3bNPqkgEwb7ueODeOiyhR8SKsH7ySW
QAOSxgrLBblGLWcDF9fjMeYaUnI34pHviCKeVxfgsxDR+Jg11F78sPdYLOF6ipBe
/5qTYucsY20B2EKtlscD0mSYBRwbVrSQt2RYbTCwaibxWUC13VV+YEk0NAv9Mm04
6sKO
-----END CERTIFICATE-----`
