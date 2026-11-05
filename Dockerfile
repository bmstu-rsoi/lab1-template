FROM ubuntu:24.04 AS builder

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y --no-install-recommends \
    clang-19 \
    clang-tools-19 \
    cmake \
    ninja-build \
    libdrogon-dev \
    libpq-dev \
    libjsoncpp-dev \
    uuid-dev \
    zlib1g-dev \
    libsqlite3-dev \
    libmariadb-dev \
    libbrotli-dev \
    libhiredis-dev \
    libyaml-cpp-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .

RUN cmake -S . -B build -G Ninja \
    -DCMAKE_CXX_COMPILER=clang++-19 \
    -DCMAKE_CXX_COMPILER_CLANG_SCAN_DEPS=clang-scan-deps-19 \
    -DBUILD_TESTING=OFF \
    -DCMAKE_BUILD_TYPE=Release

RUN cmake --build build --parallel


FROM ubuntu:24.04 AS runtime

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y --no-install-recommends \
    libdrogon1 \
    libpq5 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY postgres/init.sql /docker-entrypoint-initdb.d/init.sql

COPY --from=builder /app/build/src/person_service ./person_service

EXPOSE 8080

CMD ["./person_service"]