# Go 基硫

## 1 Build-in

## 2. Basic Types

### Declartion

### Zero Value

### Costant, iota, and Enum

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
fmt.Println(Q2)     // ["April" "May" "June"]
fmt.Println(summer) // ["June" "July" "August"]
```

![Slice](slice.png)

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
    fmt.Println(a[i])
}

for i, v := range a {
    fmt.Println(i, v)
}

for _, v := range a {
    fmt.Println(v)
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
fmt.Println(ages["bob"])    // 0 (zero-value)

a, ok := ages["bob"]
fmt.Println(a, ok)          // 0, false
```

### 3.6 Map Travel

與 array 同，用 `for-range`

```go {.line-numbers}
m := map[int]string{
    1: "a",
    2: "b",
}

for key := range m {
    fmt.Println(key, m[key])
}

for key, val := range m {
    fmt.Println(key, val)
}

for _, val := range m {
    fmt.Println(val)
}
```


## Functions

## 1. 宣告

```go {.line-numbers}
func name(parameter-list) (result-list) {
    body
}
```

```go {.line-numbers}
func hypot(x float64, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}

fmt.Println(hypot(3, 4)) // "5"
```

### 1.1 Grouping 相同型別

```go {.line-numbers}
func f(i int, j int, k int, s string, t string) { /* ... */ } // original
func f(i, j, k int, s, t string)                { /* ... */ } // simplify
```

### 1.2 回傳值

```go {.line-numbers}
func add(x int, y int) int { return x+y }
func sub(x, y int) (z int) { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int { return 0 }
```

Go 的 function 可以一次回傳多個值 (tuple)

```go {.line-numbers}
func swap(x, y int) (int, int) {
    return y, x
}

a, b := 1, 2        // a = 1, b = 2
a, b = swap(a, b)   // a = 2, b = 1
```

### 1.3 Variadic Functions

function 的參數個數可以是不固定的。eg:

1. 宣告

    ```go {.line-numbers}
    func sum(vals ...int) int {
        total := 0
        for _, val := range vals {
            total += val
        }
        return total
    }

    fmt.Println(sum())           //  "0"
    fmt.Println(sum(3))          //  "3"
    fmt.Println(sum(1, 2, 3, 4)) //  "10"
    ```

1. 如何將 slice 傳入:

    ```go {.line-numbers}
    values := []int{1, 2, 3, 4}
    fmt.Println(sum(values...)) // "10"
    ```

### 1.4 空白 Body

可以定義 function 但沒有 body。通常是用另一種程式語言來實作，比如 C or Javascript in WASM。越是底層的工作越容易看到這樣子的做法。

## 2. Recursion 遞迴

```go {.line-numbers}
func gcd(a, b int) int {
    if b == 0 {
        return a
    }

    return gcd(b, a % b)
}

fmt.Println(gcd(24, 128)) // 8
```

## 3. Pass by Value (Call by Value)

Go 在傳遞參數時，是以 **by value** 的方式進行，也就是說在傳入 function 前，會產生一份新的資料，給 function 使用，也因此 function 修改時，也是修改此新的資料。

此時要特別注意傳入的資料型別：

- Aggregate Types (Array, Struct)，在 Java 的定義下，是屬於 Value Types，也就是說會產生一筆新的資料給 function，function 做任何修改，都**不會**異動到原本的資料，如果 array/struct 資料很龐大時，會造成記憶體的浪費。
- Reference Types (Pointer, Slice, Map, Function, Channel)，一樣在傳入 function 時，會複製新的值給 function，只是這新的值，只是 copy 原本的參照值(reference, 可以當作記憶體位址)，因此 function 做任何修改時，也都是透過原來的參照值在做資料異動，會修改到原本的資料，要特別小心。

### 3.1 Pass by Value with Struct and Struct Pointer (ex06_01)

@import "ex06_01/main.go" {class=line-numbers}

### 3.2 Pass by Value with Aray and Slice (ex06_02)

@import "ex06_02/main.go" {class=line-numbers}

## 4. Signature

一個 function 的型別，通常也稱做 **Signature**。兩個 function 有相同的 signature，需滿足以下兩個條件：

1. 參數 (parameters) 資料型別與順序相同，與參數名稱無關。
1. 回傳的值的資料型別與順序相同

eg:

```go {.line-numbers}
func add(x int, y int) int { return x+y }
func sub(x, y int) (z int) { z= x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int { return 0 }

fmt.Printf("%T\n", add)   // "func(int, int) int"
fmt.Printf("%T\n", sub)   // "func(int, int) int"
fmt.Printf("%T\n", first) // "func(int, int) int"
fmt.Printf("%T\n", zero)  // "func(int, int) int"
```

在 Go 的 function 也可以當作參數與回傳值。也因此 Go 也算是一種 first-class lanaugage.

### 4.1 First-Class (ex06_04)

function 也是一種資料型別，可以當作變數，或當作另一個 function 的參數及回傳值。
以 Go 來說，**signature** 是 Function 的資料型別。當宣告 funcation 沒有指定 name 時，則稱為 **anonymous function**

### 4.2 Assignment

Function 可以當作一個值，assign 給某個變數。

```go {.line-numbers}
func square(n int) int { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }

var f func(int) int     // signature
fmt.Printf("%T\n", f)   // "func(int) int"

f = square
fmt.Println(f(3))       // "9"

f = negative
fmt.Println(f(3))       // "-3"

f = product // cannot use product (type func(int, int) int) as type func(int) int in assignment
```

### 4.3 As Parameter and Return

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
fmt.Printf("%T\n", k1)              // func(int) int
fmt.Println(k1(10))                 // -100 negative(square(10))

k2 := compose(negative, square)
fmt.Printf("%T\n", k2)              // func(int) int
fmt.Println(k2(10))                 // 100 square(negative(10))
```


### Variadic Function

### First-Class Function

## Flow Control

### if

### for

### switch

## Go Wildcard
