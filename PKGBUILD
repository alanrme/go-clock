# This is an example PKGBUILD file. Use this as a start to creating your own,
# and remove these comments. For more information, see 'man PKGBUILD'.
# NOTE: Please fill out the license field for your package! If it is unknown,
# then please put 'unknown'.

# Maintainer: Alan D'souza <hello@xilog.xyz>
pkgname=go-clock
pkgver=1.0.0
pkgrel=1
epoch=
pkgdesc="go-clock command line utility to make a large clock using Unicode characters"
arch=()
url="https://github.com/XilogOfficial/go-clock/raw/master/"
license=('MIT')
groups=()
depends=()
makedepends=()
checkdepends=()
optdepends=()
provides=()
conflicts=()
replaces=()
backup=()
options=()
install=
changelog=
source=("$url/$pkgname-$pkgver.tar.gz")
noextract=()
md5sums=()
validpgpkeys=()

prepare(){
    cd "$pkgname-$pkgver"
    mkdir -p build
}

build(){
    cd "$pkgname-$pkgver"
    go build -o build .
}

check() {
  cd "$pkgname-$pkgver"
  go test ./...
}

package() {
  cd "$pkgname-$pkgver"
  install -Dm755 build/$pkgname "$pkgdir"/usr/bin/$pkgname
}
