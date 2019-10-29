FROM envoyproxy/envoy:latest
COPY ./envoy.grpcproxy.yaml /etc/envoy/envoy.grpcproxy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.grpcproxy.yaml