day1:
	echo -e "Chronal Calibration"
	go build -o ./chronal-calibration/app chronal-calibration/*.go
	cat chronal-calibration/input.txt | chronal-calibration/app

ims:
	go build -o ./inventory-management-system/app inventory-management-system/*.go