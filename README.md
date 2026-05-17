# Thai Baht Text Converter

A Go project for converting decimal numbers into Thai Baht text format.

This project was developed for the Backend Take Home assignment.

---

# Features

- Convert decimal numbers to Thai Baht text
- Support Thai currency format:
  - บาท
  - สตางค์
  - ถ้วน
- Handle Thai language special cases:
  - สิบ
  - ยี่สิบ
  - เอ็ด
- Support large numbers (ล้าน, ล้านล้าน)
- Unit tests included

---

# Example Outputs

| Input | Output |
|---|---|
| 1234 | หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน |
| 33333.75 | สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์ |
| 1000000 | หนึ่งล้านบาทถ้วน |
| 0.01 | ศูนย์บาทหนึ่งสตางค์ |

---

# Clone Repository

```bash
git clone https://github.com/saran21332/ThaibathText.git
```

# Enter Project Directory

```bash
cd bath_to_text
```

---

# Install Dependencies

```bash
go mod tidy
```

---

# Run the Project

```bash
go run main.go
```

---
`


