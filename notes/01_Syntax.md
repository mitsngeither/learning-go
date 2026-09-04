- biến viết hoa chữ cái đầu là biến có thể được import(tức là export biến này = cách viết hoa chữ đầu)
- := là nội suy, lấy kiểu dữ liệu của value (bên phải) cho biến (bên trái) và không chi tiết
  ex: k := 3 là int, k có uintN, intN như int8
- int8: 8 bit, vân vân,các int chạy từ 0
- uint: k dấu, chạy từ số âm đến số dương
- byte: alias cho uint8
- rune: alias cho int32, đại diện cho Unicode code point
- chú ý khi ép kiểu type dài về type ngắn hơn, ví dụ float64 về float32
- << và >> là dịch bit, cũng là cơ chế PHÉP NHÂN NHANH NHẤT

* X << Y tương dương X nhân 2 mũ Y (ex 2 << 2 là 2 nhân 2 mũ 2 = 2 x 4 = 8)
* X >> Y tương tự nhưng thay nhân = chia nhưng lấy phần nguyên?(ex 2 >> 3 là 2 chia 2 mũ 3 = lấy nguyên của 0.25 = 0)

- Go k có while, chạy theo kiểu chỉ gán điều kiện stop
- if trong Go có short statement, miễn tận cùng của dòng khai if là điều kiện là ok; switch case cũng kiểu vậy
  ex:

  ```go
  if v := math.Pow(x, n); v < lim {
      return v
  }
  ```

- switch case trong go không cần break

# defer: chế lập lịch thực thi đồng bộ

- chạy ngay trước khi hàm kết thúc, dù có được đặt trước
- stacking defer: first in last out

# anonymous func như js, viết ra xong chạy luôn

# Pointer: Go có con trỏ tường minh

- Khi nào được gọi mới là tham chiếu (tức dùng \*), còn không luôn luôn là tham trị

```go
    p:= &i //p là con trỏ trỏ đến vùng nhớ của i
    fmt.Println(*p) //đọc giá trị mà p đang trỏ tới (tức giá trị của i)
    *p = 22 //set giá trị mà p đang trỏ tới = 22 (tức set i = 22)
    fmt.Println(i) //i nhận value mới
```

# Struct: Go không có OOP, chỉ có Struct, đây là 1 type

