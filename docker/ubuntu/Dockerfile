FROM ubuntu

ARG chargen_udp_port=10019

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update -y && \
    apt-get install -y vim iproute2 xinetd ntp miniupnpd iptables nftables && \
    sed -i -E "s/(disable\s+=\s+).+/\1no/" /etc/xinetd.d/chargen && \
    sed -i -E "s/(port\s+=\s+).+/\1${chargen_udp_port}/" /etc/xinetd.d/chargen-udp && \
    sed -i -E "s/(disable\s+=\s+).+/\1no/" /etc/xinetd.d/chargen-udp && \
    cp /etc/miniupnpd/miniupnpd.conf /etc/miniupnpd.conf && \
    sed -i -E "s/^#(ext_ifname=)eth1/\1eth0/" /etc/miniupnpd.conf && \
    sed -i -E "s/^#(listening_ip=eth0)/\1/" /etc/miniupnpd.conf && \
    sed -i -E "s/^#(enable_natpmp=).+/\1yes/" /etc/miniupnpd.conf
