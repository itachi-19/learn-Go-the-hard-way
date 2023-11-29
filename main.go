package main

//import "fmt"
import "reflect"
//import "fmt"

//Reverse reverses a slice.
// var Reverse = func(slice interface{}) {
//     arr := reflect.ValueOf(slice)
//     if arr.Kind() == reflect.Pointer {
//         s := arr.Elem()
//         swapper := reflect.Swapper(s.Interface())
//         for i, j := 0, s.Len() - 1; i < j; {
//             swapper(i, j)
//             i++
//             j--
//             // fmt.Println(s.Index(i), s.Index(j))
//         }
//     }
// }

var Reverse = func(slice interface{}) {
    arr := reflect.ValueOf(slice)
    if arr.Kind() == reflect.Pointer {
        s := arr.Elem()
        for i, j := 0, s.Len() - 1; i < j; {
            x, y := s.Index(i).Interface(), s.Index(j).Interface()
            s.Index(i).Set(reflect.ValueOf(y))
            s.Index(j).Set(reflect.ValueOf(x))
            i++
            j--
        }
    }
}

func main() {
	println("Please edit main.go,and complete the 'Reverse' function to pass the test.\nYou should use reflect package to reflect the slice type and make it applly to any type.\nTo run test,please run 'go test'\nIf you pass the test,please run 'git checkout l2' ")
}
