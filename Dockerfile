# ใช้ base image ของ Golang
FROM golang:1.24

# ตั้งค่า working directory
WORKDIR /app

# คัดลอกไฟล์ทั้งหมดไปยัง container
COPY . .

# ติดตั้ง dependencies
RUN go mod tidy

# คอมไพล์โค้ด Go
RUN go build -o main .

# คำสั่งเริ่มรัน API
CMD ["/app/main"]
