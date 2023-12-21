FROM ubuntu

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update -y && \
    apt-get install -y vim xinetd ntp miniupnpd iptables nftables && \
    sed -i -E "s/(disable\s+=\s+)yes/\1no/" /etc/xinetd.d/chargen && \
    cp /etc/miniupnpd/miniupnpd.conf /etc/miniupnpd.conf && \
    sed -i -E "s/^#(ext_ifname=)eth1/\1eth0/" /etc/miniupnpd.conf && \
    sed -i -E "s/^#(listening_ip=eth0)/\1/" /etc/miniupnpd.conf && \
    sed -i -E "s/^#(enable_natpmp=).+/\1yes/" /etc/miniupnpd.conf
