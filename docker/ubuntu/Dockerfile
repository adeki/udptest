FROM ubuntu

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update -y && apt-get install -y xinetd ntp && sed -i -E "s/(disable\s+=\s+)yes/\1no/" /etc/xinetd.d/chargen
