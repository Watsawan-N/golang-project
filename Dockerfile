# ใช้ Microsoft SQL Server 2019 Image
FROM mcr.microsoft.com/mssql/server:2019-latest

# กำหนด Environment Variables
ENV ACCEPT_EULA=Y
ENV SA_PASSWORD=YourStrong!Passw0rd
ENV MSSQL_PID=Developer

# เปิดพอร์ต SQL Server (ค่าเริ่มต้นคือ 1433)
EXPOSE 1433

# รัน SQL Server
CMD ["/opt/mssql/bin/sqlservr"]
