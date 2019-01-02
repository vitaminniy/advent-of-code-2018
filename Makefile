chronal-calibration:
	echo -e "Chronal Calibration"
	go build -o ./chronal-calibration/app chronal-calibration/*.go
	chronal-calibration/app -p chronal-calibration/input.txt && rm chronal-calibration/app && echo

ims:
	echo "Inventory Management System"
	go build -o ./inventory-management-system/app inventory-management-system/*.go
	inventory-management-system/app -p inventory-management-system/input.txt && rm inventory-management-system/app && echo

day3:
	echo "No Matter How You Slice It"
	go build -o ./no-matter-how-you-slice-it/app no-matter-how-you-slice-it/*.go
	no-matter-how-you-slice-it/app -p no-matter-how-you-slice-it/input.txt && rm no-matter-how-you-slice-it/app && echo

repose-record:
	echo "Repose Record"
	go build -o ./repose-record/app repose-record/*.go
	repose-record/app -p repose-record/input.txt && rm repose-record/app && echo

alchemical-reduction:
	echo "Alchemical Reduction"
	go build -o ./alchemical-reduction/app alchemical-reduction/*.go
	alchemical-reduction/app -p alchemical-reduction/input.txt && rm alchemical-reduction/app && echo

chronal-coordinates:
	echo "Chronal Coordinates"
	go build -o ./chronal-coordinates/app chronal-coordinates/*.go
	chronal-coordinates/app -p chronal-coordinates/input.txt && rm chronal-coordinates/app && echo

.PHONY: chronal-calibration ims day3 repose-record alchemical-reduction chronal-coordinates