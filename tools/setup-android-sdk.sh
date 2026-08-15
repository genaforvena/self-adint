#!/usr/bin/env bash
# One-shot Android build toolchain setup for the self-adint pilot test app.
# Installs JDK17 + Android cmdline-tools + platform/build-tools into ~/android-sdk.
# Idempotent-ish: skips downloads already present. Safe to re-run.
set -euo pipefail

SDK="$HOME/android-sdk"
CLT_ZIP="$HOME/android-ndk-dl/commandlinetools-linux.zip"
CLT_URL="https://dl.google.com/android/repository/commandlinetools-linux-11076708_latest.zip"
PLATFORM="platforms;android-34"
BUILDTOOLS="build-tools;34.0.0"

echo "== [1/5] JDK17 =="
if ! dpkg -s openjdk-17-jdk-headless >/dev/null 2>&1; then
  sudo DEBIAN_FRONTEND=noninteractive apt-get update -qq
  sudo DEBIAN_FRONTEND=noninteractive apt-get install -y -qq openjdk-17-jdk-headless
fi
JAVA17="$(dirname "$(dirname "$(readlink -f "$(command -v javac 2>/dev/null || echo /usr/lib/jvm/java-17-openjdk-amd64/bin/javac)")")")"
export JAVA_HOME="/usr/lib/jvm/java-17-openjdk-amd64"
echo "JAVA_HOME=$JAVA_HOME"
"$JAVA_HOME/bin/java" -version 2>&1 | head -1

echo "== [2/5] cmdline-tools download =="
mkdir -p "$SDK/cmdline-tools"
if [ ! -s "$CLT_ZIP" ]; then
  curl -fL --retry 3 -o "$CLT_ZIP" "$CLT_URL"
fi
echo "zip: $(du -h "$CLT_ZIP" | cut -f1)"

echo "== [3/5] unpack cmdline-tools =="
if [ ! -x "$SDK/cmdline-tools/latest/bin/sdkmanager" ]; then
  rm -rf "$SDK/cmdline-tools/tmp" "$SDK/cmdline-tools/latest"
  mkdir -p "$SDK/cmdline-tools/tmp"
  unzip -q "$CLT_ZIP" -d "$SDK/cmdline-tools/tmp"
  mv "$SDK/cmdline-tools/tmp/cmdline-tools" "$SDK/cmdline-tools/latest"
  rmdir "$SDK/cmdline-tools/tmp" 2>/dev/null || true
fi
SDKMAN="$SDK/cmdline-tools/latest/bin/sdkmanager"
echo "sdkmanager: $SDKMAN"

echo "== [4/5] licenses =="
yes | "$SDKMAN" --sdk_root="$SDK" --licenses >/dev/null 2>&1 || true

echo "== [5/5] platform-tools + $PLATFORM + $BUILDTOOLS =="
"$SDKMAN" --sdk_root="$SDK" "platform-tools" "$PLATFORM" "$BUILDTOOLS"

echo "== DONE =="
ls "$SDK"
"$SDK/platform-tools/adb" version 2>&1 | head -1 || true
