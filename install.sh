# Install Script
# Usage: curl -s $addr | sh

platform='unknown'
host='https://connect.segment.com.s3-us-west-2.amazonaws.com'
binary='connect-kafka'
install_dir='/usr/local/bin'
arch='amd64'

if [[ "$OSTYPE" == "linux-gnu" ]]; then
  platform='linux'
elif [[ "$OSTYPE" == "darwin"* ]]; then
  platform='darwin'
elif [[ "$OSTYPE" == "cygwin" ]]; then
  platform='windows'
elif [[ "$OSTYPE" == "msys" ]]; then
  platform='windows'
elif [[ "$OSTYPE" == "win32" ]]; then
  echo 'Platform is not supported!'
  exit 1
elif [[ "$OSTYPE" == "freebsd"* ]]; then
  platform='freebsd'
else
  echo 'Platform is not supported!'
  exit 1
fi

echo "Installing $binary for $platform/$arch..."
echo "Debug $host/$binary-$platform-$arch"

curl -s "$host/$binary-$platform-$arch" >> "$install_dir/$binary"
chmod +x $install_dir/$binary

size=$(wc -c <"$install_dir/$binary")

echo "Size: $size"

echo "$binary was installed successfully to $install_dir"
