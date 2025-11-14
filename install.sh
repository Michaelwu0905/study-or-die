#!/bin/bash

# project name
APP_NAME="study-or-die"
echo "Building $APP_NAME"

# Check wether golang was install
if ! command -v go >/dev/null 2>&1; then
  echo "Error: Go is not installed."
  exit 1
fi

# 构建可执行文件
go build -o $APP_NAME
if [ $? -ne 0]; then
  echo "Build failed. Fix the Go errors and retry."
  exit 1
fi

# Move to /usr/local/bin
echo "Installing to your /usr/local/bin ..."

sudo mv $APP_NAME /usr/local/bin/
sudo chmod +x /usr/local/bin/$APP_NAME

# Check if available in PATH
if command -v $APP_NAME >/dev/null 2>&1; then
  echo ""
  echo "Successfully installed!"
  echo "You can now run it anywhere!"
  echo ""
  echo "    $APP_NAME -f 10 -t 5"
  echo ""
else
  echo "Installation failed: $APP_NAME not found in PATH."
fi
