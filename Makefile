day1:
	echo -e "Chronal Calibration"
	go build -o ./chronal-calibration/app chronal-calibration/*.go
	cat chronal-calibration/input.txt | chronal-calibration/app

ims:
	go build -o ./inventory-management-system/app inventory-management-system/*.go

day3:
	echo -e "No Matter How You Slice It"
	go build -o ./no-matter-how-you-slice-it/app no-matter-how-you-slice-it/*.go
	no-matter-how-you-slice-it/app -p no-matter-how-you-slice-it/input.txt
