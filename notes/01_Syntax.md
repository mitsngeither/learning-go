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

# Array: Array của Go luôn phải khai báo số phần tử trong mảng - kích thước cố định không đổi được

# Slice: Giống Array nhưng không cần khai báo số phần tử, chỉ là array con của array cho trước, có con trỏ tới array gốc -> đổi slice thì đổi luôn array gốc

- Nên làm kiểu slice trỏ tới chính nó, thay vì trỏ tới array gốc (slice default)
- Slice có thể khai báo capacity(sức chứa) > len
- len() là số ptu hiện có trong slice, cap() là sức chứa tối đa của slice
- Slice mặc định chưa khai báo gì thì là nil
- Cap trong slice là cơ số 2, dù len lúc đó có thể là 5, cap tự tăng nên thường k cần quản lý
- range là 1 for each loop của slice, hay đi với for
- Khi khai báo biến trong Go LÀ PHẢI DÙNG, nếu k dùng thì dùng \_

# Map: như mấy thằng kia, có 1 số API riêng

# Có thể dùng func như 1 biến bình thường trong hàm

# Go không có OOP, không có class nhưng vẫn khai báo được method

- Là cách để giới hạn scope hoạt động của hàm

# Interface: là 1 tập hợp các định nghĩa method
