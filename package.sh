# The more popular operating systems that Go supports.
platforms=(darwin openbsd freebsd linux)
arch=amd64
bucket='connect.segment.com'
install="install-connect-kafka.sh"
name="connect-kafka"
host='https://connect.segment.com.s3-us-west-2.amazonaws.com'

if hash gpg 2>/dev/null; then
  echo 'gpg was found, skipping...'
else
  echo 'gpg is not installed. Using homebrew to install it...'
  brew install gpg
fi


build() {
  local platform=$1

  rm -rf target
  mkdir -p target

  echo "Building for $platform on $arch"
  GOOS=$platform GOARCH=$arch go build -ldflags "-s -w" -o target/$name-$platform-$arch
}

build_all() {
  for i in "${platforms[@]}"
  do
    build $i
  done
}

upload() {
  echo "Uploading artifacts to s3..."

  # Upload the targets to the production S3 bucket.
  aws-vault exec production -- aws s3 cp target/ s3://$bucket/ --recursive
  aws-vault exec production -- aws s3 cp install.sh s3://$bucket/$install

  echo "\n\nInstall script available at $host/$install"
  echo "\nTo install, run:"
  echo "-------------------------------------------"
  echo "   $ curl -s $host/$install | sh         "
  echo "-------------------------------------------"
}

case $1 in
  linux)
    build 'linux'
  ;;
  *)
    build_all
    upload
  ;;
esac
