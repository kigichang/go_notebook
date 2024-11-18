---
export_on_save:
  markdown: true
markdown:
  image_dir: assets
  path: README.md
  ignore_from_front_matter: true
  absolute_image_path: false
---
# Go 基礎

[The Go Programming Language Specification](https://go.dev/ref/spec)

## 1. Built-in

### 1.1 Keywords

Go 有以下關鍵字，不能用來當作變數名稱。後續會提到如何使用這些關鍵字。

```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

see [Go Keywords](https://golang.org/ref/spec#Keywords)

### 1.2 Functions

Go 內建的函式，後續的內容會提到如何使用這些內建函數。

- `append`: appending slices
- `copy`: copying slices
- `clear`: fill zero values into variable
- `close`: close channel
- `complex`: construct complex number
- `real`: get real part of complex number
- `imag`: get imaginary part of complex number
- `delete`: delete element from map
- `len`: length of array, slice, string, map, channel
- `cap`: capacity of slice, channel
- `make`: create slice, map, channel
- `new`: allocate memory
- `max`: maximum value of two values
- `min`: minimum value of two values
- `panic`: stop normal execution
- `recover`: recover from panic
- `print`: print to standard output
- `println`: print to standard output with newline

see [Go Built-in Functions](https://go.dev/ref/spec#Built-in_functions)

## 2. Primitive Types

## 2.1 Boolean Types

只有 `true` 及 `false`，且不能像 C 數字或 null 來當 boolean 使用。

## 2.2 Numeric Types

| Type       | Description                                              |
| ---------- | -------------------------------------------------------- |
| uint8      | 無符號 8-bit 整數 (0 to 255)                               |
| uint16     | 無符號 16-bit 整數 (0 to 65535)                            |
| uint32     | 無符號 32-bit 整數 (0 to 4294967295)                       |
| uint64     | 無符號 64-bit 整數 (0 to 18446744073709551615)             |
| int8       | 8-bit 整數 (-128 to 127)                                  |
| int16      | 16-bit 整數 (-32768 to 32767)                             |
| int32      | 32-bit 整數 (-2147483648 to 2147483647)                   |
| int64      | 64-bit 整數 (-9223372036854775808 to 9223372036854775807) |
| float32    | 32-bit 浮點數                                             |
| float64    | 64-bit 浮點數                                             |
| complex64  | float32 複數                                              |
| complex128 | float64 複數                                              |
| byte       | 位元組，uint8 別名                                         |
| rune       | Unicode 字元，int32 別名                                   |
| uint       | 無符號整數，平台相依，uint32 或 uint64                       |
| int        | 整數，平台相依，可能是 int32 或 int64                        |
| uintptr    | 指標位址                                                  |

## 2.3 變數宣告

### 2.3.1 完整宣告

```go
var name type = expression
```

```go
var a int = 0
```

1. 宣告變數，不給初始值，可以省略 `= expression`

    ```go
    var a int
    ```

1. 宣告變數，給初始值時，則可以省略資料型別

    ```go
    var a = 10
    ```

常用的方式如下：

1. 宣告變數，不給初始值

    ```go
    var s string
    println(s) // ""
    ```

1. 一次宣告多個變數，並給初始值(可省略型別)

    ```go
    var i, j, k int                 // int, int, int
    var b, f, s = true, 2.3, "four" // bool, float64, string
    ```

1. 如果函式有回傳多組值時

    ```go
    var f, err = os.Open(name) // os.Open returns a file and an error
    ```

1. 檢查 map 是否存在指定的 key

    ```go
    var val, ok = ages["bob"]
    ```

1. 檢查是否可以轉換型別

    ```go {.line-numbers}
    var guessType, ok = a.(guessType)
    ```

### 2.3.2 使用 `:=` 簡寫

使用 `:=` 省略 `var`

```go {.line-numbers}
name := expression
```

使用方式如下：

1. 省略型別宣告

    ```go {.line-numbers}
    i, j := 0, 1
    ```

1. 接收函式回傳值

    ```go {.line-numbers}
    anim := gif.GIF{LoopCount: nframes}
    freq := rand.Float64() * 3.0
    t := 0.0

    f, err := os.Open(name)
    if err != nil {
    return err
    }

    // ...use f...

    f.Close()

    val, ok := ages["bob"]
    guessType, ok := a.(guessType)
    ```

使用 `:=` 時，左邊的變數名稱，至少要有一個是新的。

1. 至少要有一個是新的變數名稱

    ```go
    in, err := os.Open(infile)
    // ...
    out, err := os.Create(outfile)
    ```

    以上，雖然 `err` 重覆，但 `out` 是新的變數名稱，compile 會過。

1. 都是舊的

    ```go
    f, err := os.Open(infile)
    // ...
    f, err := os.Create(outfile) // compile error: no new variables
    ```

    以上，`f` 與 `err` 都是舊的變數，所以在第二次，還是使用 `:=` 時，compile 會錯。通常 compile 會報錯，都不是什麼大問題，修正就好了。

### 2.4 預設型別

當省略型別時，Go 的編譯器會自動帶入預設型別。

- 整數: __int__
- 浮點數: __float64__
- 複數型別: __complex128__

### 2.5 數值宣告

以往宣告很大數值時，無法像 excel 每千分位，用 `,` 來區隔。現在 Go 也支援這項功能，可以使用 `_` 來區隔。

```go
var x int64 = 123_456_789
var y float64 = 12_345.678_9
```

### 2.6 Zero Value

每一種資料型別在宣告時，沒有給定值的話，則 Go 會給予一個初始值，這個初始值則稱為該型別的 __Zero Value__。

| Type           | Zero Value |
| -------------- | ---------- |
| Boolean        | false      |
| Integer        | 0          |
| Floating Point | 0.0        |
| string         | ""         |
| pointer        | nil        |
| interface      | nil        |
| map            | nil        |
| slice          | nil        |
| struct         | 依每一個 field 的資料型別，給定對應的 zero value. |

### 2.7 複數

GO 的複數

- complex64: 由兩個 float32 組成
- complex128: 由兩個 float64 組成

1. 複數宣告

    @import "declaration/main.go" {line_begin=28 line_end=30}

1. 使用 `complex` function 宣告， `real` 與 `imag` function

    @import "declaration/main.go" {line_begin=33 line_end=38}

### 2.8 Constant and itoa

與 C 相同，利用 `const` 這個關鍵字來宣告常數。

@import "const_iota/main.go" {line_begin=2 line_end=6}

可以使用 `iota` 定義連續的常數，`iota` 是一個特殊的常數，從 0 開始，每次宣告時，會自動遞增。
可以利用 __\___ 萬用字元，不用宣告變數，只是為了讓 `iota` 遞增。

@import "const_iota/main.go" {line_begin=7 line_end=21}

see [Go Wiki: Iota](https://go.dev/wiki/Iota)

## 3. Functions

### 3.1 宣告

使用 `func` 關鍵字來宣告 function。

```go
func name(parameter-list) (result-list) {
    body
}
```

如:

```go
func hypot(x float64, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}

println(hypot(3, 4)) // "5"
```

### 3.2 Grouping 相同型別

如果參數型別相同，可以將型別寫在最後一個參數後面，簡化宣告。

```go
func f(i int, j int, k int, s string, t string) { /* ... */ } // original
func f(i, j, k int, s, t string)                { /* ... */ } // simplify
```

### 3.3 回傳值

Go 函式的回傳值，也可以像參數命名；如果回傳值有命名，等同在函式內宣告了一個變數，並且給需對應的 Zero Value。

```go
func add(x int, y int) int { return x+y }
func sub(x, y int) (z int) { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int { return 0 }
```

Go 的 function 可以一次回傳多個值。

```go {.line-numbers}
func swap(x, y int) (int, int) {
    return y, x
}

a, b := 1, 2        // a = 1, b = 2
a, b = swap(a, b)   // a = 2, b = 1
```

### 3.4 Variadic Functions

function 的參數個數可以是不固定的。eg:

1. 宣告

    ```go
    func sum(vals ...int) int {
        total := 0
        for _, val := range vals {
            total += val
        }
        return total
    }

    println(sum())           //  "0"
    println(sum(3))          //  "3"
    println(sum(1, 2, 3, 4)) //  "10"
    ```

1. 如何將 slice 傳入:

    ```go
    values := []int{1, 2, 3, 4}
    println(sum(values...)) // "10"
    ```

### 3.5 Recursion 遞迴

```go {.line-numbers}
func gcd(a, b int) int {
    if b == 0 {
        return a
    }

    return gcd(b, a % b)
}

println(gcd(24, 128)) // 8
```

### 3.6 Signature

一個 function 的型別，通常也稱做 __Signature__。兩個 function 有相同的 signature，需滿足以下兩個條件：

1. 參數 (parameters) 資料型別與順序相同，與參數名稱無關。
1. 回傳的值的資料型別與順序相同

eg:

```go
func add(x int, y int) int { return x+y }
func sub(x, y int) (z int) { z= x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int { return 0 }

printf("%T\n", add)   // "func(int, int) int"
printf("%T\n", sub)   // "func(int, int) int"
printf("%T\n", first) // "func(int, int) int"
printf("%T\n", zero)  // "func(int, int) int"
```

### 3.7 First-Class Language

在 Go 的 function 是一種資料型別，可以當作參數與回傳值。
以 Go 來說，__signature__ 是 Function 的資料型別。當宣告 funcation 沒有命名時，則稱為 __anonymous function__。

#### 3.7.1 Assignment

Function 可以當作一個值，assign 給某個變數。

```go {.line-numbers}
func square(n int) int { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }

var f func(int) int     // signature

f = square
println(f(3))       // "9"
f = negative
println(f(3))       // "-3"

f = product // cannot use product (type func(int, int) int) as type func(int) int in assignment
```

### 3.7.2 As Parameter and Return

Function 可以當作參數與回傳值。

```go {.line-numbers}
func square(n int) int { return n * n }
func negative(n int) int { return -n }

func compose(f, g func(int) int) func(int) int {
    return func(a int) int {        // anonymous function
        return g(f(a))
    }
}

k1 := compose(square, negative)
printf("%T\n", k1)              // func(int) int
println(k1(10))                 // -100 negative(square(10))

k2 := compose(negative, square)
printf("%T\n", k2)              // func(int) int
println(k2(10))                 // 100 square(negative(10))
```

## Array and Slice

Slice 與 Array 類似，與 Array 最大不同點是 Slice 的長度是可變動的，但 Array 是固定的。

Slice 組成元素：

- Pointer: 指向內部的 array。
- Length: 目前該 slice 可操作的最大長度。
- Capacity: 實際內部 array 的大小，可視做目前最大的容量。

### 2.1 Slice Declaration (ex05_02)

Slice 的 __zero value__ 是 __nil__。宣告的方式可以是：`[]T` T 是指資料型別, 或用 `make`，指定 Length 及 Capacity。如:

@import "ex05_02/main.go" {class=line-numbers}

### 2.2 Array and Slice Relation (ex05_03)

實際上，Slice 底層有一組 Array 存放資料，Slice 的 Pointer 會指向該 Array。。

```go {.line-numbers}
months := [...]string{1: "January", /* ... */, 12: "December"}
Q2 := months[4:7]
summer := months[6:9]
println(Q2)     // ["April" "May" "June"]
println(summer) // ["June" "July" "August"]
```

```ditaa
                         +-------------+ 
Q2 = Months[4:7]         | ""          |               summer = months[6:9]
+--------+               +-------------+               +--------+
| Data   |---+           | "January"   |           +---| Data   |
| len: 3 |   |           +-------------+           |   | len: 3 |
| cap: 9 |   |           | "February"  |           |   | cap: 9 |
+--------+   |           +-------------+           |   +--------+
             |           | "March"     |           | 
             +-------+-> +-------------+           |
             |       |   | "April"     |           |
             |       |   +-------------+           | 
             |       |   | "May"       |           | 
             | len=3 |   +-------------+ <-+-------+
             |       |   | "June"      |   |       |
             |       +-- +-------------+   |       |
             |           | "July"      |   | len=3 |
       cap=9 |           +-------------+   |       | 
             |           | "August"    |   |       | cap=7
             |           +-------------+ --+       |
             |           | "September" |           |
             |           +-------------+           |
             |           | "October"   |           |
             |           +-------------+           |
             |           | "November"  |           |
             |           +-------------+           |
             |           | "December"  |           |
             +---------- +-------------+ ----------+
```

1. `Q2`是取 `April`, `May`, `June` 值，Pointer 會指到 `April` 為開啟。 Q2 只取三個月份，因此 Length 是 __3__；全部資料有 12 筆，Q2 是從 `April` 開始，捨棄前面 4 筆資料，因此 Capacity 為 __9___ (12 - 4 = 9)。

由於 Array, Struct __都不是__ Reference Types，因此在傳入 Function 時，都會 clone 一份新的資料，給 Function 使用，也因此如果 array/struct 的資料很龐大時，就會造成記憶體上的浪費。因此在設計上，Function 的參數有 array 時，可以改用 slice, struct 請用 pointer。

由於 slice 是用 pointer 指到 array, 因此修改 slice 的值時，也會異動到原本的 array.

@import "ex05_03/main.go" {class=line-numbers}

### 2.3 Slice Append (ex05_04)

可以使用 `append` 新增資料進 slice

@import "ex05_04/main.go" {class=line-numbers}

在上述範例中，

1. `s1` 已經沒有空間做 `append`，因此產生了一組新的記憶體空間，也因為這樣，才沒有更動到 `s`。
1. 但 `s2` 還有空間做 `append`, 可以用原來的位址來操作，因此會修改到原來的 `s`。[^append]

在實作上，儘可能利用 __slice__ 而非 array。

1. 可避免因 pass by value，而造成記憶體的浪費。
1. 避免上述 puzzle。

[^append]: 在進行 append 時，會先檢查 capacity 是否有足夠空間，來加入新的資料，如果沒有時，則會再產生一組新的記憶體空間，先將舊的資料，__copy__ 進新的空間，再把新的資料加入。也因此，如果要大量 append 資料時，應該先計算好可能的容量大小，以免一直在做 copy 的動作，影響效能。

### 2.4 Slice Travel

與 array 同，用 `for-range`

```go {.line-numbers}
a := []int{1, 2, 3}
for i := range a {
    println(a[i])
}

for i, v := range a {
    println(i, v)
}

for _, v := range a {
    println(v)
}
```

## string and rune

### []byte to string and string to []byte

### number to string and string to number

## Map

Key-Value 結構，也就是 hashtable 的結構。

### 3.1 Map Declaration

```go {.line-numbers}
ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
```

也可使用 `make` 來產生 空白 map.

```go {.line-numbers}
ages := make(map[string]int) // mapping from strings to ints
```

### 3.2 Put

```go {.line-numbers}
ages["alice"] = 32      // alice = 32
ages["alice"]++         // alice = 33
```

### 3.3 Delete

```go {.line-numbers}
delete(ages, "cat")
```

### 3.5 Get

Map 在取值時，如果 key 不存在，會回值 value 型別的 __zero value__，也因此無法直接從回傳值來判斷該 key 是否存在。可以利用 `value, ok := map[key]` 的方式，透過驗証 `ok` 來判斷 key 是否存在。

```go {.line-numbers}
println(ages["bob"])    // 0 (zero-value)

a, ok := ages["bob"]
println(a, ok)          // 0, false
```

### 3.6 Map Travel

與 array 同，用 `for-range`

```go {.line-numbers}
m := map[int]string{
    1: "a",
    2: "b",
}

for key := range m {
    println(key, m[key])
}

for key, val := range m {
    println(key, val)
}

for _, val := range m {
    println(val)
}
```


## Flow Control

### if

### for

### switch

## 10. Go Wildcard

Go 的 wildcard 是 `_`，可以用在以下情境：

1. 因為 Go compiler 會檢查沒有使用的變數，如果不想使用該數值時，可以使用 `_` 來取代。

```go {.line-numbers}
_ = test()
```

1. 在宣告函式時，有些參數確定不會被用到，可以在宣告時使用 `_`。

```go {.line-numbers}
func mytest(_ int, str string) {

}
```
