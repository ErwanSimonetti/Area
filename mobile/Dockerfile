FROM ubuntu:18.04

# Prerequisites
RUN apt update && apt install -y curl git unzip xz-utils zip libglu1-mesa openjdk-8-jdk wget && apt install -y android-tools-adb
RUN bash -c 'apt-get update && apt-get -y dist-upgrade'
RUN apt install --assume-yes clang cmake ninja-build pkg-config libgtk-3-dev android-sdk-platform-tools-common


# Set up new user
RUN useradd -ms /bin/bash mobile
USER mobile
WORKDIR /home/mobile

# Prepare Android directories and system variables
RUN mkdir -p Android/sdk
ENV ANDROID_SDK_ROOT /home/mobile/Android/sdk
RUN mkdir -p .android && touch .android/repositories.cfg

# Set up Android SDK
RUN wget -O sdk-tools.zip https://dl.google.com/android/repository/sdk-tools-linux-4333796.zip
RUN unzip sdk-tools.zip && rm sdk-tools.zip
RUN mv tools Android/sdk/tools
RUN cd Android/sdk/tools/bin && yes | ./sdkmanager --licenses
RUN cd Android/sdk/tools/bin && ./sdkmanager "build-tools;29.0.2" "patcher;v4" "platform-tools" "platforms;android-29" "sources;android-29"
ENV PATH "$PATH:/home/mobile/Android/sdk/platform-tools"

# Download Flutter SDK
RUN git clone https://github.com/flutter/flutter.git
ENV PATH "$PATH:/home/mobile/flutter/bin"

COPY workspace .

# Run basic check to download Dark SDK
RUN git config --global --add safe.directory /home/mobile/flutter
RUN /home/mobile/Android/sdk/tools/bin/sdkmanager --install "cmdline-tools;latest"
RUN flutter doctor
# RUN flutter build apk --split-per-abi
RUN flutter run
