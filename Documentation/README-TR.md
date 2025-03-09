# Golang Modules User Guide

Bu rehber, Go programlama dilinde modül oluşturma, bağımlılık yönetimi ve modül sürümlendirme süreçlerini kapsar. Golang Modülleri, bağımsız paketler oluşturmayı ve projeler arasında kod paylaşımını kolaylaştırır.

---

## 📌 Golang Modülleri Nedir?

Golang Modülleri, Go dilinde bağımlılık yönetimi ve proje sürümlendirme için kullanılan bir sistemdir. Go 1.11 sürümünde tanıtıldı ve 1.13 sürümünden itibaren varsayılan bağımlılık yönetim aracı haline geldi. Modüller ile şunları yapabilirsiniz:

- Projelerinizi bağımsız paketlere ayırabilirsiniz.
- Harici bağımlılıkları yönetebilirsiniz (sürüm kontrolü, indirme, güncelleme).
- Farklı projeler arasında kod paylaşımını kolaylaştırabilirsiniz.

---

## 🛠️ 1. Modül Oluşturma: İlk Adımlar

### 📂 Klasör Yapısı:

```plaintext
learn-golang-modules
    ├── go.mod
    ├── main.go
    ├── goodbye
    ├── greeting
    └── hello
```

Bu yapıda "goodbye", "greeting" ve "hello" adında üç ayrı modül oluşturulmuştur.

### 📋 Adımlar:

#### A. Proje Klasörünü Oluşturun:
```bash
mkdir learn-golang-modules
cd learn-golang-modules
```

#### B. Modülü Başlatın:
```bash
go mod init https://github.com/miraccan00/learn-golang-modules.git
```
- `https://github.com/miraccan00/learn-golang-modules.git`: Modül adı (gerçek bir domain olması gerekmez).

#### C. go.mod Dosyası:
Otomatik olarak aşağıdaki içerikle oluşturulur:

```go
module https://github.com/miraccan00/learn-golang-modules.git

go 1.24
```

---

## 🚀 2. İlk Go Kodunuzu Yazın

### main.go Dosyasını Oluşturun:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go Modules!")
}
```

### Uygulamayı Çalıştırın:
```bash
go run main.go
```

**Beklenen Çıktı:**
```plaintext
Hello, Go Modules!
```

---

## 🔗 3. Harici Bir Paket Ekleme
Örneğin, `github.com/fatih/color` paketini ekleyelim.

```bash
go get github.com/fatih/color
```

### Kodu Güncelleyin:

```go
package main

import (
    "fmt"
    "github.com/fatih/color"
)

func main() {
    fmt.Println("Normal text")
    color.Red("This is red text")
    color.Green("And this is green!")
}
```

### Tekrar Çalıştırın:
```bash
go run main.go
```

---

## 🏷️ 4. Modül Sürümlendirme

### Git Etiketleri Kullanarak:

```bash
# Sürüm oluştur
git tag v1.0.0

# Uzak depoya gönder
git push origin v1.0.0
```

### GitHub Üzerinde Yeni Sürüm Oluşturma:

1. GitHub depo sayfasına gidin.
2. **Releases** bölümünde yeni bir **Tag** oluşturun.
3. Etiket adını `v1.0.0` gibi belirleyin.

---

## 🧠 Faydalı Go Mod Komutları

| Komut | Açıklama |
|---------|-------------|
| `go mod init <module_name>` | Yeni bir modül başlatır. |
| `go mod tidy` | Kullanılmayan bağımlılıkları temizler. |
| `go get <package>` | Harici bir paketi indirir. |
| `go test ./...` | Tüm paketlerdeki testleri çalıştırır. |

---

## Bu Paketi Nasıl Kullanırım?
```bash
go get https://github.com/miraccan00/learn-golang-modules@v1.0.0
```


## 🎯 Sonuç

Bu rehber ile şunları öğrendiniz:
- Sıfırdan Go modülleri oluşturmayı.
- Harici bağımlılıkları yönetmeyi.
- Git etiketlerini kullanarak sürüm yönetimini gerçekleştirmeyi.

Daha fazla bilgi için resmi Go dokümantasyonunu ziyaret edebilirsiniz: [Golang Modules](https://golang.org/doc/go1.11#modules)

