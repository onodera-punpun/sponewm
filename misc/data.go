package misc

var (
	DecorTopAPng    []byte
	DecorTopIPng    []byte
	DecorBottomAPng []byte
	DecorBottomIPng []byte
	DecorLeftAPng   []byte
	DecorLeftIPng   []byte
	DecorRightAPng  []byte
	DecorRightIPng  []byte
)

func ReadData() {
	DecorTopAPng = DataFile("active_top.png")
	DecorTopIPng = DataFile("inactive_top.png")
	DecorBottomAPng = DataFile("active_bottom.png")
	DecorBottomIPng = DataFile("inactive_bottom.png")
	DecorLeftAPng = DataFile("active_left.png")
	DecorLeftIPng = DataFile("inactive_left.png")
	DecorRightAPng = DataFile("active_right.png")
	DecorRightIPng = DataFile("inactive_right.png")
}
