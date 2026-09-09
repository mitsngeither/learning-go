## Viết User Story

- Focus on "noun" in user story. Ex:

* As an `user`, I can browse all `restaurants`.
* As an `user`, I can browse all `foods` in `restaurant`.

- So we have: user, restaurant & food are entities
- Entity name should be in singular form
- Table name should be in plural form
- High priority on important/core features

- Entity properties:

* Some of them can be found on list UI
* More on details screens

- Not all of element on UI are properties. They could be computed fields.
- Each of them should have ID, status, created_at, updated_at.

- Relationship between entities - cần bảng trung gian
- Primary key đảm bảo dữ liệu k trùng, ảnh hưởng vị trí vật lý của dữ liệu trong DB
- đánh index cột nào thì tìm kiếm trên cột đó sẽ nhanh hơn nhiều
- bản chất index là chỉ mục - k tac1 động vị trí vật lý của DB
- index tạo ra 1 danh sách vị trí, có vùng lưu trữ riêng indexStorage - có tradeoff, DB tốn dung lượng hơn, nếu quá nhiều loại index + bảng to -> ảnh hưởng tốc độ thay đổi data như insert/delete/update vì luôn phải xử lý lại chỗ index
- lưu ý cơ chế check query trên cột chưa đánh index

```sql
SELECT
    start_time,
    query_time,
    rows_sent,
    rows_examined,
    CONVERT(sql_text USING utf8mb4) AS query_content
FROM mysql.slow_log
ORDER BY start_time DESC;
```

- nếu có 1 cặp key dạng x_id_and_y_id (tức khóa đánh trên multi columns) thì khi queue phải theo thứ tự key từ trái sang phải như dưới thì mới nhanh được, nếu chỉ có y_id thì k có tác dụng - lưu ý để tránh phải đánh nhiều index :

```sql
SELECT * FROM notes WHERE x_id = 1 //(and y_id = 1)
```

- mặc định primary key LÀ index VẬT LÝ
- không dùng foreign key nữa nếu hệ thống k cần consistant data - là 1 khóa ngoại - dùng để tham chiếu qua, bản chất khi insert/update/delete thì luôn check bảng nó tham chiếu/bị tham chiếu tới
- microservice thì k cần dùng foreign key nữa - bỏ dần => join bảng KHÔNG cần foreign key
- thường đánh index value dạng số/ số value đếm được

```sql
SELECT uln.*, n.title FROM user_like_notes uln INNER JOIN  notes n ON n.id = uln-note_id;
```
