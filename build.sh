#!/bin/bash

# Define the version
VERSION="v1.0.0"

# Define the output directory
OUTPUT_DIR="./dist"

# Define the platforms
PLATFORMS=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64")

# Clean up the output directory
rm -rf ${OUTPUT_DIR}
mkdir -p ${OUTPUT_DIR}

# Build for each platform
for PLATFORM in "${PLATFORMS[@]}"; do
  OS=$(echo $PLATFORM | cut -d'/' -f1)
  ARCH=$(echo $PLATFORM | cut -d'/' -f2)
  OUTPUT_NAME="terraform-provider-bridge_${VERSION}_${OS}_${ARCH}"

  if [ $OS == "windows" ]; then
    OUTPUT_NAME+='.exe'
  fi

  GOOS=$OS GOARCH=$ARCH go build -o ${OUTPUT_DIR}/${OUTPUT_NAME} .

  if [ $? -ne 0 ]; then
    echo "An error occurred while building for ${PLATFORM}"
    exit 1
  fi
done

echo "Builds completed successfully."
