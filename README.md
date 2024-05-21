# DNS Resolver

Bu basit bir DNS çözümleme programıdır. Program, kullanıcıdan alınan bir alan adı (domain) için DNS çözümleme yaparak A kaydını (IP adresini) bulmaktadır. Program, `github.com/miekg/dns` kütüphanesini kullanarak DNS sorgularını gerçekleştirir.

## Kullanım

Programın kullanımı oldukça basittir. Sadece çalıştırıldığında kullanıcıdan bir alan adı alır ve bu alan adı için DNS çözümleme yapar. Sonuçları ekrana yazdırır.

```bash
go run main.go
```

## Gereksinimler

Programın çalışması için Go dilinin yüklü olması gerekmektedir. Ayrıca, `github.com/miekg/dns` kütüphanesine bağımlıdır. Bağımlılıkları yüklemek için:

```bash
go get github.com/miekg/dns
```

## Örnek Çıktı

![!altext](https://i.ibb.co/1T7xGZs/Screenshot-from-2024-05-21-20-14-48.png")

