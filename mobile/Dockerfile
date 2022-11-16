FROM ubuntu:18.04
WORKDIR /app

RUN apt update && apt install -y curl git unzip xz-utils zip libglu1-mesa openjdk-11-jdk wget software-properties-common
RUN add-apt-repository ppa:maarten-fonville/android-studio
RUN apt update && apt install -y android-studio

RUN mkdir /root/Android
RUN mkdir /root/Android/sdk
ENV ANDROID_HOME /root/Android/sdk
RUN mkdir /root/.android && touch /root/.android/repositories.cfg

RUN wget -O sdk-tools.zip https://dl.google.com/android/repository/commandlinetools-linux-8512546_latest.zip
RUN unzip sdk-tools.zip && rm sdk-tools.zip
RUN mkdir $ANDROID_HOME/cmdline-tools/
RUN mkdir $ANDROID_HOME/cmdline-tools/latest
RUN mv cmdline-tools/* $ANDROID_HOME/cmdline-tools/latest
RUN rm -rf cmdline-tools
ENV PATH $PATH:$ANDROID_HOME/cmdline-tools/latest/bin
RUN yes | sdkmanager --licenses
RUN sdkmanager "cmdline-tools;latest"

RUN cd /root && git clone https://github.com/flutter/flutter.git
ENV PATH "$PATH:/root/flutter/bin"

RUN flutter upgrade
RUN flutter doctor

# ENTRYPOINT ["tail", "-f", "/dev/null"]

# CMD ./start.sh
CMD flutter build apk --split-per-abi