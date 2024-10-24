package builder

import (
	"fmt"

	internal "github.com/hieupc05/goDesignPattern/builder/exp1/internal"
)

func main() {
	normal := internal.NewIBuilder("normal")

	director := internal.NewDirector(normal)

	normalHouse := director.BuildHouse()

	fmt.Println("window: ", normalHouse.GetWindowType())
	fmt.Println("door: ", normalHouse.GetDoorType())
	fmt.Println("numOfFloor: ", normalHouse.GetFloor())

	igloo := internal.NewIBuilder("igloo")
	director = internal.SetBuilder(igloo)
	IglooHouse := director.BuildHouse()
	fmt.Println("Igloo House", IglooHouse)

}

/*
## Hướng xây dựng bao gồm
+ Product -> Sản phẩm cuối
+ Interface -> Phương thức struct tuân theo
+ Concerate builder 1 -> Mẫu builder 1
+ concerate builder 2 -> Mẫu builder 2
...
+ Director : thằng chỉ định gắn interface với struct, nắm giữ các hàm biến mẫu thành product
*/
