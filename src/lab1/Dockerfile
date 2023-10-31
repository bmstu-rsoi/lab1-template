FROM mcr.microsoft.com/dotnet/aspnet:5.0 AS base
WORKDIR /app
EXPOSE 8080

FROM mcr.microsoft.com/dotnet/sdk:5.0 AS build
WORKDIR /src
COPY ["WebApplication3.csproj", "."]
RUN dotnet restore "./WebApplication3.csproj"
COPY . .
WORKDIR "/src/."
RUN dotnet build "WebApplication3.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "WebApplication3.csproj" -c Release -o /app/publish

FROM base AS final
ENV CON_STRING="Host = postgres; Port = 5432; Database = persons; Username = postgres; Password = 1234"
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "lab1.dll", "--urls=http://0.0.0.0:8080"]



