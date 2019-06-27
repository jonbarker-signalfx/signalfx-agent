FROM owasp/dependency-check:5.1.0
USER root

RUN mv /etc/apt/sources.list.d/microsoft-prod.list /etc/apt/sources.list.d/microsoft-prod.list.disabled && \
    apt-get update && \
    apt-get install -y --no-install-recommends apt-transport-https wget && \
    mv /etc/apt/sources.list.d/microsoft-prod.list.disabled /etc/apt/sources.list.d/microsoft-prod.list

ARG GO_VERSION=1.12.1
RUN wget -nv https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    echo 'analyzer.golang.path=/usr/local/go/bin/go' > /usr/share/dependency-check/go.properties && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

COPY suppression.xml /usr/share/dependency-check/

ENV PATH=$PATH:/usr/share/dependency-check/bin:/usr/local/go/bin

VOLUME ["/bundle" "/src" "/usr/share/dependency-check/data" "/reports"]

WORKDIR /

ENTRYPOINT ["dependency-check.sh"]
CMD ["--help"]
