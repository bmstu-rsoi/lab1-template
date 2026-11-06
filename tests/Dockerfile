FROM ubuntu:24.04 AS dependencies

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
    libgtest-dev \
    && rm -rf /var/lib/apt/lists/*


FROM dependencies AS builder

WORKDIR /app

COPY . .

RUN cmake -S . -B build -G Ninja \
    -DCMAKE_CXX_COMPILER=clang++-19 \
    -DCMAKE_CXX_COMPILER_CLANG_SCAN_DEPS=clang-scan-deps-19 \
    -DBUILD_TESTING=ON \
    -DCMAKE_BUILD_TYPE=Release

RUN cmake --build build --parallel


FROM dependencies AS tester

WORKDIR /app

COPY --from=builder /app/build ./build

CMD ["ctest", "--test-dir", "build", "--output-on-failure"]