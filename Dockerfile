FROM registry.fedoraproject.org/fedora:32 AS builder
RUN dnf install -y golang-bin; dnf clean all;
WORKDIR /opt
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /opt/bin

FROM registry.fedoraproject.org/fedora:32
ADD wait-for-postgres.sh /app/
RUN dnf install -y golang-bin postgresql; dnf clean all;
COPY --from=builder /opt/bin/rhcos-scanner /app/
WORKDIR /app
ENTRYPOINT ["wait-for-postgres.sh", "--" "./rhcos-scanner"]
