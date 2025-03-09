📌 Golang Modules Nedir?

Golang Modules, Go dilinde bağımlılık yönetimi ve proje versiyonlaması için kullanılan bir sistemdir. Go 1.11 sürümü ile tanıtıldı ve 1.13 sürümünden itibaren varsayılan bağımlılık yönetim aracı oldu. Modüller sayesinde:

- Projelerimizi bağımsız paketler halinde organize edebiliriz.
- Dış bağımlılıkları yönetebiliriz (versiyon kontrolü, indirme, güncelleme).
- Farklı projeler arasında kod paylaşımını kolaylaştırabiliriz.


🛠️ 1. Modül Oluşturma: İlk Adım

Öncelikle basit bir Go modülü oluşturarak başlayalım.

📂 Klasör Yapısı:

```
go-modules-example/
    ├── go.mod
    └── main.go
```

A. Proje Klasörünü Oluştur:
```
mkdir go-modules-example
cd go-modules-example
```
B. Modülü Başlat:
```
go mod init example.com/myfirstmodule
```
    - example.com/myfirstmodule: Bu senin modül adın. Gerçek bir domain olmak zorunda değil; organizasyon yapısı gibi düşünebiliriz.

C. go.mod Dosyası:

Otomatik olarak oluşturulur ve içeriği şöyle olur:

```
module example.com/myfirstmodule

go 1.20
```