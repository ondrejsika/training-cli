FROM debian:10-slim
COPY training-cli /
ENTRYPOINT [ "/training-cli" ]
