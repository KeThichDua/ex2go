## cơ bản 2 context và time

- Một kiến thức cực quan trọng trong golang là time và context. Vì ý nghĩa lớn của nó trong thực hiện các action. Việc tách time và context ra là bài số 2 là điều cực quan trọng.
- Nắm vững kiến thức của time và context là tiền đề cho việc đi tiếp những bài tiếp theo.
- Nội dung được biên soạn bởi [maTe](fb.com/0.anhsang.0)

1. Tạo 1 chương trình cứ 3s in ra dòng chữ: `time now: {milliseconds}`:
Sau khi in được 3 lần thì in ra dòng chữ `kết thúc`

2. Việt 1 đoạn chương trình tính ra ngày hiện tại theo timestamp.
Điểm mốc từ mức 0  tại 1/1/1970

3. Thực hiện 1 chương trình với 1 vòng lặp for và 3 lần sleep mỗi lần sleep 3sec
Nhưng sau 3s thì kết thúc hàm đấy
Sử dụng và tìm hiểu context. Nêu được tác dụng của context trong chương trình.

4. Tính ra số phút(number_of_minutes) của mốc thời gian sau `1592190294764144364`


5. Tính ra số ngày trong tuần(dạng string và number) của mốc thời gian sau `1592190385`

6. Trong golang mặc định thì thời gian dạng số được sử dụng với các loại mốc đơn vị nào?

7. Tạo 1 đối tượng context với 1 value là timestamp hiện tại tính theo `ns`
chạy qua 1 hàm như sau. Sau 3s in ra hiệu của thời gian của thời gian hiện tại với biến dữ liệu trong context. in ra màn hình kết quả.
```go
func x(ctx context.Context) {

}
```
goi y sư dụng một method của time để tính tạo ra 1 tính năng như setTimeout trong js



8. Tạo 1 interval time sao cho cứ 100ms in ra dùng chữ `${time.Now().Unix()} done`

9. Tạo 1 đoạn code sử dụng `time.AfterFunc()`, sau 100ms thì in ra dòng chữ `i'm study`
