#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

# WARNING: change <your_api_project_name> with real value

FROM mcr.microsoft.com/dotnet/aspnet:5.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:5.0 AS build
WORKDIR /src
COPY ["nuget.config", "."]
COPY ["src/<your_api_project_name>/<your_api_project_name>.csproj", "src/<your_api_project_name>/"]

ARG PAT=secret
RUN sed -i "s|</configuration>|<packageSourceCredentials><genocs><add key=\"Username\" value=\"PAT\" /><add key=\"ClearTextPassword\" value=\"${PAT}\" /></genocs></packageSourceCredentials></configuration>|" nuget.config

RUN dotnet restore "src/<your_api_project_name>/<your_api_project_name>.csproj"
COPY . .
WORKDIR "/src/src/<your_api_project_name>"
RUN dotnet build "<your_api_project_name>.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "<your_api_project_name>.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "<your_api_project_name>.dll"]