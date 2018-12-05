chronal-calibration:
	echo -e "Chronal Calibration"
	go build -o ./chronal-calibration/app chronal-calibration/*.go
	chronal-calibration/app -p chronal-calibration/input.txt && rm chronal-calibration/app && echo

ims:
	echo -e "Inventory Management System"
	go build -o ./inventory-management-system/app inventory-management-system/*.go
	inventory-management-system/app -p inventory-management-system/input.txt && rm inventory-management-system/app && echo

day3:
	echo -e "No Matter How You Slice It"
	go build -o ./no-matter-how-you-slice-it/app no-matter-how-you-slice-it/*.go
	no-matter-how-you-slice-it/app -p no-matter-how-you-slice-it/input.txt

repose-record:
	echo "Repose Record"
	go build -o ./repose-record/app repose-record/*.go
	repose-record/app -p repose-record/input.txt && rm repose-record/app && echo

.PHONY: chronal-calibration ims repose-record