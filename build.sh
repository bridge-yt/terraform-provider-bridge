#!/bin/bash

# Define the version
VERSION="1.0.0"

# Define the provider name
PROVIDER_NAME="bridge"

# Define the output directory
OUTPUT_DIR="./dist"

# Define the platforms
PLATFORMS=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64")

# Clean up the output directory
rm -rf ${OUTPUT_DIR}
mkdir -p ${OUTPUT_DIR}

# Build and package for each platform
for PLATFORM in "${PLATFORMS[@]}"; do
  OS=$(echo $PLATFORM | cut -d'/' -f1)
  ARCH=$(echo $PLATFORM | cut -d'/' -f2)
  BINARY_NAME="terraform-provider-${PROVIDER_NAME}_v${VERSION}"
  ARCHIVE_NAME="terraform-provider-${PROVIDER_NAME}_${VERSION}_${OS}_${ARCH}.zip"

  if [ $OS == "windows" ]; then
    BINARY_NAME+='.exe'
  fi

  echo "Building for ${PLATFORM}..."
  GOOS=$OS GOARCH=$ARCH go build -o ${OUTPUT_DIR}/${BINARY_NAME} .

  if [ $? -ne 0 ]; then
    echo "An error occurred while building for ${PLATFORM}"
    exit 1
  fi

  echo "Packaging ${ARCHIVE_NAME}..."
  (cd ${OUTPUT_DIR} && zip ${ARCHIVE_NAME} ${BINARY_NAME} && rm ${BINARY_NAME})
done

echo "Builds and packaging completed successfully. Archives are in the ${OUTPUT_DIR} directory."

# Generate the manifest file
MANIFEST_FILE="${OUTPUT_DIR}/terraform-provider-${PROVIDER_NAME}_${VERSION}_manifest.json"
echo "Generating manifest file..."
cat <<EOL > ${MANIFEST_FILE}
{
  "provider": {
    "name": "${PROVIDER_NAME}",
    "version": "${VERSION}",
    "platforms": [
    $(for ((i=0; i<${#PLATFORMS[@]}; i++)); do
        OS=$(echo ${PLATFORMS[$i]} | cut -d'/' -f1)
        ARCH=$(echo ${PLATFORMS[$i]} | cut -d'/' -f2)
        echo "      { \"os\": \"${OS}\", \"arch\": \"${ARCH}\" }$(if [ $i -lt $((${#PLATFORMS[@]} - 1)) ]; then echo ","; fi)"
      done)
    ]
  }
}
EOL

# Debug: list files in output directory
echo "Files in ${OUTPUT_DIR}:"
ls -al ${OUTPUT_DIR}

# Generate SHA256SUMS file
SHA256SUMS_FILE="${OUTPUT_DIR}/terraform-provider-${PROVIDER_NAME}_${VERSION}_SHA256SUMS"
echo "Generating SHA256SUMS file at ${SHA256SUMS_FILE}..."
(cd ${OUTPUT_DIR} && shasum -a 256 *.zip ${MANIFEST_FILE} > "${SHA256SUMS_FILE}")

# Debug: check if SHA256SUMS file was created
if [ -f "${SHA256SUMS_FILE}" ]; then
  echo "SHA256SUMS file created successfully."
else
  echo "Failed to create SHA256SUMS file."
  exit 1
fi

# Sign the SHA256SUMS file
gpg --output "${SHA256SUMS_FILE}.sig" --detach-sign "${SHA256SUMS_FILE}"

# Debug: check if SHA256SUMS.sig file was created
if [ -f "${SHA256SUMS_FILE}.sig" ]; then
  echo "SHA256SUMS.sig file created successfully."
else
  echo "Failed to create SHA256SUMS.sig file."
  exit 1
fi

echo "SHA256SUMS and SHA256SUMS.sig generated successfully."
cd ..

echo "All files are ready in the ${OUTPUT_DIR} directory."
