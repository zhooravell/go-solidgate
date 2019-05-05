SolidGate golang client
=======================
> SolidGate api golang wrapper

[![codecov][scrutinizer-image]][scrutinizer-link] [![License][license-image]][license-link] [![Build Status][travis-image]][travis-link] [![codecov][codecov-image]][codecov-link] 

![](https://solid.ng/wp-content/uploads/2017/07/solid_logo.png)

## Installing

``` sh
$ go get github.com/zhooravell/go-solidgate
```
``` sh
$ dep ensure -add github.com/zhooravell/go-solidgate
```

## Using
Initialize client
``` go

merchantID := "merchantID"
privateKey := "privateKey"

solidGateClient := solidgate.NewSolidGateClient(
    merchantID,
    &http.Client{},
    solidgate.NewSha512Signer(merchantID, []byte(privateKey)),
    "https://pay.signedpay.com/api/v1",
)
```
InitPayment transaction
``` go
ip := net.ParseIP("8.8.8.8")
email, _ := mail.ParseAddress("jondou@gmail.com")

initPaymentRequest := solidgate.NewInitPaymentRequest(
    1050,
    "USD",
    email,
    "UKR",
    &ip,
    "Premium package",
    "777",
    "WEB",
)

initPaymentResponse, err := solidGateClient.InitPayment(context.Background(), initPaymentRequest)

if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", initPaymentResponse)
```
Charge transaction
``` go
ip := net.ParseIP("8.8.8.8")
email, _ := mail.ParseAddress("jondou@gmail.com")

chargeRequest := solidgate.NewChargeRequest(
    1050,
    "USD",
    "123",
    "01",
    "2024",
    "JOHN SNOW",
    "4111111111111111",
    email,
    "UKR",
    &ip,
    "Premium package",
    "777",
    "WEB",
)

chargeResponse, err := solidGateClient.Charge(context.Background(), chargeRequest)

if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", chargeResponse)
```
Recurring transaction
``` go
ip := net.ParseIP("8.8.8.8")
email, _ := mail.ParseAddress("jondou@gmail.com")

recurringRequest := solidgate.NewRecurringRequest(
    1050,
    "USD",
    "7ats8da7sd8-a66dfa7-a9s9das89t",
    email,
    &ip,
    "Premium package",
    "777",
    "WEB",
)

recurringResponse, err := solidGateClient.Recurring(context.Background(), recurringRequest)

if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", recurringResponse)
```
Refund transaction
``` go
refundRequest := solidgate.NewRefundRequest("777", 1050)
refundResponse, err := solidGateClient.Refund(context.Background(), refundRequest)

if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", refundResponse)
```

## Source(s)

* [SolidGate](https://solid.ng/)
* [SolidGate Documentation](https://solidgate.atlassian.net/wiki/spaces/API/pages/4718593/EN)

[license-link]: https://github.com/zhooravell/go-solidgate/blob/master/LICENSE
[license-image]: https://img.shields.io/dub/l/vibe-d.svg

[travis-link]: https://travis-ci.com/zhooravell/go-solidgate
[travis-image]: https://travis-ci.com/zhooravell/go-solidgate.svg?branch=master

[codecov-link]: https://codecov.io/gh/zhooravell/go-solidgate
[codecov-image]: https://codecov.io/gh/zhooravell/go-solidgate/branch/master/graph/badge.svg

[scrutinizer-link]: https://scrutinizer-ci.com/g/zhooravell/go-solidgate/?branch=master
[scrutinizer-image]: https://scrutinizer-ci.com/g/zhooravell/go-solidgate/badges/quality-score.png?b=master