# Thai Baht Text Converter

A simple Go project for converting decimal numbers into Thai Baht text format.

This project was developed for the Backend Take Home assignment.

---

## Features

- Convert numbers to Thai Baht text
- Support decimal values using `shopspring/decimal`
- Handle Thai language special cases:
  - สิบ
  - ยี่สิบ
  - เอ็ด
- Support large numbers (ล้าน, ล้านล้าน)
- Unit tests included

---

## Example

| Input | Output |
|---|---|
| 1234 | หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน |
| 33333.75 | สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์ |
| 1000000 | หนึ่งล้านบาทถ้วน |


## Install Dependencies

```bash
go mod tidy
```

---

## Run the Project

```bash
go run main.go
```

---

## Run Tests

```bash
go test ./...
```

---


## Author

Your Name
